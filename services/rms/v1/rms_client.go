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

// CreateAggregationAuthorization 创建资源聚合器授权
//
// 给资源聚合器帐号授予从源帐号收集数据的权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) CreateAggregationAuthorization(request *model.CreateAggregationAuthorizationRequest) (*model.CreateAggregationAuthorizationResponse, error) {
	requestDef := GenReqDefForCreateAggregationAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAggregationAuthorizationResponse), nil
	}
}

// CreateAggregationAuthorizationInvoker 创建资源聚合器授权
func (c *RmsClient) CreateAggregationAuthorizationInvoker(request *model.CreateAggregationAuthorizationRequest) *CreateAggregationAuthorizationInvoker {
	requestDef := GenReqDefForCreateAggregationAuthorization()
	return &CreateAggregationAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConfigurationAggregator 创建资源聚合器
//
// 创建资源聚合器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) CreateConfigurationAggregator(request *model.CreateConfigurationAggregatorRequest) (*model.CreateConfigurationAggregatorResponse, error) {
	requestDef := GenReqDefForCreateConfigurationAggregator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConfigurationAggregatorResponse), nil
	}
}

// CreateConfigurationAggregatorInvoker 创建资源聚合器
func (c *RmsClient) CreateConfigurationAggregatorInvoker(request *model.CreateConfigurationAggregatorRequest) *CreateConfigurationAggregatorInvoker {
	requestDef := GenReqDefForCreateConfigurationAggregator()
	return &CreateConfigurationAggregatorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAggregationAuthorization 删除资源聚合器授权
//
// 删除指定资源聚合器帐号的授权。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) DeleteAggregationAuthorization(request *model.DeleteAggregationAuthorizationRequest) (*model.DeleteAggregationAuthorizationResponse, error) {
	requestDef := GenReqDefForDeleteAggregationAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAggregationAuthorizationResponse), nil
	}
}

// DeleteAggregationAuthorizationInvoker 删除资源聚合器授权
func (c *RmsClient) DeleteAggregationAuthorizationInvoker(request *model.DeleteAggregationAuthorizationRequest) *DeleteAggregationAuthorizationInvoker {
	requestDef := GenReqDefForDeleteAggregationAuthorization()
	return &DeleteAggregationAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConfigurationAggregator 删除资源聚合器
//
// 删除资源聚合器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) DeleteConfigurationAggregator(request *model.DeleteConfigurationAggregatorRequest) (*model.DeleteConfigurationAggregatorResponse, error) {
	requestDef := GenReqDefForDeleteConfigurationAggregator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConfigurationAggregatorResponse), nil
	}
}

// DeleteConfigurationAggregatorInvoker 删除资源聚合器
func (c *RmsClient) DeleteConfigurationAggregatorInvoker(request *model.DeleteConfigurationAggregatorRequest) *DeleteConfigurationAggregatorInvoker {
	requestDef := GenReqDefForDeleteConfigurationAggregator()
	return &DeleteConfigurationAggregatorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePendingAggregationRequest 删除聚合器帐号中挂起的授权请求
//
// 删除聚合器帐号中挂起的授权请求。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) DeletePendingAggregationRequest(request *model.DeletePendingAggregationRequestRequest) (*model.DeletePendingAggregationRequestResponse, error) {
	requestDef := GenReqDefForDeletePendingAggregationRequest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePendingAggregationRequestResponse), nil
	}
}

// DeletePendingAggregationRequestInvoker 删除聚合器帐号中挂起的授权请求
func (c *RmsClient) DeletePendingAggregationRequestInvoker(request *model.DeletePendingAggregationRequestRequest) *DeletePendingAggregationRequestInvoker {
	requestDef := GenReqDefForDeletePendingAggregationRequest()
	return &DeletePendingAggregationRequestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAggregateDiscoveredResources 查询聚合器中资源的列表
//
// 查询资源聚合器中特定资源的列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) ListAggregateDiscoveredResources(request *model.ListAggregateDiscoveredResourcesRequest) (*model.ListAggregateDiscoveredResourcesResponse, error) {
	requestDef := GenReqDefForListAggregateDiscoveredResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAggregateDiscoveredResourcesResponse), nil
	}
}

// ListAggregateDiscoveredResourcesInvoker 查询聚合器中资源的列表
func (c *RmsClient) ListAggregateDiscoveredResourcesInvoker(request *model.ListAggregateDiscoveredResourcesRequest) *ListAggregateDiscoveredResourcesInvoker {
	requestDef := GenReqDefForListAggregateDiscoveredResources()
	return &ListAggregateDiscoveredResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAggregationAuthorizations 查询资源聚合器授权列表
//
// 查询授权过的资源聚合器列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) ListAggregationAuthorizations(request *model.ListAggregationAuthorizationsRequest) (*model.ListAggregationAuthorizationsResponse, error) {
	requestDef := GenReqDefForListAggregationAuthorizations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAggregationAuthorizationsResponse), nil
	}
}

// ListAggregationAuthorizationsInvoker 查询资源聚合器授权列表
func (c *RmsClient) ListAggregationAuthorizationsInvoker(request *model.ListAggregationAuthorizationsRequest) *ListAggregationAuthorizationsInvoker {
	requestDef := GenReqDefForListAggregationAuthorizations()
	return &ListAggregationAuthorizationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigurationAggregators 查询资源聚合器列表
//
// 查询资源聚合器列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) ListConfigurationAggregators(request *model.ListConfigurationAggregatorsRequest) (*model.ListConfigurationAggregatorsResponse, error) {
	requestDef := GenReqDefForListConfigurationAggregators()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigurationAggregatorsResponse), nil
	}
}

// ListConfigurationAggregatorsInvoker 查询资源聚合器列表
func (c *RmsClient) ListConfigurationAggregatorsInvoker(request *model.ListConfigurationAggregatorsRequest) *ListConfigurationAggregatorsInvoker {
	requestDef := GenReqDefForListConfigurationAggregators()
	return &ListConfigurationAggregatorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPendingAggregationRequests 查询所有挂起的聚合请求列表
//
// 查询所有挂起的聚合请求列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) ListPendingAggregationRequests(request *model.ListPendingAggregationRequestsRequest) (*model.ListPendingAggregationRequestsResponse, error) {
	requestDef := GenReqDefForListPendingAggregationRequests()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPendingAggregationRequestsResponse), nil
	}
}

// ListPendingAggregationRequestsInvoker 查询所有挂起的聚合请求列表
func (c *RmsClient) ListPendingAggregationRequestsInvoker(request *model.ListPendingAggregationRequestsRequest) *ListPendingAggregationRequestsInvoker {
	requestDef := GenReqDefForListPendingAggregationRequests()
	return &ListPendingAggregationRequestsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunAggregateResourceQuery 对指定聚合器执行高级查询
//
// 对指定聚合器执行高级查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) RunAggregateResourceQuery(request *model.RunAggregateResourceQueryRequest) (*model.RunAggregateResourceQueryResponse, error) {
	requestDef := GenReqDefForRunAggregateResourceQuery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunAggregateResourceQueryResponse), nil
	}
}

// RunAggregateResourceQueryInvoker 对指定聚合器执行高级查询
func (c *RmsClient) RunAggregateResourceQueryInvoker(request *model.RunAggregateResourceQueryRequest) *RunAggregateResourceQueryInvoker {
	requestDef := GenReqDefForRunAggregateResourceQuery()
	return &RunAggregateResourceQueryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAggregateDiscoveredResourceCounts 查询聚合器中帐号资源的计数
//
// 查询聚合器中帐号资源的计数，支持通过过滤器和GroupByKey来统计资源数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) ShowAggregateDiscoveredResourceCounts(request *model.ShowAggregateDiscoveredResourceCountsRequest) (*model.ShowAggregateDiscoveredResourceCountsResponse, error) {
	requestDef := GenReqDefForShowAggregateDiscoveredResourceCounts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAggregateDiscoveredResourceCountsResponse), nil
	}
}

// ShowAggregateDiscoveredResourceCountsInvoker 查询聚合器中帐号资源的计数
func (c *RmsClient) ShowAggregateDiscoveredResourceCountsInvoker(request *model.ShowAggregateDiscoveredResourceCountsRequest) *ShowAggregateDiscoveredResourceCountsInvoker {
	requestDef := GenReqDefForShowAggregateDiscoveredResourceCounts()
	return &ShowAggregateDiscoveredResourceCountsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAggregateResourceConfig 查询源帐号中资源的详情
//
// 查询源帐号中特定资源的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) ShowAggregateResourceConfig(request *model.ShowAggregateResourceConfigRequest) (*model.ShowAggregateResourceConfigResponse, error) {
	requestDef := GenReqDefForShowAggregateResourceConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAggregateResourceConfigResponse), nil
	}
}

// ShowAggregateResourceConfigInvoker 查询源帐号中资源的详情
func (c *RmsClient) ShowAggregateResourceConfigInvoker(request *model.ShowAggregateResourceConfigRequest) *ShowAggregateResourceConfigInvoker {
	requestDef := GenReqDefForShowAggregateResourceConfig()
	return &ShowAggregateResourceConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConfigurationAggregator 查询指定资源聚合器
//
// 查询指定资源聚合器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) ShowConfigurationAggregator(request *model.ShowConfigurationAggregatorRequest) (*model.ShowConfigurationAggregatorResponse, error) {
	requestDef := GenReqDefForShowConfigurationAggregator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigurationAggregatorResponse), nil
	}
}

// ShowConfigurationAggregatorInvoker 查询指定资源聚合器
func (c *RmsClient) ShowConfigurationAggregatorInvoker(request *model.ShowConfigurationAggregatorRequest) *ShowConfigurationAggregatorInvoker {
	requestDef := GenReqDefForShowConfigurationAggregator()
	return &ShowConfigurationAggregatorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConfigurationAggregatorSourcesStatus 查询指定资源聚合器聚合帐号的状态信息
//
// 查询指定资源聚合器聚合帐号的状态信息，状态包括验证源帐号和聚合器帐号之间授权的信息。如果失败，状态包含相关的错误码或消息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) ShowConfigurationAggregatorSourcesStatus(request *model.ShowConfigurationAggregatorSourcesStatusRequest) (*model.ShowConfigurationAggregatorSourcesStatusResponse, error) {
	requestDef := GenReqDefForShowConfigurationAggregatorSourcesStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigurationAggregatorSourcesStatusResponse), nil
	}
}

// ShowConfigurationAggregatorSourcesStatusInvoker 查询指定资源聚合器聚合帐号的状态信息
func (c *RmsClient) ShowConfigurationAggregatorSourcesStatusInvoker(request *model.ShowConfigurationAggregatorSourcesStatusRequest) *ShowConfigurationAggregatorSourcesStatusInvoker {
	requestDef := GenReqDefForShowConfigurationAggregatorSourcesStatus()
	return &ShowConfigurationAggregatorSourcesStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateConfigurationAggregator 更新资源聚合器
//
// 更新资源聚合器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) UpdateConfigurationAggregator(request *model.UpdateConfigurationAggregatorRequest) (*model.UpdateConfigurationAggregatorResponse, error) {
	requestDef := GenReqDefForUpdateConfigurationAggregator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConfigurationAggregatorResponse), nil
	}
}

// UpdateConfigurationAggregatorInvoker 更新资源聚合器
func (c *RmsClient) UpdateConfigurationAggregatorInvoker(request *model.UpdateConfigurationAggregatorRequest) *UpdateConfigurationAggregatorInvoker {
	requestDef := GenReqDefForUpdateConfigurationAggregator()
	return &UpdateConfigurationAggregatorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceHistory 查询资源历史
//
// 查询资源与资源关系的变更历史
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// CreateOrganizationPolicyAssignment 创建或更新组织合规规则
//
// 创建或更新组织合规规则，如果规则名称已存在，则为更新操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) CreateOrganizationPolicyAssignment(request *model.CreateOrganizationPolicyAssignmentRequest) (*model.CreateOrganizationPolicyAssignmentResponse, error) {
	requestDef := GenReqDefForCreateOrganizationPolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrganizationPolicyAssignmentResponse), nil
	}
}

// CreateOrganizationPolicyAssignmentInvoker 创建或更新组织合规规则
func (c *RmsClient) CreateOrganizationPolicyAssignmentInvoker(request *model.CreateOrganizationPolicyAssignmentRequest) *CreateOrganizationPolicyAssignmentInvoker {
	requestDef := GenReqDefForCreateOrganizationPolicyAssignment()
	return &CreateOrganizationPolicyAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePolicyAssignments 创建合规规则
//
// 创建新的合规规则
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// DeleteOrganizationPolicyAssignment 删除组织合规规则
//
// 删除组织合规规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) DeleteOrganizationPolicyAssignment(request *model.DeleteOrganizationPolicyAssignmentRequest) (*model.DeleteOrganizationPolicyAssignmentResponse, error) {
	requestDef := GenReqDefForDeleteOrganizationPolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteOrganizationPolicyAssignmentResponse), nil
	}
}

// DeleteOrganizationPolicyAssignmentInvoker 删除组织合规规则
func (c *RmsClient) DeleteOrganizationPolicyAssignmentInvoker(request *model.DeleteOrganizationPolicyAssignmentRequest) *DeleteOrganizationPolicyAssignmentInvoker {
	requestDef := GenReqDefForDeleteOrganizationPolicyAssignment()
	return &DeleteOrganizationPolicyAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePolicyAssignment 删除合规规则
//
// 根据规则ID删除此规则
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListOrganizationPolicyAssignments 查询组织合规规则列表
//
// 查询组织合规规则列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) ListOrganizationPolicyAssignments(request *model.ListOrganizationPolicyAssignmentsRequest) (*model.ListOrganizationPolicyAssignmentsResponse, error) {
	requestDef := GenReqDefForListOrganizationPolicyAssignments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrganizationPolicyAssignmentsResponse), nil
	}
}

// ListOrganizationPolicyAssignmentsInvoker 查询组织合规规则列表
func (c *RmsClient) ListOrganizationPolicyAssignmentsInvoker(request *model.ListOrganizationPolicyAssignmentsRequest) *ListOrganizationPolicyAssignmentsInvoker {
	requestDef := GenReqDefForListOrganizationPolicyAssignments()
	return &ListOrganizationPolicyAssignmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyAssignments 列出合规规则
//
// 列出用户的合规规则
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowOrganizationPolicyAssignment 查询指定组织合规规则
//
// 查询指定组织合规规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) ShowOrganizationPolicyAssignment(request *model.ShowOrganizationPolicyAssignmentRequest) (*model.ShowOrganizationPolicyAssignmentResponse, error) {
	requestDef := GenReqDefForShowOrganizationPolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOrganizationPolicyAssignmentResponse), nil
	}
}

// ShowOrganizationPolicyAssignmentInvoker 查询指定组织合规规则
func (c *RmsClient) ShowOrganizationPolicyAssignmentInvoker(request *model.ShowOrganizationPolicyAssignmentRequest) *ShowOrganizationPolicyAssignmentInvoker {
	requestDef := GenReqDefForShowOrganizationPolicyAssignment()
	return &ShowOrganizationPolicyAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOrganizationPolicyAssignmentDetailedStatus 查询组织内每个成员帐号合规规则部署的详细状态
//
// 查询组织内每个成员帐号合规规则部署的详细状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) ShowOrganizationPolicyAssignmentDetailedStatus(request *model.ShowOrganizationPolicyAssignmentDetailedStatusRequest) (*model.ShowOrganizationPolicyAssignmentDetailedStatusResponse, error) {
	requestDef := GenReqDefForShowOrganizationPolicyAssignmentDetailedStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOrganizationPolicyAssignmentDetailedStatusResponse), nil
	}
}

// ShowOrganizationPolicyAssignmentDetailedStatusInvoker 查询组织内每个成员帐号合规规则部署的详细状态
func (c *RmsClient) ShowOrganizationPolicyAssignmentDetailedStatusInvoker(request *model.ShowOrganizationPolicyAssignmentDetailedStatusRequest) *ShowOrganizationPolicyAssignmentDetailedStatusInvoker {
	requestDef := GenReqDefForShowOrganizationPolicyAssignmentDetailedStatus()
	return &ShowOrganizationPolicyAssignmentDetailedStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOrganizationPolicyAssignmentStatuses 查询组织合规规则部署状态
//
// 查询组织合规规则部署状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) ShowOrganizationPolicyAssignmentStatuses(request *model.ShowOrganizationPolicyAssignmentStatusesRequest) (*model.ShowOrganizationPolicyAssignmentStatusesResponse, error) {
	requestDef := GenReqDefForShowOrganizationPolicyAssignmentStatuses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOrganizationPolicyAssignmentStatusesResponse), nil
	}
}

// ShowOrganizationPolicyAssignmentStatusesInvoker 查询组织合规规则部署状态
func (c *RmsClient) ShowOrganizationPolicyAssignmentStatusesInvoker(request *model.ShowOrganizationPolicyAssignmentStatusesRequest) *ShowOrganizationPolicyAssignmentStatusesInvoker {
	requestDef := GenReqDefForShowOrganizationPolicyAssignmentStatuses()
	return &ShowOrganizationPolicyAssignmentStatusesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPolicyAssignment 获取单个合规规则
//
// 根据规则ID获取单个规则
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// UpdatePolicyState 更新合规评估结果
//
// 更新用户自定义合规规则的合规评估结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) UpdatePolicyState(request *model.UpdatePolicyStateRequest) (*model.UpdatePolicyStateResponse, error) {
	requestDef := GenReqDefForUpdatePolicyState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePolicyStateResponse), nil
	}
}

// UpdatePolicyStateInvoker 更新合规评估结果
func (c *RmsClient) UpdatePolicyStateInvoker(request *model.UpdatePolicyStateRequest) *UpdatePolicyStateInvoker {
	requestDef := GenReqDefForUpdatePolicyState()
	return &UpdatePolicyStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStoredQuery 创建高级查询
//
// 创建新的高级查询
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowResourceRelationsDetail 列举资源关系详情
//
// 指定资源ID，查询该资源与其他资源的关联关系，可以指定关系方向为“in”或者“out”，需要当帐号有rms:resources:getRelation权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) ShowResourceRelationsDetail(request *model.ShowResourceRelationsDetailRequest) (*model.ShowResourceRelationsDetailResponse, error) {
	requestDef := GenReqDefForShowResourceRelationsDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceRelationsDetailResponse), nil
	}
}

// ShowResourceRelationsDetailInvoker 列举资源关系详情
func (c *RmsClient) ShowResourceRelationsDetailInvoker(request *model.ShowResourceRelationsDetailRequest) *ShowResourceRelationsDetailInvoker {
	requestDef := GenReqDefForShowResourceRelationsDetail()
	return &ShowResourceRelationsDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CollectAllResourcesSummary 列举资源概要
//
// 查询当前帐号的资源概览。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) CollectAllResourcesSummary(request *model.CollectAllResourcesSummaryRequest) (*model.CollectAllResourcesSummaryResponse, error) {
	requestDef := GenReqDefForCollectAllResourcesSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectAllResourcesSummaryResponse), nil
	}
}

// CollectAllResourcesSummaryInvoker 列举资源概要
func (c *RmsClient) CollectAllResourcesSummaryInvoker(request *model.CollectAllResourcesSummaryRequest) *CollectAllResourcesSummaryInvoker {
	requestDef := GenReqDefForCollectAllResourcesSummary()
	return &CollectAllResourcesSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountAllResources 查询资源数量
//
// 查询当前帐号的资源数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) CountAllResources(request *model.CountAllResourcesRequest) (*model.CountAllResourcesResponse, error) {
	requestDef := GenReqDefForCountAllResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountAllResourcesResponse), nil
	}
}

// CountAllResourcesInvoker 查询资源数量
func (c *RmsClient) CountAllResourcesInvoker(request *model.CountAllResourcesRequest) *CountAllResourcesInvoker {
	requestDef := GenReqDefForCountAllResources()
	return &CountAllResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllResources 列举所有资源
//
// 返回当前用户下所有资源，需要当前用户有rms:resources:list权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListAllTags 列举资源标签
//
// 查询当前帐号下所有资源的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) ListAllTags(request *model.ListAllTagsRequest) (*model.ListAllTagsResponse, error) {
	requestDef := GenReqDefForListAllTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllTagsResponse), nil
	}
}

// ListAllTagsInvoker 列举资源标签
func (c *RmsClient) ListAllTagsInvoker(request *model.ListAllTagsRequest) *ListAllTagsInvoker {
	requestDef := GenReqDefForListAllTags()
	return &ListAllTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProviders 列举云服务
//
// 查询RMS支持的云服务、资源、区域列表
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowResourceDetail 查询帐号下的单个资源
//
// 查询当前帐号下的单个资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RmsClient) ShowResourceDetail(request *model.ShowResourceDetailRequest) (*model.ShowResourceDetailResponse, error) {
	requestDef := GenReqDefForShowResourceDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceDetailResponse), nil
	}
}

// ShowResourceDetailInvoker 查询帐号下的单个资源
func (c *RmsClient) ShowResourceDetailInvoker(request *model.ShowResourceDetailRequest) *ShowResourceDetailInvoker {
	requestDef := GenReqDefForShowResourceDetail()
	return &ShowResourceDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTrackerConfig 创建或更新记录器
//
// 创建或更新资源记录器，只能存在一个资源记录器
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
