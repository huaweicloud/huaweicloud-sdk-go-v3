package v3

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dwr/v3/model"
)

type DwrClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDwrClient(hcClient *httpclient.HcHttpClient) *DwrClient {
	return &DwrClient{HcClient: hcClient}
}

func DwrClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AcceptServiceContract 同意服务协议
//
// 本接口用于使用工作流时需要同意服务使用协议。该函数具有幂等性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) AcceptServiceContract(request *model.AcceptServiceContractRequest) (*model.AcceptServiceContractResponse, error) {
	requestDef := GenReqDefForAcceptServiceContract()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AcceptServiceContractResponse), nil
	}
}

// AcceptServiceContractInvoker 同意服务协议
func (c *DwrClient) AcceptServiceContractInvoker(request *model.AcceptServiceContractRequest) *AcceptServiceContractInvoker {
	requestDef := GenReqDefForAcceptServiceContract()
	return &AcceptServiceContractInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AsyncInvokeApiStartWorkflow API异步启动工作流
//
// 本接口用于API方式异步启动已有工作流，产生工作流实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) AsyncInvokeApiStartWorkflow(request *model.AsyncInvokeApiStartWorkflowRequest) (*model.AsyncInvokeApiStartWorkflowResponse, error) {
	requestDef := GenReqDefForAsyncInvokeApiStartWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AsyncInvokeApiStartWorkflowResponse), nil
	}
}

// AsyncInvokeApiStartWorkflowInvoker API异步启动工作流
func (c *DwrClient) AsyncInvokeApiStartWorkflowInvoker(request *model.AsyncInvokeApiStartWorkflowRequest) *AsyncInvokeApiStartWorkflowInvoker {
	requestDef := GenReqDefForAsyncInvokeApiStartWorkflow()
	return &AsyncInvokeApiStartWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckWorkflowAuthentication 查询授权
//
// 本接口用于查询授权，查询由DWR服务自动帮助用户创建工作流运行时需要的函数服务权限，以及函数服务运行时的权限。该函数具有幂等性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) CheckWorkflowAuthentication(request *model.CheckWorkflowAuthenticationRequest) (*model.CheckWorkflowAuthenticationResponse, error) {
	requestDef := GenReqDefForCheckWorkflowAuthentication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckWorkflowAuthenticationResponse), nil
	}
}

// CheckWorkflowAuthenticationInvoker 查询授权
func (c *DwrClient) CheckWorkflowAuthenticationInvoker(request *model.CheckWorkflowAuthenticationRequest) *CheckWorkflowAuthenticationInvoker {
	requestDef := GenReqDefForCheckWorkflowAuthentication()
	return &CheckWorkflowAuthenticationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMyActionTemplate 创建第三方算子模板
//
// 创建第三方算子模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) CreateMyActionTemplate(request *model.CreateMyActionTemplateRequest) (*model.CreateMyActionTemplateResponse, error) {
	requestDef := GenReqDefForCreateMyActionTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMyActionTemplateResponse), nil
	}
}

// CreateMyActionTemplateInvoker 创建第三方算子模板
func (c *DwrClient) CreateMyActionTemplateInvoker(request *model.CreateMyActionTemplateRequest) *CreateMyActionTemplateInvoker {
	requestDef := GenReqDefForCreateMyActionTemplate()
	return &CreateMyActionTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkflowAuthentication 开通授权
//
// 本接口用于开通授权，由DWR服务自动帮助用户创建工作流运行时需要的函数服务权限，以及函数服务运行时的权限。该函数具有幂等性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) CreateWorkflowAuthentication(request *model.CreateWorkflowAuthenticationRequest) (*model.CreateWorkflowAuthenticationResponse, error) {
	requestDef := GenReqDefForCreateWorkflowAuthentication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkflowAuthenticationResponse), nil
	}
}

// CreateWorkflowAuthenticationInvoker 开通授权
func (c *DwrClient) CreateWorkflowAuthenticationInvoker(request *model.CreateWorkflowAuthenticationRequest) *CreateWorkflowAuthenticationInvoker {
	requestDef := GenReqDefForCreateWorkflowAuthentication()
	return &CreateWorkflowAuthenticationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMyActionTemplate 删除第三方算子模板
//
// 本接口用于标记删除提交的第三方算子模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) DeleteMyActionTemplate(request *model.DeleteMyActionTemplateRequest) (*model.DeleteMyActionTemplateResponse, error) {
	requestDef := GenReqDefForDeleteMyActionTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMyActionTemplateResponse), nil
	}
}

// DeleteMyActionTemplateInvoker 删除第三方算子模板
func (c *DwrClient) DeleteMyActionTemplateInvoker(request *model.DeleteMyActionTemplateRequest) *DeleteMyActionTemplateInvoker {
	requestDef := GenReqDefForDeleteMyActionTemplate()
	return &DeleteMyActionTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMyActionTemplate 查询第三方算子列表
//
// 本接口用于查询提交注册过的三方算子列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) ListMyActionTemplate(request *model.ListMyActionTemplateRequest) (*model.ListMyActionTemplateResponse, error) {
	requestDef := GenReqDefForListMyActionTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMyActionTemplateResponse), nil
	}
}

// ListMyActionTemplateInvoker 查询第三方算子列表
func (c *DwrClient) ListMyActionTemplateInvoker(request *model.ListMyActionTemplateRequest) *ListMyActionTemplateInvoker {
	requestDef := GenReqDefForListMyActionTemplate()
	return &ListMyActionTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSystemTemplates 查询华为云内置算子列表
//
// 本接口用于按名称查询系统内置算子列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) ListSystemTemplates(request *model.ListSystemTemplatesRequest) (*model.ListSystemTemplatesResponse, error) {
	requestDef := GenReqDefForListSystemTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSystemTemplatesResponse), nil
	}
}

// ListSystemTemplatesInvoker 查询华为云内置算子列表
func (c *DwrClient) ListSystemTemplatesInvoker(request *model.ListSystemTemplatesRequest) *ListSystemTemplatesInvoker {
	requestDef := GenReqDefForListSystemTemplates()
	return &ListSystemTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkflowInstance 本接口用于查询用户工作流的实例列表
//
// 本接口用于查询用户工作流的实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) ListWorkflowInstance(request *model.ListWorkflowInstanceRequest) (*model.ListWorkflowInstanceResponse, error) {
	requestDef := GenReqDefForListWorkflowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkflowInstanceResponse), nil
	}
}

// ListWorkflowInstanceInvoker 本接口用于查询用户工作流的实例列表
func (c *DwrClient) ListWorkflowInstanceInvoker(request *model.ListWorkflowInstanceRequest) *ListWorkflowInstanceInvoker {
	requestDef := GenReqDefForListWorkflowInstance()
	return &ListWorkflowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreWorkflowExecution 恢复一个执行失败状态的工作流实例
//
// 本接口用于恢复一个执行失败状态的工作流实例。恢复后，工作流实例将从上次失败的状态处继续执行，而工作流步骤中已经执行成功的状态不会再执行。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) RestoreWorkflowExecution(request *model.RestoreWorkflowExecutionRequest) (*model.RestoreWorkflowExecutionResponse, error) {
	requestDef := GenReqDefForRestoreWorkflowExecution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreWorkflowExecutionResponse), nil
	}
}

// RestoreWorkflowExecutionInvoker 恢复一个执行失败状态的工作流实例
func (c *DwrClient) RestoreWorkflowExecutionInvoker(request *model.RestoreWorkflowExecutionRequest) *RestoreWorkflowExecutionInvoker {
	requestDef := GenReqDefForRestoreWorkflowExecution()
	return &RestoreWorkflowExecutionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPublicActionList 查询已发布算子列表
//
// 本接口用于查询开放的算子列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) ShowPublicActionList(request *model.ShowPublicActionListRequest) (*model.ShowPublicActionListResponse, error) {
	requestDef := GenReqDefForShowPublicActionList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicActionListResponse), nil
	}
}

// ShowPublicActionListInvoker 查询已发布算子列表
func (c *DwrClient) ShowPublicActionListInvoker(request *model.ShowPublicActionListRequest) *ShowPublicActionListInvoker {
	requestDef := GenReqDefForShowPublicActionList()
	return &ShowPublicActionListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPublicTemplateInfo 查询已发布算子模板详情
//
// 本接口用于按名称查询开放的算子详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) ShowPublicTemplateInfo(request *model.ShowPublicTemplateInfoRequest) (*model.ShowPublicTemplateInfoResponse, error) {
	requestDef := GenReqDefForShowPublicTemplateInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicTemplateInfoResponse), nil
	}
}

// ShowPublicTemplateInfoInvoker 查询已发布算子模板详情
func (c *DwrClient) ShowPublicTemplateInfoInvoker(request *model.ShowPublicTemplateInfoRequest) *ShowPublicTemplateInfoInvoker {
	requestDef := GenReqDefForShowPublicTemplateInfo()
	return &ShowPublicTemplateInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServiceContract 查询服务协议
//
// 本接口用于查询使用工作流时同意的服务协议。该函数具有幂等性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) ShowServiceContract(request *model.ShowServiceContractRequest) (*model.ShowServiceContractResponse, error) {
	requestDef := GenReqDefForShowServiceContract()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServiceContractResponse), nil
	}
}

// ShowServiceContractInvoker 查询服务协议
func (c *DwrClient) ShowServiceContractInvoker(request *model.ShowServiceContractRequest) *ShowServiceContractInvoker {
	requestDef := GenReqDefForShowServiceContract()
	return &ShowServiceContractInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSystemTemplateDetail 查询华为云内置算子模板信息
//
// 本接口用于按名称查询系统内置算子详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) ShowSystemTemplateDetail(request *model.ShowSystemTemplateDetailRequest) (*model.ShowSystemTemplateDetailResponse, error) {
	requestDef := GenReqDefForShowSystemTemplateDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSystemTemplateDetailResponse), nil
	}
}

// ShowSystemTemplateDetailInvoker 查询华为云内置算子模板信息
func (c *DwrClient) ShowSystemTemplateDetailInvoker(request *model.ShowSystemTemplateDetailRequest) *ShowSystemTemplateDetailInvoker {
	requestDef := GenReqDefForShowSystemTemplateDetail()
	return &ShowSystemTemplateDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowThirdTemplateInfo 查询公共Action模板详情
//
// 本接口用于按名称查询第三方模板详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) ShowThirdTemplateInfo(request *model.ShowThirdTemplateInfoRequest) (*model.ShowThirdTemplateInfoResponse, error) {
	requestDef := GenReqDefForShowThirdTemplateInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowThirdTemplateInfoResponse), nil
	}
}

// ShowThirdTemplateInfoInvoker 查询公共Action模板详情
func (c *DwrClient) ShowThirdTemplateInfoInvoker(request *model.ShowThirdTemplateInfoRequest) *ShowThirdTemplateInfoInvoker {
	requestDef := GenReqDefForShowThirdTemplateInfo()
	return &ShowThirdTemplateInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkflowInstance 本接口用于查询指定工作流实例详细
//
// 本接口用于查询指定工作流实例详细。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) ShowWorkflowInstance(request *model.ShowWorkflowInstanceRequest) (*model.ShowWorkflowInstanceResponse, error) {
	requestDef := GenReqDefForShowWorkflowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkflowInstanceResponse), nil
	}
}

// ShowWorkflowInstanceInvoker 本接口用于查询指定工作流实例详细
func (c *DwrClient) ShowWorkflowInstanceInvoker(request *model.ShowWorkflowInstanceRequest) *ShowWorkflowInstanceInvoker {
	requestDef := GenReqDefForShowWorkflowInstance()
	return &ShowWorkflowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMyActionTemplate 更新第三方算子模板
//
// 本接口用于修改第三方算子和将三方算子提交审核
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) UpdateMyActionTemplate(request *model.UpdateMyActionTemplateRequest) (*model.UpdateMyActionTemplateResponse, error) {
	requestDef := GenReqDefForUpdateMyActionTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMyActionTemplateResponse), nil
	}
}

// UpdateMyActionTemplateInvoker 更新第三方算子模板
func (c *DwrClient) UpdateMyActionTemplateInvoker(request *model.UpdateMyActionTemplateRequest) *UpdateMyActionTemplateInvoker {
	requestDef := GenReqDefForUpdateMyActionTemplate()
	return &UpdateMyActionTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMyActionTemplateToDeprecated 禁用第三方算子模板
//
// 本接口用于申请禁用第三方算子。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) UpdateMyActionTemplateToDeprecated(request *model.UpdateMyActionTemplateToDeprecatedRequest) (*model.UpdateMyActionTemplateToDeprecatedResponse, error) {
	requestDef := GenReqDefForUpdateMyActionTemplateToDeprecated()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMyActionTemplateToDeprecatedResponse), nil
	}
}

// UpdateMyActionTemplateToDeprecatedInvoker 禁用第三方算子模板
func (c *DwrClient) UpdateMyActionTemplateToDeprecatedInvoker(request *model.UpdateMyActionTemplateToDeprecatedRequest) *UpdateMyActionTemplateToDeprecatedInvoker {
	requestDef := GenReqDefForUpdateMyActionTemplateToDeprecated()
	return &UpdateMyActionTemplateToDeprecatedInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkflow 创建工作流
//
// 本接口用于通过Body体直接创建工作流
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) CreateWorkflow(request *model.CreateWorkflowRequest) (*model.CreateWorkflowResponse, error) {
	requestDef := GenReqDefForCreateWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkflowResponse), nil
	}
}

// CreateWorkflowInvoker 创建工作流
func (c *DwrClient) CreateWorkflowInvoker(request *model.CreateWorkflowRequest) *CreateWorkflowInvoker {
	requestDef := GenReqDefForCreateWorkflow()
	return &CreateWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWorkflow 删除工作流
//
// 本接口用于删除工作流。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) DeleteWorkflow(request *model.DeleteWorkflowRequest) (*model.DeleteWorkflowResponse, error) {
	requestDef := GenReqDefForDeleteWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWorkflowResponse), nil
	}
}

// DeleteWorkflowInvoker 删除工作流
func (c *DwrClient) DeleteWorkflowInvoker(request *model.DeleteWorkflowRequest) *DeleteWorkflowInvoker {
	requestDef := GenReqDefForDeleteWorkflow()
	return &DeleteWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkflows 查询工作流列表
//
// 本接口用于查询工作流列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) ListWorkflows(request *model.ListWorkflowsRequest) (*model.ListWorkflowsResponse, error) {
	requestDef := GenReqDefForListWorkflows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkflowsResponse), nil
	}
}

// ListWorkflowsInvoker 查询工作流列表
func (c *DwrClient) ListWorkflowsInvoker(request *model.ListWorkflowsRequest) *ListWorkflowsInvoker {
	requestDef := GenReqDefForListWorkflows()
	return &ListWorkflowsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkflowInfo 查询工作流信息
//
// 本接口用于根据工作流名称查询工作流详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) ShowWorkflowInfo(request *model.ShowWorkflowInfoRequest) (*model.ShowWorkflowInfoResponse, error) {
	requestDef := GenReqDefForShowWorkflowInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkflowInfoResponse), nil
	}
}

// ShowWorkflowInfoInvoker 查询工作流信息
func (c *DwrClient) ShowWorkflowInfoInvoker(request *model.ShowWorkflowInfoRequest) *ShowWorkflowInfoInvoker {
	requestDef := GenReqDefForShowWorkflowInfo()
	return &ShowWorkflowInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkflow 更新工作流
//
// # Update Workflow
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DwrClient) UpdateWorkflow(request *model.UpdateWorkflowRequest) (*model.UpdateWorkflowResponse, error) {
	requestDef := GenReqDefForUpdateWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkflowResponse), nil
	}
}

// UpdateWorkflowInvoker 更新工作流
func (c *DwrClient) UpdateWorkflowInvoker(request *model.UpdateWorkflowRequest) *UpdateWorkflowInvoker {
	requestDef := GenReqDefForUpdateWorkflow()
	return &UpdateWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
