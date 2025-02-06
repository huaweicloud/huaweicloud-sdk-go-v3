package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/functiongraph/v2/model"
)

type FunctionGraphClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewFunctionGraphClient(hcClient *httpclient.HcHttpClient) *FunctionGraphClient {
	return &FunctionGraphClient{HcClient: hcClient}
}

func FunctionGraphClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AsyncInvokeFunction 异步执行函数
//
// 异步执行函数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) AsyncInvokeFunction(request *model.AsyncInvokeFunctionRequest) (*model.AsyncInvokeFunctionResponse, error) {
	requestDef := GenReqDefForAsyncInvokeFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AsyncInvokeFunctionResponse), nil
	}
}

// AsyncInvokeFunctionInvoker 异步执行函数
func (c *FunctionGraphClient) AsyncInvokeFunctionInvoker(request *model.AsyncInvokeFunctionRequest) *AsyncInvokeFunctionInvoker {
	requestDef := GenReqDefForAsyncInvokeFunction()
	return &AsyncInvokeFunctionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteFunctionTriggers 删除指定函数的所有触发器
//
// 删除指定函数所有触发器设置。
//
// 在提供函数版本且非latest的情况下，删除对应函数版本的触发器。
// 在提供函数别名的情况下，删除对应函数别名的触发器。
// 在不提供函数版本（也不提供别名）或版本为latest的情况下，删除该函数所有的触发器（包括所有版本和别名）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) BatchDeleteFunctionTriggers(request *model.BatchDeleteFunctionTriggersRequest) (*model.BatchDeleteFunctionTriggersResponse, error) {
	requestDef := GenReqDefForBatchDeleteFunctionTriggers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteFunctionTriggersResponse), nil
	}
}

// BatchDeleteFunctionTriggersInvoker 删除指定函数的所有触发器
func (c *FunctionGraphClient) BatchDeleteFunctionTriggersInvoker(request *model.BatchDeleteFunctionTriggersRequest) *BatchDeleteFunctionTriggersInvoker {
	requestDef := GenReqDefForBatchDeleteFunctionTriggers()
	return &BatchDeleteFunctionTriggersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteWorkflows 删除函数流
//
// 删除函数流
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) BatchDeleteWorkflows(request *model.BatchDeleteWorkflowsRequest) (*model.BatchDeleteWorkflowsResponse, error) {
	requestDef := GenReqDefForBatchDeleteWorkflows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteWorkflowsResponse), nil
	}
}

// BatchDeleteWorkflowsInvoker 删除函数流
func (c *FunctionGraphClient) BatchDeleteWorkflowsInvoker(request *model.BatchDeleteWorkflowsRequest) *BatchDeleteWorkflowsInvoker {
	requestDef := GenReqDefForBatchDeleteWorkflows()
	return &BatchDeleteWorkflowsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelAsyncInvocation 停止函数异步调用请求
//
// -| 当前仅支持参数recursive为false且force为true的函数。 在1：N的函数做并发异步调用的场景下调用停止异步请求接口时，同一函数实例同时在执行的其他请求也会被一并停止并返回4208 function invocation canceled
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// CreateCallbackWorkflow 回调工作流
//
// 回调工作流
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) CreateCallbackWorkflow(request *model.CreateCallbackWorkflowRequest) (*model.CreateCallbackWorkflowResponse, error) {
	requestDef := GenReqDefForCreateCallbackWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCallbackWorkflowResponse), nil
	}
}

// CreateCallbackWorkflowInvoker 回调工作流
func (c *FunctionGraphClient) CreateCallbackWorkflowInvoker(request *model.CreateCallbackWorkflowRequest) *CreateCallbackWorkflowInvoker {
	requestDef := GenReqDefForCreateCallbackWorkflow()
	return &CreateCallbackWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDependencyVersion 创建依赖包版本
//
// 创建依赖包版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) CreateDependencyVersion(request *model.CreateDependencyVersionRequest) (*model.CreateDependencyVersionResponse, error) {
	requestDef := GenReqDefForCreateDependencyVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDependencyVersionResponse), nil
	}
}

// CreateDependencyVersionInvoker 创建依赖包版本
func (c *FunctionGraphClient) CreateDependencyVersionInvoker(request *model.CreateDependencyVersionRequest) *CreateDependencyVersionInvoker {
	requestDef := GenReqDefForCreateDependencyVersion()
	return &CreateDependencyVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEvent 创建测试事件
//
// 创建测试事件
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// CreateFunction 创建函数
//
// 创建指定的函数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) CreateFunction(request *model.CreateFunctionRequest) (*model.CreateFunctionResponse, error) {
	requestDef := GenReqDefForCreateFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFunctionResponse), nil
	}
}

// CreateFunctionInvoker 创建函数
func (c *FunctionGraphClient) CreateFunctionInvoker(request *model.CreateFunctionRequest) *CreateFunctionInvoker {
	requestDef := GenReqDefForCreateFunction()
	return &CreateFunctionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFunctionApp 创建应用程序
//
// 创建应用程序（该功能目前仅支持华北-北京四、华东-上海一）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) CreateFunctionApp(request *model.CreateFunctionAppRequest) (*model.CreateFunctionAppResponse, error) {
	requestDef := GenReqDefForCreateFunctionApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFunctionAppResponse), nil
	}
}

// CreateFunctionAppInvoker 创建应用程序
func (c *FunctionGraphClient) CreateFunctionAppInvoker(request *model.CreateFunctionAppRequest) *CreateFunctionAppInvoker {
	requestDef := GenReqDefForCreateFunctionApp()
	return &CreateFunctionAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFunctionTrigger 创建触发器
//
// 创建触发器。
//
// - 可以创建的触发器类型包括TIMER、APIG、CTS、DDS、DMS、DIS、LTS、OBS、SMN、KAFKA。
// - DDS和KAFKA触发器创建时默认为DISABLED状态，其他触发器默认为ACTIVE状态。
// - TIMER、DDS、DMS、KAFKA、LTS触发器支持禁用，其他触发器不支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) CreateFunctionTrigger(request *model.CreateFunctionTriggerRequest) (*model.CreateFunctionTriggerResponse, error) {
	requestDef := GenReqDefForCreateFunctionTrigger()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFunctionTriggerResponse), nil
	}
}

// CreateFunctionTriggerInvoker 创建触发器
func (c *FunctionGraphClient) CreateFunctionTriggerInvoker(request *model.CreateFunctionTriggerRequest) *CreateFunctionTriggerInvoker {
	requestDef := GenReqDefForCreateFunctionTrigger()
	return &CreateFunctionTriggerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFunctionVersion 发布函数版本
//
// 发布函数版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) CreateFunctionVersion(request *model.CreateFunctionVersionRequest) (*model.CreateFunctionVersionResponse, error) {
	requestDef := GenReqDefForCreateFunctionVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFunctionVersionResponse), nil
	}
}

// CreateFunctionVersionInvoker 发布函数版本
func (c *FunctionGraphClient) CreateFunctionVersionInvoker(request *model.CreateFunctionVersionRequest) *CreateFunctionVersionInvoker {
	requestDef := GenReqDefForCreateFunctionVersion()
	return &CreateFunctionVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTags 创建资源标签
//
// 创建资源标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) CreateTags(request *model.CreateTagsRequest) (*model.CreateTagsResponse, error) {
	requestDef := GenReqDefForCreateTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagsResponse), nil
	}
}

// CreateTagsInvoker 创建资源标签
func (c *FunctionGraphClient) CreateTagsInvoker(request *model.CreateTagsRequest) *CreateTagsInvoker {
	requestDef := GenReqDefForCreateTags()
	return &CreateTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVersionAlias 创建函数版本别名
//
// 创建函数灰度版本别名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) CreateVersionAlias(request *model.CreateVersionAliasRequest) (*model.CreateVersionAliasResponse, error) {
	requestDef := GenReqDefForCreateVersionAlias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVersionAliasResponse), nil
	}
}

// CreateVersionAliasInvoker 创建函数版本别名
func (c *FunctionGraphClient) CreateVersionAliasInvoker(request *model.CreateVersionAliasRequest) *CreateVersionAliasInvoker {
	requestDef := GenReqDefForCreateVersionAlias()
	return &CreateVersionAliasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVpcEndpoint 创建下沉入口
//
// 创建下沉入口。（该功能目前仅支持华北-北京四、华东-上海一、华东-上海二、西南-贵阳一）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) CreateVpcEndpoint(request *model.CreateVpcEndpointRequest) (*model.CreateVpcEndpointResponse, error) {
	requestDef := GenReqDefForCreateVpcEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpcEndpointResponse), nil
	}
}

// CreateVpcEndpointInvoker 创建下沉入口
func (c *FunctionGraphClient) CreateVpcEndpointInvoker(request *model.CreateVpcEndpointRequest) *CreateVpcEndpointInvoker {
	requestDef := GenReqDefForCreateVpcEndpoint()
	return &CreateVpcEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkflow 创建函数流
//
// 创建函数流
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) CreateWorkflow(request *model.CreateWorkflowRequest) (*model.CreateWorkflowResponse, error) {
	requestDef := GenReqDefForCreateWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkflowResponse), nil
	}
}

// CreateWorkflowInvoker 创建函数流
func (c *FunctionGraphClient) CreateWorkflowInvoker(request *model.CreateWorkflowRequest) *CreateWorkflowInvoker {
	requestDef := GenReqDefForCreateWorkflow()
	return &CreateWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDependencyVersion 删除依赖包版本
//
// 删除依赖包版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) DeleteDependencyVersion(request *model.DeleteDependencyVersionRequest) (*model.DeleteDependencyVersionResponse, error) {
	requestDef := GenReqDefForDeleteDependencyVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDependencyVersionResponse), nil
	}
}

// DeleteDependencyVersionInvoker 删除依赖包版本
func (c *FunctionGraphClient) DeleteDependencyVersionInvoker(request *model.DeleteDependencyVersionRequest) *DeleteDependencyVersionInvoker {
	requestDef := GenReqDefForDeleteDependencyVersion()
	return &DeleteDependencyVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEvent 删除指定测试事件
//
// 删除指定测试事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) DeleteEvent(request *model.DeleteEventRequest) (*model.DeleteEventResponse, error) {
	requestDef := GenReqDefForDeleteEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEventResponse), nil
	}
}

// DeleteEventInvoker 删除指定测试事件
func (c *FunctionGraphClient) DeleteEventInvoker(request *model.DeleteEventRequest) *DeleteEventInvoker {
	requestDef := GenReqDefForDeleteEvent()
	return &DeleteEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFunction 删除函数/版本
//
// 删除指定的函数或者特定的版本（不允许删除latest版本）。
//
// 如果URN中包含函数版本或者别名，则删除特定的函数版本或者别名指向的版本以及该版本关联的trigger。
// 如果URN中不包含版本或者别名，则删除整个函数，包含所有版本以及别名，触发器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) DeleteFunction(request *model.DeleteFunctionRequest) (*model.DeleteFunctionResponse, error) {
	requestDef := GenReqDefForDeleteFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFunctionResponse), nil
	}
}

// DeleteFunctionInvoker 删除函数/版本
func (c *FunctionGraphClient) DeleteFunctionInvoker(request *model.DeleteFunctionRequest) *DeleteFunctionInvoker {
	requestDef := GenReqDefForDeleteFunction()
	return &DeleteFunctionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFunctionApp 删除应用程序
//
// 删除应用程序（该功能目前仅支持华北-北京四、华东-上海一）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) DeleteFunctionApp(request *model.DeleteFunctionAppRequest) (*model.DeleteFunctionAppResponse, error) {
	requestDef := GenReqDefForDeleteFunctionApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFunctionAppResponse), nil
	}
}

// DeleteFunctionAppInvoker 删除应用程序
func (c *FunctionGraphClient) DeleteFunctionAppInvoker(request *model.DeleteFunctionAppRequest) *DeleteFunctionAppInvoker {
	requestDef := GenReqDefForDeleteFunctionApp()
	return &DeleteFunctionAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFunctionAsyncInvokeConfig 删除函数异步配置信息
//
// 删除函数异步配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) DeleteFunctionAsyncInvokeConfig(request *model.DeleteFunctionAsyncInvokeConfigRequest) (*model.DeleteFunctionAsyncInvokeConfigResponse, error) {
	requestDef := GenReqDefForDeleteFunctionAsyncInvokeConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFunctionAsyncInvokeConfigResponse), nil
	}
}

// DeleteFunctionAsyncInvokeConfigInvoker 删除函数异步配置信息
func (c *FunctionGraphClient) DeleteFunctionAsyncInvokeConfigInvoker(request *model.DeleteFunctionAsyncInvokeConfigRequest) *DeleteFunctionAsyncInvokeConfigInvoker {
	requestDef := GenReqDefForDeleteFunctionAsyncInvokeConfig()
	return &DeleteFunctionAsyncInvokeConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFunctionTrigger 删除触发器
//
// 删除触发器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) DeleteFunctionTrigger(request *model.DeleteFunctionTriggerRequest) (*model.DeleteFunctionTriggerResponse, error) {
	requestDef := GenReqDefForDeleteFunctionTrigger()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFunctionTriggerResponse), nil
	}
}

// DeleteFunctionTriggerInvoker 删除触发器
func (c *FunctionGraphClient) DeleteFunctionTriggerInvoker(request *model.DeleteFunctionTriggerRequest) *DeleteFunctionTriggerInvoker {
	requestDef := GenReqDefForDeleteFunctionTrigger()
	return &DeleteFunctionTriggerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTags 删除资源标签
//
// 删除资源标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) DeleteTags(request *model.DeleteTagsRequest) (*model.DeleteTagsResponse, error) {
	requestDef := GenReqDefForDeleteTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagsResponse), nil
	}
}

// DeleteTagsInvoker 删除资源标签
func (c *FunctionGraphClient) DeleteTagsInvoker(request *model.DeleteTagsRequest) *DeleteTagsInvoker {
	requestDef := GenReqDefForDeleteTags()
	return &DeleteTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVersionAlias 删除函数版本别名
//
// 删除函数版本别名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) DeleteVersionAlias(request *model.DeleteVersionAliasRequest) (*model.DeleteVersionAliasResponse, error) {
	requestDef := GenReqDefForDeleteVersionAlias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVersionAliasResponse), nil
	}
}

// DeleteVersionAliasInvoker 删除函数版本别名
func (c *FunctionGraphClient) DeleteVersionAliasInvoker(request *model.DeleteVersionAliasRequest) *DeleteVersionAliasInvoker {
	requestDef := GenReqDefForDeleteVersionAlias()
	return &DeleteVersionAliasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVpcEndpoint 删除下沉入口
//
// 删除下沉入口。（该功能目前仅支持华北-北京四、华东-上海一、华东-上海二、西南-贵阳一）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) DeleteVpcEndpoint(request *model.DeleteVpcEndpointRequest) (*model.DeleteVpcEndpointResponse, error) {
	requestDef := GenReqDefForDeleteVpcEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpcEndpointResponse), nil
	}
}

// DeleteVpcEndpointInvoker 删除下沉入口
func (c *FunctionGraphClient) DeleteVpcEndpointInvoker(request *model.DeleteVpcEndpointRequest) *DeleteVpcEndpointInvoker {
	requestDef := GenReqDefForDeleteVpcEndpoint()
	return &DeleteVpcEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableAsyncStatusLog 允许异步状态通知
//
// 允许异步状态通知。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) EnableAsyncStatusLog(request *model.EnableAsyncStatusLogRequest) (*model.EnableAsyncStatusLogResponse, error) {
	requestDef := GenReqDefForEnableAsyncStatusLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableAsyncStatusLogResponse), nil
	}
}

// EnableAsyncStatusLogInvoker 允许异步状态通知
func (c *FunctionGraphClient) EnableAsyncStatusLogInvoker(request *model.EnableAsyncStatusLogRequest) *EnableAsyncStatusLogInvoker {
	requestDef := GenReqDefForEnableAsyncStatusLog()
	return &EnableAsyncStatusLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableLtsLogs 开通lts日志上报功能
//
// 开通lts日志上报功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) EnableLtsLogs(request *model.EnableLtsLogsRequest) (*model.EnableLtsLogsResponse, error) {
	requestDef := GenReqDefForEnableLtsLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableLtsLogsResponse), nil
	}
}

// EnableLtsLogsInvoker 开通lts日志上报功能
func (c *FunctionGraphClient) EnableLtsLogsInvoker(request *model.EnableLtsLogsRequest) *EnableLtsLogsInvoker {
	requestDef := GenReqDefForEnableLtsLogs()
	return &EnableLtsLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportFunction 导出函数
//
// 导出函数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ExportFunction(request *model.ExportFunctionRequest) (*model.ExportFunctionResponse, error) {
	requestDef := GenReqDefForExportFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportFunctionResponse), nil
	}
}

// ExportFunctionInvoker 导出函数
func (c *FunctionGraphClient) ExportFunctionInvoker(request *model.ExportFunctionRequest) *ExportFunctionInvoker {
	requestDef := GenReqDefForExportFunction()
	return &ExportFunctionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportFunction 导入函数
//
// 导入函数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ImportFunction(request *model.ImportFunctionRequest) (*model.ImportFunctionResponse, error) {
	requestDef := GenReqDefForImportFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportFunctionResponse), nil
	}
}

// ImportFunctionInvoker 导入函数
func (c *FunctionGraphClient) ImportFunctionInvoker(request *model.ImportFunctionRequest) *ImportFunctionInvoker {
	requestDef := GenReqDefForImportFunction()
	return &ImportFunctionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InvokeFunction 同步执行函数
//
// 同步调用指的是客户端请求需要明确等到响应结果，也就是说这样的请求必须得调用到用户的函数，并且等到调用完成才返回。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) InvokeFunction(request *model.InvokeFunctionRequest) (*model.InvokeFunctionResponse, error) {
	requestDef := GenReqDefForInvokeFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InvokeFunctionResponse), nil
	}
}

// InvokeFunctionInvoker 同步执行函数
func (c *FunctionGraphClient) InvokeFunctionInvoker(request *model.InvokeFunctionRequest) *InvokeFunctionInvoker {
	requestDef := GenReqDefForInvokeFunction()
	return &InvokeFunctionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListActiveAsyncInvocations 获取函数活跃异步调用请求列表
//
// 获取函数异步调用活跃请求列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ListActiveAsyncInvocations(request *model.ListActiveAsyncInvocationsRequest) (*model.ListActiveAsyncInvocationsResponse, error) {
	requestDef := GenReqDefForListActiveAsyncInvocations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListActiveAsyncInvocationsResponse), nil
	}
}

// ListActiveAsyncInvocationsInvoker 获取函数活跃异步调用请求列表
func (c *FunctionGraphClient) ListActiveAsyncInvocationsInvoker(request *model.ListActiveAsyncInvocationsRequest) *ListActiveAsyncInvocationsInvoker {
	requestDef := GenReqDefForListActiveAsyncInvocations()
	return &ListActiveAsyncInvocationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppTemplates 查询应用程序模板列表
//
// 查询应用程序模板列表（该功能目前仅支持华北-北京四、华东-上海一）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ListAppTemplates(request *model.ListAppTemplatesRequest) (*model.ListAppTemplatesResponse, error) {
	requestDef := GenReqDefForListAppTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppTemplatesResponse), nil
	}
}

// ListAppTemplatesInvoker 查询应用程序模板列表
func (c *FunctionGraphClient) ListAppTemplatesInvoker(request *model.ListAppTemplatesRequest) *ListAppTemplatesInvoker {
	requestDef := GenReqDefForListAppTemplates()
	return &ListAppTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAsyncInvocations 获取函数异步调用请求列表
//
// 获取函数异步调用请求列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ListAsyncInvocations(request *model.ListAsyncInvocationsRequest) (*model.ListAsyncInvocationsResponse, error) {
	requestDef := GenReqDefForListAsyncInvocations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAsyncInvocationsResponse), nil
	}
}

// ListAsyncInvocationsInvoker 获取函数异步调用请求列表
func (c *FunctionGraphClient) ListAsyncInvocationsInvoker(request *model.ListAsyncInvocationsRequest) *ListAsyncInvocationsInvoker {
	requestDef := GenReqDefForListAsyncInvocations()
	return &ListAsyncInvocationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBridgeFunctions 获取指定函数绑定的servicebridge函数列表
//
// 获取指定函数绑定的servicebridge函数列表信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ListBridgeFunctions(request *model.ListBridgeFunctionsRequest) (*model.ListBridgeFunctionsResponse, error) {
	requestDef := GenReqDefForListBridgeFunctions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBridgeFunctionsResponse), nil
	}
}

// ListBridgeFunctionsInvoker 获取指定函数绑定的servicebridge函数列表
func (c *FunctionGraphClient) ListBridgeFunctionsInvoker(request *model.ListBridgeFunctionsRequest) *ListBridgeFunctionsInvoker {
	requestDef := GenReqDefForListBridgeFunctions()
	return &ListBridgeFunctionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBridgeVersions 获取servicebridge可用的版本
//
// 获取servicebridge可用的版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ListBridgeVersions(request *model.ListBridgeVersionsRequest) (*model.ListBridgeVersionsResponse, error) {
	requestDef := GenReqDefForListBridgeVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBridgeVersionsResponse), nil
	}
}

// ListBridgeVersionsInvoker 获取servicebridge可用的版本
func (c *FunctionGraphClient) ListBridgeVersionsInvoker(request *model.ListBridgeVersionsRequest) *ListBridgeVersionsInvoker {
	requestDef := GenReqDefForListBridgeVersions()
	return &ListBridgeVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDependencies 获取依赖包列表
//
// 获取依赖包列表
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListDependencyVersion 获取依赖包版本列表
//
// 获取依赖包版本列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ListDependencyVersion(request *model.ListDependencyVersionRequest) (*model.ListDependencyVersionResponse, error) {
	requestDef := GenReqDefForListDependencyVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDependencyVersionResponse), nil
	}
}

// ListDependencyVersionInvoker 获取依赖包版本列表
func (c *FunctionGraphClient) ListDependencyVersionInvoker(request *model.ListDependencyVersionRequest) *ListDependencyVersionInvoker {
	requestDef := GenReqDefForListDependencyVersion()
	return &ListDependencyVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEvents 获取指定函数的测试事件列表
//
// 获取指定函数的测试事件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ListEvents(request *model.ListEventsRequest) (*model.ListEventsResponse, error) {
	requestDef := GenReqDefForListEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventsResponse), nil
	}
}

// ListEventsInvoker 获取指定函数的测试事件列表
func (c *FunctionGraphClient) ListEventsInvoker(request *model.ListEventsRequest) *ListEventsInvoker {
	requestDef := GenReqDefForListEvents()
	return &ListEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFunctionApplications 查询应用程序列表
//
// 查询应用程序列表（该功能目前仅支持华北-北京四、华东-上海一）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ListFunctionApplications(request *model.ListFunctionApplicationsRequest) (*model.ListFunctionApplicationsResponse, error) {
	requestDef := GenReqDefForListFunctionApplications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionApplicationsResponse), nil
	}
}

// ListFunctionApplicationsInvoker 查询应用程序列表
func (c *FunctionGraphClient) ListFunctionApplicationsInvoker(request *model.ListFunctionApplicationsRequest) *ListFunctionApplicationsInvoker {
	requestDef := GenReqDefForListFunctionApplications()
	return &ListFunctionApplicationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFunctionAsMetric 获取按指定指标排序的函数列表
//
// 按指定指标排序的函数列表。
//
// 默认统计按错误次数指标统计最近一天失败次数最多的前10个函数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ListFunctionAsMetric(request *model.ListFunctionAsMetricRequest) (*model.ListFunctionAsMetricResponse, error) {
	requestDef := GenReqDefForListFunctionAsMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionAsMetricResponse), nil
	}
}

// ListFunctionAsMetricInvoker 获取按指定指标排序的函数列表
func (c *FunctionGraphClient) ListFunctionAsMetricInvoker(request *model.ListFunctionAsMetricRequest) *ListFunctionAsMetricInvoker {
	requestDef := GenReqDefForListFunctionAsMetric()
	return &ListFunctionAsMetricInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFunctionAsyncInvokeConfig 获取函数异步配置列表
//
// 获取指定函数所有版本的异步配置列表。。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListFunctionReservedInstances 获取函数预留实例数量
//
// 获取函数预留实例数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ListFunctionReservedInstances(request *model.ListFunctionReservedInstancesRequest) (*model.ListFunctionReservedInstancesResponse, error) {
	requestDef := GenReqDefForListFunctionReservedInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionReservedInstancesResponse), nil
	}
}

// ListFunctionReservedInstancesInvoker 获取函数预留实例数量
func (c *FunctionGraphClient) ListFunctionReservedInstancesInvoker(request *model.ListFunctionReservedInstancesRequest) *ListFunctionReservedInstancesInvoker {
	requestDef := GenReqDefForListFunctionReservedInstances()
	return &ListFunctionReservedInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFunctionStatistics 获取指定时间段的函数运行指标
//
// 获取指定时间段的函数运行指标。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListFunctionTags 查询函数标签列表
//
// 查询函数标签列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ListFunctionTags(request *model.ListFunctionTagsRequest) (*model.ListFunctionTagsResponse, error) {
	requestDef := GenReqDefForListFunctionTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionTagsResponse), nil
	}
}

// ListFunctionTagsInvoker 查询函数标签列表
func (c *FunctionGraphClient) ListFunctionTagsInvoker(request *model.ListFunctionTagsRequest) *ListFunctionTagsInvoker {
	requestDef := GenReqDefForListFunctionTags()
	return &ListFunctionTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFunctionTemplate 获取函数模板列表
//
// 获取函数模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ListFunctionTemplate(request *model.ListFunctionTemplateRequest) (*model.ListFunctionTemplateResponse, error) {
	requestDef := GenReqDefForListFunctionTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionTemplateResponse), nil
	}
}

// ListFunctionTemplateInvoker 获取函数模板列表
func (c *FunctionGraphClient) ListFunctionTemplateInvoker(request *model.ListFunctionTemplateRequest) *ListFunctionTemplateInvoker {
	requestDef := GenReqDefForListFunctionTemplate()
	return &ListFunctionTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFunctionTriggers 获取指定函数的所有触发器
//
// 获取指定函数的所有触发器设置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ListFunctionTriggers(request *model.ListFunctionTriggersRequest) (*model.ListFunctionTriggersResponse, error) {
	requestDef := GenReqDefForListFunctionTriggers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionTriggersResponse), nil
	}
}

// ListFunctionTriggersInvoker 获取指定函数的所有触发器
func (c *FunctionGraphClient) ListFunctionTriggersInvoker(request *model.ListFunctionTriggersRequest) *ListFunctionTriggersInvoker {
	requestDef := GenReqDefForListFunctionTriggers()
	return &ListFunctionTriggersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFunctionVersions 获取指定函数的版本列表
//
// 获取指定函数的版本列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ListFunctionVersions(request *model.ListFunctionVersionsRequest) (*model.ListFunctionVersionsResponse, error) {
	requestDef := GenReqDefForListFunctionVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionVersionsResponse), nil
	}
}

// ListFunctionVersionsInvoker 获取指定函数的版本列表
func (c *FunctionGraphClient) ListFunctionVersionsInvoker(request *model.ListFunctionVersionsRequest) *ListFunctionVersionsInvoker {
	requestDef := GenReqDefForListFunctionVersions()
	return &ListFunctionVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFunctions 获取函数列表
//
// 获取函数列表
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListReservedInstanceConfigs 获取函数预留实例配置列表
//
// 获取函数预留实例配置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ListReservedInstanceConfigs(request *model.ListReservedInstanceConfigsRequest) (*model.ListReservedInstanceConfigsResponse, error) {
	requestDef := GenReqDefForListReservedInstanceConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListReservedInstanceConfigsResponse), nil
	}
}

// ListReservedInstanceConfigsInvoker 获取函数预留实例配置列表
func (c *FunctionGraphClient) ListReservedInstanceConfigsInvoker(request *model.ListReservedInstanceConfigsRequest) *ListReservedInstanceConfigsInvoker {
	requestDef := GenReqDefForListReservedInstanceConfigs()
	return &ListReservedInstanceConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStatistics 租户函数统计信息
//
// 租户函数统计信息。
//
// 返回三类的统计信息，函数格式和大小使用情况包括配额和使用量，流量报告。
// 通过查询参数filter可以进行过滤，查询参数period可以指定返回的时间段。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListVersionAliases 获取指定函数所有版本别名列表
//
// 获取函数版本别名列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ListVersionAliases(request *model.ListVersionAliasesRequest) (*model.ListVersionAliasesResponse, error) {
	requestDef := GenReqDefForListVersionAliases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVersionAliasesResponse), nil
	}
}

// ListVersionAliasesInvoker 获取指定函数所有版本别名列表
func (c *FunctionGraphClient) ListVersionAliasesInvoker(request *model.ListVersionAliasesRequest) *ListVersionAliasesInvoker {
	requestDef := GenReqDefForListVersionAliases()
	return &ListVersionAliasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkflow 查询函数流
//
// 查询函数流
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ListWorkflow(request *model.ListWorkflowRequest) (*model.ListWorkflowResponse, error) {
	requestDef := GenReqDefForListWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkflowResponse), nil
	}
}

// ListWorkflowInvoker 查询函数流
func (c *FunctionGraphClient) ListWorkflowInvoker(request *model.ListWorkflowRequest) *ListWorkflowInvoker {
	requestDef := GenReqDefForListWorkflow()
	return &ListWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkflowExecutions 获取指定函数流执行实例列表
//
// 获取指定函数流执行实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// RetryWorkFlow 重试函数流
//
// 重试函数流
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) RetryWorkFlow(request *model.RetryWorkFlowRequest) (*model.RetryWorkFlowResponse, error) {
	requestDef := GenReqDefForRetryWorkFlow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryWorkFlowResponse), nil
	}
}

// RetryWorkFlowInvoker 重试函数流
func (c *FunctionGraphClient) RetryWorkFlowInvoker(request *model.RetryWorkFlowRequest) *RetryWorkFlowInvoker {
	requestDef := GenReqDefForRetryWorkFlow()
	return &RetryWorkFlowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppTemplate 查询应用程序模板详情
//
// 查询应用程序模板详情（该功能目前仅支持华北-北京四、华东-上海一）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ShowAppTemplate(request *model.ShowAppTemplateRequest) (*model.ShowAppTemplateResponse, error) {
	requestDef := GenReqDefForShowAppTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppTemplateResponse), nil
	}
}

// ShowAppTemplateInvoker 查询应用程序模板详情
func (c *FunctionGraphClient) ShowAppTemplateInvoker(request *model.ShowAppTemplateRequest) *ShowAppTemplateInvoker {
	requestDef := GenReqDefForShowAppTemplate()
	return &ShowAppTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDependencyVersion 获取依赖包版本详情
//
// 获取依赖包版本详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ShowDependencyVersion(request *model.ShowDependencyVersionRequest) (*model.ShowDependencyVersionResponse, error) {
	requestDef := GenReqDefForShowDependencyVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDependencyVersionResponse), nil
	}
}

// ShowDependencyVersionInvoker 获取依赖包版本详情
func (c *FunctionGraphClient) ShowDependencyVersionInvoker(request *model.ShowDependencyVersionRequest) *ShowDependencyVersionInvoker {
	requestDef := GenReqDefForShowDependencyVersion()
	return &ShowDependencyVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEvent 获取测试事件详细信息
//
// 获取测试事件详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowFuncReservedInstanceMetrics 查询函数实例使用情况指标
//
// 查询函数实例使用情况指标。
//
// - 指标单位为分钟：
//     当查询时间范围小于1小时,指标周期为1分钟
//     当查询时间范围小于1天,指标周期为30分钟
//     当查询时间范围大于1天,指标周期为180分钟
// - 指标分为如下几类：reservedinstancenum（预留实例使用）、concurrency（实例使用/并发）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ShowFuncReservedInstanceMetrics(request *model.ShowFuncReservedInstanceMetricsRequest) (*model.ShowFuncReservedInstanceMetricsResponse, error) {
	requestDef := GenReqDefForShowFuncReservedInstanceMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFuncReservedInstanceMetricsResponse), nil
	}
}

// ShowFuncReservedInstanceMetricsInvoker 查询函数实例使用情况指标
func (c *FunctionGraphClient) ShowFuncReservedInstanceMetricsInvoker(request *model.ShowFuncReservedInstanceMetricsRequest) *ShowFuncReservedInstanceMetricsInvoker {
	requestDef := GenReqDefForShowFuncReservedInstanceMetrics()
	return &ShowFuncReservedInstanceMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFuncSnapshotState 查询函数快照制作状态
//
// 查询函数快照制作状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ShowFuncSnapshotState(request *model.ShowFuncSnapshotStateRequest) (*model.ShowFuncSnapshotStateResponse, error) {
	requestDef := GenReqDefForShowFuncSnapshotState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFuncSnapshotStateResponse), nil
	}
}

// ShowFuncSnapshotStateInvoker 查询函数快照制作状态
func (c *FunctionGraphClient) ShowFuncSnapshotStateInvoker(request *model.ShowFuncSnapshotStateRequest) *ShowFuncSnapshotStateInvoker {
	requestDef := GenReqDefForShowFuncSnapshotState()
	return &ShowFuncSnapshotStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFunctionApp 查询应用程序详情
//
// 查询应用程序详情（该功能目前仅支持华北-北京四、华东-上海一）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ShowFunctionApp(request *model.ShowFunctionAppRequest) (*model.ShowFunctionAppResponse, error) {
	requestDef := GenReqDefForShowFunctionApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFunctionAppResponse), nil
	}
}

// ShowFunctionAppInvoker 查询应用程序详情
func (c *FunctionGraphClient) ShowFunctionAppInvoker(request *model.ShowFunctionAppRequest) *ShowFunctionAppInvoker {
	requestDef := GenReqDefForShowFunctionApp()
	return &ShowFunctionAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFunctionAsyncInvokeConfig 获取函数异步配置信息
//
// 获取指定函数某一版本的异步配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ShowFunctionAsyncInvokeConfig(request *model.ShowFunctionAsyncInvokeConfigRequest) (*model.ShowFunctionAsyncInvokeConfigResponse, error) {
	requestDef := GenReqDefForShowFunctionAsyncInvokeConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFunctionAsyncInvokeConfigResponse), nil
	}
}

// ShowFunctionAsyncInvokeConfigInvoker 获取函数异步配置信息
func (c *FunctionGraphClient) ShowFunctionAsyncInvokeConfigInvoker(request *model.ShowFunctionAsyncInvokeConfigRequest) *ShowFunctionAsyncInvokeConfigInvoker {
	requestDef := GenReqDefForShowFunctionAsyncInvokeConfig()
	return &ShowFunctionAsyncInvokeConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFunctionCode 获取指定函数代码
//
// 获取指定函数的代码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ShowFunctionCode(request *model.ShowFunctionCodeRequest) (*model.ShowFunctionCodeResponse, error) {
	requestDef := GenReqDefForShowFunctionCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFunctionCodeResponse), nil
	}
}

// ShowFunctionCodeInvoker 获取指定函数代码
func (c *FunctionGraphClient) ShowFunctionCodeInvoker(request *model.ShowFunctionCodeRequest) *ShowFunctionCodeInvoker {
	requestDef := GenReqDefForShowFunctionCode()
	return &ShowFunctionCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFunctionConfig 获取函数的metadata
//
// 获取指定函数的metadata。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ShowFunctionConfig(request *model.ShowFunctionConfigRequest) (*model.ShowFunctionConfigResponse, error) {
	requestDef := GenReqDefForShowFunctionConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFunctionConfigResponse), nil
	}
}

// ShowFunctionConfigInvoker 获取函数的metadata
func (c *FunctionGraphClient) ShowFunctionConfigInvoker(request *model.ShowFunctionConfigRequest) *ShowFunctionConfigInvoker {
	requestDef := GenReqDefForShowFunctionConfig()
	return &ShowFunctionConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFunctionMetrics 查询函数实例流量指标
//
// 查询函数流量指标。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ShowFunctionMetrics(request *model.ShowFunctionMetricsRequest) (*model.ShowFunctionMetricsResponse, error) {
	requestDef := GenReqDefForShowFunctionMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFunctionMetricsResponse), nil
	}
}

// ShowFunctionMetricsInvoker 查询函数实例流量指标
func (c *FunctionGraphClient) ShowFunctionMetricsInvoker(request *model.ShowFunctionMetricsRequest) *ShowFunctionMetricsInvoker {
	requestDef := GenReqDefForShowFunctionMetrics()
	return &ShowFunctionMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFunctionTemplate 获取指定函数模板
//
// 获取指定函数模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ShowFunctionTemplate(request *model.ShowFunctionTemplateRequest) (*model.ShowFunctionTemplateResponse, error) {
	requestDef := GenReqDefForShowFunctionTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFunctionTemplateResponse), nil
	}
}

// ShowFunctionTemplateInvoker 获取指定函数模板
func (c *FunctionGraphClient) ShowFunctionTemplateInvoker(request *model.ShowFunctionTemplateRequest) *ShowFunctionTemplateInvoker {
	requestDef := GenReqDefForShowFunctionTemplate()
	return &ShowFunctionTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFunctionTrigger 获取指定触发器的信息
//
// 获取特定触发器的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ShowFunctionTrigger(request *model.ShowFunctionTriggerRequest) (*model.ShowFunctionTriggerResponse, error) {
	requestDef := GenReqDefForShowFunctionTrigger()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFunctionTriggerResponse), nil
	}
}

// ShowFunctionTriggerInvoker 获取指定触发器的信息
func (c *FunctionGraphClient) ShowFunctionTriggerInvoker(request *model.ShowFunctionTriggerRequest) *ShowFunctionTriggerInvoker {
	requestDef := GenReqDefForShowFunctionTrigger()
	return &ShowFunctionTriggerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLtsLogDetails 获取指定函数的lts日志组日志流配置
//
// 获取指定函数的lts日志组日志流配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ShowLtsLogDetails(request *model.ShowLtsLogDetailsRequest) (*model.ShowLtsLogDetailsResponse, error) {
	requestDef := GenReqDefForShowLtsLogDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLtsLogDetailsResponse), nil
	}
}

// ShowLtsLogDetailsInvoker 获取指定函数的lts日志组日志流配置
func (c *FunctionGraphClient) ShowLtsLogDetailsInvoker(request *model.ShowLtsLogDetailsRequest) *ShowLtsLogDetailsInvoker {
	requestDef := GenReqDefForShowLtsLogDetails()
	return &ShowLtsLogDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectAsyncStatusLogInfo 查询异步日志详情
//
// 查询异步日志详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ShowProjectAsyncStatusLogInfo(request *model.ShowProjectAsyncStatusLogInfoRequest) (*model.ShowProjectAsyncStatusLogInfoResponse, error) {
	requestDef := GenReqDefForShowProjectAsyncStatusLogInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectAsyncStatusLogInfoResponse), nil
	}
}

// ShowProjectAsyncStatusLogInfoInvoker 查询异步日志详情
func (c *FunctionGraphClient) ShowProjectAsyncStatusLogInfoInvoker(request *model.ShowProjectAsyncStatusLogInfoRequest) *ShowProjectAsyncStatusLogInfoInvoker {
	requestDef := GenReqDefForShowProjectAsyncStatusLogInfo()
	return &ShowProjectAsyncStatusLogInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectTagsList 查询资源标签
//
// 查询资源标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ShowProjectTagsList(request *model.ShowProjectTagsListRequest) (*model.ShowProjectTagsListResponse, error) {
	requestDef := GenReqDefForShowProjectTagsList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectTagsListResponse), nil
	}
}

// ShowProjectTagsListInvoker 查询资源标签
func (c *FunctionGraphClient) ShowProjectTagsListInvoker(request *model.ShowProjectTagsListRequest) *ShowProjectTagsListInvoker {
	requestDef := GenReqDefForShowProjectTagsList()
	return &ShowProjectTagsListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResInstanceInfo 查询资源实例
//
// 查询资源实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ShowResInstanceInfo(request *model.ShowResInstanceInfoRequest) (*model.ShowResInstanceInfoResponse, error) {
	requestDef := GenReqDefForShowResInstanceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResInstanceInfoResponse), nil
	}
}

// ShowResInstanceInfoInvoker 查询资源实例
func (c *FunctionGraphClient) ShowResInstanceInfoInvoker(request *model.ShowResInstanceInfoRequest) *ShowResInstanceInfoInvoker {
	requestDef := GenReqDefForShowResInstanceInfo()
	return &ShowResInstanceInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTenantMetric 获取函数流指标
//
// 获取函数流指标
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowTracing 获取函数调用链配置
//
// 获取函数调用链配置
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowVersionAlias 获取函数版本的指定别名信息
//
// 获取函数指定的版本别名信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ShowVersionAlias(request *model.ShowVersionAliasRequest) (*model.ShowVersionAliasResponse, error) {
	requestDef := GenReqDefForShowVersionAlias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVersionAliasResponse), nil
	}
}

// ShowVersionAliasInvoker 获取函数版本的指定别名信息
func (c *FunctionGraphClient) ShowVersionAliasInvoker(request *model.ShowVersionAliasRequest) *ShowVersionAliasInvoker {
	requestDef := GenReqDefForShowVersionAlias()
	return &ShowVersionAliasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkFlow 获取指定函数流实例的元数据
//
// 获取指定函数流实例的元数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ShowWorkFlow(request *model.ShowWorkFlowRequest) (*model.ShowWorkFlowResponse, error) {
	requestDef := GenReqDefForShowWorkFlow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkFlowResponse), nil
	}
}

// ShowWorkFlowInvoker 获取指定函数流实例的元数据
func (c *FunctionGraphClient) ShowWorkFlowInvoker(request *model.ShowWorkFlowRequest) *ShowWorkFlowInvoker {
	requestDef := GenReqDefForShowWorkFlow()
	return &ShowWorkFlowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkFlowMetric 获取指定函数流指标
//
// 获取指定函数流指标
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ShowWorkFlowMetric(request *model.ShowWorkFlowMetricRequest) (*model.ShowWorkFlowMetricResponse, error) {
	requestDef := GenReqDefForShowWorkFlowMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkFlowMetricResponse), nil
	}
}

// ShowWorkFlowMetricInvoker 获取指定函数流指标
func (c *FunctionGraphClient) ShowWorkFlowMetricInvoker(request *model.ShowWorkFlowMetricRequest) *ShowWorkFlowMetricInvoker {
	requestDef := GenReqDefForShowWorkFlowMetric()
	return &ShowWorkFlowMetricInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkflowExecution 获取指定函数流执行实例
//
// 获取指定函数流执行实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowWorkflowExecutionForPage 分页获取指定函数流执行实例列表
//
// 分页获取指定函数流执行实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) ShowWorkflowExecutionForPage(request *model.ShowWorkflowExecutionForPageRequest) (*model.ShowWorkflowExecutionForPageResponse, error) {
	requestDef := GenReqDefForShowWorkflowExecutionForPage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkflowExecutionForPageResponse), nil
	}
}

// ShowWorkflowExecutionForPageInvoker 分页获取指定函数流执行实例列表
func (c *FunctionGraphClient) ShowWorkflowExecutionForPageInvoker(request *model.ShowWorkflowExecutionForPageRequest) *ShowWorkflowExecutionForPageInvoker {
	requestDef := GenReqDefForShowWorkflowExecutionForPage()
	return &ShowWorkflowExecutionForPageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartSyncWorkflowExecution 同步执行函数流
//
// 以同步执行方式启动函数流（仅快速模式函数流支持）
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// 以异步执行方式启动函数流
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// StopWorkFlow 停止函数流
//
// 停止函数流
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) StopWorkFlow(request *model.StopWorkFlowRequest) (*model.StopWorkFlowResponse, error) {
	requestDef := GenReqDefForStopWorkFlow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopWorkFlowResponse), nil
	}
}

// StopWorkFlowInvoker 停止函数流
func (c *FunctionGraphClient) StopWorkFlowInvoker(request *model.StopWorkFlowRequest) *StopWorkFlowInvoker {
	requestDef := GenReqDefForStopWorkFlow()
	return &StopWorkFlowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEvent 更新测试事件详细信息
//
// 更新测试事件详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateEvent(request *model.UpdateEventRequest) (*model.UpdateEventResponse, error) {
	requestDef := GenReqDefForUpdateEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEventResponse), nil
	}
}

// UpdateEventInvoker 更新测试事件详细信息
func (c *FunctionGraphClient) UpdateEventInvoker(request *model.UpdateEventRequest) *UpdateEventInvoker {
	requestDef := GenReqDefForUpdateEvent()
	return &UpdateEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFuncSnapshot 禁用/启动函数快照
//
// 禁用/启动函数快照，仅支持java运行时函数，且为非latest版本才能开启函数快照功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateFuncSnapshot(request *model.UpdateFuncSnapshotRequest) (*model.UpdateFuncSnapshotResponse, error) {
	requestDef := GenReqDefForUpdateFuncSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFuncSnapshotResponse), nil
	}
}

// UpdateFuncSnapshotInvoker 禁用/启动函数快照
func (c *FunctionGraphClient) UpdateFuncSnapshotInvoker(request *model.UpdateFuncSnapshotRequest) *UpdateFuncSnapshotInvoker {
	requestDef := GenReqDefForUpdateFuncSnapshot()
	return &UpdateFuncSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFunctionAsyncInvokeConfig 设置函数异步配置信息
//
// 设置函数异步配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateFunctionAsyncInvokeConfig(request *model.UpdateFunctionAsyncInvokeConfigRequest) (*model.UpdateFunctionAsyncInvokeConfigResponse, error) {
	requestDef := GenReqDefForUpdateFunctionAsyncInvokeConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFunctionAsyncInvokeConfigResponse), nil
	}
}

// UpdateFunctionAsyncInvokeConfigInvoker 设置函数异步配置信息
func (c *FunctionGraphClient) UpdateFunctionAsyncInvokeConfigInvoker(request *model.UpdateFunctionAsyncInvokeConfigRequest) *UpdateFunctionAsyncInvokeConfigInvoker {
	requestDef := GenReqDefForUpdateFunctionAsyncInvokeConfig()
	return &UpdateFunctionAsyncInvokeConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFunctionCode 修改函数代码
//
// 修改指定的函数的代码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateFunctionCode(request *model.UpdateFunctionCodeRequest) (*model.UpdateFunctionCodeResponse, error) {
	requestDef := GenReqDefForUpdateFunctionCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFunctionCodeResponse), nil
	}
}

// UpdateFunctionCodeInvoker 修改函数代码
func (c *FunctionGraphClient) UpdateFunctionCodeInvoker(request *model.UpdateFunctionCodeRequest) *UpdateFunctionCodeInvoker {
	requestDef := GenReqDefForUpdateFunctionCode()
	return &UpdateFunctionCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFunctionCollectState 更新函数置顶状态
//
// 更新函数置顶状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateFunctionCollectState(request *model.UpdateFunctionCollectStateRequest) (*model.UpdateFunctionCollectStateResponse, error) {
	requestDef := GenReqDefForUpdateFunctionCollectState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFunctionCollectStateResponse), nil
	}
}

// UpdateFunctionCollectStateInvoker 更新函数置顶状态
func (c *FunctionGraphClient) UpdateFunctionCollectStateInvoker(request *model.UpdateFunctionCollectStateRequest) *UpdateFunctionCollectStateInvoker {
	requestDef := GenReqDefForUpdateFunctionCollectState()
	return &UpdateFunctionCollectStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFunctionConfig 修改函数的metadata信息
//
// 修改指定的函数的metadata信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateFunctionConfig(request *model.UpdateFunctionConfigRequest) (*model.UpdateFunctionConfigResponse, error) {
	requestDef := GenReqDefForUpdateFunctionConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFunctionConfigResponse), nil
	}
}

// UpdateFunctionConfigInvoker 修改函数的metadata信息
func (c *FunctionGraphClient) UpdateFunctionConfigInvoker(request *model.UpdateFunctionConfigRequest) *UpdateFunctionConfigInvoker {
	requestDef := GenReqDefForUpdateFunctionConfig()
	return &UpdateFunctionConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFunctionMaxInstanceConfig 更新函数最大实例数
//
// 更新函数最大实例数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateFunctionMaxInstanceConfig(request *model.UpdateFunctionMaxInstanceConfigRequest) (*model.UpdateFunctionMaxInstanceConfigResponse, error) {
	requestDef := GenReqDefForUpdateFunctionMaxInstanceConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFunctionMaxInstanceConfigResponse), nil
	}
}

// UpdateFunctionMaxInstanceConfigInvoker 更新函数最大实例数
func (c *FunctionGraphClient) UpdateFunctionMaxInstanceConfigInvoker(request *model.UpdateFunctionMaxInstanceConfigRequest) *UpdateFunctionMaxInstanceConfigInvoker {
	requestDef := GenReqDefForUpdateFunctionMaxInstanceConfig()
	return &UpdateFunctionMaxInstanceConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFunctionReservedInstancesCount 修改函数预留实例数量
//
// 修改函数预留实例数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateFunctionReservedInstancesCount(request *model.UpdateFunctionReservedInstancesCountRequest) (*model.UpdateFunctionReservedInstancesCountResponse, error) {
	requestDef := GenReqDefForUpdateFunctionReservedInstancesCount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFunctionReservedInstancesCountResponse), nil
	}
}

// UpdateFunctionReservedInstancesCountInvoker 修改函数预留实例数量
func (c *FunctionGraphClient) UpdateFunctionReservedInstancesCountInvoker(request *model.UpdateFunctionReservedInstancesCountRequest) *UpdateFunctionReservedInstancesCountInvoker {
	requestDef := GenReqDefForUpdateFunctionReservedInstancesCount()
	return &UpdateFunctionReservedInstancesCountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTracing 修改函数调用链配置
//
// 修改函数调用链配置,开通/修改传入aksk，关闭aksk传空
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// UpdateTrigger 更新触发器
//
// 更新触发器
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// UpdateVersionAlias 修改函数版本别名信息
//
// 修改函数版本别名信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateVersionAlias(request *model.UpdateVersionAliasRequest) (*model.UpdateVersionAliasResponse, error) {
	requestDef := GenReqDefForUpdateVersionAlias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVersionAliasResponse), nil
	}
}

// UpdateVersionAliasInvoker 修改函数版本别名信息
func (c *FunctionGraphClient) UpdateVersionAliasInvoker(request *model.UpdateVersionAliasRequest) *UpdateVersionAliasInvoker {
	requestDef := GenReqDefForUpdateVersionAlias()
	return &UpdateVersionAliasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkFlow 修改指定函数流实例的元数据
//
// 修改指定函数流实例的元数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *FunctionGraphClient) UpdateWorkFlow(request *model.UpdateWorkFlowRequest) (*model.UpdateWorkFlowResponse, error) {
	requestDef := GenReqDefForUpdateWorkFlow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkFlowResponse), nil
	}
}

// UpdateWorkFlowInvoker 修改指定函数流实例的元数据
func (c *FunctionGraphClient) UpdateWorkFlowInvoker(request *model.UpdateWorkFlowRequest) *UpdateWorkFlowInvoker {
	requestDef := GenReqDefForUpdateWorkFlow()
	return &UpdateWorkFlowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
