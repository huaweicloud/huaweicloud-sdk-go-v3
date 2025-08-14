package v5

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/hss/v5/model"
)

type AddBaselineWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddBaselineWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddBaselineWhiteListInvoker) Invoke() (*model.AddBaselineWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddBaselineWhiteListResponse), nil
	}
}

type AddCceIntegrationProtectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddCceIntegrationProtectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddCceIntegrationProtectionInvoker) Invoke() (*model.AddCceIntegrationProtectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddCceIntegrationProtectionResponse), nil
	}
}

type AddHostsGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddHostsGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddHostsGroupInvoker) Invoke() (*model.AddHostsGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddHostsGroupResponse), nil
	}
}

type AddLoginWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddLoginWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddLoginWhiteListInvoker) Invoke() (*model.AddLoginWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddLoginWhiteListResponse), nil
	}
}

type AddPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddPolicyInvoker) Invoke() (*model.AddPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddPolicyResponse), nil
	}
}

type AddProtectionPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddProtectionPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddProtectionPolicyInvoker) Invoke() (*model.AddProtectionPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddProtectionPolicyResponse), nil
	}
}

type AddSystemUserWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddSystemUserWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddSystemUserWhiteListInvoker) Invoke() (*model.AddSystemUserWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddSystemUserWhiteListResponse), nil
	}
}

type AssociatePolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociatePolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociatePolicyGroupInvoker) Invoke() (*model.AssociatePolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociatePolicyGroupResponse), nil
	}
}

type BatchAddAccountsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddAccountsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAddAccountsInvoker) Invoke() (*model.BatchAddAccountsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddAccountsResponse), nil
	}
}

type BatchCreateTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateTagsInvoker) Invoke() (*model.BatchCreateTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateTagsResponse), nil
	}
}

type BatchScanSwrImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchScanSwrImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchScanSwrImageInvoker) Invoke() (*model.BatchScanSwrImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchScanSwrImageResponse), nil
	}
}

type BatchStartProtectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchStartProtectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchStartProtectionInvoker) Invoke() (*model.BatchStartProtectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchStartProtectionResponse), nil
	}
}

type ChangeBaselineWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeBaselineWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeBaselineWhiteListInvoker) Invoke() (*model.ChangeBaselineWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeBaselineWhiteListResponse), nil
	}
}

type ChangeBlockedIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeBlockedIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeBlockedIpInvoker) Invoke() (*model.ChangeBlockedIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeBlockedIpResponse), nil
	}
}

type ChangeCheckRuleActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeCheckRuleActionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeCheckRuleActionInvoker) Invoke() (*model.ChangeCheckRuleActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeCheckRuleActionResponse), nil
	}
}

type ChangeClusterEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeClusterEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeClusterEventsInvoker) Invoke() (*model.ChangeClusterEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeClusterEventsResponse), nil
	}
}

type ChangeClusterProtectionPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeClusterProtectionPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeClusterProtectionPolicyInvoker) Invoke() (*model.ChangeClusterProtectionPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeClusterProtectionPolicyResponse), nil
	}
}

type ChangeEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeEventInvoker) Invoke() (*model.ChangeEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeEventResponse), nil
	}
}

type ChangeHostsGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeHostsGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeHostsGroupInvoker) Invoke() (*model.ChangeHostsGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeHostsGroupResponse), nil
	}
}

type ChangeIsolatedFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeIsolatedFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeIsolatedFileInvoker) Invoke() (*model.ChangeIsolatedFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeIsolatedFileResponse), nil
	}
}

type ChangePasswordComplexityStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangePasswordComplexityStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangePasswordComplexityStatusInvoker) Invoke() (*model.ChangePasswordComplexityStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangePasswordComplexityStatusResponse), nil
	}
}

type ChangeVulStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeVulStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeVulStatusInvoker) Invoke() (*model.ChangeVulStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeVulStatusResponse), nil
	}
}

type CreateClusterProtectionPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterProtectionPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateClusterProtectionPolicyInvoker) Invoke() (*model.CreateClusterProtectionPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterProtectionPolicyResponse), nil
	}
}

type CreateClustersInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClustersInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateClustersInfoInvoker) Invoke() (*model.CreateClustersInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClustersInfoResponse), nil
	}
}

type CreateContainerNetworkPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateContainerNetworkPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateContainerNetworkPolicyInvoker) Invoke() (*model.CreateContainerNetworkPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateContainerNetworkPolicyResponse), nil
	}
}

type CreateDecoyPortPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDecoyPortPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDecoyPortPolicyInvoker) Invoke() (*model.CreateDecoyPortPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDecoyPortPolicyResponse), nil
	}
}

type CreateGlobalAssetScanTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGlobalAssetScanTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGlobalAssetScanTaskInvoker) Invoke() (*model.CreateGlobalAssetScanTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGlobalAssetScanTaskResponse), nil
	}
}

type CreateQuotasOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateQuotasOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateQuotasOrderInvoker) Invoke() (*model.CreateQuotasOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateQuotasOrderResponse), nil
	}
}

type CreateSecurityGroupPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecurityGroupPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSecurityGroupPolicyInvoker) Invoke() (*model.CreateSecurityGroupPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecurityGroupPolicyResponse), nil
	}
}

type DeleteAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAccountInvoker) Invoke() (*model.DeleteAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAccountResponse), nil
	}
}

type DeleteBaselineWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBaselineWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBaselineWhiteListInvoker) Invoke() (*model.DeleteBaselineWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBaselineWhiteListResponse), nil
	}
}

type DeleteClusterProtectionPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClusterProtectionPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteClusterProtectionPolicyInvoker) Invoke() (*model.DeleteClusterProtectionPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClusterProtectionPolicyResponse), nil
	}
}

type DeleteContainerNetworkPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteContainerNetworkPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteContainerNetworkPolicyInvoker) Invoke() (*model.DeleteContainerNetworkPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteContainerNetworkPolicyResponse), nil
	}
}

type DeleteDecoyPortHostPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDecoyPortHostPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDecoyPortHostPolicyInvoker) Invoke() (*model.DeleteDecoyPortHostPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDecoyPortHostPolicyResponse), nil
	}
}

type DeleteDecoyPortPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDecoyPortPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDecoyPortPolicyInvoker) Invoke() (*model.DeleteDecoyPortPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDecoyPortPolicyResponse), nil
	}
}

type DeleteHostsGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHostsGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHostsGroupInvoker) Invoke() (*model.DeleteHostsGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHostsGroupResponse), nil
	}
}

type DeleteIsolatedFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteIsolatedFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteIsolatedFileInvoker) Invoke() (*model.DeleteIsolatedFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteIsolatedFileResponse), nil
	}
}

type DeletePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePolicyInvoker) Invoke() (*model.DeletePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePolicyResponse), nil
	}
}

type DeleteProtectionPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProtectionPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteProtectionPolicyInvoker) Invoke() (*model.DeleteProtectionPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProtectionPolicyResponse), nil
	}
}

type DeleteResourceInstanceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResourceInstanceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteResourceInstanceTagInvoker) Invoke() (*model.DeleteResourceInstanceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResourceInstanceTagResponse), nil
	}
}

type DeleteSecurityGroupPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecurityGroupPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSecurityGroupPolicyInvoker) Invoke() (*model.DeleteSecurityGroupPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecurityGroupPolicyResponse), nil
	}
}

type ExportContainerListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportContainerListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportContainerListInvoker) Invoke() (*model.ExportContainerListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportContainerListResponse), nil
	}
}

type ListAccountsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccountsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccountsInvoker) Invoke() (*model.ListAccountsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccountsResponse), nil
	}
}

type ListAgentInstallScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAgentInstallScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAgentInstallScriptInvoker) Invoke() (*model.ListAgentInstallScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAgentInstallScriptResponse), nil
	}
}

type ListAlarmWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlarmWhiteListInvoker) Invoke() (*model.ListAlarmWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmWhiteListResponse), nil
	}
}

type ListAntivirusHandleHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAntivirusHandleHistoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAntivirusHandleHistoryInvoker) Invoke() (*model.ListAntivirusHandleHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAntivirusHandleHistoryResponse), nil
	}
}

type ListAppChangeHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppChangeHistoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppChangeHistoriesInvoker) Invoke() (*model.ListAppChangeHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppChangeHistoriesResponse), nil
	}
}

type ListAppStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppStatisticsInvoker) Invoke() (*model.ListAppStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppStatisticsResponse), nil
	}
}

type ListAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppsInvoker) Invoke() (*model.ListAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppsResponse), nil
	}
}

type ListAutoLaunchChangeHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAutoLaunchChangeHistoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAutoLaunchChangeHistoriesInvoker) Invoke() (*model.ListAutoLaunchChangeHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAutoLaunchChangeHistoriesResponse), nil
	}
}

type ListAutoLaunchStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAutoLaunchStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAutoLaunchStatisticsInvoker) Invoke() (*model.ListAutoLaunchStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAutoLaunchStatisticsResponse), nil
	}
}

type ListAutoLaunchsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAutoLaunchsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAutoLaunchsInvoker) Invoke() (*model.ListAutoLaunchsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAutoLaunchsResponse), nil
	}
}

type ListBackupVaultsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackupVaultsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBackupVaultsInvoker) Invoke() (*model.ListBackupVaultsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackupVaultsResponse), nil
	}
}

type ListBaselineWhiteListsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBaselineWhiteListsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBaselineWhiteListsInvoker) Invoke() (*model.ListBaselineWhiteListsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBaselineWhiteListsResponse), nil
	}
}

type ListBlockedIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBlockedIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBlockedIpInvoker) Invoke() (*model.ListBlockedIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBlockedIpResponse), nil
	}
}

type ListCceClusterConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCceClusterConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCceClusterConfigInvoker) Invoke() (*model.ListCceClusterConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCceClusterConfigResponse), nil
	}
}

type ListCceClusterDetectRiskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCceClusterDetectRiskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCceClusterDetectRiskInvoker) Invoke() (*model.ListCceClusterDetectRiskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCceClusterDetectRiskResponse), nil
	}
}

type ListCheckFeatureRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCheckFeatureRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCheckFeatureRuleInvoker) Invoke() (*model.ListCheckFeatureRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCheckFeatureRuleResponse), nil
	}
}

type ListClusterAuditLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterAuditLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClusterAuditLogsInvoker) Invoke() (*model.ListClusterAuditLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterAuditLogsResponse), nil
	}
}

type ListClusterEventLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterEventLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClusterEventLogsInvoker) Invoke() (*model.ListClusterEventLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterEventLogsResponse), nil
	}
}

type ListClusterEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClusterEventsInvoker) Invoke() (*model.ListClusterEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterEventsResponse), nil
	}
}

type ListClusterProtectOverviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterProtectOverviewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClusterProtectOverviewInvoker) Invoke() (*model.ListClusterProtectOverviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterProtectOverviewResponse), nil
	}
}

type ListClusterProtectPolicyTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterProtectPolicyTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClusterProtectPolicyTemplatesInvoker) Invoke() (*model.ListClusterProtectPolicyTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterProtectPolicyTemplatesResponse), nil
	}
}

type ListClusterProtectionDefaultPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterProtectionDefaultPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClusterProtectionDefaultPolicyInvoker) Invoke() (*model.ListClusterProtectionDefaultPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterProtectionDefaultPolicyResponse), nil
	}
}

type ListClusterProtectionInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterProtectionInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClusterProtectionInfoInvoker) Invoke() (*model.ListClusterProtectionInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterProtectionInfoResponse), nil
	}
}

type ListClusterProtectionItemInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterProtectionItemInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClusterProtectionItemInvoker) Invoke() (*model.ListClusterProtectionItemResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterProtectionItemResponse), nil
	}
}

type ListClusterProtectionPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterProtectionPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClusterProtectionPolicyInvoker) Invoke() (*model.ListClusterProtectionPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterProtectionPolicyResponse), nil
	}
}

type ListClusterProtectionPolicyDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterProtectionPolicyDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClusterProtectionPolicyDetailInvoker) Invoke() (*model.ListClusterProtectionPolicyDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterProtectionPolicyDetailResponse), nil
	}
}

type ListCommonTipsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCommonTipsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCommonTipsInvoker) Invoke() (*model.ListCommonTipsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCommonTipsResponse), nil
	}
}

type ListContainerCmdLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListContainerCmdLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListContainerCmdLogsInvoker) Invoke() (*model.ListContainerCmdLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListContainerCmdLogsResponse), nil
	}
}

type ListContainerImageLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListContainerImageLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListContainerImageLogsInvoker) Invoke() (*model.ListContainerImageLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListContainerImageLogsResponse), nil
	}
}

type ListContainerImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListContainerImagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListContainerImagesInvoker) Invoke() (*model.ListContainerImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListContainerImagesResponse), nil
	}
}

type ListContainerLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListContainerLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListContainerLogsInvoker) Invoke() (*model.ListContainerLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListContainerLogsResponse), nil
	}
}

type ListContainerNetworkClustersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListContainerNetworkClustersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListContainerNetworkClustersInvoker) Invoke() (*model.ListContainerNetworkClustersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListContainerNetworkClustersResponse), nil
	}
}

type ListContainerNetworkNodeListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListContainerNetworkNodeListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListContainerNetworkNodeListInvoker) Invoke() (*model.ListContainerNetworkNodeListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListContainerNetworkNodeListResponse), nil
	}
}

type ListContainerNetworkPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListContainerNetworkPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListContainerNetworkPolicyInvoker) Invoke() (*model.ListContainerNetworkPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListContainerNetworkPolicyResponse), nil
	}
}

type ListContainerNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListContainerNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListContainerNodesInvoker) Invoke() (*model.ListContainerNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListContainerNodesResponse), nil
	}
}

type ListContainersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListContainersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListContainersInvoker) Invoke() (*model.ListContainersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListContainersResponse), nil
	}
}

type ListDecoyPortPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDecoyPortPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDecoyPortPolicyInvoker) Invoke() (*model.ListDecoyPortPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDecoyPortPolicyResponse), nil
	}
}

type ListDownloadExportedFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDownloadExportedFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDownloadExportedFileInvoker) Invoke() (*model.ListDownloadExportedFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDownloadExportedFileResponse), nil
	}
}

type ListEventHandleHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventHandleHistoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEventHandleHistoryInvoker) Invoke() (*model.ListEventHandleHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventHandleHistoryResponse), nil
	}
}

type ListGlobalAssetScanTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalAssetScanTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGlobalAssetScanTaskInvoker) Invoke() (*model.ListGlobalAssetScanTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalAssetScanTaskResponse), nil
	}
}

type ListHostGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHostGroupsInvoker) Invoke() (*model.ListHostGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostGroupsResponse), nil
	}
}

type ListHostStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHostStatusInvoker) Invoke() (*model.ListHostStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostStatusResponse), nil
	}
}

type ListImageLocalInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageLocalInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageLocalInvoker) Invoke() (*model.ListImageLocalResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageLocalResponse), nil
	}
}

type ListImageRiskConfigRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageRiskConfigRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageRiskConfigRulesInvoker) Invoke() (*model.ListImageRiskConfigRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageRiskConfigRulesResponse), nil
	}
}

type ListImageRiskConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageRiskConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageRiskConfigsInvoker) Invoke() (*model.ListImageRiskConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageRiskConfigsResponse), nil
	}
}

type ListImageVulnerabilitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageVulnerabilitiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageVulnerabilitiesInvoker) Invoke() (*model.ListImageVulnerabilitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageVulnerabilitiesResponse), nil
	}
}

type ListIsolatedFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIsolatedFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIsolatedFileInvoker) Invoke() (*model.ListIsolatedFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIsolatedFileResponse), nil
	}
}

type ListJarPackageHostInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJarPackageHostInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListJarPackageHostInfoInvoker) Invoke() (*model.ListJarPackageHostInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJarPackageHostInfoResponse), nil
	}
}

type ListJarPackageStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJarPackageStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListJarPackageStatisticsInvoker) Invoke() (*model.ListJarPackageStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJarPackageStatisticsResponse), nil
	}
}

type ListK8sCronJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListK8sCronJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListK8sCronJobsInvoker) Invoke() (*model.ListK8sCronJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListK8sCronJobsResponse), nil
	}
}

type ListK8sDaemonSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListK8sDaemonSetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListK8sDaemonSetsInvoker) Invoke() (*model.ListK8sDaemonSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListK8sDaemonSetsResponse), nil
	}
}

type ListK8sDeploymentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListK8sDeploymentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListK8sDeploymentsInvoker) Invoke() (*model.ListK8sDeploymentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListK8sDeploymentsResponse), nil
	}
}

type ListK8sJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListK8sJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListK8sJobsInvoker) Invoke() (*model.ListK8sJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListK8sJobsResponse), nil
	}
}

type ListK8sPodsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListK8sPodsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListK8sPodsInvoker) Invoke() (*model.ListK8sPodsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListK8sPodsResponse), nil
	}
}

type ListK8sStatefulSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListK8sStatefulSetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListK8sStatefulSetsInvoker) Invoke() (*model.ListK8sStatefulSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListK8sStatefulSetsResponse), nil
	}
}

type ListKernelModuleHostInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListKernelModuleHostInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListKernelModuleHostInfoInvoker) Invoke() (*model.ListKernelModuleHostInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListKernelModuleHostInfoResponse), nil
	}
}

type ListKernelModuleStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListKernelModuleStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListKernelModuleStatisticsInvoker) Invoke() (*model.ListKernelModuleStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListKernelModuleStatisticsResponse), nil
	}
}

type ListKubernetesClusterDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListKubernetesClusterDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListKubernetesClusterDetailsInvoker) Invoke() (*model.ListKubernetesClusterDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListKubernetesClusterDetailsResponse), nil
	}
}

type ListKubernetesEndpointDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListKubernetesEndpointDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListKubernetesEndpointDetailsInvoker) Invoke() (*model.ListKubernetesEndpointDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListKubernetesEndpointDetailsResponse), nil
	}
}

type ListKubernetesServiceDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListKubernetesServiceDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListKubernetesServiceDetailsInvoker) Invoke() (*model.ListKubernetesServiceDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListKubernetesServiceDetailsResponse), nil
	}
}

type ListLoginCommonIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLoginCommonIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLoginCommonIpInvoker) Invoke() (*model.ListLoginCommonIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLoginCommonIpResponse), nil
	}
}

type ListLoginCommonLocationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLoginCommonLocationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLoginCommonLocationInvoker) Invoke() (*model.ListLoginCommonLocationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLoginCommonLocationResponse), nil
	}
}

type ListLoginWhiteIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLoginWhiteIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLoginWhiteIpInvoker) Invoke() (*model.ListLoginWhiteIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLoginWhiteIpResponse), nil
	}
}

type ListLoginWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLoginWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLoginWhiteListInvoker) Invoke() (*model.ListLoginWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLoginWhiteListResponse), nil
	}
}

type ListNamespacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNamespacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNamespacesInvoker) Invoke() (*model.ListNamespacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNamespacesResponse), nil
	}
}

type ListOperationLogsByVaultNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOperationLogsByVaultNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOperationLogsByVaultNameInvoker) Invoke() (*model.ListOperationLogsByVaultNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOperationLogsByVaultNameResponse), nil
	}
}

type ListOrganizationTreeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOrganizationTreeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOrganizationTreeInvoker) Invoke() (*model.ListOrganizationTreeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOrganizationTreeResponse), nil
	}
}

type ListPasswordComplexityInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPasswordComplexityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPasswordComplexityInvoker) Invoke() (*model.ListPasswordComplexityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPasswordComplexityResponse), nil
	}
}

type ListPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPoliciesInvoker) Invoke() (*model.ListPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPoliciesResponse), nil
	}
}

type ListPolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyGroupInvoker) Invoke() (*model.ListPolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyGroupResponse), nil
	}
}

type ListPortHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPortHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPortHostInvoker) Invoke() (*model.ListPortHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPortHostResponse), nil
	}
}

type ListPortStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPortStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPortStatisticsInvoker) Invoke() (*model.ListPortStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPortStatisticsResponse), nil
	}
}

type ListPortsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPortsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPortsInvoker) Invoke() (*model.ListPortsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPortsResponse), nil
	}
}

type ListProcessStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProcessStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProcessStatisticsInvoker) Invoke() (*model.ListProcessStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProcessStatisticsResponse), nil
	}
}

type ListProcessesHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProcessesHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProcessesHostInvoker) Invoke() (*model.ListProcessesHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProcessesHostResponse), nil
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

type ListProtectionPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProtectionPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProtectionPolicyInvoker) Invoke() (*model.ListProtectionPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProtectionPolicyResponse), nil
	}
}

type ListProtectionServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProtectionServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProtectionServerInvoker) Invoke() (*model.ListProtectionServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProtectionServerResponse), nil
	}
}

type ListProtectionServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProtectionServersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProtectionServersInvoker) Invoke() (*model.ListProtectionServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProtectionServersResponse), nil
	}
}

type ListQueryExportTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQueryExportTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQueryExportTaskInvoker) Invoke() (*model.ListQueryExportTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQueryExportTaskResponse), nil
	}
}

type ListQuotasDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotasDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuotasDetailInvoker) Invoke() (*model.ListQuotasDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotasDetailResponse), nil
	}
}

type ListRansomwareProtectionNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRansomwareProtectionNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRansomwareProtectionNodesInvoker) Invoke() (*model.ListRansomwareProtectionNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRansomwareProtectionNodesResponse), nil
	}
}

type ListRaspEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRaspEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRaspEventsInvoker) Invoke() (*model.ListRaspEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRaspEventsResponse), nil
	}
}

type ListRaspPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRaspPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRaspPoliciesInvoker) Invoke() (*model.ListRaspPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRaspPoliciesResponse), nil
	}
}

type ListRiskConfigCheckRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRiskConfigCheckRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRiskConfigCheckRulesInvoker) Invoke() (*model.ListRiskConfigCheckRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRiskConfigCheckRulesResponse), nil
	}
}

type ListRiskConfigHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRiskConfigHostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRiskConfigHostsInvoker) Invoke() (*model.ListRiskConfigHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRiskConfigHostsResponse), nil
	}
}

type ListRiskConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRiskConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRiskConfigsInvoker) Invoke() (*model.ListRiskConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRiskConfigsResponse), nil
	}
}

type ListSecurityEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityEventsInvoker) Invoke() (*model.ListSecurityEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityEventsResponse), nil
	}
}

type ListSecurityGroupPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityGroupPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityGroupPoliciesInvoker) Invoke() (*model.ListSecurityGroupPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityGroupPoliciesResponse), nil
	}
}

type ListSecurityGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityGroupsInvoker) Invoke() (*model.ListSecurityGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityGroupsResponse), nil
	}
}

type ListSwrImageRepositoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSwrImageRepositoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSwrImageRepositoryInvoker) Invoke() (*model.ListSwrImageRepositoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSwrImageRepositoryResponse), nil
	}
}

type ListSystemUserWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSystemUserWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSystemUserWhiteListInvoker) Invoke() (*model.ListSystemUserWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSystemUserWhiteListResponse), nil
	}
}

type ListTwoFactorLoginHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTwoFactorLoginHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTwoFactorLoginHostInvoker) Invoke() (*model.ListTwoFactorLoginHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTwoFactorLoginHostResponse), nil
	}
}

type ListUserChangeHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserChangeHistoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserChangeHistoriesInvoker) Invoke() (*model.ListUserChangeHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserChangeHistoriesResponse), nil
	}
}

type ListUserStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserStatisticsInvoker) Invoke() (*model.ListUserStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserStatisticsResponse), nil
	}
}

type ListUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUsersInvoker) Invoke() (*model.ListUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsersResponse), nil
	}
}

type ListVulHandleHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulHandleHistoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulHandleHistoryInvoker) Invoke() (*model.ListVulHandleHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulHandleHistoryResponse), nil
	}
}

type ListVulnerabilityCveInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulnerabilityCveInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulnerabilityCveInvoker) Invoke() (*model.ListVulnerabilityCveResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulnerabilityCveResponse), nil
	}
}

type ListWeakPasswordUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWeakPasswordUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWeakPasswordUsersInvoker) Invoke() (*model.ListWeakPasswordUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWeakPasswordUsersResponse), nil
	}
}

type ListWebAppAndServiceStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWebAppAndServiceStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWebAppAndServiceStatisticsInvoker) Invoke() (*model.ListWebAppAndServiceStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWebAppAndServiceStatisticsResponse), nil
	}
}

type ListWebAppAndServicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWebAppAndServicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWebAppAndServicesInvoker) Invoke() (*model.ListWebAppAndServicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWebAppAndServicesResponse), nil
	}
}

type ListWebFrameworkHostInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWebFrameworkHostInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWebFrameworkHostInfoInvoker) Invoke() (*model.ListWebFrameworkHostInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWebFrameworkHostInfoResponse), nil
	}
}

type ListWebFrameworkStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWebFrameworkStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWebFrameworkStatisticsInvoker) Invoke() (*model.ListWebFrameworkStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWebFrameworkStatisticsResponse), nil
	}
}

type ListWebSiteHostInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWebSiteHostInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWebSiteHostInfoInvoker) Invoke() (*model.ListWebSiteHostInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWebSiteHostInfoResponse), nil
	}
}

type ListWebSiteStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWebSiteStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWebSiteStatisticsInvoker) Invoke() (*model.ListWebSiteStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWebSiteStatisticsResponse), nil
	}
}

type ListWorkLoadsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkLoadsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkLoadsInvoker) Invoke() (*model.ListWorkLoadsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkLoadsResponse), nil
	}
}

type ModifyDecoyPortPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyDecoyPortPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyDecoyPortPolicyInvoker) Invoke() (*model.ModifyDecoyPortPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyDecoyPortPolicyResponse), nil
	}
}

type RemoveAlarmWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveAlarmWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RemoveAlarmWhiteListInvoker) Invoke() (*model.RemoveAlarmWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveAlarmWhiteListResponse), nil
	}
}

type RemoveLoginWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveLoginWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RemoveLoginWhiteListInvoker) Invoke() (*model.RemoveLoginWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveLoginWhiteListResponse), nil
	}
}

type RemoveSystemUserWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveSystemUserWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RemoveSystemUserWhiteListInvoker) Invoke() (*model.RemoveSystemUserWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveSystemUserWhiteListResponse), nil
	}
}

type RunHostAssetManualCollectInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunHostAssetManualCollectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunHostAssetManualCollectInvoker) Invoke() (*model.RunHostAssetManualCollectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunHostAssetManualCollectResponse), nil
	}
}

type RunImageSynchronizeInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunImageSynchronizeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunImageSynchronizeInvoker) Invoke() (*model.RunImageSynchronizeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunImageSynchronizeResponse), nil
	}
}

type SetTwoFactorLoginConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetTwoFactorLoginConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetTwoFactorLoginConfigInvoker) Invoke() (*model.SetTwoFactorLoginConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetTwoFactorLoginConfigResponse), nil
	}
}

type ShowAgentStatisticsStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAgentStatisticsStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAgentStatisticsStatusInvoker) Invoke() (*model.ShowAgentStatisticsStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgentStatisticsStatusResponse), nil
	}
}

type ShowAppRaspSwitchStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppRaspSwitchStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppRaspSwitchStatusInvoker) Invoke() (*model.ShowAppRaspSwitchStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppRaspSwitchStatusResponse), nil
	}
}

type ShowAssetStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAssetStatisticInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAssetStatisticInvoker) Invoke() (*model.ShowAssetStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssetStatisticResponse), nil
	}
}

type ShowBackupPolicyInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBackupPolicyInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBackupPolicyInfoInvoker) Invoke() (*model.ShowBackupPolicyInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBackupPolicyInfoResponse), nil
	}
}

type ShowBaselineScanStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBaselineScanStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBaselineScanStatusInvoker) Invoke() (*model.ShowBaselineScanStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBaselineScanStatusResponse), nil
	}
}

type ShowBaselineWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBaselineWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBaselineWhiteListInvoker) Invoke() (*model.ShowBaselineWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBaselineWhiteListResponse), nil
	}
}

type ShowCheckRuleDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCheckRuleDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCheckRuleDetailInvoker) Invoke() (*model.ShowCheckRuleDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCheckRuleDetailResponse), nil
	}
}

type ShowClusterAssetStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterAssetStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClusterAssetStatisticsInvoker) Invoke() (*model.ShowClusterAssetStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterAssetStatisticsResponse), nil
	}
}

type ShowClusterProtectPolicyTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterProtectPolicyTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClusterProtectPolicyTemplateInvoker) Invoke() (*model.ShowClusterProtectPolicyTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterProtectPolicyTemplateResponse), nil
	}
}

type ShowCommonPortInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCommonPortInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCommonPortInvoker) Invoke() (*model.ShowCommonPortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCommonPortResponse), nil
	}
}

type ShowContainerNetworkInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowContainerNetworkInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowContainerNetworkInfoInvoker) Invoke() (*model.ShowContainerNetworkInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowContainerNetworkInfoResponse), nil
	}
}

type ShowContainerNodeStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowContainerNodeStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowContainerNodeStatisticsInvoker) Invoke() (*model.ShowContainerNodeStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowContainerNodeStatisticsResponse), nil
	}
}

type ShowContainerProtectionStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowContainerProtectionStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowContainerProtectionStatusInvoker) Invoke() (*model.ShowContainerProtectionStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowContainerProtectionStatusResponse), nil
	}
}

type ShowDecoyPortPolicyDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDecoyPortPolicyDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDecoyPortPolicyDetailsInvoker) Invoke() (*model.ShowDecoyPortPolicyDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDecoyPortPolicyDetailsResponse), nil
	}
}

type ShowHostAssetManualCollectStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHostAssetManualCollectStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHostAssetManualCollectStatusInvoker) Invoke() (*model.ShowHostAssetManualCollectStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHostAssetManualCollectStatusResponse), nil
	}
}

type ShowHostProtectionStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHostProtectionStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHostProtectionStatusInvoker) Invoke() (*model.ShowHostProtectionStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHostProtectionStatusResponse), nil
	}
}

type ShowImageAssetStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageAssetStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowImageAssetStatisticsInvoker) Invoke() (*model.ShowImageAssetStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageAssetStatisticsResponse), nil
	}
}

type ShowImageCheckRuleDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageCheckRuleDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowImageCheckRuleDetailInvoker) Invoke() (*model.ShowImageCheckRuleDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageCheckRuleDetailResponse), nil
	}
}

type ShowK8sContainerDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowK8sContainerDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowK8sContainerDetailInvoker) Invoke() (*model.ShowK8sContainerDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowK8sContainerDetailResponse), nil
	}
}

type ShowK8sPodDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowK8sPodDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowK8sPodDetailInvoker) Invoke() (*model.ShowK8sPodDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowK8sPodDetailResponse), nil
	}
}

type ShowKubernetesEndpointInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKubernetesEndpointInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowKubernetesEndpointInfoInvoker) Invoke() (*model.ShowKubernetesEndpointInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKubernetesEndpointInfoResponse), nil
	}
}

type ShowKubernetesServiceInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKubernetesServiceInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowKubernetesServiceInfoInvoker) Invoke() (*model.ShowKubernetesServiceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKubernetesServiceInfoResponse), nil
	}
}

type ShowLatestExportTaskByTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLatestExportTaskByTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLatestExportTaskByTypeInvoker) Invoke() (*model.ShowLatestExportTaskByTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLatestExportTaskByTypeResponse), nil
	}
}

type ShowNetworkStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNetworkStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNetworkStatisticsInvoker) Invoke() (*model.ShowNetworkStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNetworkStatisticsResponse), nil
	}
}

type ShowOsStatisticsInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOsStatisticsInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOsStatisticsInfoInvoker) Invoke() (*model.ShowOsStatisticsInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOsStatisticsInfoResponse), nil
	}
}

type ShowPageNoticesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPageNoticesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPageNoticesInvoker) Invoke() (*model.ShowPageNoticesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPageNoticesResponse), nil
	}
}

type ShowProductdataOfferingInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProductdataOfferingInfosInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProductdataOfferingInfosInvoker) Invoke() (*model.ShowProductdataOfferingInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProductdataOfferingInfosResponse), nil
	}
}

type ShowQuotaStatisticsInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotaStatisticsInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowQuotaStatisticsInfoInvoker) Invoke() (*model.ShowQuotaStatisticsInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotaStatisticsInfoResponse), nil
	}
}

type ShowRaspPolicyDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRaspPolicyDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRaspPolicyDetailInvoker) Invoke() (*model.ShowRaspPolicyDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRaspPolicyDetailResponse), nil
	}
}

type ShowRaspProtectStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRaspProtectStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRaspProtectStatisticsInvoker) Invoke() (*model.ShowRaspProtectStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRaspProtectStatisticsResponse), nil
	}
}

type ShowRaspServerDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRaspServerDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRaspServerDetailInvoker) Invoke() (*model.ShowRaspServerDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRaspServerDetailResponse), nil
	}
}

type ShowResourceQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceQuotasInvoker) Invoke() (*model.ShowResourceQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceQuotasResponse), nil
	}
}

type ShowRiskConfigDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRiskConfigDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRiskConfigDetailInvoker) Invoke() (*model.ShowRiskConfigDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRiskConfigDetailResponse), nil
	}
}

type ShowSingleBackupPolicyInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSingleBackupPolicyInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSingleBackupPolicyInfoInvoker) Invoke() (*model.ShowSingleBackupPolicyInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSingleBackupPolicyInfoResponse), nil
	}
}

type StartProtectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartProtectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartProtectionInvoker) Invoke() (*model.StartProtectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartProtectionResponse), nil
	}
}

type StopProtectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopProtectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopProtectionInvoker) Invoke() (*model.StopProtectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopProtectionResponse), nil
	}
}

type SwitchClusterProtectionModeInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchClusterProtectionModeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchClusterProtectionModeInvoker) Invoke() (*model.SwitchClusterProtectionModeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchClusterProtectionModeResponse), nil
	}
}

type SwitchContainerProtectStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchContainerProtectStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchContainerProtectStatusInvoker) Invoke() (*model.SwitchContainerProtectStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchContainerProtectStatusResponse), nil
	}
}

type SwitchDecoyPortHostPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchDecoyPortHostPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchDecoyPortHostPolicyInvoker) Invoke() (*model.SwitchDecoyPortHostPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchDecoyPortHostPolicyResponse), nil
	}
}

type SwitchHostsProtectStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchHostsProtectStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchHostsProtectStatusInvoker) Invoke() (*model.SwitchHostsProtectStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchHostsProtectStatusResponse), nil
	}
}

type SwitchRaspInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchRaspInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchRaspInvoker) Invoke() (*model.SwitchRaspResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchRaspResponse), nil
	}
}

type SyncClusterListInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncClusterListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SyncClusterListInvoker) Invoke() (*model.SyncClusterListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncClusterListResponse), nil
	}
}

type SyncClusterProtectionEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncClusterProtectionEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SyncClusterProtectionEventInvoker) Invoke() (*model.SyncClusterProtectionEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncClusterProtectionEventResponse), nil
	}
}

type SyncContainerNetworkNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncContainerNetworkNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SyncContainerNetworkNodeInvoker) Invoke() (*model.SyncContainerNetworkNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncContainerNetworkNodeResponse), nil
	}
}

type SyncContainerNetworkPolicyListInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncContainerNetworkPolicyListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SyncContainerNetworkPolicyListInvoker) Invoke() (*model.SyncContainerNetworkPolicyListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncContainerNetworkPolicyListResponse), nil
	}
}

type SyncSecurityGroupPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncSecurityGroupPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SyncSecurityGroupPoliciesInvoker) Invoke() (*model.SyncSecurityGroupPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncSecurityGroupPoliciesResponse), nil
	}
}

type UpdateBackupPolicyInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBackupPolicyInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateBackupPolicyInfoInvoker) Invoke() (*model.UpdateBackupPolicyInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBackupPolicyInfoResponse), nil
	}
}

type UpdateContainerNetworkPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateContainerNetworkPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateContainerNetworkPolicyInvoker) Invoke() (*model.UpdateContainerNetworkPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateContainerNetworkPolicyResponse), nil
	}
}

type UpdatePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePolicyInvoker) Invoke() (*model.UpdatePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePolicyResponse), nil
	}
}

type UpdateProtectionPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProtectionPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProtectionPolicyInvoker) Invoke() (*model.UpdateProtectionPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProtectionPolicyResponse), nil
	}
}

type UpdateSecurityGroupPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecurityGroupPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSecurityGroupPolicyInvoker) Invoke() (*model.UpdateSecurityGroupPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecurityGroupPolicyResponse), nil
	}
}

type UpdateSystemUserWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSystemUserWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSystemUserWhiteListInvoker) Invoke() (*model.UpdateSystemUserWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSystemUserWhiteListResponse), nil
	}
}

type ChangeAntivirusPayPerScanStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeAntivirusPayPerScanStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeAntivirusPayPerScanStatusInvoker) Invoke() (*model.ChangeAntivirusPayPerScanStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeAntivirusPayPerScanStatusResponse), nil
	}
}

type ChangeAntivirusPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeAntivirusPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeAntivirusPolicyInvoker) Invoke() (*model.ChangeAntivirusPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeAntivirusPolicyResponse), nil
	}
}

type CreateAntiVirusPaidTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAntiVirusPaidTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAntiVirusPaidTaskInvoker) Invoke() (*model.CreateAntiVirusPaidTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAntiVirusPaidTaskResponse), nil
	}
}

type CreateAntiVirusPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAntiVirusPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAntiVirusPolicyInvoker) Invoke() (*model.CreateAntiVirusPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAntiVirusPolicyResponse), nil
	}
}

type CreateAntiVirusTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAntiVirusTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAntiVirusTaskInvoker) Invoke() (*model.CreateAntiVirusTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAntiVirusTaskResponse), nil
	}
}

type DeleteAntivirusPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAntivirusPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAntivirusPolicyInvoker) Invoke() (*model.DeleteAntivirusPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAntivirusPolicyResponse), nil
	}
}

type ExportAntiVirusResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportAntiVirusResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportAntiVirusResultInvoker) Invoke() (*model.ExportAntiVirusResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportAntiVirusResultResponse), nil
	}
}

type HandleAntiVirusResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *HandleAntiVirusResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *HandleAntiVirusResultInvoker) Invoke() (*model.HandleAntiVirusResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.HandleAntiVirusResultResponse), nil
	}
}

type ListAntiVirusHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAntiVirusHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAntiVirusHostInvoker) Invoke() (*model.ListAntiVirusHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAntiVirusHostResponse), nil
	}
}

type ListAntiVirusPaidHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAntiVirusPaidHostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAntiVirusPaidHostsInvoker) Invoke() (*model.ListAntiVirusPaidHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAntiVirusPaidHostsResponse), nil
	}
}

type ListAntiVirusPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAntiVirusPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAntiVirusPolicyInvoker) Invoke() (*model.ListAntiVirusPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAntiVirusPolicyResponse), nil
	}
}

type ListAntiVirusResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAntiVirusResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAntiVirusResultInvoker) Invoke() (*model.ListAntiVirusResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAntiVirusResultResponse), nil
	}
}

type ListAntiVirusTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAntiVirusTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAntiVirusTaskInvoker) Invoke() (*model.ListAntiVirusTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAntiVirusTaskResponse), nil
	}
}

type ShowAntivirusFreeQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAntivirusFreeQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAntivirusFreeQuotaInvoker) Invoke() (*model.ShowAntivirusFreeQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAntivirusFreeQuotaResponse), nil
	}
}

type ShowAntivirusPayPerScanStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAntivirusPayPerScanStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAntivirusPayPerScanStatusInvoker) Invoke() (*model.ShowAntivirusPayPerScanStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAntivirusPayPerScanStatusResponse), nil
	}
}

type ShowAntivirusStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAntivirusStatisticInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAntivirusStatisticInvoker) Invoke() (*model.ShowAntivirusStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAntivirusStatisticResponse), nil
	}
}

type SwitchAntivirusTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchAntivirusTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchAntivirusTaskInvoker) Invoke() (*model.SwitchAntivirusTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchAntivirusTaskResponse), nil
	}
}

type AddAppWhitelistPolicyHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddAppWhitelistPolicyHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddAppWhitelistPolicyHostInvoker) Invoke() (*model.AddAppWhitelistPolicyHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddAppWhitelistPolicyHostResponse), nil
	}
}

type AddAppWhitelistPolicyProcessInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddAppWhitelistPolicyProcessInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddAppWhitelistPolicyProcessInvoker) Invoke() (*model.AddAppWhitelistPolicyProcessResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddAppWhitelistPolicyProcessResponse), nil
	}
}

type ChangeAppWhitelistPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeAppWhitelistPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeAppWhitelistPolicyInvoker) Invoke() (*model.ChangeAppWhitelistPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeAppWhitelistPolicyResponse), nil
	}
}

type ChangeAppWhitelistPolicyProcessStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeAppWhitelistPolicyProcessStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeAppWhitelistPolicyProcessStatusInvoker) Invoke() (*model.ChangeAppWhitelistPolicyProcessStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeAppWhitelistPolicyProcessStatusResponse), nil
	}
}

type CreateAppWhitelistPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppWhitelistPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAppWhitelistPolicyInvoker) Invoke() (*model.CreateAppWhitelistPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppWhitelistPolicyResponse), nil
	}
}

type DeleteAppWhitelistPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppWhitelistPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAppWhitelistPolicyInvoker) Invoke() (*model.DeleteAppWhitelistPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppWhitelistPolicyResponse), nil
	}
}

type DeleteAppWhitelistPolicyHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppWhitelistPolicyHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAppWhitelistPolicyHostInvoker) Invoke() (*model.DeleteAppWhitelistPolicyHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppWhitelistPolicyHostResponse), nil
	}
}

type ListAppWhitelistEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppWhitelistEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppWhitelistEventInvoker) Invoke() (*model.ListAppWhitelistEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppWhitelistEventResponse), nil
	}
}

type ListAppWhitelistHostStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppWhitelistHostStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppWhitelistHostStatusInvoker) Invoke() (*model.ListAppWhitelistHostStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppWhitelistHostStatusResponse), nil
	}
}

type ListAppWhitelistPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppWhitelistPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppWhitelistPolicyInvoker) Invoke() (*model.ListAppWhitelistPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppWhitelistPolicyResponse), nil
	}
}

type ListAppWhitelistPolicyHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppWhitelistPolicyHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppWhitelistPolicyHostInvoker) Invoke() (*model.ListAppWhitelistPolicyHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppWhitelistPolicyHostResponse), nil
	}
}

type ListAppWhitelistPolicyProcessInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppWhitelistPolicyProcessInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppWhitelistPolicyProcessInvoker) Invoke() (*model.ListAppWhitelistPolicyProcessResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppWhitelistPolicyProcessResponse), nil
	}
}

type ListAppWhitelistPolicyProcessExtendInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppWhitelistPolicyProcessExtendInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppWhitelistPolicyProcessExtendInvoker) Invoke() (*model.ListAppWhitelistPolicyProcessExtendResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppWhitelistPolicyProcessExtendResponse), nil
	}
}

type ShowAppWhitelistAgentStaticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppWhitelistAgentStaticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppWhitelistAgentStaticsInvoker) Invoke() (*model.ShowAppWhitelistAgentStaticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppWhitelistAgentStaticsResponse), nil
	}
}

type ShowAppWhitelistPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppWhitelistPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppWhitelistPolicyInvoker) Invoke() (*model.ShowAppWhitelistPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppWhitelistPolicyResponse), nil
	}
}

type SwitchAppWhitelistPolicyHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchAppWhitelistPolicyHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchAppWhitelistPolicyHostInvoker) Invoke() (*model.SwitchAppWhitelistPolicyHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchAppWhitelistPolicyHostResponse), nil
	}
}

type SwitchAppWhitelistPolicyLearnStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchAppWhitelistPolicyLearnStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchAppWhitelistPolicyLearnStatusInvoker) Invoke() (*model.SwitchAppWhitelistPolicyLearnStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchAppWhitelistPolicyLearnStatusResponse), nil
	}
}

type ListHandleAffectBaselineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHandleAffectBaselineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHandleAffectBaselineInvoker) Invoke() (*model.ListHandleAffectBaselineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHandleAffectBaselineResponse), nil
	}
}

type CopyBaselinePolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyBaselinePolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CopyBaselinePolicyGroupInvoker) Invoke() (*model.CopyBaselinePolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyBaselinePolicyGroupResponse), nil
	}
}

type ShowBaselineDirectoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBaselineDirectoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBaselineDirectoryInvoker) Invoke() (*model.ShowBaselineDirectoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBaselineDirectoryResponse), nil
	}
}

type ListClusterRiskAffectResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterRiskAffectResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClusterRiskAffectResourcesInvoker) Invoke() (*model.ListClusterRiskAffectResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterRiskAffectResourcesResponse), nil
	}
}

type ListClusterRisksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterRisksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClusterRisksInvoker) Invoke() (*model.ListClusterRisksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterRisksResponse), nil
	}
}

type ShowClusterScanStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterScanStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClusterScanStatisticsInvoker) Invoke() (*model.ShowClusterScanStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterScanStatisticsResponse), nil
	}
}

type ListProjectConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectConfigsInvoker) Invoke() (*model.ListProjectConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectConfigsResponse), nil
	}
}

type ModifyProjectConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyProjectConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyProjectConfigsInvoker) Invoke() (*model.ModifyProjectConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyProjectConfigsResponse), nil
	}
}

type SaveBrowsingHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *SaveBrowsingHistoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SaveBrowsingHistoryInvoker) Invoke() (*model.SaveBrowsingHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SaveBrowsingHistoryResponse), nil
	}
}

type BatchDeleteAgentDaemonsetInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteAgentDaemonsetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteAgentDaemonsetInvoker) Invoke() (*model.BatchDeleteAgentDaemonsetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteAgentDaemonsetResponse), nil
	}
}

type BatchUpgradeAgentDaemonsetInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpgradeAgentDaemonsetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpgradeAgentDaemonsetInvoker) Invoke() (*model.BatchUpgradeAgentDaemonsetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpgradeAgentDaemonsetResponse), nil
	}
}

type CreateAgentDaemonsetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAgentDaemonsetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAgentDaemonsetInvoker) Invoke() (*model.CreateAgentDaemonsetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAgentDaemonsetResponse), nil
	}
}

type CreateMultiCloudClustersInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMultiCloudClustersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMultiCloudClustersInvoker) Invoke() (*model.CreateMultiCloudClustersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMultiCloudClustersResponse), nil
	}
}

type DeleteAgentDaemonsetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAgentDaemonsetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAgentDaemonsetInvoker) Invoke() (*model.DeleteAgentDaemonsetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAgentDaemonsetResponse), nil
	}
}

type DeleteCicdConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCicdConfigurationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCicdConfigurationsInvoker) Invoke() (*model.DeleteCicdConfigurationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCicdConfigurationsResponse), nil
	}
}

type ExportImageSecurityReportTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportImageSecurityReportTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportImageSecurityReportTaskInvoker) Invoke() (*model.ExportImageSecurityReportTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportImageSecurityReportTaskResponse), nil
	}
}

type ListAssociateRegistriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAssociateRegistriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAssociateRegistriesInvoker) Invoke() (*model.ListAssociateRegistriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAssociateRegistriesResponse), nil
	}
}

type ListCicdConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCicdConfigurationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCicdConfigurationsInvoker) Invoke() (*model.ListCicdConfigurationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCicdConfigurationsResponse), nil
	}
}

type ListMultiCloudClustersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMultiCloudClustersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMultiCloudClustersInvoker) Invoke() (*model.ListMultiCloudClustersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMultiCloudClustersResponse), nil
	}
}

type ModifyCicdConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyCicdConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyCicdConfigurationInvoker) Invoke() (*model.ModifyCicdConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyCicdConfigurationResponse), nil
	}
}

type ParseMultiCloudClusterConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ParseMultiCloudClusterConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ParseMultiCloudClusterConfigInvoker) Invoke() (*model.ParseMultiCloudClusterConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ParseMultiCloudClusterConfigResponse), nil
	}
}

type RemoveMultiCloudClustersInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveMultiCloudClustersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RemoveMultiCloudClustersInvoker) Invoke() (*model.RemoveMultiCloudClustersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveMultiCloudClustersResponse), nil
	}
}

type ShowAgentDaemonsetDeployTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAgentDaemonsetDeployTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAgentDaemonsetDeployTemplateInvoker) Invoke() (*model.ShowAgentDaemonsetDeployTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgentDaemonsetDeployTemplateResponse), nil
	}
}

type ShowAgentDaemonsetDetailInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAgentDaemonsetDetailInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAgentDaemonsetDetailInfoInvoker) Invoke() (*model.ShowAgentDaemonsetDetailInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgentDaemonsetDetailInfoResponse), nil
	}
}

type ShowCicdConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCicdConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCicdConfigurationInvoker) Invoke() (*model.ShowCicdConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCicdConfigurationResponse), nil
	}
}

type ShowMultiCloudClusterImageCommandInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMultiCloudClusterImageCommandInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMultiCloudClusterImageCommandInvoker) Invoke() (*model.ShowMultiCloudClusterImageCommandResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMultiCloudClusterImageCommandResponse), nil
	}
}

type ShowMultiCloudClusterProxyScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMultiCloudClusterProxyScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMultiCloudClusterProxyScriptInvoker) Invoke() (*model.ShowMultiCloudClusterProxyScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMultiCloudClusterProxyScriptResponse), nil
	}
}

type SyncMultiCloudClusterStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncMultiCloudClusterStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SyncMultiCloudClusterStatusInvoker) Invoke() (*model.SyncMultiCloudClusterStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncMultiCloudClusterStatusResponse), nil
	}
}

type UpdateAgentDaemonsetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAgentDaemonsetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAgentDaemonsetInvoker) Invoke() (*model.UpdateAgentDaemonsetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAgentDaemonsetResponse), nil
	}
}

type UpdateMultiCloudClustersInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMultiCloudClustersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMultiCloudClustersInvoker) Invoke() (*model.UpdateMultiCloudClustersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMultiCloudClustersResponse), nil
	}
}

type ListFileEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFileEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFileEventsInvoker) Invoke() (*model.ListFileEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFileEventsResponse), nil
	}
}

type ListFileHostEventDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFileHostEventDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFileHostEventDetailsInvoker) Invoke() (*model.ListFileHostEventDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFileHostEventDetailsResponse), nil
	}
}

type ListFileHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFileHostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFileHostsInvoker) Invoke() (*model.ListFileHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFileHostsResponse), nil
	}
}

type ShowFileStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFileStatisticInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFileStatisticInvoker) Invoke() (*model.ShowFileStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFileStatisticResponse), nil
	}
}

type ListIacFileRiskPathsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIacFileRiskPathsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIacFileRiskPathsInvoker) Invoke() (*model.ListIacFileRiskPathsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIacFileRiskPathsResponse), nil
	}
}

type ListIacFileRisksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIacFileRisksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIacFileRisksInvoker) Invoke() (*model.ListIacFileRisksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIacFileRisksResponse), nil
	}
}

type ListIacFilesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIacFilesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIacFilesInvoker) Invoke() (*model.ListIacFilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIacFilesResponse), nil
	}
}

type CreateTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTaskInvoker) Invoke() (*model.CreateTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTaskResponse), nil
	}
}

type ListTaskResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTaskResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTaskResourcesInvoker) Invoke() (*model.ListTaskResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTaskResourcesResponse), nil
	}
}

type ListTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTasksInvoker) Invoke() (*model.ListTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTasksResponse), nil
	}
}

type ShowTaskStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskStatisticsInvoker) Invoke() (*model.ShowTaskStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskStatisticsResponse), nil
	}
}

type ChangeVulScanPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeVulScanPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeVulScanPolicyInvoker) Invoke() (*model.ChangeVulScanPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeVulScanPolicyResponse), nil
	}
}

type CreateVulnerabilityScanTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVulnerabilityScanTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateVulnerabilityScanTaskInvoker) Invoke() (*model.CreateVulnerabilityScanTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVulnerabilityScanTaskResponse), nil
	}
}

type ExportHandledVulnerabilitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportHandledVulnerabilitiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportHandledVulnerabilitiesInvoker) Invoke() (*model.ExportHandledVulnerabilitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportHandledVulnerabilitiesResponse), nil
	}
}

type ExportVulHandleHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportVulHandleHistoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportVulHandleHistoryInvoker) Invoke() (*model.ExportVulHandleHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportVulHandleHistoryResponse), nil
	}
}

type ExportVulsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportVulsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportVulsInvoker) Invoke() (*model.ExportVulsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportVulsResponse), nil
	}
}

type ListHostVulsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostVulsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHostVulsInvoker) Invoke() (*model.ListHostVulsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostVulsResponse), nil
	}
}

type ListVulContainerAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulContainerAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulContainerAppsInvoker) Invoke() (*model.ListVulContainerAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulContainerAppsResponse), nil
	}
}

type ListVulContainersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulContainersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulContainersInvoker) Invoke() (*model.ListVulContainersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulContainersResponse), nil
	}
}

type ListVulHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulHostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulHostsInvoker) Invoke() (*model.ListVulHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulHostsResponse), nil
	}
}

type ListVulScanTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulScanTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulScanTaskInvoker) Invoke() (*model.ListVulScanTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulScanTaskResponse), nil
	}
}

type ListVulScanTaskHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulScanTaskHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulScanTaskHostInvoker) Invoke() (*model.ListVulScanTaskHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulScanTaskHostResponse), nil
	}
}

type ListVulnerabilitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulnerabilitiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulnerabilitiesInvoker) Invoke() (*model.ListVulnerabilitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulnerabilitiesResponse), nil
	}
}

type RecordUserViewVulTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecordUserViewVulTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecordUserViewVulTaskInvoker) Invoke() (*model.RecordUserViewVulTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecordUserViewVulTaskResponse), nil
	}
}

type ShowVulScanPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVulScanPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVulScanPolicyInvoker) Invoke() (*model.ShowVulScanPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVulScanPolicyResponse), nil
	}
}

type ShowVulStaticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVulStaticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVulStaticsInvoker) Invoke() (*model.ShowVulStaticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVulStaticsResponse), nil
	}
}

type ShowVulTaskStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVulTaskStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVulTaskStatisticsInvoker) Invoke() (*model.ShowVulTaskStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVulTaskStatisticsResponse), nil
	}
}

type BatchStartWebTamperProtectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchStartWebTamperProtectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchStartWebTamperProtectionInvoker) Invoke() (*model.BatchStartWebTamperProtectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchStartWebTamperProtectionResponse), nil
	}
}

type ExportWebTamperHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportWebTamperHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportWebTamperHostInvoker) Invoke() (*model.ExportWebTamperHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportWebTamperHostResponse), nil
	}
}

type ListHostProtectHistoryInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostProtectHistoryInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHostProtectHistoryInfoInvoker) Invoke() (*model.ListHostProtectHistoryInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostProtectHistoryInfoResponse), nil
	}
}

type ListHostRaspProtectHistoryInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostRaspProtectHistoryInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHostRaspProtectHistoryInfoInvoker) Invoke() (*model.ListHostRaspProtectHistoryInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostRaspProtectHistoryInfoResponse), nil
	}
}

type ListWebTamperHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWebTamperHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWebTamperHostInvoker) Invoke() (*model.ListWebTamperHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWebTamperHostResponse), nil
	}
}

type ListWtpProtectHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWtpProtectHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWtpProtectHostInvoker) Invoke() (*model.ListWtpProtectHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWtpProtectHostResponse), nil
	}
}

type SetRaspSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetRaspSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetRaspSwitchInvoker) Invoke() (*model.SetRaspSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetRaspSwitchResponse), nil
	}
}

type SetWtpProtectionStatusInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetWtpProtectionStatusInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetWtpProtectionStatusInfoInvoker) Invoke() (*model.SetWtpProtectionStatusInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetWtpProtectionStatusInfoResponse), nil
	}
}

type ShowWebTamperHostPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWebTamperHostPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWebTamperHostPolicyInvoker) Invoke() (*model.ShowWebTamperHostPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWebTamperHostPolicyResponse), nil
	}
}

type ShowWebTamperRaspPathInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWebTamperRaspPathInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWebTamperRaspPathInvoker) Invoke() (*model.ShowWebTamperRaspPathResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWebTamperRaspPathResponse), nil
	}
}

type UpdateWebTamperHostPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWebTamperHostPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWebTamperHostPolicyInvoker) Invoke() (*model.UpdateWebTamperHostPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWebTamperHostPolicyResponse), nil
	}
}

type UpdateWebTamperRaspPathInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWebTamperRaspPathInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWebTamperRaspPathInvoker) Invoke() (*model.UpdateWebTamperRaspPathResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWebTamperRaspPathResponse), nil
	}
}
