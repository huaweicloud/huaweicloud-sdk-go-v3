package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/smnglobal/v2/model"
)

type CreateSubscriptionUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSubscriptionUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSubscriptionUserInvoker) Invoke() (*model.CreateSubscriptionUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSubscriptionUserResponse), nil
	}
}

type DeleteSubscriptionUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSubscriptionUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSubscriptionUserInvoker) Invoke() (*model.DeleteSubscriptionUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSubscriptionUserResponse), nil
	}
}

type ListSubscriptionUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubscriptionUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubscriptionUserInvoker) Invoke() (*model.ListSubscriptionUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubscriptionUserResponse), nil
	}
}

type UpdateSubscriptionUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSubscriptionUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSubscriptionUserInvoker) Invoke() (*model.UpdateSubscriptionUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSubscriptionUserResponse), nil
	}
}
