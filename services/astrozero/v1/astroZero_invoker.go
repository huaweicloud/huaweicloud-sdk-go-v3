package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/astrozero/v1/model"
)

type ShowOrderStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOrderStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOrderStatusInvoker) Invoke() (*model.ShowOrderStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOrderStatusResponse), nil
	}
}
