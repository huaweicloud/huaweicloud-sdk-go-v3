package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/waf/v1/model"
)

type ApplyCertificateToHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplyCertificateToHostInvoker) Invoke() (*model.ApplyCertificateToHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplyCertificateToHostResponse), nil
	}
}

type CreateAntiTamperRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAntiTamperRuleInvoker) Invoke() (*model.CreateAntiTamperRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAntiTamperRuleResponse), nil
	}
}

type CreateCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCertificateInvoker) Invoke() (*model.CreateCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCertificateResponse), nil
	}
}

type CreateGeoipRuleInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreatePremiumHostInvoker) Invoke() (*model.CreatePremiumHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePremiumHostResponse), nil
	}
}

type CreatePrivacyRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePrivacyRuleInvoker) Invoke() (*model.CreatePrivacyRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePrivacyRuleResponse), nil
	}
}

type CreateValueListInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreateWhiteblackipRuleInvoker) Invoke() (*model.CreateWhiteblackipRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWhiteblackipRuleResponse), nil
	}
}

type DeleteAntitamperRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAntitamperRuleInvoker) Invoke() (*model.DeleteAntitamperRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAntitamperRuleResponse), nil
	}
}

type DeleteCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCertificateInvoker) Invoke() (*model.DeleteCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCertificateResponse), nil
	}
}

type DeleteGeoipRuleInvoker struct {
	*invoker.BaseInvoker
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

func (i *DeletePrivacyRuleInvoker) Invoke() (*model.DeletePrivacyRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePrivacyRuleResponse), nil
	}
}

type DeleteValueListInvoker struct {
	*invoker.BaseInvoker
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

func (i *DeleteWhiteBlackIpRuleInvoker) Invoke() (*model.DeleteWhiteBlackIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWhiteBlackIpRuleResponse), nil
	}
}

type ListAntitamperRuleInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListBandwidthTimelineInvoker) Invoke() (*model.ListBandwidthTimelineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBandwidthTimelineResponse), nil
	}
}

type ListCertificatesInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListCompositeHostsInvoker) Invoke() (*model.ListCompositeHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCompositeHostsResponse), nil
	}
}

type ListEventInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListPrivacyRuleInvoker) Invoke() (*model.ListPrivacyRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPrivacyRuleResponse), nil
	}
}

type ListQpsTimelineInvoker struct {
	*invoker.BaseInvoker
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

func (i *RenameInstanceInvoker) Invoke() (*model.RenameInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RenameInstanceResponse), nil
	}
}

type ShowCertificateInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowConsoleConfigInvoker) Invoke() (*model.ShowConsoleConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConsoleConfigResponse), nil
	}
}

type ShowEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEventInvoker) Invoke() (*model.ShowEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEventResponse), nil
	}
}

type ShowHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHostInvoker) Invoke() (*model.ShowHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHostResponse), nil
	}
}

type ShowInstanceInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowPremiumHostInvoker) Invoke() (*model.ShowPremiumHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPremiumHostResponse), nil
	}
}

type ShowSourceIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSourceIpInvoker) Invoke() (*model.ShowSourceIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSourceIpResponse), nil
	}
}

type UpdateAlertNoticeConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAlertNoticeConfigInvoker) Invoke() (*model.UpdateAlertNoticeConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAlertNoticeConfigResponse), nil
	}
}

type UpdateCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCertificateInvoker) Invoke() (*model.UpdateCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCertificateResponse), nil
	}
}

type UpdateGeoipRuleInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateHostProtectStatusInvoker) Invoke() (*model.UpdateHostProtectStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHostProtectStatusResponse), nil
	}
}

type UpdateIpGroupInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdatePrivacyRuleInvoker) Invoke() (*model.UpdatePrivacyRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePrivacyRuleResponse), nil
	}
}

type UpdateValueListInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateWhiteblackipRuleInvoker) Invoke() (*model.UpdateWhiteblackipRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWhiteblackipRuleResponse), nil
	}
}
