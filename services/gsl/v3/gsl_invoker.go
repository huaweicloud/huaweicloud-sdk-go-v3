package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/gsl/v3/model"
)

type BatchSetAttributesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSetAttributesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchSetAttributesInvoker) Invoke() (*model.BatchSetAttributesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSetAttributesResponse), nil
	}
}

type CreateAttributeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAttributeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAttributeInvoker) Invoke() (*model.CreateAttributeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAttributeResponse), nil
	}
}

type DisableAttributeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableAttributeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableAttributeInvoker) Invoke() (*model.DisableAttributeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableAttributeResponse), nil
	}
}

type EnableAttributeInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableAttributeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableAttributeInvoker) Invoke() (*model.EnableAttributeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableAttributeResponse), nil
	}
}

type ListAttributesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAttributesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAttributesInvoker) Invoke() (*model.ListAttributesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAttributesResponse), nil
	}
}

type UpdateAttributeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAttributeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAttributeInvoker) Invoke() (*model.UpdateAttributeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAttributeResponse), nil
	}
}

type ListBackPoolMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackPoolMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBackPoolMembersInvoker) Invoke() (*model.ListBackPoolMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackPoolMembersResponse), nil
	}
}

type ListBackPoolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackPoolsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBackPoolsInvoker) Invoke() (*model.ListBackPoolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackPoolsResponse), nil
	}
}

type AddNetworkSwitchPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddNetworkSwitchPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddNetworkSwitchPolicyInvoker) Invoke() (*model.AddNetworkSwitchPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddNetworkSwitchPolicyResponse), nil
	}
}

type ListNetworkSwitchPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNetworkSwitchPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNetworkSwitchPoliciesInvoker) Invoke() (*model.ListNetworkSwitchPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNetworkSwitchPoliciesResponse), nil
	}
}

type ListProPricePlansInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProPricePlansInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProPricePlansInvoker) Invoke() (*model.ListProPricePlansResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProPricePlansResponse), nil
	}
}

type DeleteRealNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRealNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRealNameInvoker) Invoke() (*model.DeleteRealNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRealNameResponse), nil
	}
}

type EnableSimCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableSimCardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableSimCardInvoker) Invoke() (*model.EnableSimCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableSimCardResponse), nil
	}
}

type ListSimCardFlowPerDayInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSimCardFlowPerDayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSimCardFlowPerDayInvoker) Invoke() (*model.ListSimCardFlowPerDayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSimCardFlowPerDayResponse), nil
	}
}

type ListSimCardsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSimCardsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSimCardsInvoker) Invoke() (*model.ListSimCardsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSimCardsResponse), nil
	}
}

type RegisterImeiInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterImeiInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RegisterImeiInvoker) Invoke() (*model.RegisterImeiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterImeiResponse), nil
	}
}

type ResetSimCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetSimCardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetSimCardInvoker) Invoke() (*model.ResetSimCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetSimCardResponse), nil
	}
}

type SetExceedCutNetInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetExceedCutNetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetExceedCutNetInvoker) Invoke() (*model.SetExceedCutNetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetExceedCutNetResponse), nil
	}
}

type SetSpeedValueInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetSpeedValueInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetSpeedValueInvoker) Invoke() (*model.SetSpeedValueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetSpeedValueResponse), nil
	}
}

type ShowMonthUsagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMonthUsagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMonthUsagesInvoker) Invoke() (*model.ShowMonthUsagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMonthUsagesResponse), nil
	}
}

type ShowRealNamedInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRealNamedInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRealNamedInvoker) Invoke() (*model.ShowRealNamedResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRealNamedResponse), nil
	}
}

type ShowSimCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSimCardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSimCardInvoker) Invoke() (*model.ShowSimCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSimCardResponse), nil
	}
}

type StartStopNetInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartStopNetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartStopNetInvoker) Invoke() (*model.StartStopNetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartStopNetResponse), nil
	}
}

type StopSimCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopSimCardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopSimCardInvoker) Invoke() (*model.StopSimCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopSimCardResponse), nil
	}
}

type ListSimDeviceMultiplyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSimDeviceMultiplyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSimDeviceMultiplyInvoker) Invoke() (*model.ListSimDeviceMultiplyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSimDeviceMultiplyResponse), nil
	}
}

type SetNetworkSwitchPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetNetworkSwitchPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetNetworkSwitchPolicyInvoker) Invoke() (*model.SetNetworkSwitchPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetNetworkSwitchPolicyResponse), nil
	}
}

type SwitchNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchNetworkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchNetworkInvoker) Invoke() (*model.SwitchNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchNetworkResponse), nil
	}
}

type ListSimPoolMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSimPoolMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSimPoolMembersInvoker) Invoke() (*model.ListSimPoolMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSimPoolMembersResponse), nil
	}
}

type ListSimPoolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSimPoolsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSimPoolsInvoker) Invoke() (*model.ListSimPoolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSimPoolsResponse), nil
	}
}

type ListFlowBySimCardsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlowBySimCardsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFlowBySimCardsInvoker) Invoke() (*model.ListFlowBySimCardsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlowBySimCardsResponse), nil
	}
}

type ListSimPricePlansInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSimPricePlansInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSimPricePlansInvoker) Invoke() (*model.ListSimPricePlansResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSimPricePlansResponse), nil
	}
}

type ListSmsDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSmsDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSmsDetailsInvoker) Invoke() (*model.ListSmsDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSmsDetailsResponse), nil
	}
}

type SendSmsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendSmsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SendSmsInvoker) Invoke() (*model.SendSmsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendSmsResponse), nil
	}
}

type BatchSetTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSetTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchSetTagsInvoker) Invoke() (*model.BatchSetTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSetTagsResponse), nil
	}
}

type CreateTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTagInvoker) Invoke() (*model.CreateTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTagResponse), nil
	}
}

type DeleteTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTagInvoker) Invoke() (*model.DeleteTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTagResponse), nil
	}
}

type ListTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagsInvoker) Invoke() (*model.ListTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagsResponse), nil
	}
}

type ListWorkOrderDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkOrderDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkOrderDetailsInvoker) Invoke() (*model.ListWorkOrderDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkOrderDetailsResponse), nil
	}
}

type ListWorkOrdersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkOrdersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkOrdersInvoker) Invoke() (*model.ListWorkOrdersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkOrdersResponse), nil
	}
}
