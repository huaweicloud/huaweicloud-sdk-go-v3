package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/edgesec/v2/model"
)

type CreateDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDomainInvoker) Invoke() (*model.CreateDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDomainResponse), nil
	}
}

type CreateHttpReferenceTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHttpReferenceTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHttpReferenceTableInvoker) Invoke() (*model.CreateHttpReferenceTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHttpReferenceTableResponse), nil
	}
}

type DeleteDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDomainInvoker) Invoke() (*model.DeleteDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDomainResponse), nil
	}
}

type DeleteHttpReferenceTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHttpReferenceTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHttpReferenceTableInvoker) Invoke() (*model.DeleteHttpReferenceTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHttpReferenceTableResponse), nil
	}
}

type ShowDomainDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDomainDetailInvoker) Invoke() (*model.ShowDomainDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainDetailResponse), nil
	}
}

type ShowDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDomainsInvoker) Invoke() (*model.ShowDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainsResponse), nil
	}
}

type ShowHttpOverviewsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpOverviewsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpOverviewsInvoker) Invoke() (*model.ShowHttpOverviewsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpOverviewsResponse), nil
	}
}

type ShowHttpReferenceTablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpReferenceTablesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpReferenceTablesInvoker) Invoke() (*model.ShowHttpReferenceTablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpReferenceTablesResponse), nil
	}
}

type ShowHttpStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpStatisticsInvoker) Invoke() (*model.ShowHttpStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpStatisticsResponse), nil
	}
}

type UpdateDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDomainInvoker) Invoke() (*model.UpdateDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainResponse), nil
	}
}

type UpdateHttpReferenceTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHttpReferenceTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHttpReferenceTableInvoker) Invoke() (*model.UpdateHttpReferenceTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHttpReferenceTableResponse), nil
	}
}

type DownloadDdosAttackLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadDdosAttackLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadDdosAttackLogsInvoker) Invoke() (*model.DownloadDdosAttackLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadDdosAttackLogsResponse), nil
	}
}

type ShowDdosAttackLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDdosAttackLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDdosAttackLogsInvoker) Invoke() (*model.ShowDdosAttackLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDdosAttackLogsResponse), nil
	}
}

type ShowDdosAttackTimelineStatsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDdosAttackTimelineStatsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDdosAttackTimelineStatsInvoker) Invoke() (*model.ShowDdosAttackTimelineStatsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDdosAttackTimelineStatsResponse), nil
	}
}

type ApplyHttpPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplyHttpPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ApplyHttpPolicyInvoker) Invoke() (*model.ApplyHttpPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplyHttpPolicyResponse), nil
	}
}

type CreateHttpAccessControlRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHttpAccessControlRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHttpAccessControlRuleInvoker) Invoke() (*model.CreateHttpAccessControlRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHttpAccessControlRuleResponse), nil
	}
}

type CreateHttpGeoIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHttpGeoIpRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHttpGeoIpRuleInvoker) Invoke() (*model.CreateHttpGeoIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHttpGeoIpRuleResponse), nil
	}
}

type CreateHttpIgnoreRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHttpIgnoreRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHttpIgnoreRuleInvoker) Invoke() (*model.CreateHttpIgnoreRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHttpIgnoreRuleResponse), nil
	}
}

type CreateHttpPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHttpPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHttpPolicyInvoker) Invoke() (*model.CreateHttpPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHttpPolicyResponse), nil
	}
}

type CreateHttpPunishmentRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHttpPunishmentRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHttpPunishmentRuleInvoker) Invoke() (*model.CreateHttpPunishmentRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHttpPunishmentRuleResponse), nil
	}
}

type DeleteHttpAccessControlRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHttpAccessControlRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHttpAccessControlRuleInvoker) Invoke() (*model.DeleteHttpAccessControlRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHttpAccessControlRuleResponse), nil
	}
}

type DeleteHttpGeoIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHttpGeoIpRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHttpGeoIpRuleInvoker) Invoke() (*model.DeleteHttpGeoIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHttpGeoIpRuleResponse), nil
	}
}

type DeleteHttpIgnoreRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHttpIgnoreRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHttpIgnoreRuleInvoker) Invoke() (*model.DeleteHttpIgnoreRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHttpIgnoreRuleResponse), nil
	}
}

type DeleteHttpPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHttpPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHttpPolicyInvoker) Invoke() (*model.DeleteHttpPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHttpPolicyResponse), nil
	}
}

type DeleteHttpPunishmentRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHttpPunishmentRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHttpPunishmentRuleInvoker) Invoke() (*model.DeleteHttpPunishmentRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHttpPunishmentRuleResponse), nil
	}
}

type ResetHttpIgnoreRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetHttpIgnoreRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetHttpIgnoreRuleInvoker) Invoke() (*model.ResetHttpIgnoreRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetHttpIgnoreRuleResponse), nil
	}
}

type ShowHttpAccessControlRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpAccessControlRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpAccessControlRuleInvoker) Invoke() (*model.ShowHttpAccessControlRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpAccessControlRuleResponse), nil
	}
}

type ShowHttpAccessControlRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpAccessControlRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpAccessControlRulesInvoker) Invoke() (*model.ShowHttpAccessControlRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpAccessControlRulesResponse), nil
	}
}

type ShowHttpAttackDistributionStatsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpAttackDistributionStatsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpAttackDistributionStatsInvoker) Invoke() (*model.ShowHttpAttackDistributionStatsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpAttackDistributionStatsResponse), nil
	}
}

type ShowHttpAttackTimelineStatsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpAttackTimelineStatsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpAttackTimelineStatsInvoker) Invoke() (*model.ShowHttpAttackTimelineStatsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpAttackTimelineStatsResponse), nil
	}
}

type ShowHttpAttackTopStatsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpAttackTopStatsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpAttackTopStatsInvoker) Invoke() (*model.ShowHttpAttackTopStatsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpAttackTopStatsResponse), nil
	}
}

type ShowHttpGeoIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpGeoIpRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpGeoIpRuleInvoker) Invoke() (*model.ShowHttpGeoIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpGeoIpRuleResponse), nil
	}
}

type ShowHttpGeoIpRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpGeoIpRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpGeoIpRulesInvoker) Invoke() (*model.ShowHttpGeoIpRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpGeoIpRulesResponse), nil
	}
}

type ShowHttpIgnoreRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpIgnoreRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpIgnoreRuleInvoker) Invoke() (*model.ShowHttpIgnoreRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpIgnoreRuleResponse), nil
	}
}

type ShowHttpIgnoreRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpIgnoreRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpIgnoreRulesInvoker) Invoke() (*model.ShowHttpIgnoreRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpIgnoreRulesResponse), nil
	}
}

type ShowHttpPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpPoliciesInvoker) Invoke() (*model.ShowHttpPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpPoliciesResponse), nil
	}
}

type ShowHttpPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpPolicyInvoker) Invoke() (*model.ShowHttpPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpPolicyResponse), nil
	}
}

type ShowHttpPunishmentRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpPunishmentRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpPunishmentRuleInvoker) Invoke() (*model.ShowHttpPunishmentRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpPunishmentRuleResponse), nil
	}
}

type ShowHttpPunishmentRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpPunishmentRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpPunishmentRulesInvoker) Invoke() (*model.ShowHttpPunishmentRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpPunishmentRulesResponse), nil
	}
}

type UpdateHttpAccessControlRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHttpAccessControlRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHttpAccessControlRuleInvoker) Invoke() (*model.UpdateHttpAccessControlRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHttpAccessControlRuleResponse), nil
	}
}

type UpdateHttpGeoIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHttpGeoIpRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHttpGeoIpRuleInvoker) Invoke() (*model.UpdateHttpGeoIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHttpGeoIpRuleResponse), nil
	}
}

type UpdateHttpIgnoreRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHttpIgnoreRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHttpIgnoreRuleInvoker) Invoke() (*model.UpdateHttpIgnoreRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHttpIgnoreRuleResponse), nil
	}
}

type UpdateHttpPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHttpPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHttpPolicyInvoker) Invoke() (*model.UpdateHttpPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHttpPolicyResponse), nil
	}
}

type UpdateHttpPolicyRuleStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHttpPolicyRuleStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHttpPolicyRuleStatusInvoker) Invoke() (*model.UpdateHttpPolicyRuleStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHttpPolicyRuleStatusResponse), nil
	}
}

type UpdateHttpPunishmentRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHttpPunishmentRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHttpPunishmentRuleInvoker) Invoke() (*model.UpdateHttpPunishmentRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHttpPunishmentRuleResponse), nil
	}
}

type CreateHttpCcRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHttpCcRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHttpCcRuleInvoker) Invoke() (*model.CreateHttpCcRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHttpCcRuleResponse), nil
	}
}

type DeleteHttpCcRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHttpCcRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHttpCcRuleInvoker) Invoke() (*model.DeleteHttpCcRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHttpCcRuleResponse), nil
	}
}

type ShowHttpCcRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpCcRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpCcRuleInvoker) Invoke() (*model.ShowHttpCcRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpCcRuleResponse), nil
	}
}

type ShowHttpCcRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpCcRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpCcRulesInvoker) Invoke() (*model.ShowHttpCcRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpCcRulesResponse), nil
	}
}

type UpdateHttpCcRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHttpCcRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHttpCcRuleInvoker) Invoke() (*model.UpdateHttpCcRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHttpCcRuleResponse), nil
	}
}

type CreateHttpBlockTrustIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHttpBlockTrustIpRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHttpBlockTrustIpRuleInvoker) Invoke() (*model.CreateHttpBlockTrustIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHttpBlockTrustIpRuleResponse), nil
	}
}

type DeleteHttpBlockTrustIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHttpBlockTrustIpRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHttpBlockTrustIpRuleInvoker) Invoke() (*model.DeleteHttpBlockTrustIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHttpBlockTrustIpRuleResponse), nil
	}
}

type ShowHttpBlockTrustIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpBlockTrustIpRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpBlockTrustIpRuleInvoker) Invoke() (*model.ShowHttpBlockTrustIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpBlockTrustIpRuleResponse), nil
	}
}

type ShowHttpBlockTrustIpRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpBlockTrustIpRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpBlockTrustIpRulesInvoker) Invoke() (*model.ShowHttpBlockTrustIpRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpBlockTrustIpRulesResponse), nil
	}
}

type UpdateHttpBlockTrustIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHttpBlockTrustIpRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHttpBlockTrustIpRuleInvoker) Invoke() (*model.UpdateHttpBlockTrustIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHttpBlockTrustIpRuleResponse), nil
	}
}

type CreateHttpIpGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHttpIpGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHttpIpGroupInvoker) Invoke() (*model.CreateHttpIpGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHttpIpGroupResponse), nil
	}
}

type DeleteHttpIpGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHttpIpGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHttpIpGroupInvoker) Invoke() (*model.DeleteHttpIpGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHttpIpGroupResponse), nil
	}
}

type ShowHttpIpGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpIpGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpIpGroupInvoker) Invoke() (*model.ShowHttpIpGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpIpGroupResponse), nil
	}
}

type ShowHttpIpGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpIpGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpIpGroupsInvoker) Invoke() (*model.ShowHttpIpGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpIpGroupsResponse), nil
	}
}

type UpdateHttpIpGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHttpIpGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHttpIpGroupInvoker) Invoke() (*model.UpdateHttpIpGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHttpIpGroupResponse), nil
	}
}
