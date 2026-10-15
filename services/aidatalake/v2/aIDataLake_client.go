package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/aidatalake/v2/model"
)

type AIDataLakeClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewAIDataLakeClient(hcClient *httpclient.HcHttpClient) *AIDataLakeClient {
	return &AIDataLakeClient{HcClient: hcClient}
}

func AIDataLakeClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CancelAuraDataframeJob 取消Aura Dataframe Job
//
// 取消正在运行中的Dataframe Job。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CancelAuraDataframeJob(request *model.CancelAuraDataframeJobRequest) (*model.CancelAuraDataframeJobResponse, error) {
	requestDef := GenReqDefForCancelAuraDataframeJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelAuraDataframeJobResponse), nil
	}
}

// CancelAuraDataframeJobInvoker 取消Aura Dataframe Job
func (c *AIDataLakeClient) CancelAuraDataframeJobInvoker(request *model.CancelAuraDataframeJobRequest) *CancelAuraDataframeJobInvoker {
	requestDef := GenReqDefForCancelAuraDataframeJob()
	return &CancelAuraDataframeJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAuraDataframeJob 创建Aura Dataframe Job
//
// 创建Dataframe Job。
// 用户通过此接口在指定端点创建Dataframe Job。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CreateAuraDataframeJob(request *model.CreateAuraDataframeJobRequest) (*model.CreateAuraDataframeJobResponse, error) {
	requestDef := GenReqDefForCreateAuraDataframeJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuraDataframeJobResponse), nil
	}
}

// CreateAuraDataframeJobInvoker 创建Aura Dataframe Job
func (c *AIDataLakeClient) CreateAuraDataframeJobInvoker(request *model.CreateAuraDataframeJobRequest) *CreateAuraDataframeJobInvoker {
	requestDef := GenReqDefForCreateAuraDataframeJob()
	return &CreateAuraDataframeJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuraDataframeJob 查看Aura Dataframe Job列表
//
// 查询Aura Dataframe Job列表。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListAuraDataframeJob(request *model.ListAuraDataframeJobRequest) (*model.ListAuraDataframeJobResponse, error) {
	requestDef := GenReqDefForListAuraDataframeJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuraDataframeJobResponse), nil
	}
}

// ListAuraDataframeJobInvoker 查看Aura Dataframe Job列表
func (c *AIDataLakeClient) ListAuraDataframeJobInvoker(request *model.ListAuraDataframeJobRequest) *ListAuraDataframeJobInvoker {
	requestDef := GenReqDefForListAuraDataframeJob()
	return &ListAuraDataframeJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuraDataframeJob 查看Aura Dataframe Job详情
//
// 查询Aura Dataframe Job信息。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowAuraDataframeJob(request *model.ShowAuraDataframeJobRequest) (*model.ShowAuraDataframeJobResponse, error) {
	requestDef := GenReqDefForShowAuraDataframeJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuraDataframeJobResponse), nil
	}
}

// ShowAuraDataframeJobInvoker 查看Aura Dataframe Job详情
func (c *AIDataLakeClient) ShowAuraDataframeJobInvoker(request *model.ShowAuraDataframeJobRequest) *ShowAuraDataframeJobInvoker {
	requestDef := GenReqDefForShowAuraDataframeJob()
	return &ShowAuraDataframeJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelAuraSqlStatement 取消SQL执行
//
// 取消SQL语句执行。 用户可通过该接口取消SQL语句执行，输入为SQL Session id、statement id，返回操作结果。 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CancelAuraSqlStatement(request *model.CancelAuraSqlStatementRequest) (*model.CancelAuraSqlStatementResponse, error) {
	requestDef := GenReqDefForCancelAuraSqlStatement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelAuraSqlStatementResponse), nil
	}
}

// CancelAuraSqlStatementInvoker 取消SQL执行
func (c *AIDataLakeClient) CancelAuraSqlStatementInvoker(request *model.CancelAuraSqlStatementRequest) *CancelAuraSqlStatementInvoker {
	requestDef := GenReqDefForCancelAuraSqlStatement()
	return &CancelAuraSqlStatementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CloseAuraSqlSession 关闭Session
//
// 关闭SQL Session。
// 用户可通过此接口关闭SQL Session，输入为Session id，返回操作结果。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CloseAuraSqlSession(request *model.CloseAuraSqlSessionRequest) (*model.CloseAuraSqlSessionResponse, error) {
	requestDef := GenReqDefForCloseAuraSqlSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CloseAuraSqlSessionResponse), nil
	}
}

// CloseAuraSqlSessionInvoker 关闭Session
func (c *AIDataLakeClient) CloseAuraSqlSessionInvoker(request *model.CloseAuraSqlSessionRequest) *CloseAuraSqlSessionInvoker {
	requestDef := GenReqDefForCloseAuraSqlSession()
	return &CloseAuraSqlSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAuraSqlSession 创建Session
//
// 创建SQL Session。
// 用户通过此接口在指定端点创建SQL Session，通过输入端点id、LakeFormation配置，返回Session信息。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CreateAuraSqlSession(request *model.CreateAuraSqlSessionRequest) (*model.CreateAuraSqlSessionResponse, error) {
	requestDef := GenReqDefForCreateAuraSqlSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuraSqlSessionResponse), nil
	}
}

// CreateAuraSqlSessionInvoker 创建Session
func (c *AIDataLakeClient) CreateAuraSqlSessionInvoker(request *model.CreateAuraSqlSessionRequest) *CreateAuraSqlSessionInvoker {
	requestDef := GenReqDefForCreateAuraSqlSession()
	return &CreateAuraSqlSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteAuraSqlStatement 在指定Session中执行SQL
//
// 在指定的Session下执行SQL语句。
// 用户可通过此接口执行SQL语句，输入为SQL语句、SQL Session id、绑定参数等信息，返回语句执行结果或statement id。
// 此接口支持异步和同步两种执行模式，由入参is_sync参数决定。如果是异步执行，配套使用[查看SQL执行结果](ShowAuraV2SqlStatementResult.xml)接口查询语句结果。
// 如果是同步执行，接口会等待一段时间接收结果，如果执行完毕会返回执行结果，如果没有执行完毕，会返回一个运行中状态，后续可以使用[查看SQL执行结果](ShowAuraV2SqlStatementResult.xml)接口查询语句结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ExecuteAuraSqlStatement(request *model.ExecuteAuraSqlStatementRequest) (*model.ExecuteAuraSqlStatementResponse, error) {
	requestDef := GenReqDefForExecuteAuraSqlStatement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteAuraSqlStatementResponse), nil
	}
}

// ExecuteAuraSqlStatementInvoker 在指定Session中执行SQL
func (c *AIDataLakeClient) ExecuteAuraSqlStatementInvoker(request *model.ExecuteAuraSqlStatementRequest) *ExecuteAuraSqlStatementInvoker {
	requestDef := GenReqDefForExecuteAuraSqlStatement()
	return &ExecuteAuraSqlStatementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuraSessionStatementRecords 查看指定Session下的SQL执行记录
//
// 查询指定Session下的SQL执行记录。
// 输入workspace_id，session_id，statement_id（可选），status（可选），分页查询参数limit和marker；输出此会话下SQL执行记录。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListAuraSessionStatementRecords(request *model.ListAuraSessionStatementRecordsRequest) (*model.ListAuraSessionStatementRecordsResponse, error) {
	requestDef := GenReqDefForListAuraSessionStatementRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuraSessionStatementRecordsResponse), nil
	}
}

// ListAuraSessionStatementRecordsInvoker 查看指定Session下的SQL执行记录
func (c *AIDataLakeClient) ListAuraSessionStatementRecordsInvoker(request *model.ListAuraSessionStatementRecordsRequest) *ListAuraSessionStatementRecordsInvoker {
	requestDef := GenReqDefForListAuraSessionStatementRecords()
	return &ListAuraSessionStatementRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuraSqlSessions 查看Session列表
//
// 查询SQL Session列表。
// 输入workspace_id，endpoint_name（可选），status（可选），session_id（可选），start_time（可选），end_time（可选）,分页查询参数limit和marker；输出SQL会话列表，包含会话id、会话创建时间、结束时间、会话状态等信息。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListAuraSqlSessions(request *model.ListAuraSqlSessionsRequest) (*model.ListAuraSqlSessionsResponse, error) {
	requestDef := GenReqDefForListAuraSqlSessions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuraSqlSessionsResponse), nil
	}
}

// ListAuraSqlSessionsInvoker 查看Session列表
func (c *AIDataLakeClient) ListAuraSqlSessionsInvoker(request *model.ListAuraSqlSessionsRequest) *ListAuraSqlSessionsInvoker {
	requestDef := GenReqDefForListAuraSqlSessions()
	return &ListAuraSqlSessionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuraStatementRecords 查看SQL执行记录
//
// 查询SQL执行记录。 查询指定Session下的SQL执行记录。 输入workspace_id，session_id，statement_id（可选），status（可选），分页查询参数limit和marker；输出此会话下SQL执行记录。 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListAuraStatementRecords(request *model.ListAuraStatementRecordsRequest) (*model.ListAuraStatementRecordsResponse, error) {
	requestDef := GenReqDefForListAuraStatementRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuraStatementRecordsResponse), nil
	}
}

// ListAuraStatementRecordsInvoker 查看SQL执行记录
func (c *AIDataLakeClient) ListAuraStatementRecordsInvoker(request *model.ListAuraStatementRecordsRequest) *ListAuraStatementRecordsInvoker {
	requestDef := GenReqDefForListAuraStatementRecords()
	return &ListAuraStatementRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuraSqlSession 查看Session详情
//
// 查询SQL Session信息。
// 用户可通过此接口查询SQL Session信息，输入为Session id，返回操作结果。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowAuraSqlSession(request *model.ShowAuraSqlSessionRequest) (*model.ShowAuraSqlSessionResponse, error) {
	requestDef := GenReqDefForShowAuraSqlSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuraSqlSessionResponse), nil
	}
}

// ShowAuraSqlSessionInvoker 查看Session详情
func (c *AIDataLakeClient) ShowAuraSqlSessionInvoker(request *model.ShowAuraSqlSessionRequest) *ShowAuraSqlSessionInvoker {
	requestDef := GenReqDefForShowAuraSqlSession()
	return &ShowAuraSqlSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuraSqlStatementResult 查看SQL执行结果。
//
// 查询SQL语句执行结果。
// 用户可通过该接口查询SQL语句执行结果，输入为SQL Session id、statement id，返回执行结果。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowAuraSqlStatementResult(request *model.ShowAuraSqlStatementResultRequest) (*model.ShowAuraSqlStatementResultResponse, error) {
	requestDef := GenReqDefForShowAuraSqlStatementResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuraSqlStatementResultResponse), nil
	}
}

// ShowAuraSqlStatementResultInvoker 查看SQL执行结果。
func (c *AIDataLakeClient) ShowAuraSqlStatementResultInvoker(request *model.ShowAuraSqlStatementResultRequest) *ShowAuraSqlStatementResultInvoker {
	requestDef := GenReqDefForShowAuraSqlStatementResult()
	return &ShowAuraSqlStatementResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuraStatementOperatorMetrics 查看SQL执行算子监控记录
//
// 查看SQL执行算子监控记录。 用户可通过该接口查看SQL执行算子监控记录，输入为SQL Session id、statement id，返回操作结果。 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowAuraStatementOperatorMetrics(request *model.ShowAuraStatementOperatorMetricsRequest) (*model.ShowAuraStatementOperatorMetricsResponse, error) {
	requestDef := GenReqDefForShowAuraStatementOperatorMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuraStatementOperatorMetricsResponse), nil
	}
}

// ShowAuraStatementOperatorMetricsInvoker 查看SQL执行算子监控记录
func (c *AIDataLakeClient) ShowAuraStatementOperatorMetricsInvoker(request *model.ShowAuraStatementOperatorMetricsRequest) *ShowAuraStatementOperatorMetricsInvoker {
	requestDef := GenReqDefForShowAuraStatementOperatorMetrics()
	return &ShowAuraStatementOperatorMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuraStatementQueryMetrics 查看语句监控详情
//
// 查询语句监控详情信息。
// 输入workspace_id，session_id，statement_id，输出此会话下SQL监控数据。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowAuraStatementQueryMetrics(request *model.ShowAuraStatementQueryMetricsRequest) (*model.ShowAuraStatementQueryMetricsResponse, error) {
	requestDef := GenReqDefForShowAuraStatementQueryMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuraStatementQueryMetricsResponse), nil
	}
}

// ShowAuraStatementQueryMetricsInvoker 查看语句监控详情
func (c *AIDataLakeClient) ShowAuraStatementQueryMetricsInvoker(request *model.ShowAuraStatementQueryMetricsRequest) *ShowAuraStatementQueryMetricsInvoker {
	requestDef := GenReqDefForShowAuraStatementQueryMetrics()
	return &ShowAuraStatementQueryMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlinkJobAutoscalingHistories 查询Flink作业自动调优历史
//
// 查询Flink作业自动调优历史记录，包含每次扩缩容操作的详细信息。此接口为同步接口，支持分页查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListFlinkJobAutoscalingHistories(request *model.ListFlinkJobAutoscalingHistoriesRequest) (*model.ListFlinkJobAutoscalingHistoriesResponse, error) {
	requestDef := GenReqDefForListFlinkJobAutoscalingHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlinkJobAutoscalingHistoriesResponse), nil
	}
}

// ListFlinkJobAutoscalingHistoriesInvoker 查询Flink作业自动调优历史
func (c *AIDataLakeClient) ListFlinkJobAutoscalingHistoriesInvoker(request *model.ListFlinkJobAutoscalingHistoriesRequest) *ListFlinkJobAutoscalingHistoriesInvoker {
	requestDef := GenReqDefForListFlinkJobAutoscalingHistories()
	return &ListFlinkJobAutoscalingHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlinkJobAutoscalingPolicy 查询Flink作业自动调优配置
//
// 查询Flink作业当前的自动调优配置。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowFlinkJobAutoscalingPolicy(request *model.ShowFlinkJobAutoscalingPolicyRequest) (*model.ShowFlinkJobAutoscalingPolicyResponse, error) {
	requestDef := GenReqDefForShowFlinkJobAutoscalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlinkJobAutoscalingPolicyResponse), nil
	}
}

// ShowFlinkJobAutoscalingPolicyInvoker 查询Flink作业自动调优配置
func (c *AIDataLakeClient) ShowFlinkJobAutoscalingPolicyInvoker(request *model.ShowFlinkJobAutoscalingPolicyRequest) *ShowFlinkJobAutoscalingPolicyInvoker {
	requestDef := GenReqDefForShowFlinkJobAutoscalingPolicy()
	return &ShowFlinkJobAutoscalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFlinkJobAutoscalingPolicy 热更新Flink作业自动调优配置
//
// 热更新Flink作业自动调优配置。此接口为异步接口，作业运行中时配置将自动应用，无需重启作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) UpdateFlinkJobAutoscalingPolicy(request *model.UpdateFlinkJobAutoscalingPolicyRequest) (*model.UpdateFlinkJobAutoscalingPolicyResponse, error) {
	requestDef := GenReqDefForUpdateFlinkJobAutoscalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFlinkJobAutoscalingPolicyResponse), nil
	}
}

// UpdateFlinkJobAutoscalingPolicyInvoker 热更新Flink作业自动调优配置
func (c *AIDataLakeClient) UpdateFlinkJobAutoscalingPolicyInvoker(request *model.UpdateFlinkJobAutoscalingPolicyRequest) *UpdateFlinkJobAutoscalingPolicyInvoker {
	requestDef := GenReqDefForUpdateFlinkJobAutoscalingPolicy()
	return &UpdateFlinkJobAutoscalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelFlinkJob 取消Flink作业执行
//
// 取消Flink作业。此接口为异步接口，响应仅表示取消请求已接收，作业实际停止状态可通过[查询Flink作业状态](ShowFlinkJobState.xml)接口查询。
// 支持多种停止策略：
// - NONE：立即停止不保存状态。
// - STOP_WITH_SAVEPOINT：停止前创建快照。
// - STOP_WITH_DRAIN：Drain模式停止，触发窗口计算后再创建快照停止。
// - FORCE_CANCEL：强制取消，仅当作业已处于CANCELLING状态时可再次触发取消。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CancelFlinkJob(request *model.CancelFlinkJobRequest) (*model.CancelFlinkJobResponse, error) {
	requestDef := GenReqDefForCancelFlinkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelFlinkJobResponse), nil
	}
}

// CancelFlinkJobInvoker 取消Flink作业执行
func (c *AIDataLakeClient) CancelFlinkJobInvoker(request *model.CancelFlinkJobRequest) *CancelFlinkJobInvoker {
	requestDef := GenReqDefForCancelFlinkJob()
	return &CancelFlinkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlinkJobEvents 查询Flink作业事件详情列表
//
// 查询Flink作业的运行事件列表。此接口为同步接口，响应中包含作业运行期间产生的事件详情，可用于问题定位和运行状态追踪。
// 可通过[查看Flink作业详情](ShowFlinkJob.xml)接口获取作业完整信息，或通过[查询Flink作业状态](ShowFlinkJobState.xml)接口获取作业当前运行状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListFlinkJobEvents(request *model.ListFlinkJobEventsRequest) (*model.ListFlinkJobEventsResponse, error) {
	requestDef := GenReqDefForListFlinkJobEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlinkJobEventsResponse), nil
	}
}

// ListFlinkJobEventsInvoker 查询Flink作业事件详情列表
func (c *AIDataLakeClient) ListFlinkJobEventsInvoker(request *model.ListFlinkJobEventsRequest) *ListFlinkJobEventsInvoker {
	requestDef := GenReqDefForListFlinkJobEvents()
	return &ListFlinkJobEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlinkJobs 查询Flink作业列表
//
// 查询工作空间下Flink作业列表。此接口为同步接口，响应中包含作业的概要信息，支持分页查询和多条件筛选。
// 可通过[查看Flink作业详情](ShowFlinkJob.xml)接口获取指定作业的完整信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListFlinkJobs(request *model.ListFlinkJobsRequest) (*model.ListFlinkJobsResponse, error) {
	requestDef := GenReqDefForListFlinkJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlinkJobsResponse), nil
	}
}

// ListFlinkJobsInvoker 查询Flink作业列表
func (c *AIDataLakeClient) ListFlinkJobsInvoker(request *model.ListFlinkJobsRequest) *ListFlinkJobsInvoker {
	requestDef := GenReqDefForListFlinkJobs()
	return &ListFlinkJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunFlinkJob 创建Flink作业
//
// 创建并启动Flink作业实例。此接口为异步接口，响应仅表示作业创建请求已接收，作业实际运行状态可通过[查询Flink作业状态](ShowFlinkJobState.xml)接口查询。
// 支持多种作业类型：
// - FLINK_SQL_JOB：Flink SQL作业，通过SQL脚本定义数据处理逻辑。
// - FLINK_JAR_JOB：Flink JAR作业，通过自定义JAR包定义数据处理逻辑。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) RunFlinkJob(request *model.RunFlinkJobRequest) (*model.RunFlinkJobResponse, error) {
	requestDef := GenReqDefForRunFlinkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunFlinkJobResponse), nil
	}
}

// RunFlinkJobInvoker 创建Flink作业
func (c *AIDataLakeClient) RunFlinkJobInvoker(request *model.RunFlinkJobRequest) *RunFlinkJobInvoker {
	requestDef := GenReqDefForRunFlinkJob()
	return &RunFlinkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlinkJob 查看Flink作业详情
//
// 获取Flink作业实例的详细信息。此接口为同步接口，响应中包含作业的完整配置和运行状态信息。
// 可通过[查询Flink作业事件详情列表](ListFlinkJobEvents.xml)接口获取作业运行事件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowFlinkJob(request *model.ShowFlinkJobRequest) (*model.ShowFlinkJobResponse, error) {
	requestDef := GenReqDefForShowFlinkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlinkJobResponse), nil
	}
}

// ShowFlinkJobInvoker 查看Flink作业详情
func (c *AIDataLakeClient) ShowFlinkJobInvoker(request *model.ShowFlinkJobRequest) *ShowFlinkJobInvoker {
	requestDef := GenReqDefForShowFlinkJob()
	return &ShowFlinkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlinkJobState 查询Flink作业状态
//
// 查询Flink作业的运行状态。此接口为同步接口，响应中包含作业当前状态。
// 可通过[查看Flink作业详情](ShowFlinkJob.xml)接口获取作业完整信息，或通过[查询Flink作业事件详情列表](ListFlinkJobEvents.xml)接口获取作业运行事件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowFlinkJobState(request *model.ShowFlinkJobStateRequest) (*model.ShowFlinkJobStateResponse, error) {
	requestDef := GenReqDefForShowFlinkJobState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlinkJobStateResponse), nil
	}
}

// ShowFlinkJobStateInvoker 查询Flink作业状态
func (c *AIDataLakeClient) ShowFlinkJobStateInvoker(request *model.ShowFlinkJobStateRequest) *ShowFlinkJobStateInvoker {
	requestDef := GenReqDefForShowFlinkJobState()
	return &ShowFlinkJobStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFlinkSavepoint 创建Flink作业Savepoint
//
// 为运行中的Flink作业触发Savepoint，用于保存作业当前状态。此接口为异步接口，响应仅表示快照创建请求已接收，快照实际生成状态可通过[查询Flink作业Savepoint列表](ListFlinkSavepoint.xml)接口查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CreateFlinkSavepoint(request *model.CreateFlinkSavepointRequest) (*model.CreateFlinkSavepointResponse, error) {
	requestDef := GenReqDefForCreateFlinkSavepoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFlinkSavepointResponse), nil
	}
}

// CreateFlinkSavepointInvoker 创建Flink作业Savepoint
func (c *AIDataLakeClient) CreateFlinkSavepointInvoker(request *model.CreateFlinkSavepointRequest) *CreateFlinkSavepointInvoker {
	requestDef := GenReqDefForCreateFlinkSavepoint()
	return &CreateFlinkSavepointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFlinkSavepoint 删除Flink作业Savepoint
//
// 删除指定Savepoint的元数据记录，不影响OBS中的实际Savepoint数据。该接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) DeleteFlinkSavepoint(request *model.DeleteFlinkSavepointRequest) (*model.DeleteFlinkSavepointResponse, error) {
	requestDef := GenReqDefForDeleteFlinkSavepoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFlinkSavepointResponse), nil
	}
}

// DeleteFlinkSavepointInvoker 删除Flink作业Savepoint
func (c *AIDataLakeClient) DeleteFlinkSavepointInvoker(request *model.DeleteFlinkSavepointRequest) *DeleteFlinkSavepointInvoker {
	requestDef := GenReqDefForDeleteFlinkSavepoint()
	return &DeleteFlinkSavepointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlinkSavepoint 查询Flink作业Savepoint列表
//
// 查询指定Flink作业的Savepoint记录，返回分页列表。此接口为同步接口，响应中包含Savepoint的详细信息。
// 可通过[查询Flink作业Savepoint详情](ShowFlinkSavepoint.xml)接口获取指定Savepoint的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListFlinkSavepoint(request *model.ListFlinkSavepointRequest) (*model.ListFlinkSavepointResponse, error) {
	requestDef := GenReqDefForListFlinkSavepoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlinkSavepointResponse), nil
	}
}

// ListFlinkSavepointInvoker 查询Flink作业Savepoint列表
func (c *AIDataLakeClient) ListFlinkSavepointInvoker(request *model.ListFlinkSavepointRequest) *ListFlinkSavepointInvoker {
	requestDef := GenReqDefForListFlinkSavepoint()
	return &ListFlinkSavepointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlinkSavepoint 查询Flink作业Savepoint详情
//
// 获取指定Savepoint的详细信息。此接口为同步接口，响应中包含Savepoint的完整状态和配置信息。
// 可通过[查询Flink作业Savepoint列表](ListFlinkSavepoint.xml)接口获取Savepoint列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowFlinkSavepoint(request *model.ShowFlinkSavepointRequest) (*model.ShowFlinkSavepointResponse, error) {
	requestDef := GenReqDefForShowFlinkSavepoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlinkSavepointResponse), nil
	}
}

// ShowFlinkSavepointInvoker 查询Flink作业Savepoint详情
func (c *AIDataLakeClient) ShowFlinkSavepointInvoker(request *model.ShowFlinkSavepointRequest) *ShowFlinkSavepointInvoker {
	requestDef := GenReqDefForShowFlinkSavepoint()
	return &ShowFlinkSavepointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFederation 创建Ray联邦端点
//
// 创建Ray联邦端点。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CreateFederation(request *model.CreateFederationRequest) (*model.CreateFederationResponse, error) {
	requestDef := GenReqDefForCreateFederation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFederationResponse), nil
	}
}

// CreateFederationInvoker 创建Ray联邦端点
func (c *AIDataLakeClient) CreateFederationInvoker(request *model.CreateFederationRequest) *CreateFederationInvoker {
	requestDef := GenReqDefForCreateFederation()
	return &CreateFederationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFederation 删除Ray联邦端点
//
// 删除Ray联邦端点。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) DeleteFederation(request *model.DeleteFederationRequest) (*model.DeleteFederationResponse, error) {
	requestDef := GenReqDefForDeleteFederation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFederationResponse), nil
	}
}

// DeleteFederationInvoker 删除Ray联邦端点
func (c *AIDataLakeClient) DeleteFederationInvoker(request *model.DeleteFederationRequest) *DeleteFederationInvoker {
	requestDef := GenReqDefForDeleteFederation()
	return &DeleteFederationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFederations 查询Ray联邦端点列表
//
// 查询工作空间下的Ray联邦端点列表。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListFederations(request *model.ListFederationsRequest) (*model.ListFederationsResponse, error) {
	requestDef := GenReqDefForListFederations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFederationsResponse), nil
	}
}

// ListFederationsInvoker 查询Ray联邦端点列表
func (c *AIDataLakeClient) ListFederationsInvoker(request *model.ListFederationsRequest) *ListFederationsInvoker {
	requestDef := GenReqDefForListFederations()
	return &ListFederationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFederation 查询Ray联邦端点详情
//
// 查询Ray联邦端点详情。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowFederation(request *model.ShowFederationRequest) (*model.ShowFederationResponse, error) {
	requestDef := GenReqDefForShowFederation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFederationResponse), nil
	}
}

// ShowFederationInvoker 查询Ray联邦端点详情
func (c *AIDataLakeClient) ShowFederationInvoker(request *model.ShowFederationRequest) *ShowFederationInvoker {
	requestDef := GenReqDefForShowFederation()
	return &ShowFederationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFederation 更新Ray联邦端点
//
// 更新Ray联邦端点。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) UpdateFederation(request *model.UpdateFederationRequest) (*model.UpdateFederationResponse, error) {
	requestDef := GenReqDefForUpdateFederation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFederationResponse), nil
	}
}

// UpdateFederationInvoker 更新Ray联邦端点
func (c *AIDataLakeClient) UpdateFederationInvoker(request *model.UpdateFederationRequest) *UpdateFederationInvoker {
	requestDef := GenReqDefForUpdateFederation()
	return &UpdateFederationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelRayJob 取消作业运行
//
// 取消作业运行。主要使用于取消运行Ray Job场景。本接口为异步接口，当前取消作业运行请求下发成功后会返回job_id，此时取消作业运行并没有立即完成，需要通过调用[查看Ray作业详情](ShowRayJob.xml)查询Job状态，当Job状态为CANCELED时代表取消作业运行成功。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CancelRayJob(request *model.CancelRayJobRequest) (*model.CancelRayJobResponse, error) {
	requestDef := GenReqDefForCancelRayJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelRayJobResponse), nil
	}
}

// CancelRayJobInvoker 取消作业运行
func (c *AIDataLakeClient) CancelRayJobInvoker(request *model.CancelRayJobRequest) *CancelRayJobInvoker {
	requestDef := GenReqDefForCancelRayJob()
	return &CancelRayJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRayJobs 查询Ray作业列表
//
// 列举所有Ray作业。列举工作空间下的作业，分页返回。本接口为同步接口，当前列举所有Ray作业请求下发后会返回结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListRayJobs(request *model.ListRayJobsRequest) (*model.ListRayJobsResponse, error) {
	requestDef := GenReqDefForListRayJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRayJobsResponse), nil
	}
}

// ListRayJobsInvoker 查询Ray作业列表
func (c *AIDataLakeClient) ListRayJobsInvoker(request *model.ListRayJobsRequest) *ListRayJobsInvoker {
	requestDef := GenReqDefForListRayJobs()
	return &ListRayJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunRayJob 运行作业
//
// 运行Ray作业。本接口为异步接口，当前运行作业请求下发成功后会返回job_id，此时运行作业并没有立即完成，需要通过调用[查看Ray作业详情](ShowRayJob.xml)查询Job状态，当Job状态为SUCCEEDED时代表运行作业成功。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) RunRayJob(request *model.RunRayJobRequest) (*model.RunRayJobResponse, error) {
	requestDef := GenReqDefForRunRayJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunRayJobResponse), nil
	}
}

// RunRayJobInvoker 运行作业
func (c *AIDataLakeClient) RunRayJobInvoker(request *model.RunRayJobRequest) *RunRayJobInvoker {
	requestDef := GenReqDefForRunRayJob()
	return &RunRayJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRayJob 查看Ray作业详情
//
// 查看Ray作业详情。本接口为同步接口，当前查看Ray作业详情请求下发成功后会返回结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowRayJob(request *model.ShowRayJobRequest) (*model.ShowRayJobResponse, error) {
	requestDef := GenReqDefForShowRayJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRayJobResponse), nil
	}
}

// ShowRayJobInvoker 查看Ray作业详情
func (c *AIDataLakeClient) ShowRayJobInvoker(request *model.ShowRayJobRequest) *ShowRayJobInvoker {
	requestDef := GenReqDefForShowRayJob()
	return &ShowRayJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CloseRaySession 关闭交互式会话
//
// 关闭一个已存在的交互式会话，将会话状态设置为成功。本接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CloseRaySession(request *model.CloseRaySessionRequest) (*model.CloseRaySessionResponse, error) {
	requestDef := GenReqDefForCloseRaySession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CloseRaySessionResponse), nil
	}
}

// CloseRaySessionInvoker 关闭交互式会话
func (c *AIDataLakeClient) CloseRaySessionInvoker(request *model.CloseRaySessionRequest) *CloseRaySessionInvoker {
	requestDef := GenReqDefForCloseRaySession()
	return &CloseRaySessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRaySession 创建交互式会话
//
// 创建一个交互式会话。本接口为同步接口，用于Notebook等交互式工具连接Ray集群。联邦用户认证场景下，会返回含联邦用户身份对应的凭据名称。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CreateRaySession(request *model.CreateRaySessionRequest) (*model.CreateRaySessionResponse, error) {
	requestDef := GenReqDefForCreateRaySession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRaySessionResponse), nil
	}
}

// CreateRaySessionInvoker 创建交互式会话
func (c *AIDataLakeClient) CreateRaySessionInvoker(request *model.CreateRaySessionRequest) *CreateRaySessionInvoker {
	requestDef := GenReqDefForCreateRaySession()
	return &CreateRaySessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelSparkJob 取消Spark作业执行
//
// 取消正在执行的Spark作业，此接口为同步接口。调用成功后，作业将被终止执行，直接返回取消结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CancelSparkJob(request *model.CancelSparkJobRequest) (*model.CancelSparkJobResponse, error) {
	requestDef := GenReqDefForCancelSparkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelSparkJobResponse), nil
	}
}

// CancelSparkJobInvoker 取消Spark作业执行
func (c *AIDataLakeClient) CancelSparkJobInvoker(request *model.CancelSparkJobRequest) *CancelSparkJobInvoker {
	requestDef := GenReqDefForCancelSparkJob()
	return &CancelSparkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSparkJobs 查询Spark作业列表
//
// 查询工作空间下Spark作业列表，此接口为同步接口。支持按作业ID、作业名称、作业状态、作业类型、创建时间等条件过滤查询，支持分页查询，调用成功后直接返回作业列表数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListSparkJobs(request *model.ListSparkJobsRequest) (*model.ListSparkJobsResponse, error) {
	requestDef := GenReqDefForListSparkJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSparkJobsResponse), nil
	}
}

// ListSparkJobsInvoker 查询Spark作业列表
func (c *AIDataLakeClient) ListSparkJobsInvoker(request *model.ListSparkJobsRequest) *ListSparkJobsInvoker {
	requestDef := GenReqDefForListSparkJobs()
	return &ListSparkJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunSparkJob 启动Spark作业
//
// 启动Spark作业，此接口为异步接口。支持Spark Jar作业、Python作业和SQL Script作业。调用该接口后，作业将提交到队列等待执行，返回作业ID后需通过查询作业状态接口确认作业是否成功启动。查询作业状态请参见[查询Spark作业的状态](ShowSparkJobState.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) RunSparkJob(request *model.RunSparkJobRequest) (*model.RunSparkJobResponse, error) {
	requestDef := GenReqDefForRunSparkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunSparkJobResponse), nil
	}
}

// RunSparkJobInvoker 启动Spark作业
func (c *AIDataLakeClient) RunSparkJobInvoker(request *model.RunSparkJobRequest) *RunSparkJobInvoker {
	requestDef := GenReqDefForRunSparkJob()
	return &RunSparkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSparkJob 查看Spark作业详情
//
// 查看指定Spark作业的详细信息，此接口为同步接口。包括作业ID、作业名称、作业状态、作业配置、资源配置、镜像配置等完整信息，调用成功后直接返回作业详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowSparkJob(request *model.ShowSparkJobRequest) (*model.ShowSparkJobResponse, error) {
	requestDef := GenReqDefForShowSparkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSparkJobResponse), nil
	}
}

// ShowSparkJobInvoker 查看Spark作业详情
func (c *AIDataLakeClient) ShowSparkJobInvoker(request *model.ShowSparkJobRequest) *ShowSparkJobInvoker {
	requestDef := GenReqDefForShowSparkJob()
	return &ShowSparkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSparkJobState 查询Spark作业的状态
//
// 查询Spark作业的状态，此接口为同步接口。可通过作业ID查询Spark作业的当前执行状态，包括排队中、运行中、已成功、已失败等状态，调用后立即返回作业当前状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowSparkJobState(request *model.ShowSparkJobStateRequest) (*model.ShowSparkJobStateResponse, error) {
	requestDef := GenReqDefForShowSparkJobState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSparkJobStateResponse), nil
	}
}

// ShowSparkJobStateInvoker 查询Spark作业的状态
func (c *AIDataLakeClient) ShowSparkJobStateInvoker(request *model.ShowSparkJobStateRequest) *ShowSparkJobStateInvoker {
	requestDef := GenReqDefForShowSparkJobState()
	return &ShowSparkJobStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSparkOperatorState 查询Spark异步操作状态
//
// 查询Spark异步操作的执行状态，用于获取异步操作（如重启集群等）的当前状态和执行结果。该接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowSparkOperatorState(request *model.ShowSparkOperatorStateRequest) (*model.ShowSparkOperatorStateResponse, error) {
	requestDef := GenReqDefForShowSparkOperatorState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSparkOperatorStateResponse), nil
	}
}

// ShowSparkOperatorStateInvoker 查询Spark异步操作状态
func (c *AIDataLakeClient) ShowSparkOperatorStateInvoker(request *model.ShowSparkOperatorStateRequest) *ShowSparkOperatorStateInvoker {
	requestDef := GenReqDefForShowSparkOperatorState()
	return &ShowSparkOperatorStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelSparkSql 取消SparkSql作业执行
//
// 取消正在运行或排队中的SparkSql作业，此接口为同步接口。只能取消处于QUEUED或RUNNING状态的作业，调用成功后直接返回取消结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CancelSparkSql(request *model.CancelSparkSqlRequest) (*model.CancelSparkSqlResponse, error) {
	requestDef := GenReqDefForCancelSparkSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelSparkSqlResponse), nil
	}
}

// CancelSparkSqlInvoker 取消SparkSql作业执行
func (c *AIDataLakeClient) CancelSparkSqlInvoker(request *model.CancelSparkSqlRequest) *CancelSparkSqlInvoker {
	requestDef := GenReqDefForCancelSparkSql()
	return &CancelSparkSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSparkSqls 查询SparkSql作业列表
//
// 查询工作空间下SparkSql作业列表，此接口为同步接口。支持按作业状态、创建时间、SQL片段等条件进行过滤查询，调用成功后直接返回作业列表数据。可通过[查看SparkSql作业详情](ShowSparkSql.xml)接口查看作业详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListSparkSqls(request *model.ListSparkSqlsRequest) (*model.ListSparkSqlsResponse, error) {
	requestDef := GenReqDefForListSparkSqls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSparkSqlsResponse), nil
	}
}

// ListSparkSqlsInvoker 查询SparkSql作业列表
func (c *AIDataLakeClient) ListSparkSqlsInvoker(request *model.ListSparkSqlsRequest) *ListSparkSqlsInvoker {
	requestDef := GenReqDefForListSparkSqls()
	return &ListSparkSqlsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PreviewSparkSqlResult 预览SparkSql作业查询结果
//
// 预览SparkSql作业的查询结果，此接口为同步接口。仅适用于执行成功的DQL类型作业，可查看作业返回的数据内容，调用成功后直接返回查询结果数据。可通过[查看SparkSql作业详情](ShowSparkSql.xml)接口查看作业详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) PreviewSparkSqlResult(request *model.PreviewSparkSqlResultRequest) (*model.PreviewSparkSqlResultResponse, error) {
	requestDef := GenReqDefForPreviewSparkSqlResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PreviewSparkSqlResultResponse), nil
	}
}

// PreviewSparkSqlResultInvoker 预览SparkSql作业查询结果
func (c *AIDataLakeClient) PreviewSparkSqlResultInvoker(request *model.PreviewSparkSqlResultRequest) *PreviewSparkSqlResultInvoker {
	requestDef := GenReqDefForPreviewSparkSqlResult()
	return &PreviewSparkSqlResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartSparkSqlCluster 重启SparkSql集群
//
// 重启SparkSql集群，该接口为异步接口，接口调用成功后会返回操作ID（operation_id），您可以通过查询Spark异步操作状态接口查询操作执行结果，详情请参见[查询Spark异步操作状态](ShowSparkOperatorState.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) RestartSparkSqlCluster(request *model.RestartSparkSqlClusterRequest) (*model.RestartSparkSqlClusterResponse, error) {
	requestDef := GenReqDefForRestartSparkSqlCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartSparkSqlClusterResponse), nil
	}
}

// RestartSparkSqlClusterInvoker 重启SparkSql集群
func (c *AIDataLakeClient) RestartSparkSqlClusterInvoker(request *model.RestartSparkSqlClusterRequest) *RestartSparkSqlClusterInvoker {
	requestDef := GenReqDefForRestartSparkSqlCluster()
	return &RestartSparkSqlClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunSparkSql 执行SparkSql作业
//
// 执行SparkSql作业，此接口为异步接口。接口调用成功后会返回作业ID（statement_id）,您可以通过查询作业状态接口查询作业执行结果，详情请参见[查询SparkSql作业的状态](ShowSparkSqlState.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) RunSparkSql(request *model.RunSparkSqlRequest) (*model.RunSparkSqlResponse, error) {
	requestDef := GenReqDefForRunSparkSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunSparkSqlResponse), nil
	}
}

// RunSparkSqlInvoker 执行SparkSql作业
func (c *AIDataLakeClient) RunSparkSqlInvoker(request *model.RunSparkSqlRequest) *RunSparkSqlInvoker {
	requestDef := GenReqDefForRunSparkSql()
	return &RunSparkSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSparkSql 查看SparkSql作业详情
//
// 查看SparkSql作业的详细信息，此接口为同步接口。包括作业状态、SQL内容、执行参数等，调用成功后直接返回作业详细信息。可通过[执行SparkSql作业](RunSparkSql.xml)接口创建作业，通过[查询SparkSql作业列表](ListSparkSqls.xml)接口查询作业列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowSparkSql(request *model.ShowSparkSqlRequest) (*model.ShowSparkSqlResponse, error) {
	requestDef := GenReqDefForShowSparkSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSparkSqlResponse), nil
	}
}

// ShowSparkSqlInvoker 查看SparkSql作业详情
func (c *AIDataLakeClient) ShowSparkSqlInvoker(request *model.ShowSparkSqlRequest) *ShowSparkSqlInvoker {
	requestDef := GenReqDefForShowSparkSql()
	return &ShowSparkSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSparkSqlState 查询SparkSql作业的状态
//
// 查询SparkSql作业的状态，此接口为同步接口。可通过查询SparkSql作业列表接口获取statement_id，调用成功后直接返回作业当前状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowSparkSqlState(request *model.ShowSparkSqlStateRequest) (*model.ShowSparkSqlStateResponse, error) {
	requestDef := GenReqDefForShowSparkSqlState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSparkSqlStateResponse), nil
	}
}

// ShowSparkSqlStateInvoker 查询SparkSql作业的状态
func (c *AIDataLakeClient) ShowSparkSqlStateInvoker(request *model.ShowSparkSqlStateRequest) *ShowSparkSqlStateInvoker {
	requestDef := GenReqDefForShowSparkSqlState()
	return &ShowSparkSqlStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
