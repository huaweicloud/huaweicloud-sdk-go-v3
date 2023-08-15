package v5

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/hss/v5/model"
)

type AddHostsGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddHostsGroupInvoker) Invoke() (*model.AddHostsGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddHostsGroupResponse), nil
	}
}

type AssociatePolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociatePolicyGroupInvoker) Invoke() (*model.AssociatePolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociatePolicyGroupResponse), nil
	}
}

type BatchCreateTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateTagsInvoker) Invoke() (*model.BatchCreateTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateTagsResponse), nil
	}
}

type ChangeEventInvoker struct {
	*invoker.BaseInvoker
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

func (i *ChangeHostsGroupInvoker) Invoke() (*model.ChangeHostsGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeHostsGroupResponse), nil
	}
}

type ChangeVulStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeVulStatusInvoker) Invoke() (*model.ChangeVulStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeVulStatusResponse), nil
	}
}

type DeleteHostsGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHostsGroupInvoker) Invoke() (*model.DeleteHostsGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHostsGroupResponse), nil
	}
}

type DeleteResourceInstanceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResourceInstanceTagInvoker) Invoke() (*model.DeleteResourceInstanceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResourceInstanceTagResponse), nil
	}
}

type ListAlarmWhiteListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmWhiteListInvoker) Invoke() (*model.ListAlarmWhiteListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmWhiteListResponse), nil
	}
}

type ListAppChangeHistoriesInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListAutoLaunchsInvoker) Invoke() (*model.ListAutoLaunchsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAutoLaunchsResponse), nil
	}
}

type ListHostGroupsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListHostVulsInvoker) Invoke() (*model.ListHostVulsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostVulsResponse), nil
	}
}

type ListJarPackageHostInfoInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListJarPackageStatisticsInvoker) Invoke() (*model.ListJarPackageStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJarPackageStatisticsResponse), nil
	}
}

type ListPasswordComplexityInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPasswordComplexityInvoker) Invoke() (*model.ListPasswordComplexityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPasswordComplexityResponse), nil
	}
}

type ListPolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyGroupInvoker) Invoke() (*model.ListPolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyGroupResponse), nil
	}
}

type ListPortStatisticsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListProcessStatisticsInvoker) Invoke() (*model.ListProcessStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProcessStatisticsResponse), nil
	}
}

type ListProtectionPolicyInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListProtectionServerInvoker) Invoke() (*model.ListProtectionServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProtectionServerResponse), nil
	}
}

type ListQuotasDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotasDetailInvoker) Invoke() (*model.ListQuotasDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotasDetailResponse), nil
	}
}

type ListRiskConfigCheckRulesInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListSecurityEventsInvoker) Invoke() (*model.ListSecurityEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityEventsResponse), nil
	}
}

type ListUserChangeHistoriesInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListUsersInvoker) Invoke() (*model.ListUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsersResponse), nil
	}
}

type ListVulHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulHostsInvoker) Invoke() (*model.ListVulHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulHostsResponse), nil
	}
}

type ListVulnerabilitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulnerabilitiesInvoker) Invoke() (*model.ListVulnerabilitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulnerabilitiesResponse), nil
	}
}

type ListWeakPasswordUsersInvoker struct {
	*invoker.BaseInvoker
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

func (i *SetWtpProtectionStatusInfoInvoker) Invoke() (*model.SetWtpProtectionStatusInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetWtpProtectionStatusInfoResponse), nil
	}
}

type ShowAssetStatisticInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowCheckRuleDetailInvoker) Invoke() (*model.ShowCheckRuleDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCheckRuleDetailResponse), nil
	}
}

type ShowResourceQuotasInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowRiskConfigDetailInvoker) Invoke() (*model.ShowRiskConfigDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRiskConfigDetailResponse), nil
	}
}

type StartProtectionInvoker struct {
	*invoker.BaseInvoker
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

func (i *StopProtectionInvoker) Invoke() (*model.StopProtectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopProtectionResponse), nil
	}
}

type SwitchHostsProtectStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchHostsProtectStatusInvoker) Invoke() (*model.SwitchHostsProtectStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchHostsProtectStatusResponse), nil
	}
}

type UpdateBackupPolicyInfoInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateProtectionPolicyInvoker) Invoke() (*model.UpdateProtectionPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProtectionPolicyResponse), nil
	}
}
