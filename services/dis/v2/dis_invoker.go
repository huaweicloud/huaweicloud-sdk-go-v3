package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dis/v2/model"
)

type BatchCreateTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateTagsInvoker) Invoke() (*model.BatchCreateTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateTagsResponse), nil
	}
}

type BatchDeleteTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteTagsInvoker) Invoke() (*model.BatchDeleteTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteTagsResponse), nil
	}
}

type BatchStartTransferTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchStartTransferTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchStartTransferTaskInvoker) Invoke() (*model.BatchStartTransferTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchStartTransferTaskResponse), nil
	}
}

type BatchStopTransferTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchStopTransferTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchStopTransferTaskInvoker) Invoke() (*model.BatchStopTransferTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchStopTransferTaskResponse), nil
	}
}

type ConsumeRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConsumeRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ConsumeRecordsInvoker) Invoke() (*model.ConsumeRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConsumeRecordsResponse), nil
	}
}

type CreateObsTransferTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateObsTransferTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateObsTransferTaskInvoker) Invoke() (*model.CreateObsTransferTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateObsTransferTaskResponse), nil
	}
}

type CreateStreamInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStreamInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateStreamInvoker) Invoke() (*model.CreateStreamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStreamResponse), nil
	}
}

type CreateTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTagInvoker) Invoke() (*model.CreateTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTagResponse), nil
	}
}

type DeleteStreamInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStreamInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteStreamInvoker) Invoke() (*model.DeleteStreamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStreamResponse), nil
	}
}

type DeleteTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTagInvoker) Invoke() (*model.DeleteTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTagResponse), nil
	}
}

type DeleteTransferTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTransferTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTransferTaskInvoker) Invoke() (*model.DeleteTransferTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTransferTaskResponse), nil
	}
}

type ListPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPoliciesInvoker) Invoke() (*model.ListPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPoliciesResponse), nil
	}
}

type ListResourcesByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourcesByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourcesByTagsInvoker) Invoke() (*model.ListResourcesByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourcesByTagsResponse), nil
	}
}

type ListStreamsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStreamsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStreamsInvoker) Invoke() (*model.ListStreamsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStreamsResponse), nil
	}
}

type ListTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagsInvoker) Invoke() (*model.ListTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagsResponse), nil
	}
}

type ListTransferTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTransferTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTransferTasksInvoker) Invoke() (*model.ListTransferTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTransferTasksResponse), nil
	}
}

type SendRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SendRecordsInvoker) Invoke() (*model.SendRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendRecordsResponse), nil
	}
}

type ShowCursorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCursorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCursorInvoker) Invoke() (*model.ShowCursorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCursorResponse), nil
	}
}

type ShowPartitionMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPartitionMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPartitionMetricsInvoker) Invoke() (*model.ShowPartitionMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPartitionMetricsResponse), nil
	}
}

type ShowStreamInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStreamInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStreamInvoker) Invoke() (*model.ShowStreamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStreamResponse), nil
	}
}

type ShowStreamMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStreamMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStreamMetricsInvoker) Invoke() (*model.ShowStreamMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStreamMetricsResponse), nil
	}
}

type ShowStreamTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStreamTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStreamTagsInvoker) Invoke() (*model.ShowStreamTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStreamTagsResponse), nil
	}
}

type ShowTransferTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTransferTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTransferTaskInvoker) Invoke() (*model.ShowTransferTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTransferTaskResponse), nil
	}
}

type UpdatePartitionCountInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePartitionCountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePartitionCountInvoker) Invoke() (*model.UpdatePartitionCountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePartitionCountResponse), nil
	}
}

type UpdateStreamInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStreamInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateStreamInvoker) Invoke() (*model.UpdateStreamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStreamResponse), nil
	}
}

type CreateAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAppInvoker) Invoke() (*model.CreateAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppResponse), nil
	}
}

type DeleteAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAppInvoker) Invoke() (*model.DeleteAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppResponse), nil
	}
}

type ListAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppInvoker) Invoke() (*model.ListAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppResponse), nil
	}
}

type ShowAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppInvoker) Invoke() (*model.ShowAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppResponse), nil
	}
}

type ShowConsumerStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConsumerStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConsumerStateInvoker) Invoke() (*model.ShowConsumerStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConsumerStateResponse), nil
	}
}

type CommitCheckpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *CommitCheckpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CommitCheckpointInvoker) Invoke() (*model.CommitCheckpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommitCheckpointResponse), nil
	}
}

type DeleteCheckpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCheckpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCheckpointInvoker) Invoke() (*model.DeleteCheckpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCheckpointResponse), nil
	}
}

type ShowCheckpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCheckpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCheckpointInvoker) Invoke() (*model.ShowCheckpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCheckpointResponse), nil
	}
}
