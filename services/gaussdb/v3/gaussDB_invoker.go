package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/gaussdb/v3/model"
)

type AddDatabasePermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddDatabasePermissionInvoker) Invoke() (*model.AddDatabasePermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddDatabasePermissionResponse), nil
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

type CancelGaussMySqlInstanceEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelGaussMySqlInstanceEipInvoker) Invoke() (*model.CancelGaussMySqlInstanceEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelGaussMySqlInstanceEipResponse), nil
	}
}

type CancelScheduleTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelScheduleTaskInvoker) Invoke() (*model.CancelScheduleTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelScheduleTaskResponse), nil
	}
}

type ChangeGaussMySqlInstanceSpecificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeGaussMySqlInstanceSpecificationInvoker) Invoke() (*model.ChangeGaussMySqlInstanceSpecificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeGaussMySqlInstanceSpecificationResponse), nil
	}
}

type ChangeGaussMySqlProxySpecificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeGaussMySqlProxySpecificationInvoker) Invoke() (*model.ChangeGaussMySqlProxySpecificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeGaussMySqlProxySpecificationResponse), nil
	}
}

type CheckResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckResourceInvoker) Invoke() (*model.CheckResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckResourceResponse), nil
	}
}

type CopyConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyConfigurationsInvoker) Invoke() (*model.CopyConfigurationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyConfigurationsResponse), nil
	}
}

type CopyInstanceConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyInstanceConfigurationsInvoker) Invoke() (*model.CopyInstanceConfigurationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyInstanceConfigurationsResponse), nil
	}
}

type CreateAccessControlInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAccessControlInvoker) Invoke() (*model.CreateAccessControlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAccessControlResponse), nil
	}
}

type CreateGaussMySqlBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGaussMySqlBackupInvoker) Invoke() (*model.CreateGaussMySqlBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGaussMySqlBackupResponse), nil
	}
}

type CreateGaussMySqlConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGaussMySqlConfigurationInvoker) Invoke() (*model.CreateGaussMySqlConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGaussMySqlConfigurationResponse), nil
	}
}

type CreateGaussMySqlDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGaussMySqlDatabaseInvoker) Invoke() (*model.CreateGaussMySqlDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGaussMySqlDatabaseResponse), nil
	}
}

type CreateGaussMySqlDatabaseUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGaussMySqlDatabaseUserInvoker) Invoke() (*model.CreateGaussMySqlDatabaseUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGaussMySqlDatabaseUserResponse), nil
	}
}

type CreateGaussMySqlInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGaussMySqlInstanceInvoker) Invoke() (*model.CreateGaussMySqlInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGaussMySqlInstanceResponse), nil
	}
}

type CreateGaussMySqlProxyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGaussMySqlProxyInvoker) Invoke() (*model.CreateGaussMySqlProxyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGaussMySqlProxyResponse), nil
	}
}

type CreateGaussMySqlReadonlyNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGaussMySqlReadonlyNodeInvoker) Invoke() (*model.CreateGaussMySqlReadonlyNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGaussMySqlReadonlyNodeResponse), nil
	}
}

type CreateGaussMysqlDnsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGaussMysqlDnsInvoker) Invoke() (*model.CreateGaussMysqlDnsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGaussMysqlDnsResponse), nil
	}
}

type CreateRestoreTablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRestoreTablesInvoker) Invoke() (*model.CreateRestoreTablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRestoreTablesResponse), nil
	}
}

type DeleteDatabasePermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDatabasePermissionInvoker) Invoke() (*model.DeleteDatabasePermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDatabasePermissionResponse), nil
	}
}

type DeleteGaussMySqlBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGaussMySqlBackupInvoker) Invoke() (*model.DeleteGaussMySqlBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGaussMySqlBackupResponse), nil
	}
}

type DeleteGaussMySqlConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGaussMySqlConfigurationInvoker) Invoke() (*model.DeleteGaussMySqlConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGaussMySqlConfigurationResponse), nil
	}
}

type DeleteGaussMySqlDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGaussMySqlDatabaseInvoker) Invoke() (*model.DeleteGaussMySqlDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGaussMySqlDatabaseResponse), nil
	}
}

type DeleteGaussMySqlDatabaseUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGaussMySqlDatabaseUserInvoker) Invoke() (*model.DeleteGaussMySqlDatabaseUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGaussMySqlDatabaseUserResponse), nil
	}
}

type DeleteGaussMySqlInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGaussMySqlInstanceInvoker) Invoke() (*model.DeleteGaussMySqlInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGaussMySqlInstanceResponse), nil
	}
}

type DeleteGaussMySqlProxyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGaussMySqlProxyInvoker) Invoke() (*model.DeleteGaussMySqlProxyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGaussMySqlProxyResponse), nil
	}
}

type DeleteGaussMySqlReadonlyNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGaussMySqlReadonlyNodeInvoker) Invoke() (*model.DeleteGaussMySqlReadonlyNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGaussMySqlReadonlyNodeResponse), nil
	}
}

type DeleteScheduleTasKInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScheduleTasKInvoker) Invoke() (*model.DeleteScheduleTasKResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScheduleTasKResponse), nil
	}
}

type DeleteTaskRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTaskRecordInvoker) Invoke() (*model.DeleteTaskRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTaskRecordResponse), nil
	}
}

type DescribeBackupEncryptStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribeBackupEncryptStatusInvoker) Invoke() (*model.DescribeBackupEncryptStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribeBackupEncryptStatusResponse), nil
	}
}

type ExpandGaussMySqlInstanceVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandGaussMySqlInstanceVolumeInvoker) Invoke() (*model.ExpandGaussMySqlInstanceVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandGaussMySqlInstanceVolumeResponse), nil
	}
}

type ExpandGaussMySqlProxyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandGaussMySqlProxyInvoker) Invoke() (*model.ExpandGaussMySqlProxyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandGaussMySqlProxyResponse), nil
	}
}

type InvokeGaussMySqlInstanceSwitchOverInvoker struct {
	*invoker.BaseInvoker
}

func (i *InvokeGaussMySqlInstanceSwitchOverInvoker) Invoke() (*model.InvokeGaussMySqlInstanceSwitchOverResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InvokeGaussMySqlInstanceSwitchOverResponse), nil
	}
}

type ListAuditLogDownloadLinkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditLogDownloadLinkInvoker) Invoke() (*model.ListAuditLogDownloadLinkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditLogDownloadLinkResponse), nil
	}
}

type ListConfigurationsDifferencesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConfigurationsDifferencesInvoker) Invoke() (*model.ListConfigurationsDifferencesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigurationsDifferencesResponse), nil
	}
}

type ListConfigurationsInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConfigurationsInstancesInvoker) Invoke() (*model.ListConfigurationsInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigurationsInstancesResponse), nil
	}
}

type ListEnterpriseProjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnterpriseProjectsInvoker) Invoke() (*model.ListEnterpriseProjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnterpriseProjectsResponse), nil
	}
}

type ListGaussMySqlConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGaussMySqlConfigurationsInvoker) Invoke() (*model.ListGaussMySqlConfigurationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGaussMySqlConfigurationsResponse), nil
	}
}

type ListGaussMySqlDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGaussMySqlDatabaseInvoker) Invoke() (*model.ListGaussMySqlDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGaussMySqlDatabaseResponse), nil
	}
}

type ListGaussMySqlDatabaseCharsetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGaussMySqlDatabaseCharsetsInvoker) Invoke() (*model.ListGaussMySqlDatabaseCharsetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGaussMySqlDatabaseCharsetsResponse), nil
	}
}

type ListGaussMySqlDatabaseUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGaussMySqlDatabaseUserInvoker) Invoke() (*model.ListGaussMySqlDatabaseUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGaussMySqlDatabaseUserResponse), nil
	}
}

type ListGaussMySqlDedicatedResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGaussMySqlDedicatedResourcesInvoker) Invoke() (*model.ListGaussMySqlDedicatedResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGaussMySqlDedicatedResourcesResponse), nil
	}
}

type ListGaussMySqlInstanceDetailInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGaussMySqlInstanceDetailInfoInvoker) Invoke() (*model.ListGaussMySqlInstanceDetailInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGaussMySqlInstanceDetailInfoResponse), nil
	}
}

type ListGaussMySqlInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGaussMySqlInstancesInvoker) Invoke() (*model.ListGaussMySqlInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGaussMySqlInstancesResponse), nil
	}
}

type ListImmediateJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImmediateJobsInvoker) Invoke() (*model.ListImmediateJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImmediateJobsResponse), nil
	}
}

type ListInstanceConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceConfigurationsInvoker) Invoke() (*model.ListInstanceConfigurationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceConfigurationsResponse), nil
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

type ListLtsErrorLogDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLtsErrorLogDetailsInvoker) Invoke() (*model.ListLtsErrorLogDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLtsErrorLogDetailsResponse), nil
	}
}

type ListLtsSlowlogDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLtsSlowlogDetailsInvoker) Invoke() (*model.ListLtsSlowlogDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLtsSlowlogDetailsResponse), nil
	}
}

type ListModifyHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListModifyHistoryInvoker) Invoke() (*model.ListModifyHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListModifyHistoryResponse), nil
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

type ListScheduleJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScheduleJobsInvoker) Invoke() (*model.ListScheduleJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScheduleJobsResponse), nil
	}
}

type ModifyBackupEncryptStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyBackupEncryptStatusInvoker) Invoke() (*model.ModifyBackupEncryptStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyBackupEncryptStatusResponse), nil
	}
}

type ModifyGaussMySqlProxyRouteModeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyGaussMySqlProxyRouteModeInvoker) Invoke() (*model.ModifyGaussMySqlProxyRouteModeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyGaussMySqlProxyRouteModeResponse), nil
	}
}

type ModifyGaussMysqlDnsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyGaussMysqlDnsInvoker) Invoke() (*model.ModifyGaussMysqlDnsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyGaussMysqlDnsResponse), nil
	}
}

type ResetGaussMySqlDatabasePasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetGaussMySqlDatabasePasswordInvoker) Invoke() (*model.ResetGaussMySqlDatabasePasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetGaussMySqlDatabasePasswordResponse), nil
	}
}

type ResetGaussMySqlPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetGaussMySqlPasswordInvoker) Invoke() (*model.ResetGaussMySqlPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetGaussMySqlPasswordResponse), nil
	}
}

type RestartGaussMySqlInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartGaussMySqlInstanceInvoker) Invoke() (*model.RestartGaussMySqlInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartGaussMySqlInstanceResponse), nil
	}
}

type RestartGaussMySqlNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartGaussMySqlNodeInvoker) Invoke() (*model.RestartGaussMySqlNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartGaussMySqlNodeResponse), nil
	}
}

type RestoreOldInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreOldInstanceInvoker) Invoke() (*model.RestoreOldInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreOldInstanceResponse), nil
	}
}

type SetGaussMySqlProxyWeightInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetGaussMySqlProxyWeightInvoker) Invoke() (*model.SetGaussMySqlProxyWeightResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetGaussMySqlProxyWeightResponse), nil
	}
}

type SetGaussMySqlQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetGaussMySqlQuotasInvoker) Invoke() (*model.SetGaussMySqlQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetGaussMySqlQuotasResponse), nil
	}
}

type ShowAuditLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuditLogInvoker) Invoke() (*model.ShowAuditLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditLogResponse), nil
	}
}

type ShowAutoScalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoScalingPolicyInvoker) Invoke() (*model.ShowAutoScalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAutoScalingPolicyResponse), nil
	}
}

type ShowBackupRestoreTimeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBackupRestoreTimeInvoker) Invoke() (*model.ShowBackupRestoreTimeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBackupRestoreTimeResponse), nil
	}
}

type ShowDedicatedResourceInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDedicatedResourceInfoInvoker) Invoke() (*model.ShowDedicatedResourceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDedicatedResourceInfoResponse), nil
	}
}

type ShowGaussMySqlBackupListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlBackupListInvoker) Invoke() (*model.ShowGaussMySqlBackupListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlBackupListResponse), nil
	}
}

type ShowGaussMySqlBackupPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlBackupPolicyInvoker) Invoke() (*model.ShowGaussMySqlBackupPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlBackupPolicyResponse), nil
	}
}

type ShowGaussMySqlConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlConfigurationInvoker) Invoke() (*model.ShowGaussMySqlConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlConfigurationResponse), nil
	}
}

type ShowGaussMySqlEngineVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlEngineVersionInvoker) Invoke() (*model.ShowGaussMySqlEngineVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlEngineVersionResponse), nil
	}
}

type ShowGaussMySqlFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlFlavorsInvoker) Invoke() (*model.ShowGaussMySqlFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlFlavorsResponse), nil
	}
}

type ShowGaussMySqlIncrementalBackupListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlIncrementalBackupListInvoker) Invoke() (*model.ShowGaussMySqlIncrementalBackupListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlIncrementalBackupListResponse), nil
	}
}

type ShowGaussMySqlInstanceInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlInstanceInfoInvoker) Invoke() (*model.ShowGaussMySqlInstanceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlInstanceInfoResponse), nil
	}
}

type ShowGaussMySqlJobInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlJobInfoInvoker) Invoke() (*model.ShowGaussMySqlJobInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlJobInfoResponse), nil
	}
}

type ShowGaussMySqlProjectQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlProjectQuotasInvoker) Invoke() (*model.ShowGaussMySqlProjectQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlProjectQuotasResponse), nil
	}
}

type ShowGaussMySqlProxyFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlProxyFlavorsInvoker) Invoke() (*model.ShowGaussMySqlProxyFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlProxyFlavorsResponse), nil
	}
}

type ShowGaussMySqlProxyListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlProxyListInvoker) Invoke() (*model.ShowGaussMySqlProxyListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlProxyListResponse), nil
	}
}

type ShowGaussMySqlQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlQuotasInvoker) Invoke() (*model.ShowGaussMySqlQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlQuotasResponse), nil
	}
}

type ShowInstanceDatabaseVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceDatabaseVersionInvoker) Invoke() (*model.ShowInstanceDatabaseVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceDatabaseVersionResponse), nil
	}
}

type ShowInstanceMonitorExtendInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceMonitorExtendInvoker) Invoke() (*model.ShowInstanceMonitorExtendResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceMonitorExtendResponse), nil
	}
}

type ShowIntelligentDiagnosisAbnormalCountOfInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIntelligentDiagnosisAbnormalCountOfInstancesInvoker) Invoke() (*model.ShowIntelligentDiagnosisAbnormalCountOfInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIntelligentDiagnosisAbnormalCountOfInstancesResponse), nil
	}
}

type ShowIntelligentDiagnosisInstanceInfosPerMetricInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIntelligentDiagnosisInstanceInfosPerMetricInvoker) Invoke() (*model.ShowIntelligentDiagnosisInstanceInfosPerMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIntelligentDiagnosisInstanceInfosPerMetricResponse), nil
	}
}

type ShowRestoreTablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRestoreTablesInvoker) Invoke() (*model.ShowRestoreTablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRestoreTablesResponse), nil
	}
}

type ShrinkGaussMySqlProxyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShrinkGaussMySqlProxyInvoker) Invoke() (*model.ShrinkGaussMySqlProxyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShrinkGaussMySqlProxyResponse), nil
	}
}

type SwitchAccessControlInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchAccessControlInvoker) Invoke() (*model.SwitchAccessControlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchAccessControlResponse), nil
	}
}

type SwitchGaussMySqlConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchGaussMySqlConfigurationInvoker) Invoke() (*model.SwitchGaussMySqlConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchGaussMySqlConfigurationResponse), nil
	}
}

type SwitchGaussMySqlInstanceSslInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchGaussMySqlInstanceSslInvoker) Invoke() (*model.SwitchGaussMySqlInstanceSslResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchGaussMySqlInstanceSslResponse), nil
	}
}

type UpdateAuditLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAuditLogInvoker) Invoke() (*model.UpdateAuditLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuditLogResponse), nil
	}
}

type UpdateAutoScalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAutoScalingPolicyInvoker) Invoke() (*model.UpdateAutoScalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAutoScalingPolicyResponse), nil
	}
}

type UpdateBackupOffsitePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBackupOffsitePolicyInvoker) Invoke() (*model.UpdateBackupOffsitePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBackupOffsitePolicyResponse), nil
	}
}

type UpdateGaussMySqlBackupPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGaussMySqlBackupPolicyInvoker) Invoke() (*model.UpdateGaussMySqlBackupPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGaussMySqlBackupPolicyResponse), nil
	}
}

type UpdateGaussMySqlConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGaussMySqlConfigurationInvoker) Invoke() (*model.UpdateGaussMySqlConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGaussMySqlConfigurationResponse), nil
	}
}

type UpdateGaussMySqlDatabaseCommentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGaussMySqlDatabaseCommentInvoker) Invoke() (*model.UpdateGaussMySqlDatabaseCommentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGaussMySqlDatabaseCommentResponse), nil
	}
}

type UpdateGaussMySqlDatabaseUserCommentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGaussMySqlDatabaseUserCommentInvoker) Invoke() (*model.UpdateGaussMySqlDatabaseUserCommentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGaussMySqlDatabaseUserCommentResponse), nil
	}
}

type UpdateGaussMySqlInstanceAliasInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGaussMySqlInstanceAliasInvoker) Invoke() (*model.UpdateGaussMySqlInstanceAliasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGaussMySqlInstanceAliasResponse), nil
	}
}

type UpdateGaussMySqlInstanceEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGaussMySqlInstanceEipInvoker) Invoke() (*model.UpdateGaussMySqlInstanceEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGaussMySqlInstanceEipResponse), nil
	}
}

type UpdateGaussMySqlInstanceInternalIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGaussMySqlInstanceInternalIpInvoker) Invoke() (*model.UpdateGaussMySqlInstanceInternalIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGaussMySqlInstanceInternalIpResponse), nil
	}
}

type UpdateGaussMySqlInstanceNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGaussMySqlInstanceNameInvoker) Invoke() (*model.UpdateGaussMySqlInstanceNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGaussMySqlInstanceNameResponse), nil
	}
}

type UpdateGaussMySqlInstanceOpsWindowInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGaussMySqlInstanceOpsWindowInvoker) Invoke() (*model.UpdateGaussMySqlInstanceOpsWindowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGaussMySqlInstanceOpsWindowResponse), nil
	}
}

type UpdateGaussMySqlInstancePortInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGaussMySqlInstancePortInvoker) Invoke() (*model.UpdateGaussMySqlInstancePortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGaussMySqlInstancePortResponse), nil
	}
}

type UpdateGaussMySqlInstanceSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGaussMySqlInstanceSecurityGroupInvoker) Invoke() (*model.UpdateGaussMySqlInstanceSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGaussMySqlInstanceSecurityGroupResponse), nil
	}
}

type UpdateGaussMySqlQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGaussMySqlQuotasInvoker) Invoke() (*model.UpdateGaussMySqlQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGaussMySqlQuotasResponse), nil
	}
}

type UpdateInstanceConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceConfigurationsInvoker) Invoke() (*model.UpdateInstanceConfigurationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceConfigurationsResponse), nil
	}
}

type UpdateInstanceMonitorInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceMonitorInvoker) Invoke() (*model.UpdateInstanceMonitorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceMonitorResponse), nil
	}
}

type UpdateProxyConnectionPoolTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProxyConnectionPoolTypeInvoker) Invoke() (*model.UpdateProxyConnectionPoolTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProxyConnectionPoolTypeResponse), nil
	}
}

type UpdateProxyNewConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProxyNewConfigurationsInvoker) Invoke() (*model.UpdateProxyNewConfigurationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProxyNewConfigurationsResponse), nil
	}
}

type UpdateProxyPortInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProxyPortInvoker) Invoke() (*model.UpdateProxyPortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProxyPortResponse), nil
	}
}

type UpdateProxySessionConsistenceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProxySessionConsistenceInvoker) Invoke() (*model.UpdateProxySessionConsistenceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProxySessionConsistenceResponse), nil
	}
}

type UpdateTransactionSplitStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTransactionSplitStatusInvoker) Invoke() (*model.UpdateTransactionSplitStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTransactionSplitStatusResponse), nil
	}
}

type UpgradeGaussMySqlInstanceDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeGaussMySqlInstanceDatabaseInvoker) Invoke() (*model.UpgradeGaussMySqlInstanceDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeGaussMySqlInstanceDatabaseResponse), nil
	}
}

type DeleteSqlFilterRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSqlFilterRuleInvoker) Invoke() (*model.DeleteSqlFilterRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSqlFilterRuleResponse), nil
	}
}

type SetSqlFilterRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetSqlFilterRuleInvoker) Invoke() (*model.SetSqlFilterRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetSqlFilterRuleResponse), nil
	}
}

type ShowSqlFilterControlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlFilterControlInvoker) Invoke() (*model.ShowSqlFilterControlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlFilterControlResponse), nil
	}
}

type ShowSqlFilterRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlFilterRuleInvoker) Invoke() (*model.ShowSqlFilterRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlFilterRuleResponse), nil
	}
}

type UpdateSqlFilterControlInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSqlFilterControlInvoker) Invoke() (*model.UpdateSqlFilterControlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSqlFilterControlResponse), nil
	}
}
