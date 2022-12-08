package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/aos/v1/model"
)

type AosClient struct {
	HcClient *http_client.HcHttpClient
}

func NewAosClient(hcClient *http_client.HcHttpClient) *AosClient {
	return &AosClient{HcClient: hcClient}
}

func AosClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// ApplyExecutionPlan 此命令用于执行已有的执行计划(execution plan)
//
// 此命令用于执行已有的执行计划(execution plan)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AosClient) ApplyExecutionPlan(request *model.ApplyExecutionPlanRequest) (*model.ApplyExecutionPlanResponse, error) {
	requestDef := GenReqDefForApplyExecutionPlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApplyExecutionPlanResponse), nil
	}
}

// ApplyExecutionPlanInvoker 此命令用于执行已有的执行计划(execution plan)
func (c *AosClient) ApplyExecutionPlanInvoker(request *model.ApplyExecutionPlanRequest) *ApplyExecutionPlanInvoker {
	requestDef := GenReqDefForApplyExecutionPlan()
	return &ApplyExecutionPlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateExecutionPlan 此命令用于生成一个执行计划(execution plan)
//
// 此命令用于生成一个执行计划(execution plan)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AosClient) CreateExecutionPlan(request *model.CreateExecutionPlanRequest) (*model.CreateExecutionPlanResponse, error) {
	requestDef := GenReqDefForCreateExecutionPlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateExecutionPlanResponse), nil
	}
}

// CreateExecutionPlanInvoker 此命令用于生成一个执行计划(execution plan)
func (c *AosClient) CreateExecutionPlanInvoker(request *model.CreateExecutionPlanRequest) *CreateExecutionPlanInvoker {
	requestDef := GenReqDefForCreateExecutionPlan()
	return &CreateExecutionPlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteExecutionPlan 此命令用于删除已有的执行计划(execution plan)
//
// 此命令用于删除已有的执行计划(execution plan)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AosClient) DeleteExecutionPlan(request *model.DeleteExecutionPlanRequest) (*model.DeleteExecutionPlanResponse, error) {
	requestDef := GenReqDefForDeleteExecutionPlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteExecutionPlanResponse), nil
	}
}

// DeleteExecutionPlanInvoker 此命令用于删除已有的执行计划(execution plan)
func (c *AosClient) DeleteExecutionPlanInvoker(request *model.DeleteExecutionPlanRequest) *DeleteExecutionPlanInvoker {
	requestDef := GenReqDefForDeleteExecutionPlan()
	return &DeleteExecutionPlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteStack 删除资源栈
//
// 删除资源栈
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AosClient) DeleteStack(request *model.DeleteStackRequest) (*model.DeleteStackResponse, error) {
	requestDef := GenReqDefForDeleteStack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStackResponse), nil
	}
}

// DeleteStackInvoker 删除资源栈
func (c *AosClient) DeleteStackInvoker(request *model.DeleteStackRequest) *DeleteStackInvoker {
	requestDef := GenReqDefForDeleteStack()
	return &DeleteStackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribeExecutionPlan 描述执行计划当前的状态，返回执行计划的元数据
//
// 描述执行计划当前的状态，返回执行计划的元数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AosClient) DescribeExecutionPlan(request *model.DescribeExecutionPlanRequest) (*model.DescribeExecutionPlanResponse, error) {
	requestDef := GenReqDefForDescribeExecutionPlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DescribeExecutionPlanResponse), nil
	}
}

// DescribeExecutionPlanInvoker 描述执行计划当前的状态，返回执行计划的元数据
func (c *AosClient) DescribeExecutionPlanInvoker(request *model.DescribeExecutionPlanRequest) *DescribeExecutionPlanInvoker {
	requestDef := GenReqDefForDescribeExecutionPlan()
	return &DescribeExecutionPlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EstimateExecutionPlanPrice 预估执行计划的价格
//
// 预估执行计划的价格
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AosClient) EstimateExecutionPlanPrice(request *model.EstimateExecutionPlanPriceRequest) (*model.EstimateExecutionPlanPriceResponse, error) {
	requestDef := GenReqDefForEstimateExecutionPlanPrice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EstimateExecutionPlanPriceResponse), nil
	}
}

// EstimateExecutionPlanPriceInvoker 预估执行计划的价格
func (c *AosClient) EstimateExecutionPlanPriceInvoker(request *model.EstimateExecutionPlanPriceRequest) *EstimateExecutionPlanPriceInvoker {
	requestDef := GenReqDefForEstimateExecutionPlanPrice()
	return &EstimateExecutionPlanPriceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetExecutionPlan 此接口用于获取执行计划的详细内容
//
// 此接口用于获取执行计划的详细内容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AosClient) GetExecutionPlan(request *model.GetExecutionPlanRequest) (*model.GetExecutionPlanResponse, error) {
	requestDef := GenReqDefForGetExecutionPlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetExecutionPlanResponse), nil
	}
}

// GetExecutionPlanInvoker 此接口用于获取执行计划的详细内容
func (c *AosClient) GetExecutionPlanInvoker(request *model.GetExecutionPlanRequest) *GetExecutionPlanInvoker {
	requestDef := GenReqDefForGetExecutionPlan()
	return &GetExecutionPlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetStackMetadata 描述栈的状态，返回栈的元数据
//
// 描述栈的状态，返回栈的元数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AosClient) GetStackMetadata(request *model.GetStackMetadataRequest) (*model.GetStackMetadataResponse, error) {
	requestDef := GenReqDefForGetStackMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetStackMetadataResponse), nil
	}
}

// GetStackMetadataInvoker 描述栈的状态，返回栈的元数据
func (c *AosClient) GetStackMetadataInvoker(request *model.GetStackMetadataRequest) *GetStackMetadataInvoker {
	requestDef := GenReqDefForGetStackMetadata()
	return &GetStackMetadataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetStackTemplate 获取资源栈模板
//
// 获取资源栈当前使用的模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AosClient) GetStackTemplate(request *model.GetStackTemplateRequest) (*model.GetStackTemplateResponse, error) {
	requestDef := GenReqDefForGetStackTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetStackTemplateResponse), nil
	}
}

// GetStackTemplateInvoker 获取资源栈模板
func (c *AosClient) GetStackTemplateInvoker(request *model.GetStackTemplateRequest) *GetStackTemplateInvoker {
	requestDef := GenReqDefForGetStackTemplate()
	return &GetStackTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListExecutionPlans 列举执行计划
//
// 列举执行计划
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AosClient) ListExecutionPlans(request *model.ListExecutionPlansRequest) (*model.ListExecutionPlansResponse, error) {
	requestDef := GenReqDefForListExecutionPlans()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListExecutionPlansResponse), nil
	}
}

// ListExecutionPlansInvoker 列举执行计划
func (c *AosClient) ListExecutionPlansInvoker(request *model.ListExecutionPlansRequest) *ListExecutionPlansInvoker {
	requestDef := GenReqDefForListExecutionPlans()
	return &ListExecutionPlansInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStackEvents 获取栈的细节更新状态
//
// 获取栈的细节更新状态，可以获取整个栈从生成到当前时间点的所有状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AosClient) ListStackEvents(request *model.ListStackEventsRequest) (*model.ListStackEventsResponse, error) {
	requestDef := GenReqDefForListStackEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStackEventsResponse), nil
	}
}

// ListStackEventsInvoker 获取栈的细节更新状态
func (c *AosClient) ListStackEventsInvoker(request *model.ListStackEventsRequest) *ListStackEventsInvoker {
	requestDef := GenReqDefForListStackEvents()
	return &ListStackEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStackOutputs 列举资源栈的输出
//
// 列举资源栈的输出
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AosClient) ListStackOutputs(request *model.ListStackOutputsRequest) (*model.ListStackOutputsResponse, error) {
	requestDef := GenReqDefForListStackOutputs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStackOutputsResponse), nil
	}
}

// ListStackOutputsInvoker 列举资源栈的输出
func (c *AosClient) ListStackOutputsInvoker(request *model.ListStackOutputsRequest) *ListStackOutputsInvoker {
	requestDef := GenReqDefForListStackOutputs()
	return &ListStackOutputsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStackResources 获取资源栈的资源列表
//
// 获取资源栈的资源列表，可以获取整个栈从生成到当前时间点的所有状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AosClient) ListStackResources(request *model.ListStackResourcesRequest) (*model.ListStackResourcesResponse, error) {
	requestDef := GenReqDefForListStackResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStackResourcesResponse), nil
	}
}

// ListStackResourcesInvoker 获取资源栈的资源列表
func (c *AosClient) ListStackResourcesInvoker(request *model.ListStackResourcesRequest) *ListStackResourcesInvoker {
	requestDef := GenReqDefForListStackResources()
	return &ListStackResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ParseTemplateVariables 此命令用于解析模板参数
//
// 此命令用于解析模板参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AosClient) ParseTemplateVariables(request *model.ParseTemplateVariablesRequest) (*model.ParseTemplateVariablesResponse, error) {
	requestDef := GenReqDefForParseTemplateVariables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ParseTemplateVariablesResponse), nil
	}
}

// ParseTemplateVariablesInvoker 此命令用于解析模板参数
func (c *AosClient) ParseTemplateVariablesInvoker(request *model.ParseTemplateVariablesRequest) *ParseTemplateVariablesInvoker {
	requestDef := GenReqDefForParseTemplateVariables()
	return &ParseTemplateVariablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ContinueRollbackStack 继续回滚资源栈
//
// 如果资源栈开启了自动回滚，在部署失败的时候则会自动回滚。但是自动回滚依然有可能失败，用户可以根据错误信息修复后，调用ContinueRollbackStack触发继续回滚，即重试回滚
//
// * 如果资源栈当前可以回滚，即处于&#x60;ROLLBACK_FAILED&#x60;，则返回202与对应生成的deploymentId，否则将不允许回滚并返回响应的错误码
// * 继续回滚也有可能会回滚失败。如果失败，用户可以从ListStackEvents获取对应的log，解决后可再次调用ContinueRollbackStack去继续触发回滚
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AosClient) ContinueRollbackStack(request *model.ContinueRollbackStackRequest) (*model.ContinueRollbackStackResponse, error) {
	requestDef := GenReqDefForContinueRollbackStack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ContinueRollbackStackResponse), nil
	}
}

// ContinueRollbackStackInvoker 继续回滚资源栈
func (c *AosClient) ContinueRollbackStackInvoker(request *model.ContinueRollbackStackRequest) *ContinueRollbackStackInvoker {
	requestDef := GenReqDefForContinueRollbackStack()
	return &ContinueRollbackStackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStack 创建资源栈
//
// CreateStack用于生成一个资源栈
//
// * 当请求中不含有模板（template）、参数（vars）等信息，将生成一个无任何资源的空资源栈，返回资源栈ID（stack_id）
// * 当请求中携带了模板（template）、参数（vars）等信息，则会同时创建并部署资源栈，返回资源栈ID（stack_id）和部署ID（deployment_id）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AosClient) CreateStack(request *model.CreateStackRequest) (*model.CreateStackResponse, error) {
	requestDef := GenReqDefForCreateStack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStackResponse), nil
	}
}

// CreateStackInvoker 创建资源栈
func (c *AosClient) CreateStackInvoker(request *model.CreateStackRequest) *CreateStackInvoker {
	requestDef := GenReqDefForCreateStack()
	return &CreateStackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeployStack 部署一个已有的资源栈
//
// 部署一个已有的资源栈
//
// * 用户可以使用此API更新模板、参数等并触发一个新的部署
//
// * 此API会直接触发部署，如果用户希望先确认部署会发生的时间，请使用执行计划，即使用CreateExecutionPlan以创建执行计划、使用GetExecutionPlan以获取执行计划
//
// * 此API为全量API，即用户每次部署都需要给予所想要使用的template、vars的全量
//
// * 当触发的部署失败时，如果资源栈开启了自动回滚，会触发自动回滚的流程，否则就会停留在部署失败时的状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AosClient) DeployStack(request *model.DeployStackRequest) (*model.DeployStackResponse, error) {
	requestDef := GenReqDefForDeployStack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeployStackResponse), nil
	}
}

// DeployStackInvoker 部署一个已有的资源栈
func (c *AosClient) DeployStackInvoker(request *model.DeployStackRequest) *DeployStackInvoker {
	requestDef := GenReqDefForDeployStack()
	return &DeployStackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStacks 列举资源栈
//
// ListStacks 列举当前局点下用户所有的资源栈
//
//   * 默认按照生成时间排序，最早生成的在最前
//   * 注意：目前暂时返回全量资源栈信息，即不支持分页
//   * 如果没有任何资源栈，则返回空list
//
// ListStacks返回的只有摘要信息（具体摘要信息见ListStacksResponseBody），如果用户需要详细的资源栈元数据请调用GetStackMetadata
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AosClient) ListStacks(request *model.ListStacksRequest) (*model.ListStacksResponse, error) {
	requestDef := GenReqDefForListStacks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStacksResponse), nil
	}
}

// ListStacksInvoker 列举资源栈
func (c *AosClient) ListStacksInvoker(request *model.ListStacksRequest) *ListStacksInvoker {
	requestDef := GenReqDefForListStacks()
	return &ListStacksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
