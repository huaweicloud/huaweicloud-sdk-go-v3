package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eg/v1/model"
)

type CreateChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateChannelInvoker) Invoke() (*model.CreateChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateChannelResponse), nil
	}
}

type CreateEventSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEventSourceInvoker) Invoke() (*model.CreateEventSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEventSourceResponse), nil
	}
}

type CreateSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSubscriptionInvoker) Invoke() (*model.CreateSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSubscriptionResponse), nil
	}
}

type CreateSubscriptionTargetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSubscriptionTargetInvoker) Invoke() (*model.CreateSubscriptionTargetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSubscriptionTargetResponse), nil
	}
}

type DeleteChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteChannelInvoker) Invoke() (*model.DeleteChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteChannelResponse), nil
	}
}

type DeleteEventSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEventSourceInvoker) Invoke() (*model.DeleteEventSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEventSourceResponse), nil
	}
}

type DeleteSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSubscriptionInvoker) Invoke() (*model.DeleteSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSubscriptionResponse), nil
	}
}

type DeleteSubscriptionTargetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSubscriptionTargetInvoker) Invoke() (*model.DeleteSubscriptionTargetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSubscriptionTargetResponse), nil
	}
}

type ListChannelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListChannelsInvoker) Invoke() (*model.ListChannelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListChannelsResponse), nil
	}
}

type ListEventSourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventSourcesInvoker) Invoke() (*model.ListEventSourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventSourcesResponse), nil
	}
}

type ListEventTargetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventTargetInvoker) Invoke() (*model.ListEventTargetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventTargetResponse), nil
	}
}

type ListQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotasInvoker) Invoke() (*model.ListQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotasResponse), nil
	}
}

type ListSubscriptionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubscriptionsInvoker) Invoke() (*model.ListSubscriptionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubscriptionsResponse), nil
	}
}

type OperateSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *OperateSubscriptionInvoker) Invoke() (*model.OperateSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.OperateSubscriptionResponse), nil
	}
}

type PutEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *PutEventsInvoker) Invoke() (*model.PutEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PutEventsResponse), nil
	}
}

type ShowDetailOfChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailOfChannelInvoker) Invoke() (*model.ShowDetailOfChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailOfChannelResponse), nil
	}
}

type ShowDetailOfEventSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailOfEventSourceInvoker) Invoke() (*model.ShowDetailOfEventSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailOfEventSourceResponse), nil
	}
}

type ShowDetailOfSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailOfSubscriptionInvoker) Invoke() (*model.ShowDetailOfSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailOfSubscriptionResponse), nil
	}
}

type ShowDetailOfSubscriptionTargetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailOfSubscriptionTargetInvoker) Invoke() (*model.ShowDetailOfSubscriptionTargetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailOfSubscriptionTargetResponse), nil
	}
}

type UpdateChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateChannelInvoker) Invoke() (*model.UpdateChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateChannelResponse), nil
	}
}

type UpdateEventSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEventSourceInvoker) Invoke() (*model.UpdateEventSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEventSourceResponse), nil
	}
}

type UpdateSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSubscriptionInvoker) Invoke() (*model.UpdateSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSubscriptionResponse), nil
	}
}

type UpdateSubscriptionSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSubscriptionSourceInvoker) Invoke() (*model.UpdateSubscriptionSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSubscriptionSourceResponse), nil
	}
}

type UpdateSubscriptionTargetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSubscriptionTargetInvoker) Invoke() (*model.UpdateSubscriptionTargetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSubscriptionTargetResponse), nil
	}
}
