package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rms/v1/model"
)

type RmsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewRmsClient(hcClient *http_client.HcHttpClient) *RmsClient {
	return &RmsClient{HcClient: hcClient}
}

func RmsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// ShowResourceHistory 查询资源历史
//
// 查询资源与资源关系的变更历史
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ShowResourceHistory(request *model.ShowResourceHistoryRequest) (*model.ShowResourceHistoryResponse, error) {
	requestDef := GenReqDefForShowResourceHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceHistoryResponse), nil
	}
}

// ShowResourceHistoryInvoker 查询资源历史
func (c *RmsClient) ShowResourceHistoryInvoker(request *model.ShowResourceHistoryRequest) *ShowResourceHistoryInvoker {
	requestDef := GenReqDefForShowResourceHistory()
	return &ShowResourceHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePolicyAssignments 创建合规规则
//
// 创建新的合规规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) CreatePolicyAssignments(request *model.CreatePolicyAssignmentsRequest) (*model.CreatePolicyAssignmentsResponse, error) {
	requestDef := GenReqDefForCreatePolicyAssignments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePolicyAssignmentsResponse), nil
	}
}

// CreatePolicyAssignmentsInvoker 创建合规规则
func (c *RmsClient) CreatePolicyAssignmentsInvoker(request *model.CreatePolicyAssignmentsRequest) *CreatePolicyAssignmentsInvoker {
	requestDef := GenReqDefForCreatePolicyAssignments()
	return &CreatePolicyAssignmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePolicyAssignment 删除合规规则
//
// 根据规则ID删除此规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) DeletePolicyAssignment(request *model.DeletePolicyAssignmentRequest) (*model.DeletePolicyAssignmentResponse, error) {
	requestDef := GenReqDefForDeletePolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePolicyAssignmentResponse), nil
	}
}

// DeletePolicyAssignmentInvoker 删除合规规则
func (c *RmsClient) DeletePolicyAssignmentInvoker(request *model.DeletePolicyAssignmentRequest) *DeletePolicyAssignmentInvoker {
	requestDef := GenReqDefForDeletePolicyAssignment()
	return &DeletePolicyAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisablePolicyAssignment 停用合规规则
//
// 根据规则ID停用此规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) DisablePolicyAssignment(request *model.DisablePolicyAssignmentRequest) (*model.DisablePolicyAssignmentResponse, error) {
	requestDef := GenReqDefForDisablePolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisablePolicyAssignmentResponse), nil
	}
}

// DisablePolicyAssignmentInvoker 停用合规规则
func (c *RmsClient) DisablePolicyAssignmentInvoker(request *model.DisablePolicyAssignmentRequest) *DisablePolicyAssignmentInvoker {
	requestDef := GenReqDefForDisablePolicyAssignment()
	return &DisablePolicyAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnablePolicyAssignment 启用合规规则
//
// 根据规则ID启用此规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) EnablePolicyAssignment(request *model.EnablePolicyAssignmentRequest) (*model.EnablePolicyAssignmentResponse, error) {
	requestDef := GenReqDefForEnablePolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnablePolicyAssignmentResponse), nil
	}
}

// EnablePolicyAssignmentInvoker 启用合规规则
func (c *RmsClient) EnablePolicyAssignmentInvoker(request *model.EnablePolicyAssignmentRequest) *EnablePolicyAssignmentInvoker {
	requestDef := GenReqDefForEnablePolicyAssignment()
	return &EnablePolicyAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBuiltInPolicyDefinitions 列出内置策略
//
// 列出用户的内置策略
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListBuiltInPolicyDefinitions(request *model.ListBuiltInPolicyDefinitionsRequest) (*model.ListBuiltInPolicyDefinitionsResponse, error) {
	requestDef := GenReqDefForListBuiltInPolicyDefinitions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBuiltInPolicyDefinitionsResponse), nil
	}
}

// ListBuiltInPolicyDefinitionsInvoker 列出内置策略
func (c *RmsClient) ListBuiltInPolicyDefinitionsInvoker(request *model.ListBuiltInPolicyDefinitionsRequest) *ListBuiltInPolicyDefinitionsInvoker {
	requestDef := GenReqDefForListBuiltInPolicyDefinitions()
	return &ListBuiltInPolicyDefinitionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyAssignments 列出合规规则
//
// 列出用户的合规规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListPolicyAssignments(request *model.ListPolicyAssignmentsRequest) (*model.ListPolicyAssignmentsResponse, error) {
	requestDef := GenReqDefForListPolicyAssignments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyAssignmentsResponse), nil
	}
}

// ListPolicyAssignmentsInvoker 列出合规规则
func (c *RmsClient) ListPolicyAssignmentsInvoker(request *model.ListPolicyAssignmentsRequest) *ListPolicyAssignmentsInvoker {
	requestDef := GenReqDefForListPolicyAssignments()
	return &ListPolicyAssignmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyStatesByAssignmentId 获取规则的合规结果
//
// 根据规则ID查询所有的合规结果
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListPolicyStatesByAssignmentId(request *model.ListPolicyStatesByAssignmentIdRequest) (*model.ListPolicyStatesByAssignmentIdResponse, error) {
	requestDef := GenReqDefForListPolicyStatesByAssignmentId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyStatesByAssignmentIdResponse), nil
	}
}

// ListPolicyStatesByAssignmentIdInvoker 获取规则的合规结果
func (c *RmsClient) ListPolicyStatesByAssignmentIdInvoker(request *model.ListPolicyStatesByAssignmentIdRequest) *ListPolicyStatesByAssignmentIdInvoker {
	requestDef := GenReqDefForListPolicyStatesByAssignmentId()
	return &ListPolicyStatesByAssignmentIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyStatesByDomainId 获取用户的合规结果
//
// 查询用户所有的合规结果
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListPolicyStatesByDomainId(request *model.ListPolicyStatesByDomainIdRequest) (*model.ListPolicyStatesByDomainIdResponse, error) {
	requestDef := GenReqDefForListPolicyStatesByDomainId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyStatesByDomainIdResponse), nil
	}
}

// ListPolicyStatesByDomainIdInvoker 获取用户的合规结果
func (c *RmsClient) ListPolicyStatesByDomainIdInvoker(request *model.ListPolicyStatesByDomainIdRequest) *ListPolicyStatesByDomainIdInvoker {
	requestDef := GenReqDefForListPolicyStatesByDomainId()
	return &ListPolicyStatesByDomainIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyStatesByResourceId 获取资源的合规结果
//
// 根据资源ID查询所有合规结果
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListPolicyStatesByResourceId(request *model.ListPolicyStatesByResourceIdRequest) (*model.ListPolicyStatesByResourceIdResponse, error) {
	requestDef := GenReqDefForListPolicyStatesByResourceId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyStatesByResourceIdResponse), nil
	}
}

// ListPolicyStatesByResourceIdInvoker 获取资源的合规结果
func (c *RmsClient) ListPolicyStatesByResourceIdInvoker(request *model.ListPolicyStatesByResourceIdRequest) *ListPolicyStatesByResourceIdInvoker {
	requestDef := GenReqDefForListPolicyStatesByResourceId()
	return &ListPolicyStatesByResourceIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunEvaluationByPolicyAssignmentId 运行合规评估
//
// 根据规则ID评估此规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) RunEvaluationByPolicyAssignmentId(request *model.RunEvaluationByPolicyAssignmentIdRequest) (*model.RunEvaluationByPolicyAssignmentIdResponse, error) {
	requestDef := GenReqDefForRunEvaluationByPolicyAssignmentId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunEvaluationByPolicyAssignmentIdResponse), nil
	}
}

// RunEvaluationByPolicyAssignmentIdInvoker 运行合规评估
func (c *RmsClient) RunEvaluationByPolicyAssignmentIdInvoker(request *model.RunEvaluationByPolicyAssignmentIdRequest) *RunEvaluationByPolicyAssignmentIdInvoker {
	requestDef := GenReqDefForRunEvaluationByPolicyAssignmentId()
	return &RunEvaluationByPolicyAssignmentIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBuiltInPolicyDefinition 查询单个内置策略
//
// 根据策略ID查询单个内置策略
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ShowBuiltInPolicyDefinition(request *model.ShowBuiltInPolicyDefinitionRequest) (*model.ShowBuiltInPolicyDefinitionResponse, error) {
	requestDef := GenReqDefForShowBuiltInPolicyDefinition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBuiltInPolicyDefinitionResponse), nil
	}
}

// ShowBuiltInPolicyDefinitionInvoker 查询单个内置策略
func (c *RmsClient) ShowBuiltInPolicyDefinitionInvoker(request *model.ShowBuiltInPolicyDefinitionRequest) *ShowBuiltInPolicyDefinitionInvoker {
	requestDef := GenReqDefForShowBuiltInPolicyDefinition()
	return &ShowBuiltInPolicyDefinitionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEvaluationStateByAssignmentId 获取规则的评估状态
//
// 根据规则ID查询此规则的评估状态
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ShowEvaluationStateByAssignmentId(request *model.ShowEvaluationStateByAssignmentIdRequest) (*model.ShowEvaluationStateByAssignmentIdResponse, error) {
	requestDef := GenReqDefForShowEvaluationStateByAssignmentId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEvaluationStateByAssignmentIdResponse), nil
	}
}

// ShowEvaluationStateByAssignmentIdInvoker 获取规则的评估状态
func (c *RmsClient) ShowEvaluationStateByAssignmentIdInvoker(request *model.ShowEvaluationStateByAssignmentIdRequest) *ShowEvaluationStateByAssignmentIdInvoker {
	requestDef := GenReqDefForShowEvaluationStateByAssignmentId()
	return &ShowEvaluationStateByAssignmentIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPolicyAssignment 获取单个合规规则
//
// 根据规则ID获取单个规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ShowPolicyAssignment(request *model.ShowPolicyAssignmentRequest) (*model.ShowPolicyAssignmentResponse, error) {
	requestDef := GenReqDefForShowPolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPolicyAssignmentResponse), nil
	}
}

// ShowPolicyAssignmentInvoker 获取单个合规规则
func (c *RmsClient) ShowPolicyAssignmentInvoker(request *model.ShowPolicyAssignmentRequest) *ShowPolicyAssignmentInvoker {
	requestDef := GenReqDefForShowPolicyAssignment()
	return &ShowPolicyAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePolicyAssignment 更新合规规则
//
// 更新用户的合规规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) UpdatePolicyAssignment(request *model.UpdatePolicyAssignmentRequest) (*model.UpdatePolicyAssignmentResponse, error) {
	requestDef := GenReqDefForUpdatePolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePolicyAssignmentResponse), nil
	}
}

// UpdatePolicyAssignmentInvoker 更新合规规则
func (c *RmsClient) UpdatePolicyAssignmentInvoker(request *model.UpdatePolicyAssignmentRequest) *UpdatePolicyAssignmentInvoker {
	requestDef := GenReqDefForUpdatePolicyAssignment()
	return &UpdatePolicyAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStoredQuery 创建高级查询
//
// 创建新的高级查询
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) CreateStoredQuery(request *model.CreateStoredQueryRequest) (*model.CreateStoredQueryResponse, error) {
	requestDef := GenReqDefForCreateStoredQuery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStoredQueryResponse), nil
	}
}

// CreateStoredQueryInvoker 创建高级查询
func (c *RmsClient) CreateStoredQueryInvoker(request *model.CreateStoredQueryRequest) *CreateStoredQueryInvoker {
	requestDef := GenReqDefForCreateStoredQuery()
	return &CreateStoredQueryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteStoredQuery 删除高级查询
//
// 删除单个高级查询
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) DeleteStoredQuery(request *model.DeleteStoredQueryRequest) (*model.DeleteStoredQueryResponse, error) {
	requestDef := GenReqDefForDeleteStoredQuery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStoredQueryResponse), nil
	}
}

// DeleteStoredQueryInvoker 删除高级查询
func (c *RmsClient) DeleteStoredQueryInvoker(request *model.DeleteStoredQueryRequest) *DeleteStoredQueryInvoker {
	requestDef := GenReqDefForDeleteStoredQuery()
	return &DeleteStoredQueryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSchemas 列举高级查询Schema
//
// List Schemas
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListSchemas(request *model.ListSchemasRequest) (*model.ListSchemasResponse, error) {
	requestDef := GenReqDefForListSchemas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSchemasResponse), nil
	}
}

// ListSchemasInvoker 列举高级查询Schema
func (c *RmsClient) ListSchemasInvoker(request *model.ListSchemasRequest) *ListSchemasInvoker {
	requestDef := GenReqDefForListSchemas()
	return &ListSchemasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStoredQueries 列出高级查询
//
// 列举所有高级查询
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListStoredQueries(request *model.ListStoredQueriesRequest) (*model.ListStoredQueriesResponse, error) {
	requestDef := GenReqDefForListStoredQueries()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStoredQueriesResponse), nil
	}
}

// ListStoredQueriesInvoker 列出高级查询
func (c *RmsClient) ListStoredQueriesInvoker(request *model.ListStoredQueriesRequest) *ListStoredQueriesInvoker {
	requestDef := GenReqDefForListStoredQueries()
	return &ListStoredQueriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunQuery 运行高级查询
//
// 执行高级查询
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) RunQuery(request *model.RunQueryRequest) (*model.RunQueryResponse, error) {
	requestDef := GenReqDefForRunQuery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunQueryResponse), nil
	}
}

// RunQueryInvoker 运行高级查询
func (c *RmsClient) RunQueryInvoker(request *model.RunQueryRequest) *RunQueryInvoker {
	requestDef := GenReqDefForRunQuery()
	return &RunQueryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStoredQuery 查询单个高级查询
//
// Show Resource Query Language
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ShowStoredQuery(request *model.ShowStoredQueryRequest) (*model.ShowStoredQueryResponse, error) {
	requestDef := GenReqDefForShowStoredQuery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStoredQueryResponse), nil
	}
}

// ShowStoredQueryInvoker 查询单个高级查询
func (c *RmsClient) ShowStoredQueryInvoker(request *model.ShowStoredQueryRequest) *ShowStoredQueryInvoker {
	requestDef := GenReqDefForShowStoredQuery()
	return &ShowStoredQueryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateStoredQuery 更新单个高级查询
//
// 更新自定义查询
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) UpdateStoredQuery(request *model.UpdateStoredQueryRequest) (*model.UpdateStoredQueryResponse, error) {
	requestDef := GenReqDefForUpdateStoredQuery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStoredQueryResponse), nil
	}
}

// UpdateStoredQueryInvoker 更新单个高级查询
func (c *RmsClient) UpdateStoredQueryInvoker(request *model.UpdateStoredQueryRequest) *UpdateStoredQueryInvoker {
	requestDef := GenReqDefForUpdateStoredQuery()
	return &UpdateStoredQueryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRegions 查询用户可见的区域
//
// 查询用户可见的区域
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListRegions(request *model.ListRegionsRequest) (*model.ListRegionsResponse, error) {
	requestDef := GenReqDefForListRegions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRegionsResponse), nil
	}
}

// ListRegionsInvoker 查询用户可见的区域
func (c *RmsClient) ListRegionsInvoker(request *model.ListRegionsRequest) *ListRegionsInvoker {
	requestDef := GenReqDefForListRegions()
	return &ListRegionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceRelations 列举资源关系
//
// 指定资源ID，查询该资源与其他资源的关联关系，可以指定关系方向为\&quot;in\&quot; 或者\&quot;out\&quot;
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ShowResourceRelations(request *model.ShowResourceRelationsRequest) (*model.ShowResourceRelationsResponse, error) {
	requestDef := GenReqDefForShowResourceRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceRelationsResponse), nil
	}
}

// ShowResourceRelationsInvoker 列举资源关系
func (c *RmsClient) ShowResourceRelationsInvoker(request *model.ShowResourceRelationsRequest) *ShowResourceRelationsInvoker {
	requestDef := GenReqDefForShowResourceRelations()
	return &ShowResourceRelationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllResources 列举所有资源
//
// 返回当前用户下所有资源，需要当前用户有rms:resources:list权限。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListAllResources(request *model.ListAllResourcesRequest) (*model.ListAllResourcesResponse, error) {
	requestDef := GenReqDefForListAllResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllResourcesResponse), nil
	}
}

// ListAllResourcesInvoker 列举所有资源
func (c *RmsClient) ListAllResourcesInvoker(request *model.ListAllResourcesRequest) *ListAllResourcesInvoker {
	requestDef := GenReqDefForListAllResources()
	return &ListAllResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProviders 列举云服务
//
// 查询RMS支持的云服务、资源、区域列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListProviders(request *model.ListProvidersRequest) (*model.ListProvidersResponse, error) {
	requestDef := GenReqDefForListProviders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProvidersResponse), nil
	}
}

// ListProvidersInvoker 列举云服务
func (c *RmsClient) ListProvidersInvoker(request *model.ListProvidersRequest) *ListProvidersInvoker {
	requestDef := GenReqDefForListProviders()
	return &ListProvidersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResources 列举指定类型的资源
//
// 返回当前租户下特定资源类型的资源，需要当前用户有rms:resources:list权限。比如查询云服务器，对应的RMS资源类型是ecs.cloudservers，其中provider为ecs，type为cloudservers。 RMS支持的服务和资源类型参见[支持的服务和区域](https://console.huaweicloud.com/eps/#/resources/supported)。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListResources(request *model.ListResourcesRequest) (*model.ListResourcesResponse, error) {
	requestDef := GenReqDefForListResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourcesResponse), nil
	}
}

// ListResourcesInvoker 列举指定类型的资源
func (c *RmsClient) ListResourcesInvoker(request *model.ListResourcesRequest) *ListResourcesInvoker {
	requestDef := GenReqDefForListResources()
	return &ListResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceById 查询单个资源
//
// 指定资源ID，返回该资源的详细信息，需要当前用户有rms:resources:get权限。比如查询云服务器，对应的RMS资源类型是ecs.cloudservers，其中provider为ecs，type为cloudservers。RMS支持的服务和资源类型参见[支持的服务和区域](https://console.huaweicloud.com/eps/#/resources/supported)。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ShowResourceById(request *model.ShowResourceByIdRequest) (*model.ShowResourceByIdResponse, error) {
	requestDef := GenReqDefForShowResourceById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceByIdResponse), nil
	}
}

// ShowResourceByIdInvoker 查询单个资源
func (c *RmsClient) ShowResourceByIdInvoker(request *model.ShowResourceByIdRequest) *ShowResourceByIdInvoker {
	requestDef := GenReqDefForShowResourceById()
	return &ShowResourceByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTrackerConfig 创建或更新记录器
//
// 创建或更新资源记录器，只能存在一个资源记录器
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) CreateTrackerConfig(request *model.CreateTrackerConfigRequest) (*model.CreateTrackerConfigResponse, error) {
	requestDef := GenReqDefForCreateTrackerConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTrackerConfigResponse), nil
	}
}

// CreateTrackerConfigInvoker 创建或更新记录器
func (c *RmsClient) CreateTrackerConfigInvoker(request *model.CreateTrackerConfigRequest) *CreateTrackerConfigInvoker {
	requestDef := GenReqDefForCreateTrackerConfig()
	return &CreateTrackerConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTrackerConfig 删除记录器
//
// 删除资源记录器
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) DeleteTrackerConfig(request *model.DeleteTrackerConfigRequest) (*model.DeleteTrackerConfigResponse, error) {
	requestDef := GenReqDefForDeleteTrackerConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTrackerConfigResponse), nil
	}
}

// DeleteTrackerConfigInvoker 删除记录器
func (c *RmsClient) DeleteTrackerConfigInvoker(request *model.DeleteTrackerConfigRequest) *DeleteTrackerConfigInvoker {
	requestDef := GenReqDefForDeleteTrackerConfig()
	return &DeleteTrackerConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTrackerConfig 查询记录器
//
// 查询资源记录器的详细信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ShowTrackerConfig(request *model.ShowTrackerConfigRequest) (*model.ShowTrackerConfigResponse, error) {
	requestDef := GenReqDefForShowTrackerConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTrackerConfigResponse), nil
	}
}

// ShowTrackerConfigInvoker 查询记录器
func (c *RmsClient) ShowTrackerConfigInvoker(request *model.ShowTrackerConfigRequest) *ShowTrackerConfigInvoker {
	requestDef := GenReqDefForShowTrackerConfig()
	return &ShowTrackerConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
