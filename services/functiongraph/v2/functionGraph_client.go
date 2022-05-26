package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
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

// AsyncInvokeFunction 异步执行函数。
//
// 异步执行函数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) AsyncInvokeFunction(request *model.AsyncInvokeFunctionRequest) (*model.AsyncInvokeFunctionResponse, error) {
	requestDef := GenReqDefForAsyncInvokeFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AsyncInvokeFunctionResponse), nil
	}
}

// AsyncInvokeFunctionInvoker 异步执行函数。
func (c *FunctionGraphClient) AsyncInvokeFunctionInvoker(request *model.AsyncInvokeFunctionRequest) *AsyncInvokeFunctionInvoker {
	requestDef := GenReqDefForAsyncInvokeFunction()
	return &AsyncInvokeFunctionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AsyncInvokeReservedFunction 函数异步执行并返回预留实例ID。
//
// 函数异步执行并返回预留实例ID用于场景指客户端请求执行比较费时任务，不需要同步等待执行完成返回结果，该方法提前返回任务执行对应的预留实例ID, 如果预留实例有异常，
// 可以通过该实例ID把对应实例删除（该接口主要针对白名单用户）。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) AsyncInvokeReservedFunction(request *model.AsyncInvokeReservedFunctionRequest) (*model.AsyncInvokeReservedFunctionResponse, error) {
	requestDef := GenReqDefForAsyncInvokeReservedFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AsyncInvokeReservedFunctionResponse), nil
	}
}

// AsyncInvokeReservedFunctionInvoker 函数异步执行并返回预留实例ID。
func (c *FunctionGraphClient) AsyncInvokeReservedFunctionInvoker(request *model.AsyncInvokeReservedFunctionRequest) *AsyncInvokeReservedFunctionInvoker {
	requestDef := GenReqDefForAsyncInvokeReservedFunction()
	return &AsyncInvokeReservedFunctionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelAsyncInvocation 停止函数异步调用请求
//
// 停止函数异步调用请求
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) CancelAsyncInvocation(request *model.CancelAsyncInvocationRequest) (*model.CancelAsyncInvocationResponse, error) {
	requestDef := GenReqDefForCancelAsyncInvocation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelAsyncInvocationResponse), nil
	}
}

// CancelAsyncInvocationInvoker 停止函数异步调用请求
func (c *FunctionGraphClient) CancelAsyncInvocationInvoker(request *model.CancelAsyncInvocationRequest) *CancelAsyncInvocationInvoker {
	requestDef := GenReqDefForCancelAsyncInvocation()
	return &CancelAsyncInvocationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDependency 创建依赖包
//
// 创建依赖包。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) CreateDependency(request *model.CreateDependencyRequest) (*model.CreateDependencyResponse, error) {
	requestDef := GenReqDefForCreateDependency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDependencyResponse), nil
	}
}

// CreateDependencyInvoker 创建依赖包
func (c *FunctionGraphClient) CreateDependencyInvoker(request *model.CreateDependencyRequest) *CreateDependencyInvoker {
	requestDef := GenReqDefForCreateDependency()
	return &CreateDependencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEvent 创建测试事件
//
// 创建测试事件。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) CreateEvent(request *model.CreateEventRequest) (*model.CreateEventResponse, error) {
	requestDef := GenReqDefForCreateEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEventResponse), nil
	}
}

// CreateEventInvoker 创建测试事件
func (c *FunctionGraphClient) CreateEventInvoker(request *model.CreateEventRequest) *CreateEventInvoker {
	requestDef := GenReqDefForCreateEvent()
	return &CreateEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFunction 创建函数。
//
// 创建指定的函数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) CreateFunction(request *model.CreateFunctionRequest) (*model.CreateFunctionResponse, error) {
	requestDef := GenReqDefForCreateFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFunctionResponse), nil
	}
}

// CreateFunctionInvoker 创建函数。
func (c *FunctionGraphClient) CreateFunctionInvoker(request *model.CreateFunctionRequest) *CreateFunctionInvoker {
	requestDef := GenReqDefForCreateFunction()
	return &CreateFunctionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFunctionVersion 发布函数版本。
//
// 发布函数版本。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) CreateFunctionVersion(request *model.CreateFunctionVersionRequest) (*model.CreateFunctionVersionResponse, error) {
	requestDef := GenReqDefForCreateFunctionVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFunctionVersionResponse), nil
	}
}

// CreateFunctionVersionInvoker 发布函数版本。
func (c *FunctionGraphClient) CreateFunctionVersionInvoker(request *model.CreateFunctionVersionRequest) *CreateFunctionVersionInvoker {
	requestDef := GenReqDefForCreateFunctionVersion()
	return &CreateFunctionVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVersionAlias 创建函数版本别名。
//
// 创建函数灰度版本别名。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) CreateVersionAlias(request *model.CreateVersionAliasRequest) (*model.CreateVersionAliasResponse, error) {
	requestDef := GenReqDefForCreateVersionAlias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVersionAliasResponse), nil
	}
}

// CreateVersionAliasInvoker 创建函数版本别名。
func (c *FunctionGraphClient) CreateVersionAliasInvoker(request *model.CreateVersionAliasRequest) *CreateVersionAliasInvoker {
	requestDef := GenReqDefForCreateVersionAlias()
	return &CreateVersionAliasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDependency 删除依赖包
//
// 删除指定的依赖包。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) DeleteDependency(request *model.DeleteDependencyRequest) (*model.DeleteDependencyResponse, error) {
	requestDef := GenReqDefForDeleteDependency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDependencyResponse), nil
	}
}

// DeleteDependencyInvoker 删除依赖包
func (c *FunctionGraphClient) DeleteDependencyInvoker(request *model.DeleteDependencyRequest) *DeleteDependencyInvoker {
	requestDef := GenReqDefForDeleteDependency()
	return &DeleteDependencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEvent 删除测试事件
//
// 删除测试事件。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) DeleteEvent(request *model.DeleteEventRequest) (*model.DeleteEventResponse, error) {
	requestDef := GenReqDefForDeleteEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEventResponse), nil
	}
}

// DeleteEventInvoker 删除测试事件
func (c *FunctionGraphClient) DeleteEventInvoker(request *model.DeleteEventRequest) *DeleteEventInvoker {
	requestDef := GenReqDefForDeleteEvent()
	return &DeleteEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFunction 删除函数/版本。
//
// 删除指定的函数或者特定的版本（不允许删除latest版本）。
//
// 如果URN中包含函数版本或者别名，则删除特定的函数版本或者别名指向的版本以及该版本关联的trigger。
// 如果URN中不包含版本或者别名，则删除整个函数，包含所有版本以及别名，触发器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) DeleteFunction(request *model.DeleteFunctionRequest) (*model.DeleteFunctionResponse, error) {
	requestDef := GenReqDefForDeleteFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFunctionResponse), nil
	}
}

// DeleteFunctionInvoker 删除函数/版本。
func (c *FunctionGraphClient) DeleteFunctionInvoker(request *model.DeleteFunctionRequest) *DeleteFunctionInvoker {
	requestDef := GenReqDefForDeleteFunction()
	return &DeleteFunctionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFunctionAsyncInvokeConfig 删除函数异步配置信息。
//
// 删除函数异步配置信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) DeleteFunctionAsyncInvokeConfig(request *model.DeleteFunctionAsyncInvokeConfigRequest) (*model.DeleteFunctionAsyncInvokeConfigResponse, error) {
	requestDef := GenReqDefForDeleteFunctionAsyncInvokeConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFunctionAsyncInvokeConfigResponse), nil
	}
}

// DeleteFunctionAsyncInvokeConfigInvoker 删除函数异步配置信息。
func (c *FunctionGraphClient) DeleteFunctionAsyncInvokeConfigInvoker(request *model.DeleteFunctionAsyncInvokeConfigRequest) *DeleteFunctionAsyncInvokeConfigInvoker {
	requestDef := GenReqDefForDeleteFunctionAsyncInvokeConfig()
	return &DeleteFunctionAsyncInvokeConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVersionAlias 删除函数版本别名。
//
// 删除函数版本别名。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) DeleteVersionAlias(request *model.DeleteVersionAliasRequest) (*model.DeleteVersionAliasResponse, error) {
	requestDef := GenReqDefForDeleteVersionAlias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVersionAliasResponse), nil
	}
}

// DeleteVersionAliasInvoker 删除函数版本别名。
func (c *FunctionGraphClient) DeleteVersionAliasInvoker(request *model.DeleteVersionAliasRequest) *DeleteVersionAliasInvoker {
	requestDef := GenReqDefForDeleteVersionAlias()
	return &DeleteVersionAliasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableLtsLogs 开通lts日志上报功能。
//
// 开通lts日志上报功能。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) EnableLtsLogs(request *model.EnableLtsLogsRequest) (*model.EnableLtsLogsResponse, error) {
	requestDef := GenReqDefForEnableLtsLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableLtsLogsResponse), nil
	}
}

// EnableLtsLogsInvoker 开通lts日志上报功能。
func (c *FunctionGraphClient) EnableLtsLogsInvoker(request *model.EnableLtsLogsRequest) *EnableLtsLogsInvoker {
	requestDef := GenReqDefForEnableLtsLogs()
	return &EnableLtsLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportFunction 导出函数。
//
// 导出函数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ExportFunction(request *model.ExportFunctionRequest) (*model.ExportFunctionResponse, error) {
	requestDef := GenReqDefForExportFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportFunctionResponse), nil
	}
}

// ExportFunctionInvoker 导出函数。
func (c *FunctionGraphClient) ExportFunctionInvoker(request *model.ExportFunctionRequest) *ExportFunctionInvoker {
	requestDef := GenReqDefForExportFunction()
	return &ExportFunctionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportFunction 导入函数。
//
// 导入函数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ImportFunction(request *model.ImportFunctionRequest) (*model.ImportFunctionResponse, error) {
	requestDef := GenReqDefForImportFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportFunctionResponse), nil
	}
}

// ImportFunctionInvoker 导入函数。
func (c *FunctionGraphClient) ImportFunctionInvoker(request *model.ImportFunctionRequest) *ImportFunctionInvoker {
	requestDef := GenReqDefForImportFunction()
	return &ImportFunctionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InvokeFunction 同步执行函数。
//
// 同步调用指的是客户端请求需要明确等到响应结果，也就是说这样的请求必须得调用到用户的函数，并且等到调用完成才返回。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) InvokeFunction(request *model.InvokeFunctionRequest) (*model.InvokeFunctionResponse, error) {
	requestDef := GenReqDefForInvokeFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InvokeFunctionResponse), nil
	}
}

// InvokeFunctionInvoker 同步执行函数。
func (c *FunctionGraphClient) InvokeFunctionInvoker(request *model.InvokeFunctionRequest) *InvokeFunctionInvoker {
	requestDef := GenReqDefForInvokeFunction()
	return &InvokeFunctionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDependencies 获取依赖包列表
//
// 获取依赖包列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ListDependencies(request *model.ListDependenciesRequest) (*model.ListDependenciesResponse, error) {
	requestDef := GenReqDefForListDependencies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDependenciesResponse), nil
	}
}

// ListDependenciesInvoker 获取依赖包列表
func (c *FunctionGraphClient) ListDependenciesInvoker(request *model.ListDependenciesRequest) *ListDependenciesInvoker {
	requestDef := GenReqDefForListDependencies()
	return &ListDependenciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEvents 获取测试事件列表
//
// 获取指定函数的测试事件列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ListEvents(request *model.ListEventsRequest) (*model.ListEventsResponse, error) {
	requestDef := GenReqDefForListEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventsResponse), nil
	}
}

// ListEventsInvoker 获取测试事件列表
func (c *FunctionGraphClient) ListEventsInvoker(request *model.ListEventsRequest) *ListEventsInvoker {
	requestDef := GenReqDefForListEvents()
	return &ListEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFunctionAsyncInvocations 获取函数异步调用请求列表
//
// 获取函数异步调用请求列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ListFunctionAsyncInvocations(request *model.ListFunctionAsyncInvocationsRequest) (*model.ListFunctionAsyncInvocationsResponse, error) {
	requestDef := GenReqDefForListFunctionAsyncInvocations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionAsyncInvocationsResponse), nil
	}
}

// ListFunctionAsyncInvocationsInvoker 获取函数异步调用请求列表
func (c *FunctionGraphClient) ListFunctionAsyncInvocationsInvoker(request *model.ListFunctionAsyncInvocationsRequest) *ListFunctionAsyncInvocationsInvoker {
	requestDef := GenReqDefForListFunctionAsyncInvocations()
	return &ListFunctionAsyncInvocationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFunctionAsyncInvokeConfig 获取函数异步配置列表
//
// 获取函数异步配置列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ListFunctionAsyncInvokeConfig(request *model.ListFunctionAsyncInvokeConfigRequest) (*model.ListFunctionAsyncInvokeConfigResponse, error) {
	requestDef := GenReqDefForListFunctionAsyncInvokeConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionAsyncInvokeConfigResponse), nil
	}
}

// ListFunctionAsyncInvokeConfigInvoker 获取函数异步配置列表
func (c *FunctionGraphClient) ListFunctionAsyncInvokeConfigInvoker(request *model.ListFunctionAsyncInvokeConfigRequest) *ListFunctionAsyncInvokeConfigInvoker {
	requestDef := GenReqDefForListFunctionAsyncInvokeConfig()
	return &ListFunctionAsyncInvokeConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFunctionStatistics 获取指定时间段的函数运行指标
//
// 获取指定时间段的函数运行指标。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ListFunctionStatistics(request *model.ListFunctionStatisticsRequest) (*model.ListFunctionStatisticsResponse, error) {
	requestDef := GenReqDefForListFunctionStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionStatisticsResponse), nil
	}
}

// ListFunctionStatisticsInvoker 获取指定时间段的函数运行指标
func (c *FunctionGraphClient) ListFunctionStatisticsInvoker(request *model.ListFunctionStatisticsRequest) *ListFunctionStatisticsInvoker {
	requestDef := GenReqDefForListFunctionStatistics()
	return &ListFunctionStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFunctionVersions 获取指定函数的版本列表。
//
// 获取指定函数的版本列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ListFunctionVersions(request *model.ListFunctionVersionsRequest) (*model.ListFunctionVersionsResponse, error) {
	requestDef := GenReqDefForListFunctionVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionVersionsResponse), nil
	}
}

// ListFunctionVersionsInvoker 获取指定函数的版本列表。
func (c *FunctionGraphClient) ListFunctionVersionsInvoker(request *model.ListFunctionVersionsRequest) *ListFunctionVersionsInvoker {
	requestDef := GenReqDefForListFunctionVersions()
	return &ListFunctionVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFunctions 获取函数列表
//
// 获取函数列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ListFunctions(request *model.ListFunctionsRequest) (*model.ListFunctionsResponse, error) {
	requestDef := GenReqDefForListFunctions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionsResponse), nil
	}
}

// ListFunctionsInvoker 获取函数列表
func (c *FunctionGraphClient) ListFunctionsInvoker(request *model.ListFunctionsRequest) *ListFunctionsInvoker {
	requestDef := GenReqDefForListFunctions()
	return &ListFunctionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotas 查询租户配额
//
// 查询租户配额
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

// ListQuotasInvoker 查询租户配额
func (c *FunctionGraphClient) ListQuotasInvoker(request *model.ListQuotasRequest) *ListQuotasInvoker {
	requestDef := GenReqDefForListQuotas()
	return &ListQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStatistics 租户函数统计信息
//
// 租户函数统计信息。
//
// 返回三类的统计信息，函数格式和大小使用情况包括配额和使用量，流量报告。
// 通过查询参数filter可以进行过滤，查询参数period可以指定返回的时间段。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ListStatistics(request *model.ListStatisticsRequest) (*model.ListStatisticsResponse, error) {
	requestDef := GenReqDefForListStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStatisticsResponse), nil
	}
}

// ListStatisticsInvoker 租户函数统计信息
func (c *FunctionGraphClient) ListStatisticsInvoker(request *model.ListStatisticsRequest) *ListStatisticsInvoker {
	requestDef := GenReqDefForListStatistics()
	return &ListStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVersionAliases 获取指定函数所有版本别名列表。
//
// 获取函数版本别名列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ListVersionAliases(request *model.ListVersionAliasesRequest) (*model.ListVersionAliasesResponse, error) {
	requestDef := GenReqDefForListVersionAliases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVersionAliasesResponse), nil
	}
}

// ListVersionAliasesInvoker 获取指定函数所有版本别名列表。
func (c *FunctionGraphClient) ListVersionAliasesInvoker(request *model.ListVersionAliasesRequest) *ListVersionAliasesInvoker {
	requestDef := GenReqDefForListVersionAliases()
	return &ListVersionAliasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDependency 获取指定依赖包
//
// 获取指定依赖包。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ShowDependency(request *model.ShowDependencyRequest) (*model.ShowDependencyResponse, error) {
	requestDef := GenReqDefForShowDependency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDependencyResponse), nil
	}
}

// ShowDependencyInvoker 获取指定依赖包
func (c *FunctionGraphClient) ShowDependencyInvoker(request *model.ShowDependencyRequest) *ShowDependencyInvoker {
	requestDef := GenReqDefForShowDependency()
	return &ShowDependencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEvent 获取测试事件详细信息
//
// 获取测试事件详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ShowEvent(request *model.ShowEventRequest) (*model.ShowEventResponse, error) {
	requestDef := GenReqDefForShowEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEventResponse), nil
	}
}

// ShowEventInvoker 获取测试事件详细信息
func (c *FunctionGraphClient) ShowEventInvoker(request *model.ShowEventRequest) *ShowEventInvoker {
	requestDef := GenReqDefForShowEvent()
	return &ShowEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFunctionAsyncInvokeConfig 获取函数异步配置信息。
//
// 获取函数异步配置信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ShowFunctionAsyncInvokeConfig(request *model.ShowFunctionAsyncInvokeConfigRequest) (*model.ShowFunctionAsyncInvokeConfigResponse, error) {
	requestDef := GenReqDefForShowFunctionAsyncInvokeConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFunctionAsyncInvokeConfigResponse), nil
	}
}

// ShowFunctionAsyncInvokeConfigInvoker 获取函数异步配置信息。
func (c *FunctionGraphClient) ShowFunctionAsyncInvokeConfigInvoker(request *model.ShowFunctionAsyncInvokeConfigRequest) *ShowFunctionAsyncInvokeConfigInvoker {
	requestDef := GenReqDefForShowFunctionAsyncInvokeConfig()
	return &ShowFunctionAsyncInvokeConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFunctionCode 获取指定函数代码。
//
// 获取指定函数的代码。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ShowFunctionCode(request *model.ShowFunctionCodeRequest) (*model.ShowFunctionCodeResponse, error) {
	requestDef := GenReqDefForShowFunctionCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFunctionCodeResponse), nil
	}
}

// ShowFunctionCodeInvoker 获取指定函数代码。
func (c *FunctionGraphClient) ShowFunctionCodeInvoker(request *model.ShowFunctionCodeRequest) *ShowFunctionCodeInvoker {
	requestDef := GenReqDefForShowFunctionCode()
	return &ShowFunctionCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFunctionConfig 获取函数的metadata。
//
// 获取指定函数的metadata。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ShowFunctionConfig(request *model.ShowFunctionConfigRequest) (*model.ShowFunctionConfigResponse, error) {
	requestDef := GenReqDefForShowFunctionConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFunctionConfigResponse), nil
	}
}

// ShowFunctionConfigInvoker 获取函数的metadata。
func (c *FunctionGraphClient) ShowFunctionConfigInvoker(request *model.ShowFunctionConfigRequest) *ShowFunctionConfigInvoker {
	requestDef := GenReqDefForShowFunctionConfig()
	return &ShowFunctionConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLtsLogDetails 获取指定函数的lts日志组日志流配置。
//
// 获取指定函数的lts日志组日志流配置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ShowLtsLogDetails(request *model.ShowLtsLogDetailsRequest) (*model.ShowLtsLogDetailsResponse, error) {
	requestDef := GenReqDefForShowLtsLogDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLtsLogDetailsResponse), nil
	}
}

// ShowLtsLogDetailsInvoker 获取指定函数的lts日志组日志流配置。
func (c *FunctionGraphClient) ShowLtsLogDetailsInvoker(request *model.ShowLtsLogDetailsRequest) *ShowLtsLogDetailsInvoker {
	requestDef := GenReqDefForShowLtsLogDetails()
	return &ShowLtsLogDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTracing 获取函数调用链配置
//
// 获取函数调用链配置
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ShowTracing(request *model.ShowTracingRequest) (*model.ShowTracingResponse, error) {
	requestDef := GenReqDefForShowTracing()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTracingResponse), nil
	}
}

// ShowTracingInvoker 获取函数调用链配置
func (c *FunctionGraphClient) ShowTracingInvoker(request *model.ShowTracingRequest) *ShowTracingInvoker {
	requestDef := GenReqDefForShowTracing()
	return &ShowTracingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVersionAlias 获取函数版本的指定别名信息。
//
// 获取函数指定的版本别名信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ShowVersionAlias(request *model.ShowVersionAliasRequest) (*model.ShowVersionAliasResponse, error) {
	requestDef := GenReqDefForShowVersionAlias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVersionAliasResponse), nil
	}
}

// ShowVersionAliasInvoker 获取函数版本的指定别名信息。
func (c *FunctionGraphClient) ShowVersionAliasInvoker(request *model.ShowVersionAliasRequest) *ShowVersionAliasInvoker {
	requestDef := GenReqDefForShowVersionAlias()
	return &ShowVersionAliasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDependency 更新依赖包指定依赖包
//
// 更新依赖包指定依赖包。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateDependency(request *model.UpdateDependencyRequest) (*model.UpdateDependencyResponse, error) {
	requestDef := GenReqDefForUpdateDependency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDependencyResponse), nil
	}
}

// UpdateDependencyInvoker 更新依赖包指定依赖包
func (c *FunctionGraphClient) UpdateDependencyInvoker(request *model.UpdateDependencyRequest) *UpdateDependencyInvoker {
	requestDef := GenReqDefForUpdateDependency()
	return &UpdateDependencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEvent 更新测试事件
//
// 更新测试事件。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateEvent(request *model.UpdateEventRequest) (*model.UpdateEventResponse, error) {
	requestDef := GenReqDefForUpdateEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEventResponse), nil
	}
}

// UpdateEventInvoker 更新测试事件
func (c *FunctionGraphClient) UpdateEventInvoker(request *model.UpdateEventRequest) *UpdateEventInvoker {
	requestDef := GenReqDefForUpdateEvent()
	return &UpdateEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFunctionAsyncInvokeConfig 设置函数异步配置信息。
//
// 设置函数异步配置信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateFunctionAsyncInvokeConfig(request *model.UpdateFunctionAsyncInvokeConfigRequest) (*model.UpdateFunctionAsyncInvokeConfigResponse, error) {
	requestDef := GenReqDefForUpdateFunctionAsyncInvokeConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFunctionAsyncInvokeConfigResponse), nil
	}
}

// UpdateFunctionAsyncInvokeConfigInvoker 设置函数异步配置信息。
func (c *FunctionGraphClient) UpdateFunctionAsyncInvokeConfigInvoker(request *model.UpdateFunctionAsyncInvokeConfigRequest) *UpdateFunctionAsyncInvokeConfigInvoker {
	requestDef := GenReqDefForUpdateFunctionAsyncInvokeConfig()
	return &UpdateFunctionAsyncInvokeConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFunctionCode 修改函数代码。
//
// 修改指定的函数的代码。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateFunctionCode(request *model.UpdateFunctionCodeRequest) (*model.UpdateFunctionCodeResponse, error) {
	requestDef := GenReqDefForUpdateFunctionCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFunctionCodeResponse), nil
	}
}

// UpdateFunctionCodeInvoker 修改函数代码。
func (c *FunctionGraphClient) UpdateFunctionCodeInvoker(request *model.UpdateFunctionCodeRequest) *UpdateFunctionCodeInvoker {
	requestDef := GenReqDefForUpdateFunctionCode()
	return &UpdateFunctionCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFunctionConfig 修改函数的metadata信息。
//
// 修改指定的函数的metadata信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateFunctionConfig(request *model.UpdateFunctionConfigRequest) (*model.UpdateFunctionConfigResponse, error) {
	requestDef := GenReqDefForUpdateFunctionConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFunctionConfigResponse), nil
	}
}

// UpdateFunctionConfigInvoker 修改函数的metadata信息。
func (c *FunctionGraphClient) UpdateFunctionConfigInvoker(request *model.UpdateFunctionConfigRequest) *UpdateFunctionConfigInvoker {
	requestDef := GenReqDefForUpdateFunctionConfig()
	return &UpdateFunctionConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFunctionReservedInstances 更新函数预留实例个数
//
// 为函数绑定预留实例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateFunctionReservedInstances(request *model.UpdateFunctionReservedInstancesRequest) (*model.UpdateFunctionReservedInstancesResponse, error) {
	requestDef := GenReqDefForUpdateFunctionReservedInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFunctionReservedInstancesResponse), nil
	}
}

// UpdateFunctionReservedInstancesInvoker 更新函数预留实例个数
func (c *FunctionGraphClient) UpdateFunctionReservedInstancesInvoker(request *model.UpdateFunctionReservedInstancesRequest) *UpdateFunctionReservedInstancesInvoker {
	requestDef := GenReqDefForUpdateFunctionReservedInstances()
	return &UpdateFunctionReservedInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTracing 修改函数调用链配置
//
// 修改函数调用链配置,开通/修改传入aksk，关闭aksk传空
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateTracing(request *model.UpdateTracingRequest) (*model.UpdateTracingResponse, error) {
	requestDef := GenReqDefForUpdateTracing()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTracingResponse), nil
	}
}

// UpdateTracingInvoker 修改函数调用链配置
func (c *FunctionGraphClient) UpdateTracingInvoker(request *model.UpdateTracingRequest) *UpdateTracingInvoker {
	requestDef := GenReqDefForUpdateTracing()
	return &UpdateTracingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVersionAlias 修改函数版本别名信息。
//
// 修改函数版本别名信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateVersionAlias(request *model.UpdateVersionAliasRequest) (*model.UpdateVersionAliasResponse, error) {
	requestDef := GenReqDefForUpdateVersionAlias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVersionAliasResponse), nil
	}
}

// UpdateVersionAliasInvoker 修改函数版本别名信息。
func (c *FunctionGraphClient) UpdateVersionAliasInvoker(request *model.UpdateVersionAliasRequest) *UpdateVersionAliasInvoker {
	requestDef := GenReqDefForUpdateVersionAlias()
	return &UpdateVersionAliasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteFunctionTriggers 删除指定函数的所有触发器。
//
// 删除指定函数所有触发器设置。
//
// 在提供函数版本且非latest的情况下，删除对应函数版本的触发器。
// 在提供函数别名的情况下，删除对应函数别名的触发器。
// 在不提供函数版本（也不提供别名）或版本为latest的情况下，删除该函数所有的触发器（包括所有版本和别名）。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) BatchDeleteFunctionTriggers(request *model.BatchDeleteFunctionTriggersRequest) (*model.BatchDeleteFunctionTriggersResponse, error) {
	requestDef := GenReqDefForBatchDeleteFunctionTriggers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteFunctionTriggersResponse), nil
	}
}

// BatchDeleteFunctionTriggersInvoker 删除指定函数的所有触发器。
func (c *FunctionGraphClient) BatchDeleteFunctionTriggersInvoker(request *model.BatchDeleteFunctionTriggersRequest) *BatchDeleteFunctionTriggersInvoker {
	requestDef := GenReqDefForBatchDeleteFunctionTriggers()
	return &BatchDeleteFunctionTriggersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFunctionTrigger 创建触发器。
//
// 创建触发器。
//
// - 可以创建的触发器类型包括TIMER、APIG、CTS、DDS、DMS、DIS、LTS、OBS、SMN、KAFKA。
// - DDS和KAFKA触发器创建时默认为DISABLED状态，其他触发器默认为ACTIVE状态。
// - TIMER、DDS、DMS、KAFKA、LTS触发器支持禁用，其他触发器不支持。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) CreateFunctionTrigger(request *model.CreateFunctionTriggerRequest) (*model.CreateFunctionTriggerResponse, error) {
	requestDef := GenReqDefForCreateFunctionTrigger()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFunctionTriggerResponse), nil
	}
}

// CreateFunctionTriggerInvoker 创建触发器。
func (c *FunctionGraphClient) CreateFunctionTriggerInvoker(request *model.CreateFunctionTriggerRequest) *CreateFunctionTriggerInvoker {
	requestDef := GenReqDefForCreateFunctionTrigger()
	return &CreateFunctionTriggerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFunctionTrigger 删除触发器。
//
// 删除触发器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) DeleteFunctionTrigger(request *model.DeleteFunctionTriggerRequest) (*model.DeleteFunctionTriggerResponse, error) {
	requestDef := GenReqDefForDeleteFunctionTrigger()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFunctionTriggerResponse), nil
	}
}

// DeleteFunctionTriggerInvoker 删除触发器。
func (c *FunctionGraphClient) DeleteFunctionTriggerInvoker(request *model.DeleteFunctionTriggerRequest) *DeleteFunctionTriggerInvoker {
	requestDef := GenReqDefForDeleteFunctionTrigger()
	return &DeleteFunctionTriggerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFunctionTriggers 获取指定函数的所有触发器。
//
// 获取指定函数的所有触发器设置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ListFunctionTriggers(request *model.ListFunctionTriggersRequest) (*model.ListFunctionTriggersResponse, error) {
	requestDef := GenReqDefForListFunctionTriggers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionTriggersResponse), nil
	}
}

// ListFunctionTriggersInvoker 获取指定函数的所有触发器。
func (c *FunctionGraphClient) ListFunctionTriggersInvoker(request *model.ListFunctionTriggersRequest) *ListFunctionTriggersInvoker {
	requestDef := GenReqDefForListFunctionTriggers()
	return &ListFunctionTriggersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFunctionTrigger 获取指定触发器的信息。
//
// 获取特定触发器的信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ShowFunctionTrigger(request *model.ShowFunctionTriggerRequest) (*model.ShowFunctionTriggerResponse, error) {
	requestDef := GenReqDefForShowFunctionTrigger()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFunctionTriggerResponse), nil
	}
}

// ShowFunctionTriggerInvoker 获取指定触发器的信息。
func (c *FunctionGraphClient) ShowFunctionTriggerInvoker(request *model.ShowFunctionTriggerRequest) *ShowFunctionTriggerInvoker {
	requestDef := GenReqDefForShowFunctionTrigger()
	return &ShowFunctionTriggerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTrigger 更新触发器
//
// 更新触发器
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateTrigger(request *model.UpdateTriggerRequest) (*model.UpdateTriggerResponse, error) {
	requestDef := GenReqDefForUpdateTrigger()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTriggerResponse), nil
	}
}

// UpdateTriggerInvoker 更新触发器
func (c *FunctionGraphClient) UpdateTriggerInvoker(request *model.UpdateTriggerRequest) *UpdateTriggerInvoker {
	requestDef := GenReqDefForUpdateTrigger()
	return &UpdateTriggerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteWorkflows 删除工作流列表
//
// 删除工作流列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) BatchDeleteWorkflows(request *model.BatchDeleteWorkflowsRequest) (*model.BatchDeleteWorkflowsResponse, error) {
	requestDef := GenReqDefForBatchDeleteWorkflows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteWorkflowsResponse), nil
	}
}

// BatchDeleteWorkflowsInvoker 删除工作流列表
func (c *FunctionGraphClient) BatchDeleteWorkflowsInvoker(request *model.BatchDeleteWorkflowsRequest) *BatchDeleteWorkflowsInvoker {
	requestDef := GenReqDefForBatchDeleteWorkflows()
	return &BatchDeleteWorkflowsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkflow 创建工作流列表
//
// 创建工作流列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) CreateWorkflow(request *model.CreateWorkflowRequest) (*model.CreateWorkflowResponse, error) {
	requestDef := GenReqDefForCreateWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkflowResponse), nil
	}
}

// CreateWorkflowInvoker 创建工作流列表
func (c *FunctionGraphClient) CreateWorkflowInvoker(request *model.CreateWorkflowRequest) *CreateWorkflowInvoker {
	requestDef := GenReqDefForCreateWorkflow()
	return &CreateWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkflowExecutions 获取指定函数流执行实例列表
//
// 获取指定函数流执行实例列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ListWorkflowExecutions(request *model.ListWorkflowExecutionsRequest) (*model.ListWorkflowExecutionsResponse, error) {
	requestDef := GenReqDefForListWorkflowExecutions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkflowExecutionsResponse), nil
	}
}

// ListWorkflowExecutionsInvoker 获取指定函数流执行实例列表
func (c *FunctionGraphClient) ListWorkflowExecutionsInvoker(request *model.ListWorkflowExecutionsRequest) *ListWorkflowExecutionsInvoker {
	requestDef := GenReqDefForListWorkflowExecutions()
	return &ListWorkflowExecutionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkflows 查询工作流列表
//
// 查询工作流列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ListWorkflows(request *model.ListWorkflowsRequest) (*model.ListWorkflowsResponse, error) {
	requestDef := GenReqDefForListWorkflows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkflowsResponse), nil
	}
}

// ListWorkflowsInvoker 查询工作流列表
func (c *FunctionGraphClient) ListWorkflowsInvoker(request *model.ListWorkflowsRequest) *ListWorkflowsInvoker {
	requestDef := GenReqDefForListWorkflows()
	return &ListWorkflowsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetryWorkFlow 重试工作流
//
// 重试工作流
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) RetryWorkFlow(request *model.RetryWorkFlowRequest) (*model.RetryWorkFlowResponse, error) {
	requestDef := GenReqDefForRetryWorkFlow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryWorkFlowResponse), nil
	}
}

// RetryWorkFlowInvoker 重试工作流
func (c *FunctionGraphClient) RetryWorkFlowInvoker(request *model.RetryWorkFlowRequest) *RetryWorkFlowInvoker {
	requestDef := GenReqDefForRetryWorkFlow()
	return &RetryWorkFlowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTenantMetric 获取函数流指标
//
// 获取函数流指标
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ShowTenantMetric(request *model.ShowTenantMetricRequest) (*model.ShowTenantMetricResponse, error) {
	requestDef := GenReqDefForShowTenantMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTenantMetricResponse), nil
	}
}

// ShowTenantMetricInvoker 获取函数流指标
func (c *FunctionGraphClient) ShowTenantMetricInvoker(request *model.ShowTenantMetricRequest) *ShowTenantMetricInvoker {
	requestDef := GenReqDefForShowTenantMetric()
	return &ShowTenantMetricInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkFlow 获取指定函数流实例
//
// 获取指定函数流实例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ShowWorkFlow(request *model.ShowWorkFlowRequest) (*model.ShowWorkFlowResponse, error) {
	requestDef := GenReqDefForShowWorkFlow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkFlowResponse), nil
	}
}

// ShowWorkFlowInvoker 获取指定函数流实例
func (c *FunctionGraphClient) ShowWorkFlowInvoker(request *model.ShowWorkFlowRequest) *ShowWorkFlowInvoker {
	requestDef := GenReqDefForShowWorkFlow()
	return &ShowWorkFlowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkFlowMetric 获取指定工作流指标
//
// 获取指定工作流指标
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ShowWorkFlowMetric(request *model.ShowWorkFlowMetricRequest) (*model.ShowWorkFlowMetricResponse, error) {
	requestDef := GenReqDefForShowWorkFlowMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkFlowMetricResponse), nil
	}
}

// ShowWorkFlowMetricInvoker 获取指定工作流指标
func (c *FunctionGraphClient) ShowWorkFlowMetricInvoker(request *model.ShowWorkFlowMetricRequest) *ShowWorkFlowMetricInvoker {
	requestDef := GenReqDefForShowWorkFlowMetric()
	return &ShowWorkFlowMetricInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkflowExecution 获取指定函数流执行实例
//
// 获取指定函数流执行实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) ShowWorkflowExecution(request *model.ShowWorkflowExecutionRequest) (*model.ShowWorkflowExecutionResponse, error) {
	requestDef := GenReqDefForShowWorkflowExecution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkflowExecutionResponse), nil
	}
}

// ShowWorkflowExecutionInvoker 获取指定函数流执行实例
func (c *FunctionGraphClient) ShowWorkflowExecutionInvoker(request *model.ShowWorkflowExecutionRequest) *ShowWorkflowExecutionInvoker {
	requestDef := GenReqDefForShowWorkflowExecution()
	return &ShowWorkflowExecutionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartSyncWorkflowExecution 同步执行函数流
//
// 同步执行函数流
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) StartSyncWorkflowExecution(request *model.StartSyncWorkflowExecutionRequest) (*model.StartSyncWorkflowExecutionResponse, error) {
	requestDef := GenReqDefForStartSyncWorkflowExecution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartSyncWorkflowExecutionResponse), nil
	}
}

// StartSyncWorkflowExecutionInvoker 同步执行函数流
func (c *FunctionGraphClient) StartSyncWorkflowExecutionInvoker(request *model.StartSyncWorkflowExecutionRequest) *StartSyncWorkflowExecutionInvoker {
	requestDef := GenReqDefForStartSyncWorkflowExecution()
	return &StartSyncWorkflowExecutionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartWorkflowExecution 开始执行函数流
//
// 开始执行函数流
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) StartWorkflowExecution(request *model.StartWorkflowExecutionRequest) (*model.StartWorkflowExecutionResponse, error) {
	requestDef := GenReqDefForStartWorkflowExecution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartWorkflowExecutionResponse), nil
	}
}

// StartWorkflowExecutionInvoker 开始执行函数流
func (c *FunctionGraphClient) StartWorkflowExecutionInvoker(request *model.StartWorkflowExecutionRequest) *StartWorkflowExecutionInvoker {
	requestDef := GenReqDefForStartWorkflowExecution()
	return &StartWorkflowExecutionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopWorkFlow 停止工作流
//
// 停止工作流
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) StopWorkFlow(request *model.StopWorkFlowRequest) (*model.StopWorkFlowResponse, error) {
	requestDef := GenReqDefForStopWorkFlow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopWorkFlowResponse), nil
	}
}

// StopWorkFlowInvoker 停止工作流
func (c *FunctionGraphClient) StopWorkFlowInvoker(request *model.StopWorkFlowRequest) *StopWorkFlowInvoker {
	requestDef := GenReqDefForStopWorkFlow()
	return &StopWorkFlowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkFlow 修改指定函数流实例
//
// 修改指定函数流实例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateWorkFlow(request *model.UpdateWorkFlowRequest) (*model.UpdateWorkFlowResponse, error) {
	requestDef := GenReqDefForUpdateWorkFlow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkFlowResponse), nil
	}
}

// UpdateWorkFlowInvoker 修改指定函数流实例
func (c *FunctionGraphClient) UpdateWorkFlowInvoker(request *model.UpdateWorkFlowRequest) *UpdateWorkFlowInvoker {
	requestDef := GenReqDefForUpdateWorkFlow()
	return &UpdateWorkFlowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
