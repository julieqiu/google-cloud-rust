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
//
// Code generated by sidekick. DO NOT EDIT.

/// The messages and enums that are part of this client library.
pub mod model;

use gax::error::{Error, HttpError};
use google_cloud_auth::{Credential, CredentialConfig};
use std::sync::Arc;

const DEFAULT_HOST: &str = "https://cloud.googleapis.com/";

/// A `Result` alias where the `Err` case is an [Error].
pub type Result<T> = std::result::Result<T, Error>;

struct InnerClient {
    http_client: reqwest::Client,
    cred: Credential,
    endpoint: String,
}

#[derive(Default)]
pub struct ConfigBuilder {
    pub(crate) endpoint: Option<String>,
    pub(crate) client: Option<reqwest::Client>,
    pub(crate) cred: Option<Credential>,
}

impl ConfigBuilder {
    /// Returns a default [ConfigBuilder].
    pub fn new() -> Self {
        Self::default()
    }

    /// Sets an endpoint that overrides the default endpoint for a service.
    pub fn set_endpoint<T: Into<String>>(mut self, v: T) -> Self {
        self.endpoint = Some(v.into());
        self
    }

    pub(crate) fn default_client() -> reqwest::Client {
        reqwest::Client::builder().build().unwrap()
    }

    pub(crate) async fn default_credential() -> Result<Credential> {
        let cc = CredentialConfig::builder()
            .scopes(vec![
                "https://www.googleapis.com/auth/cloud-platform".to_string()
            ])
            .build()
            .map_err(Error::authentication)?;
        Credential::find_default(cc)
            .await
            .map_err(Error::authentication)
    }
}

#[derive(serde::Serialize)]
#[allow(dead_code)]
struct NoBody {}

/// An abstract interface that provides location-related information for
/// a service. Service-specific metadata is provided through the
/// [Location.metadata][google.cloud.location.Location.metadata] field.
#[derive(Clone)]
pub struct LocationsClient {
    inner: Arc<InnerClient>,
}

impl LocationsClient {
    pub async fn new() -> Result<Self> {
        Self::new_with_config(ConfigBuilder::new()).await
    }

    pub async fn new_with_config(conf: ConfigBuilder) -> Result<Self> {
        let inner = InnerClient {
            http_client: conf.client.unwrap_or(ConfigBuilder::default_client()),
            cred: conf
                .cred
                .unwrap_or(ConfigBuilder::default_credential().await?),
            endpoint: conf.endpoint.unwrap_or(DEFAULT_HOST.to_string()),
        };
        Ok(Self {
            inner: Arc::new(inner),
        })
    }

    /// Lists information about the supported locations for this service.
    pub async fn list_locations(
        &self,
        req: crate::model::ListLocationsRequest,
    ) -> Result<crate::model::ListLocationsResponse> {
        let inner_client = self.inner.clone();
        let builder = inner_client
            .http_client
            .get(format!("{}/v1/{}", inner_client.endpoint, req.name,))
            .query(&[("alt", "json")]);
        let builder =
            gax::query_parameter::add(builder, "filter", &req.filter).map_err(Error::other)?;
        let builder =
            gax::query_parameter::add(builder, "pageSize", &req.page_size).map_err(Error::other)?;
        let builder = gax::query_parameter::add(builder, "pageToken", &req.page_token)
            .map_err(Error::other)?;
        self.execute(builder, None::<NoBody>).await
    }

    /// Gets information about a location.
    pub async fn get_location(
        &self,
        req: crate::model::GetLocationRequest,
    ) -> Result<crate::model::Location> {
        let inner_client = self.inner.clone();
        let builder = inner_client
            .http_client
            .get(format!("{}/v1/{}", inner_client.endpoint, req.name,))
            .query(&[("alt", "json")]);
        self.execute(builder, None::<NoBody>).await
    }

    async fn execute<I: serde::ser::Serialize, O: serde::de::DeserializeOwned>(
        &self,
        mut builder: reqwest::RequestBuilder,
        body: Option<I>,
    ) -> Result<O> {
        let inner_client = self.inner.clone();
        builder = builder.bearer_auth(
            &inner_client
                .cred
                .access_token()
                .await
                .map_err(Error::authentication)?
                .value,
        );
        if let Some(body) = body {
            builder = builder.json(&body);
        }
        let resp = builder.send().await.map_err(Error::io)?;
        if !resp.status().is_success() {
            let status = resp.status().as_u16();
            let headers = gax::error::convert_headers(resp.headers());
            let body = resp.bytes().await.map_err(Error::io)?;
            return Err(HttpError::new(status, headers, Some(body)).into());
        }
        let response = resp.json::<O>().await.map_err(Error::serde)?;
        Ok(response)
    }
}
