package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/nat/v2/model"
)

type BatchCreateNatGatewayDnatRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateNatGatewayDnatRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateNatGatewayDnatRulesInvoker) Invoke() (*model.BatchCreateNatGatewayDnatRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateNatGatewayDnatRulesResponse), nil
	}
}

type CreateNatGatewayDnatRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNatGatewayDnatRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNatGatewayDnatRuleInvoker) Invoke() (*model.CreateNatGatewayDnatRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNatGatewayDnatRuleResponse), nil
	}
}

type CreatePrivateDnatInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePrivateDnatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePrivateDnatInvoker) Invoke() (*model.CreatePrivateDnatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePrivateDnatResponse), nil
	}
}

type DeleteNatGatewayDnatRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNatGatewayDnatRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteNatGatewayDnatRuleInvoker) Invoke() (*model.DeleteNatGatewayDnatRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNatGatewayDnatRuleResponse), nil
	}
}

type DeletePrivateDnatInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePrivateDnatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePrivateDnatInvoker) Invoke() (*model.DeletePrivateDnatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePrivateDnatResponse), nil
	}
}

type ListNatGatewayDnatRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNatGatewayDnatRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNatGatewayDnatRulesInvoker) Invoke() (*model.ListNatGatewayDnatRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNatGatewayDnatRulesResponse), nil
	}
}

type ListPrivateDnatsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPrivateDnatsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPrivateDnatsInvoker) Invoke() (*model.ListPrivateDnatsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPrivateDnatsResponse), nil
	}
}

type ShowNatGatewayDnatRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNatGatewayDnatRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNatGatewayDnatRuleInvoker) Invoke() (*model.ShowNatGatewayDnatRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNatGatewayDnatRuleResponse), nil
	}
}

type ShowPrivateDnatInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPrivateDnatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPrivateDnatInvoker) Invoke() (*model.ShowPrivateDnatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPrivateDnatResponse), nil
	}
}

type UpdateNatGatewayDnatRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNatGatewayDnatRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNatGatewayDnatRuleInvoker) Invoke() (*model.UpdateNatGatewayDnatRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNatGatewayDnatRuleResponse), nil
	}
}

type UpdatePrivateDnatInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePrivateDnatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePrivateDnatInvoker) Invoke() (*model.UpdatePrivateDnatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePrivateDnatResponse), nil
	}
}

type BatchCreateDeleteTransitIpTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateDeleteTransitIpTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateDeleteTransitIpTagsInvoker) Invoke() (*model.BatchCreateDeleteTransitIpTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateDeleteTransitIpTagsResponse), nil
	}
}

type CreateTransitIpTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTransitIpTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTransitIpTagInvoker) Invoke() (*model.CreateTransitIpTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTransitIpTagResponse), nil
	}
}

type DeleteTransitIpTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTransitIpTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTransitIpTagInvoker) Invoke() (*model.DeleteTransitIpTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTransitIpTagResponse), nil
	}
}

type ListTransitIpTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTransitIpTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTransitIpTagsInvoker) Invoke() (*model.ListTransitIpTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTransitIpTagsResponse), nil
	}
}

type ListTransitIpsByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTransitIpsByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTransitIpsByTagsInvoker) Invoke() (*model.ListTransitIpsByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTransitIpsByTagsResponse), nil
	}
}

type ShowTransitIpTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTransitIpTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTransitIpTagsInvoker) Invoke() (*model.ShowTransitIpTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTransitIpTagsResponse), nil
	}
}

type BatchCreateDeleteNatGatewayTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateDeleteNatGatewayTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateDeleteNatGatewayTagInvoker) Invoke() (*model.BatchCreateDeleteNatGatewayTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateDeleteNatGatewayTagResponse), nil
	}
}

type BatchCreateDeletePrivateNatTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateDeletePrivateNatTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateDeletePrivateNatTagsInvoker) Invoke() (*model.BatchCreateDeletePrivateNatTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateDeletePrivateNatTagsResponse), nil
	}
}

type CreateNatGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNatGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNatGatewayInvoker) Invoke() (*model.CreateNatGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNatGatewayResponse), nil
	}
}

type CreateNatGatewayTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNatGatewayTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNatGatewayTagInvoker) Invoke() (*model.CreateNatGatewayTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNatGatewayTagResponse), nil
	}
}

type CreatePrivateNatInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePrivateNatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePrivateNatInvoker) Invoke() (*model.CreatePrivateNatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePrivateNatResponse), nil
	}
}

type CreatePrivateNatTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePrivateNatTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePrivateNatTagInvoker) Invoke() (*model.CreatePrivateNatTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePrivateNatTagResponse), nil
	}
}

type DeleteNatGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNatGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteNatGatewayInvoker) Invoke() (*model.DeleteNatGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNatGatewayResponse), nil
	}
}

type DeleteNatGatewayTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNatGatewayTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteNatGatewayTagInvoker) Invoke() (*model.DeleteNatGatewayTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNatGatewayTagResponse), nil
	}
}

type DeletePrivateNatInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePrivateNatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePrivateNatInvoker) Invoke() (*model.DeletePrivateNatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePrivateNatResponse), nil
	}
}

type DeletePrivateNatTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePrivateNatTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePrivateNatTagInvoker) Invoke() (*model.DeletePrivateNatTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePrivateNatTagResponse), nil
	}
}

type ListNatGatewayByTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNatGatewayByTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNatGatewayByTagInvoker) Invoke() (*model.ListNatGatewayByTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNatGatewayByTagResponse), nil
	}
}

type ListNatGatewayTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNatGatewayTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNatGatewayTagInvoker) Invoke() (*model.ListNatGatewayTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNatGatewayTagResponse), nil
	}
}

type ListNatGatewaysInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNatGatewaysInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNatGatewaysInvoker) Invoke() (*model.ListNatGatewaysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNatGatewaysResponse), nil
	}
}

type ListPrivateNatTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPrivateNatTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPrivateNatTagsInvoker) Invoke() (*model.ListPrivateNatTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPrivateNatTagsResponse), nil
	}
}

type ListPrivateNatsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPrivateNatsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPrivateNatsInvoker) Invoke() (*model.ListPrivateNatsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPrivateNatsResponse), nil
	}
}

type ListPrivateNatsByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPrivateNatsByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPrivateNatsByTagsInvoker) Invoke() (*model.ListPrivateNatsByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPrivateNatsByTagsResponse), nil
	}
}

type ShowNatGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNatGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNatGatewayInvoker) Invoke() (*model.ShowNatGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNatGatewayResponse), nil
	}
}

type ShowNatGatewayTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNatGatewayTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNatGatewayTagInvoker) Invoke() (*model.ShowNatGatewayTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNatGatewayTagResponse), nil
	}
}

type ShowPrivateNatInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPrivateNatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPrivateNatInvoker) Invoke() (*model.ShowPrivateNatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPrivateNatResponse), nil
	}
}

type ShowPrivateNatTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPrivateNatTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPrivateNatTagsInvoker) Invoke() (*model.ShowPrivateNatTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPrivateNatTagsResponse), nil
	}
}

type UpdateNatGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNatGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNatGatewayInvoker) Invoke() (*model.UpdateNatGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNatGatewayResponse), nil
	}
}

type UpdatePrivateNatInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePrivateNatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePrivateNatInvoker) Invoke() (*model.UpdatePrivateNatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePrivateNatResponse), nil
	}
}

type CreateTransitIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTransitIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTransitIpInvoker) Invoke() (*model.CreateTransitIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTransitIpResponse), nil
	}
}

type DeleteTransitIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTransitIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTransitIpInvoker) Invoke() (*model.DeleteTransitIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTransitIpResponse), nil
	}
}

type ListTransitIpsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTransitIpsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTransitIpsInvoker) Invoke() (*model.ListTransitIpsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTransitIpsResponse), nil
	}
}

type ShowTransitIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTransitIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTransitIpInvoker) Invoke() (*model.ShowTransitIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTransitIpResponse), nil
	}
}

type CreateNatGatewaySnatRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNatGatewaySnatRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNatGatewaySnatRuleInvoker) Invoke() (*model.CreateNatGatewaySnatRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNatGatewaySnatRuleResponse), nil
	}
}

type CreatePrivateSnatInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePrivateSnatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePrivateSnatInvoker) Invoke() (*model.CreatePrivateSnatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePrivateSnatResponse), nil
	}
}

type DeleteNatGatewaySnatRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNatGatewaySnatRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteNatGatewaySnatRuleInvoker) Invoke() (*model.DeleteNatGatewaySnatRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNatGatewaySnatRuleResponse), nil
	}
}

type DeletePrivateSnatInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePrivateSnatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePrivateSnatInvoker) Invoke() (*model.DeletePrivateSnatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePrivateSnatResponse), nil
	}
}

type ListNatGatewaySnatRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNatGatewaySnatRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNatGatewaySnatRulesInvoker) Invoke() (*model.ListNatGatewaySnatRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNatGatewaySnatRulesResponse), nil
	}
}

type ListPrivateSnatsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPrivateSnatsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPrivateSnatsInvoker) Invoke() (*model.ListPrivateSnatsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPrivateSnatsResponse), nil
	}
}

type ShowNatGatewaySnatRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNatGatewaySnatRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNatGatewaySnatRuleInvoker) Invoke() (*model.ShowNatGatewaySnatRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNatGatewaySnatRuleResponse), nil
	}
}

type ShowPrivateSnatInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPrivateSnatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPrivateSnatInvoker) Invoke() (*model.ShowPrivateSnatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPrivateSnatResponse), nil
	}
}

type UpdateNatGatewaySnatRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNatGatewaySnatRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNatGatewaySnatRuleInvoker) Invoke() (*model.UpdateNatGatewaySnatRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNatGatewaySnatRuleResponse), nil
	}
}

type UpdatePrivateSnatInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePrivateSnatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePrivateSnatInvoker) Invoke() (*model.UpdatePrivateSnatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePrivateSnatResponse), nil
	}
}
