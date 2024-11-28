package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartsinspector/v2/model"
)

type CreatePurchaseOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePurchaseOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePurchaseOrderInvoker) Invoke() (*model.CreatePurchaseOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePurchaseOrderResponse), nil
	}
}

type UpdatePurchaseOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePurchaseOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePurchaseOrderInvoker) Invoke() (*model.UpdatePurchaseOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePurchaseOrderResponse), nil
	}
}
