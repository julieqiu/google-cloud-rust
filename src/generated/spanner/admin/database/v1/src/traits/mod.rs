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

use gax::error::Error;

pub(crate) mod dyntraits;

/// Cloud Spanner Database Admin API
///
/// The Cloud Spanner Database Admin API can be used to:
///
/// * create, drop, and list databases
/// * update the schema of pre-existing databases
/// * create, delete, copy and list backups for a database
/// * restore a database from an existing backup
///
/// # Mocking
///
/// Application developers may use this trait to mock the spanner clients.
///
/// Services gain new RPCs routinely. Consequently, this trait gains new methods
/// too. To avoid breaking applications the trait provides a default
/// implementation for each method. These implementations return an error.
pub trait DatabaseAdmin: std::fmt::Debug + Send + Sync {
    /// Lists Cloud Spanner databases.
    fn list_databases(
        &self,
        _req: crate::model::ListDatabasesRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<crate::model::ListDatabasesResponse>> + Send
    {
        std::future::ready::<crate::Result<crate::model::ListDatabasesResponse>>(Err(Error::other(
            "unimplemented",
        )))
    }

    /// Creates a new Cloud Spanner database and starts to prepare it for serving.
    /// The returned [long-running operation][google.longrunning.Operation] will
    /// have a name of the format `<database_name>/operations/<operation_id>` and
    /// can be used to track preparation of the database. The
    /// [metadata][google.longrunning.Operation.metadata] field type is
    /// [CreateDatabaseMetadata][google.spanner.admin.database.v1.CreateDatabaseMetadata].
    /// The [response][google.longrunning.Operation.response] field type is
    /// [Database][google.spanner.admin.database.v1.Database], if successful.
    ///
    /// [google.longrunning.Operation]: longrunning::model::Operation
    /// [google.longrunning.Operation.metadata]: longrunning::model::Operation::metadata
    /// [google.longrunning.Operation.response]: longrunning::model::Operation::result
    /// [google.spanner.admin.database.v1.CreateDatabaseMetadata]: crate::model::CreateDatabaseMetadata
    /// [google.spanner.admin.database.v1.Database]: crate::model::Database
    fn create_database(
        &self,
        _req: crate::model::CreateDatabaseRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<longrunning::model::Operation>> + Send
    {
        std::future::ready::<crate::Result<longrunning::model::Operation>>(Err(Error::other(
            "unimplemented",
        )))
    }

    /// Gets the state of a Cloud Spanner database.
    fn get_database(
        &self,
        _req: crate::model::GetDatabaseRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<crate::model::Database>> + Send {
        std::future::ready::<crate::Result<crate::model::Database>>(Err(Error::other(
            "unimplemented",
        )))
    }

    /// Updates a Cloud Spanner database. The returned
    /// [long-running operation][google.longrunning.Operation] can be used to track
    /// the progress of updating the database. If the named database does not
    /// exist, returns `NOT_FOUND`.
    ///
    /// While the operation is pending:
    ///
    /// * The database's
    ///   [reconciling][google.spanner.admin.database.v1.Database.reconciling]
    ///   field is set to true.
    /// * Cancelling the operation is best-effort. If the cancellation succeeds,
    ///   the operation metadata's
    ///   [cancel_time][google.spanner.admin.database.v1.UpdateDatabaseMetadata.cancel_time]
    ///   is set, the updates are reverted, and the operation terminates with a
    ///   `CANCELLED` status.
    /// * New UpdateDatabase requests will return a `FAILED_PRECONDITION` error
    ///   until the pending operation is done (returns successfully or with
    ///   error).
    /// * Reading the database via the API continues to give the pre-request
    ///   values.
    ///
    /// Upon completion of the returned operation:
    ///
    /// * The new values are in effect and readable via the API.
    /// * The database's
    ///   [reconciling][google.spanner.admin.database.v1.Database.reconciling]
    ///   field becomes false.
    ///
    /// The returned [long-running operation][google.longrunning.Operation] will
    /// have a name of the format
    /// `projects/<project>/instances/<instance>/databases/<database>/operations/<operation_id>`
    /// and can be used to track the database modification. The
    /// [metadata][google.longrunning.Operation.metadata] field type is
    /// [UpdateDatabaseMetadata][google.spanner.admin.database.v1.UpdateDatabaseMetadata].
    /// The [response][google.longrunning.Operation.response] field type is
    /// [Database][google.spanner.admin.database.v1.Database], if successful.
    ///
    /// [google.longrunning.Operation]: longrunning::model::Operation
    /// [google.longrunning.Operation.metadata]: longrunning::model::Operation::metadata
    /// [google.longrunning.Operation.response]: longrunning::model::Operation::result
    /// [google.spanner.admin.database.v1.Database]: crate::model::Database
    /// [google.spanner.admin.database.v1.Database.reconciling]: crate::model::Database::reconciling
    /// [google.spanner.admin.database.v1.UpdateDatabaseMetadata]: crate::model::UpdateDatabaseMetadata
    /// [google.spanner.admin.database.v1.UpdateDatabaseMetadata.cancel_time]: crate::model::UpdateDatabaseMetadata::cancel_time
    fn update_database(
        &self,
        _req: crate::model::UpdateDatabaseRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<longrunning::model::Operation>> + Send
    {
        std::future::ready::<crate::Result<longrunning::model::Operation>>(Err(Error::other(
            "unimplemented",
        )))
    }

    /// Updates the schema of a Cloud Spanner database by
    /// creating/altering/dropping tables, columns, indexes, etc. The returned
    /// [long-running operation][google.longrunning.Operation] will have a name of
    /// the format `<database_name>/operations/<operation_id>` and can be used to
    /// track execution of the schema change(s). The
    /// [metadata][google.longrunning.Operation.metadata] field type is
    /// [UpdateDatabaseDdlMetadata][google.spanner.admin.database.v1.UpdateDatabaseDdlMetadata].
    /// The operation has no response.
    ///
    /// [google.longrunning.Operation]: longrunning::model::Operation
    /// [google.longrunning.Operation.metadata]: longrunning::model::Operation::metadata
    /// [google.spanner.admin.database.v1.UpdateDatabaseDdlMetadata]: crate::model::UpdateDatabaseDdlMetadata
    fn update_database_ddl(
        &self,
        _req: crate::model::UpdateDatabaseDdlRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<longrunning::model::Operation>> + Send
    {
        std::future::ready::<crate::Result<longrunning::model::Operation>>(Err(Error::other(
            "unimplemented",
        )))
    }

    /// Drops (aka deletes) a Cloud Spanner database.
    /// Completed backups for the database will be retained according to their
    /// `expire_time`.
    /// Note: Cloud Spanner might continue to accept requests for a few seconds
    /// after the database has been deleted.
    fn drop_database(
        &self,
        _req: crate::model::DropDatabaseRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<wkt::Empty>> + Send {
        std::future::ready::<crate::Result<wkt::Empty>>(Err(Error::other("unimplemented")))
    }

    /// Returns the schema of a Cloud Spanner database as a list of formatted
    /// DDL statements. This method does not show pending schema updates, those may
    /// be queried using the [Operations][google.longrunning.Operations] API.
    ///
    /// [google.longrunning.Operations]: longrunning::traits::Operations
    fn get_database_ddl(
        &self,
        _req: crate::model::GetDatabaseDdlRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<crate::model::GetDatabaseDdlResponse>> + Send
    {
        std::future::ready::<crate::Result<crate::model::GetDatabaseDdlResponse>>(Err(
            Error::other("unimplemented"),
        ))
    }

    /// Sets the access control policy on a database or backup resource.
    /// Replaces any existing policy.
    ///
    /// Authorization requires `spanner.databases.setIamPolicy`
    /// permission on [resource][google.iam.v1.SetIamPolicyRequest.resource].
    /// For backups, authorization requires `spanner.backups.setIamPolicy`
    /// permission on [resource][google.iam.v1.SetIamPolicyRequest.resource].
    ///
    /// [google.iam.v1.SetIamPolicyRequest.resource]: iam_v1::model::SetIamPolicyRequest::resource
    fn set_iam_policy(
        &self,
        _req: iam_v1::model::SetIamPolicyRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<iam_v1::model::Policy>> + Send {
        std::future::ready::<crate::Result<iam_v1::model::Policy>>(Err(Error::other(
            "unimplemented",
        )))
    }

    /// Gets the access control policy for a database or backup resource.
    /// Returns an empty policy if a database or backup exists but does not have a
    /// policy set.
    ///
    /// Authorization requires `spanner.databases.getIamPolicy` permission on
    /// [resource][google.iam.v1.GetIamPolicyRequest.resource].
    /// For backups, authorization requires `spanner.backups.getIamPolicy`
    /// permission on [resource][google.iam.v1.GetIamPolicyRequest.resource].
    ///
    /// [google.iam.v1.GetIamPolicyRequest.resource]: iam_v1::model::GetIamPolicyRequest::resource
    fn get_iam_policy(
        &self,
        _req: iam_v1::model::GetIamPolicyRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<iam_v1::model::Policy>> + Send {
        std::future::ready::<crate::Result<iam_v1::model::Policy>>(Err(Error::other(
            "unimplemented",
        )))
    }

    /// Returns permissions that the caller has on the specified database or backup
    /// resource.
    ///
    /// Attempting this RPC on a non-existent Cloud Spanner database will
    /// result in a NOT_FOUND error if the user has
    /// `spanner.databases.list` permission on the containing Cloud
    /// Spanner instance. Otherwise returns an empty set of permissions.
    /// Calling this method on a backup that does not exist will
    /// result in a NOT_FOUND error if the user has
    /// `spanner.backups.list` permission on the containing instance.
    fn test_iam_permissions(
        &self,
        _req: iam_v1::model::TestIamPermissionsRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<iam_v1::model::TestIamPermissionsResponse>> + Send
    {
        std::future::ready::<crate::Result<iam_v1::model::TestIamPermissionsResponse>>(Err(
            Error::other("unimplemented"),
        ))
    }

    /// Starts creating a new Cloud Spanner Backup.
    /// The returned backup [long-running operation][google.longrunning.Operation]
    /// will have a name of the format
    /// `projects/<project>/instances/<instance>/backups/<backup>/operations/<operation_id>`
    /// and can be used to track creation of the backup. The
    /// [metadata][google.longrunning.Operation.metadata] field type is
    /// [CreateBackupMetadata][google.spanner.admin.database.v1.CreateBackupMetadata].
    /// The [response][google.longrunning.Operation.response] field type is
    /// [Backup][google.spanner.admin.database.v1.Backup], if successful.
    /// Cancelling the returned operation will stop the creation and delete the
    /// backup. There can be only one pending backup creation per database. Backup
    /// creation of different databases can run concurrently.
    ///
    /// [google.longrunning.Operation]: longrunning::model::Operation
    /// [google.longrunning.Operation.metadata]: longrunning::model::Operation::metadata
    /// [google.longrunning.Operation.response]: longrunning::model::Operation::result
    /// [google.spanner.admin.database.v1.Backup]: crate::model::Backup
    /// [google.spanner.admin.database.v1.CreateBackupMetadata]: crate::model::CreateBackupMetadata
    fn create_backup(
        &self,
        _req: crate::model::CreateBackupRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<longrunning::model::Operation>> + Send
    {
        std::future::ready::<crate::Result<longrunning::model::Operation>>(Err(Error::other(
            "unimplemented",
        )))
    }

    /// Starts copying a Cloud Spanner Backup.
    /// The returned backup [long-running operation][google.longrunning.Operation]
    /// will have a name of the format
    /// `projects/<project>/instances/<instance>/backups/<backup>/operations/<operation_id>`
    /// and can be used to track copying of the backup. The operation is associated
    /// with the destination backup.
    /// The [metadata][google.longrunning.Operation.metadata] field type is
    /// [CopyBackupMetadata][google.spanner.admin.database.v1.CopyBackupMetadata].
    /// The [response][google.longrunning.Operation.response] field type is
    /// [Backup][google.spanner.admin.database.v1.Backup], if successful.
    /// Cancelling the returned operation will stop the copying and delete the
    /// destination backup. Concurrent CopyBackup requests can run on the same
    /// source backup.
    ///
    /// [google.longrunning.Operation]: longrunning::model::Operation
    /// [google.longrunning.Operation.metadata]: longrunning::model::Operation::metadata
    /// [google.longrunning.Operation.response]: longrunning::model::Operation::result
    /// [google.spanner.admin.database.v1.Backup]: crate::model::Backup
    /// [google.spanner.admin.database.v1.CopyBackupMetadata]: crate::model::CopyBackupMetadata
    fn copy_backup(
        &self,
        _req: crate::model::CopyBackupRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<longrunning::model::Operation>> + Send
    {
        std::future::ready::<crate::Result<longrunning::model::Operation>>(Err(Error::other(
            "unimplemented",
        )))
    }

    /// Gets metadata on a pending or completed
    /// [Backup][google.spanner.admin.database.v1.Backup].
    ///
    /// [google.spanner.admin.database.v1.Backup]: crate::model::Backup
    fn get_backup(
        &self,
        _req: crate::model::GetBackupRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<crate::model::Backup>> + Send {
        std::future::ready::<crate::Result<crate::model::Backup>>(Err(Error::other(
            "unimplemented",
        )))
    }

    /// Updates a pending or completed
    /// [Backup][google.spanner.admin.database.v1.Backup].
    ///
    /// [google.spanner.admin.database.v1.Backup]: crate::model::Backup
    fn update_backup(
        &self,
        _req: crate::model::UpdateBackupRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<crate::model::Backup>> + Send {
        std::future::ready::<crate::Result<crate::model::Backup>>(Err(Error::other(
            "unimplemented",
        )))
    }

    /// Deletes a pending or completed
    /// [Backup][google.spanner.admin.database.v1.Backup].
    ///
    /// [google.spanner.admin.database.v1.Backup]: crate::model::Backup
    fn delete_backup(
        &self,
        _req: crate::model::DeleteBackupRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<wkt::Empty>> + Send {
        std::future::ready::<crate::Result<wkt::Empty>>(Err(Error::other("unimplemented")))
    }

    /// Lists completed and pending backups.
    /// Backups returned are ordered by `create_time` in descending order,
    /// starting from the most recent `create_time`.
    fn list_backups(
        &self,
        _req: crate::model::ListBackupsRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<crate::model::ListBackupsResponse>> + Send
    {
        std::future::ready::<crate::Result<crate::model::ListBackupsResponse>>(Err(Error::other(
            "unimplemented",
        )))
    }

    /// Create a new database by restoring from a completed backup. The new
    /// database must be in the same project and in an instance with the same
    /// instance configuration as the instance containing
    /// the backup. The returned database [long-running
    /// operation][google.longrunning.Operation] has a name of the format
    /// `projects/<project>/instances/<instance>/databases/<database>/operations/<operation_id>`,
    /// and can be used to track the progress of the operation, and to cancel it.
    /// The [metadata][google.longrunning.Operation.metadata] field type is
    /// [RestoreDatabaseMetadata][google.spanner.admin.database.v1.RestoreDatabaseMetadata].
    /// The [response][google.longrunning.Operation.response] type
    /// is [Database][google.spanner.admin.database.v1.Database], if
    /// successful. Cancelling the returned operation will stop the restore and
    /// delete the database.
    /// There can be only one database being restored into an instance at a time.
    /// Once the restore operation completes, a new restore operation can be
    /// initiated, without waiting for the optimize operation associated with the
    /// first restore to complete.
    ///
    /// [google.longrunning.Operation]: longrunning::model::Operation
    /// [google.longrunning.Operation.metadata]: longrunning::model::Operation::metadata
    /// [google.longrunning.Operation.response]: longrunning::model::Operation::result
    /// [google.spanner.admin.database.v1.Database]: crate::model::Database
    /// [google.spanner.admin.database.v1.RestoreDatabaseMetadata]: crate::model::RestoreDatabaseMetadata
    fn restore_database(
        &self,
        _req: crate::model::RestoreDatabaseRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<longrunning::model::Operation>> + Send
    {
        std::future::ready::<crate::Result<longrunning::model::Operation>>(Err(Error::other(
            "unimplemented",
        )))
    }

    /// Lists database [longrunning-operations][google.longrunning.Operation].
    /// A database operation has a name of the form
    /// `projects/<project>/instances/<instance>/databases/<database>/operations/<operation>`.
    /// The long-running operation
    /// [metadata][google.longrunning.Operation.metadata] field type
    /// `metadata.type_url` describes the type of the metadata. Operations returned
    /// include those that have completed/failed/canceled within the last 7 days,
    /// and pending operations.
    ///
    /// [google.longrunning.Operation]: longrunning::model::Operation
    /// [google.longrunning.Operation.metadata]: longrunning::model::Operation::metadata
    fn list_database_operations(
        &self,
        _req: crate::model::ListDatabaseOperationsRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<crate::model::ListDatabaseOperationsResponse>>
           + Send {
        std::future::ready::<crate::Result<crate::model::ListDatabaseOperationsResponse>>(Err(
            Error::other("unimplemented"),
        ))
    }

    /// Lists the backup [long-running operations][google.longrunning.Operation] in
    /// the given instance. A backup operation has a name of the form
    /// `projects/<project>/instances/<instance>/backups/<backup>/operations/<operation>`.
    /// The long-running operation
    /// [metadata][google.longrunning.Operation.metadata] field type
    /// `metadata.type_url` describes the type of the metadata. Operations returned
    /// include those that have completed/failed/canceled within the last 7 days,
    /// and pending operations. Operations returned are ordered by
    /// `operation.metadata.value.progress.start_time` in descending order starting
    /// from the most recently started operation.
    ///
    /// [google.longrunning.Operation]: longrunning::model::Operation
    /// [google.longrunning.Operation.metadata]: longrunning::model::Operation::metadata
    fn list_backup_operations(
        &self,
        _req: crate::model::ListBackupOperationsRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<crate::model::ListBackupOperationsResponse>>
           + Send {
        std::future::ready::<crate::Result<crate::model::ListBackupOperationsResponse>>(Err(
            Error::other("unimplemented"),
        ))
    }

    /// Lists Cloud Spanner database roles.
    fn list_database_roles(
        &self,
        _req: crate::model::ListDatabaseRolesRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<crate::model::ListDatabaseRolesResponse>> + Send
    {
        std::future::ready::<crate::Result<crate::model::ListDatabaseRolesResponse>>(Err(
            Error::other("unimplemented"),
        ))
    }

    /// Creates a new backup schedule.
    fn create_backup_schedule(
        &self,
        _req: crate::model::CreateBackupScheduleRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<crate::model::BackupSchedule>> + Send {
        std::future::ready::<crate::Result<crate::model::BackupSchedule>>(Err(Error::other(
            "unimplemented",
        )))
    }

    /// Gets backup schedule for the input schedule name.
    fn get_backup_schedule(
        &self,
        _req: crate::model::GetBackupScheduleRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<crate::model::BackupSchedule>> + Send {
        std::future::ready::<crate::Result<crate::model::BackupSchedule>>(Err(Error::other(
            "unimplemented",
        )))
    }

    /// Updates a backup schedule.
    fn update_backup_schedule(
        &self,
        _req: crate::model::UpdateBackupScheduleRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<crate::model::BackupSchedule>> + Send {
        std::future::ready::<crate::Result<crate::model::BackupSchedule>>(Err(Error::other(
            "unimplemented",
        )))
    }

    /// Deletes a backup schedule.
    fn delete_backup_schedule(
        &self,
        _req: crate::model::DeleteBackupScheduleRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<wkt::Empty>> + Send {
        std::future::ready::<crate::Result<wkt::Empty>>(Err(Error::other("unimplemented")))
    }

    /// Lists all the backup schedules for the database.
    fn list_backup_schedules(
        &self,
        _req: crate::model::ListBackupSchedulesRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<crate::model::ListBackupSchedulesResponse>> + Send
    {
        std::future::ready::<crate::Result<crate::model::ListBackupSchedulesResponse>>(Err(
            Error::other("unimplemented"),
        ))
    }

    /// Provides the [Operations][google.longrunning.Operations] service functionality in this service.
    ///
    /// [google.longrunning.Operations]: longrunning::traits::Operations
    fn list_operations(
        &self,
        _req: longrunning::model::ListOperationsRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<longrunning::model::ListOperationsResponse>>
           + Send {
        std::future::ready::<crate::Result<longrunning::model::ListOperationsResponse>>(Err(
            Error::other("unimplemented"),
        ))
    }

    /// Provides the [Operations][google.longrunning.Operations] service functionality in this service.
    ///
    /// [google.longrunning.Operations]: longrunning::traits::Operations
    fn get_operation(
        &self,
        _req: longrunning::model::GetOperationRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<longrunning::model::Operation>> + Send
    {
        std::future::ready::<crate::Result<longrunning::model::Operation>>(Err(Error::other(
            "unimplemented",
        )))
    }

    /// Provides the [Operations][google.longrunning.Operations] service functionality in this service.
    ///
    /// [google.longrunning.Operations]: longrunning::traits::Operations
    fn delete_operation(
        &self,
        _req: longrunning::model::DeleteOperationRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<wkt::Empty>> + Send {
        std::future::ready::<crate::Result<wkt::Empty>>(Err(Error::other("unimplemented")))
    }

    /// Provides the [Operations][google.longrunning.Operations] service functionality in this service.
    ///
    /// [google.longrunning.Operations]: longrunning::traits::Operations
    fn cancel_operation(
        &self,
        _req: longrunning::model::CancelOperationRequest,
        _options: gax::options::RequestOptions,
    ) -> impl std::future::Future<Output = crate::Result<wkt::Empty>> + Send {
        std::future::ready::<crate::Result<wkt::Empty>>(Err(Error::other("unimplemented")))
    }

    /// Returns the polling policy.
    fn get_polling_policy(
        &self,
        options: &gax::options::RequestOptions,
    ) -> std::sync::Arc<dyn gax::polling_policy::PollingPolicy>;

    /// Returns the polling backoff policy.
    fn get_polling_backoff_policy(
        &self,
        options: &gax::options::RequestOptions,
    ) -> std::sync::Arc<dyn gax::polling_backoff_policy::PollingBackoffPolicy>;
}