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

/// Implements a [Autokey](crate::traits::) decorator for logging and tracing.
#[derive(Clone, Debug)]
pub struct Autokey<T>
where
    T: crate::traits::Autokey + std::fmt::Debug + Send + Sync,
{
    inner: T,
}

impl<T> Autokey<T>
where
    T: crate::traits::Autokey + std::fmt::Debug + Send + Sync,
{
    pub fn new(inner: T) -> Self {
        Self { inner }
    }
}

impl<T> crate::traits::Autokey for Autokey<T>
where
    T: crate::traits::Autokey + std::fmt::Debug + Send + Sync,
{
    #[tracing::instrument(ret)]
    async fn create_key_handle(
        &self,
        req: crate::model::CreateKeyHandleRequest,
        options: gax::options::RequestOptions,
    ) -> Result<longrunning::model::Operation> {
        self.inner.create_key_handle(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_key_handle(
        &self,
        req: crate::model::GetKeyHandleRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::KeyHandle> {
        self.inner.get_key_handle(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn list_key_handles(
        &self,
        req: crate::model::ListKeyHandlesRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::ListKeyHandlesResponse> {
        self.inner.list_key_handles(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn list_locations(
        &self,
        req: location::model::ListLocationsRequest,
        options: gax::options::RequestOptions,
    ) -> Result<location::model::ListLocationsResponse> {
        self.inner.list_locations(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_location(
        &self,
        req: location::model::GetLocationRequest,
        options: gax::options::RequestOptions,
    ) -> Result<location::model::Location> {
        self.inner.get_location(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn set_iam_policy(
        &self,
        req: iam_v1::model::SetIamPolicyRequest,
        options: gax::options::RequestOptions,
    ) -> Result<iam_v1::model::Policy> {
        self.inner.set_iam_policy(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_iam_policy(
        &self,
        req: iam_v1::model::GetIamPolicyRequest,
        options: gax::options::RequestOptions,
    ) -> Result<iam_v1::model::Policy> {
        self.inner.get_iam_policy(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn test_iam_permissions(
        &self,
        req: iam_v1::model::TestIamPermissionsRequest,
        options: gax::options::RequestOptions,
    ) -> Result<iam_v1::model::TestIamPermissionsResponse> {
        self.inner.test_iam_permissions(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_operation(
        &self,
        req: longrunning::model::GetOperationRequest,
        options: gax::options::RequestOptions,
    ) -> Result<longrunning::model::Operation> {
        self.inner.get_operation(req, options).await
    }

    fn get_polling_policy(
        &self,
        options: &gax::options::RequestOptions,
    ) -> std::sync::Arc<dyn gax::polling_policy::PollingPolicy> {
        self.inner.get_polling_policy(options)
    }

    fn get_polling_backoff_policy(
        &self,
        options: &gax::options::RequestOptions,
    ) -> std::sync::Arc<dyn gax::polling_backoff_policy::PollingBackoffPolicy> {
        self.inner.get_polling_backoff_policy(options)
    }
}

/// Implements a [AutokeyAdmin](crate::traits::) decorator for logging and tracing.
#[derive(Clone, Debug)]
pub struct AutokeyAdmin<T>
where
    T: crate::traits::AutokeyAdmin + std::fmt::Debug + Send + Sync,
{
    inner: T,
}

impl<T> AutokeyAdmin<T>
where
    T: crate::traits::AutokeyAdmin + std::fmt::Debug + Send + Sync,
{
    pub fn new(inner: T) -> Self {
        Self { inner }
    }
}

impl<T> crate::traits::AutokeyAdmin for AutokeyAdmin<T>
where
    T: crate::traits::AutokeyAdmin + std::fmt::Debug + Send + Sync,
{
    #[tracing::instrument(ret)]
    async fn update_autokey_config(
        &self,
        req: crate::model::UpdateAutokeyConfigRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::AutokeyConfig> {
        self.inner.update_autokey_config(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_autokey_config(
        &self,
        req: crate::model::GetAutokeyConfigRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::AutokeyConfig> {
        self.inner.get_autokey_config(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn show_effective_autokey_config(
        &self,
        req: crate::model::ShowEffectiveAutokeyConfigRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::ShowEffectiveAutokeyConfigResponse> {
        self.inner.show_effective_autokey_config(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn list_locations(
        &self,
        req: location::model::ListLocationsRequest,
        options: gax::options::RequestOptions,
    ) -> Result<location::model::ListLocationsResponse> {
        self.inner.list_locations(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_location(
        &self,
        req: location::model::GetLocationRequest,
        options: gax::options::RequestOptions,
    ) -> Result<location::model::Location> {
        self.inner.get_location(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn set_iam_policy(
        &self,
        req: iam_v1::model::SetIamPolicyRequest,
        options: gax::options::RequestOptions,
    ) -> Result<iam_v1::model::Policy> {
        self.inner.set_iam_policy(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_iam_policy(
        &self,
        req: iam_v1::model::GetIamPolicyRequest,
        options: gax::options::RequestOptions,
    ) -> Result<iam_v1::model::Policy> {
        self.inner.get_iam_policy(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn test_iam_permissions(
        &self,
        req: iam_v1::model::TestIamPermissionsRequest,
        options: gax::options::RequestOptions,
    ) -> Result<iam_v1::model::TestIamPermissionsResponse> {
        self.inner.test_iam_permissions(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_operation(
        &self,
        req: longrunning::model::GetOperationRequest,
        options: gax::options::RequestOptions,
    ) -> Result<longrunning::model::Operation> {
        self.inner.get_operation(req, options).await
    }
}

/// Implements a [EkmService](crate::traits::) decorator for logging and tracing.
#[derive(Clone, Debug)]
pub struct EkmService<T>
where
    T: crate::traits::EkmService + std::fmt::Debug + Send + Sync,
{
    inner: T,
}

impl<T> EkmService<T>
where
    T: crate::traits::EkmService + std::fmt::Debug + Send + Sync,
{
    pub fn new(inner: T) -> Self {
        Self { inner }
    }
}

impl<T> crate::traits::EkmService for EkmService<T>
where
    T: crate::traits::EkmService + std::fmt::Debug + Send + Sync,
{
    #[tracing::instrument(ret)]
    async fn list_ekm_connections(
        &self,
        req: crate::model::ListEkmConnectionsRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::ListEkmConnectionsResponse> {
        self.inner.list_ekm_connections(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_ekm_connection(
        &self,
        req: crate::model::GetEkmConnectionRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::EkmConnection> {
        self.inner.get_ekm_connection(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn create_ekm_connection(
        &self,
        req: crate::model::CreateEkmConnectionRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::EkmConnection> {
        self.inner.create_ekm_connection(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn update_ekm_connection(
        &self,
        req: crate::model::UpdateEkmConnectionRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::EkmConnection> {
        self.inner.update_ekm_connection(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_ekm_config(
        &self,
        req: crate::model::GetEkmConfigRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::EkmConfig> {
        self.inner.get_ekm_config(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn update_ekm_config(
        &self,
        req: crate::model::UpdateEkmConfigRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::EkmConfig> {
        self.inner.update_ekm_config(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn verify_connectivity(
        &self,
        req: crate::model::VerifyConnectivityRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::VerifyConnectivityResponse> {
        self.inner.verify_connectivity(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn list_locations(
        &self,
        req: location::model::ListLocationsRequest,
        options: gax::options::RequestOptions,
    ) -> Result<location::model::ListLocationsResponse> {
        self.inner.list_locations(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_location(
        &self,
        req: location::model::GetLocationRequest,
        options: gax::options::RequestOptions,
    ) -> Result<location::model::Location> {
        self.inner.get_location(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn set_iam_policy(
        &self,
        req: iam_v1::model::SetIamPolicyRequest,
        options: gax::options::RequestOptions,
    ) -> Result<iam_v1::model::Policy> {
        self.inner.set_iam_policy(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_iam_policy(
        &self,
        req: iam_v1::model::GetIamPolicyRequest,
        options: gax::options::RequestOptions,
    ) -> Result<iam_v1::model::Policy> {
        self.inner.get_iam_policy(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn test_iam_permissions(
        &self,
        req: iam_v1::model::TestIamPermissionsRequest,
        options: gax::options::RequestOptions,
    ) -> Result<iam_v1::model::TestIamPermissionsResponse> {
        self.inner.test_iam_permissions(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_operation(
        &self,
        req: longrunning::model::GetOperationRequest,
        options: gax::options::RequestOptions,
    ) -> Result<longrunning::model::Operation> {
        self.inner.get_operation(req, options).await
    }
}

/// Implements a [KeyManagementService](crate::traits::) decorator for logging and tracing.
#[derive(Clone, Debug)]
pub struct KeyManagementService<T>
where
    T: crate::traits::KeyManagementService + std::fmt::Debug + Send + Sync,
{
    inner: T,
}

impl<T> KeyManagementService<T>
where
    T: crate::traits::KeyManagementService + std::fmt::Debug + Send + Sync,
{
    pub fn new(inner: T) -> Self {
        Self { inner }
    }
}

impl<T> crate::traits::KeyManagementService for KeyManagementService<T>
where
    T: crate::traits::KeyManagementService + std::fmt::Debug + Send + Sync,
{
    #[tracing::instrument(ret)]
    async fn list_key_rings(
        &self,
        req: crate::model::ListKeyRingsRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::ListKeyRingsResponse> {
        self.inner.list_key_rings(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn list_crypto_keys(
        &self,
        req: crate::model::ListCryptoKeysRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::ListCryptoKeysResponse> {
        self.inner.list_crypto_keys(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn list_crypto_key_versions(
        &self,
        req: crate::model::ListCryptoKeyVersionsRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::ListCryptoKeyVersionsResponse> {
        self.inner.list_crypto_key_versions(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn list_import_jobs(
        &self,
        req: crate::model::ListImportJobsRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::ListImportJobsResponse> {
        self.inner.list_import_jobs(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_key_ring(
        &self,
        req: crate::model::GetKeyRingRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::KeyRing> {
        self.inner.get_key_ring(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_crypto_key(
        &self,
        req: crate::model::GetCryptoKeyRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::CryptoKey> {
        self.inner.get_crypto_key(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_crypto_key_version(
        &self,
        req: crate::model::GetCryptoKeyVersionRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::CryptoKeyVersion> {
        self.inner.get_crypto_key_version(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_public_key(
        &self,
        req: crate::model::GetPublicKeyRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::PublicKey> {
        self.inner.get_public_key(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_import_job(
        &self,
        req: crate::model::GetImportJobRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::ImportJob> {
        self.inner.get_import_job(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn create_key_ring(
        &self,
        req: crate::model::CreateKeyRingRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::KeyRing> {
        self.inner.create_key_ring(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn create_crypto_key(
        &self,
        req: crate::model::CreateCryptoKeyRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::CryptoKey> {
        self.inner.create_crypto_key(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn create_crypto_key_version(
        &self,
        req: crate::model::CreateCryptoKeyVersionRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::CryptoKeyVersion> {
        self.inner.create_crypto_key_version(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn import_crypto_key_version(
        &self,
        req: crate::model::ImportCryptoKeyVersionRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::CryptoKeyVersion> {
        self.inner.import_crypto_key_version(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn create_import_job(
        &self,
        req: crate::model::CreateImportJobRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::ImportJob> {
        self.inner.create_import_job(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn update_crypto_key(
        &self,
        req: crate::model::UpdateCryptoKeyRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::CryptoKey> {
        self.inner.update_crypto_key(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn update_crypto_key_version(
        &self,
        req: crate::model::UpdateCryptoKeyVersionRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::CryptoKeyVersion> {
        self.inner.update_crypto_key_version(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn update_crypto_key_primary_version(
        &self,
        req: crate::model::UpdateCryptoKeyPrimaryVersionRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::CryptoKey> {
        self.inner
            .update_crypto_key_primary_version(req, options)
            .await
    }

    #[tracing::instrument(ret)]
    async fn destroy_crypto_key_version(
        &self,
        req: crate::model::DestroyCryptoKeyVersionRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::CryptoKeyVersion> {
        self.inner.destroy_crypto_key_version(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn restore_crypto_key_version(
        &self,
        req: crate::model::RestoreCryptoKeyVersionRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::CryptoKeyVersion> {
        self.inner.restore_crypto_key_version(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn encrypt(
        &self,
        req: crate::model::EncryptRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::EncryptResponse> {
        self.inner.encrypt(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn decrypt(
        &self,
        req: crate::model::DecryptRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::DecryptResponse> {
        self.inner.decrypt(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn raw_encrypt(
        &self,
        req: crate::model::RawEncryptRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::RawEncryptResponse> {
        self.inner.raw_encrypt(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn raw_decrypt(
        &self,
        req: crate::model::RawDecryptRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::RawDecryptResponse> {
        self.inner.raw_decrypt(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn asymmetric_sign(
        &self,
        req: crate::model::AsymmetricSignRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::AsymmetricSignResponse> {
        self.inner.asymmetric_sign(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn asymmetric_decrypt(
        &self,
        req: crate::model::AsymmetricDecryptRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::AsymmetricDecryptResponse> {
        self.inner.asymmetric_decrypt(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn mac_sign(
        &self,
        req: crate::model::MacSignRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::MacSignResponse> {
        self.inner.mac_sign(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn mac_verify(
        &self,
        req: crate::model::MacVerifyRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::MacVerifyResponse> {
        self.inner.mac_verify(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn generate_random_bytes(
        &self,
        req: crate::model::GenerateRandomBytesRequest,
        options: gax::options::RequestOptions,
    ) -> Result<crate::model::GenerateRandomBytesResponse> {
        self.inner.generate_random_bytes(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn list_locations(
        &self,
        req: location::model::ListLocationsRequest,
        options: gax::options::RequestOptions,
    ) -> Result<location::model::ListLocationsResponse> {
        self.inner.list_locations(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_location(
        &self,
        req: location::model::GetLocationRequest,
        options: gax::options::RequestOptions,
    ) -> Result<location::model::Location> {
        self.inner.get_location(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn set_iam_policy(
        &self,
        req: iam_v1::model::SetIamPolicyRequest,
        options: gax::options::RequestOptions,
    ) -> Result<iam_v1::model::Policy> {
        self.inner.set_iam_policy(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_iam_policy(
        &self,
        req: iam_v1::model::GetIamPolicyRequest,
        options: gax::options::RequestOptions,
    ) -> Result<iam_v1::model::Policy> {
        self.inner.get_iam_policy(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn test_iam_permissions(
        &self,
        req: iam_v1::model::TestIamPermissionsRequest,
        options: gax::options::RequestOptions,
    ) -> Result<iam_v1::model::TestIamPermissionsResponse> {
        self.inner.test_iam_permissions(req, options).await
    }

    #[tracing::instrument(ret)]
    async fn get_operation(
        &self,
        req: longrunning::model::GetOperationRequest,
        options: gax::options::RequestOptions,
    ) -> Result<longrunning::model::Operation> {
        self.inner.get_operation(req, options).await
    }
}