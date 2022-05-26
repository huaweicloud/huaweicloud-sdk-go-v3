package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/apm/v1/model"
)

type ListAkSkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAkSkInvoker) Invoke() (*model.ListAkSkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAkSkResponse), nil
	}
}

type ListBusinessInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBusinessInvoker) Invoke() (*model.ListBusinessResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBusinessResponse), nil
	}
}

type ShowMasterAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMasterAddressInvoker) Invoke() (*model.ShowMasterAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMasterAddressResponse), nil
	}
}
