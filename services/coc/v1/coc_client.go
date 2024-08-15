package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/coc/v1/model"
)

type CocClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCocClient(hcClient *httpclient.HcHttpClient) *CocClient {
	return &CocClient{HcClient: hcClient}
}

func CocClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// ListApplications 查询应用
//
// 查询应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListApplications(request *model.ListApplicationsRequest) (*model.ListApplicationsResponse, error) {
	requestDef := GenReqDefForListApplications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationsResponse), nil
	}
}

// ListApplicationsInvoker 查询应用
func (c *CocClient) ListApplicationsInvoker(request *model.ListApplicationsRequest) *ListApplicationsInvoker {
	requestDef := GenReqDefForListApplications()
	return &ListApplicationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceCompliant 获取节点合规性报告
//
// 分页获取节点合规性报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListInstanceCompliant(request *model.ListInstanceCompliantRequest) (*model.ListInstanceCompliantResponse, error) {
	requestDef := GenReqDefForListInstanceCompliant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceCompliantResponse), nil
	}
}

// ListInstanceCompliantInvoker 获取节点合规性报告
func (c *CocClient) ListInstanceCompliantInvoker(request *model.ListInstanceCompliantRequest) *ListInstanceCompliantInvoker {
	requestDef := GenReqDefForListInstanceCompliant()
	return &ListInstanceCompliantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstancePatchItems 分页获取节点补丁详情
//
// 分页获取节点补丁详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ShowInstancePatchItems(request *model.ShowInstancePatchItemsRequest) (*model.ShowInstancePatchItemsResponse, error) {
	requestDef := GenReqDefForShowInstancePatchItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstancePatchItemsResponse), nil
	}
}

// ShowInstancePatchItemsInvoker 分页获取节点补丁详情
func (c *CocClient) ShowInstancePatchItemsInvoker(request *model.ShowInstancePatchItemsRequest) *ShowInstancePatchItemsInvoker {
	requestDef := GenReqDefForShowInstancePatchItems()
	return &ShowInstancePatchItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateReportCustomEvent 支持用户自主接入告警数据
//
// 支持租户将自开发的监控系统按照标准化集成至COC，集成后告警会按照标准格式上报至COC告警中心
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateReportCustomEvent(request *model.CreateReportCustomEventRequest) (*model.CreateReportCustomEventResponse, error) {
	requestDef := GenReqDefForCreateReportCustomEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateReportCustomEventResponse), nil
	}
}

// CreateReportCustomEventInvoker 支持用户自主接入告警数据
func (c *CocClient) CreateReportCustomEventInvoker(request *model.CreateReportCustomEventRequest) *CreateReportCustomEventInvoker {
	requestDef := GenReqDefForCreateReportCustomEvent()
	return &CreateReportCustomEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateReportPrometheusEvent Prometheus事件接入
//
// # Prometheus事件接入
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateReportPrometheusEvent(request *model.CreateReportPrometheusEventRequest) (*model.CreateReportPrometheusEventResponse, error) {
	requestDef := GenReqDefForCreateReportPrometheusEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateReportPrometheusEventResponse), nil
	}
}

// CreateReportPrometheusEventInvoker Prometheus事件接入
func (c *CocClient) CreateReportPrometheusEventInvoker(request *model.CreateReportPrometheusEventRequest) *CreateReportPrometheusEventInvoker {
	requestDef := GenReqDefForCreateReportPrometheusEvent()
	return &CreateReportPrometheusEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCocIncident CreateExternalIncident 创建事件单
//
// # CreateExternalIncident 创建事件单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateCocIncident(request *model.CreateCocIncidentRequest) (*model.CreateCocIncidentResponse, error) {
	requestDef := GenReqDefForCreateCocIncident()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCocIncidentResponse), nil
	}
}

// CreateCocIncidentInvoker CreateExternalIncident 创建事件单
func (c *CocClient) CreateCocIncidentInvoker(request *model.CreateCocIncidentRequest) *CreateCocIncidentInvoker {
	requestDef := GenReqDefForCreateCocIncident()
	return &CreateCocIncidentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// HandleCocIncident HandleCocIncident处理事件单
//
// # HandleCocIncident 处理事件单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) HandleCocIncident(request *model.HandleCocIncidentRequest) (*model.HandleCocIncidentResponse, error) {
	requestDef := GenReqDefForHandleCocIncident()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.HandleCocIncidentResponse), nil
	}
}

// HandleCocIncidentInvoker HandleCocIncident处理事件单
func (c *CocClient) HandleCocIncidentInvoker(request *model.HandleCocIncidentRequest) *HandleCocIncidentInvoker {
	requestDef := GenReqDefForHandleCocIncident()
	return &HandleCocIncidentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCocTicketOperationHistories GetCocTicketOperationHistories 获取事件单历史
//
// # ListCocTicketOperationHistories  获取事件单历史
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListCocTicketOperationHistories(request *model.ListCocTicketOperationHistoriesRequest) (*model.ListCocTicketOperationHistoriesResponse, error) {
	requestDef := GenReqDefForListCocTicketOperationHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCocTicketOperationHistoriesResponse), nil
	}
}

// ListCocTicketOperationHistoriesInvoker GetCocTicketOperationHistories 获取事件单历史
func (c *CocClient) ListCocTicketOperationHistoriesInvoker(request *model.ListCocTicketOperationHistoriesRequest) *ListCocTicketOperationHistoriesInvoker {
	requestDef := GenReqDefForListCocTicketOperationHistories()
	return &ListCocTicketOperationHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCocIncidentDetail GetCocIncidentDetail 获取事件单详细
//
// # ShowCocIncidentDetail  获取事件单详细
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ShowCocIncidentDetail(request *model.ShowCocIncidentDetailRequest) (*model.ShowCocIncidentDetailResponse, error) {
	requestDef := GenReqDefForShowCocIncidentDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCocIncidentDetailResponse), nil
	}
}

// ShowCocIncidentDetailInvoker GetCocIncidentDetail 获取事件单详细
func (c *CocClient) ShowCocIncidentDetailInvoker(request *model.ShowCocIncidentDetailRequest) *ShowCocIncidentDetailInvoker {
	requestDef := GenReqDefForShowCocIncidentDetail()
	return &ShowCocIncidentDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCocIssues CreateExternalIssues 创建问题单
//
// # CreateExternalIssues 创建问题单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateCocIssues(request *model.CreateCocIssuesRequest) (*model.CreateCocIssuesResponse, error) {
	requestDef := GenReqDefForCreateCocIssues()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCocIssuesResponse), nil
	}
}

// CreateCocIssuesInvoker CreateExternalIssues 创建问题单
func (c *CocClient) CreateCocIssuesInvoker(request *model.CreateCocIssuesRequest) *CreateCocIssuesInvoker {
	requestDef := GenReqDefForCreateCocIssues()
	return &CreateCocIssuesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCocIssuesDetail GetCocIssuesDetail 获取事件单详细
//
// # ShowCocIssuesDetail  获取事件单详细
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ShowCocIssuesDetail(request *model.ShowCocIssuesDetailRequest) (*model.ShowCocIssuesDetailResponse, error) {
	requestDef := GenReqDefForShowCocIssuesDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCocIssuesDetailResponse), nil
	}
}

// ShowCocIssuesDetailInvoker GetCocIssuesDetail 获取事件单详细
func (c *CocClient) ShowCocIssuesDetailInvoker(request *model.ShowCocIssuesDetailRequest) *ShowCocIssuesDetailInvoker {
	requestDef := GenReqDefForShowCocIssuesDetail()
	return &ShowCocIssuesDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuthorizableTicketsExternal 查询COC可授权单列表
//
// 查询COC可授权单列表（变更单号、事件单号、warroom和告警）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListAuthorizableTicketsExternal(request *model.ListAuthorizableTicketsExternalRequest) (*model.ListAuthorizableTicketsExternalResponse, error) {
	requestDef := GenReqDefForListAuthorizableTicketsExternal()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuthorizableTicketsExternalResponse), nil
	}
}

// ListAuthorizableTicketsExternalInvoker 查询COC可授权单列表
func (c *CocClient) ListAuthorizableTicketsExternalInvoker(request *model.ListAuthorizableTicketsExternalRequest) *ListAuthorizableTicketsExternalInvoker {
	requestDef := GenReqDefForListAuthorizableTicketsExternal()
	return &ListAuthorizableTicketsExternalInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResource 查询用户所有资源
//
// 查询用户所有资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListResource(request *model.ListResourceRequest) (*model.ListResourceResponse, error) {
	requestDef := GenReqDefForListResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceResponse), nil
	}
}

// ListResourceInvoker 查询用户所有资源
func (c *CocClient) ListResourceInvoker(request *model.ListResourceRequest) *ListResourceInvoker {
	requestDef := GenReqDefForListResource()
	return &ListResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SyncResource 从RMS同步用户所有资源
//
// 从RMS同步用户所有资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) SyncResource(request *model.SyncResourceRequest) (*model.SyncResourceResponse, error) {
	requestDef := GenReqDefForSyncResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncResourceResponse), nil
	}
}

// SyncResourceInvoker 从RMS同步用户所有资源
func (c *CocClient) SyncResourceInvoker(request *model.SyncResourceRequest) *SyncResourceInvoker {
	requestDef := GenReqDefForSyncResource()
	return &SyncResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetScriptJobBatch 展示批次详情
//
// 查询：批次详情，分页获取批次中的实例列表。
// 过滤条件：分页参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) GetScriptJobBatch(request *model.GetScriptJobBatchRequest) (*model.GetScriptJobBatchResponse, error) {
	requestDef := GenReqDefForGetScriptJobBatch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetScriptJobBatchResponse), nil
	}
}

// GetScriptJobBatchInvoker 展示批次详情
func (c *CocClient) GetScriptJobBatchInvoker(request *model.GetScriptJobBatchRequest) *GetScriptJobBatchInvoker {
	requestDef := GenReqDefForGetScriptJobBatch()
	return &GetScriptJobBatchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetScriptJobInfo 展示脚本工单基本信息
//
// 查询执行：基本信息
// 执行类型、执行名称、创建人、创建时间、结束时间、执行状态、标签（脚本id，脚本名，执行脚本参数，执行用户，超时时长、成功率阈值）
//
// 不同的任务类型消费标签中的不同key
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) GetScriptJobInfo(request *model.GetScriptJobInfoRequest) (*model.GetScriptJobInfoResponse, error) {
	requestDef := GenReqDefForGetScriptJobInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetScriptJobInfoResponse), nil
	}
}

// GetScriptJobInfoInvoker 展示脚本工单基本信息
func (c *CocClient) GetScriptJobInfoInvoker(request *model.GetScriptJobInfoRequest) *GetScriptJobInfoInvoker {
	requestDef := GenReqDefForGetScriptJobInfo()
	return &GetScriptJobInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetScriptJobStatistics 展示实例状态统计信息
//
// 查询：实例状态统计信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) GetScriptJobStatistics(request *model.GetScriptJobStatisticsRequest) (*model.GetScriptJobStatisticsResponse, error) {
	requestDef := GenReqDefForGetScriptJobStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetScriptJobStatisticsResponse), nil
	}
}

// GetScriptJobStatisticsInvoker 展示实例状态统计信息
func (c *CocClient) GetScriptJobStatisticsInvoker(request *model.GetScriptJobStatisticsRequest) *GetScriptJobStatisticsInvoker {
	requestDef := GenReqDefForGetScriptJobStatistics()
	return &GetScriptJobStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScriptJobBatches 展示批次列表
//
// 查询：批次列表
// 返回：批次index、批次标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListScriptJobBatches(request *model.ListScriptJobBatchesRequest) (*model.ListScriptJobBatchesResponse, error) {
	requestDef := GenReqDefForListScriptJobBatches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScriptJobBatchesResponse), nil
	}
}

// ListScriptJobBatchesInvoker 展示批次列表
func (c *CocClient) ListScriptJobBatchesInvoker(request *model.ListScriptJobBatchesRequest) *ListScriptJobBatchesInvoker {
	requestDef := GenReqDefForListScriptJobBatches()
	return &ListScriptJobBatchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScriptJobs 展示工单列表
//
// 查询作业工单列表，分页查询
// 过滤：创建时间开始，创建时间结束、创建人
// 返回：id、脚本名称、区域、创建人、创建时间、结束时间、总耗时、状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListScriptJobs(request *model.ListScriptJobsRequest) (*model.ListScriptJobsResponse, error) {
	requestDef := GenReqDefForListScriptJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScriptJobsResponse), nil
	}
}

// ListScriptJobsInvoker 展示工单列表
func (c *CocClient) ListScriptJobsInvoker(request *model.ListScriptJobsRequest) *ListScriptJobsInvoker {
	requestDef := GenReqDefForListScriptJobs()
	return &ListScriptJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// OperateScriptJob 操作脚本工单
//
// 操作类型：取消实例、跳过批次、取消整个工单、暂停整个工单、继续整个工单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) OperateScriptJob(request *model.OperateScriptJobRequest) (*model.OperateScriptJobResponse, error) {
	requestDef := GenReqDefForOperateScriptJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.OperateScriptJobResponse), nil
	}
}

// OperateScriptJobInvoker 操作脚本工单
func (c *CocClient) OperateScriptJobInvoker(request *model.OperateScriptJobRequest) *OperateScriptJobInvoker {
	requestDef := GenReqDefForOperateScriptJob()
	return &OperateScriptJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateScript 创建脚本
//
// 创建作业脚本：自定义脚本
// - 脚本有标签属性，表示是高危脚本。创建时候不需要对脚本进行是否是高危的二次校验。
// - 进行租户隔离；北向接口创建的脚本，审批人字段不填写，默认不需要审批
// - 约束条件：
// - 脚本名称：同一租户下，脚本名称不能重复，最大字符64个字符，支持中文+字母+数字+下划线。
// - 脚本内容最大100kb。
// - 脚本参数个数最多20个。
// - 脚本描述：最大256个字符。
// - 单个参数的参数名称 64个字符，只支持字母+数字+下划线。
// - 单个参数的值最大1024个字符，正则表达式如下：^((?!\\.{2,})[a-zA-Z0-9_\\-/\\.\\*\\x20\\?:\&quot;,&#x3D;+@\\\\\\[\\{\\]\\}])*$。
// - 审批人最多支持5人。
// - 脚本输出的日志总量只支持1MB。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateScript(request *model.CreateScriptRequest) (*model.CreateScriptResponse, error) {
	requestDef := GenReqDefForCreateScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScriptResponse), nil
	}
}

// CreateScriptInvoker 创建脚本
func (c *CocClient) CreateScriptInvoker(request *model.CreateScriptRequest) *CreateScriptInvoker {
	requestDef := GenReqDefForCreateScript()
	return &CreateScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteScript 删除自定义脚本
//
// 删除作业脚本：自定义脚本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) DeleteScript(request *model.DeleteScriptRequest) (*model.DeleteScriptResponse, error) {
	requestDef := GenReqDefForDeleteScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScriptResponse), nil
	}
}

// DeleteScriptInvoker 删除自定义脚本
func (c *CocClient) DeleteScriptInvoker(request *model.DeleteScriptRequest) *DeleteScriptInvoker {
	requestDef := GenReqDefForDeleteScript()
	return &DeleteScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteScript 执行自定义脚本
//
// 执行脚本
//
// 脚本入参、超时时间、执行用户、资源受限
// 脚本入参支持20个。
// 单次下发的机器支持200个。
// 单次批次内机器数量最大10个。
// 最大批次数量为20批。
// 脚本输出的日志总量只支持1MB。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ExecuteScript(request *model.ExecuteScriptRequest) (*model.ExecuteScriptResponse, error) {
	requestDef := GenReqDefForExecuteScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteScriptResponse), nil
	}
}

// ExecuteScriptInvoker 执行自定义脚本
func (c *CocClient) ExecuteScriptInvoker(request *model.ExecuteScriptRequest) *ExecuteScriptInvoker {
	requestDef := GenReqDefForExecuteScript()
	return &ExecuteScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetScript 获取自定义脚本详情
//
// 获取脚本详情
// 约束条件：
// 只能查询自定义脚本详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) GetScript(request *model.GetScriptRequest) (*model.GetScriptResponse, error) {
	requestDef := GenReqDefForGetScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetScriptResponse), nil
	}
}

// GetScriptInvoker 获取自定义脚本详情
func (c *CocClient) GetScriptInvoker(request *model.GetScriptRequest) *GetScriptInvoker {
	requestDef := GenReqDefForGetScript()
	return &GetScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScripts 查询脚本列表
//
// 作业脚本列表：自定义脚本
//
// limit最大为100
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListScripts(request *model.ListScriptsRequest) (*model.ListScriptsResponse, error) {
	requestDef := GenReqDefForListScripts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScriptsResponse), nil
	}
}

// ListScriptsInvoker 查询脚本列表
func (c *CocClient) ListScriptsInvoker(request *model.ListScriptsRequest) *ListScriptsInvoker {
	requestDef := GenReqDefForListScripts()
	return &ListScriptsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateScript 修改脚本
//
// 修改作业脚本：自定义脚本
// 约束条件：
// 脚本名称：同一租户下，脚本名称不能重复，最大字符64个字符，支持中文+字母+数字+下划线。
// 脚本内容最大4096个字符。
// 脚本参数个数最多20个。
// 脚本描述：最大256个字符。
// 单个参数的参数名称 64个字符，只支持字母+数字+下划线。
// 单个参数的值最大1024个字符，正则表达式如下：^((?!.{2,})[a-zA-Z0-9_-/.*\\x20?:\&quot;,&#x3D;+@\\[{]}])*$。
// 修改的脚本如果有审批人，在修改之后，需要再次选择审批人，查询审批
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) UpdateScript(request *model.UpdateScriptRequest) (*model.UpdateScriptResponse, error) {
	requestDef := GenReqDefForUpdateScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateScriptResponse), nil
	}
}

// UpdateScriptInvoker 修改脚本
func (c *CocClient) UpdateScriptInvoker(request *model.UpdateScriptRequest) *UpdateScriptInvoker {
	requestDef := GenReqDefForUpdateScript()
	return &UpdateScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecutePublicScript 执行公共脚本
//
// 执行公共脚本
// 脚本入参、超时时间、执行用户、资源受限
// 脚本入参支持20个。
// 单次下发的机器支持200个。
// 单次批次内机器数量最大10个。
// 最大批次数量为20批。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ExecutePublicScript(request *model.ExecutePublicScriptRequest) (*model.ExecutePublicScriptResponse, error) {
	requestDef := GenReqDefForExecutePublicScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecutePublicScriptResponse), nil
	}
}

// ExecutePublicScriptInvoker 执行公共脚本
func (c *CocClient) ExecutePublicScriptInvoker(request *model.ExecutePublicScriptRequest) *ExecutePublicScriptInvoker {
	requestDef := GenReqDefForExecutePublicScript()
	return &ExecutePublicScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetPublicScript 展示公共脚本详情
//
// 展示公共脚本详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) GetPublicScript(request *model.GetPublicScriptRequest) (*model.GetPublicScriptResponse, error) {
	requestDef := GenReqDefForGetPublicScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetPublicScriptResponse), nil
	}
}

// GetPublicScriptInvoker 展示公共脚本详情
func (c *CocClient) GetPublicScriptInvoker(request *model.GetPublicScriptRequest) *GetPublicScriptInvoker {
	requestDef := GenReqDefForGetPublicScript()
	return &GetPublicScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPublicScripts 获取公共脚本列表
//
// 获取公共脚本列表，分页逻辑：采用limit+marker方式，提高分页效率。用自增id作为marker参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListPublicScripts(request *model.ListPublicScriptsRequest) (*model.ListPublicScriptsResponse, error) {
	requestDef := GenReqDefForListPublicScripts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublicScriptsResponse), nil
	}
}

// ListPublicScriptsInvoker 获取公共脚本列表
func (c *CocClient) ListPublicScriptsInvoker(request *model.ListPublicScriptsRequest) *ListPublicScriptsInvoker {
	requestDef := GenReqDefForListPublicScripts()
	return &ListPublicScriptsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWarRoom 创建租户区WarRoom
//
// 创建租户区WarRoom
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateWarRoom(request *model.CreateWarRoomRequest) (*model.CreateWarRoomResponse, error) {
	requestDef := GenReqDefForCreateWarRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWarRoomResponse), nil
	}
}

// CreateWarRoomInvoker 创建租户区WarRoom
func (c *CocClient) CreateWarRoomInvoker(request *model.CreateWarRoomRequest) *CreateWarRoomInvoker {
	requestDef := GenReqDefForCreateWarRoom()
	return &CreateWarRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWarRooms 查询租户区WarRoom信息列表
//
// 查询租户区WarRoom信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListWarRooms(request *model.ListWarRoomsRequest) (*model.ListWarRoomsResponse, error) {
	requestDef := GenReqDefForListWarRooms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWarRoomsResponse), nil
	}
}

// ListWarRoomsInvoker 查询租户区WarRoom信息列表
func (c *CocClient) ListWarRoomsInvoker(request *model.ListWarRoomsRequest) *ListWarRoomsInvoker {
	requestDef := GenReqDefForListWarRooms()
	return &ListWarRoomsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
