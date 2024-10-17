package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/sa/v1/model"
)

type CheckProductHealthyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckProductHealthyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckProductHealthyInvoker) Invoke() (*model.CheckProductHealthyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckProductHealthyResponse), nil
	}
}

type ImportEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportEventsInvoker) Invoke() (*model.ImportEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportEventsResponse), nil
	}
}
