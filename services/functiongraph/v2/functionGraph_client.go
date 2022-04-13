package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/functiongraph/v2/model"
)

type FunctionGraphClient struct {
	HcClient *http_client.HcHttpClient
}

func NewFunctionGraphClient(hcClient *http_client.HcHttpClient) *FunctionGraphClient {
	return &FunctionGraphClient{HcClient: hcClient}
}

func FunctionGraphClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//异步执行函数。
func (c *FunctionGraphClient) AsyncInvokeFunction(request *model.AsyncInvokeFunctionRequest) (*model.AsyncInvokeFunctionResponse, error) {
	requestDef := GenReqDefForAsyncInvokeFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AsyncInvokeFunctionResponse), nil
	}
}

//函数异步执行并返回预留实例ID用于场景指客户端请求执行比较费时任务，不需要同步等待执行完成返回结果，该方法提前返回任务执行对应的预留实例ID, 如果预留实例有异常， 可以通过该实例ID把对应实例删除（该接口主要针对白名单用户）。
func (c *FunctionGraphClient) AsyncInvokeReservedFunction(request *model.AsyncInvokeReservedFunctionRequest) (*model.AsyncInvokeReservedFunctionResponse, error) {
	requestDef := GenReqDefForAsyncInvokeReservedFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AsyncInvokeReservedFunctionResponse), nil
	}
}

//创建依赖包。
func (c *FunctionGraphClient) CreateDependency(request *model.CreateDependencyRequest) (*model.CreateDependencyResponse, error) {
	requestDef := GenReqDefForCreateDependency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDependencyResponse), nil
	}
}

//创建测试事件。
func (c *FunctionGraphClient) CreateEvent(request *model.CreateEventRequest) (*model.CreateEventResponse, error) {
	requestDef := GenReqDefForCreateEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEventResponse), nil
	}
}

//创建指定的函数。
func (c *FunctionGraphClient) CreateFunction(request *model.CreateFunctionRequest) (*model.CreateFunctionResponse, error) {
	requestDef := GenReqDefForCreateFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFunctionResponse), nil
	}
}

//发布函数版本。
func (c *FunctionGraphClient) CreateFunctionVersion(request *model.CreateFunctionVersionRequest) (*model.CreateFunctionVersionResponse, error) {
	requestDef := GenReqDefForCreateFunctionVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFunctionVersionResponse), nil
	}
}

//创建函数灰度版本别名。
func (c *FunctionGraphClient) CreateVersionAlias(request *model.CreateVersionAliasRequest) (*model.CreateVersionAliasResponse, error) {
	requestDef := GenReqDefForCreateVersionAlias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVersionAliasResponse), nil
	}
}

//删除指定的依赖包。
func (c *FunctionGraphClient) DeleteDependency(request *model.DeleteDependencyRequest) (*model.DeleteDependencyResponse, error) {
	requestDef := GenReqDefForDeleteDependency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDependencyResponse), nil
	}
}

//删除测试事件。
func (c *FunctionGraphClient) DeleteEvent(request *model.DeleteEventRequest) (*model.DeleteEventResponse, error) {
	requestDef := GenReqDefForDeleteEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEventResponse), nil
	}
}

//删除指定的函数或者特定的版本（不允许删除latest版本）。  如果URN中包含函数版本或者别名，则删除特定的函数版本或者别名指向的版本以及该版本关联的trigger。 如果URN中不包含版本或者别名，则删除整个函数，包含所有版本以及别名，触发器。
func (c *FunctionGraphClient) DeleteFunction(request *model.DeleteFunctionRequest) (*model.DeleteFunctionResponse, error) {
	requestDef := GenReqDefForDeleteFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFunctionResponse), nil
	}
}

//删除函数异步配置信息。
func (c *FunctionGraphClient) DeleteFunctionAsyncInvokeConfig(request *model.DeleteFunctionAsyncInvokeConfigRequest) (*model.DeleteFunctionAsyncInvokeConfigResponse, error) {
	requestDef := GenReqDefForDeleteFunctionAsyncInvokeConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFunctionAsyncInvokeConfigResponse), nil
	}
}

//删除函数版本别名。
func (c *FunctionGraphClient) DeleteVersionAlias(request *model.DeleteVersionAliasRequest) (*model.DeleteVersionAliasResponse, error) {
	requestDef := GenReqDefForDeleteVersionAlias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVersionAliasResponse), nil
	}
}

//开通lts日志上报功能。
func (c *FunctionGraphClient) EnableLtsLogs(request *model.EnableLtsLogsRequest) (*model.EnableLtsLogsResponse, error) {
	requestDef := GenReqDefForEnableLtsLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableLtsLogsResponse), nil
	}
}

//导出函数。
func (c *FunctionGraphClient) ExportFunction(request *model.ExportFunctionRequest) (*model.ExportFunctionResponse, error) {
	requestDef := GenReqDefForExportFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportFunctionResponse), nil
	}
}

//导入函数。
func (c *FunctionGraphClient) ImportFunction(request *model.ImportFunctionRequest) (*model.ImportFunctionResponse, error) {
	requestDef := GenReqDefForImportFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportFunctionResponse), nil
	}
}

//同步调用指的是客户端请求需要明确等到响应结果，也就是说这样的请求必须得调用到用户的函数，并且等到调用完成才返回。
func (c *FunctionGraphClient) InvokeFunction(request *model.InvokeFunctionRequest) (*model.InvokeFunctionResponse, error) {
	requestDef := GenReqDefForInvokeFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InvokeFunctionResponse), nil
	}
}

//获取依赖包列表。
func (c *FunctionGraphClient) ListDependencies(request *model.ListDependenciesRequest) (*model.ListDependenciesResponse, error) {
	requestDef := GenReqDefForListDependencies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDependenciesResponse), nil
	}
}

//获取指定函数的测试事件列表。
func (c *FunctionGraphClient) ListEvents(request *model.ListEventsRequest) (*model.ListEventsResponse, error) {
	requestDef := GenReqDefForListEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventsResponse), nil
	}
}

//获取函数异步调用请求列表
func (c *FunctionGraphClient) ListFunctionAsyncInvocations(request *model.ListFunctionAsyncInvocationsRequest) (*model.ListFunctionAsyncInvocationsResponse, error) {
	requestDef := GenReqDefForListFunctionAsyncInvocations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionAsyncInvocationsResponse), nil
	}
}

//获取函数异步配置列表。
func (c *FunctionGraphClient) ListFunctionAsyncInvokeConfig(request *model.ListFunctionAsyncInvokeConfigRequest) (*model.ListFunctionAsyncInvokeConfigResponse, error) {
	requestDef := GenReqDefForListFunctionAsyncInvokeConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionAsyncInvokeConfigResponse), nil
	}
}

//获取指定时间段的函数运行指标。
func (c *FunctionGraphClient) ListFunctionStatistics(request *model.ListFunctionStatisticsRequest) (*model.ListFunctionStatisticsResponse, error) {
	requestDef := GenReqDefForListFunctionStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionStatisticsResponse), nil
	}
}

//获取指定函数的版本列表。
func (c *FunctionGraphClient) ListFunctionVersions(request *model.ListFunctionVersionsRequest) (*model.ListFunctionVersionsResponse, error) {
	requestDef := GenReqDefForListFunctionVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionVersionsResponse), nil
	}
}

//获取函数列表
func (c *FunctionGraphClient) ListFunctions(request *model.ListFunctionsRequest) (*model.ListFunctionsResponse, error) {
	requestDef := GenReqDefForListFunctions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionsResponse), nil
	}
}

//查询租户配额
func (c *FunctionGraphClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

//租户函数统计信息。  返回三类的统计信息，函数格式和大小使用情况包括配额和使用量，流量报告。 通过查询参数filter可以进行过滤，查询参数period可以指定返回的时间段。
func (c *FunctionGraphClient) ListStatistics(request *model.ListStatisticsRequest) (*model.ListStatisticsResponse, error) {
	requestDef := GenReqDefForListStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStatisticsResponse), nil
	}
}

//获取函数版本别名列表。
func (c *FunctionGraphClient) ListVersionAliases(request *model.ListVersionAliasesRequest) (*model.ListVersionAliasesResponse, error) {
	requestDef := GenReqDefForListVersionAliases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVersionAliasesResponse), nil
	}
}

//获取指定依赖包。
func (c *FunctionGraphClient) ShowDependency(request *model.ShowDependencyRequest) (*model.ShowDependencyResponse, error) {
	requestDef := GenReqDefForShowDependency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDependencyResponse), nil
	}
}

//获取测试事件详细信息。
func (c *FunctionGraphClient) ShowEvent(request *model.ShowEventRequest) (*model.ShowEventResponse, error) {
	requestDef := GenReqDefForShowEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEventResponse), nil
	}
}

//获取函数异步配置信息。
func (c *FunctionGraphClient) ShowFunctionAsyncInvokeConfig(request *model.ShowFunctionAsyncInvokeConfigRequest) (*model.ShowFunctionAsyncInvokeConfigResponse, error) {
	requestDef := GenReqDefForShowFunctionAsyncInvokeConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFunctionAsyncInvokeConfigResponse), nil
	}
}

//获取指定函数的代码。
func (c *FunctionGraphClient) ShowFunctionCode(request *model.ShowFunctionCodeRequest) (*model.ShowFunctionCodeResponse, error) {
	requestDef := GenReqDefForShowFunctionCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFunctionCodeResponse), nil
	}
}

//获取指定函数的metadata。
func (c *FunctionGraphClient) ShowFunctionConfig(request *model.ShowFunctionConfigRequest) (*model.ShowFunctionConfigResponse, error) {
	requestDef := GenReqDefForShowFunctionConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFunctionConfigResponse), nil
	}
}

//获取指定函数的lts日志组日志流配置。
func (c *FunctionGraphClient) ShowLtsLogDetails(request *model.ShowLtsLogDetailsRequest) (*model.ShowLtsLogDetailsResponse, error) {
	requestDef := GenReqDefForShowLtsLogDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLtsLogDetailsResponse), nil
	}
}

//获取函数调用链配置
func (c *FunctionGraphClient) ShowTracing(request *model.ShowTracingRequest) (*model.ShowTracingResponse, error) {
	requestDef := GenReqDefForShowTracing()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTracingResponse), nil
	}
}

//获取函数指定的版本别名信息。
func (c *FunctionGraphClient) ShowVersionAlias(request *model.ShowVersionAliasRequest) (*model.ShowVersionAliasResponse, error) {
	requestDef := GenReqDefForShowVersionAlias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVersionAliasResponse), nil
	}
}

//更新依赖包指定依赖包。
func (c *FunctionGraphClient) UpdateDependency(request *model.UpdateDependencyRequest) (*model.UpdateDependencyResponse, error) {
	requestDef := GenReqDefForUpdateDependency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDependencyResponse), nil
	}
}

//更新测试事件。
func (c *FunctionGraphClient) UpdateEvent(request *model.UpdateEventRequest) (*model.UpdateEventResponse, error) {
	requestDef := GenReqDefForUpdateEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEventResponse), nil
	}
}

//设置函数异步配置信息。
func (c *FunctionGraphClient) UpdateFunctionAsyncInvokeConfig(request *model.UpdateFunctionAsyncInvokeConfigRequest) (*model.UpdateFunctionAsyncInvokeConfigResponse, error) {
	requestDef := GenReqDefForUpdateFunctionAsyncInvokeConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFunctionAsyncInvokeConfigResponse), nil
	}
}

//修改指定的函数的代码。
func (c *FunctionGraphClient) UpdateFunctionCode(request *model.UpdateFunctionCodeRequest) (*model.UpdateFunctionCodeResponse, error) {
	requestDef := GenReqDefForUpdateFunctionCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFunctionCodeResponse), nil
	}
}

//修改指定的函数的metadata信息。
func (c *FunctionGraphClient) UpdateFunctionConfig(request *model.UpdateFunctionConfigRequest) (*model.UpdateFunctionConfigResponse, error) {
	requestDef := GenReqDefForUpdateFunctionConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFunctionConfigResponse), nil
	}
}

//为函数绑定预留实例
func (c *FunctionGraphClient) UpdateFunctionReservedInstances(request *model.UpdateFunctionReservedInstancesRequest) (*model.UpdateFunctionReservedInstancesResponse, error) {
	requestDef := GenReqDefForUpdateFunctionReservedInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFunctionReservedInstancesResponse), nil
	}
}

//修改函数调用链配置,开通/修改传入aksk，关闭aksk传空
func (c *FunctionGraphClient) UpdateTracing(request *model.UpdateTracingRequest) (*model.UpdateTracingResponse, error) {
	requestDef := GenReqDefForUpdateTracing()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTracingResponse), nil
	}
}

//修改函数版本别名信息。
func (c *FunctionGraphClient) UpdateVersionAlias(request *model.UpdateVersionAliasRequest) (*model.UpdateVersionAliasResponse, error) {
	requestDef := GenReqDefForUpdateVersionAlias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVersionAliasResponse), nil
	}
}

//删除指定函数所有触发器设置。  在提供函数版本且非latest的情况下，删除对应函数版本的触发器。 在提供函数别名的情况下，删除对应函数别名的触发器。 在不提供函数版本（也不提供别名）或版本为latest的情况下，删除该函数所有的触发器（包括所有版本和别名）。
func (c *FunctionGraphClient) BatchDeleteFunctionTriggers(request *model.BatchDeleteFunctionTriggersRequest) (*model.BatchDeleteFunctionTriggersResponse, error) {
	requestDef := GenReqDefForBatchDeleteFunctionTriggers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteFunctionTriggersResponse), nil
	}
}

//创建触发器。  - 可以创建的触发器类型包括TIMER、APIG、CTS、DDS、DMS、DIS、LTS、OBS、SMN、KAFKA。 - DDS和KAFKA触发器创建时默认为DISABLED状态，其他触发器默认为ACTIVE状态。 - TIMER、DDS、DMS、KAFKA、LTS触发器支持禁用，其他触发器不支持。
func (c *FunctionGraphClient) CreateFunctionTrigger(request *model.CreateFunctionTriggerRequest) (*model.CreateFunctionTriggerResponse, error) {
	requestDef := GenReqDefForCreateFunctionTrigger()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFunctionTriggerResponse), nil
	}
}

//删除触发器。
func (c *FunctionGraphClient) DeleteFunctionTrigger(request *model.DeleteFunctionTriggerRequest) (*model.DeleteFunctionTriggerResponse, error) {
	requestDef := GenReqDefForDeleteFunctionTrigger()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFunctionTriggerResponse), nil
	}
}

//获取指定函数的所有触发器设置。
func (c *FunctionGraphClient) ListFunctionTriggers(request *model.ListFunctionTriggersRequest) (*model.ListFunctionTriggersResponse, error) {
	requestDef := GenReqDefForListFunctionTriggers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionTriggersResponse), nil
	}
}

//获取特定触发器的信息。
func (c *FunctionGraphClient) ShowFunctionTrigger(request *model.ShowFunctionTriggerRequest) (*model.ShowFunctionTriggerResponse, error) {
	requestDef := GenReqDefForShowFunctionTrigger()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFunctionTriggerResponse), nil
	}
}

//更新触发器
func (c *FunctionGraphClient) UpdateTrigger(request *model.UpdateTriggerRequest) (*model.UpdateTriggerResponse, error) {
	requestDef := GenReqDefForUpdateTrigger()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTriggerResponse), nil
	}
}

//删除工作流列表
func (c *FunctionGraphClient) BatchDeleteWorkflows(request *model.BatchDeleteWorkflowsRequest) (*model.BatchDeleteWorkflowsResponse, error) {
	requestDef := GenReqDefForBatchDeleteWorkflows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteWorkflowsResponse), nil
	}
}

//创建工作流列表
func (c *FunctionGraphClient) CreateWorkflow(request *model.CreateWorkflowRequest) (*model.CreateWorkflowResponse, error) {
	requestDef := GenReqDefForCreateWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkflowResponse), nil
	}
}

//获取指定函数流执行实例列表
func (c *FunctionGraphClient) ListWorkflowExecutions(request *model.ListWorkflowExecutionsRequest) (*model.ListWorkflowExecutionsResponse, error) {
	requestDef := GenReqDefForListWorkflowExecutions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkflowExecutionsResponse), nil
	}
}

//查询工作流列表
func (c *FunctionGraphClient) ListWorkflows(request *model.ListWorkflowsRequest) (*model.ListWorkflowsResponse, error) {
	requestDef := GenReqDefForListWorkflows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkflowsResponse), nil
	}
}

//重试工作流
func (c *FunctionGraphClient) RetryWorkFlow(request *model.RetryWorkFlowRequest) (*model.RetryWorkFlowResponse, error) {
	requestDef := GenReqDefForRetryWorkFlow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryWorkFlowResponse), nil
	}
}

//获取函数流指标
func (c *FunctionGraphClient) ShowTenantMetric(request *model.ShowTenantMetricRequest) (*model.ShowTenantMetricResponse, error) {
	requestDef := GenReqDefForShowTenantMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTenantMetricResponse), nil
	}
}

//获取指定函数流实例
func (c *FunctionGraphClient) ShowWorkFlow(request *model.ShowWorkFlowRequest) (*model.ShowWorkFlowResponse, error) {
	requestDef := GenReqDefForShowWorkFlow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkFlowResponse), nil
	}
}

//获取指定工作流指标
func (c *FunctionGraphClient) ShowWorkFlowMetric(request *model.ShowWorkFlowMetricRequest) (*model.ShowWorkFlowMetricResponse, error) {
	requestDef := GenReqDefForShowWorkFlowMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkFlowMetricResponse), nil
	}
}

//获取指定函数流执行实例。
func (c *FunctionGraphClient) ShowWorkflowExecution(request *model.ShowWorkflowExecutionRequest) (*model.ShowWorkflowExecutionResponse, error) {
	requestDef := GenReqDefForShowWorkflowExecution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkflowExecutionResponse), nil
	}
}

//开始执行函数流
func (c *FunctionGraphClient) StartWorkflowExecution(request *model.StartWorkflowExecutionRequest) (*model.StartWorkflowExecutionResponse, error) {
	requestDef := GenReqDefForStartWorkflowExecution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartWorkflowExecutionResponse), nil
	}
}

//停止工作流
func (c *FunctionGraphClient) StopWorkFlow(request *model.StopWorkFlowRequest) (*model.StopWorkFlowResponse, error) {
	requestDef := GenReqDefForStopWorkFlow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopWorkFlowResponse), nil
	}
}

//修改指定函数流实例
func (c *FunctionGraphClient) UpdateWorkFlow(request *model.UpdateWorkFlowRequest) (*model.UpdateWorkFlowResponse, error) {
	requestDef := GenReqDefForUpdateWorkFlow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkFlowResponse), nil
	}
}
