package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/waf/v1/model"
)

type ApplyCertificateToHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplyCertificateToHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ApplyCertificateToHostInvoker) Invoke() (*model.ApplyCertificateToHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplyCertificateToHostResponse), nil
	}
}

type ChangePrepaidCloudWafInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangePrepaidCloudWafInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangePrepaidCloudWafInvoker) Invoke() (*model.ChangePrepaidCloudWafResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangePrepaidCloudWafResponse), nil
	}
}

type CreateAntiTamperRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAntiTamperRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAntiTamperRuleInvoker) Invoke() (*model.CreateAntiTamperRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAntiTamperRuleResponse), nil
	}
}

type CreateAnticrawlerRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAnticrawlerRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAnticrawlerRuleInvoker) Invoke() (*model.CreateAnticrawlerRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAnticrawlerRuleResponse), nil
	}
}

type CreateAntileakageRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAntileakageRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAntileakageRuleInvoker) Invoke() (*model.CreateAntileakageRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAntileakageRuleResponse), nil
	}
}

type CreateCcRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCcRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCcRuleInvoker) Invoke() (*model.CreateCcRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCcRuleResponse), nil
	}
}

type CreateCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCertificateInvoker) Invoke() (*model.CreateCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCertificateResponse), nil
	}
}

type CreateCloudWafPostPaidResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCloudWafPostPaidResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCloudWafPostPaidResourceInvoker) Invoke() (*model.CreateCloudWafPostPaidResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCloudWafPostPaidResourceResponse), nil
	}
}

type CreateCustomRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCustomRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCustomRuleInvoker) Invoke() (*model.CreateCustomRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCustomRuleResponse), nil
	}
}

type CreateGeoipRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGeoipRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGeoipRuleInvoker) Invoke() (*model.CreateGeoipRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGeoipRuleResponse), nil
	}
}

type CreateHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHostInvoker) Invoke() (*model.CreateHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHostResponse), nil
	}
}

type CreateIgnoreRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateIgnoreRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateIgnoreRuleInvoker) Invoke() (*model.CreateIgnoreRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIgnoreRuleResponse), nil
	}
}

type CreateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceInvoker) Invoke() (*model.CreateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceResponse), nil
	}
}

type CreateIpGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateIpGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateIpGroupInvoker) Invoke() (*model.CreateIpGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIpGroupResponse), nil
	}
}

type CreatePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePolicyInvoker) Invoke() (*model.CreatePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePolicyResponse), nil
	}
}

type CreatePremiumHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePremiumHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePremiumHostInvoker) Invoke() (*model.CreatePremiumHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePremiumHostResponse), nil
	}
}

type CreatePrepaidCloudWafInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePrepaidCloudWafInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePrepaidCloudWafInvoker) Invoke() (*model.CreatePrepaidCloudWafResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePrepaidCloudWafResponse), nil
	}
}

type CreatePrivacyRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePrivacyRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePrivacyRuleInvoker) Invoke() (*model.CreatePrivacyRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePrivacyRuleResponse), nil
	}
}

type CreatePunishmentRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePunishmentRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePunishmentRuleInvoker) Invoke() (*model.CreatePunishmentRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePunishmentRuleResponse), nil
	}
}

type CreateValueListInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateValueListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateValueListInvoker) Invoke() (*model.CreateValueListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateValueListResponse), nil
	}
}

type CreateWhiteblackipRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWhiteblackipRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWhiteblackipRuleInvoker) Invoke() (*model.CreateWhiteblackipRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWhiteblackipRuleResponse), nil
	}
}

type DeleteAnticrawlerRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAnticrawlerRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAnticrawlerRuleInvoker) Invoke() (*model.DeleteAnticrawlerRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAnticrawlerRuleResponse), nil
	}
}

type DeleteAntileakageRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAntileakageRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAntileakageRuleInvoker) Invoke() (*model.DeleteAntileakageRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAntileakageRuleResponse), nil
	}
}

type DeleteAntitamperRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAntitamperRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAntitamperRuleInvoker) Invoke() (*model.DeleteAntitamperRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAntitamperRuleResponse), nil
	}
}

type DeleteCcRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCcRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCcRuleInvoker) Invoke() (*model.DeleteCcRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCcRuleResponse), nil
	}
}

type DeleteCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCertificateInvoker) Invoke() (*model.DeleteCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCertificateResponse), nil
	}
}

type DeleteCloudWafPostPaidResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCloudWafPostPaidResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCloudWafPostPaidResourceInvoker) Invoke() (*model.DeleteCloudWafPostPaidResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCloudWafPostPaidResourceResponse), nil
	}
}

type DeleteCustomRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCustomRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCustomRuleInvoker) Invoke() (*model.DeleteCustomRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCustomRuleResponse), nil
	}
}

type DeleteGeoipRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGeoipRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGeoipRuleInvoker) Invoke() (*model.DeleteGeoipRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGeoipRuleResponse), nil
	}
}

type DeleteHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHostInvoker) Invoke() (*model.DeleteHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHostResponse), nil
	}
}

type DeleteIgnoreRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteIgnoreRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteIgnoreRuleInvoker) Invoke() (*model.DeleteIgnoreRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteIgnoreRuleResponse), nil
	}
}

type DeleteInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceInvoker) Invoke() (*model.DeleteInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceResponse), nil
	}
}

type DeleteIpGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteIpGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteIpGroupInvoker) Invoke() (*model.DeleteIpGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteIpGroupResponse), nil
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

type DeletePremiumHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePremiumHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePremiumHostInvoker) Invoke() (*model.DeletePremiumHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePremiumHostResponse), nil
	}
}

type DeletePrivacyRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePrivacyRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePrivacyRuleInvoker) Invoke() (*model.DeletePrivacyRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePrivacyRuleResponse), nil
	}
}

type DeletePunishmentRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePunishmentRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePunishmentRuleInvoker) Invoke() (*model.DeletePunishmentRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePunishmentRuleResponse), nil
	}
}

type DeleteValueListInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteValueListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteValueListInvoker) Invoke() (*model.DeleteValueListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteValueListResponse), nil
	}
}

type DeleteWhiteBlackIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWhiteBlackIpRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteWhiteBlackIpRuleInvoker) Invoke() (*model.DeleteWhiteBlackIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWhiteBlackIpRuleResponse), nil
	}
}

type ListAnticrawlerRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAnticrawlerRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAnticrawlerRulesInvoker) Invoke() (*model.ListAnticrawlerRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAnticrawlerRulesResponse), nil
	}
}

type ListAntileakageRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAntileakageRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAntileakageRulesInvoker) Invoke() (*model.ListAntileakageRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAntileakageRulesResponse), nil
	}
}

type ListAntitamperRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAntitamperRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAntitamperRuleInvoker) Invoke() (*model.ListAntitamperRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAntitamperRuleResponse), nil
	}
}

type ListBandwidthTimelineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBandwidthTimelineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBandwidthTimelineInvoker) Invoke() (*model.ListBandwidthTimelineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBandwidthTimelineResponse), nil
	}
}

type ListCcRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCcRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCcRulesInvoker) Invoke() (*model.ListCcRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCcRulesResponse), nil
	}
}

type ListCertificatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCertificatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCertificatesInvoker) Invoke() (*model.ListCertificatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCertificatesResponse), nil
	}
}

type ListCompositeHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCompositeHostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCompositeHostsInvoker) Invoke() (*model.ListCompositeHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCompositeHostsResponse), nil
	}
}

type ListCustomRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCustomRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCustomRulesInvoker) Invoke() (*model.ListCustomRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCustomRulesResponse), nil
	}
}

type ListEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEventInvoker) Invoke() (*model.ListEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventResponse), nil
	}
}

type ListGeoipRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGeoipRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGeoipRuleInvoker) Invoke() (*model.ListGeoipRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGeoipRuleResponse), nil
	}
}

type ListHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHostInvoker) Invoke() (*model.ListHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostResponse), nil
	}
}

type ListHostRouteInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostRouteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHostRouteInvoker) Invoke() (*model.ListHostRouteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostRouteResponse), nil
	}
}

type ListIgnoreRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIgnoreRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIgnoreRuleInvoker) Invoke() (*model.ListIgnoreRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIgnoreRuleResponse), nil
	}
}

type ListInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceInvoker) Invoke() (*model.ListInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceResponse), nil
	}
}

type ListIpGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIpGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIpGroupInvoker) Invoke() (*model.ListIpGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIpGroupResponse), nil
	}
}

type ListNoticeConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNoticeConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNoticeConfigsInvoker) Invoke() (*model.ListNoticeConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNoticeConfigsResponse), nil
	}
}

type ListOverviewsClassificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOverviewsClassificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOverviewsClassificationInvoker) Invoke() (*model.ListOverviewsClassificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOverviewsClassificationResponse), nil
	}
}

type ListPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyInvoker) Invoke() (*model.ListPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyResponse), nil
	}
}

type ListPremiumHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPremiumHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPremiumHostInvoker) Invoke() (*model.ListPremiumHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPremiumHostResponse), nil
	}
}

type ListPrivacyRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPrivacyRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPrivacyRuleInvoker) Invoke() (*model.ListPrivacyRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPrivacyRuleResponse), nil
	}
}

type ListPunishmentRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPunishmentRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPunishmentRulesInvoker) Invoke() (*model.ListPunishmentRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPunishmentRulesResponse), nil
	}
}

type ListQpsTimelineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQpsTimelineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQpsTimelineInvoker) Invoke() (*model.ListQpsTimelineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQpsTimelineResponse), nil
	}
}

type ListRequestTimelineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRequestTimelineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRequestTimelineInvoker) Invoke() (*model.ListRequestTimelineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRequestTimelineResponse), nil
	}
}

type ListStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStatisticsInvoker) Invoke() (*model.ListStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStatisticsResponse), nil
	}
}

type ListTopAbnormalInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTopAbnormalInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTopAbnormalInvoker) Invoke() (*model.ListTopAbnormalResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTopAbnormalResponse), nil
	}
}

type ListValueListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListValueListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListValueListInvoker) Invoke() (*model.ListValueListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListValueListResponse), nil
	}
}

type ListWhiteblackipRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWhiteblackipRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWhiteblackipRuleInvoker) Invoke() (*model.ListWhiteblackipRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWhiteblackipRuleResponse), nil
	}
}

type MigrateCompositeHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *MigrateCompositeHostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *MigrateCompositeHostsInvoker) Invoke() (*model.MigrateCompositeHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.MigrateCompositeHostsResponse), nil
	}
}

type RenameInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RenameInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RenameInstanceInvoker) Invoke() (*model.RenameInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RenameInstanceResponse), nil
	}
}

type ShowAnticrawlerRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAnticrawlerRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAnticrawlerRuleInvoker) Invoke() (*model.ShowAnticrawlerRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAnticrawlerRuleResponse), nil
	}
}

type ShowAntileakageRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAntileakageRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAntileakageRuleInvoker) Invoke() (*model.ShowAntileakageRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAntileakageRuleResponse), nil
	}
}

type ShowAntitamperRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAntitamperRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAntitamperRuleInvoker) Invoke() (*model.ShowAntitamperRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAntitamperRuleResponse), nil
	}
}

type ShowCcRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCcRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCcRuleInvoker) Invoke() (*model.ShowCcRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCcRuleResponse), nil
	}
}

type ShowCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCertificateInvoker) Invoke() (*model.ShowCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCertificateResponse), nil
	}
}

type ShowCompositeHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCompositeHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCompositeHostInvoker) Invoke() (*model.ShowCompositeHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCompositeHostResponse), nil
	}
}

type ShowConsoleConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConsoleConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConsoleConfigInvoker) Invoke() (*model.ShowConsoleConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConsoleConfigResponse), nil
	}
}

type ShowCustomRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCustomRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCustomRuleInvoker) Invoke() (*model.ShowCustomRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCustomRuleResponse), nil
	}
}

type ShowEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEventInvoker) Invoke() (*model.ShowEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEventResponse), nil
	}
}

type ShowGeoipRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGeoipRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGeoipRuleInvoker) Invoke() (*model.ShowGeoipRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGeoipRuleResponse), nil
	}
}

type ShowHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHostInvoker) Invoke() (*model.ShowHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHostResponse), nil
	}
}

type ShowIgnoreRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIgnoreRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIgnoreRuleInvoker) Invoke() (*model.ShowIgnoreRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIgnoreRuleResponse), nil
	}
}

type ShowInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceInvoker) Invoke() (*model.ShowInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceResponse), nil
	}
}

type ShowIpGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIpGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIpGroupInvoker) Invoke() (*model.ShowIpGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIpGroupResponse), nil
	}
}

type ShowLtsInfoConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLtsInfoConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLtsInfoConfigInvoker) Invoke() (*model.ShowLtsInfoConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLtsInfoConfigResponse), nil
	}
}

type ShowPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPolicyInvoker) Invoke() (*model.ShowPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPolicyResponse), nil
	}
}

type ShowPremiumHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPremiumHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPremiumHostInvoker) Invoke() (*model.ShowPremiumHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPremiumHostResponse), nil
	}
}

type ShowPrivacyRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPrivacyRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPrivacyRuleInvoker) Invoke() (*model.ShowPrivacyRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPrivacyRuleResponse), nil
	}
}

type ShowPunishmentRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPunishmentRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPunishmentRuleInvoker) Invoke() (*model.ShowPunishmentRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPunishmentRuleResponse), nil
	}
}

type ShowSourceIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSourceIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSourceIpInvoker) Invoke() (*model.ShowSourceIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSourceIpResponse), nil
	}
}

type ShowSubscriptionInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSubscriptionInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSubscriptionInfoInvoker) Invoke() (*model.ShowSubscriptionInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSubscriptionInfoResponse), nil
	}
}

type ShowValueListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowValueListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowValueListInvoker) Invoke() (*model.ShowValueListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowValueListResponse), nil
	}
}

type ShowWhiteBlackIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWhiteBlackIpRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWhiteBlackIpRuleInvoker) Invoke() (*model.ShowWhiteBlackIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWhiteBlackIpRuleResponse), nil
	}
}

type UpdateAlertNoticeConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAlertNoticeConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAlertNoticeConfigInvoker) Invoke() (*model.UpdateAlertNoticeConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAlertNoticeConfigResponse), nil
	}
}

type UpdateAntiTamperRuleRefreshInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAntiTamperRuleRefreshInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAntiTamperRuleRefreshInvoker) Invoke() (*model.UpdateAntiTamperRuleRefreshResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAntiTamperRuleRefreshResponse), nil
	}
}

type UpdateAnticrawlerRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAnticrawlerRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAnticrawlerRuleInvoker) Invoke() (*model.UpdateAnticrawlerRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAnticrawlerRuleResponse), nil
	}
}

type UpdateAnticrawlerRuleTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAnticrawlerRuleTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAnticrawlerRuleTypeInvoker) Invoke() (*model.UpdateAnticrawlerRuleTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAnticrawlerRuleTypeResponse), nil
	}
}

type UpdateAntileakageRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAntileakageRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAntileakageRuleInvoker) Invoke() (*model.UpdateAntileakageRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAntileakageRuleResponse), nil
	}
}

type UpdateCcRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCcRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCcRuleInvoker) Invoke() (*model.UpdateCcRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCcRuleResponse), nil
	}
}

type UpdateCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCertificateInvoker) Invoke() (*model.UpdateCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCertificateResponse), nil
	}
}

type UpdateCustomRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCustomRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCustomRuleInvoker) Invoke() (*model.UpdateCustomRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCustomRuleResponse), nil
	}
}

type UpdateGeoipRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGeoipRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateGeoipRuleInvoker) Invoke() (*model.UpdateGeoipRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGeoipRuleResponse), nil
	}
}

type UpdateHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHostInvoker) Invoke() (*model.UpdateHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHostResponse), nil
	}
}

type UpdateHostProtectStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHostProtectStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHostProtectStatusInvoker) Invoke() (*model.UpdateHostProtectStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHostProtectStatusResponse), nil
	}
}

type UpdateIgnoreRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateIgnoreRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateIgnoreRuleInvoker) Invoke() (*model.UpdateIgnoreRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIgnoreRuleResponse), nil
	}
}

type UpdateIpGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateIpGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateIpGroupInvoker) Invoke() (*model.UpdateIpGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIpGroupResponse), nil
	}
}

type UpdateLtsInfoConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLtsInfoConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLtsInfoConfigInvoker) Invoke() (*model.UpdateLtsInfoConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLtsInfoConfigResponse), nil
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

type UpdatePolicyProtectHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePolicyProtectHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePolicyProtectHostInvoker) Invoke() (*model.UpdatePolicyProtectHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePolicyProtectHostResponse), nil
	}
}

type UpdatePolicyRuleStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePolicyRuleStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePolicyRuleStatusInvoker) Invoke() (*model.UpdatePolicyRuleStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePolicyRuleStatusResponse), nil
	}
}

type UpdatePremiumHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePremiumHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePremiumHostInvoker) Invoke() (*model.UpdatePremiumHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePremiumHostResponse), nil
	}
}

type UpdatePremiumHostProtectStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePremiumHostProtectStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePremiumHostProtectStatusInvoker) Invoke() (*model.UpdatePremiumHostProtectStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePremiumHostProtectStatusResponse), nil
	}
}

type UpdatePrivacyRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePrivacyRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePrivacyRuleInvoker) Invoke() (*model.UpdatePrivacyRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePrivacyRuleResponse), nil
	}
}

type UpdatePunishmentRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePunishmentRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePunishmentRuleInvoker) Invoke() (*model.UpdatePunishmentRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePunishmentRuleResponse), nil
	}
}

type UpdateValueListInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateValueListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateValueListInvoker) Invoke() (*model.UpdateValueListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateValueListResponse), nil
	}
}

type UpdateWhiteblackipRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWhiteblackipRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWhiteblackipRuleInvoker) Invoke() (*model.UpdateWhiteblackipRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWhiteblackipRuleResponse), nil
	}
}
