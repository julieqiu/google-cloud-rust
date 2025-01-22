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

use crate::Result;
use std::sync::Arc;

/// An implementation of [crate::traits::TraceService] to make requests with.
///
/// `TraceService` has various configuration parameters, but the defaults
/// are set to work with most applications.
///
/// `TraceService` holds a connection pool internally, it is advised to
/// create one and the reuse it.  You do not need to wrap `TraceService` in
/// an [Rc](std::rc::Rc) or [Arc] to reuse it, because it already uses an `Arc`
/// internally.
///
/// Service for collecting and viewing traces and spans within a trace.
///
/// A trace is a collection of spans corresponding to a single
/// operation or a set of operations in an application.
///
/// A span is an individual timed event which forms a node of the trace tree.
/// A single trace can contain spans from multiple services.
#[derive(Clone, Debug)]
pub struct TraceService {
    inner: Arc<dyn crate::traits::dyntraits::TraceService>,
}

impl TraceService {
    /// Creates a new client with the default configuration.
    pub async fn new() -> Result<Self> {
        Self::new_with_config(gax::options::ClientConfig::default()).await
    }

    /// Creates a new client with the specified configuration.
    pub async fn new_with_config(conf: gax::options::ClientConfig) -> Result<Self> {
        let inner = Self::build_inner(conf).await?;
        Ok(Self { inner })
    }

    /// Creates a new client from the provided stub.
    ///
    /// The most common case for calling this function is when mocking the
    /// client.
    pub fn from_stub<T>(stub: T) -> Self
    where
        T: crate::traits::TraceService + 'static,
    {
        Self {
            inner: Arc::new(stub),
        }
    }

    async fn build_inner(
        conf: gax::options::ClientConfig,
    ) -> Result<Arc<dyn crate::traits::dyntraits::TraceService>> {
        if conf.tracing_enabled() {
            return Ok(Arc::new(Self::build_with_tracing(conf).await?));
        }
        Ok(Arc::new(Self::build_transport(conf).await?))
    }

    async fn build_transport(
        conf: gax::options::ClientConfig,
    ) -> Result<impl crate::traits::TraceService> {
        crate::transport::TraceService::new(conf).await
    }

    async fn build_with_tracing(
        conf: gax::options::ClientConfig,
    ) -> Result<impl crate::traits::TraceService> {
        Self::build_transport(conf)
            .await
            .map(crate::tracing::TraceService::new)
    }

    /// Batch writes new spans to new or existing traces. You cannot update
    /// existing spans.
    pub fn batch_write_spans(
        &self,
        name: impl Into<std::string::String>,
    ) -> crate::builders::trace_service::BatchWriteSpans {
        crate::builders::trace_service::BatchWriteSpans::new(self.inner.clone())
            .set_name(name.into())
    }

    /// Creates a new span.
    pub fn create_span(
        &self,
        name: impl Into<std::string::String>,
    ) -> crate::builders::trace_service::CreateSpan {
        crate::builders::trace_service::CreateSpan::new(self.inner.clone()).set_name(name.into())
    }
}