// Copyright 2025 Google LLC
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
//
// Code generated by sidekick. DO NOT EDIT.

pub mod recommender {
    use crate::Result;

    /// A builder for [Recommender][super::super::client::Recommender].
    ///
    /// ```
    /// # tokio_test::block_on(async {
    /// # use google_cloud_recommender_v1::*;
    /// # use builder::recommender::ClientBuilder;
    /// # use client::Recommender;
    /// let builder : ClientBuilder = Recommender::builder();
    /// let client = builder
    ///     .with_endpoint("https://recommender.googleapis.com")
    ///     .build().await?;
    /// # gax::Result::<()>::Ok(()) });
    /// ```
    pub type ClientBuilder =
        gax::client_builder::ClientBuilder<client::Factory, gaxi::options::Credentials>;

    pub(crate) mod client {
        use super::super::super::client::Recommender;
        pub struct Factory;
        impl gax::client_builder::internal::ClientFactory for Factory {
            type Client = Recommender;
            type Credentials = gaxi::options::Credentials;
            async fn build(self, config: gaxi::options::ClientConfig) -> gax::Result<Self::Client> {
                Self::Client::new(config).await
            }
        }
    }

    /// Common implementation for [super::super::client::Recommender] request builders.
    #[derive(Clone, Debug)]
    pub(crate) struct RequestBuilder<R: std::default::Default> {
        stub: std::sync::Arc<dyn super::super::stub::dynamic::Recommender>,
        request: R,
        options: gax::options::RequestOptions,
    }

    impl<R> RequestBuilder<R>
    where
        R: std::default::Default,
    {
        pub(crate) fn new(
            stub: std::sync::Arc<dyn super::super::stub::dynamic::Recommender>,
        ) -> Self {
            Self {
                stub,
                request: R::default(),
                options: gax::options::RequestOptions::default(),
            }
        }
    }

    /// The request builder for [Recommender::list_insights][super::super::client::Recommender::list_insights] calls.
    #[derive(Clone, Debug)]
    pub struct ListInsights(RequestBuilder<crate::model::ListInsightsRequest>);

    impl ListInsights {
        pub(crate) fn new(
            stub: std::sync::Arc<dyn super::super::stub::dynamic::Recommender>,
        ) -> Self {
            Self(RequestBuilder::new(stub))
        }

        /// Sets the full request, replacing any prior values.
        pub fn with_request<V: Into<crate::model::ListInsightsRequest>>(mut self, v: V) -> Self {
            self.0.request = v.into();
            self
        }

        /// Sets all the options, replacing any prior values.
        pub fn with_options<V: Into<gax::options::RequestOptions>>(mut self, v: V) -> Self {
            self.0.options = v.into();
            self
        }

        /// Sends the request.
        pub async fn send(self) -> Result<crate::model::ListInsightsResponse> {
            (*self.0.stub)
                .list_insights(self.0.request, self.0.options)
                .await
                .map(gax::response::Response::into_body)
        }

        /// Streams the responses back.
        pub async fn paginator(
            self,
        ) -> impl gax::paginator::Paginator<crate::model::ListInsightsResponse, gax::error::Error>
        {
            use std::clone::Clone;
            let token = self.0.request.page_token.clone();
            let execute = move |token: String| {
                let mut builder = self.clone();
                builder.0.request = builder.0.request.set_page_token(token);
                builder.send()
            };
            gax::paginator::internal::new_paginator(token, execute)
        }

        /// Sets the value of [parent][crate::model::ListInsightsRequest::parent].
        ///
        /// This is a **required** field for requests.
        pub fn set_parent<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.parent = v.into();
            self
        }

        /// Sets the value of [page_size][crate::model::ListInsightsRequest::page_size].
        pub fn set_page_size<T: Into<i32>>(mut self, v: T) -> Self {
            self.0.request.page_size = v.into();
            self
        }

        /// Sets the value of [page_token][crate::model::ListInsightsRequest::page_token].
        pub fn set_page_token<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.page_token = v.into();
            self
        }

        /// Sets the value of [filter][crate::model::ListInsightsRequest::filter].
        pub fn set_filter<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.filter = v.into();
            self
        }
    }

    #[doc(hidden)]
    impl gax::options::internal::RequestBuilder for ListInsights {
        fn request_options(&mut self) -> &mut gax::options::RequestOptions {
            &mut self.0.options
        }
    }

    /// The request builder for [Recommender::get_insight][super::super::client::Recommender::get_insight] calls.
    #[derive(Clone, Debug)]
    pub struct GetInsight(RequestBuilder<crate::model::GetInsightRequest>);

    impl GetInsight {
        pub(crate) fn new(
            stub: std::sync::Arc<dyn super::super::stub::dynamic::Recommender>,
        ) -> Self {
            Self(RequestBuilder::new(stub))
        }

        /// Sets the full request, replacing any prior values.
        pub fn with_request<V: Into<crate::model::GetInsightRequest>>(mut self, v: V) -> Self {
            self.0.request = v.into();
            self
        }

        /// Sets all the options, replacing any prior values.
        pub fn with_options<V: Into<gax::options::RequestOptions>>(mut self, v: V) -> Self {
            self.0.options = v.into();
            self
        }

        /// Sends the request.
        pub async fn send(self) -> Result<crate::model::Insight> {
            (*self.0.stub)
                .get_insight(self.0.request, self.0.options)
                .await
                .map(gax::response::Response::into_body)
        }

        /// Sets the value of [name][crate::model::GetInsightRequest::name].
        ///
        /// This is a **required** field for requests.
        pub fn set_name<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.name = v.into();
            self
        }
    }

    #[doc(hidden)]
    impl gax::options::internal::RequestBuilder for GetInsight {
        fn request_options(&mut self) -> &mut gax::options::RequestOptions {
            &mut self.0.options
        }
    }

    /// The request builder for [Recommender::mark_insight_accepted][super::super::client::Recommender::mark_insight_accepted] calls.
    #[derive(Clone, Debug)]
    pub struct MarkInsightAccepted(RequestBuilder<crate::model::MarkInsightAcceptedRequest>);

    impl MarkInsightAccepted {
        pub(crate) fn new(
            stub: std::sync::Arc<dyn super::super::stub::dynamic::Recommender>,
        ) -> Self {
            Self(RequestBuilder::new(stub))
        }

        /// Sets the full request, replacing any prior values.
        pub fn with_request<V: Into<crate::model::MarkInsightAcceptedRequest>>(
            mut self,
            v: V,
        ) -> Self {
            self.0.request = v.into();
            self
        }

        /// Sets all the options, replacing any prior values.
        pub fn with_options<V: Into<gax::options::RequestOptions>>(mut self, v: V) -> Self {
            self.0.options = v.into();
            self
        }

        /// Sends the request.
        pub async fn send(self) -> Result<crate::model::Insight> {
            (*self.0.stub)
                .mark_insight_accepted(self.0.request, self.0.options)
                .await
                .map(gax::response::Response::into_body)
        }

        /// Sets the value of [name][crate::model::MarkInsightAcceptedRequest::name].
        ///
        /// This is a **required** field for requests.
        pub fn set_name<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.name = v.into();
            self
        }

        /// Sets the value of [state_metadata][crate::model::MarkInsightAcceptedRequest::state_metadata].
        pub fn set_state_metadata<T, K, V>(mut self, v: T) -> Self
        where
            T: std::iter::IntoIterator<Item = (K, V)>,
            K: std::convert::Into<std::string::String>,
            V: std::convert::Into<std::string::String>,
        {
            self.0.request.state_metadata =
                v.into_iter().map(|(k, v)| (k.into(), v.into())).collect();
            self
        }

        /// Sets the value of [etag][crate::model::MarkInsightAcceptedRequest::etag].
        ///
        /// This is a **required** field for requests.
        pub fn set_etag<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.etag = v.into();
            self
        }
    }

    #[doc(hidden)]
    impl gax::options::internal::RequestBuilder for MarkInsightAccepted {
        fn request_options(&mut self) -> &mut gax::options::RequestOptions {
            &mut self.0.options
        }
    }

    /// The request builder for [Recommender::list_recommendations][super::super::client::Recommender::list_recommendations] calls.
    #[derive(Clone, Debug)]
    pub struct ListRecommendations(RequestBuilder<crate::model::ListRecommendationsRequest>);

    impl ListRecommendations {
        pub(crate) fn new(
            stub: std::sync::Arc<dyn super::super::stub::dynamic::Recommender>,
        ) -> Self {
            Self(RequestBuilder::new(stub))
        }

        /// Sets the full request, replacing any prior values.
        pub fn with_request<V: Into<crate::model::ListRecommendationsRequest>>(
            mut self,
            v: V,
        ) -> Self {
            self.0.request = v.into();
            self
        }

        /// Sets all the options, replacing any prior values.
        pub fn with_options<V: Into<gax::options::RequestOptions>>(mut self, v: V) -> Self {
            self.0.options = v.into();
            self
        }

        /// Sends the request.
        pub async fn send(self) -> Result<crate::model::ListRecommendationsResponse> {
            (*self.0.stub)
                .list_recommendations(self.0.request, self.0.options)
                .await
                .map(gax::response::Response::into_body)
        }

        /// Streams the responses back.
        pub async fn paginator(
            self,
        ) -> impl gax::paginator::Paginator<crate::model::ListRecommendationsResponse, gax::error::Error>
        {
            use std::clone::Clone;
            let token = self.0.request.page_token.clone();
            let execute = move |token: String| {
                let mut builder = self.clone();
                builder.0.request = builder.0.request.set_page_token(token);
                builder.send()
            };
            gax::paginator::internal::new_paginator(token, execute)
        }

        /// Sets the value of [parent][crate::model::ListRecommendationsRequest::parent].
        ///
        /// This is a **required** field for requests.
        pub fn set_parent<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.parent = v.into();
            self
        }

        /// Sets the value of [page_size][crate::model::ListRecommendationsRequest::page_size].
        pub fn set_page_size<T: Into<i32>>(mut self, v: T) -> Self {
            self.0.request.page_size = v.into();
            self
        }

        /// Sets the value of [page_token][crate::model::ListRecommendationsRequest::page_token].
        pub fn set_page_token<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.page_token = v.into();
            self
        }

        /// Sets the value of [filter][crate::model::ListRecommendationsRequest::filter].
        pub fn set_filter<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.filter = v.into();
            self
        }
    }

    #[doc(hidden)]
    impl gax::options::internal::RequestBuilder for ListRecommendations {
        fn request_options(&mut self) -> &mut gax::options::RequestOptions {
            &mut self.0.options
        }
    }

    /// The request builder for [Recommender::get_recommendation][super::super::client::Recommender::get_recommendation] calls.
    #[derive(Clone, Debug)]
    pub struct GetRecommendation(RequestBuilder<crate::model::GetRecommendationRequest>);

    impl GetRecommendation {
        pub(crate) fn new(
            stub: std::sync::Arc<dyn super::super::stub::dynamic::Recommender>,
        ) -> Self {
            Self(RequestBuilder::new(stub))
        }

        /// Sets the full request, replacing any prior values.
        pub fn with_request<V: Into<crate::model::GetRecommendationRequest>>(
            mut self,
            v: V,
        ) -> Self {
            self.0.request = v.into();
            self
        }

        /// Sets all the options, replacing any prior values.
        pub fn with_options<V: Into<gax::options::RequestOptions>>(mut self, v: V) -> Self {
            self.0.options = v.into();
            self
        }

        /// Sends the request.
        pub async fn send(self) -> Result<crate::model::Recommendation> {
            (*self.0.stub)
                .get_recommendation(self.0.request, self.0.options)
                .await
                .map(gax::response::Response::into_body)
        }

        /// Sets the value of [name][crate::model::GetRecommendationRequest::name].
        ///
        /// This is a **required** field for requests.
        pub fn set_name<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.name = v.into();
            self
        }
    }

    #[doc(hidden)]
    impl gax::options::internal::RequestBuilder for GetRecommendation {
        fn request_options(&mut self) -> &mut gax::options::RequestOptions {
            &mut self.0.options
        }
    }

    /// The request builder for [Recommender::mark_recommendation_dismissed][super::super::client::Recommender::mark_recommendation_dismissed] calls.
    #[derive(Clone, Debug)]
    pub struct MarkRecommendationDismissed(
        RequestBuilder<crate::model::MarkRecommendationDismissedRequest>,
    );

    impl MarkRecommendationDismissed {
        pub(crate) fn new(
            stub: std::sync::Arc<dyn super::super::stub::dynamic::Recommender>,
        ) -> Self {
            Self(RequestBuilder::new(stub))
        }

        /// Sets the full request, replacing any prior values.
        pub fn with_request<V: Into<crate::model::MarkRecommendationDismissedRequest>>(
            mut self,
            v: V,
        ) -> Self {
            self.0.request = v.into();
            self
        }

        /// Sets all the options, replacing any prior values.
        pub fn with_options<V: Into<gax::options::RequestOptions>>(mut self, v: V) -> Self {
            self.0.options = v.into();
            self
        }

        /// Sends the request.
        pub async fn send(self) -> Result<crate::model::Recommendation> {
            (*self.0.stub)
                .mark_recommendation_dismissed(self.0.request, self.0.options)
                .await
                .map(gax::response::Response::into_body)
        }

        /// Sets the value of [name][crate::model::MarkRecommendationDismissedRequest::name].
        ///
        /// This is a **required** field for requests.
        pub fn set_name<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.name = v.into();
            self
        }

        /// Sets the value of [etag][crate::model::MarkRecommendationDismissedRequest::etag].
        pub fn set_etag<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.etag = v.into();
            self
        }
    }

    #[doc(hidden)]
    impl gax::options::internal::RequestBuilder for MarkRecommendationDismissed {
        fn request_options(&mut self) -> &mut gax::options::RequestOptions {
            &mut self.0.options
        }
    }

    /// The request builder for [Recommender::mark_recommendation_claimed][super::super::client::Recommender::mark_recommendation_claimed] calls.
    #[derive(Clone, Debug)]
    pub struct MarkRecommendationClaimed(
        RequestBuilder<crate::model::MarkRecommendationClaimedRequest>,
    );

    impl MarkRecommendationClaimed {
        pub(crate) fn new(
            stub: std::sync::Arc<dyn super::super::stub::dynamic::Recommender>,
        ) -> Self {
            Self(RequestBuilder::new(stub))
        }

        /// Sets the full request, replacing any prior values.
        pub fn with_request<V: Into<crate::model::MarkRecommendationClaimedRequest>>(
            mut self,
            v: V,
        ) -> Self {
            self.0.request = v.into();
            self
        }

        /// Sets all the options, replacing any prior values.
        pub fn with_options<V: Into<gax::options::RequestOptions>>(mut self, v: V) -> Self {
            self.0.options = v.into();
            self
        }

        /// Sends the request.
        pub async fn send(self) -> Result<crate::model::Recommendation> {
            (*self.0.stub)
                .mark_recommendation_claimed(self.0.request, self.0.options)
                .await
                .map(gax::response::Response::into_body)
        }

        /// Sets the value of [name][crate::model::MarkRecommendationClaimedRequest::name].
        ///
        /// This is a **required** field for requests.
        pub fn set_name<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.name = v.into();
            self
        }

        /// Sets the value of [state_metadata][crate::model::MarkRecommendationClaimedRequest::state_metadata].
        pub fn set_state_metadata<T, K, V>(mut self, v: T) -> Self
        where
            T: std::iter::IntoIterator<Item = (K, V)>,
            K: std::convert::Into<std::string::String>,
            V: std::convert::Into<std::string::String>,
        {
            self.0.request.state_metadata =
                v.into_iter().map(|(k, v)| (k.into(), v.into())).collect();
            self
        }

        /// Sets the value of [etag][crate::model::MarkRecommendationClaimedRequest::etag].
        ///
        /// This is a **required** field for requests.
        pub fn set_etag<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.etag = v.into();
            self
        }
    }

    #[doc(hidden)]
    impl gax::options::internal::RequestBuilder for MarkRecommendationClaimed {
        fn request_options(&mut self) -> &mut gax::options::RequestOptions {
            &mut self.0.options
        }
    }

    /// The request builder for [Recommender::mark_recommendation_succeeded][super::super::client::Recommender::mark_recommendation_succeeded] calls.
    #[derive(Clone, Debug)]
    pub struct MarkRecommendationSucceeded(
        RequestBuilder<crate::model::MarkRecommendationSucceededRequest>,
    );

    impl MarkRecommendationSucceeded {
        pub(crate) fn new(
            stub: std::sync::Arc<dyn super::super::stub::dynamic::Recommender>,
        ) -> Self {
            Self(RequestBuilder::new(stub))
        }

        /// Sets the full request, replacing any prior values.
        pub fn with_request<V: Into<crate::model::MarkRecommendationSucceededRequest>>(
            mut self,
            v: V,
        ) -> Self {
            self.0.request = v.into();
            self
        }

        /// Sets all the options, replacing any prior values.
        pub fn with_options<V: Into<gax::options::RequestOptions>>(mut self, v: V) -> Self {
            self.0.options = v.into();
            self
        }

        /// Sends the request.
        pub async fn send(self) -> Result<crate::model::Recommendation> {
            (*self.0.stub)
                .mark_recommendation_succeeded(self.0.request, self.0.options)
                .await
                .map(gax::response::Response::into_body)
        }

        /// Sets the value of [name][crate::model::MarkRecommendationSucceededRequest::name].
        ///
        /// This is a **required** field for requests.
        pub fn set_name<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.name = v.into();
            self
        }

        /// Sets the value of [state_metadata][crate::model::MarkRecommendationSucceededRequest::state_metadata].
        pub fn set_state_metadata<T, K, V>(mut self, v: T) -> Self
        where
            T: std::iter::IntoIterator<Item = (K, V)>,
            K: std::convert::Into<std::string::String>,
            V: std::convert::Into<std::string::String>,
        {
            self.0.request.state_metadata =
                v.into_iter().map(|(k, v)| (k.into(), v.into())).collect();
            self
        }

        /// Sets the value of [etag][crate::model::MarkRecommendationSucceededRequest::etag].
        ///
        /// This is a **required** field for requests.
        pub fn set_etag<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.etag = v.into();
            self
        }
    }

    #[doc(hidden)]
    impl gax::options::internal::RequestBuilder for MarkRecommendationSucceeded {
        fn request_options(&mut self) -> &mut gax::options::RequestOptions {
            &mut self.0.options
        }
    }

    /// The request builder for [Recommender::mark_recommendation_failed][super::super::client::Recommender::mark_recommendation_failed] calls.
    #[derive(Clone, Debug)]
    pub struct MarkRecommendationFailed(
        RequestBuilder<crate::model::MarkRecommendationFailedRequest>,
    );

    impl MarkRecommendationFailed {
        pub(crate) fn new(
            stub: std::sync::Arc<dyn super::super::stub::dynamic::Recommender>,
        ) -> Self {
            Self(RequestBuilder::new(stub))
        }

        /// Sets the full request, replacing any prior values.
        pub fn with_request<V: Into<crate::model::MarkRecommendationFailedRequest>>(
            mut self,
            v: V,
        ) -> Self {
            self.0.request = v.into();
            self
        }

        /// Sets all the options, replacing any prior values.
        pub fn with_options<V: Into<gax::options::RequestOptions>>(mut self, v: V) -> Self {
            self.0.options = v.into();
            self
        }

        /// Sends the request.
        pub async fn send(self) -> Result<crate::model::Recommendation> {
            (*self.0.stub)
                .mark_recommendation_failed(self.0.request, self.0.options)
                .await
                .map(gax::response::Response::into_body)
        }

        /// Sets the value of [name][crate::model::MarkRecommendationFailedRequest::name].
        ///
        /// This is a **required** field for requests.
        pub fn set_name<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.name = v.into();
            self
        }

        /// Sets the value of [state_metadata][crate::model::MarkRecommendationFailedRequest::state_metadata].
        pub fn set_state_metadata<T, K, V>(mut self, v: T) -> Self
        where
            T: std::iter::IntoIterator<Item = (K, V)>,
            K: std::convert::Into<std::string::String>,
            V: std::convert::Into<std::string::String>,
        {
            self.0.request.state_metadata =
                v.into_iter().map(|(k, v)| (k.into(), v.into())).collect();
            self
        }

        /// Sets the value of [etag][crate::model::MarkRecommendationFailedRequest::etag].
        ///
        /// This is a **required** field for requests.
        pub fn set_etag<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.etag = v.into();
            self
        }
    }

    #[doc(hidden)]
    impl gax::options::internal::RequestBuilder for MarkRecommendationFailed {
        fn request_options(&mut self) -> &mut gax::options::RequestOptions {
            &mut self.0.options
        }
    }

    /// The request builder for [Recommender::get_recommender_config][super::super::client::Recommender::get_recommender_config] calls.
    #[derive(Clone, Debug)]
    pub struct GetRecommenderConfig(RequestBuilder<crate::model::GetRecommenderConfigRequest>);

    impl GetRecommenderConfig {
        pub(crate) fn new(
            stub: std::sync::Arc<dyn super::super::stub::dynamic::Recommender>,
        ) -> Self {
            Self(RequestBuilder::new(stub))
        }

        /// Sets the full request, replacing any prior values.
        pub fn with_request<V: Into<crate::model::GetRecommenderConfigRequest>>(
            mut self,
            v: V,
        ) -> Self {
            self.0.request = v.into();
            self
        }

        /// Sets all the options, replacing any prior values.
        pub fn with_options<V: Into<gax::options::RequestOptions>>(mut self, v: V) -> Self {
            self.0.options = v.into();
            self
        }

        /// Sends the request.
        pub async fn send(self) -> Result<crate::model::RecommenderConfig> {
            (*self.0.stub)
                .get_recommender_config(self.0.request, self.0.options)
                .await
                .map(gax::response::Response::into_body)
        }

        /// Sets the value of [name][crate::model::GetRecommenderConfigRequest::name].
        ///
        /// This is a **required** field for requests.
        pub fn set_name<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.name = v.into();
            self
        }
    }

    #[doc(hidden)]
    impl gax::options::internal::RequestBuilder for GetRecommenderConfig {
        fn request_options(&mut self) -> &mut gax::options::RequestOptions {
            &mut self.0.options
        }
    }

    /// The request builder for [Recommender::update_recommender_config][super::super::client::Recommender::update_recommender_config] calls.
    #[derive(Clone, Debug)]
    pub struct UpdateRecommenderConfig(
        RequestBuilder<crate::model::UpdateRecommenderConfigRequest>,
    );

    impl UpdateRecommenderConfig {
        pub(crate) fn new(
            stub: std::sync::Arc<dyn super::super::stub::dynamic::Recommender>,
        ) -> Self {
            Self(RequestBuilder::new(stub))
        }

        /// Sets the full request, replacing any prior values.
        pub fn with_request<V: Into<crate::model::UpdateRecommenderConfigRequest>>(
            mut self,
            v: V,
        ) -> Self {
            self.0.request = v.into();
            self
        }

        /// Sets all the options, replacing any prior values.
        pub fn with_options<V: Into<gax::options::RequestOptions>>(mut self, v: V) -> Self {
            self.0.options = v.into();
            self
        }

        /// Sends the request.
        pub async fn send(self) -> Result<crate::model::RecommenderConfig> {
            (*self.0.stub)
                .update_recommender_config(self.0.request, self.0.options)
                .await
                .map(gax::response::Response::into_body)
        }

        /// Sets the value of [recommender_config][crate::model::UpdateRecommenderConfigRequest::recommender_config].
        ///
        /// This is a **required** field for requests.
        pub fn set_recommender_config<
            T: Into<std::option::Option<crate::model::RecommenderConfig>>,
        >(
            mut self,
            v: T,
        ) -> Self {
            self.0.request.recommender_config = v.into();
            self
        }

        /// Sets the value of [update_mask][crate::model::UpdateRecommenderConfigRequest::update_mask].
        pub fn set_update_mask<T: Into<std::option::Option<wkt::FieldMask>>>(
            mut self,
            v: T,
        ) -> Self {
            self.0.request.update_mask = v.into();
            self
        }

        /// Sets the value of [validate_only][crate::model::UpdateRecommenderConfigRequest::validate_only].
        pub fn set_validate_only<T: Into<bool>>(mut self, v: T) -> Self {
            self.0.request.validate_only = v.into();
            self
        }
    }

    #[doc(hidden)]
    impl gax::options::internal::RequestBuilder for UpdateRecommenderConfig {
        fn request_options(&mut self) -> &mut gax::options::RequestOptions {
            &mut self.0.options
        }
    }

    /// The request builder for [Recommender::get_insight_type_config][super::super::client::Recommender::get_insight_type_config] calls.
    #[derive(Clone, Debug)]
    pub struct GetInsightTypeConfig(RequestBuilder<crate::model::GetInsightTypeConfigRequest>);

    impl GetInsightTypeConfig {
        pub(crate) fn new(
            stub: std::sync::Arc<dyn super::super::stub::dynamic::Recommender>,
        ) -> Self {
            Self(RequestBuilder::new(stub))
        }

        /// Sets the full request, replacing any prior values.
        pub fn with_request<V: Into<crate::model::GetInsightTypeConfigRequest>>(
            mut self,
            v: V,
        ) -> Self {
            self.0.request = v.into();
            self
        }

        /// Sets all the options, replacing any prior values.
        pub fn with_options<V: Into<gax::options::RequestOptions>>(mut self, v: V) -> Self {
            self.0.options = v.into();
            self
        }

        /// Sends the request.
        pub async fn send(self) -> Result<crate::model::InsightTypeConfig> {
            (*self.0.stub)
                .get_insight_type_config(self.0.request, self.0.options)
                .await
                .map(gax::response::Response::into_body)
        }

        /// Sets the value of [name][crate::model::GetInsightTypeConfigRequest::name].
        ///
        /// This is a **required** field for requests.
        pub fn set_name<T: Into<std::string::String>>(mut self, v: T) -> Self {
            self.0.request.name = v.into();
            self
        }
    }

    #[doc(hidden)]
    impl gax::options::internal::RequestBuilder for GetInsightTypeConfig {
        fn request_options(&mut self) -> &mut gax::options::RequestOptions {
            &mut self.0.options
        }
    }

    /// The request builder for [Recommender::update_insight_type_config][super::super::client::Recommender::update_insight_type_config] calls.
    #[derive(Clone, Debug)]
    pub struct UpdateInsightTypeConfig(
        RequestBuilder<crate::model::UpdateInsightTypeConfigRequest>,
    );

    impl UpdateInsightTypeConfig {
        pub(crate) fn new(
            stub: std::sync::Arc<dyn super::super::stub::dynamic::Recommender>,
        ) -> Self {
            Self(RequestBuilder::new(stub))
        }

        /// Sets the full request, replacing any prior values.
        pub fn with_request<V: Into<crate::model::UpdateInsightTypeConfigRequest>>(
            mut self,
            v: V,
        ) -> Self {
            self.0.request = v.into();
            self
        }

        /// Sets all the options, replacing any prior values.
        pub fn with_options<V: Into<gax::options::RequestOptions>>(mut self, v: V) -> Self {
            self.0.options = v.into();
            self
        }

        /// Sends the request.
        pub async fn send(self) -> Result<crate::model::InsightTypeConfig> {
            (*self.0.stub)
                .update_insight_type_config(self.0.request, self.0.options)
                .await
                .map(gax::response::Response::into_body)
        }

        /// Sets the value of [insight_type_config][crate::model::UpdateInsightTypeConfigRequest::insight_type_config].
        ///
        /// This is a **required** field for requests.
        pub fn set_insight_type_config<
            T: Into<std::option::Option<crate::model::InsightTypeConfig>>,
        >(
            mut self,
            v: T,
        ) -> Self {
            self.0.request.insight_type_config = v.into();
            self
        }

        /// Sets the value of [update_mask][crate::model::UpdateInsightTypeConfigRequest::update_mask].
        pub fn set_update_mask<T: Into<std::option::Option<wkt::FieldMask>>>(
            mut self,
            v: T,
        ) -> Self {
            self.0.request.update_mask = v.into();
            self
        }

        /// Sets the value of [validate_only][crate::model::UpdateInsightTypeConfigRequest::validate_only].
        pub fn set_validate_only<T: Into<bool>>(mut self, v: T) -> Self {
            self.0.request.validate_only = v.into();
            self
        }
    }

    #[doc(hidden)]
    impl gax::options::internal::RequestBuilder for UpdateInsightTypeConfig {
        fn request_options(&mut self) -> &mut gax::options::RequestOptions {
            &mut self.0.options
        }
    }
}
