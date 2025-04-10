package v5

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/hss/v5/model"
)

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
