package v1

import (
	http_client "github.com/dysodeng/huaweicloud-sdk-go-v3/core"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/aom/v1/model"
)

type AomClient struct {
	HcClient *http_client.HcHttpClient
}

func NewAomClient(hcClient *http_client.HcHttpClient) *AomClient {
	return &AomClient{HcClient: hcClient}
}

func AomClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateFastExecuteScript 快速创建并执行脚本
//
// 该接口用于创建快速执行脚本的任务，可以指定脚本类型，执行用户，脚本参数，执行机器，脚本内容，在用户指定的机器上执行脚本。（注：接口目前开放的region为：华东-苏州二零一）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) CreateFastExecuteScript(request *model.CreateFastExecuteScriptRequest) (*model.CreateFastExecuteScriptResponse, error) {
	requestDef := GenReqDefForCreateFastExecuteScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFastExecuteScriptResponse), nil
	}
}

// CreateFastExecuteScriptInvoker 快速创建并执行脚本
func (c *AomClient) CreateFastExecuteScriptInvoker(request *model.CreateFastExecuteScriptRequest) *CreateFastExecuteScriptInvoker {
	requestDef := GenReqDefForCreateFastExecuteScript()
	return &CreateFastExecuteScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkflow 创建任务
//
// 该接口用于创建工作流（任务），返回工作流详情。任务类型取决于模板名称和&#39;input&#39;参数。（注：接口目前开放的region为：华北-北京四,华东-上海一,华东-上海二,华南-广州）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) CreateWorkflow(request *model.CreateWorkflowRequest) (*model.CreateWorkflowResponse, error) {
	requestDef := GenReqDefForCreateWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkflowResponse), nil
	}
}

// CreateWorkflowInvoker 创建任务
func (c *AomClient) CreateWorkflowInvoker(request *model.CreateWorkflowRequest) *CreateWorkflowInvoker {
	requestDef := GenReqDefForCreateWorkflow()
	return &CreateWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteWorkflow 执行工作流
//
// 该接口可下发执行指定的任务。（注：接口目前开放的region为：华北-北京四,华东-上海一,华东-上海二,华南-广州）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) ExecuteWorkflow(request *model.ExecuteWorkflowRequest) (*model.ExecuteWorkflowResponse, error) {
	requestDef := GenReqDefForExecuteWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteWorkflowResponse), nil
	}
}

// ExecuteWorkflowInvoker 执行工作流
func (c *AomClient) ExecuteWorkflowInvoker(request *model.ExecuteWorkflowRequest) *ExecuteWorkflowInvoker {
	requestDef := GenReqDefForExecuteWorkflow()
	return &ExecuteWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllJobByName 作业管理主页模糊查询
//
// 该接口可查询已创建的作业，可指定作业名称和作业创建人去精确查询，返回作业列表信息。（注：接口目前开放的region为：华北-北京四,华东-上海一,华东-上海二,华南-广州）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) ListAllJobByName(request *model.ListAllJobByNameRequest) (*model.ListAllJobByNameResponse, error) {
	requestDef := GenReqDefForListAllJobByName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllJobByNameResponse), nil
	}
}

// ListAllJobByNameInvoker 作业管理主页模糊查询
func (c *AomClient) ListAllJobByNameInvoker(request *model.ListAllJobByNameRequest) *ListAllJobByNameInvoker {
	requestDef := GenReqDefForListAllJobByName()
	return &ListAllJobByNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllScriptByName 脚本查询
//
// 该接口是脚本主页查询，可指定脚本名称和脚本创建人进行精确查询，返回包含脚本基本信息的列表数据。（注：接口目前开放的region为：华北-北京四,华东-上海一,华东-上海二,华南-广州）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) ListAllScriptByName(request *model.ListAllScriptByNameRequest) (*model.ListAllScriptByNameResponse, error) {
	requestDef := GenReqDefForListAllScriptByName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllScriptByNameResponse), nil
	}
}

// ListAllScriptByNameInvoker 脚本查询
func (c *AomClient) ListAllScriptByNameInvoker(request *model.ListAllScriptByNameRequest) *ListAllScriptByNameInvoker {
	requestDef := GenReqDefForListAllScriptByName()
	return &ListAllScriptByNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllVersionByVersionId 脚本版本查询
//
// 该接口可查询指定脚本ID下的所有版本，返回该名称的脚本版本列表信息。（注：接口目前开放的region为：华北-北京四,华东-上海一,华东-上海二,华南-广州）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) ListAllVersionByVersionId(request *model.ListAllVersionByVersionIdRequest) (*model.ListAllVersionByVersionIdResponse, error) {
	requestDef := GenReqDefForListAllVersionByVersionId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllVersionByVersionIdResponse), nil
	}
}

// ListAllVersionByVersionIdInvoker 脚本版本查询
func (c *AomClient) ListAllVersionByVersionIdInvoker(request *model.ListAllVersionByVersionIdRequest) *ListAllVersionByVersionIdInvoker {
	requestDef := GenReqDefForListAllVersionByVersionId()
	return &ListAllVersionByVersionIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTemplateByJobId 根据作业id查询方案(自定义模板)列表
//
// 该接口可根据作业ID查询执行方案，分页返回执行方案列表。（注：接口目前开放的region为：华北-北京四,华东-上海一,华东-上海二,华南-广州）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) ListTemplateByJobId(request *model.ListTemplateByJobIdRequest) (*model.ListTemplateByJobIdResponse, error) {
	requestDef := GenReqDefForListTemplateByJobId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplateByJobIdResponse), nil
	}
}

// ListTemplateByJobIdInvoker 根据作业id查询方案(自定义模板)列表
func (c *AomClient) ListTemplateByJobIdInvoker(request *model.ListTemplateByJobIdRequest) *ListTemplateByJobIdInvoker {
	requestDef := GenReqDefForListTemplateByJobId()
	return &ListTemplateByJobIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkflow 查询任务列表
//
// 该接口可返回已经创建的任务列表，可按任务名称，任务状态，任务类型，执行人，更新时间为查询条件分页查询任务。（注：接口目前开放的region为：华北-北京四,华东-上海一,华东-上海二,华南-广州）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) ListWorkflow(request *model.ListWorkflowRequest) (*model.ListWorkflowResponse, error) {
	requestDef := GenReqDefForListWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkflowResponse), nil
	}
}

// ListWorkflowInvoker 查询任务列表
func (c *AomClient) ListWorkflowInvoker(request *model.ListWorkflowRequest) *ListWorkflowInvoker {
	requestDef := GenReqDefForListWorkflow()
	return &ListWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkflowExecutions 获取任务执行历史
//
// 该接口可获取执行任务的执行历史。（注：接口目前开放的region为：华北-北京四,华东-上海一,华东-上海二,华南-广州）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) ListWorkflowExecutions(request *model.ListWorkflowExecutionsRequest) (*model.ListWorkflowExecutionsResponse, error) {
	requestDef := GenReqDefForListWorkflowExecutions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkflowExecutionsResponse), nil
	}
}

// ListWorkflowExecutionsInvoker 获取任务执行历史
func (c *AomClient) ListWorkflowExecutionsInvoker(request *model.ListWorkflowExecutionsRequest) *ListWorkflowExecutionsInvoker {
	requestDef := GenReqDefForListWorkflowExecutions()
	return &ListWorkflowExecutionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchTemplateById 获取方案信息
//
// 该接口可根据执行方案id查询执行方案详情。（注：接口目前开放的region为：华北-北京四,华东-上海一,华东-上海二,华南-广州）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) SearchTemplateById(request *model.SearchTemplateByIdRequest) (*model.SearchTemplateByIdResponse, error) {
	requestDef := GenReqDefForSearchTemplateById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchTemplateByIdResponse), nil
	}
}

// SearchTemplateByIdInvoker 获取方案信息
func (c *AomClient) SearchTemplateByIdInvoker(request *model.SearchTemplateByIdRequest) *SearchTemplateByIdInvoker {
	requestDef := GenReqDefForSearchTemplateById()
	return &SearchTemplateByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchWorkflowExecutionDetail 获取工作流执行中的执行详情
//
// 该接口可获取任务的执行详情，可指定工作流ID和执行ID去查询对应的任务，返回任务执行详情。（注：接口目前开放的region为：华北-北京四,华东-上海一,华东-上海二,华南-广州）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) SearchWorkflowExecutionDetail(request *model.SearchWorkflowExecutionDetailRequest) (*model.SearchWorkflowExecutionDetailResponse, error) {
	requestDef := GenReqDefForSearchWorkflowExecutionDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchWorkflowExecutionDetailResponse), nil
	}
}

// SearchWorkflowExecutionDetailInvoker 获取工作流执行中的执行详情
func (c *AomClient) SearchWorkflowExecutionDetailInvoker(request *model.SearchWorkflowExecutionDetailRequest) *SearchWorkflowExecutionDetailInvoker {
	requestDef := GenReqDefForSearchWorkflowExecutionDetail()
	return &SearchWorkflowExecutionDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartPausingWorkflowExecutions 对暂停中的任务进行操作
//
// 该接口可对任务进行失败重试、失败跳过、暂停继续操作，返回操作结果。（注：接口目前开放的region为：华北-北京四,华东-上海一,华东-上海二,华南-广州）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) StartPausingWorkflowExecutions(request *model.StartPausingWorkflowExecutionsRequest) (*model.StartPausingWorkflowExecutionsResponse, error) {
	requestDef := GenReqDefForStartPausingWorkflowExecutions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartPausingWorkflowExecutionsResponse), nil
	}
}

// StartPausingWorkflowExecutionsInvoker 对暂停中的任务进行操作
func (c *AomClient) StartPausingWorkflowExecutionsInvoker(request *model.StartPausingWorkflowExecutionsRequest) *StartPausingWorkflowExecutionsInvoker {
	requestDef := GenReqDefForStartPausingWorkflowExecutions()
	return &StartPausingWorkflowExecutionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopExecution 终止任务执行
//
// 该接口可终止正在执行的任务，指定工作流ID和执行ID去终止对应的任务，返回终止操作状态。（注：接口目前开放的region为：华北-北京四,华东-上海一,华东-上海二,华南-广州）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) StopExecution(request *model.StopExecutionRequest) (*model.StopExecutionResponse, error) {
	requestDef := GenReqDefForStopExecution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopExecutionResponse), nil
	}
}

// StopExecutionInvoker 终止任务执行
func (c *AomClient) StopExecutionInvoker(request *model.StopExecutionRequest) *StopExecutionInvoker {
	requestDef := GenReqDefForStopExecution()
	return &StopExecutionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkflowTriggerStatus 更新任务
//
// 更新定时任务的启停状态，可启动定时任务或停止定时任务，返回操作任务结果。（注：接口目前开放的region为：华北-北京四,华东-上海一,华东-上海二,华南-广州）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) UpdateWorkflowTriggerStatus(request *model.UpdateWorkflowTriggerStatusRequest) (*model.UpdateWorkflowTriggerStatusResponse, error) {
	requestDef := GenReqDefForUpdateWorkflowTriggerStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkflowTriggerStatusResponse), nil
	}
}

// UpdateWorkflowTriggerStatusInvoker 更新任务
func (c *AomClient) UpdateWorkflowTriggerStatusInvoker(request *model.UpdateWorkflowTriggerStatusRequest) *UpdateWorkflowTriggerStatusInvoker {
	requestDef := GenReqDefForUpdateWorkflowTriggerStatus()
	return &UpdateWorkflowTriggerStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
