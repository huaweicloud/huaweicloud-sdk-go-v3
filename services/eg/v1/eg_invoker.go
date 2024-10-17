package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eg/v1/model"
)

type CheckPutEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckPutEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckPutEventsInvoker) Invoke() (*model.CheckPutEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckPutEventsResponse), nil
	}
}

type CreateAgenciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAgenciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAgenciesInvoker) Invoke() (*model.CreateAgenciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAgenciesResponse), nil
	}
}

type CreateChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateChannelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateChannelInvoker) Invoke() (*model.CreateChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateChannelResponse), nil
	}
}

type CreateConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateConnectionInvoker) Invoke() (*model.CreateConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConnectionResponse), nil
	}
}

type CreateEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEndpointInvoker) Invoke() (*model.CreateEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEndpointResponse), nil
	}
}

type CreateEventSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEventSourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEventSourceInvoker) Invoke() (*model.CreateEventSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEventSourceResponse), nil
	}
}

type CreateEventStreamingInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEventStreamingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEventStreamingInvoker) Invoke() (*model.CreateEventStreamingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEventStreamingResponse), nil
	}
}

type CreateSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSubscriptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateSubscriptionTargetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteChannelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteChannelInvoker) Invoke() (*model.DeleteChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteChannelResponse), nil
	}
}

type DeleteConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteConnectionInvoker) Invoke() (*model.DeleteConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConnectionResponse), nil
	}
}

type DeleteEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEndpointInvoker) Invoke() (*model.DeleteEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEndpointResponse), nil
	}
}

type DeleteEventSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEventSourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEventSourceInvoker) Invoke() (*model.DeleteEventSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEventSourceResponse), nil
	}
}

type DeleteEventStreamingInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEventStreamingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEventStreamingInvoker) Invoke() (*model.DeleteEventStreamingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEventStreamingResponse), nil
	}
}

type DeleteSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSubscriptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteSubscriptionTargetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSubscriptionTargetInvoker) Invoke() (*model.DeleteSubscriptionTargetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSubscriptionTargetResponse), nil
	}
}

type ListAgenciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAgenciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAgenciesInvoker) Invoke() (*model.ListAgenciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAgenciesResponse), nil
	}
}

type ListChannelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListChannelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListChannelsInvoker) Invoke() (*model.ListChannelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListChannelsResponse), nil
	}
}

type ListConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConnectionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConnectionsInvoker) Invoke() (*model.ListConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConnectionsResponse), nil
	}
}

type ListEndpointsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEndpointsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEndpointsInvoker) Invoke() (*model.ListEndpointsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEndpointsResponse), nil
	}
}

type ListEventSourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventSourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEventSourcesInvoker) Invoke() (*model.ListEventSourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventSourcesResponse), nil
	}
}

type ListEventStreamingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventStreamingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEventStreamingInvoker) Invoke() (*model.ListEventStreamingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventStreamingResponse), nil
	}
}

type ListEventTargetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventTargetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEventTargetInvoker) Invoke() (*model.ListEventTargetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventTargetResponse), nil
	}
}

type ListPubMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPubMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPubMetricsInvoker) Invoke() (*model.ListPubMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPubMetricsResponse), nil
	}
}

type ListQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuotasInvoker) Invoke() (*model.ListQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotasResponse), nil
	}
}

type ListSubMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubMetricsInvoker) Invoke() (*model.ListSubMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubMetricsResponse), nil
	}
}

type ListSubscriptionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubscriptionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubscriptionsInvoker) Invoke() (*model.ListSubscriptionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubscriptionsResponse), nil
	}
}

type ListTracedEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTracedEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTracedEventsInvoker) Invoke() (*model.ListTracedEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTracedEventsResponse), nil
	}
}

type ListTriggersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTriggersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTriggersInvoker) Invoke() (*model.ListTriggersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTriggersResponse), nil
	}
}

type ListWorkflowTriggersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkflowTriggersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkflowTriggersInvoker) Invoke() (*model.ListWorkflowTriggersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkflowTriggersResponse), nil
	}
}

type OperateSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *OperateSubscriptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *PutEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PutEventsInvoker) Invoke() (*model.PutEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PutEventsResponse), nil
	}
}

type PutOfficialEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *PutOfficialEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PutOfficialEventsInvoker) Invoke() (*model.PutOfficialEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PutOfficialEventsResponse), nil
	}
}

type ResumeEventStreamingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResumeEventStreamingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResumeEventStreamingInvoker) Invoke() (*model.ResumeEventStreamingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResumeEventStreamingResponse), nil
	}
}

type ShowDetailOfChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailOfChannelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailOfChannelInvoker) Invoke() (*model.ShowDetailOfChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailOfChannelResponse), nil
	}
}

type ShowDetailOfConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailOfConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailOfConnectionInvoker) Invoke() (*model.ShowDetailOfConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailOfConnectionResponse), nil
	}
}

type ShowDetailOfEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailOfEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailOfEventInvoker) Invoke() (*model.ShowDetailOfEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailOfEventResponse), nil
	}
}

type ShowDetailOfEventSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailOfEventSourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailOfEventSourceInvoker) Invoke() (*model.ShowDetailOfEventSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailOfEventSourceResponse), nil
	}
}

type ShowDetailOfEventTraceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailOfEventTraceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailOfEventTraceInvoker) Invoke() (*model.ShowDetailOfEventTraceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailOfEventTraceResponse), nil
	}
}

type ShowDetailOfSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailOfSubscriptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowDetailOfSubscriptionTargetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailOfSubscriptionTargetInvoker) Invoke() (*model.ShowDetailOfSubscriptionTargetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailOfSubscriptionTargetResponse), nil
	}
}

type ShowEventStreamingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEventStreamingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEventStreamingInvoker) Invoke() (*model.ShowEventStreamingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEventStreamingResponse), nil
	}
}

type UpdateChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateChannelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateChannelInvoker) Invoke() (*model.UpdateChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateChannelResponse), nil
	}
}

type UpdateConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateConnectionInvoker) Invoke() (*model.UpdateConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateConnectionResponse), nil
	}
}

type UpdateEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEndpointInvoker) Invoke() (*model.UpdateEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEndpointResponse), nil
	}
}

type UpdateEventSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEventSourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEventSourceInvoker) Invoke() (*model.UpdateEventSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEventSourceResponse), nil
	}
}

type UpdateEventStreamingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEventStreamingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEventStreamingInvoker) Invoke() (*model.UpdateEventStreamingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEventStreamingResponse), nil
	}
}

type UpdateSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSubscriptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateSubscriptionSourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateSubscriptionTargetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSubscriptionTargetInvoker) Invoke() (*model.UpdateSubscriptionTargetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSubscriptionTargetResponse), nil
	}
}

type ListApiVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApiVersionsInvoker) Invoke() (*model.ListApiVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiVersionsResponse), nil
	}
}

type ListObsBucketsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListObsBucketsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListObsBucketsInvoker) Invoke() (*model.ListObsBucketsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListObsBucketsResponse), nil
	}
}
