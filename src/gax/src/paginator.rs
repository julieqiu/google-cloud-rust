// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

use futures::stream::unfold;
use futures::{Stream, StreamExt};
use pin_project::pin_project;
use std::future::Future;
use std::pin::Pin;

/// This module contains implementation details. It is not part of the public
/// API. Types inside may be changed or removed without warnings. Applications
/// should not use any types contained within.
#[doc(hidden)]
pub mod internal {
    use super::*;

    /// Describes a type that can be iterated over asynchronously when used with
    /// [super::Paginator].
    pub trait PageableResponse {
        type PageItem: Send;

        // Consumes the [PageableResponse] and returns the items associated with the
        // current page.
        fn items(self) -> Vec<Self::PageItem>;

        /// Returns the next page token.
        fn next_page_token(&self) -> String;
    }

    /// Creates a new `impl Paginator<T, E>` given the initial page token and a function
    /// to fetch the next response.
    pub fn new_paginator<T, E, F>(
        seed_token: String,
        execute: impl Fn(String) -> F + Clone + Send + 'static,
    ) -> impl Paginator<T, E>
    where
        T: internal::PageableResponse,
        F: Future<Output = Result<T, E>> + Send + 'static,
    {
        PaginatorImpl::new(seed_token, execute)
    }
}

mod sealed {
    pub trait Paginator {}
}

/// An adapter that converts list RPCs as defined by [AIP-4233](https://google.aip.dev/client-libraries/4233)
/// into a [futures::Stream] that can be iterated over in an async fashion.
pub trait Paginator<T, E>: Send + sealed::Paginator
where
    T: internal::PageableResponse,
{
    /// Creates a new [ItemPaginator] from an existing [Paginator].
    fn items(self) -> impl ItemPaginator<T, E>;

    /// Returns the next mutation of the wrapped stream.
    fn next(&mut self) -> impl Future<Output = Option<Result<T, E>>> + Send;

    #[cfg(feature = "unstable-stream")]
    /// Convert the paginator to a stream.
    ///
    /// This API is gated by the `unstable-stream` feature.
    fn into_stream(self) -> impl futures::Stream<Item = Result<T, E>> + Unpin;
}

#[pin_project]
struct PaginatorImpl<T, E> {
    #[pin]
    stream: Pin<Box<dyn Stream<Item = Result<T, E>> + Send>>,
}

type ControlFlow = std::ops::ControlFlow<(), String>;

impl<T, E> PaginatorImpl<T, E>
where
    T: internal::PageableResponse,
{
    /// Creates a new [Paginator] given the initial page token and a function
    /// to fetch the next response.
    pub fn new<F>(
        seed_token: String,
        execute: impl Fn(String) -> F + Clone + Send + 'static,
    ) -> Self
    where
        F: Future<Output = Result<T, E>> + Send + 'static,
    {
        let stream = unfold(ControlFlow::Continue(seed_token), move |state| {
            let execute = execute.clone();
            async move {
                let token = match state {
                    ControlFlow::Continue(token) => token,
                    ControlFlow::Break(_) => return None,
                };
                match execute(token).await {
                    Ok(page_resp) => {
                        let tok = page_resp.next_page_token();
                        let next_state = if tok.is_empty() {
                            ControlFlow::Break(())
                        } else {
                            ControlFlow::Continue(tok)
                        };
                        Some((Ok(page_resp), next_state))
                    }
                    Err(e) => Some((Err(e), ControlFlow::Break(()))),
                }
            }
        });
        Self {
            stream: Box::pin(stream),
        }
    }
}

impl<T, E> Paginator<T, E> for PaginatorImpl<T, E>
where
    T: internal::PageableResponse,
{
    /// Creates a new [ItemPaginator] from an existing [Paginator].
    fn items(self) -> impl ItemPaginator<T, E> {
        ItemPaginatorImpl::new(self)
    }

    /// Returns the next mutation of the wrapped stream.
    async fn next(&mut self) -> Option<Result<T, E>> {
        self.stream.next().await
    }

    #[cfg(feature = "unstable-stream")]
    /// Convert the paginator to a stream.
    ///
    /// This API is gated by the `unstable-stream` feature.
    fn into_stream(self) -> impl futures::Stream<Item = Result<T, E>> + Unpin {
        Box::pin(unfold(Some(self), move |state| async move {
            if let Some(mut paginator) = state {
                if let Some(pr) = paginator.next().await {
                    return Some((pr, Some(paginator)));
                }
            };
            None
        }))
    }
}

impl<T, E> sealed::Paginator for PaginatorImpl<T, E> where T: internal::PageableResponse {}

impl<T, E> std::fmt::Debug for PaginatorImpl<T, E> {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("Paginator").finish()
    }
}

pub trait ItemPaginator<T, E>: Send + sealed::Paginator
where
    T: internal::PageableResponse,
{
    /// Returns the next mutation of the wrapped stream.
    ///
    /// Enable the `unstable-stream` feature to interact with a [`futures::stream::Stream`].
    ///
    /// [`futures::stream::Stream`]: https://docs.rs/futures/latest/futures/stream/trait.Stream.html
    fn next(&mut self) -> impl Future<Output = Option<Result<T::PageItem, E>>> + Send;

    #[cfg(feature = "unstable-stream")]
    /// Convert the paginator to a stream.
    ///
    /// This API is gated by the `unstable-stream` feature.
    fn into_stream(self) -> impl futures::Stream<Item = Result<T::PageItem, E>> + Unpin;
}

/// An adapter that converts a [Paginator] into a stream of individual page
/// items.
#[pin_project]
struct ItemPaginatorImpl<T, E>
where
    T: internal::PageableResponse,
{
    #[pin]
    stream: PaginatorImpl<T, E>,
    current_items: Option<std::vec::IntoIter<T::PageItem>>,
}

impl<T, E> ItemPaginatorImpl<T, E>
where
    T: internal::PageableResponse,
{
    /// Creates a new [ItemPaginator] from an existing [Paginator].
    fn new(paginator: PaginatorImpl<T, E>) -> Self {
        Self {
            stream: paginator,
            current_items: None,
        }
    }
}

impl<T, E> ItemPaginator<T, E> for ItemPaginatorImpl<T, E>
where
    T: internal::PageableResponse,
{
    /// Returns the next mutation of the wrapped stream.
    ///
    /// Enable the `unstable-stream` feature to interact with a [`futures::stream::Stream`].
    ///
    /// [`futures::stream::Stream`]: https://docs.rs/futures/latest/futures/stream/trait.Stream.html
    async fn next(&mut self) -> Option<Result<T::PageItem, E>> {
        loop {
            if let Some(ref mut iter) = self.current_items {
                if let Some(item) = iter.next() {
                    return Some(Ok(item));
                }
            }

            let next_page = self.stream.next().await;
            match next_page {
                Some(Ok(page)) => {
                    self.current_items = Some(page.items().into_iter());
                }
                Some(Err(e)) => {
                    return Some(Err(e));
                }
                None => return None,
            }
        }
    }

    #[cfg(feature = "unstable-stream")]
    /// Convert the paginator to a stream.
    ///
    /// This API is gated by the `unstable-stream` feature.
    fn into_stream(self) -> impl Stream<Item = Result<T::PageItem, E>> + Unpin {
        Box::pin(unfold(Some(self), move |state| async move {
            if let Some(mut paginator) = state {
                if let Some(pr) = paginator.next().await {
                    return Some((pr, Some(paginator)));
                }
            };
            None
        }))
    }
}

impl<T, E> sealed::Paginator for ItemPaginatorImpl<T, E> where T: internal::PageableResponse {}

#[cfg(test)]
mod tests {
    use super::internal::*;
    use super::*;
    use std::collections::VecDeque;
    use std::sync::{Arc, Mutex};

    #[derive(Clone, Default)]
    struct TestRequest {
        page_token: String,
    }

    struct TestResponse {
        items: Vec<PageItem>,
        next_page_token: String,
    }

    #[derive(Clone)]
    struct PageItem {
        name: String,
    }

    impl PageableResponse for TestResponse {
        type PageItem = PageItem;

        fn items(self) -> Vec<Self::PageItem> {
            self.items
        }

        fn next_page_token(&self) -> String {
            self.next_page_token.clone()
        }
    }

    #[derive(Clone)]
    struct Client {
        inner: Arc<InnerClient>,
    }

    struct InnerClient {
        data: Arc<Mutex<Vec<TestResponse>>>,
    }

    impl Client {
        async fn execute(
            data: Arc<Mutex<Vec<TestResponse>>>,
            _: TestRequest,
        ) -> Result<TestResponse, Box<dyn std::error::Error>> {
            // This is where we could run a request with a client
            let mut responses = data.lock().unwrap();
            let resp: TestResponse = responses.remove(0);
            Ok(resp)
        }

        async fn list_rpc(
            &self,
            req: TestRequest,
        ) -> Result<TestResponse, Box<dyn std::error::Error>> {
            let inner = self.inner.clone();
            Client::execute(inner.data.clone(), req).await
        }

        fn list_rpc_stream(
            &self,
            req: TestRequest,
        ) -> impl Paginator<TestResponse, Box<dyn std::error::Error>> {
            let client = self.clone();
            let tok = req.page_token.clone();
            let execute = move |token| {
                let mut req = req.clone();
                let client = client.clone();
                req.page_token = token;
                async move { client.list_rpc(req).await }
            };
            new_paginator(tok, execute)
        }
    }

    #[tokio::test]
    async fn test_paginator() {
        let seed = "token1".to_string();
        let mut responses = VecDeque::new();
        responses.push_back(TestResponse {
            items: vec![
                PageItem {
                    name: "item1".to_string(),
                },
                PageItem {
                    name: "item2".to_string(),
                },
            ],
            next_page_token: "token2".to_string(),
        });
        responses.push_back(TestResponse {
            items: vec![PageItem {
                name: "item3".to_string(),
            }],
            next_page_token: "".to_string(),
        });
        let mut expected_tokens = VecDeque::new();
        expected_tokens.push_back("token1".to_string());
        expected_tokens.push_back("token2".to_string());

        let state = Arc::new(Mutex::new(responses));
        let tokens = Arc::new(Mutex::new(expected_tokens));

        let execute = move |token: String| {
            let expected_token = tokens.clone().lock().unwrap().pop_front().unwrap();
            assert_eq!(token, expected_token);
            let resp: TestResponse = state.clone().lock().unwrap().pop_front().unwrap();
            async move { Ok::<_, Box<dyn std::error::Error>>(resp) }
        };

        let mut resps = vec![];
        let mut paginator = new_paginator(seed, execute);
        while let Some(resp) = paginator.next().await {
            if let Ok(resp) = resp {
                resps.push(resp)
            }
        }
        assert_eq!(resps.len(), 2);
        assert_eq!(resps[0].items[0].name, "item1");
        assert_eq!(resps[0].items[1].name, "item2");
    }

    #[tokio::test]
    async fn test_paginator_as_client() {
        let responses = vec![
            TestResponse {
                items: vec![
                    PageItem {
                        name: "item1".to_string(),
                    },
                    PageItem {
                        name: "item2".to_string(),
                    },
                ],
                next_page_token: "token1".to_string(),
            },
            TestResponse {
                items: vec![PageItem {
                    name: "item3".to_string(),
                }],
                next_page_token: "".to_string(),
            },
        ];

        let client = Client {
            inner: Arc::new(InnerClient {
                data: Arc::new(Mutex::new(responses)),
            }),
        };
        let mut resps = vec![];
        let mut paginator = client.list_rpc_stream(TestRequest::default());
        while let Some(resp) = paginator.next().await {
            if let Ok(resp) = resp {
                resps.push(resp)
            }
        }
        assert_eq!(resps.len(), 2);
        assert_eq!(resps[0].items[0].name, "item1");
        assert_eq!(resps[0].items[1].name, "item2");
        assert_eq!(resps[1].items[0].name, "item3");
    }

    #[tokio::test]
    async fn test_paginator_error() {
        let execute = |_| async { Err::<TestResponse, Box<dyn std::error::Error>>("err".into()) };

        let mut paginator = new_paginator(String::new(), execute);
        let mut count = 0;
        while let Some(resp) = paginator.next().await {
            match resp {
                Ok(_) => {
                    panic!("Should not succeed");
                }
                Err(e) => {
                    assert_eq!(e.to_string(), "err");
                    count += 1;
                }
            }
        }
        assert_eq!(count, 1);
    }

    #[cfg(feature = "unstable-stream")]
    #[tokio::test]
    async fn test_paginator_into_stream() {
        let responses = vec![
            TestResponse {
                items: vec![
                    PageItem {
                        name: "item1".to_string(),
                    },
                    PageItem {
                        name: "item2".to_string(),
                    },
                ],
                next_page_token: "token1".to_string(),
            },
            TestResponse {
                items: vec![PageItem {
                    name: "item3".to_string(),
                }],
                next_page_token: "".to_string(),
            },
        ];

        let client = Client {
            inner: Arc::new(InnerClient {
                data: Arc::new(Mutex::new(responses)),
            }),
        };
        let mut resps = vec![];
        let mut stream = client.list_rpc_stream(TestRequest::default()).into_stream();
        while let Some(resp) = stream.next().await {
            if let Ok(resp) = resp {
                resps.push(resp)
            }
        }
        assert_eq!(resps.len(), 2);
        assert_eq!(resps[0].items[0].name, "item1");
        assert_eq!(resps[0].items[1].name, "item2");
        assert_eq!(resps[1].items[0].name, "item3");
    }

    #[cfg(feature = "unstable-stream")]
    #[tokio::test]
    async fn test_item_paginator_into_stream() {
        let responses = vec![
            TestResponse {
                items: vec![
                    PageItem {
                        name: "item1".to_string(),
                    },
                    PageItem {
                        name: "item2".to_string(),
                    },
                ],
                next_page_token: "token1".to_string(),
            },
            TestResponse {
                items: vec![PageItem {
                    name: "item3".to_string(),
                }],
                next_page_token: "".to_string(),
            },
        ];

        let client = Client {
            inner: Arc::new(InnerClient {
                data: Arc::new(Mutex::new(responses)),
            }),
        };
        let mut items = vec![];
        let mut stream = client
            .list_rpc_stream(TestRequest::default())
            .items()
            .into_stream();
        while let Some(item) = stream.next().await {
            if let Ok(item) = item {
                items.push(item)
            }
        }
        assert_eq!(items.len(), 3);
        assert_eq!(items[0].name, "item1");
        assert_eq!(items[1].name, "item2");
        assert_eq!(items[2].name, "item3");
    }
}
