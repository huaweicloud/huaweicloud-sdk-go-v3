package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dis/v2/model"
)

type DisClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDisClient(hcClient *httpclient.HcHttpClient) *DisClient {
	return &DisClient{HcClient: hcClient}
}

func DisClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BatchCreateTags 批量添加资源标签
//
// 该接口用于批量添加资源（通道等）标签。此接口为幂等接口：创建时如果请求体中存在重复key则报错。创建时，不允许设置重复key数据，如果数据库已存在该key，就覆盖value的值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) BatchCreateTags(request *model.BatchCreateTagsRequest) (*model.BatchCreateTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateTagsResponse), nil
	}
}

// BatchCreateTagsInvoker 批量添加资源标签
func (c *DisClient) BatchCreateTagsInvoker(request *model.BatchCreateTagsRequest) *BatchCreateTagsInvoker {
	requestDef := GenReqDefForBatchCreateTags()
	return &BatchCreateTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteTags 批量删除资源标签
//
// 该接口用于批量删除资源（通道等）标签。此接口为幂等接口：删除时，如果删除的标签不存在，默认处理成功；删除时不对标签字符集范围做校验。删除时tags结构体不能缺失，key不能为空，或者空字符串。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) BatchDeleteTags(request *model.BatchDeleteTagsRequest) (*model.BatchDeleteTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteTagsResponse), nil
	}
}

// BatchDeleteTagsInvoker 批量删除资源标签
func (c *DisClient) BatchDeleteTagsInvoker(request *model.BatchDeleteTagsRequest) *BatchDeleteTagsInvoker {
	requestDef := GenReqDefForBatchDeleteTags()
	return &BatchDeleteTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchStartTransferTask 批量启动转储任务
//
// 此接口用于批量启动转储任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) BatchStartTransferTask(request *model.BatchStartTransferTaskRequest) (*model.BatchStartTransferTaskResponse, error) {
	requestDef := GenReqDefForBatchStartTransferTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStartTransferTaskResponse), nil
	}
}

// BatchStartTransferTaskInvoker 批量启动转储任务
func (c *DisClient) BatchStartTransferTaskInvoker(request *model.BatchStartTransferTaskRequest) *BatchStartTransferTaskInvoker {
	requestDef := GenReqDefForBatchStartTransferTask()
	return &BatchStartTransferTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchStopTransferTask 批量暂停转储任务
//
// 此接口用于批量暂停转储任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) BatchStopTransferTask(request *model.BatchStopTransferTaskRequest) (*model.BatchStopTransferTaskResponse, error) {
	requestDef := GenReqDefForBatchStopTransferTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStopTransferTaskResponse), nil
	}
}

// BatchStopTransferTaskInvoker 批量暂停转储任务
func (c *DisClient) BatchStopTransferTaskInvoker(request *model.BatchStopTransferTaskRequest) *BatchStopTransferTaskInvoker {
	requestDef := GenReqDefForBatchStopTransferTask()
	return &BatchStopTransferTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ConsumeRecords 下载数据
//
// 本接口用于从DIS通道中下载数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) ConsumeRecords(request *model.ConsumeRecordsRequest) (*model.ConsumeRecordsResponse, error) {
	requestDef := GenReqDefForConsumeRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ConsumeRecordsResponse), nil
	}
}

// ConsumeRecordsInvoker 下载数据
func (c *DisClient) ConsumeRecordsInvoker(request *model.ConsumeRecordsRequest) *ConsumeRecordsInvoker {
	requestDef := GenReqDefForConsumeRecords()
	return &ConsumeRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateObsTransferTask 添加OBS转储任务
//
// 本接口用于添加OBS转储任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) CreateObsTransferTask(request *model.CreateObsTransferTaskRequest) (*model.CreateObsTransferTaskResponse, error) {
	requestDef := GenReqDefForCreateObsTransferTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateObsTransferTaskResponse), nil
	}
}

// CreateObsTransferTaskInvoker 添加OBS转储任务
func (c *DisClient) CreateObsTransferTaskInvoker(request *model.CreateObsTransferTaskRequest) *CreateObsTransferTaskInvoker {
	requestDef := GenReqDefForCreateObsTransferTask()
	return &CreateObsTransferTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStream 创建通道
//
// 本接口用于创建通道。
//
// - 创建通道时，需指定通道类型（普通、高级）、分区数量。
// - 一个账号默认最多可以创建10个高级通道分区和50个普通通道分区，可提交工单增加配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) CreateStream(request *model.CreateStreamRequest) (*model.CreateStreamResponse, error) {
	requestDef := GenReqDefForCreateStream()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStreamResponse), nil
	}
}

// CreateStreamInvoker 创建通道
func (c *DisClient) CreateStreamInvoker(request *model.CreateStreamRequest) *CreateStreamInvoker {
	requestDef := GenReqDefForCreateStream()
	return &CreateStreamInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTag 给指定通道添加标签
//
// 本接口用于给指定通道添加标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) CreateTag(request *model.CreateTagRequest) (*model.CreateTagResponse, error) {
	requestDef := GenReqDefForCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagResponse), nil
	}
}

// CreateTagInvoker 给指定通道添加标签
func (c *DisClient) CreateTagInvoker(request *model.CreateTagRequest) *CreateTagInvoker {
	requestDef := GenReqDefForCreateTag()
	return &CreateTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteStream 删除指定通道
//
// 本接口用于删除指定通道。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) DeleteStream(request *model.DeleteStreamRequest) (*model.DeleteStreamResponse, error) {
	requestDef := GenReqDefForDeleteStream()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStreamResponse), nil
	}
}

// DeleteStreamInvoker 删除指定通道
func (c *DisClient) DeleteStreamInvoker(request *model.DeleteStreamRequest) *DeleteStreamInvoker {
	requestDef := GenReqDefForDeleteStream()
	return &DeleteStreamInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTag 删除指定通道的标签
//
// 该接口用于删除指定通道的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) DeleteTag(request *model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	requestDef := GenReqDefForDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResponse), nil
	}
}

// DeleteTagInvoker 删除指定通道的标签
func (c *DisClient) DeleteTagInvoker(request *model.DeleteTagRequest) *DeleteTagInvoker {
	requestDef := GenReqDefForDeleteTag()
	return &DeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTransferTask 删除转储任务
//
// 该接口用于删除转储任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) DeleteTransferTask(request *model.DeleteTransferTaskRequest) (*model.DeleteTransferTaskResponse, error) {
	requestDef := GenReqDefForDeleteTransferTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTransferTaskResponse), nil
	}
}

// DeleteTransferTaskInvoker 删除转储任务
func (c *DisClient) DeleteTransferTaskInvoker(request *model.DeleteTransferTaskRequest) *DeleteTransferTaskInvoker {
	requestDef := GenReqDefForDeleteTransferTask()
	return &DeleteTransferTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicies 查询权限策略列表
//
// 本接口用于查询指定通道的权限策略列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) ListPolicies(request *model.ListPoliciesRequest) (*model.ListPoliciesResponse, error) {
	requestDef := GenReqDefForListPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPoliciesResponse), nil
	}
}

// ListPoliciesInvoker 查询权限策略列表
func (c *DisClient) ListPoliciesInvoker(request *model.ListPoliciesRequest) *ListPoliciesInvoker {
	requestDef := GenReqDefForListPolicies()
	return &ListPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourcesByTags 使用标签过滤资源（通道等）
//
// 该接口用于使用标签过滤资源（通道等）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) ListResourcesByTags(request *model.ListResourcesByTagsRequest) (*model.ListResourcesByTagsResponse, error) {
	requestDef := GenReqDefForListResourcesByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourcesByTagsResponse), nil
	}
}

// ListResourcesByTagsInvoker 使用标签过滤资源（通道等）
func (c *DisClient) ListResourcesByTagsInvoker(request *model.ListResourcesByTagsRequest) *ListResourcesByTagsInvoker {
	requestDef := GenReqDefForListResourcesByTags()
	return &ListResourcesByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStreams 查询通道列表
//
// 本接口用户查询当前租户创建的所有通道。
//
// 查询时，需要指定从哪个通道开始返回通道列表和单次请求需要返回的最大数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) ListStreams(request *model.ListStreamsRequest) (*model.ListStreamsResponse, error) {
	requestDef := GenReqDefForListStreams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStreamsResponse), nil
	}
}

// ListStreamsInvoker 查询通道列表
func (c *DisClient) ListStreamsInvoker(request *model.ListStreamsRequest) *ListStreamsInvoker {
	requestDef := GenReqDefForListStreams()
	return &ListStreamsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTags 查询指定区域所有标签集合
//
// 该接口用于查询指定区域所有标签集合。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

// ListTagsInvoker 查询指定区域所有标签集合
func (c *DisClient) ListTagsInvoker(request *model.ListTagsRequest) *ListTagsInvoker {
	requestDef := GenReqDefForListTags()
	return &ListTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTransferTasks 查询转储任务列表
//
// 本接口用于查询转储任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) ListTransferTasks(request *model.ListTransferTasksRequest) (*model.ListTransferTasksResponse, error) {
	requestDef := GenReqDefForListTransferTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTransferTasksResponse), nil
	}
}

// ListTransferTasksInvoker 查询转储任务列表
func (c *DisClient) ListTransferTasksInvoker(request *model.ListTransferTasksRequest) *ListTransferTasksInvoker {
	requestDef := GenReqDefForListTransferTasks()
	return &ListTransferTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendRecords 上传数据
//
// 本接口用于上传数据到DIS通道中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) SendRecords(request *model.SendRecordsRequest) (*model.SendRecordsResponse, error) {
	requestDef := GenReqDefForSendRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendRecordsResponse), nil
	}
}

// SendRecordsInvoker 上传数据
func (c *DisClient) SendRecordsInvoker(request *model.SendRecordsRequest) *SendRecordsInvoker {
	requestDef := GenReqDefForSendRecords()
	return &SendRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCursor 获取数据游标
//
// 本接口用于获取数据游标。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) ShowCursor(request *model.ShowCursorRequest) (*model.ShowCursorResponse, error) {
	requestDef := GenReqDefForShowCursor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCursorResponse), nil
	}
}

// ShowCursorInvoker 获取数据游标
func (c *DisClient) ShowCursorInvoker(request *model.ShowCursorRequest) *ShowCursorInvoker {
	requestDef := GenReqDefForShowCursor()
	return &ShowCursorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPartitionMetrics 查询分区监控
//
// 本接口用于查询通道指定分区的监控数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) ShowPartitionMetrics(request *model.ShowPartitionMetricsRequest) (*model.ShowPartitionMetricsResponse, error) {
	requestDef := GenReqDefForShowPartitionMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPartitionMetricsResponse), nil
	}
}

// ShowPartitionMetricsInvoker 查询分区监控
func (c *DisClient) ShowPartitionMetricsInvoker(request *model.ShowPartitionMetricsRequest) *ShowPartitionMetricsInvoker {
	requestDef := GenReqDefForShowPartitionMetrics()
	return &ShowPartitionMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStream 查看通道详情
//
// 本接口用于查询指定通道的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) ShowStream(request *model.ShowStreamRequest) (*model.ShowStreamResponse, error) {
	requestDef := GenReqDefForShowStream()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStreamResponse), nil
	}
}

// ShowStreamInvoker 查看通道详情
func (c *DisClient) ShowStreamInvoker(request *model.ShowStreamRequest) *ShowStreamInvoker {
	requestDef := GenReqDefForShowStream()
	return &ShowStreamInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStreamMetrics 查询通道监控
//
// 本接口用于查询指定通道的监控数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) ShowStreamMetrics(request *model.ShowStreamMetricsRequest) (*model.ShowStreamMetricsResponse, error) {
	requestDef := GenReqDefForShowStreamMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStreamMetricsResponse), nil
	}
}

// ShowStreamMetricsInvoker 查询通道监控
func (c *DisClient) ShowStreamMetricsInvoker(request *model.ShowStreamMetricsRequest) *ShowStreamMetricsInvoker {
	requestDef := GenReqDefForShowStreamMetrics()
	return &ShowStreamMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStreamTags 查询指定通道的标签信息
//
// 该接口用于查询指定通道的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) ShowStreamTags(request *model.ShowStreamTagsRequest) (*model.ShowStreamTagsResponse, error) {
	requestDef := GenReqDefForShowStreamTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStreamTagsResponse), nil
	}
}

// ShowStreamTagsInvoker 查询指定通道的标签信息
func (c *DisClient) ShowStreamTagsInvoker(request *model.ShowStreamTagsRequest) *ShowStreamTagsInvoker {
	requestDef := GenReqDefForShowStreamTags()
	return &ShowStreamTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTransferTask 查询转储任务详情
//
// 查询转储任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) ShowTransferTask(request *model.ShowTransferTaskRequest) (*model.ShowTransferTaskResponse, error) {
	requestDef := GenReqDefForShowTransferTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTransferTaskResponse), nil
	}
}

// ShowTransferTaskInvoker 查询转储任务详情
func (c *DisClient) ShowTransferTaskInvoker(request *model.ShowTransferTaskRequest) *ShowTransferTaskInvoker {
	requestDef := GenReqDefForShowTransferTask()
	return &ShowTransferTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePartitionCount 修改分区数量
//
// 本接口用于变更指定通道中的分区数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) UpdatePartitionCount(request *model.UpdatePartitionCountRequest) (*model.UpdatePartitionCountResponse, error) {
	requestDef := GenReqDefForUpdatePartitionCount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePartitionCountResponse), nil
	}
}

// UpdatePartitionCountInvoker 修改分区数量
func (c *DisClient) UpdatePartitionCountInvoker(request *model.UpdatePartitionCountRequest) *UpdatePartitionCountInvoker {
	requestDef := GenReqDefForUpdatePartitionCount()
	return &UpdatePartitionCountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateStream 更新通道信息
//
// 本接口用于更新指定通道的通道信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) UpdateStream(request *model.UpdateStreamRequest) (*model.UpdateStreamResponse, error) {
	requestDef := GenReqDefForUpdateStream()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStreamResponse), nil
	}
}

// UpdateStreamInvoker 更新通道信息
func (c *DisClient) UpdateStreamInvoker(request *model.UpdateStreamRequest) *UpdateStreamInvoker {
	requestDef := GenReqDefForUpdateStream()
	return &UpdateStreamInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApp 创建消费App
//
// 本接口用于创建消费APP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) CreateApp(request *model.CreateAppRequest) (*model.CreateAppResponse, error) {
	requestDef := GenReqDefForCreateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppResponse), nil
	}
}

// CreateAppInvoker 创建消费App
func (c *DisClient) CreateAppInvoker(request *model.CreateAppRequest) *CreateAppInvoker {
	requestDef := GenReqDefForCreateApp()
	return &CreateAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApp 删除App
//
// 本接口用于删除App。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) DeleteApp(request *model.DeleteAppRequest) (*model.DeleteAppResponse, error) {
	requestDef := GenReqDefForDeleteApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppResponse), nil
	}
}

// DeleteAppInvoker 删除App
func (c *DisClient) DeleteAppInvoker(request *model.DeleteAppRequest) *DeleteAppInvoker {
	requestDef := GenReqDefForDeleteApp()
	return &DeleteAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApp 查询App列表
//
// 本接口用于查询APP列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) ListApp(request *model.ListAppRequest) (*model.ListAppResponse, error) {
	requestDef := GenReqDefForListApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppResponse), nil
	}
}

// ListAppInvoker 查询App列表
func (c *DisClient) ListAppInvoker(request *model.ListAppRequest) *ListAppInvoker {
	requestDef := GenReqDefForListApp()
	return &ListAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApp 查看App详情
//
// 本接口用于查询APP详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) ShowApp(request *model.ShowAppRequest) (*model.ShowAppResponse, error) {
	requestDef := GenReqDefForShowApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppResponse), nil
	}
}

// ShowAppInvoker 查看App详情
func (c *DisClient) ShowAppInvoker(request *model.ShowAppRequest) *ShowAppInvoker {
	requestDef := GenReqDefForShowApp()
	return &ShowAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConsumerState 查看App消费状态
//
// 本接口用于查询APP消费状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) ShowConsumerState(request *model.ShowConsumerStateRequest) (*model.ShowConsumerStateResponse, error) {
	requestDef := GenReqDefForShowConsumerState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConsumerStateResponse), nil
	}
}

// ShowConsumerStateInvoker 查看App消费状态
func (c *DisClient) ShowConsumerStateInvoker(request *model.ShowConsumerStateRequest) *ShowConsumerStateInvoker {
	requestDef := GenReqDefForShowConsumerState()
	return &ShowConsumerStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CommitCheckpoint 提交Checkpoint
//
// 本接口用于提交Checkpoint。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) CommitCheckpoint(request *model.CommitCheckpointRequest) (*model.CommitCheckpointResponse, error) {
	requestDef := GenReqDefForCommitCheckpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CommitCheckpointResponse), nil
	}
}

// CommitCheckpointInvoker 提交Checkpoint
func (c *DisClient) CommitCheckpointInvoker(request *model.CommitCheckpointRequest) *CommitCheckpointInvoker {
	requestDef := GenReqDefForCommitCheckpoint()
	return &CommitCheckpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCheckpoint 删除Checkpoint
//
// 本接口用于删除Checkpoint。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) DeleteCheckpoint(request *model.DeleteCheckpointRequest) (*model.DeleteCheckpointResponse, error) {
	requestDef := GenReqDefForDeleteCheckpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCheckpointResponse), nil
	}
}

// DeleteCheckpointInvoker 删除Checkpoint
func (c *DisClient) DeleteCheckpointInvoker(request *model.DeleteCheckpointRequest) *DeleteCheckpointInvoker {
	requestDef := GenReqDefForDeleteCheckpoint()
	return &DeleteCheckpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCheckpoint 查询Checkpoint详情
//
// 本接口用于查询Checkpoint详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DisClient) ShowCheckpoint(request *model.ShowCheckpointRequest) (*model.ShowCheckpointResponse, error) {
	requestDef := GenReqDefForShowCheckpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCheckpointResponse), nil
	}
}

// ShowCheckpointInvoker 查询Checkpoint详情
func (c *DisClient) ShowCheckpointInvoker(request *model.ShowCheckpointRequest) *ShowCheckpointInvoker {
	requestDef := GenReqDefForShowCheckpoint()
	return &ShowCheckpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
