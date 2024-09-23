package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dds/v3/model"
)

type AddReadonlyNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddReadonlyNodeInvoker) Invoke() (*model.AddReadonlyNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddReadonlyNodeResponse), nil
	}
}

type AddShardingNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddShardingNodeInvoker) Invoke() (*model.AddShardingNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddShardingNodeResponse), nil
	}
}

type AttachEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachEipInvoker) Invoke() (*model.AttachEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachEipResponse), nil
	}
}

type AttachInternalIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachInternalIpInvoker) Invoke() (*model.AttachInternalIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachInternalIpResponse), nil
	}
}

type BatchTagActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchTagActionInvoker) Invoke() (*model.BatchTagActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchTagActionResponse), nil
	}
}

type CancelEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelEipInvoker) Invoke() (*model.CancelEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelEipResponse), nil
	}
}

type ChangeOpsWindowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeOpsWindowInvoker) Invoke() (*model.ChangeOpsWindowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeOpsWindowResponse), nil
	}
}

type CheckPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckPasswordInvoker) Invoke() (*model.CheckPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckPasswordResponse), nil
	}
}

type CheckWeakPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckWeakPasswordInvoker) Invoke() (*model.CheckWeakPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckWeakPasswordResponse), nil
	}
}

type CompareConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CompareConfigurationInvoker) Invoke() (*model.CompareConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CompareConfigurationResponse), nil
	}
}

type CopyConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyConfigurationInvoker) Invoke() (*model.CopyConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyConfigurationResponse), nil
	}
}

type CreateConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConfigurationInvoker) Invoke() (*model.CreateConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConfigurationResponse), nil
	}
}

type CreateDatabaseRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatabaseRoleInvoker) Invoke() (*model.CreateDatabaseRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatabaseRoleResponse), nil
	}
}

type CreateDatabaseUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatabaseUserInvoker) Invoke() (*model.CreateDatabaseUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatabaseUserResponse), nil
	}
}

type CreateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceInvoker) Invoke() (*model.CreateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceResponse), nil
	}
}

type CreateIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateIpInvoker) Invoke() (*model.CreateIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIpResponse), nil
	}
}

type CreateKillOpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateKillOpRuleInvoker) Invoke() (*model.CreateKillOpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateKillOpRuleResponse), nil
	}
}

type CreateManualBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateManualBackupInvoker) Invoke() (*model.CreateManualBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateManualBackupResponse), nil
	}
}

type DeleteAuditLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAuditLogInvoker) Invoke() (*model.DeleteAuditLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAuditLogResponse), nil
	}
}

type DeleteConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConfigurationInvoker) Invoke() (*model.DeleteConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConfigurationResponse), nil
	}
}

type DeleteDatabaseRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDatabaseRoleInvoker) Invoke() (*model.DeleteDatabaseRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDatabaseRoleResponse), nil
	}
}

type DeleteDatabaseUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDatabaseUserInvoker) Invoke() (*model.DeleteDatabaseUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDatabaseUserResponse), nil
	}
}

type DeleteInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceInvoker) Invoke() (*model.DeleteInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceResponse), nil
	}
}

type DeleteKillOpRuleListInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteKillOpRuleListInvoker) Invoke() (*model.DeleteKillOpRuleListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteKillOpRuleListResponse), nil
	}
}

type DeleteLtsConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLtsConfigInvoker) Invoke() (*model.DeleteLtsConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLtsConfigResponse), nil
	}
}

type DeleteManualBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteManualBackupInvoker) Invoke() (*model.DeleteManualBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteManualBackupResponse), nil
	}
}

type DeleteReadonlyNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteReadonlyNodeInvoker) Invoke() (*model.DeleteReadonlyNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteReadonlyNodeResponse), nil
	}
}

type DeleteSessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSessionInvoker) Invoke() (*model.DeleteSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSessionResponse), nil
	}
}

type DownloadErrorlogInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadErrorlogInvoker) Invoke() (*model.DownloadErrorlogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadErrorlogResponse), nil
	}
}

type DownloadSlowlogInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadSlowlogInvoker) Invoke() (*model.DownloadSlowlogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadSlowlogResponse), nil
	}
}

type ExpandReplicasetNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandReplicasetNodeInvoker) Invoke() (*model.ExpandReplicasetNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandReplicasetNodeResponse), nil
	}
}

type ListAppliedInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppliedInstancesInvoker) Invoke() (*model.ListAppliedInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppliedInstancesResponse), nil
	}
}

type ListAuditlogLinksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditlogLinksInvoker) Invoke() (*model.ListAuditlogLinksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditlogLinksResponse), nil
	}
}

type ListAuditlogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditlogsInvoker) Invoke() (*model.ListAuditlogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditlogsResponse), nil
	}
}

type ListAz2MigrateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAz2MigrateInvoker) Invoke() (*model.ListAz2MigrateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAz2MigrateResponse), nil
	}
}

type ListBackupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackupsInvoker) Invoke() (*model.ListBackupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackupsResponse), nil
	}
}

type ListConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConfigurationsInvoker) Invoke() (*model.ListConfigurationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigurationsResponse), nil
	}
}

type ListDatabaseRolesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabaseRolesInvoker) Invoke() (*model.ListDatabaseRolesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabaseRolesResponse), nil
	}
}

type ListDatabaseUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabaseUsersInvoker) Invoke() (*model.ListDatabaseUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabaseUsersResponse), nil
	}
}

type ListDatabasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabasesInvoker) Invoke() (*model.ListDatabasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabasesResponse), nil
	}
}

type ListDatastoreVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatastoreVersionsInvoker) Invoke() (*model.ListDatastoreVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatastoreVersionsResponse), nil
	}
}

type ListErrorLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListErrorLogsInvoker) Invoke() (*model.ListErrorLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListErrorLogsResponse), nil
	}
}

type ListFlavorInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlavorInfosInvoker) Invoke() (*model.ListFlavorInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlavorInfosResponse), nil
	}
}

type ListFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlavorsInvoker) Invoke() (*model.ListFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlavorsResponse), nil
	}
}

type ListInstanceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceTagsInvoker) Invoke() (*model.ListInstanceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceTagsResponse), nil
	}
}

type ListInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesInvoker) Invoke() (*model.ListInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesResponse), nil
	}
}

type ListInstancesByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesByTagsInvoker) Invoke() (*model.ListInstancesByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesByTagsResponse), nil
	}
}

type ListLtsConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLtsConfigsInvoker) Invoke() (*model.ListLtsConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLtsConfigsResponse), nil
	}
}

type ListLtsErrorLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLtsErrorLogsInvoker) Invoke() (*model.ListLtsErrorLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLtsErrorLogsResponse), nil
	}
}

type ListLtsSlowLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLtsSlowLogsInvoker) Invoke() (*model.ListLtsSlowLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLtsSlowLogsResponse), nil
	}
}

type ListProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectTagsInvoker) Invoke() (*model.ListProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectTagsResponse), nil
	}
}

type ListRecycleInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecycleInstancesInvoker) Invoke() (*model.ListRecycleInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecycleInstancesResponse), nil
	}
}

type ListRestoreCollectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRestoreCollectionsInvoker) Invoke() (*model.ListRestoreCollectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRestoreCollectionsResponse), nil
	}
}

type ListRestoreDatabasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRestoreDatabasesInvoker) Invoke() (*model.ListRestoreDatabasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRestoreDatabasesResponse), nil
	}
}

type ListRestoreTimesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRestoreTimesInvoker) Invoke() (*model.ListRestoreTimesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRestoreTimesResponse), nil
	}
}

type ListSessionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSessionsInvoker) Invoke() (*model.ListSessionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSessionsResponse), nil
	}
}

type ListSlowLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSlowLogsInvoker) Invoke() (*model.ListSlowLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSlowLogsResponse), nil
	}
}

type ListSslCertDownloadAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSslCertDownloadAddressInvoker) Invoke() (*model.ListSslCertDownloadAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSslCertDownloadAddressResponse), nil
	}
}

type ListStorageTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStorageTypeInvoker) Invoke() (*model.ListStorageTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStorageTypeResponse), nil
	}
}

type ListTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTasksInvoker) Invoke() (*model.ListTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTasksResponse), nil
	}
}

type MigrateAzInvoker struct {
	*invoker.BaseInvoker
}

func (i *MigrateAzInvoker) Invoke() (*model.MigrateAzResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.MigrateAzResponse), nil
	}
}

type ResetConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetConfigurationInvoker) Invoke() (*model.ResetConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetConfigurationResponse), nil
	}
}

type ResetPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetPasswordInvoker) Invoke() (*model.ResetPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetPasswordResponse), nil
	}
}

type ResizeInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeInstanceInvoker) Invoke() (*model.ResizeInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeInstanceResponse), nil
	}
}

type ResizeInstanceVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeInstanceVolumeInvoker) Invoke() (*model.ResizeInstanceVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeInstanceVolumeResponse), nil
	}
}

type RestartInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartInstanceInvoker) Invoke() (*model.RestartInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartInstanceResponse), nil
	}
}

type RestoreInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreInstanceInvoker) Invoke() (*model.RestoreInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreInstanceResponse), nil
	}
}

type RestoreInstanceFromCollectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreInstanceFromCollectionInvoker) Invoke() (*model.RestoreInstanceFromCollectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreInstanceFromCollectionResponse), nil
	}
}

type RestoreNewInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreNewInstanceInvoker) Invoke() (*model.RestoreNewInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreNewInstanceResponse), nil
	}
}

type SetAuditlogPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetAuditlogPolicyInvoker) Invoke() (*model.SetAuditlogPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetAuditlogPolicyResponse), nil
	}
}

type SetBackupPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetBackupPolicyInvoker) Invoke() (*model.SetBackupPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetBackupPolicyResponse), nil
	}
}

type SetBalancerSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetBalancerSwitchInvoker) Invoke() (*model.SetBalancerSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetBalancerSwitchResponse), nil
	}
}

type SetBalancerWindowInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetBalancerWindowInvoker) Invoke() (*model.SetBalancerWindowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetBalancerWindowResponse), nil
	}
}

type SetRecyclePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetRecyclePolicyInvoker) Invoke() (*model.SetRecyclePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetRecyclePolicyResponse), nil
	}
}

type ShowAuditlogPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuditlogPolicyInvoker) Invoke() (*model.ShowAuditlogPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditlogPolicyResponse), nil
	}
}

type ShowBackupDownloadLinkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBackupDownloadLinkInvoker) Invoke() (*model.ShowBackupDownloadLinkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBackupDownloadLinkResponse), nil
	}
}

type ShowBackupPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBackupPolicyInvoker) Invoke() (*model.ShowBackupPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBackupPolicyResponse), nil
	}
}

type ShowClientNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClientNetworkInvoker) Invoke() (*model.ShowClientNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClientNetworkResponse), nil
	}
}

type ShowConfigurationAppliedHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConfigurationAppliedHistoryInvoker) Invoke() (*model.ShowConfigurationAppliedHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConfigurationAppliedHistoryResponse), nil
	}
}

type ShowConfigurationModifyHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConfigurationModifyHistoryInvoker) Invoke() (*model.ShowConfigurationModifyHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConfigurationModifyHistoryResponse), nil
	}
}

type ShowConfigurationParameterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConfigurationParameterInvoker) Invoke() (*model.ShowConfigurationParameterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConfigurationParameterResponse), nil
	}
}

type ShowConnectionStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConnectionStatisticsInvoker) Invoke() (*model.ShowConnectionStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConnectionStatisticsResponse), nil
	}
}

type ShowDiskUsageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDiskUsageInvoker) Invoke() (*model.ShowDiskUsageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDiskUsageResponse), nil
	}
}

type ShowEntityConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEntityConfigurationInvoker) Invoke() (*model.ShowEntityConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEntityConfigurationResponse), nil
	}
}

type ShowJobDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobDetailInvoker) Invoke() (*model.ShowJobDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobDetailResponse), nil
	}
}

type ShowKillOpRuleRuleListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKillOpRuleRuleListInvoker) Invoke() (*model.ShowKillOpRuleRuleListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKillOpRuleRuleListResponse), nil
	}
}

type ShowQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotasInvoker) Invoke() (*model.ShowQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotasResponse), nil
	}
}

type ShowRecyclePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecyclePolicyInvoker) Invoke() (*model.ShowRecyclePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecyclePolicyResponse), nil
	}
}

type ShowReplSetNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowReplSetNameInvoker) Invoke() (*model.ShowReplSetNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowReplSetNameResponse), nil
	}
}

type ShowSecondLevelMonitoringStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecondLevelMonitoringStatusInvoker) Invoke() (*model.ShowSecondLevelMonitoringStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecondLevelMonitoringStatusResponse), nil
	}
}

type ShowShardingBalancerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowShardingBalancerInvoker) Invoke() (*model.ShowShardingBalancerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowShardingBalancerResponse), nil
	}
}

type ShowSlowlogDesensitizationSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSlowlogDesensitizationSwitchInvoker) Invoke() (*model.ShowSlowlogDesensitizationSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSlowlogDesensitizationSwitchResponse), nil
	}
}

type ShowUpgradeDurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUpgradeDurationInvoker) Invoke() (*model.ShowUpgradeDurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUpgradeDurationResponse), nil
	}
}

type ShrinkInstanceNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShrinkInstanceNodesInvoker) Invoke() (*model.ShrinkInstanceNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShrinkInstanceNodesResponse), nil
	}
}

type StopBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopBackupInvoker) Invoke() (*model.StopBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopBackupResponse), nil
	}
}

type SwitchConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchConfigurationInvoker) Invoke() (*model.SwitchConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchConfigurationResponse), nil
	}
}

type SwitchInstancePrimaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchInstancePrimaryInvoker) Invoke() (*model.SwitchInstancePrimaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchInstancePrimaryResponse), nil
	}
}

type SwitchSecondLevelMonitoringInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchSecondLevelMonitoringInvoker) Invoke() (*model.SwitchSecondLevelMonitoringResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchSecondLevelMonitoringResponse), nil
	}
}

type SwitchSlowlogDesensitizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchSlowlogDesensitizationInvoker) Invoke() (*model.SwitchSlowlogDesensitizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchSlowlogDesensitizationResponse), nil
	}
}

type SwitchSslInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchSslInvoker) Invoke() (*model.SwitchSslResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchSslResponse), nil
	}
}

type SwitchoverReplicaSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchoverReplicaSetInvoker) Invoke() (*model.SwitchoverReplicaSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchoverReplicaSetResponse), nil
	}
}

type UpdateClientNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClientNetworkInvoker) Invoke() (*model.UpdateClientNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClientNetworkResponse), nil
	}
}

type UpdateConfigurationParameterInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConfigurationParameterInvoker) Invoke() (*model.UpdateConfigurationParameterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateConfigurationParameterResponse), nil
	}
}

type UpdateEntityConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEntityConfigurationInvoker) Invoke() (*model.UpdateEntityConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEntityConfigurationResponse), nil
	}
}

type UpdateInstanceNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceNameInvoker) Invoke() (*model.UpdateInstanceNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceNameResponse), nil
	}
}

type UpdateInstancePortInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstancePortInvoker) Invoke() (*model.UpdateInstancePortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstancePortResponse), nil
	}
}

type UpdateInstanceRemarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceRemarkInvoker) Invoke() (*model.UpdateInstanceRemarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceRemarkResponse), nil
	}
}

type UpdateKillOpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateKillOpRuleInvoker) Invoke() (*model.UpdateKillOpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateKillOpRuleResponse), nil
	}
}

type UpdateLtsConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLtsConfigInvoker) Invoke() (*model.UpdateLtsConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLtsConfigResponse), nil
	}
}

type UpdateReplSetNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateReplSetNameInvoker) Invoke() (*model.UpdateReplSetNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateReplSetNameResponse), nil
	}
}

type UpdateSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecurityGroupInvoker) Invoke() (*model.UpdateSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecurityGroupResponse), nil
	}
}

type UpgradeDatabaseVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeDatabaseVersionInvoker) Invoke() (*model.UpgradeDatabaseVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeDatabaseVersionResponse), nil
	}
}

type ListApiVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionInvoker) Invoke() (*model.ListApiVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiVersionResponse), nil
	}
}

type ShowApiVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApiVersionInvoker) Invoke() (*model.ShowApiVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApiVersionResponse), nil
	}
}
