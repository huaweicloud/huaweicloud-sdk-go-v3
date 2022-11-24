package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/nat/v2/model"
)

type BatchCreateNatGatewayDnatRulesInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreateNatGatewayDnatRuleInvoker) Invoke() (*model.CreateNatGatewayDnatRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNatGatewayDnatRuleResponse), nil
	}
}

type DeleteNatGatewayDnatRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNatGatewayDnatRuleInvoker) Invoke() (*model.DeleteNatGatewayDnatRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNatGatewayDnatRuleResponse), nil
	}
}

type ListNatGatewayDnatRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNatGatewayDnatRulesInvoker) Invoke() (*model.ListNatGatewayDnatRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNatGatewayDnatRulesResponse), nil
	}
}

type ShowNatGatewayDnatRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNatGatewayDnatRuleInvoker) Invoke() (*model.ShowNatGatewayDnatRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNatGatewayDnatRuleResponse), nil
	}
}

type UpdateNatGatewayDnatRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNatGatewayDnatRuleInvoker) Invoke() (*model.UpdateNatGatewayDnatRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNatGatewayDnatRuleResponse), nil
	}
}

type CreateNatGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNatGatewayInvoker) Invoke() (*model.CreateNatGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNatGatewayResponse), nil
	}
}

type DeleteNatGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNatGatewayInvoker) Invoke() (*model.DeleteNatGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNatGatewayResponse), nil
	}
}

type ListNatGatewaysInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNatGatewaysInvoker) Invoke() (*model.ListNatGatewaysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNatGatewaysResponse), nil
	}
}

type ShowNatGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNatGatewayInvoker) Invoke() (*model.ShowNatGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNatGatewayResponse), nil
	}
}

type UpdateNatGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNatGatewayInvoker) Invoke() (*model.UpdateNatGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNatGatewayResponse), nil
	}
}

type CreateNatGatewaySnatRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNatGatewaySnatRuleInvoker) Invoke() (*model.CreateNatGatewaySnatRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNatGatewaySnatRuleResponse), nil
	}
}

type DeleteNatGatewaySnatRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNatGatewaySnatRuleInvoker) Invoke() (*model.DeleteNatGatewaySnatRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNatGatewaySnatRuleResponse), nil
	}
}

type ListNatGatewaySnatRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNatGatewaySnatRulesInvoker) Invoke() (*model.ListNatGatewaySnatRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNatGatewaySnatRulesResponse), nil
	}
}

type ShowNatGatewaySnatRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNatGatewaySnatRuleInvoker) Invoke() (*model.ShowNatGatewaySnatRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNatGatewaySnatRuleResponse), nil
	}
}

type UpdateNatGatewaySnatRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNatGatewaySnatRuleInvoker) Invoke() (*model.UpdateNatGatewaySnatRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNatGatewaySnatRuleResponse), nil
	}
}
