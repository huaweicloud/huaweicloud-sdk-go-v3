package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/aad/v2/model"
)

type AddWafWhiteIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddWafWhiteIpRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddWafWhiteIpRuleInvoker) Invoke() (*model.AddWafWhiteIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddWafWhiteIpRuleResponse), nil
	}
}

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

type DeleteWafWhiteIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWafWhiteIpRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteWafWhiteIpRuleInvoker) Invoke() (*model.DeleteWafWhiteIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWafWhiteIpRuleResponse), nil
	}
}

type ListDDoSAttackEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDDoSAttackEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDDoSAttackEventInvoker) Invoke() (*model.ListDDoSAttackEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDDoSAttackEventResponse), nil
	}
}

type ListDDoSBlackHoleEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDDoSBlackHoleEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDDoSBlackHoleEventInvoker) Invoke() (*model.ListDDoSBlackHoleEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDDoSBlackHoleEventResponse), nil
	}
}

type ListDDoSConnectionNumberInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDDoSConnectionNumberInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDDoSConnectionNumberInvoker) Invoke() (*model.ListDDoSConnectionNumberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDDoSConnectionNumberResponse), nil
	}
}

type ListDDoSFlowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDDoSFlowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDDoSFlowInvoker) Invoke() (*model.ListDDoSFlowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDDoSFlowResponse), nil
	}
}

type ListFrequencyControlRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFrequencyControlRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFrequencyControlRuleInvoker) Invoke() (*model.ListFrequencyControlRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFrequencyControlRuleResponse), nil
	}
}

type ListGlobalConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGlobalConfigInvoker) Invoke() (*model.ListGlobalConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalConfigResponse), nil
	}
}

type ListInstanceDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceDomainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceDomainsInvoker) Invoke() (*model.ListInstanceDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceDomainsResponse), nil
	}
}

type ListSourceIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSourceIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSourceIpInvoker) Invoke() (*model.ListSourceIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSourceIpResponse), nil
	}
}

type ListWafAttackEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWafAttackEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWafAttackEventInvoker) Invoke() (*model.ListWafAttackEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWafAttackEventResponse), nil
	}
}

type ListWafBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWafBandwidthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWafBandwidthInvoker) Invoke() (*model.ListWafBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWafBandwidthResponse), nil
	}
}

type ListWafCustomRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWafCustomRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWafCustomRuleInvoker) Invoke() (*model.ListWafCustomRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWafCustomRuleResponse), nil
	}
}

type ListWafGeoIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWafGeoIpRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWafGeoIpRuleInvoker) Invoke() (*model.ListWafGeoIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWafGeoIpRuleResponse), nil
	}
}

type ListWafQpsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWafQpsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWafQpsInvoker) Invoke() (*model.ListWafQpsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWafQpsResponse), nil
	}
}

type ListWafWhiteIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWafWhiteIpRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWafWhiteIpRuleInvoker) Invoke() (*model.ListWafWhiteIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWafWhiteIpRuleResponse), nil
	}
}

type ListWhiteBlackIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWhiteBlackIpRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWhiteBlackIpRuleInvoker) Invoke() (*model.ListWhiteBlackIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWhiteBlackIpRuleResponse), nil
	}
}

type ShowAlarmConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAlarmConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAlarmConfigInvoker) Invoke() (*model.ShowAlarmConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAlarmConfigResponse), nil
	}
}

type ShowDDoSPeakInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDDoSPeakInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDDoSPeakInvoker) Invoke() (*model.ShowDDoSPeakResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDDoSPeakResponse), nil
	}
}

type ShowDomainCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDomainCertificateInvoker) Invoke() (*model.ShowDomainCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainCertificateResponse), nil
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

type ShowDomainNameConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainNameConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDomainNameConfigInvoker) Invoke() (*model.ShowDomainNameConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainNameConfigResponse), nil
	}
}

type ShowFlowBlockInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlowBlockInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFlowBlockInvoker) Invoke() (*model.ShowFlowBlockResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlowBlockResponse), nil
	}
}

type ShowInstanceByInstanceIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceByInstanceIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceByInstanceIdInvoker) Invoke() (*model.ShowInstanceByInstanceIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceByInstanceIdResponse), nil
	}
}

type ShowWafPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWafPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWafPolicyInvoker) Invoke() (*model.ShowWafPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWafPolicyResponse), nil
	}
}

type ShowWafQpsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWafQpsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWafQpsInvoker) Invoke() (*model.ShowWafQpsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWafQpsResponse), nil
	}
}

type UpdateDomainConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDomainConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDomainConfigInvoker) Invoke() (*model.UpdateDomainConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainConfigResponse), nil
	}
}

type UpdateForwardRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateForwardRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateForwardRuleInvoker) Invoke() (*model.UpdateForwardRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateForwardRuleResponse), nil
	}
}

type UpgradeInstanceSpecInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeInstanceSpecInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpgradeInstanceSpecInvoker) Invoke() (*model.UpgradeInstanceSpecResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeInstanceSpecResponse), nil
	}
}
