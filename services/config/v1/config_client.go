package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/config/v1/model"
)

type ConfigClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewConfigClient(hcClient *httpclient.HcHttpClient) *ConfigClient {
	return &ConfigClient{HcClient: hcClient}
}

func ConfigClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// CreateAggregationAuthorization 创建资源聚合器授权
//
// 给资源聚合器帐号授予从源帐号收集数据的权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) CreateAggregationAuthorization(request *model.CreateAggregationAuthorizationRequest) (*model.CreateAggregationAuthorizationResponse, error) {
	requestDef := GenReqDefForCreateAggregationAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAggregationAuthorizationResponse), nil
	}
}

// CreateAggregationAuthorizationInvoker 创建资源聚合器授权
func (c *ConfigClient) CreateAggregationAuthorizationInvoker(request *model.CreateAggregationAuthorizationRequest) *CreateAggregationAuthorizationInvoker {
	requestDef := GenReqDefForCreateAggregationAuthorization()
	return &CreateAggregationAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConfigurationAggregator 创建资源聚合器
//
// 创建资源聚合器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) CreateConfigurationAggregator(request *model.CreateConfigurationAggregatorRequest) (*model.CreateConfigurationAggregatorResponse, error) {
	requestDef := GenReqDefForCreateConfigurationAggregator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConfigurationAggregatorResponse), nil
	}
}

// CreateConfigurationAggregatorInvoker 创建资源聚合器
func (c *ConfigClient) CreateConfigurationAggregatorInvoker(request *model.CreateConfigurationAggregatorRequest) *CreateConfigurationAggregatorInvoker {
	requestDef := GenReqDefForCreateConfigurationAggregator()
	return &CreateConfigurationAggregatorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAggregationAuthorization 删除资源聚合器授权
//
// 删除指定资源聚合器帐号的授权。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) DeleteAggregationAuthorization(request *model.DeleteAggregationAuthorizationRequest) (*model.DeleteAggregationAuthorizationResponse, error) {
	requestDef := GenReqDefForDeleteAggregationAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAggregationAuthorizationResponse), nil
	}
}

// DeleteAggregationAuthorizationInvoker 删除资源聚合器授权
func (c *ConfigClient) DeleteAggregationAuthorizationInvoker(request *model.DeleteAggregationAuthorizationRequest) *DeleteAggregationAuthorizationInvoker {
	requestDef := GenReqDefForDeleteAggregationAuthorization()
	return &DeleteAggregationAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConfigurationAggregator 删除资源聚合器
//
// 删除资源聚合器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) DeleteConfigurationAggregator(request *model.DeleteConfigurationAggregatorRequest) (*model.DeleteConfigurationAggregatorResponse, error) {
	requestDef := GenReqDefForDeleteConfigurationAggregator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConfigurationAggregatorResponse), nil
	}
}

// DeleteConfigurationAggregatorInvoker 删除资源聚合器
func (c *ConfigClient) DeleteConfigurationAggregatorInvoker(request *model.DeleteConfigurationAggregatorRequest) *DeleteConfigurationAggregatorInvoker {
	requestDef := GenReqDefForDeleteConfigurationAggregator()
	return &DeleteConfigurationAggregatorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePendingAggregationRequest 删除聚合器帐号中挂起的授权请求
//
// 删除聚合器帐号中挂起的授权请求。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) DeletePendingAggregationRequest(request *model.DeletePendingAggregationRequestRequest) (*model.DeletePendingAggregationRequestResponse, error) {
	requestDef := GenReqDefForDeletePendingAggregationRequest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePendingAggregationRequestResponse), nil
	}
}

// DeletePendingAggregationRequestInvoker 删除聚合器帐号中挂起的授权请求
func (c *ConfigClient) DeletePendingAggregationRequestInvoker(request *model.DeletePendingAggregationRequestRequest) *DeletePendingAggregationRequestInvoker {
	requestDef := GenReqDefForDeletePendingAggregationRequest()
	return &DeletePendingAggregationRequestInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAggregateComplianceByPolicyAssignment 查询聚合合规规则列表
//
// 查询合规和不合规规则的列表，其中包含合规和不合规规则的资源数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListAggregateComplianceByPolicyAssignment(request *model.ListAggregateComplianceByPolicyAssignmentRequest) (*model.ListAggregateComplianceByPolicyAssignmentResponse, error) {
	requestDef := GenReqDefForListAggregateComplianceByPolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAggregateComplianceByPolicyAssignmentResponse), nil
	}
}

// ListAggregateComplianceByPolicyAssignmentInvoker 查询聚合合规规则列表
func (c *ConfigClient) ListAggregateComplianceByPolicyAssignmentInvoker(request *model.ListAggregateComplianceByPolicyAssignmentRequest) *ListAggregateComplianceByPolicyAssignmentInvoker {
	requestDef := GenReqDefForListAggregateComplianceByPolicyAssignment()
	return &ListAggregateComplianceByPolicyAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAggregateDiscoveredResources 查询聚合器中资源的列表
//
// 查询资源聚合器中特定资源的列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListAggregateDiscoveredResources(request *model.ListAggregateDiscoveredResourcesRequest) (*model.ListAggregateDiscoveredResourcesResponse, error) {
	requestDef := GenReqDefForListAggregateDiscoveredResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAggregateDiscoveredResourcesResponse), nil
	}
}

// ListAggregateDiscoveredResourcesInvoker 查询聚合器中资源的列表
func (c *ConfigClient) ListAggregateDiscoveredResourcesInvoker(request *model.ListAggregateDiscoveredResourcesRequest) *ListAggregateDiscoveredResourcesInvoker {
	requestDef := GenReqDefForListAggregateDiscoveredResources()
	return &ListAggregateDiscoveredResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAggregationAuthorizations 查询资源聚合器授权列表
//
// 查询授权过的资源聚合器列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListAggregationAuthorizations(request *model.ListAggregationAuthorizationsRequest) (*model.ListAggregationAuthorizationsResponse, error) {
	requestDef := GenReqDefForListAggregationAuthorizations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAggregationAuthorizationsResponse), nil
	}
}

// ListAggregationAuthorizationsInvoker 查询资源聚合器授权列表
func (c *ConfigClient) ListAggregationAuthorizationsInvoker(request *model.ListAggregationAuthorizationsRequest) *ListAggregationAuthorizationsInvoker {
	requestDef := GenReqDefForListAggregationAuthorizations()
	return &ListAggregationAuthorizationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigurationAggregators 查询资源聚合器列表
//
// 查询资源聚合器列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListConfigurationAggregators(request *model.ListConfigurationAggregatorsRequest) (*model.ListConfigurationAggregatorsResponse, error) {
	requestDef := GenReqDefForListConfigurationAggregators()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigurationAggregatorsResponse), nil
	}
}

// ListConfigurationAggregatorsInvoker 查询资源聚合器列表
func (c *ConfigClient) ListConfigurationAggregatorsInvoker(request *model.ListConfigurationAggregatorsRequest) *ListConfigurationAggregatorsInvoker {
	requestDef := GenReqDefForListConfigurationAggregators()
	return &ListConfigurationAggregatorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPendingAggregationRequests 查询所有挂起的聚合请求列表
//
// 查询所有挂起的聚合请求列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListPendingAggregationRequests(request *model.ListPendingAggregationRequestsRequest) (*model.ListPendingAggregationRequestsResponse, error) {
	requestDef := GenReqDefForListPendingAggregationRequests()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPendingAggregationRequestsResponse), nil
	}
}

// ListPendingAggregationRequestsInvoker 查询所有挂起的聚合请求列表
func (c *ConfigClient) ListPendingAggregationRequestsInvoker(request *model.ListPendingAggregationRequestsRequest) *ListPendingAggregationRequestsInvoker {
	requestDef := GenReqDefForListPendingAggregationRequests()
	return &ListPendingAggregationRequestsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunAggregateResourceQuery 对指定聚合器执行高级查询
//
// 对指定聚合器执行高级查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) RunAggregateResourceQuery(request *model.RunAggregateResourceQueryRequest) (*model.RunAggregateResourceQueryResponse, error) {
	requestDef := GenReqDefForRunAggregateResourceQuery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunAggregateResourceQueryResponse), nil
	}
}

// RunAggregateResourceQueryInvoker 对指定聚合器执行高级查询
func (c *ConfigClient) RunAggregateResourceQueryInvoker(request *model.RunAggregateResourceQueryRequest) *RunAggregateResourceQueryInvoker {
	requestDef := GenReqDefForRunAggregateResourceQuery()
	return &RunAggregateResourceQueryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAggregateComplianceDetailsByPolicyAssignment 查询指定聚合合规规则的评估结果详情
//
// 返回指定聚合合规规则的评估结果详情。包含评估了哪些资源，以及每个资源是否符合规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowAggregateComplianceDetailsByPolicyAssignment(request *model.ShowAggregateComplianceDetailsByPolicyAssignmentRequest) (*model.ShowAggregateComplianceDetailsByPolicyAssignmentResponse, error) {
	requestDef := GenReqDefForShowAggregateComplianceDetailsByPolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAggregateComplianceDetailsByPolicyAssignmentResponse), nil
	}
}

// ShowAggregateComplianceDetailsByPolicyAssignmentInvoker 查询指定聚合合规规则的评估结果详情
func (c *ConfigClient) ShowAggregateComplianceDetailsByPolicyAssignmentInvoker(request *model.ShowAggregateComplianceDetailsByPolicyAssignmentRequest) *ShowAggregateComplianceDetailsByPolicyAssignmentInvoker {
	requestDef := GenReqDefForShowAggregateComplianceDetailsByPolicyAssignment()
	return &ShowAggregateComplianceDetailsByPolicyAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAggregateDiscoveredResourceCounts 查询聚合器中帐号资源的计数
//
// 查询聚合器中帐号资源的计数，支持通过过滤器和GroupByKey来统计资源数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowAggregateDiscoveredResourceCounts(request *model.ShowAggregateDiscoveredResourceCountsRequest) (*model.ShowAggregateDiscoveredResourceCountsResponse, error) {
	requestDef := GenReqDefForShowAggregateDiscoveredResourceCounts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAggregateDiscoveredResourceCountsResponse), nil
	}
}

// ShowAggregateDiscoveredResourceCountsInvoker 查询聚合器中帐号资源的计数
func (c *ConfigClient) ShowAggregateDiscoveredResourceCountsInvoker(request *model.ShowAggregateDiscoveredResourceCountsRequest) *ShowAggregateDiscoveredResourceCountsInvoker {
	requestDef := GenReqDefForShowAggregateDiscoveredResourceCounts()
	return &ShowAggregateDiscoveredResourceCountsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAggregatePolicyAssignmentDetail 查询指定聚合合规规则详情
//
// 返回指定聚合合规规则详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowAggregatePolicyAssignmentDetail(request *model.ShowAggregatePolicyAssignmentDetailRequest) (*model.ShowAggregatePolicyAssignmentDetailResponse, error) {
	requestDef := GenReqDefForShowAggregatePolicyAssignmentDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAggregatePolicyAssignmentDetailResponse), nil
	}
}

// ShowAggregatePolicyAssignmentDetailInvoker 查询指定聚合合规规则详情
func (c *ConfigClient) ShowAggregatePolicyAssignmentDetailInvoker(request *model.ShowAggregatePolicyAssignmentDetailRequest) *ShowAggregatePolicyAssignmentDetailInvoker {
	requestDef := GenReqDefForShowAggregatePolicyAssignmentDetail()
	return &ShowAggregatePolicyAssignmentDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAggregatePolicyStateComplianceSummary 查询聚合器中一个或多个帐户的合规概况
//
// 查询聚合器中一个或多个帐户的合规和不合规规则数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowAggregatePolicyStateComplianceSummary(request *model.ShowAggregatePolicyStateComplianceSummaryRequest) (*model.ShowAggregatePolicyStateComplianceSummaryResponse, error) {
	requestDef := GenReqDefForShowAggregatePolicyStateComplianceSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAggregatePolicyStateComplianceSummaryResponse), nil
	}
}

// ShowAggregatePolicyStateComplianceSummaryInvoker 查询聚合器中一个或多个帐户的合规概况
func (c *ConfigClient) ShowAggregatePolicyStateComplianceSummaryInvoker(request *model.ShowAggregatePolicyStateComplianceSummaryRequest) *ShowAggregatePolicyStateComplianceSummaryInvoker {
	requestDef := GenReqDefForShowAggregatePolicyStateComplianceSummary()
	return &ShowAggregatePolicyStateComplianceSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAggregateResourceConfig 查询源帐号中资源的详情
//
// 查询源帐号中特定资源的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowAggregateResourceConfig(request *model.ShowAggregateResourceConfigRequest) (*model.ShowAggregateResourceConfigResponse, error) {
	requestDef := GenReqDefForShowAggregateResourceConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAggregateResourceConfigResponse), nil
	}
}

// ShowAggregateResourceConfigInvoker 查询源帐号中资源的详情
func (c *ConfigClient) ShowAggregateResourceConfigInvoker(request *model.ShowAggregateResourceConfigRequest) *ShowAggregateResourceConfigInvoker {
	requestDef := GenReqDefForShowAggregateResourceConfig()
	return &ShowAggregateResourceConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConfigurationAggregator 查询指定资源聚合器
//
// 查询指定资源聚合器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowConfigurationAggregator(request *model.ShowConfigurationAggregatorRequest) (*model.ShowConfigurationAggregatorResponse, error) {
	requestDef := GenReqDefForShowConfigurationAggregator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigurationAggregatorResponse), nil
	}
}

// ShowConfigurationAggregatorInvoker 查询指定资源聚合器
func (c *ConfigClient) ShowConfigurationAggregatorInvoker(request *model.ShowConfigurationAggregatorRequest) *ShowConfigurationAggregatorInvoker {
	requestDef := GenReqDefForShowConfigurationAggregator()
	return &ShowConfigurationAggregatorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConfigurationAggregatorSourcesStatus 查询指定资源聚合器聚合帐号的状态信息
//
// 查询指定资源聚合器聚合帐号的状态信息，状态包括验证源帐号和聚合器帐号之间授权的信息。如果失败，状态包含相关的错误码或消息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowConfigurationAggregatorSourcesStatus(request *model.ShowConfigurationAggregatorSourcesStatusRequest) (*model.ShowConfigurationAggregatorSourcesStatusResponse, error) {
	requestDef := GenReqDefForShowConfigurationAggregatorSourcesStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigurationAggregatorSourcesStatusResponse), nil
	}
}

// ShowConfigurationAggregatorSourcesStatusInvoker 查询指定资源聚合器聚合帐号的状态信息
func (c *ConfigClient) ShowConfigurationAggregatorSourcesStatusInvoker(request *model.ShowConfigurationAggregatorSourcesStatusRequest) *ShowConfigurationAggregatorSourcesStatusInvoker {
	requestDef := GenReqDefForShowConfigurationAggregatorSourcesStatus()
	return &ShowConfigurationAggregatorSourcesStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateConfigurationAggregator 更新资源聚合器
//
// 更新资源聚合器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) UpdateConfigurationAggregator(request *model.UpdateConfigurationAggregatorRequest) (*model.UpdateConfigurationAggregatorResponse, error) {
	requestDef := GenReqDefForUpdateConfigurationAggregator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConfigurationAggregatorResponse), nil
	}
}

// UpdateConfigurationAggregatorInvoker 更新资源聚合器
func (c *ConfigClient) UpdateConfigurationAggregatorInvoker(request *model.UpdateConfigurationAggregatorRequest) *UpdateConfigurationAggregatorInvoker {
	requestDef := GenReqDefForUpdateConfigurationAggregator()
	return &UpdateConfigurationAggregatorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CollectConformancePackComplianceSummary 列举合规规则包的结果概览
//
// 列举用户的合规规则包的合规结果概览。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) CollectConformancePackComplianceSummary(request *model.CollectConformancePackComplianceSummaryRequest) (*model.CollectConformancePackComplianceSummaryResponse, error) {
	requestDef := GenReqDefForCollectConformancePackComplianceSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectConformancePackComplianceSummaryResponse), nil
	}
}

// CollectConformancePackComplianceSummaryInvoker 列举合规规则包的结果概览
func (c *ConfigClient) CollectConformancePackComplianceSummaryInvoker(request *model.CollectConformancePackComplianceSummaryRequest) *CollectConformancePackComplianceSummaryInvoker {
	requestDef := GenReqDefForCollectConformancePackComplianceSummary()
	return &CollectConformancePackComplianceSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConformancePack 创建合规规则包
//
// 创建新的合规规则包。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) CreateConformancePack(request *model.CreateConformancePackRequest) (*model.CreateConformancePackResponse, error) {
	requestDef := GenReqDefForCreateConformancePack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConformancePackResponse), nil
	}
}

// CreateConformancePackInvoker 创建合规规则包
func (c *ConfigClient) CreateConformancePackInvoker(request *model.CreateConformancePackRequest) *CreateConformancePackInvoker {
	requestDef := GenReqDefForCreateConformancePack()
	return &CreateConformancePackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOrganizationConformancePack 创建组织合规规则包
//
// 创建新的组织合规规则包。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) CreateOrganizationConformancePack(request *model.CreateOrganizationConformancePackRequest) (*model.CreateOrganizationConformancePackResponse, error) {
	requestDef := GenReqDefForCreateOrganizationConformancePack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrganizationConformancePackResponse), nil
	}
}

// CreateOrganizationConformancePackInvoker 创建组织合规规则包
func (c *ConfigClient) CreateOrganizationConformancePackInvoker(request *model.CreateOrganizationConformancePackRequest) *CreateOrganizationConformancePackInvoker {
	requestDef := GenReqDefForCreateOrganizationConformancePack()
	return &CreateOrganizationConformancePackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConformancePack 删除合规规则包
//
// 删除用户的合规规则包。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) DeleteConformancePack(request *model.DeleteConformancePackRequest) (*model.DeleteConformancePackResponse, error) {
	requestDef := GenReqDefForDeleteConformancePack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConformancePackResponse), nil
	}
}

// DeleteConformancePackInvoker 删除合规规则包
func (c *ConfigClient) DeleteConformancePackInvoker(request *model.DeleteConformancePackRequest) *DeleteConformancePackInvoker {
	requestDef := GenReqDefForDeleteConformancePack()
	return &DeleteConformancePackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteOrganizationConformancePack 删除组织合规规则包
//
// 删除用户的组织合规规则包。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) DeleteOrganizationConformancePack(request *model.DeleteOrganizationConformancePackRequest) (*model.DeleteOrganizationConformancePackResponse, error) {
	requestDef := GenReqDefForDeleteOrganizationConformancePack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteOrganizationConformancePackResponse), nil
	}
}

// DeleteOrganizationConformancePackInvoker 删除组织合规规则包
func (c *ConfigClient) DeleteOrganizationConformancePackInvoker(request *model.DeleteOrganizationConformancePackRequest) *DeleteOrganizationConformancePackInvoker {
	requestDef := GenReqDefForDeleteOrganizationConformancePack()
	return &DeleteOrganizationConformancePackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBuiltInConformancePackTemplates 列举预定义合规规则包模板
//
// 列举预定义的合规规则包的模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListBuiltInConformancePackTemplates(request *model.ListBuiltInConformancePackTemplatesRequest) (*model.ListBuiltInConformancePackTemplatesResponse, error) {
	requestDef := GenReqDefForListBuiltInConformancePackTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBuiltInConformancePackTemplatesResponse), nil
	}
}

// ListBuiltInConformancePackTemplatesInvoker 列举预定义合规规则包模板
func (c *ConfigClient) ListBuiltInConformancePackTemplatesInvoker(request *model.ListBuiltInConformancePackTemplatesRequest) *ListBuiltInConformancePackTemplatesInvoker {
	requestDef := GenReqDefForListBuiltInConformancePackTemplates()
	return &ListBuiltInConformancePackTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConformancePackComplianceByPackId 列举合规规则包的评估结果
//
// 列举合规规则包的合规规则评估结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListConformancePackComplianceByPackId(request *model.ListConformancePackComplianceByPackIdRequest) (*model.ListConformancePackComplianceByPackIdResponse, error) {
	requestDef := GenReqDefForListConformancePackComplianceByPackId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConformancePackComplianceByPackIdResponse), nil
	}
}

// ListConformancePackComplianceByPackIdInvoker 列举合规规则包的评估结果
func (c *ConfigClient) ListConformancePackComplianceByPackIdInvoker(request *model.ListConformancePackComplianceByPackIdRequest) *ListConformancePackComplianceByPackIdInvoker {
	requestDef := GenReqDefForListConformancePackComplianceByPackId()
	return &ListConformancePackComplianceByPackIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConformancePackComplianceDetailsByPackId 列举合规规则包的评估结果详情
//
// 列举合规规则包的合规规则评估结果详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListConformancePackComplianceDetailsByPackId(request *model.ListConformancePackComplianceDetailsByPackIdRequest) (*model.ListConformancePackComplianceDetailsByPackIdResponse, error) {
	requestDef := GenReqDefForListConformancePackComplianceDetailsByPackId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConformancePackComplianceDetailsByPackIdResponse), nil
	}
}

// ListConformancePackComplianceDetailsByPackIdInvoker 列举合规规则包的评估结果详情
func (c *ConfigClient) ListConformancePackComplianceDetailsByPackIdInvoker(request *model.ListConformancePackComplianceDetailsByPackIdRequest) *ListConformancePackComplianceDetailsByPackIdInvoker {
	requestDef := GenReqDefForListConformancePackComplianceDetailsByPackId()
	return &ListConformancePackComplianceDetailsByPackIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConformancePackComplianceScores 列举合规规则包分数
//
// 列举用户的合规规则包分数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListConformancePackComplianceScores(request *model.ListConformancePackComplianceScoresRequest) (*model.ListConformancePackComplianceScoresResponse, error) {
	requestDef := GenReqDefForListConformancePackComplianceScores()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConformancePackComplianceScoresResponse), nil
	}
}

// ListConformancePackComplianceScoresInvoker 列举合规规则包分数
func (c *ConfigClient) ListConformancePackComplianceScoresInvoker(request *model.ListConformancePackComplianceScoresRequest) *ListConformancePackComplianceScoresInvoker {
	requestDef := GenReqDefForListConformancePackComplianceScores()
	return &ListConformancePackComplianceScoresInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConformancePacks 列举合规规则包
//
// 列举用户的合规规则包。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListConformancePacks(request *model.ListConformancePacksRequest) (*model.ListConformancePacksResponse, error) {
	requestDef := GenReqDefForListConformancePacks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConformancePacksResponse), nil
	}
}

// ListConformancePacksInvoker 列举合规规则包
func (c *ConfigClient) ListConformancePacksInvoker(request *model.ListConformancePacksRequest) *ListConformancePacksInvoker {
	requestDef := GenReqDefForListConformancePacks()
	return &ListConformancePacksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOrganizationConformancePackStatuses 查看组织合规规则包部署状态
//
// 列举用户的组织合规规则包部署状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListOrganizationConformancePackStatuses(request *model.ListOrganizationConformancePackStatusesRequest) (*model.ListOrganizationConformancePackStatusesResponse, error) {
	requestDef := GenReqDefForListOrganizationConformancePackStatuses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrganizationConformancePackStatusesResponse), nil
	}
}

// ListOrganizationConformancePackStatusesInvoker 查看组织合规规则包部署状态
func (c *ConfigClient) ListOrganizationConformancePackStatusesInvoker(request *model.ListOrganizationConformancePackStatusesRequest) *ListOrganizationConformancePackStatusesInvoker {
	requestDef := GenReqDefForListOrganizationConformancePackStatuses()
	return &ListOrganizationConformancePackStatusesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOrganizationConformancePacks 列举组织合规规则包
//
// 列举用户的组织合规规则包。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListOrganizationConformancePacks(request *model.ListOrganizationConformancePacksRequest) (*model.ListOrganizationConformancePacksResponse, error) {
	requestDef := GenReqDefForListOrganizationConformancePacks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrganizationConformancePacksResponse), nil
	}
}

// ListOrganizationConformancePacksInvoker 列举组织合规规则包
func (c *ConfigClient) ListOrganizationConformancePacksInvoker(request *model.ListOrganizationConformancePacksRequest) *ListOrganizationConformancePacksInvoker {
	requestDef := GenReqDefForListOrganizationConformancePacks()
	return &ListOrganizationConformancePacksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBuiltInConformancePackTemplate 查看预定义合规规则包模板
//
// 根据ID获取单个预定义合规规则包模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowBuiltInConformancePackTemplate(request *model.ShowBuiltInConformancePackTemplateRequest) (*model.ShowBuiltInConformancePackTemplateResponse, error) {
	requestDef := GenReqDefForShowBuiltInConformancePackTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBuiltInConformancePackTemplateResponse), nil
	}
}

// ShowBuiltInConformancePackTemplateInvoker 查看预定义合规规则包模板
func (c *ConfigClient) ShowBuiltInConformancePackTemplateInvoker(request *model.ShowBuiltInConformancePackTemplateRequest) *ShowBuiltInConformancePackTemplateInvoker {
	requestDef := GenReqDefForShowBuiltInConformancePackTemplate()
	return &ShowBuiltInConformancePackTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConformancePack 查看合规规则包
//
// 根据ID获取单个合规规则包。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowConformancePack(request *model.ShowConformancePackRequest) (*model.ShowConformancePackResponse, error) {
	requestDef := GenReqDefForShowConformancePack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConformancePackResponse), nil
	}
}

// ShowConformancePackInvoker 查看合规规则包
func (c *ConfigClient) ShowConformancePackInvoker(request *model.ShowConformancePackRequest) *ShowConformancePackInvoker {
	requestDef := GenReqDefForShowConformancePack()
	return &ShowConformancePackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOrganizationConformancePack 查看组织合规规则包
//
// 根据ID获取单个组织合规规则包详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowOrganizationConformancePack(request *model.ShowOrganizationConformancePackRequest) (*model.ShowOrganizationConformancePackResponse, error) {
	requestDef := GenReqDefForShowOrganizationConformancePack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOrganizationConformancePackResponse), nil
	}
}

// ShowOrganizationConformancePackInvoker 查看组织合规规则包
func (c *ConfigClient) ShowOrganizationConformancePackInvoker(request *model.ShowOrganizationConformancePackRequest) *ShowOrganizationConformancePackInvoker {
	requestDef := GenReqDefForShowOrganizationConformancePack()
	return &ShowOrganizationConformancePackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOrganizationConformancePackDetailedStatuses 查看组织合规规则包部署详细状态
//
// 查看指定组织合规规则包在成员帐号中的部署状态详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowOrganizationConformancePackDetailedStatuses(request *model.ShowOrganizationConformancePackDetailedStatusesRequest) (*model.ShowOrganizationConformancePackDetailedStatusesResponse, error) {
	requestDef := GenReqDefForShowOrganizationConformancePackDetailedStatuses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOrganizationConformancePackDetailedStatusesResponse), nil
	}
}

// ShowOrganizationConformancePackDetailedStatusesInvoker 查看组织合规规则包部署详细状态
func (c *ConfigClient) ShowOrganizationConformancePackDetailedStatusesInvoker(request *model.ShowOrganizationConformancePackDetailedStatusesRequest) *ShowOrganizationConformancePackDetailedStatusesInvoker {
	requestDef := GenReqDefForShowOrganizationConformancePackDetailedStatuses()
	return &ShowOrganizationConformancePackDetailedStatusesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateConformancePack 更新合规规则包
//
// 更新用户的合规规则包。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) UpdateConformancePack(request *model.UpdateConformancePackRequest) (*model.UpdateConformancePackResponse, error) {
	requestDef := GenReqDefForUpdateConformancePack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConformancePackResponse), nil
	}
}

// UpdateConformancePackInvoker 更新合规规则包
func (c *ConfigClient) UpdateConformancePackInvoker(request *model.UpdateConformancePackRequest) *UpdateConformancePackInvoker {
	requestDef := GenReqDefForUpdateConformancePack()
	return &UpdateConformancePackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateOrganizationConformancePack 更新组织合规规则包
//
// 更新用户的组织合规规则包。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) UpdateOrganizationConformancePack(request *model.UpdateOrganizationConformancePackRequest) (*model.UpdateOrganizationConformancePackResponse, error) {
	requestDef := GenReqDefForUpdateOrganizationConformancePack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateOrganizationConformancePackResponse), nil
	}
}

// UpdateOrganizationConformancePackInvoker 更新组织合规规则包
func (c *ConfigClient) UpdateOrganizationConformancePackInvoker(request *model.UpdateOrganizationConformancePackRequest) *UpdateOrganizationConformancePackInvoker {
	requestDef := GenReqDefForUpdateOrganizationConformancePack()
	return &UpdateOrganizationConformancePackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceHistory 查询资源历史
//
// 查询资源与资源关系的变更历史
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowResourceHistory(request *model.ShowResourceHistoryRequest) (*model.ShowResourceHistoryResponse, error) {
	requestDef := GenReqDefForShowResourceHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceHistoryResponse), nil
	}
}

// ShowResourceHistoryInvoker 查询资源历史
func (c *ConfigClient) ShowResourceHistoryInvoker(request *model.ShowResourceHistoryRequest) *ShowResourceHistoryInvoker {
	requestDef := GenReqDefForShowResourceHistory()
	return &ShowResourceHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateRemediationExceptions 批量创建修正例外
//
// 批量创建合规规则修正例外。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) BatchCreateRemediationExceptions(request *model.BatchCreateRemediationExceptionsRequest) (*model.BatchCreateRemediationExceptionsResponse, error) {
	requestDef := GenReqDefForBatchCreateRemediationExceptions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateRemediationExceptionsResponse), nil
	}
}

// BatchCreateRemediationExceptionsInvoker 批量创建修正例外
func (c *ConfigClient) BatchCreateRemediationExceptionsInvoker(request *model.BatchCreateRemediationExceptionsRequest) *BatchCreateRemediationExceptionsInvoker {
	requestDef := GenReqDefForBatchCreateRemediationExceptions()
	return &BatchCreateRemediationExceptionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteRemediationExceptions 批量删除修正例外
//
// 批量删除合规规则修正例外。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) BatchDeleteRemediationExceptions(request *model.BatchDeleteRemediationExceptionsRequest) (*model.BatchDeleteRemediationExceptionsResponse, error) {
	requestDef := GenReqDefForBatchDeleteRemediationExceptions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteRemediationExceptionsResponse), nil
	}
}

// BatchDeleteRemediationExceptionsInvoker 批量删除修正例外
func (c *ConfigClient) BatchDeleteRemediationExceptionsInvoker(request *model.BatchDeleteRemediationExceptionsRequest) *BatchDeleteRemediationExceptionsInvoker {
	requestDef := GenReqDefForBatchDeleteRemediationExceptions()
	return &BatchDeleteRemediationExceptionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CollectRemediationExecutionStatusesSummary 列举修正最新记录
//
// 列举合规规则修正最新记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) CollectRemediationExecutionStatusesSummary(request *model.CollectRemediationExecutionStatusesSummaryRequest) (*model.CollectRemediationExecutionStatusesSummaryResponse, error) {
	requestDef := GenReqDefForCollectRemediationExecutionStatusesSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectRemediationExecutionStatusesSummaryResponse), nil
	}
}

// CollectRemediationExecutionStatusesSummaryInvoker 列举修正最新记录
func (c *ConfigClient) CollectRemediationExecutionStatusesSummaryInvoker(request *model.CollectRemediationExecutionStatusesSummaryRequest) *CollectRemediationExecutionStatusesSummaryInvoker {
	requestDef := GenReqDefForCollectRemediationExecutionStatusesSummary()
	return &CollectRemediationExecutionStatusesSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOrUpdateRemediationConfiguration 创建或更新修正配置
//
// 创建或更新合规规则修正配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) CreateOrUpdateRemediationConfiguration(request *model.CreateOrUpdateRemediationConfigurationRequest) (*model.CreateOrUpdateRemediationConfigurationResponse, error) {
	requestDef := GenReqDefForCreateOrUpdateRemediationConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrUpdateRemediationConfigurationResponse), nil
	}
}

// CreateOrUpdateRemediationConfigurationInvoker 创建或更新修正配置
func (c *ConfigClient) CreateOrUpdateRemediationConfigurationInvoker(request *model.CreateOrUpdateRemediationConfigurationRequest) *CreateOrUpdateRemediationConfigurationInvoker {
	requestDef := GenReqDefForCreateOrUpdateRemediationConfiguration()
	return &CreateOrUpdateRemediationConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOrganizationPolicyAssignment 创建组织合规规则
//
// 创建组织合规规则，如果规则名称已存在，则为更新操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) CreateOrganizationPolicyAssignment(request *model.CreateOrganizationPolicyAssignmentRequest) (*model.CreateOrganizationPolicyAssignmentResponse, error) {
	requestDef := GenReqDefForCreateOrganizationPolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrganizationPolicyAssignmentResponse), nil
	}
}

// CreateOrganizationPolicyAssignmentInvoker 创建组织合规规则
func (c *ConfigClient) CreateOrganizationPolicyAssignmentInvoker(request *model.CreateOrganizationPolicyAssignmentRequest) *CreateOrganizationPolicyAssignmentInvoker {
	requestDef := GenReqDefForCreateOrganizationPolicyAssignment()
	return &CreateOrganizationPolicyAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePolicyAssignments 创建合规规则
//
// 创建新的合规规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) CreatePolicyAssignments(request *model.CreatePolicyAssignmentsRequest) (*model.CreatePolicyAssignmentsResponse, error) {
	requestDef := GenReqDefForCreatePolicyAssignments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePolicyAssignmentsResponse), nil
	}
}

// CreatePolicyAssignmentsInvoker 创建合规规则
func (c *ConfigClient) CreatePolicyAssignmentsInvoker(request *model.CreatePolicyAssignmentsRequest) *CreatePolicyAssignmentsInvoker {
	requestDef := GenReqDefForCreatePolicyAssignments()
	return &CreatePolicyAssignmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteOrganizationPolicyAssignment 删除组织合规规则
//
// 删除组织合规规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) DeleteOrganizationPolicyAssignment(request *model.DeleteOrganizationPolicyAssignmentRequest) (*model.DeleteOrganizationPolicyAssignmentResponse, error) {
	requestDef := GenReqDefForDeleteOrganizationPolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteOrganizationPolicyAssignmentResponse), nil
	}
}

// DeleteOrganizationPolicyAssignmentInvoker 删除组织合规规则
func (c *ConfigClient) DeleteOrganizationPolicyAssignmentInvoker(request *model.DeleteOrganizationPolicyAssignmentRequest) *DeleteOrganizationPolicyAssignmentInvoker {
	requestDef := GenReqDefForDeleteOrganizationPolicyAssignment()
	return &DeleteOrganizationPolicyAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePolicyAssignment 删除合规规则
//
// 根据规则ID删除此规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) DeletePolicyAssignment(request *model.DeletePolicyAssignmentRequest) (*model.DeletePolicyAssignmentResponse, error) {
	requestDef := GenReqDefForDeletePolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePolicyAssignmentResponse), nil
	}
}

// DeletePolicyAssignmentInvoker 删除合规规则
func (c *ConfigClient) DeletePolicyAssignmentInvoker(request *model.DeletePolicyAssignmentRequest) *DeletePolicyAssignmentInvoker {
	requestDef := GenReqDefForDeletePolicyAssignment()
	return &DeletePolicyAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRemediationConfiguration 删除修正配置
//
// 删除合规规则修正配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) DeleteRemediationConfiguration(request *model.DeleteRemediationConfigurationRequest) (*model.DeleteRemediationConfigurationResponse, error) {
	requestDef := GenReqDefForDeleteRemediationConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRemediationConfigurationResponse), nil
	}
}

// DeleteRemediationConfigurationInvoker 删除修正配置
func (c *ConfigClient) DeleteRemediationConfigurationInvoker(request *model.DeleteRemediationConfigurationRequest) *DeleteRemediationConfigurationInvoker {
	requestDef := GenReqDefForDeleteRemediationConfiguration()
	return &DeleteRemediationConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisablePolicyAssignment 停用合规规则
//
// 根据规则ID停用此规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) DisablePolicyAssignment(request *model.DisablePolicyAssignmentRequest) (*model.DisablePolicyAssignmentResponse, error) {
	requestDef := GenReqDefForDisablePolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisablePolicyAssignmentResponse), nil
	}
}

// DisablePolicyAssignmentInvoker 停用合规规则
func (c *ConfigClient) DisablePolicyAssignmentInvoker(request *model.DisablePolicyAssignmentRequest) *DisablePolicyAssignmentInvoker {
	requestDef := GenReqDefForDisablePolicyAssignment()
	return &DisablePolicyAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnablePolicyAssignment 启用合规规则
//
// 根据规则ID启用此规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) EnablePolicyAssignment(request *model.EnablePolicyAssignmentRequest) (*model.EnablePolicyAssignmentResponse, error) {
	requestDef := GenReqDefForEnablePolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnablePolicyAssignmentResponse), nil
	}
}

// EnablePolicyAssignmentInvoker 启用合规规则
func (c *ConfigClient) EnablePolicyAssignmentInvoker(request *model.EnablePolicyAssignmentRequest) *EnablePolicyAssignmentInvoker {
	requestDef := GenReqDefForEnablePolicyAssignment()
	return &EnablePolicyAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBuiltInPolicyDefinitions 列出内置策略
//
// 列出用户的内置策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListBuiltInPolicyDefinitions(request *model.ListBuiltInPolicyDefinitionsRequest) (*model.ListBuiltInPolicyDefinitionsResponse, error) {
	requestDef := GenReqDefForListBuiltInPolicyDefinitions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBuiltInPolicyDefinitionsResponse), nil
	}
}

// ListBuiltInPolicyDefinitionsInvoker 列出内置策略
func (c *ConfigClient) ListBuiltInPolicyDefinitionsInvoker(request *model.ListBuiltInPolicyDefinitionsRequest) *ListBuiltInPolicyDefinitionsInvoker {
	requestDef := GenReqDefForListBuiltInPolicyDefinitions()
	return &ListBuiltInPolicyDefinitionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOrganizationPolicyAssignments 查询组织合规规则列表
//
// 查询组织合规规则列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListOrganizationPolicyAssignments(request *model.ListOrganizationPolicyAssignmentsRequest) (*model.ListOrganizationPolicyAssignmentsResponse, error) {
	requestDef := GenReqDefForListOrganizationPolicyAssignments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrganizationPolicyAssignmentsResponse), nil
	}
}

// ListOrganizationPolicyAssignmentsInvoker 查询组织合规规则列表
func (c *ConfigClient) ListOrganizationPolicyAssignmentsInvoker(request *model.ListOrganizationPolicyAssignmentsRequest) *ListOrganizationPolicyAssignmentsInvoker {
	requestDef := GenReqDefForListOrganizationPolicyAssignments()
	return &ListOrganizationPolicyAssignmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyAssignments 列出合规规则
//
// 列出用户的合规规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListPolicyAssignments(request *model.ListPolicyAssignmentsRequest) (*model.ListPolicyAssignmentsResponse, error) {
	requestDef := GenReqDefForListPolicyAssignments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyAssignmentsResponse), nil
	}
}

// ListPolicyAssignmentsInvoker 列出合规规则
func (c *ConfigClient) ListPolicyAssignmentsInvoker(request *model.ListPolicyAssignmentsRequest) *ListPolicyAssignmentsInvoker {
	requestDef := GenReqDefForListPolicyAssignments()
	return &ListPolicyAssignmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyStatesByAssignmentId 获取规则的合规结果
//
// 根据规则ID查询所有的合规结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListPolicyStatesByAssignmentId(request *model.ListPolicyStatesByAssignmentIdRequest) (*model.ListPolicyStatesByAssignmentIdResponse, error) {
	requestDef := GenReqDefForListPolicyStatesByAssignmentId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyStatesByAssignmentIdResponse), nil
	}
}

// ListPolicyStatesByAssignmentIdInvoker 获取规则的合规结果
func (c *ConfigClient) ListPolicyStatesByAssignmentIdInvoker(request *model.ListPolicyStatesByAssignmentIdRequest) *ListPolicyStatesByAssignmentIdInvoker {
	requestDef := GenReqDefForListPolicyStatesByAssignmentId()
	return &ListPolicyStatesByAssignmentIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyStatesByDomainId 获取用户的合规结果
//
// 查询用户所有的合规结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListPolicyStatesByDomainId(request *model.ListPolicyStatesByDomainIdRequest) (*model.ListPolicyStatesByDomainIdResponse, error) {
	requestDef := GenReqDefForListPolicyStatesByDomainId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyStatesByDomainIdResponse), nil
	}
}

// ListPolicyStatesByDomainIdInvoker 获取用户的合规结果
func (c *ConfigClient) ListPolicyStatesByDomainIdInvoker(request *model.ListPolicyStatesByDomainIdRequest) *ListPolicyStatesByDomainIdInvoker {
	requestDef := GenReqDefForListPolicyStatesByDomainId()
	return &ListPolicyStatesByDomainIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyStatesByResourceId 获取资源的合规结果
//
// 根据资源ID查询所有合规结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListPolicyStatesByResourceId(request *model.ListPolicyStatesByResourceIdRequest) (*model.ListPolicyStatesByResourceIdResponse, error) {
	requestDef := GenReqDefForListPolicyStatesByResourceId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyStatesByResourceIdResponse), nil
	}
}

// ListPolicyStatesByResourceIdInvoker 获取资源的合规结果
func (c *ConfigClient) ListPolicyStatesByResourceIdInvoker(request *model.ListPolicyStatesByResourceIdRequest) *ListPolicyStatesByResourceIdInvoker {
	requestDef := GenReqDefForListPolicyStatesByResourceId()
	return &ListPolicyStatesByResourceIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRemediationExceptions 查询修正例外
//
// 查询合规规则修正例外。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListRemediationExceptions(request *model.ListRemediationExceptionsRequest) (*model.ListRemediationExceptionsResponse, error) {
	requestDef := GenReqDefForListRemediationExceptions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRemediationExceptionsResponse), nil
	}
}

// ListRemediationExceptionsInvoker 查询修正例外
func (c *ConfigClient) ListRemediationExceptionsInvoker(request *model.ListRemediationExceptionsRequest) *ListRemediationExceptionsInvoker {
	requestDef := GenReqDefForListRemediationExceptions()
	return &ListRemediationExceptionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRemediationExecutionStatuses 查询修正执行结果
//
// 查询合规规则修正执行结果详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListRemediationExecutionStatuses(request *model.ListRemediationExecutionStatusesRequest) (*model.ListRemediationExecutionStatusesResponse, error) {
	requestDef := GenReqDefForListRemediationExecutionStatuses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRemediationExecutionStatusesResponse), nil
	}
}

// ListRemediationExecutionStatusesInvoker 查询修正执行结果
func (c *ConfigClient) ListRemediationExecutionStatusesInvoker(request *model.ListRemediationExecutionStatusesRequest) *ListRemediationExecutionStatusesInvoker {
	requestDef := GenReqDefForListRemediationExecutionStatuses()
	return &ListRemediationExecutionStatusesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunEvaluationByPolicyAssignmentId 运行合规评估
//
// 根据规则ID评估此规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) RunEvaluationByPolicyAssignmentId(request *model.RunEvaluationByPolicyAssignmentIdRequest) (*model.RunEvaluationByPolicyAssignmentIdResponse, error) {
	requestDef := GenReqDefForRunEvaluationByPolicyAssignmentId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunEvaluationByPolicyAssignmentIdResponse), nil
	}
}

// RunEvaluationByPolicyAssignmentIdInvoker 运行合规评估
func (c *ConfigClient) RunEvaluationByPolicyAssignmentIdInvoker(request *model.RunEvaluationByPolicyAssignmentIdRequest) *RunEvaluationByPolicyAssignmentIdInvoker {
	requestDef := GenReqDefForRunEvaluationByPolicyAssignmentId()
	return &RunEvaluationByPolicyAssignmentIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunRemediationExecution 运行修正执行
//
// 手动运行合规规则修正执行。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) RunRemediationExecution(request *model.RunRemediationExecutionRequest) (*model.RunRemediationExecutionResponse, error) {
	requestDef := GenReqDefForRunRemediationExecution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunRemediationExecutionResponse), nil
	}
}

// RunRemediationExecutionInvoker 运行修正执行
func (c *ConfigClient) RunRemediationExecutionInvoker(request *model.RunRemediationExecutionRequest) *RunRemediationExecutionInvoker {
	requestDef := GenReqDefForRunRemediationExecution()
	return &RunRemediationExecutionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBuiltInPolicyDefinition 查询单个内置策略
//
// 根据策略ID查询单个内置策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowBuiltInPolicyDefinition(request *model.ShowBuiltInPolicyDefinitionRequest) (*model.ShowBuiltInPolicyDefinitionResponse, error) {
	requestDef := GenReqDefForShowBuiltInPolicyDefinition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBuiltInPolicyDefinitionResponse), nil
	}
}

// ShowBuiltInPolicyDefinitionInvoker 查询单个内置策略
func (c *ConfigClient) ShowBuiltInPolicyDefinitionInvoker(request *model.ShowBuiltInPolicyDefinitionRequest) *ShowBuiltInPolicyDefinitionInvoker {
	requestDef := GenReqDefForShowBuiltInPolicyDefinition()
	return &ShowBuiltInPolicyDefinitionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEvaluationStateByAssignmentId 获取规则的评估状态
//
// 根据规则ID查询此规则的评估状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowEvaluationStateByAssignmentId(request *model.ShowEvaluationStateByAssignmentIdRequest) (*model.ShowEvaluationStateByAssignmentIdResponse, error) {
	requestDef := GenReqDefForShowEvaluationStateByAssignmentId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEvaluationStateByAssignmentIdResponse), nil
	}
}

// ShowEvaluationStateByAssignmentIdInvoker 获取规则的评估状态
func (c *ConfigClient) ShowEvaluationStateByAssignmentIdInvoker(request *model.ShowEvaluationStateByAssignmentIdRequest) *ShowEvaluationStateByAssignmentIdInvoker {
	requestDef := GenReqDefForShowEvaluationStateByAssignmentId()
	return &ShowEvaluationStateByAssignmentIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOrganizationPolicyAssignment 查询指定组织合规规则
//
// 查询指定组织合规规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowOrganizationPolicyAssignment(request *model.ShowOrganizationPolicyAssignmentRequest) (*model.ShowOrganizationPolicyAssignmentResponse, error) {
	requestDef := GenReqDefForShowOrganizationPolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOrganizationPolicyAssignmentResponse), nil
	}
}

// ShowOrganizationPolicyAssignmentInvoker 查询指定组织合规规则
func (c *ConfigClient) ShowOrganizationPolicyAssignmentInvoker(request *model.ShowOrganizationPolicyAssignmentRequest) *ShowOrganizationPolicyAssignmentInvoker {
	requestDef := GenReqDefForShowOrganizationPolicyAssignment()
	return &ShowOrganizationPolicyAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOrganizationPolicyAssignmentDetailedStatus 查询组织内每个成员帐号合规规则部署的详细状态
//
// 查询组织内每个成员帐号合规规则部署的详细状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowOrganizationPolicyAssignmentDetailedStatus(request *model.ShowOrganizationPolicyAssignmentDetailedStatusRequest) (*model.ShowOrganizationPolicyAssignmentDetailedStatusResponse, error) {
	requestDef := GenReqDefForShowOrganizationPolicyAssignmentDetailedStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOrganizationPolicyAssignmentDetailedStatusResponse), nil
	}
}

// ShowOrganizationPolicyAssignmentDetailedStatusInvoker 查询组织内每个成员帐号合规规则部署的详细状态
func (c *ConfigClient) ShowOrganizationPolicyAssignmentDetailedStatusInvoker(request *model.ShowOrganizationPolicyAssignmentDetailedStatusRequest) *ShowOrganizationPolicyAssignmentDetailedStatusInvoker {
	requestDef := GenReqDefForShowOrganizationPolicyAssignmentDetailedStatus()
	return &ShowOrganizationPolicyAssignmentDetailedStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOrganizationPolicyAssignmentStatuses 查询组织合规规则部署状态
//
// 查询组织合规规则部署状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowOrganizationPolicyAssignmentStatuses(request *model.ShowOrganizationPolicyAssignmentStatusesRequest) (*model.ShowOrganizationPolicyAssignmentStatusesResponse, error) {
	requestDef := GenReqDefForShowOrganizationPolicyAssignmentStatuses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOrganizationPolicyAssignmentStatusesResponse), nil
	}
}

// ShowOrganizationPolicyAssignmentStatusesInvoker 查询组织合规规则部署状态
func (c *ConfigClient) ShowOrganizationPolicyAssignmentStatusesInvoker(request *model.ShowOrganizationPolicyAssignmentStatusesRequest) *ShowOrganizationPolicyAssignmentStatusesInvoker {
	requestDef := GenReqDefForShowOrganizationPolicyAssignmentStatuses()
	return &ShowOrganizationPolicyAssignmentStatusesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPolicyAssignment 获取单个合规规则
//
// 根据规则ID获取单个规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowPolicyAssignment(request *model.ShowPolicyAssignmentRequest) (*model.ShowPolicyAssignmentResponse, error) {
	requestDef := GenReqDefForShowPolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPolicyAssignmentResponse), nil
	}
}

// ShowPolicyAssignmentInvoker 获取单个合规规则
func (c *ConfigClient) ShowPolicyAssignmentInvoker(request *model.ShowPolicyAssignmentRequest) *ShowPolicyAssignmentInvoker {
	requestDef := GenReqDefForShowPolicyAssignment()
	return &ShowPolicyAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRemediationConfiguration 查询修正配置
//
// 查询合规规则修正配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowRemediationConfiguration(request *model.ShowRemediationConfigurationRequest) (*model.ShowRemediationConfigurationResponse, error) {
	requestDef := GenReqDefForShowRemediationConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRemediationConfigurationResponse), nil
	}
}

// ShowRemediationConfigurationInvoker 查询修正配置
func (c *ConfigClient) ShowRemediationConfigurationInvoker(request *model.ShowRemediationConfigurationRequest) *ShowRemediationConfigurationInvoker {
	requestDef := GenReqDefForShowRemediationConfiguration()
	return &ShowRemediationConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateOrganizationPolicyAssignment 更新组织合规规则
//
// 更新组织合规规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) UpdateOrganizationPolicyAssignment(request *model.UpdateOrganizationPolicyAssignmentRequest) (*model.UpdateOrganizationPolicyAssignmentResponse, error) {
	requestDef := GenReqDefForUpdateOrganizationPolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateOrganizationPolicyAssignmentResponse), nil
	}
}

// UpdateOrganizationPolicyAssignmentInvoker 更新组织合规规则
func (c *ConfigClient) UpdateOrganizationPolicyAssignmentInvoker(request *model.UpdateOrganizationPolicyAssignmentRequest) *UpdateOrganizationPolicyAssignmentInvoker {
	requestDef := GenReqDefForUpdateOrganizationPolicyAssignment()
	return &UpdateOrganizationPolicyAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePolicyAssignment 更新合规规则
//
// 更新用户的合规规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) UpdatePolicyAssignment(request *model.UpdatePolicyAssignmentRequest) (*model.UpdatePolicyAssignmentResponse, error) {
	requestDef := GenReqDefForUpdatePolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePolicyAssignmentResponse), nil
	}
}

// UpdatePolicyAssignmentInvoker 更新合规规则
func (c *ConfigClient) UpdatePolicyAssignmentInvoker(request *model.UpdatePolicyAssignmentRequest) *UpdatePolicyAssignmentInvoker {
	requestDef := GenReqDefForUpdatePolicyAssignment()
	return &UpdatePolicyAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePolicyState 更新合规评估结果
//
// 更新用户自定义合规规则的合规评估结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) UpdatePolicyState(request *model.UpdatePolicyStateRequest) (*model.UpdatePolicyStateResponse, error) {
	requestDef := GenReqDefForUpdatePolicyState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePolicyStateResponse), nil
	}
}

// UpdatePolicyStateInvoker 更新合规评估结果
func (c *ConfigClient) UpdatePolicyStateInvoker(request *model.UpdatePolicyStateRequest) *UpdatePolicyStateInvoker {
	requestDef := GenReqDefForUpdatePolicyState()
	return &UpdatePolicyStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStoredQuery 创建高级查询
//
// 创建新的高级查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) CreateStoredQuery(request *model.CreateStoredQueryRequest) (*model.CreateStoredQueryResponse, error) {
	requestDef := GenReqDefForCreateStoredQuery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStoredQueryResponse), nil
	}
}

// CreateStoredQueryInvoker 创建高级查询
func (c *ConfigClient) CreateStoredQueryInvoker(request *model.CreateStoredQueryRequest) *CreateStoredQueryInvoker {
	requestDef := GenReqDefForCreateStoredQuery()
	return &CreateStoredQueryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteStoredQuery 删除高级查询
//
// 删除单个高级查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) DeleteStoredQuery(request *model.DeleteStoredQueryRequest) (*model.DeleteStoredQueryResponse, error) {
	requestDef := GenReqDefForDeleteStoredQuery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStoredQueryResponse), nil
	}
}

// DeleteStoredQueryInvoker 删除高级查询
func (c *ConfigClient) DeleteStoredQueryInvoker(request *model.DeleteStoredQueryRequest) *DeleteStoredQueryInvoker {
	requestDef := GenReqDefForDeleteStoredQuery()
	return &DeleteStoredQueryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSchemas 列举高级查询Schema
//
// # List Schemas
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListSchemas(request *model.ListSchemasRequest) (*model.ListSchemasResponse, error) {
	requestDef := GenReqDefForListSchemas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSchemasResponse), nil
	}
}

// ListSchemasInvoker 列举高级查询Schema
func (c *ConfigClient) ListSchemasInvoker(request *model.ListSchemasRequest) *ListSchemasInvoker {
	requestDef := GenReqDefForListSchemas()
	return &ListSchemasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStoredQueries 列出高级查询
//
// 列举所有高级查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListStoredQueries(request *model.ListStoredQueriesRequest) (*model.ListStoredQueriesResponse, error) {
	requestDef := GenReqDefForListStoredQueries()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStoredQueriesResponse), nil
	}
}

// ListStoredQueriesInvoker 列出高级查询
func (c *ConfigClient) ListStoredQueriesInvoker(request *model.ListStoredQueriesRequest) *ListStoredQueriesInvoker {
	requestDef := GenReqDefForListStoredQueries()
	return &ListStoredQueriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunQuery 运行高级查询
//
// 执行高级查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) RunQuery(request *model.RunQueryRequest) (*model.RunQueryResponse, error) {
	requestDef := GenReqDefForRunQuery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunQueryResponse), nil
	}
}

// RunQueryInvoker 运行高级查询
func (c *ConfigClient) RunQueryInvoker(request *model.RunQueryRequest) *RunQueryInvoker {
	requestDef := GenReqDefForRunQuery()
	return &RunQueryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStoredQuery 查询单个高级查询
//
// # Show Resource Query Language
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowStoredQuery(request *model.ShowStoredQueryRequest) (*model.ShowStoredQueryResponse, error) {
	requestDef := GenReqDefForShowStoredQuery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStoredQueryResponse), nil
	}
}

// ShowStoredQueryInvoker 查询单个高级查询
func (c *ConfigClient) ShowStoredQueryInvoker(request *model.ShowStoredQueryRequest) *ShowStoredQueryInvoker {
	requestDef := GenReqDefForShowStoredQuery()
	return &ShowStoredQueryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateStoredQuery 更新单个高级查询
//
// 更新自定义查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) UpdateStoredQuery(request *model.UpdateStoredQueryRequest) (*model.UpdateStoredQueryResponse, error) {
	requestDef := GenReqDefForUpdateStoredQuery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStoredQueryResponse), nil
	}
}

// UpdateStoredQueryInvoker 更新单个高级查询
func (c *ConfigClient) UpdateStoredQueryInvoker(request *model.UpdateStoredQueryRequest) *UpdateStoredQueryInvoker {
	requestDef := GenReqDefForUpdateStoredQuery()
	return &UpdateStoredQueryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRegions 查询用户可见的区域
//
// 查询用户可见的区域
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListRegions(request *model.ListRegionsRequest) (*model.ListRegionsResponse, error) {
	requestDef := GenReqDefForListRegions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRegionsResponse), nil
	}
}

// ListRegionsInvoker 查询用户可见的区域
func (c *ConfigClient) ListRegionsInvoker(request *model.ListRegionsRequest) *ListRegionsInvoker {
	requestDef := GenReqDefForListRegions()
	return &ListRegionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceRelations 列举资源关系
//
// 指定资源ID，查询该资源与其他资源的关联关系，可以指定关系方向为\&quot;in\&quot; 或者\&quot;out\&quot;。资源关系依赖开启资源记录器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowResourceRelations(request *model.ShowResourceRelationsRequest) (*model.ShowResourceRelationsResponse, error) {
	requestDef := GenReqDefForShowResourceRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceRelationsResponse), nil
	}
}

// ShowResourceRelationsInvoker 列举资源关系
func (c *ConfigClient) ShowResourceRelationsInvoker(request *model.ShowResourceRelationsRequest) *ShowResourceRelationsInvoker {
	requestDef := GenReqDefForShowResourceRelations()
	return &ShowResourceRelationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceRelationsDetail 列举资源关系详情
//
// 指定资源ID，查询该资源与其他资源的关联关系，可以指定关系方向为“in”或者“out”，需要当帐号有rms:resources:getRelation权限。资源关系依赖开启资源记录器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowResourceRelationsDetail(request *model.ShowResourceRelationsDetailRequest) (*model.ShowResourceRelationsDetailResponse, error) {
	requestDef := GenReqDefForShowResourceRelationsDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceRelationsDetailResponse), nil
	}
}

// ShowResourceRelationsDetailInvoker 列举资源关系详情
func (c *ConfigClient) ShowResourceRelationsDetailInvoker(request *model.ShowResourceRelationsDetailRequest) *ShowResourceRelationsDetailInvoker {
	requestDef := GenReqDefForShowResourceRelationsDetail()
	return &ShowResourceRelationsDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CollectAllResourcesSummary 列举资源概要
//
// 查询当前帐号的资源概览。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) CollectAllResourcesSummary(request *model.CollectAllResourcesSummaryRequest) (*model.CollectAllResourcesSummaryResponse, error) {
	requestDef := GenReqDefForCollectAllResourcesSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectAllResourcesSummaryResponse), nil
	}
}

// CollectAllResourcesSummaryInvoker 列举资源概要
func (c *ConfigClient) CollectAllResourcesSummaryInvoker(request *model.CollectAllResourcesSummaryRequest) *CollectAllResourcesSummaryInvoker {
	requestDef := GenReqDefForCollectAllResourcesSummary()
	return &CollectAllResourcesSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CollectTrackedResourcesSummary 列举资源记录器收集的资源概要
//
// 查询当前用户资源记录器收集的资源概览。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) CollectTrackedResourcesSummary(request *model.CollectTrackedResourcesSummaryRequest) (*model.CollectTrackedResourcesSummaryResponse, error) {
	requestDef := GenReqDefForCollectTrackedResourcesSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectTrackedResourcesSummaryResponse), nil
	}
}

// CollectTrackedResourcesSummaryInvoker 列举资源记录器收集的资源概要
func (c *ConfigClient) CollectTrackedResourcesSummaryInvoker(request *model.CollectTrackedResourcesSummaryRequest) *CollectTrackedResourcesSummaryInvoker {
	requestDef := GenReqDefForCollectTrackedResourcesSummary()
	return &CollectTrackedResourcesSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountAllResources 查询资源数量
//
// 查询当前帐号的资源数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) CountAllResources(request *model.CountAllResourcesRequest) (*model.CountAllResourcesResponse, error) {
	requestDef := GenReqDefForCountAllResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountAllResourcesResponse), nil
	}
}

// CountAllResourcesInvoker 查询资源数量
func (c *ConfigClient) CountAllResourcesInvoker(request *model.CountAllResourcesRequest) *CountAllResourcesInvoker {
	requestDef := GenReqDefForCountAllResources()
	return &CountAllResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountTrackedResources 查询资源记录器收集的资源数量
//
// 查询当前用户资源记录器收集的资源数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) CountTrackedResources(request *model.CountTrackedResourcesRequest) (*model.CountTrackedResourcesResponse, error) {
	requestDef := GenReqDefForCountTrackedResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountTrackedResourcesResponse), nil
	}
}

// CountTrackedResourcesInvoker 查询资源记录器收集的资源数量
func (c *ConfigClient) CountTrackedResourcesInvoker(request *model.CountTrackedResourcesRequest) *CountTrackedResourcesInvoker {
	requestDef := GenReqDefForCountTrackedResources()
	return &CountTrackedResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllResources 列举所有资源
//
// 返回当前用户下所有资源，需要当前用户有rms:resources:list权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListAllResources(request *model.ListAllResourcesRequest) (*model.ListAllResourcesResponse, error) {
	requestDef := GenReqDefForListAllResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllResourcesResponse), nil
	}
}

// ListAllResourcesInvoker 列举所有资源
func (c *ConfigClient) ListAllResourcesInvoker(request *model.ListAllResourcesRequest) *ListAllResourcesInvoker {
	requestDef := GenReqDefForListAllResources()
	return &ListAllResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllTags 列举资源标签
//
// 查询当前帐号下所有资源的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListAllTags(request *model.ListAllTagsRequest) (*model.ListAllTagsResponse, error) {
	requestDef := GenReqDefForListAllTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllTagsResponse), nil
	}
}

// ListAllTagsInvoker 列举资源标签
func (c *ConfigClient) ListAllTagsInvoker(request *model.ListAllTagsRequest) *ListAllTagsInvoker {
	requestDef := GenReqDefForListAllTags()
	return &ListAllTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProviders 列举云服务
//
// 查询Config支持的云服务、资源、区域列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListProviders(request *model.ListProvidersRequest) (*model.ListProvidersResponse, error) {
	requestDef := GenReqDefForListProviders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProvidersResponse), nil
	}
}

// ListProvidersInvoker 列举云服务
func (c *ConfigClient) ListProvidersInvoker(request *model.ListProvidersRequest) *ListProvidersInvoker {
	requestDef := GenReqDefForListProviders()
	return &ListProvidersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResources 列举指定类型的资源
//
// 返回当前租户下特定资源类型的资源，需要当前用户有rms:resources:list权限。比如查询云服务器，对应的Config资源类型是ecs.cloudservers，其中provider为ecs，type为cloudservers。 Config支持的服务和资源类型参见[支持的服务和区域](https://console.huaweicloud.com/eps/#/resources/supported)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListResources(request *model.ListResourcesRequest) (*model.ListResourcesResponse, error) {
	requestDef := GenReqDefForListResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourcesResponse), nil
	}
}

// ListResourcesInvoker 列举指定类型的资源
func (c *ConfigClient) ListResourcesInvoker(request *model.ListResourcesRequest) *ListResourcesInvoker {
	requestDef := GenReqDefForListResources()
	return &ListResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTrackedResourceTags 列举资源记录器收集的资源标签
//
// 查询当前用户资源记录器收集的资源的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListTrackedResourceTags(request *model.ListTrackedResourceTagsRequest) (*model.ListTrackedResourceTagsResponse, error) {
	requestDef := GenReqDefForListTrackedResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTrackedResourceTagsResponse), nil
	}
}

// ListTrackedResourceTagsInvoker 列举资源记录器收集的资源标签
func (c *ConfigClient) ListTrackedResourceTagsInvoker(request *model.ListTrackedResourceTagsRequest) *ListTrackedResourceTagsInvoker {
	requestDef := GenReqDefForListTrackedResourceTags()
	return &ListTrackedResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTrackedResources 列举资源记录器收集的全部资源
//
// 查询当前用户资源记录器收集的全部资源，需要当前用户有rms:resources:list权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListTrackedResources(request *model.ListTrackedResourcesRequest) (*model.ListTrackedResourcesResponse, error) {
	requestDef := GenReqDefForListTrackedResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTrackedResourcesResponse), nil
	}
}

// ListTrackedResourcesInvoker 列举资源记录器收集的全部资源
func (c *ConfigClient) ListTrackedResourcesInvoker(request *model.ListTrackedResourcesRequest) *ListTrackedResourcesInvoker {
	requestDef := GenReqDefForListTrackedResources()
	return &ListTrackedResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceById 查询单个资源
//
// 指定资源ID，返回该资源的详细信息，需要当前用户有rms:resources:get权限。比如查询云服务器，对应的Config资源类型是ecs.cloudservers，其中provider为ecs，type为cloudservers。Config支持的服务和资源类型参见[支持的服务和区域](https://console.huaweicloud.com/eps/#/resources/supported)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowResourceById(request *model.ShowResourceByIdRequest) (*model.ShowResourceByIdResponse, error) {
	requestDef := GenReqDefForShowResourceById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceByIdResponse), nil
	}
}

// ShowResourceByIdInvoker 查询单个资源
func (c *ConfigClient) ShowResourceByIdInvoker(request *model.ShowResourceByIdRequest) *ShowResourceByIdInvoker {
	requestDef := GenReqDefForShowResourceById()
	return &ShowResourceByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceDetail 查询帐号下的单个资源
//
// 查询当前帐号下的单个资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowResourceDetail(request *model.ShowResourceDetailRequest) (*model.ShowResourceDetailResponse, error) {
	requestDef := GenReqDefForShowResourceDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceDetailResponse), nil
	}
}

// ShowResourceDetailInvoker 查询帐号下的单个资源
func (c *ConfigClient) ShowResourceDetailInvoker(request *model.ShowResourceDetailRequest) *ShowResourceDetailInvoker {
	requestDef := GenReqDefForShowResourceDetail()
	return &ShowResourceDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTrackedResourceDetail 查询资源记录器收集的单个资源
//
// 查询当前用户资源记录器收集的单个资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowTrackedResourceDetail(request *model.ShowTrackedResourceDetailRequest) (*model.ShowTrackedResourceDetailResponse, error) {
	requestDef := GenReqDefForShowTrackedResourceDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTrackedResourceDetailResponse), nil
	}
}

// ShowTrackedResourceDetailInvoker 查询资源记录器收集的单个资源
func (c *ConfigClient) ShowTrackedResourceDetailInvoker(request *model.ShowTrackedResourceDetailRequest) *ShowTrackedResourceDetailInvoker {
	requestDef := GenReqDefForShowTrackedResourceDetail()
	return &ShowTrackedResourceDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountResourcesByTag 查询资源实例数量
//
// 使用标签过滤实例，标签管理服务需要提供按标签过滤各服务实例并汇总显示在列表中，需要各服务提供查询能力。注意：tags, tags_any, not_tags, not_tags_any等字段支持的tag的数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) CountResourcesByTag(request *model.CountResourcesByTagRequest) (*model.CountResourcesByTagResponse, error) {
	requestDef := GenReqDefForCountResourcesByTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountResourcesByTagResponse), nil
	}
}

// CountResourcesByTagInvoker 查询资源实例数量
func (c *ConfigClient) CountResourcesByTagInvoker(request *model.CountResourcesByTagRequest) *CountResourcesByTagInvoker {
	requestDef := GenReqDefForCountResourcesByTag()
	return &CountResourcesByTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourcesByTag 查询资源实例列表
//
// 使用标签过滤实例，标签管理服务需要提供按标签过滤各服务实例并汇总显示在列表中，需要各服务提供查询能力。注意：tags, tags_any, not_tags, not_tags_any等字段支持的tag的数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListResourcesByTag(request *model.ListResourcesByTagRequest) (*model.ListResourcesByTagResponse, error) {
	requestDef := GenReqDefForListResourcesByTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourcesByTagResponse), nil
	}
}

// ListResourcesByTagInvoker 查询资源实例列表
func (c *ConfigClient) ListResourcesByTagInvoker(request *model.ListResourcesByTagRequest) *ListResourcesByTagInvoker {
	requestDef := GenReqDefForListResourcesByTag()
	return &ListResourcesByTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTagsForResource 查询资源标签
//
// 查询指定实例的标签信息。标签管理服务需要使用该接口查询指定实例的全部标签数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListTagsForResource(request *model.ListTagsForResourceRequest) (*model.ListTagsForResourceResponse, error) {
	requestDef := GenReqDefForListTagsForResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsForResourceResponse), nil
	}
}

// ListTagsForResourceInvoker 查询资源标签
func (c *ConfigClient) ListTagsForResourceInvoker(request *model.ListTagsForResourceRequest) *ListTagsForResourceInvoker {
	requestDef := GenReqDefForListTagsForResource()
	return &ListTagsForResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTagsForResourceType 查询项目标签
//
// 查询租户在指定Project中实例类型的所有资源标签集合。标签管理服务需要能够列出当前租户全部已使用的资源标签集合，为各服务Console打资源标签和过滤实例时提供标签联想功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ListTagsForResourceType(request *model.ListTagsForResourceTypeRequest) (*model.ListTagsForResourceTypeResponse, error) {
	requestDef := GenReqDefForListTagsForResourceType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsForResourceTypeResponse), nil
	}
}

// ListTagsForResourceTypeInvoker 查询项目标签
func (c *ConfigClient) ListTagsForResourceTypeInvoker(request *model.ListTagsForResourceTypeRequest) *ListTagsForResourceTypeInvoker {
	requestDef := GenReqDefForListTagsForResourceType()
	return &ListTagsForResourceTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// TagResource 批量添加资源标签
//
// 此接口为幂等接口。为指定实例批量添加或删除标签，标签管理服务需要使用该接口批量管理实例的标签。一个资源上最多有20个标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) TagResource(request *model.TagResourceRequest) (*model.TagResourceResponse, error) {
	requestDef := GenReqDefForTagResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.TagResourceResponse), nil
	}
}

// TagResourceInvoker 批量添加资源标签
func (c *ConfigClient) TagResourceInvoker(request *model.TagResourceRequest) *TagResourceInvoker {
	requestDef := GenReqDefForTagResource()
	return &TagResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnTagResource 批量删除资源标签
//
// 此接口为幂等接口。为指定实例批量添加或删除标签，标签管理服务需要使用该接口批量管理实例的标签。一个资源上最多有20个标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) UnTagResource(request *model.UnTagResourceRequest) (*model.UnTagResourceResponse, error) {
	requestDef := GenReqDefForUnTagResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnTagResourceResponse), nil
	}
}

// UnTagResourceInvoker 批量删除资源标签
func (c *ConfigClient) UnTagResourceInvoker(request *model.UnTagResourceRequest) *UnTagResourceInvoker {
	requestDef := GenReqDefForUnTagResource()
	return &UnTagResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTrackerConfig 创建或更新记录器
//
// 创建或更新资源记录器，只能存在一个资源记录器
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) CreateTrackerConfig(request *model.CreateTrackerConfigRequest) (*model.CreateTrackerConfigResponse, error) {
	requestDef := GenReqDefForCreateTrackerConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTrackerConfigResponse), nil
	}
}

// CreateTrackerConfigInvoker 创建或更新记录器
func (c *ConfigClient) CreateTrackerConfigInvoker(request *model.CreateTrackerConfigRequest) *CreateTrackerConfigInvoker {
	requestDef := GenReqDefForCreateTrackerConfig()
	return &CreateTrackerConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTrackerConfig 删除记录器
//
// 删除资源记录器
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) DeleteTrackerConfig(request *model.DeleteTrackerConfigRequest) (*model.DeleteTrackerConfigResponse, error) {
	requestDef := GenReqDefForDeleteTrackerConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTrackerConfigResponse), nil
	}
}

// DeleteTrackerConfigInvoker 删除记录器
func (c *ConfigClient) DeleteTrackerConfigInvoker(request *model.DeleteTrackerConfigRequest) *DeleteTrackerConfigInvoker {
	requestDef := GenReqDefForDeleteTrackerConfig()
	return &DeleteTrackerConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTrackerConfig 查询记录器
//
// 查询资源记录器的详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ConfigClient) ShowTrackerConfig(request *model.ShowTrackerConfigRequest) (*model.ShowTrackerConfigResponse, error) {
	requestDef := GenReqDefForShowTrackerConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTrackerConfigResponse), nil
	}
}

// ShowTrackerConfigInvoker 查询记录器
func (c *ConfigClient) ShowTrackerConfigInvoker(request *model.ShowTrackerConfigRequest) *ShowTrackerConfigInvoker {
	requestDef := GenReqDefForShowTrackerConfig()
	return &ShowTrackerConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
