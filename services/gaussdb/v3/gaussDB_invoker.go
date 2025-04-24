package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/gaussdb/v3/model"
)

type AddDatabasePermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddDatabasePermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *BatchTagActionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CancelGaussMySqlInstanceEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CancelScheduleTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ChangeGaussMySqlInstanceSpecificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ChangeGaussMySqlProxySpecificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CheckResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CopyConfigurationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CopyInstanceConfigurationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateAccessControlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateGaussMySqlBackupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateGaussMySqlConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateGaussMySqlDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateGaussMySqlDatabaseUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateGaussMySqlInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateGaussMySqlProxyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateGaussMySqlReadonlyNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateGaussMysqlDnsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGaussMysqlDnsInvoker) Invoke() (*model.CreateGaussMysqlDnsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGaussMysqlDnsResponse), nil
	}
}

type CreateLtsConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLtsConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLtsConfigsInvoker) Invoke() (*model.CreateLtsConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLtsConfigsResponse), nil
	}
}

type CreateRestoreTablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRestoreTablesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteDatabasePermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteGaussMySqlBackupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteGaussMySqlConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteGaussMySqlDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteGaussMySqlDatabaseUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteGaussMySqlInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteGaussMySqlProxyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteGaussMySqlReadonlyNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGaussMySqlReadonlyNodeInvoker) Invoke() (*model.DeleteGaussMySqlReadonlyNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGaussMySqlReadonlyNodeResponse), nil
	}
}

type DeleteLtsConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLtsConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLtsConfigsInvoker) Invoke() (*model.DeleteLtsConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLtsConfigsResponse), nil
	}
}

type DeleteScheduleTasKInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScheduleTasKInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteScheduleTasKInvoker) Invoke() (*model.DeleteScheduleTasKResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScheduleTasKResponse), nil
	}
}

type DeleteSqlFilterRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSqlFilterRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSqlFilterRuleInvoker) Invoke() (*model.DeleteSqlFilterRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSqlFilterRuleResponse), nil
	}
}

type DeleteTaskRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTaskRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTaskRecordInvoker) Invoke() (*model.DeleteTaskRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTaskRecordResponse), nil
	}
}

type DeleteTaurusDbNodeProcessesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTaurusDbNodeProcessesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTaurusDbNodeProcessesInvoker) Invoke() (*model.DeleteTaurusDbNodeProcessesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTaurusDbNodeProcessesResponse), nil
	}
}

type DescribeBackupEncryptStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribeBackupEncryptStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DescribeBackupEncryptStatusInvoker) Invoke() (*model.DescribeBackupEncryptStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribeBackupEncryptStatusResponse), nil
	}
}

type DownloadSlowLogFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadSlowLogFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadSlowLogFileInvoker) Invoke() (*model.DownloadSlowLogFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadSlowLogFileResponse), nil
	}
}

type ExpandGaussMySqlInstanceVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandGaussMySqlInstanceVolumeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ExpandGaussMySqlProxyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *InvokeGaussMySqlInstanceSwitchOverInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListAuditLogDownloadLinkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListConfigurationsDifferencesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListConfigurationsInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListEnterpriseProjectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListGaussMySqlConfigurationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListGaussMySqlDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListGaussMySqlDatabaseCharsetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListGaussMySqlDatabaseUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListGaussMySqlDedicatedResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListGaussMySqlInstanceDetailInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGaussMySqlInstanceDetailInfoInvoker) Invoke() (*model.ListGaussMySqlInstanceDetailInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGaussMySqlInstanceDetailInfoResponse), nil
	}
}

type ListGaussMySqlInstanceDetailInfoUnifyStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGaussMySqlInstanceDetailInfoUnifyStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGaussMySqlInstanceDetailInfoUnifyStatusInvoker) Invoke() (*model.ListGaussMySqlInstanceDetailInfoUnifyStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGaussMySqlInstanceDetailInfoUnifyStatusResponse), nil
	}
}

type ListGaussMySqlInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGaussMySqlInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGaussMySqlInstancesInvoker) Invoke() (*model.ListGaussMySqlInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGaussMySqlInstancesResponse), nil
	}
}

type ListGaussMySqlInstancesUnifyStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGaussMySqlInstancesUnifyStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGaussMySqlInstancesUnifyStatusInvoker) Invoke() (*model.ListGaussMySqlInstancesUnifyStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGaussMySqlInstancesUnifyStatusResponse), nil
	}
}

type ListImmediateJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImmediateJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListInstanceConfigurationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListInstanceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListLtsErrorLogDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListLtsSlowlogDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListModifyHistoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListModifyHistoryInvoker) Invoke() (*model.ListModifyHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListModifyHistoryResponse), nil
	}
}

type ListParamsTemplateApplyHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListParamsTemplateApplyHistoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListParamsTemplateApplyHistoryInvoker) Invoke() (*model.ListParamsTemplateApplyHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListParamsTemplateApplyHistoryResponse), nil
	}
}

type ListProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListRecycleInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRecycleInstancesInvoker) Invoke() (*model.ListRecycleInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecycleInstancesResponse), nil
	}
}

type ListScheduleJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScheduleJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScheduleJobsInvoker) Invoke() (*model.ListScheduleJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScheduleJobsResponse), nil
	}
}

type ListTaurusDbNodeProcessesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTaurusDbNodeProcessesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTaurusDbNodeProcessesInvoker) Invoke() (*model.ListTaurusDbNodeProcessesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTaurusDbNodeProcessesResponse), nil
	}
}

type ModifyBackupEncryptStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyBackupEncryptStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ModifyGaussMySqlProxyRouteModeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ModifyGaussMysqlDnsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyGaussMysqlDnsInvoker) Invoke() (*model.ModifyGaussMysqlDnsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyGaussMysqlDnsResponse), nil
	}
}

type ModifyNodePriorityInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyNodePriorityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyNodePriorityInvoker) Invoke() (*model.ModifyNodePriorityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyNodePriorityResponse), nil
	}
}

type RenameInstanceNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *RenameInstanceNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RenameInstanceNodeInvoker) Invoke() (*model.RenameInstanceNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RenameInstanceNodeResponse), nil
	}
}

type ResetGaussMySqlDatabasePasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetGaussMySqlDatabasePasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ResetGaussMySqlPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *RestartGaussMySqlInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *RestartGaussMySqlNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestartGaussMySqlNodeInvoker) Invoke() (*model.RestartGaussMySqlNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartGaussMySqlNodeResponse), nil
	}
}

type RestartProxyInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartProxyInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestartProxyInstanceInvoker) Invoke() (*model.RestartProxyInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartProxyInstanceResponse), nil
	}
}

type RestoreOldInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreOldInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *SetGaussMySqlProxyWeightInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *SetGaussMySqlQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetGaussMySqlQuotasInvoker) Invoke() (*model.SetGaussMySqlQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetGaussMySqlQuotasResponse), nil
	}
}

type SetRecyclePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetRecyclePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetRecyclePolicyInvoker) Invoke() (*model.SetRecyclePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetRecyclePolicyResponse), nil
	}
}

type SetSqlFilterRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetSqlFilterRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetSqlFilterRuleInvoker) Invoke() (*model.SetSqlFilterRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetSqlFilterRuleResponse), nil
	}
}

type ShowAuditLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuditLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuditLogInvoker) Invoke() (*model.ShowAuditLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditLogResponse), nil
	}
}

type ShowAutoScalingHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoScalingHistoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAutoScalingHistoryInvoker) Invoke() (*model.ShowAutoScalingHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAutoScalingHistoryResponse), nil
	}
}

type ShowAutoScalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoScalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowBackupRestoreTimeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowDedicatedResourceInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowGaussMySqlBackupListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowGaussMySqlBackupPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowGaussMySqlConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowGaussMySqlEngineVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowGaussMySqlFlavorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowGaussMySqlIncrementalBackupListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowGaussMySqlInstanceInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGaussMySqlInstanceInfoInvoker) Invoke() (*model.ShowGaussMySqlInstanceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlInstanceInfoResponse), nil
	}
}

type ShowGaussMySqlInstanceInfoUnifyStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlInstanceInfoUnifyStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGaussMySqlInstanceInfoUnifyStatusInvoker) Invoke() (*model.ShowGaussMySqlInstanceInfoUnifyStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlInstanceInfoUnifyStatusResponse), nil
	}
}

type ShowGaussMySqlJobInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlJobInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowGaussMySqlProjectQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowGaussMySqlProxyFlavorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowGaussMySqlProxyListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowGaussMySqlQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowInstanceDatabaseVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceDatabaseVersionInvoker) Invoke() (*model.ShowInstanceDatabaseVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceDatabaseVersionResponse), nil
	}
}

type ShowInstanceEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceEipInvoker) Invoke() (*model.ShowInstanceEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceEipResponse), nil
	}
}

type ShowInstanceMonitorExtendInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceMonitorExtendInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowIntelligentDiagnosisAbnormalCountOfInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowIntelligentDiagnosisInstanceInfosPerMetricInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIntelligentDiagnosisInstanceInfosPerMetricInvoker) Invoke() (*model.ShowIntelligentDiagnosisInstanceInfosPerMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIntelligentDiagnosisInstanceInfosPerMetricResponse), nil
	}
}

type ShowLtsConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLtsConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLtsConfigsInvoker) Invoke() (*model.ShowLtsConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLtsConfigsResponse), nil
	}
}

type ShowMultiTenantInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMultiTenantInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMultiTenantInvoker) Invoke() (*model.ShowMultiTenantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMultiTenantResponse), nil
	}
}

type ShowProxyConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProxyConfigurationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProxyConfigurationsInvoker) Invoke() (*model.ShowProxyConfigurationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProxyConfigurationsResponse), nil
	}
}

type ShowProxyIpgroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProxyIpgroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProxyIpgroupInvoker) Invoke() (*model.ShowProxyIpgroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProxyIpgroupResponse), nil
	}
}

type ShowProxyVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProxyVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProxyVersionInvoker) Invoke() (*model.ShowProxyVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProxyVersionResponse), nil
	}
}

type ShowRecyclePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecyclePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRecyclePolicyInvoker) Invoke() (*model.ShowRecyclePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecyclePolicyResponse), nil
	}
}

type ShowRestoreTablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRestoreTablesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRestoreTablesInvoker) Invoke() (*model.ShowRestoreTablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRestoreTablesResponse), nil
	}
}

type ShowSlowLogStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSlowLogStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSlowLogStatisticsInvoker) Invoke() (*model.ShowSlowLogStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSlowLogStatisticsResponse), nil
	}
}

type ShowSlowlogSensitiveStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSlowlogSensitiveStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSlowlogSensitiveStatusInvoker) Invoke() (*model.ShowSlowlogSensitiveStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSlowlogSensitiveStatusResponse), nil
	}
}

type ShowSqlFilterControlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlFilterControlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowSqlFilterRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSqlFilterRuleInvoker) Invoke() (*model.ShowSqlFilterRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlFilterRuleResponse), nil
	}
}

type ShrinkGaussMySqlProxyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShrinkGaussMySqlProxyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *SwitchAccessControlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *SwitchGaussMySqlConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *SwitchGaussMySqlInstanceSslInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchGaussMySqlInstanceSslInvoker) Invoke() (*model.SwitchGaussMySqlInstanceSslResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchGaussMySqlInstanceSslResponse), nil
	}
}

type SwitchGaussMySqlProxySslInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchGaussMySqlProxySslInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchGaussMySqlProxySslInvoker) Invoke() (*model.SwitchGaussMySqlProxySslResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchGaussMySqlProxySslResponse), nil
	}
}

type UpdateAuditLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAuditLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateAutoScalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateBackupOffsitePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateGaussMySqlBackupPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateGaussMySqlConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateGaussMySqlDatabaseCommentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateGaussMySqlDatabaseUserCommentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateGaussMySqlInstanceAliasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateGaussMySqlInstanceEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateGaussMySqlInstanceInternalIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateGaussMySqlInstanceNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateGaussMySqlInstanceOpsWindowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateGaussMySqlInstancePortInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateGaussMySqlInstanceSecurityGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateGaussMySqlQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateInstanceConfigurationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateInstanceMonitorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceMonitorInvoker) Invoke() (*model.UpdateInstanceMonitorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceMonitorResponse), nil
	}
}

type UpdateMultiTenantInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMultiTenantInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMultiTenantInvoker) Invoke() (*model.UpdateMultiTenantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMultiTenantResponse), nil
	}
}

type UpdateNewNodeAutoAddSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNewNodeAutoAddSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNewNodeAutoAddSwitchInvoker) Invoke() (*model.UpdateNewNodeAutoAddSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNewNodeAutoAddSwitchResponse), nil
	}
}

type UpdateProxyConnectionPoolTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProxyConnectionPoolTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProxyConnectionPoolTypeInvoker) Invoke() (*model.UpdateProxyConnectionPoolTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProxyConnectionPoolTypeResponse), nil
	}
}

type UpdateProxyNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProxyNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProxyNameInvoker) Invoke() (*model.UpdateProxyNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProxyNameResponse), nil
	}
}

type UpdateProxyNewConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProxyNewConfigurationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateProxyPortInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateProxySessionConsistenceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProxySessionConsistenceInvoker) Invoke() (*model.UpdateProxySessionConsistenceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProxySessionConsistenceResponse), nil
	}
}

type UpdateServerlessPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateServerlessPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateServerlessPolicyInvoker) Invoke() (*model.UpdateServerlessPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateServerlessPolicyResponse), nil
	}
}

type UpdateSlowlogSensitiveSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSlowlogSensitiveSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSlowlogSensitiveSwitchInvoker) Invoke() (*model.UpdateSlowlogSensitiveSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSlowlogSensitiveSwitchResponse), nil
	}
}

type UpdateSqlFilterControlInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSqlFilterControlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSqlFilterControlInvoker) Invoke() (*model.UpdateSqlFilterControlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSqlFilterControlResponse), nil
	}
}

type UpdateTaurusNodeDataIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTaurusNodeDataIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTaurusNodeDataIpInvoker) Invoke() (*model.UpdateTaurusNodeDataIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTaurusNodeDataIpResponse), nil
	}
}

type UpdateTransactionSplitStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTransactionSplitStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpgradeGaussMySqlInstanceDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpgradeGaussMySqlInstanceDatabaseInvoker) Invoke() (*model.UpgradeGaussMySqlInstanceDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeGaussMySqlInstanceDatabaseResponse), nil
	}
}

type UpgradeProxyVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeProxyVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpgradeProxyVersionInvoker) Invoke() (*model.UpgradeProxyVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeProxyVersionResponse), nil
	}
}

type CheckClickHouseDataBaseConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckClickHouseDataBaseConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckClickHouseDataBaseConfigInvoker) Invoke() (*model.CheckClickHouseDataBaseConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckClickHouseDataBaseConfigResponse), nil
	}
}

type CheckClickHouseTableConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckClickHouseTableConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckClickHouseTableConfigInvoker) Invoke() (*model.CheckClickHouseTableConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckClickHouseTableConfigResponse), nil
	}
}

type CheckDataBaseConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckDataBaseConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckDataBaseConfigInvoker) Invoke() (*model.CheckDataBaseConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckDataBaseConfigResponse), nil
	}
}

type CheckStarRocksResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckStarRocksResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckStarRocksResourceInvoker) Invoke() (*model.CheckStarRocksResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckStarRocksResourceResponse), nil
	}
}

type CheckStarrocksParamsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckStarrocksParamsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckStarrocksParamsInvoker) Invoke() (*model.CheckStarrocksParamsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckStarrocksParamsResponse), nil
	}
}

type CheckTableConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckTableConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckTableConfigInvoker) Invoke() (*model.CheckTableConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckTableConfigResponse), nil
	}
}

type CreateClickHouseDataBaseReplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClickHouseDataBaseReplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateClickHouseDataBaseReplicationInvoker) Invoke() (*model.CreateClickHouseDataBaseReplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClickHouseDataBaseReplicationResponse), nil
	}
}

type CreateClickHouseDatabaseUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClickHouseDatabaseUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateClickHouseDatabaseUserInvoker) Invoke() (*model.CreateClickHouseDatabaseUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClickHouseDatabaseUserResponse), nil
	}
}

type CreateClickHouseInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClickHouseInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateClickHouseInstanceInvoker) Invoke() (*model.CreateClickHouseInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClickHouseInstanceResponse), nil
	}
}

type CreateStarRocksDataReplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStarRocksDataReplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateStarRocksDataReplicationInvoker) Invoke() (*model.CreateStarRocksDataReplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStarRocksDataReplicationResponse), nil
	}
}

type CreateStarRocksDatabaseUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStarRocksDatabaseUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateStarRocksDatabaseUserInvoker) Invoke() (*model.CreateStarRocksDatabaseUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStarRocksDatabaseUserResponse), nil
	}
}

type CreateStarrocksInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStarrocksInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateStarrocksInstanceInvoker) Invoke() (*model.CreateStarrocksInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStarrocksInstanceResponse), nil
	}
}

type DeleteClickHouseDataBaseConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClickHouseDataBaseConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteClickHouseDataBaseConfigInvoker) Invoke() (*model.DeleteClickHouseDataBaseConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClickHouseDataBaseConfigResponse), nil
	}
}

type DeleteClickHouseDataBaseReplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClickHouseDataBaseReplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteClickHouseDataBaseReplicationInvoker) Invoke() (*model.DeleteClickHouseDataBaseReplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClickHouseDataBaseReplicationResponse), nil
	}
}

type DeleteClickHouseDatabaseUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClickHouseDatabaseUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteClickHouseDatabaseUserInvoker) Invoke() (*model.DeleteClickHouseDatabaseUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClickHouseDatabaseUserResponse), nil
	}
}

type DeleteClickHouseInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClickHouseInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteClickHouseInstanceInvoker) Invoke() (*model.DeleteClickHouseInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClickHouseInstanceResponse), nil
	}
}

type DeleteClickHouseLtsConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClickHouseLtsConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteClickHouseLtsConfigInvoker) Invoke() (*model.DeleteClickHouseLtsConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClickHouseLtsConfigResponse), nil
	}
}

type DeleteStarRocksDataReplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStarRocksDataReplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteStarRocksDataReplicationInvoker) Invoke() (*model.DeleteStarRocksDataReplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStarRocksDataReplicationResponse), nil
	}
}

type DeleteStarRocksDatabaseUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStarRocksDatabaseUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteStarRocksDatabaseUserInvoker) Invoke() (*model.DeleteStarRocksDatabaseUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStarRocksDatabaseUserResponse), nil
	}
}

type DeleteStarrocksInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStarrocksInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteStarrocksInstanceInvoker) Invoke() (*model.DeleteStarrocksInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStarrocksInstanceResponse), nil
	}
}

type ListClickHouseDataBaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClickHouseDataBaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClickHouseDataBaseInvoker) Invoke() (*model.ListClickHouseDataBaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClickHouseDataBaseResponse), nil
	}
}

type ListClickHouseDataBaseParameterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClickHouseDataBaseParameterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClickHouseDataBaseParameterInvoker) Invoke() (*model.ListClickHouseDataBaseParameterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClickHouseDataBaseParameterResponse), nil
	}
}

type ListClickHouseDataBaseReplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClickHouseDataBaseReplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClickHouseDataBaseReplicationInvoker) Invoke() (*model.ListClickHouseDataBaseReplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClickHouseDataBaseReplicationResponse), nil
	}
}

type ListClickHouseDataBaseReplicationConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClickHouseDataBaseReplicationConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClickHouseDataBaseReplicationConfigInvoker) Invoke() (*model.ListClickHouseDataBaseReplicationConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClickHouseDataBaseReplicationConfigResponse), nil
	}
}

type ListClickHouseInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClickHouseInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClickHouseInstanceInvoker) Invoke() (*model.ListClickHouseInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClickHouseInstanceResponse), nil
	}
}

type ListClickHouseInstanceNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClickHouseInstanceNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClickHouseInstanceNodeInvoker) Invoke() (*model.ListClickHouseInstanceNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClickHouseInstanceNodeResponse), nil
	}
}

type ListHtapDataStoreInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHtapDataStoreInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHtapDataStoreInvoker) Invoke() (*model.ListHtapDataStoreResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHtapDataStoreResponse), nil
	}
}

type ListHtapFlavorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHtapFlavorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHtapFlavorInvoker) Invoke() (*model.ListHtapFlavorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHtapFlavorResponse), nil
	}
}

type ListHtapInstanceInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHtapInstanceInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHtapInstanceInfoInvoker) Invoke() (*model.ListHtapInstanceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHtapInstanceInfoResponse), nil
	}
}

type ListHtapStorageTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHtapStorageTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHtapStorageTypeInvoker) Invoke() (*model.ListHtapStorageTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHtapStorageTypeResponse), nil
	}
}

type ListStarRocksDataBasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStarRocksDataBasesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStarRocksDataBasesInvoker) Invoke() (*model.ListStarRocksDataBasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStarRocksDataBasesResponse), nil
	}
}

type ListStarRocksDataReplicationConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStarRocksDataReplicationConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStarRocksDataReplicationConfigInvoker) Invoke() (*model.ListStarRocksDataReplicationConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStarRocksDataReplicationConfigResponse), nil
	}
}

type ListStarRocksDataReplicationConfigByDataBaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStarRocksDataReplicationConfigByDataBaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStarRocksDataReplicationConfigByDataBaseInvoker) Invoke() (*model.ListStarRocksDataReplicationConfigByDataBaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStarRocksDataReplicationConfigByDataBaseResponse), nil
	}
}

type ListStarRocksDataReplicationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStarRocksDataReplicationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStarRocksDataReplicationsInvoker) Invoke() (*model.ListStarRocksDataReplicationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStarRocksDataReplicationsResponse), nil
	}
}

type ListStarRocksDbParametersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStarRocksDbParametersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStarRocksDbParametersInvoker) Invoke() (*model.ListStarRocksDbParametersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStarRocksDbParametersResponse), nil
	}
}

type ListStarrocksInstanceInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStarrocksInstanceInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStarrocksInstanceInfoInvoker) Invoke() (*model.ListStarrocksInstanceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStarrocksInstanceInfoResponse), nil
	}
}

type ModifyDataSyncInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyDataSyncInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyDataSyncInvoker) Invoke() (*model.ModifyDataSyncResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyDataSyncResponse), nil
	}
}

type PauseStarRocksDataReplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *PauseStarRocksDataReplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PauseStarRocksDataReplicationInvoker) Invoke() (*model.PauseStarRocksDataReplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PauseStarRocksDataReplicationResponse), nil
	}
}

type RebootClickHouseInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RebootClickHouseInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RebootClickHouseInstanceInvoker) Invoke() (*model.RebootClickHouseInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RebootClickHouseInstanceResponse), nil
	}
}

type ResizeClickHouseFlavorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeClickHouseFlavorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResizeClickHouseFlavorInvoker) Invoke() (*model.ResizeClickHouseFlavorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeClickHouseFlavorResponse), nil
	}
}

type ResizeClickHouseInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeClickHouseInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResizeClickHouseInstanceInvoker) Invoke() (*model.ResizeClickHouseInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeClickHouseInstanceResponse), nil
	}
}

type ResizeStarRocksFlavorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeStarRocksFlavorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResizeStarRocksFlavorInvoker) Invoke() (*model.ResizeStarRocksFlavorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeStarRocksFlavorResponse), nil
	}
}

type RestartStarrocksInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartStarrocksInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestartStarrocksInstanceInvoker) Invoke() (*model.RestartStarrocksInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartStarrocksInstanceResponse), nil
	}
}

type RestartStarrocksNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartStarrocksNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestartStarrocksNodeInvoker) Invoke() (*model.RestartStarrocksNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartStarrocksNodeResponse), nil
	}
}

type ResumeStarRocksDataReplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResumeStarRocksDataReplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResumeStarRocksDataReplicationInvoker) Invoke() (*model.ResumeStarRocksDataReplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResumeStarRocksDataReplicationResponse), nil
	}
}

type ShowClickHouseDatabaseUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClickHouseDatabaseUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClickHouseDatabaseUserInvoker) Invoke() (*model.ShowClickHouseDatabaseUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClickHouseDatabaseUserResponse), nil
	}
}

type ShowClickHouseLtsConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClickHouseLtsConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClickHouseLtsConfigInvoker) Invoke() (*model.ShowClickHouseLtsConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClickHouseLtsConfigResponse), nil
	}
}

type ShowClickHouseSlowLogDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClickHouseSlowLogDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClickHouseSlowLogDetailInvoker) Invoke() (*model.ShowClickHouseSlowLogDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClickHouseSlowLogDetailResponse), nil
	}
}

type ShowClickHouseSlowLogSensitiveStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClickHouseSlowLogSensitiveStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClickHouseSlowLogSensitiveStatusInvoker) Invoke() (*model.ShowClickHouseSlowLogSensitiveStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClickHouseSlowLogSensitiveStatusResponse), nil
	}
}

type ShowStarRocksDatabaseUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStarRocksDatabaseUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStarRocksDatabaseUserInvoker) Invoke() (*model.ShowStarRocksDatabaseUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStarRocksDatabaseUserResponse), nil
	}
}

type ShowStarrocksParamsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStarrocksParamsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStarrocksParamsInvoker) Invoke() (*model.ShowStarrocksParamsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStarrocksParamsResponse), nil
	}
}

type SyncStarRocksUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncStarRocksUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SyncStarRocksUsersInvoker) Invoke() (*model.SyncStarRocksUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncStarRocksUsersResponse), nil
	}
}

type UpdateClickHouseDataBaseConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClickHouseDataBaseConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateClickHouseDataBaseConfigInvoker) Invoke() (*model.UpdateClickHouseDataBaseConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClickHouseDataBaseConfigResponse), nil
	}
}

type UpdateClickHouseDatabaseUserPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClickHouseDatabaseUserPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateClickHouseDatabaseUserPasswordInvoker) Invoke() (*model.UpdateClickHouseDatabaseUserPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClickHouseDatabaseUserPasswordResponse), nil
	}
}

type UpdateClickHouseDatabaseUserPermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClickHouseDatabaseUserPermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateClickHouseDatabaseUserPermissionInvoker) Invoke() (*model.UpdateClickHouseDatabaseUserPermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClickHouseDatabaseUserPermissionResponse), nil
	}
}

type UpdateClickHouseLtsConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClickHouseLtsConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateClickHouseLtsConfigInvoker) Invoke() (*model.UpdateClickHouseLtsConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClickHouseLtsConfigResponse), nil
	}
}

type UpdateClickHouseSlowLogSensitiveStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClickHouseSlowLogSensitiveStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateClickHouseSlowLogSensitiveStatusInvoker) Invoke() (*model.UpdateClickHouseSlowLogSensitiveStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClickHouseSlowLogSensitiveStatusResponse), nil
	}
}

type UpdateStarRocksDatabaseUserPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStarRocksDatabaseUserPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateStarRocksDatabaseUserPasswordInvoker) Invoke() (*model.UpdateStarRocksDatabaseUserPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStarRocksDatabaseUserPasswordResponse), nil
	}
}

type UpdateStarRocksDatabaseUserPermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStarRocksDatabaseUserPermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateStarRocksDatabaseUserPermissionInvoker) Invoke() (*model.UpdateStarRocksDatabaseUserPermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStarRocksDatabaseUserPermissionResponse), nil
	}
}

type UpdateStarrocksParamsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStarrocksParamsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateStarrocksParamsInvoker) Invoke() (*model.UpdateStarrocksParamsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStarrocksParamsResponse), nil
	}
}

type UpgradeSrKernelVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeSrKernelVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpgradeSrKernelVersionInvoker) Invoke() (*model.UpgradeSrKernelVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeSrKernelVersionResponse), nil
	}
}
