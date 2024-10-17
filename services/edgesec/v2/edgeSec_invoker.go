package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/edgesec/v2/model"
)

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

type CreateDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDomainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDomainsInvoker) Invoke() (*model.CreateDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDomainsResponse), nil
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

type DeleteDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDomainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDomainsInvoker) Invoke() (*model.DeleteDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDomainsResponse), nil
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

type UpdateDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDomainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDomainsInvoker) Invoke() (*model.UpdateDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainsResponse), nil
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
