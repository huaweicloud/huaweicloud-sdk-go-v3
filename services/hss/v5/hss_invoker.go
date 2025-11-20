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

type AddVulWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddVulWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddVulWhiteListInvoker) Invoke() (*model.AddVulWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddVulWhiteListResponse), nil
	}
}

type AssociateHostAssetValueInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateHostAssetValueInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateHostAssetValueInvoker) Invoke() (*model.AssociateHostAssetValueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateHostAssetValueResponse), nil
	}
}

type AssociateHostsGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateHostsGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateHostsGroupInvoker) Invoke() (*model.AssociateHostsGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateHostsGroupResponse), nil
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

type BatchChangeEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchChangeEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchChangeEventInvoker) Invoke() (*model.BatchChangeEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchChangeEventResponse), nil
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

type BatchDeleteTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteTagsInvoker) Invoke() (*model.BatchDeleteTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteTagsResponse), nil
	}
}

type BatchModifyPortStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchModifyPortStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchModifyPortStatusInvoker) Invoke() (*model.BatchModifyPortStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchModifyPortStatusResponse), nil
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

type ChangeContainerStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeContainerStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeContainerStatusInvoker) Invoke() (*model.ChangeContainerStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeContainerStatusResponse), nil
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

type ChangeHostIgnoreStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeHostIgnoreStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeHostIgnoreStatusInvoker) Invoke() (*model.ChangeHostIgnoreStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeHostIgnoreStatusResponse), nil
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

type ChangeMalwareCollectStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeMalwareCollectStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeMalwareCollectStatusInvoker) Invoke() (*model.ChangeMalwareCollectStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeMalwareCollectStatusResponse), nil
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

type DownloadAssetFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadAssetFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadAssetFileInvoker) Invoke() (*model.DownloadAssetFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadAssetFileResponse), nil
	}
}

type DownloadEventSourceFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadEventSourceFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadEventSourceFileInvoker) Invoke() (*model.DownloadEventSourceFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadEventSourceFileResponse), nil
	}
}

type EnableTrustServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableTrustServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableTrustServiceInvoker) Invoke() (*model.EnableTrustServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableTrustServiceResponse), nil
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

type ExportEventRequestInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportEventRequestInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportEventRequestInvoker) Invoke() (*model.ExportEventRequestResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportEventRequestResponse), nil
	}
}

type ExportSecurityCheckReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportSecurityCheckReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportSecurityCheckReportInvoker) Invoke() (*model.ExportSecurityCheckReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportSecurityCheckReportResponse), nil
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

type ListAllRiskConfigCheckRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllRiskConfigCheckRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllRiskConfigCheckRulesInvoker) Invoke() (*model.ListAllRiskConfigCheckRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllRiskConfigCheckRulesResponse), nil
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

type ListDecoyPortAvailableHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDecoyPortAvailableHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDecoyPortAvailableHostInvoker) Invoke() (*model.ListDecoyPortAvailableHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDecoyPortAvailableHostResponse), nil
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

type ListDockerPluginsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDockerPluginsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDockerPluginsInvoker) Invoke() (*model.ListDockerPluginsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDockerPluginsResponse), nil
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

type ListEventAttCkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventAttCkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEventAttCkInvoker) Invoke() (*model.ListEventAttCkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventAttCkResponse), nil
	}
}

type ListEventForensicInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventForensicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEventForensicInvoker) Invoke() (*model.ListEventForensicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventForensicResponse), nil
	}
}

type ListEventOperatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventOperatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEventOperatesInvoker) Invoke() (*model.ListEventOperatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventOperatesResponse), nil
	}
}

type ListEventTopRiskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventTopRiskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEventTopRiskInvoker) Invoke() (*model.ListEventTopRiskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventTopRiskResponse), nil
	}
}

type ListEventTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEventTypeInvoker) Invoke() (*model.ListEventTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventTypeResponse), nil
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

type ListHostCheckRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostCheckRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHostCheckRulesInvoker) Invoke() (*model.ListHostCheckRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostCheckRulesResponse), nil
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

type ListJarPackageInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJarPackageInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListJarPackageInfoInvoker) Invoke() (*model.ListJarPackageInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJarPackageInfoResponse), nil
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

type ListKernelModuleInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListKernelModuleInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListKernelModuleInfoInvoker) Invoke() (*model.ListKernelModuleInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListKernelModuleInfoResponse), nil
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

type ListMalwareCollectStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMalwareCollectStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMalwareCollectStatusInvoker) Invoke() (*model.ListMalwareCollectStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMalwareCollectStatusResponse), nil
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

type ListPluginInstallScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPluginInstallScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPluginInstallScriptInvoker) Invoke() (*model.ListPluginInstallScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPluginInstallScriptResponse), nil
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

type ListResourceInstanceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceInstanceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceInstanceTagInvoker) Invoke() (*model.ListResourceInstanceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceInstanceTagResponse), nil
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

type ListSameEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSameEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSameEventsInvoker) Invoke() (*model.ListSameEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSameEventsResponse), nil
	}
}

type ListSecurityCheckPolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityCheckPolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityCheckPolicyGroupInvoker) Invoke() (*model.ListSecurityCheckPolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityCheckPolicyGroupResponse), nil
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

type ListSimilarHandledEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSimilarHandledEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSimilarHandledEventsInvoker) Invoke() (*model.ListSimilarHandledEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSimilarHandledEventsResponse), nil
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

type ListTrustServiceStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTrustServiceStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTrustServiceStatusInvoker) Invoke() (*model.ListTrustServiceStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTrustServiceStatusResponse), nil
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

type ListVulHostHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulHostHostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulHostHostsInvoker) Invoke() (*model.ListVulHostHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulHostHostsResponse), nil
	}
}

type ListVulWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulWhiteListInvoker) Invoke() (*model.ListVulWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulWhiteListResponse), nil
	}
}

type ListVulWhiteListVulOptionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulWhiteListVulOptionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulWhiteListVulOptionsInvoker) Invoke() (*model.ListVulWhiteListVulOptionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulWhiteListVulOptionsResponse), nil
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

type ListWebFrameworkInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWebFrameworkInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWebFrameworkInfoInvoker) Invoke() (*model.ListWebFrameworkInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWebFrameworkInfoResponse), nil
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

type ListWebSiteInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWebSiteInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWebSiteInfoInvoker) Invoke() (*model.ListWebSiteInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWebSiteInfoResponse), nil
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

type SetMalwareRemindersInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetMalwareRemindersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetMalwareRemindersInvoker) Invoke() (*model.SetMalwareRemindersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetMalwareRemindersResponse), nil
	}
}

type SetManualDetectInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetManualDetectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetManualDetectInvoker) Invoke() (*model.SetManualDetectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetManualDetectResponse), nil
	}
}

type ShowAccountTopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAccountTopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAccountTopInvoker) Invoke() (*model.ShowAccountTopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAccountTopResponse), nil
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

type ShowAutoLaunchTopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoLaunchTopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAutoLaunchTopInvoker) Invoke() (*model.ShowAutoLaunchTopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAutoLaunchTopResponse), nil
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

type ShowCmsVulDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCmsVulDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCmsVulDetailInvoker) Invoke() (*model.ShowCmsVulDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCmsVulDetailResponse), nil
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

type ShowDecoyPortAutoBindInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDecoyPortAutoBindInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDecoyPortAutoBindInvoker) Invoke() (*model.ShowDecoyPortAutoBindResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDecoyPortAutoBindResponse), nil
	}
}

type ShowDecoyPortHostListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDecoyPortHostListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDecoyPortHostListInvoker) Invoke() (*model.ShowDecoyPortHostListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDecoyPortHostListResponse), nil
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

type ShowEventAttackTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEventAttackTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEventAttackTagInvoker) Invoke() (*model.ShowEventAttackTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEventAttackTagResponse), nil
	}
}

type ShowEventSeverityInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEventSeverityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEventSeverityInvoker) Invoke() (*model.ShowEventSeverityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEventSeverityResponse), nil
	}
}

type ShowExportTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExportTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowExportTaskInvoker) Invoke() (*model.ShowExportTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExportTaskResponse), nil
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

type ShowJarPackageTopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJarPackageTopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJarPackageTopInvoker) Invoke() (*model.ShowJarPackageTopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJarPackageTopResponse), nil
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

type ShowKernelModuleTopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKernelModuleTopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowKernelModuleTopInvoker) Invoke() (*model.ShowKernelModuleTopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKernelModuleTopResponse), nil
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

type ShowLinuxVulDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLinuxVulDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLinuxVulDetailInvoker) Invoke() (*model.ShowLinuxVulDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLinuxVulDetailResponse), nil
	}
}

type ShowMalwareRemindersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMalwareRemindersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMalwareRemindersInvoker) Invoke() (*model.ShowMalwareRemindersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMalwareRemindersResponse), nil
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

type ShowPorcessTopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPorcessTopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPorcessTopInvoker) Invoke() (*model.ShowPorcessTopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPorcessTopResponse), nil
	}
}

type ShowPortTopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPortTopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPortTopInvoker) Invoke() (*model.ShowPortTopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPortTopResponse), nil
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

type ShowScanStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScanStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowScanStatusInvoker) Invoke() (*model.ShowScanStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScanStatusResponse), nil
	}
}

type ShowSoftwareTopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSoftwareTopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSoftwareTopInvoker) Invoke() (*model.ShowSoftwareTopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSoftwareTopResponse), nil
	}
}

type ShowVulAffectedStaticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVulAffectedStaticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVulAffectedStaticsInvoker) Invoke() (*model.ShowVulAffectedStaticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVulAffectedStaticsResponse), nil
	}
}

type ShowVulWhiteListDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVulWhiteListDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVulWhiteListDetailInvoker) Invoke() (*model.ShowVulWhiteListDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVulWhiteListDetailResponse), nil
	}
}

type ShowWebAppAndServiceTopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWebAppAndServiceTopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWebAppAndServiceTopInvoker) Invoke() (*model.ShowWebAppAndServiceTopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWebAppAndServiceTopResponse), nil
	}
}

type ShowWebFrameworkTopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWebFrameworkTopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWebFrameworkTopInvoker) Invoke() (*model.ShowWebFrameworkTopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWebFrameworkTopResponse), nil
	}
}

type ShowWebSiteTopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWebSiteTopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWebSiteTopInvoker) Invoke() (*model.ShowWebSiteTopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWebSiteTopResponse), nil
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

type SwitchDecoyPortAutoBindInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchDecoyPortAutoBindInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchDecoyPortAutoBindInvoker) Invoke() (*model.SwitchDecoyPortAutoBindResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchDecoyPortAutoBindResponse), nil
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

type SwitchDecoyPortPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchDecoyPortPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchDecoyPortPolicyInvoker) Invoke() (*model.SwitchDecoyPortPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchDecoyPortPolicyResponse), nil
	}
}

type SwitchFirewallStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchFirewallStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchFirewallStatusInvoker) Invoke() (*model.SwitchFirewallStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchFirewallStatusResponse), nil
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

type UninstallAgentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UninstallAgentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UninstallAgentsInvoker) Invoke() (*model.UninstallAgentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UninstallAgentsResponse), nil
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

type ValidateAdminInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateAdminInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ValidateAdminInvoker) Invoke() (*model.ValidateAdminResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateAdminResponse), nil
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

type AddSecurityCheckPolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddSecurityCheckPolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddSecurityCheckPolicyGroupInvoker) Invoke() (*model.AddSecurityCheckPolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddSecurityCheckPolicyGroupResponse), nil
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

type ExportBaselineSecurityCheckReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportBaselineSecurityCheckReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportBaselineSecurityCheckReportInvoker) Invoke() (*model.ExportBaselineSecurityCheckReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportBaselineSecurityCheckReportResponse), nil
	}
}

type ListCheckRuleHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCheckRuleHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCheckRuleHostInvoker) Invoke() (*model.ListCheckRuleHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCheckRuleHostResponse), nil
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

type RunBaselineDetectInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunBaselineDetectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunBaselineDetectInvoker) Invoke() (*model.RunBaselineDetectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunBaselineDetectResponse), nil
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

type ShowDefaultSecurityCheckPolicyDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDefaultSecurityCheckPolicyDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDefaultSecurityCheckPolicyDetailsInvoker) Invoke() (*model.ShowDefaultSecurityCheckPolicyDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDefaultSecurityCheckPolicyDetailsResponse), nil
	}
}

type UpdateSecurityCheckPolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecurityCheckPolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSecurityCheckPolicyGroupInvoker) Invoke() (*model.UpdateSecurityCheckPolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecurityCheckPolicyGroupResponse), nil
	}
}

type ShowBaselineOverviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBaselineOverviewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBaselineOverviewInvoker) Invoke() (*model.ShowBaselineOverviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBaselineOverviewResponse), nil
	}
}

type ShowBaselineStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBaselineStatisticInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBaselineStatisticInvoker) Invoke() (*model.ShowBaselineStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBaselineStatisticResponse), nil
	}
}

type ShowCheckRuleFixFailDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCheckRuleFixFailDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCheckRuleFixFailDetailInvoker) Invoke() (*model.ShowCheckRuleFixFailDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCheckRuleFixFailDetailResponse), nil
	}
}

type DeleteSecurityCheckPolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecurityCheckPolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSecurityCheckPolicyGroupInvoker) Invoke() (*model.DeleteSecurityCheckPolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecurityCheckPolicyGroupResponse), nil
	}
}

type ShowDefaultSecurityCheckPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDefaultSecurityCheckPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDefaultSecurityCheckPolicyInvoker) Invoke() (*model.ShowDefaultSecurityCheckPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDefaultSecurityCheckPolicyResponse), nil
	}
}

type ExportRisksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportRisksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportRisksInvoker) Invoke() (*model.ExportRisksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportRisksResponse), nil
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

type UploadReportLogoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadReportLogoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadReportLogoInvoker) Invoke() (*model.UploadReportLogoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadReportLogoResponse), nil
	}
}

type AddCicdConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddCicdConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddCicdConfigurationInvoker) Invoke() (*model.AddCicdConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddCicdConfigurationResponse), nil
	}
}

type AddRegistryInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddRegistryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddRegistryInvoker) Invoke() (*model.AddRegistryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddRegistryResponse), nil
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

type BatchDeleteRegistryInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteRegistryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteRegistryInvoker) Invoke() (*model.BatchDeleteRegistryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteRegistryResponse), nil
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

type CreateCicdConfigurationCommandInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCicdConfigurationCommandInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCicdConfigurationCommandInvoker) Invoke() (*model.CreateCicdConfigurationCommandResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCicdConfigurationCommandResponse), nil
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

type DeleteRegistryInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRegistryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRegistryInvoker) Invoke() (*model.DeleteRegistryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRegistryResponse), nil
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

type ListAgentDaemonsetClusterNodesLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAgentDaemonsetClusterNodesLabelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAgentDaemonsetClusterNodesLabelInvoker) Invoke() (*model.ListAgentDaemonsetClusterNodesLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAgentDaemonsetClusterNodesLabelResponse), nil
	}
}

type ListAgentDaemonsetInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAgentDaemonsetInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAgentDaemonsetInfoInvoker) Invoke() (*model.ListAgentDaemonsetInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAgentDaemonsetInfoResponse), nil
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

type ListCceNodesLabelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCceNodesLabelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCceNodesLabelInvoker) Invoke() (*model.ListCceNodesLabelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCceNodesLabelResponse), nil
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

type ListCicdImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCicdImagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCicdImagesInvoker) Invoke() (*model.ListCicdImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCicdImagesResponse), nil
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

type ListRegistryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRegistryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRegistryInvoker) Invoke() (*model.ListRegistryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRegistryResponse), nil
	}
}

type ListRegistryStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRegistryStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRegistryStatisticsInvoker) Invoke() (*model.ListRegistryStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRegistryStatisticsResponse), nil
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

type ShowAgentAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAgentAddressInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAgentAddressInvoker) Invoke() (*model.ShowAgentAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgentAddressResponse), nil
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

type ShowImageUploadCommandInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageUploadCommandInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowImageUploadCommandInvoker) Invoke() (*model.ShowImageUploadCommandResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageUploadCommandResponse), nil
	}
}

type ShowMultiCloudClusterAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMultiCloudClusterAuthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMultiCloudClusterAuthInvoker) Invoke() (*model.ShowMultiCloudClusterAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMultiCloudClusterAuthResponse), nil
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

type UpdateRegistryInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRegistryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRegistryInvoker) Invoke() (*model.UpdateRegistryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRegistryResponse), nil
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

type ChangeAutoOpenQuotaStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeAutoOpenQuotaStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeAutoOpenQuotaStatusInvoker) Invoke() (*model.ChangeAutoOpenQuotaStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeAutoOpenQuotaStatusResponse), nil
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

type CreateVpcEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVpcEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateVpcEndpointInvoker) Invoke() (*model.CreateVpcEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVpcEndpointResponse), nil
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

type ListAutoOpenQuotaStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAutoOpenQuotaStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAutoOpenQuotaStatusInvoker) Invoke() (*model.ListAutoOpenQuotaStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAutoOpenQuotaStatusResponse), nil
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

type ListHostsRiskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostsRiskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHostsRiskInvoker) Invoke() (*model.ListHostsRiskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostsRiskResponse), nil
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

type ShowEndpointStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEndpointStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEndpointStatusInvoker) Invoke() (*model.ShowEndpointStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEndpointStatusResponse), nil
	}
}

type ShowHostsStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHostsStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHostsStatisticsInvoker) Invoke() (*model.ShowHostsStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHostsStatisticsResponse), nil
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

type UpgradeAgentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeAgentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpgradeAgentsInvoker) Invoke() (*model.UpgradeAgentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeAgentsResponse), nil
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

type AddImageWhiteListsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddImageWhiteListsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddImageWhiteListsInvoker) Invoke() (*model.AddImageWhiteListsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddImageWhiteListsResponse), nil
	}
}

type BatchExportBaselineTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchExportBaselineTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchExportBaselineTaskInvoker) Invoke() (*model.BatchExportBaselineTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchExportBaselineTaskResponse), nil
	}
}

type BatchExportLocalVulTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchExportLocalVulTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchExportLocalVulTaskInvoker) Invoke() (*model.BatchExportLocalVulTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchExportLocalVulTaskResponse), nil
	}
}

type BatchScanLocalImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchScanLocalImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchScanLocalImageInvoker) Invoke() (*model.BatchScanLocalImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchScanLocalImageResponse), nil
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

type ChangeExtendedWeakPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeExtendedWeakPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeExtendedWeakPasswordInvoker) Invoke() (*model.ChangeExtendedWeakPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeExtendedWeakPasswordResponse), nil
	}
}

type ChangeImageWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeImageWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeImageWhiteListInvoker) Invoke() (*model.ChangeImageWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeImageWhiteListResponse), nil
	}
}

type DeleteImageWhiteListsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteImageWhiteListsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteImageWhiteListsInvoker) Invoke() (*model.DeleteImageWhiteListsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteImageWhiteListsResponse), nil
	}
}

type HandleImageVulnerabilityInvoker struct {
	*invoker.BaseInvoker
}

func (i *HandleImageVulnerabilityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *HandleImageVulnerabilityInvoker) Invoke() (*model.HandleImageVulnerabilityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.HandleImageVulnerabilityResponse), nil
	}
}

type ListCheckRuleResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCheckRuleResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCheckRuleResourcesInvoker) Invoke() (*model.ListCheckRuleResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCheckRuleResourcesResponse), nil
	}
}

type ListCheckRulesInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCheckRulesInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCheckRulesInfoInvoker) Invoke() (*model.ListCheckRulesInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCheckRulesInfoResponse), nil
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

type ListGlobalImageAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalImageAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGlobalImageAppsInvoker) Invoke() (*model.ListGlobalImageAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalImageAppsResponse), nil
	}
}

type ListGlobalImageFilesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalImageFilesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGlobalImageFilesInvoker) Invoke() (*model.ListGlobalImageFilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalImageFilesResponse), nil
	}
}

type ListGlobalMalwareInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalMalwareInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGlobalMalwareInvoker) Invoke() (*model.ListGlobalMalwareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalMalwareResponse), nil
	}
}

type ListGlobalVulnerabilitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalVulnerabilitiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGlobalVulnerabilitiesInvoker) Invoke() (*model.ListGlobalVulnerabilitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalVulnerabilitiesResponse), nil
	}
}

type ListImageAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageAppsInvoker) Invoke() (*model.ListImageAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageAppsResponse), nil
	}
}

type ListImageBasicImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageBasicImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageBasicImageInvoker) Invoke() (*model.ListImageBasicImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageBasicImageResponse), nil
	}
}

type ListImageBuildCommandRisksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageBuildCommandRisksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageBuildCommandRisksInvoker) Invoke() (*model.ListImageBuildCommandRisksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageBuildCommandRisksResponse), nil
	}
}

type ListImageBuildCommandRisksImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageBuildCommandRisksImagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageBuildCommandRisksImagesInvoker) Invoke() (*model.ListImageBuildCommandRisksImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageBuildCommandRisksImagesResponse), nil
	}
}

type ListImageFilesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageFilesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageFilesInvoker) Invoke() (*model.ListImageFilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageFilesResponse), nil
	}
}

type ListImageHandleAffectVulnerabilitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageHandleAffectVulnerabilitiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageHandleAffectVulnerabilitiesInvoker) Invoke() (*model.ListImageHandleAffectVulnerabilitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageHandleAffectVulnerabilitiesResponse), nil
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

type ListImageMalwareInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageMalwareInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageMalwareInvoker) Invoke() (*model.ListImageMalwareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageMalwareResponse), nil
	}
}

type ListImageNonCompliantAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageNonCompliantAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageNonCompliantAppInvoker) Invoke() (*model.ListImageNonCompliantAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageNonCompliantAppResponse), nil
	}
}

type ListImagePwdComplexityInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImagePwdComplexityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImagePwdComplexityInvoker) Invoke() (*model.ListImagePwdComplexityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImagePwdComplexityResponse), nil
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

type ListImageSensitiveInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageSensitiveInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageSensitiveInvoker) Invoke() (*model.ListImageSensitiveResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageSensitiveResponse), nil
	}
}

type ListImageSensitiveInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageSensitiveInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageSensitiveInfoInvoker) Invoke() (*model.ListImageSensitiveInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageSensitiveInfoResponse), nil
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

type ListImageWeakPwdUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageWeakPwdUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageWeakPwdUsersInvoker) Invoke() (*model.ListImageWeakPwdUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageWeakPwdUsersResponse), nil
	}
}

type ListImageWhiteListsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageWhiteListsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageWhiteListsInvoker) Invoke() (*model.ListImageWhiteListsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageWhiteListsResponse), nil
	}
}

type ListLocalImageAppInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLocalImageAppInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLocalImageAppInfoInvoker) Invoke() (*model.ListLocalImageAppInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLocalImageAppInfoResponse), nil
	}
}

type ListLocalImageContainersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLocalImageContainersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLocalImageContainersInvoker) Invoke() (*model.ListLocalImageContainersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLocalImageContainersResponse), nil
	}
}

type ListLocalImageHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLocalImageHostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLocalImageHostsInvoker) Invoke() (*model.ListLocalImageHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLocalImageHostsResponse), nil
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

type ListVulAffectImageAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulAffectImageAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulAffectImageAppsInvoker) Invoke() (*model.ListVulAffectImageAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulAffectImageAppsResponse), nil
	}
}

type ListVulAffectImageContainersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulAffectImageContainersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulAffectImageContainersInvoker) Invoke() (*model.ListVulAffectImageContainersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulAffectImageContainersResponse), nil
	}
}

type ListVulAffectImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulAffectImagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulAffectImagesInvoker) Invoke() (*model.ListVulAffectImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulAffectImagesResponse), nil
	}
}

type ListVulRepoImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulRepoImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulRepoImageInvoker) Invoke() (*model.ListVulRepoImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulRepoImageResponse), nil
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

type ShowExtendedWeakPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExtendedWeakPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowExtendedWeakPasswordInvoker) Invoke() (*model.ShowExtendedWeakPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExtendedWeakPasswordResponse), nil
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

type ShowImageBaselineStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageBaselineStatisticInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowImageBaselineStatisticInvoker) Invoke() (*model.ShowImageBaselineStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageBaselineStatisticResponse), nil
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

type ShowImageFilesStatInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageFilesStatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowImageFilesStatInvoker) Invoke() (*model.ShowImageFilesStatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageFilesStatResponse), nil
	}
}

type ShowImageSecurityReportStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageSecurityReportStatisticInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowImageSecurityReportStatisticInvoker) Invoke() (*model.ShowImageSecurityReportStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageSecurityReportStatisticResponse), nil
	}
}

type ShowImageWhiteListDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageWhiteListDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowImageWhiteListDetailInvoker) Invoke() (*model.ShowImageWhiteListDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageWhiteListDetailResponse), nil
	}
}

type CreateManualImageScanTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateManualImageScanTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateManualImageScanTaskInvoker) Invoke() (*model.CreateManualImageScanTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateManualImageScanTaskResponse), nil
	}
}

type ListImageScanPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageScanPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageScanPolicyInvoker) Invoke() (*model.ListImageScanPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageScanPolicyResponse), nil
	}
}

type ListImageScanTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageScanTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageScanTaskInvoker) Invoke() (*model.ListImageScanTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageScanTaskResponse), nil
	}
}

type ModifyImageScanPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyImageScanPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyImageScanPolicyInvoker) Invoke() (*model.ModifyImageScanPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyImageScanPolicyResponse), nil
	}
}

type ShowImagePayPerScanStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImagePayPerScanStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowImagePayPerScanStatisticsInvoker) Invoke() (*model.ShowImagePayPerScanStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImagePayPerScanStatisticsResponse), nil
	}
}

type StopImageScanTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopImageScanTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopImageScanTaskInvoker) Invoke() (*model.StopImageScanTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopImageScanTaskResponse), nil
	}
}

type ChangeMonthlyOperationReportTipStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeMonthlyOperationReportTipStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeMonthlyOperationReportTipStatusInvoker) Invoke() (*model.ChangeMonthlyOperationReportTipStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeMonthlyOperationReportTipStatusResponse), nil
	}
}

type ExportTaskInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportTaskInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportTaskInfoInvoker) Invoke() (*model.ExportTaskInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportTaskInfoResponse), nil
	}
}

type ListMonthlyOperationReportsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMonthlyOperationReportsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMonthlyOperationReportsInvoker) Invoke() (*model.ListMonthlyOperationReportsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMonthlyOperationReportsResponse), nil
	}
}

type ListSecurityRiskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityRiskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityRiskInvoker) Invoke() (*model.ListSecurityRiskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityRiskResponse), nil
	}
}

type ShowAgentStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAgentStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAgentStatisticsInvoker) Invoke() (*model.ShowAgentStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgentStatisticsResponse), nil
	}
}

type ShowHotInformationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHotInformationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHotInformationInvoker) Invoke() (*model.ShowHotInformationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHotInformationResponse), nil
	}
}

type ShowMonthlyOperaReportNotifyInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMonthlyOperaReportNotifyInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMonthlyOperaReportNotifyInfoInvoker) Invoke() (*model.ShowMonthlyOperaReportNotifyInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMonthlyOperaReportNotifyInfoResponse), nil
	}
}

type ShowMonthlyOperationReportDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMonthlyOperationReportDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMonthlyOperationReportDetailInvoker) Invoke() (*model.ShowMonthlyOperationReportDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMonthlyOperationReportDetailResponse), nil
	}
}

type ShowProtectStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProtectStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProtectStatisticsInvoker) Invoke() (*model.ShowProtectStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProtectStatisticsResponse), nil
	}
}

type ShowRiskScoreInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRiskScoreInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRiskScoreInvoker) Invoke() (*model.ShowRiskScoreResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRiskScoreResponse), nil
	}
}

type ShowWelfareAreaInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWelfareAreaInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWelfareAreaInfoInvoker) Invoke() (*model.ShowWelfareAreaInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWelfareAreaInfoResponse), nil
	}
}

type ListPluginInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPluginInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPluginInfoInvoker) Invoke() (*model.ListPluginInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPluginInfoResponse), nil
	}
}

type ListPluginStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPluginStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPluginStatusInvoker) Invoke() (*model.ListPluginStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPluginStatusResponse), nil
	}
}

type ListPluginsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPluginsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPluginsInvoker) Invoke() (*model.ListPluginsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPluginsResponse), nil
	}
}

type StartPluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartPluginInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartPluginInvoker) Invoke() (*model.StartPluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartPluginResponse), nil
	}
}

type StopPluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopPluginInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopPluginInvoker) Invoke() (*model.StopPluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopPluginResponse), nil
	}
}

type AddPolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddPolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddPolicyGroupInvoker) Invoke() (*model.AddPolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddPolicyGroupResponse), nil
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

type ChangePolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangePolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangePolicyGroupInvoker) Invoke() (*model.ChangePolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangePolicyGroupResponse), nil
	}
}

type DeletePolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePolicyGroupInvoker) Invoke() (*model.DeletePolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePolicyGroupResponse), nil
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

type CancelHostsQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelHostsQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelHostsQuotaInvoker) Invoke() (*model.CancelHostsQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelHostsQuotaResponse), nil
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

type ListLockedStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLockedStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLockedStatusInvoker) Invoke() (*model.ListLockedStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLockedStatusResponse), nil
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

type ListResourceIdsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceIdsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceIdsInvoker) Invoke() (*model.ListResourceIdsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceIdsResponse), nil
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

type ShowVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVersionInvoker) Invoke() (*model.ShowVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVersionResponse), nil
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

type AssociateProtectionPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateProtectionPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateProtectionPolicyInvoker) Invoke() (*model.AssociateProtectionPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateProtectionPolicyResponse), nil
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

type DeleteDuplicationInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDuplicationInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDuplicationInfoInvoker) Invoke() (*model.DeleteDuplicationInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDuplicationInfoResponse), nil
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

type ListBackedupByHostIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackedupByHostIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBackedupByHostIdInvoker) Invoke() (*model.ListBackedupByHostIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackedupByHostIdResponse), nil
	}
}

type ListBackupPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackupPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBackupPolicyInvoker) Invoke() (*model.ListBackupPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackupPolicyResponse), nil
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

type ShowBackupInfoByBackupIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBackupInfoByBackupIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBackupInfoByBackupIdInvoker) Invoke() (*model.ShowBackupInfoByBackupIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBackupInfoByBackupIdResponse), nil
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

type CreateImageSynchronizeTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateImageSynchronizeTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateImageSynchronizeTaskInvoker) Invoke() (*model.CreateImageSynchronizeTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateImageSynchronizeTaskResponse), nil
	}
}

type ListRegistryImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRegistryImagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRegistryImagesInvoker) Invoke() (*model.ListRegistryImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRegistryImagesResponse), nil
	}
}

type SendSecurityReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendSecurityReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SendSecurityReportInvoker) Invoke() (*model.SendSecurityReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendSecurityReportResponse), nil
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

type ListSecurityCheckClusterReportsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityCheckClusterReportsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityCheckClusterReportsInvoker) Invoke() (*model.ListSecurityCheckClusterReportsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityCheckClusterReportsResponse), nil
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

type ShowManualSecurityCheckStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowManualSecurityCheckStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowManualSecurityCheckStatusInvoker) Invoke() (*model.ShowManualSecurityCheckStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowManualSecurityCheckStatusResponse), nil
	}
}

type ShowSecurityCheckClusterReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecurityCheckClusterReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSecurityCheckClusterReportInvoker) Invoke() (*model.ShowSecurityCheckClusterReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecurityCheckClusterReportResponse), nil
	}
}

type ShowSecurityCheckConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecurityCheckConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSecurityCheckConfigInvoker) Invoke() (*model.ShowSecurityCheckConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecurityCheckConfigResponse), nil
	}
}

type ShowSecurityCheckHostReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecurityCheckHostReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSecurityCheckHostReportInvoker) Invoke() (*model.ShowSecurityCheckHostReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecurityCheckHostReportResponse), nil
	}
}

type StartManualSecurityCheckInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartManualSecurityCheckInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartManualSecurityCheckInvoker) Invoke() (*model.StartManualSecurityCheckResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartManualSecurityCheckResponse), nil
	}
}

type StopManualSecurityCheckInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopManualSecurityCheckInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopManualSecurityCheckInvoker) Invoke() (*model.StopManualSecurityCheckResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopManualSecurityCheckResponse), nil
	}
}

type UpdateSecurityCheckConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecurityCheckConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSecurityCheckConfigInvoker) Invoke() (*model.UpdateSecurityCheckConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecurityCheckConfigResponse), nil
	}
}

type ChangeAgentAutoUpgradeStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeAgentAutoUpgradeStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeAgentAutoUpgradeStatusInvoker) Invoke() (*model.ChangeAgentAutoUpgradeStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeAgentAutoUpgradeStatusResponse), nil
	}
}

type ChangePolicySwitchStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangePolicySwitchStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangePolicySwitchStatusInvoker) Invoke() (*model.ChangePolicySwitchStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangePolicySwitchStatusResponse), nil
	}
}

type ChangeSwitchesStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeSwitchesStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeSwitchesStatusInvoker) Invoke() (*model.ChangeSwitchesStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeSwitchesStatusResponse), nil
	}
}

type ListAgentAutoUpgradeStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAgentAutoUpgradeStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAgentAutoUpgradeStatusInvoker) Invoke() (*model.ListAgentAutoUpgradeStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAgentAutoUpgradeStatusResponse), nil
	}
}

type ListAgentVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAgentVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAgentVersionInvoker) Invoke() (*model.ListAgentVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAgentVersionResponse), nil
	}
}

type ListAlarmConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlarmConfigInvoker) Invoke() (*model.ListAlarmConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmConfigResponse), nil
	}
}

type ListAutoKillVirusStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAutoKillVirusStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAutoKillVirusStatusInvoker) Invoke() (*model.ListAutoKillVirusStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAutoKillVirusStatusResponse), nil
	}
}

type ListDictionariesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDictionariesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDictionariesInvoker) Invoke() (*model.ListDictionariesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDictionariesResponse), nil
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

type ListSystemConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSystemConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSystemConfigsInvoker) Invoke() (*model.ListSystemConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSystemConfigsResponse), nil
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

type ModifyLoginCommonIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyLoginCommonIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyLoginCommonIpInvoker) Invoke() (*model.ModifyLoginCommonIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyLoginCommonIpResponse), nil
	}
}

type ModifyLoginCommonLocationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyLoginCommonLocationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyLoginCommonLocationInvoker) Invoke() (*model.ModifyLoginCommonLocationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyLoginCommonLocationResponse), nil
	}
}

type ModifyLoginWhiteIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyLoginWhiteIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyLoginWhiteIpInvoker) Invoke() (*model.ModifyLoginWhiteIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyLoginWhiteIpResponse), nil
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

type ShowPolicySwitchStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPolicySwitchStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPolicySwitchStatusInvoker) Invoke() (*model.ShowPolicySwitchStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPolicySwitchStatusResponse), nil
	}
}

type ShowScriptFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScriptFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowScriptFileInvoker) Invoke() (*model.ShowScriptFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScriptFileResponse), nil
	}
}

type ShowSwitchesStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSwitchesStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSwitchesStatusInvoker) Invoke() (*model.ShowSwitchesStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSwitchesStatusResponse), nil
	}
}

type StartAutoKillVirusStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartAutoKillVirusStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartAutoKillVirusStatusInvoker) Invoke() (*model.StartAutoKillVirusStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartAutoKillVirusStatusResponse), nil
	}
}

type UpdateAlarmConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAlarmConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAlarmConfigInvoker) Invoke() (*model.UpdateAlarmConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAlarmConfigResponse), nil
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

type RetryTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RetryTaskInvoker) Invoke() (*model.RetryTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryTaskResponse), nil
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

type ChangeVulWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeVulWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeVulWhiteListInvoker) Invoke() (*model.ChangeVulWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeVulWhiteListResponse), nil
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

type DeleteVulWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVulWhiteListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteVulWhiteListInvoker) Invoke() (*model.DeleteVulWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVulWhiteListResponse), nil
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

type ListGeneralImageVulOperationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGeneralImageVulOperationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGeneralImageVulOperationsInvoker) Invoke() (*model.ListGeneralImageVulOperationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGeneralImageVulOperationsResponse), nil
	}
}

type ListGeneralImageVulsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGeneralImageVulsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGeneralImageVulsInvoker) Invoke() (*model.ListGeneralImageVulsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGeneralImageVulsResponse), nil
	}
}

type ListHandleVulsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHandleVulsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHandleVulsInvoker) Invoke() (*model.ListHandleVulsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHandleVulsResponse), nil
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

type ListUrgentVulnerabilitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUrgentVulnerabilitiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUrgentVulnerabilitiesInvoker) Invoke() (*model.ListUrgentVulnerabilitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUrgentVulnerabilitiesResponse), nil
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

type ListVulHostAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulHostAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulHostAppsInvoker) Invoke() (*model.ListVulHostAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulHostAppsResponse), nil
	}
}

type ListVulHostBackupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulHostBackupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulHostBackupsInvoker) Invoke() (*model.ListVulHostBackupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulHostBackupsResponse), nil
	}
}

type ListVulHostProcessInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulHostProcessInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulHostProcessInvoker) Invoke() (*model.ListVulHostProcessResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulHostProcessResponse), nil
	}
}

type ListVulHostVaultsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulHostVaultsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulHostVaultsInvoker) Invoke() (*model.ListVulHostVaultsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulHostVaultsResponse), nil
	}
}

type ListVulHostVulsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulHostVulsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulHostVulsInvoker) Invoke() (*model.ListVulHostVulsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulHostVulsResponse), nil
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

type ListVulRepairCmdsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulRepairCmdsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulRepairCmdsInvoker) Invoke() (*model.ListVulRepairCmdsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulRepairCmdsResponse), nil
	}
}

type ListVulRepairFailedDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulRepairFailedDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulRepairFailedDetailInvoker) Invoke() (*model.ListVulRepairFailedDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulRepairFailedDetailResponse), nil
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

type RecreateVulScanTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecreateVulScanTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecreateVulScanTaskInvoker) Invoke() (*model.RecreateVulScanTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecreateVulScanTaskResponse), nil
	}
}

type RestoreVulHostBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreVulHostBackupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestoreVulHostBackupInvoker) Invoke() (*model.RestoreVulHostBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreVulHostBackupResponse), nil
	}
}

type ShowVulBackupStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVulBackupStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVulBackupStatisticsInvoker) Invoke() (*model.ShowVulBackupStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVulBackupStatisticsResponse), nil
	}
}

type ShowVulReportDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVulReportDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVulReportDataInvoker) Invoke() (*model.ShowVulReportDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVulReportDataResponse), nil
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

type ShowVulScanTaskEstimatedTimeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVulScanTaskEstimatedTimeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVulScanTaskEstimatedTimeInvoker) Invoke() (*model.ShowVulScanTaskEstimatedTimeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVulScanTaskEstimatedTimeResponse), nil
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

type ShowWindosVulDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWindosVulDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWindosVulDetailInvoker) Invoke() (*model.ShowWindosVulDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWindosVulDetailResponse), nil
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

type DeleteBackupHostInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBackupHostInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBackupHostInfoInvoker) Invoke() (*model.DeleteBackupHostInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBackupHostInfoResponse), nil
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

type ListBackupHostsInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackupHostsInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBackupHostsInfoInvoker) Invoke() (*model.ListBackupHostsInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackupHostsInfoResponse), nil
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

type SetProtectDirSwitchInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetProtectDirSwitchInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetProtectDirSwitchInfoInvoker) Invoke() (*model.SetProtectDirSwitchInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetProtectDirSwitchInfoResponse), nil
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

type SetRemoteBackupInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetRemoteBackupInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetRemoteBackupInfoInvoker) Invoke() (*model.SetRemoteBackupInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetRemoteBackupInfoResponse), nil
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

type ShowRemoteBackupHostInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRemoteBackupHostInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRemoteBackupHostInfoInvoker) Invoke() (*model.ShowRemoteBackupHostInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRemoteBackupHostInfoResponse), nil
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

type ShowWtpProtectStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWtpProtectStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWtpProtectStatisticsInvoker) Invoke() (*model.ShowWtpProtectStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWtpProtectStatisticsResponse), nil
	}
}

type UpdateBackupHostInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBackupHostInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateBackupHostInfoInvoker) Invoke() (*model.UpdateBackupHostInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBackupHostInfoResponse), nil
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
