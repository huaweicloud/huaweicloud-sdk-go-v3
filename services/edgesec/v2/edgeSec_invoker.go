package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/edgesec/v2/model"
)

type CreateDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDomainsInvoker) Invoke() (*model.CreateDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDomainsResponse), nil
	}
}

type DeleteDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDomainsInvoker) Invoke() (*model.DeleteDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDomainsResponse), nil
	}
}

type ShowDomainDetailInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowDomainsInvoker) Invoke() (*model.ShowDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainsResponse), nil
	}
}

type UpdateDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDomainsInvoker) Invoke() (*model.UpdateDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainsResponse), nil
	}
}
