package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dataartsfabric/v1/model"
)

type DataArtsFabricClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDataArtsFabricClient(hcClient *httpclient.HcHttpClient) *DataArtsFabricClient {
	return &DataArtsFabricClient{HcClient: hcClient}
}

func DataArtsFabricClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateAgency 创建服务委托
//
// 为用户自动创建服务所需要的服务委托。委托需要附加必需的权限策略才能使用，创建委托会自动附加必需的权限策略，也可以指定附加需要的权限策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) CreateAgency(request *model.CreateAgencyRequest) (*model.CreateAgencyResponse, error) {
	requestDef := GenReqDefForCreateAgency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAgencyResponse), nil
	}
}

// CreateAgencyInvoker 创建服务委托
func (c *DataArtsFabricClient) CreateAgencyInvoker(request *model.CreateAgencyRequest) *CreateAgencyInvoker {
	requestDef := GenReqDefForCreateAgency()
	return &CreateAgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAgency 删除服务委托
//
// 删除用户的服务委托权限。可以通过指定权限策略来删除委托中附加的权限策略，必需的权限策略无法被删除；如果不指定权限策略，会删除整个委托。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) DeleteAgency(request *model.DeleteAgencyRequest) (*model.DeleteAgencyResponse, error) {
	requestDef := GenReqDefForDeleteAgency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAgencyResponse), nil
	}
}

// DeleteAgencyInvoker 删除服务委托
func (c *DataArtsFabricClient) DeleteAgencyInvoker(request *model.DeleteAgencyRequest) *DeleteAgencyInvoker {
	requestDef := GenReqDefForDeleteAgency()
	return &DeleteAgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAgency 查询服务委托
//
// 查询用用户服务委托详情是否满足系统所需权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) ListAgency(request *model.ListAgencyRequest) (*model.ListAgencyResponse, error) {
	requestDef := GenReqDefForListAgency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAgencyResponse), nil
	}
}

// ListAgencyInvoker 查询服务委托
func (c *DataArtsFabricClient) ListAgencyInvoker(request *model.ListAgencyRequest) *ListAgencyInvoker {
	requestDef := GenReqDefForListAgency()
	return &ListAgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAgreement 注册租户协议
//
// 注册租户协议
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) CreateAgreement(request *model.CreateAgreementRequest) (*model.CreateAgreementResponse, error) {
	requestDef := GenReqDefForCreateAgreement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAgreementResponse), nil
	}
}

// CreateAgreementInvoker 注册租户协议
func (c *DataArtsFabricClient) CreateAgreementInvoker(request *model.CreateAgreementRequest) *CreateAgreementInvoker {
	requestDef := GenReqDefForCreateAgreement()
	return &CreateAgreementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAgreement 删除用户注册协议
//
// 删除用户注册协议
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) DeleteAgreement(request *model.DeleteAgreementRequest) (*model.DeleteAgreementResponse, error) {
	requestDef := GenReqDefForDeleteAgreement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAgreementResponse), nil
	}
}

// DeleteAgreementInvoker 删除用户注册协议
func (c *DataArtsFabricClient) DeleteAgreementInvoker(request *model.DeleteAgreementRequest) *DeleteAgreementInvoker {
	requestDef := GenReqDefForDeleteAgreement()
	return &DeleteAgreementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAgreement 查询用户是否注册协议
//
// 查询用户是否注册协议
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) ShowAgreement(request *model.ShowAgreementRequest) (*model.ShowAgreementResponse, error) {
	requestDef := GenReqDefForShowAgreement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAgreementResponse), nil
	}
}

// ShowAgreementInvoker 查询用户是否注册协议
func (c *DataArtsFabricClient) ShowAgreementInvoker(request *model.ShowAgreementRequest) *ShowAgreementInvoker {
	requestDef := GenReqDefForShowAgreement()
	return &ShowAgreementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAgreementRule 查询系统协议
//
// 查询系统协议
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) ShowAgreementRule(request *model.ShowAgreementRuleRequest) (*model.ShowAgreementRuleResponse, error) {
	requestDef := GenReqDefForShowAgreementRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAgreementRuleResponse), nil
	}
}

// ShowAgreementRuleInvoker 查询系统协议
func (c *DataArtsFabricClient) ShowAgreementRuleInvoker(request *model.ShowAgreementRuleRequest) *ShowAgreementRuleInvoker {
	requestDef := GenReqDefForShowAgreementRule()
	return &ShowAgreementRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFeatures 查询用户支持特性
//
// 查询用户支持特性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) ListFeatures(request *model.ListFeaturesRequest) (*model.ListFeaturesResponse, error) {
	requestDef := GenReqDefForListFeatures()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFeaturesResponse), nil
	}
}

// ListFeaturesInvoker 查询用户支持特性
func (c *DataArtsFabricClient) ListFeaturesInvoker(request *model.ListFeaturesRequest) *ListFeaturesInvoker {
	requestDef := GenReqDefForListFeatures()
	return &ListFeaturesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEndpoint 创建Endpoint
//
// 创建Endpoint
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) CreateEndpoint(request *model.CreateEndpointRequest) (*model.CreateEndpointResponse, error) {
	requestDef := GenReqDefForCreateEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEndpointResponse), nil
	}
}

// CreateEndpointInvoker 创建Endpoint
func (c *DataArtsFabricClient) CreateEndpointInvoker(request *model.CreateEndpointRequest) *CreateEndpointInvoker {
	requestDef := GenReqDefForCreateEndpoint()
	return &CreateEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEndpoint 删除endpioint
//
// 删除endpioint
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) DeleteEndpoint(request *model.DeleteEndpointRequest) (*model.DeleteEndpointResponse, error) {
	requestDef := GenReqDefForDeleteEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEndpointResponse), nil
	}
}

// DeleteEndpointInvoker 删除endpioint
func (c *DataArtsFabricClient) DeleteEndpointInvoker(request *model.DeleteEndpointRequest) *DeleteEndpointInvoker {
	requestDef := GenReqDefForDeleteEndpoint()
	return &DeleteEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEndpoints 查询Endpoint列表
//
// 查询Endpoint列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) ListEndpoints(request *model.ListEndpointsRequest) (*model.ListEndpointsResponse, error) {
	requestDef := GenReqDefForListEndpoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointsResponse), nil
	}
}

// ListEndpointsInvoker 查询Endpoint列表
func (c *DataArtsFabricClient) ListEndpointsInvoker(request *model.ListEndpointsRequest) *ListEndpointsInvoker {
	requestDef := GenReqDefForListEndpoints()
	return &ListEndpointsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEndpoint 查询endpioint详情
//
// 查询endpioint详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) ShowEndpoint(request *model.ShowEndpointRequest) (*model.ShowEndpointResponse, error) {
	requestDef := GenReqDefForShowEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEndpointResponse), nil
	}
}

// ShowEndpointInvoker 查询endpioint详情
func (c *DataArtsFabricClient) ShowEndpointInvoker(request *model.ShowEndpointRequest) *ShowEndpointInvoker {
	requestDef := GenReqDefForShowEndpoint()
	return &ShowEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SubscribeEndpoint 订阅Endpoint
//
// 在用户Workspace下订阅Endpoint（公共Endpoint场景）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) SubscribeEndpoint(request *model.SubscribeEndpointRequest) (*model.SubscribeEndpointResponse, error) {
	requestDef := GenReqDefForSubscribeEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SubscribeEndpointResponse), nil
	}
}

// SubscribeEndpointInvoker 订阅Endpoint
func (c *DataArtsFabricClient) SubscribeEndpointInvoker(request *model.SubscribeEndpointRequest) *SubscribeEndpointInvoker {
	requestDef := GenReqDefForSubscribeEndpoint()
	return &SubscribeEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEndpoint 修改Endpoint
//
// 修改Endpoint
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) UpdateEndpoint(request *model.UpdateEndpointRequest) (*model.UpdateEndpointResponse, error) {
	requestDef := GenReqDefForUpdateEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointResponse), nil
	}
}

// UpdateEndpointInvoker 修改Endpoint
func (c *DataArtsFabricClient) UpdateEndpointInvoker(request *model.UpdateEndpointRequest) *UpdateEndpointInvoker {
	requestDef := GenReqDefForUpdateEndpoint()
	return &UpdateEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAdminHealthCheck 健康检查
//
// 服务健康检查
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) ShowAdminHealthCheck(request *model.ShowAdminHealthCheckRequest) (*model.ShowAdminHealthCheckResponse, error) {
	requestDef := GenReqDefForShowAdminHealthCheck()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAdminHealthCheckResponse), nil
	}
}

// ShowAdminHealthCheckInvoker 健康检查
func (c *DataArtsFabricClient) ShowAdminHealthCheckInvoker(request *model.ShowAdminHealthCheckRequest) *ShowAdminHealthCheckInvoker {
	requestDef := GenReqDefForShowAdminHealthCheck()
	return &ShowAdminHealthCheckInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMessageNotificationPolicy 创建消息通知策略
//
// 创建消息通知策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) CreateMessageNotificationPolicy(request *model.CreateMessageNotificationPolicyRequest) (*model.CreateMessageNotificationPolicyResponse, error) {
	requestDef := GenReqDefForCreateMessageNotificationPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMessageNotificationPolicyResponse), nil
	}
}

// CreateMessageNotificationPolicyInvoker 创建消息通知策略
func (c *DataArtsFabricClient) CreateMessageNotificationPolicyInvoker(request *model.CreateMessageNotificationPolicyRequest) *CreateMessageNotificationPolicyInvoker {
	requestDef := GenReqDefForCreateMessageNotificationPolicy()
	return &CreateMessageNotificationPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMessageNotificationPolicy 删除消息通知策略
//
// 删除消息通知策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) DeleteMessageNotificationPolicy(request *model.DeleteMessageNotificationPolicyRequest) (*model.DeleteMessageNotificationPolicyResponse, error) {
	requestDef := GenReqDefForDeleteMessageNotificationPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMessageNotificationPolicyResponse), nil
	}
}

// DeleteMessageNotificationPolicyInvoker 删除消息通知策略
func (c *DataArtsFabricClient) DeleteMessageNotificationPolicyInvoker(request *model.DeleteMessageNotificationPolicyRequest) *DeleteMessageNotificationPolicyInvoker {
	requestDef := GenReqDefForDeleteMessageNotificationPolicy()
	return &DeleteMessageNotificationPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMessageNotificationPolicy 查询消息通知策略列表
//
// 查询消息通知策略列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) ListMessageNotificationPolicy(request *model.ListMessageNotificationPolicyRequest) (*model.ListMessageNotificationPolicyResponse, error) {
	requestDef := GenReqDefForListMessageNotificationPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMessageNotificationPolicyResponse), nil
	}
}

// ListMessageNotificationPolicyInvoker 查询消息通知策略列表
func (c *DataArtsFabricClient) ListMessageNotificationPolicyInvoker(request *model.ListMessageNotificationPolicyRequest) *ListMessageNotificationPolicyInvoker {
	requestDef := GenReqDefForListMessageNotificationPolicy()
	return &ListMessageNotificationPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMetricsConfig 更新AOM监控采集配置
//
// 更新AOM监控采集配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) UpdateMetricsConfig(request *model.UpdateMetricsConfigRequest) (*model.UpdateMetricsConfigResponse, error) {
	requestDef := GenReqDefForUpdateMetricsConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMetricsConfigResponse), nil
	}
}

// UpdateMetricsConfigInvoker 更新AOM监控采集配置
func (c *DataArtsFabricClient) UpdateMetricsConfigInvoker(request *model.UpdateMetricsConfigRequest) *UpdateMetricsConfigInvoker {
	requestDef := GenReqDefForUpdateMetricsConfig()
	return &UpdateMetricsConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CleanupModel 删除未使用的模型定义
//
// 清理未使用的模型定义
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) CleanupModel(request *model.CleanupModelRequest) (*model.CleanupModelResponse, error) {
	requestDef := GenReqDefForCleanupModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CleanupModelResponse), nil
	}
}

// CleanupModelInvoker 删除未使用的模型定义
func (c *DataArtsFabricClient) CleanupModelInvoker(request *model.CleanupModelRequest) *CleanupModelInvoker {
	requestDef := GenReqDefForCleanupModel()
	return &CleanupModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateModelDefinition 创建模型
//
// 创建模型
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) CreateModelDefinition(request *model.CreateModelDefinitionRequest) (*model.CreateModelDefinitionResponse, error) {
	requestDef := GenReqDefForCreateModelDefinition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateModelDefinitionResponse), nil
	}
}

// CreateModelDefinitionInvoker 创建模型
func (c *DataArtsFabricClient) CreateModelDefinitionInvoker(request *model.CreateModelDefinitionRequest) *CreateModelDefinitionInvoker {
	requestDef := GenReqDefForCreateModelDefinition()
	return &CreateModelDefinitionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteModelVersion 删除模型版本
//
// 删除模型版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) DeleteModelVersion(request *model.DeleteModelVersionRequest) (*model.DeleteModelVersionResponse, error) {
	requestDef := GenReqDefForDeleteModelVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteModelVersionResponse), nil
	}
}

// DeleteModelVersionInvoker 删除模型版本
func (c *DataArtsFabricClient) DeleteModelVersionInvoker(request *model.DeleteModelVersionRequest) *DeleteModelVersionInvoker {
	requestDef := GenReqDefForDeleteModelVersion()
	return &DeleteModelVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBaseModels 列举基模型
//
// 列举基模型
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) ListBaseModels(request *model.ListBaseModelsRequest) (*model.ListBaseModelsResponse, error) {
	requestDef := GenReqDefForListBaseModels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBaseModelsResponse), nil
	}
}

// ListBaseModelsInvoker 列举基模型
func (c *DataArtsFabricClient) ListBaseModelsInvoker(request *model.ListBaseModelsRequest) *ListBaseModelsInvoker {
	requestDef := GenReqDefForListBaseModels()
	return &ListBaseModelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListModelVersions 查询模型的版本列表
//
// 查询模型的版本列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) ListModelVersions(request *model.ListModelVersionsRequest) (*model.ListModelVersionsResponse, error) {
	requestDef := GenReqDefForListModelVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListModelVersionsResponse), nil
	}
}

// ListModelVersionsInvoker 查询模型的版本列表
func (c *DataArtsFabricClient) ListModelVersionsInvoker(request *model.ListModelVersionsRequest) *ListModelVersionsInvoker {
	requestDef := GenReqDefForListModelVersions()
	return &ListModelVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListModels 列举模型
//
// 列举模型
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) ListModels(request *model.ListModelsRequest) (*model.ListModelsResponse, error) {
	requestDef := GenReqDefForListModels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListModelsResponse), nil
	}
}

// ListModelsInvoker 列举模型
func (c *DataArtsFabricClient) ListModelsInvoker(request *model.ListModelsRequest) *ListModelsInvoker {
	requestDef := GenReqDefForListModels()
	return &ListModelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateModelDefinition 更新模型
//
// 更新模型，会生成新版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) UpdateModelDefinition(request *model.UpdateModelDefinitionRequest) (*model.UpdateModelDefinitionResponse, error) {
	requestDef := GenReqDefForUpdateModelDefinition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateModelDefinitionResponse), nil
	}
}

// UpdateModelDefinitionInvoker 更新模型
func (c *DataArtsFabricClient) UpdateModelDefinitionInvoker(request *model.UpdateModelDefinitionRequest) *UpdateModelDefinitionInvoker {
	requestDef := GenReqDefForUpdateModelDefinition()
	return &UpdateModelDefinitionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSpecs 查询服务规格列表
//
// 查询服务规格列表，购买计算资源使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) ListSpecs(request *model.ListSpecsRequest) (*model.ListSpecsResponse, error) {
	requestDef := GenReqDefForListSpecs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSpecsResponse), nil
	}
}

// ListSpecsInvoker 查询服务规格列表
func (c *DataArtsFabricClient) ListSpecsInvoker(request *model.ListSpecsRequest) *ListSpecsInvoker {
	requestDef := GenReqDefForListSpecs()
	return &ListSpecsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateFabricWorkspaceTags 批量打资源标签
//
// 批量打资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) BatchCreateFabricWorkspaceTags(request *model.BatchCreateFabricWorkspaceTagsRequest) (*model.BatchCreateFabricWorkspaceTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateFabricWorkspaceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateFabricWorkspaceTagsResponse), nil
	}
}

// BatchCreateFabricWorkspaceTagsInvoker 批量打资源标签
func (c *DataArtsFabricClient) BatchCreateFabricWorkspaceTagsInvoker(request *model.BatchCreateFabricWorkspaceTagsRequest) *BatchCreateFabricWorkspaceTagsInvoker {
	requestDef := GenReqDefForBatchCreateFabricWorkspaceTags()
	return &BatchCreateFabricWorkspaceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteFabricWorkspaceTags 批量删除资源标签
//
// 批量删除资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) BatchDeleteFabricWorkspaceTags(request *model.BatchDeleteFabricWorkspaceTagsRequest) (*model.BatchDeleteFabricWorkspaceTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteFabricWorkspaceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteFabricWorkspaceTagsResponse), nil
	}
}

// BatchDeleteFabricWorkspaceTagsInvoker 批量删除资源标签
func (c *DataArtsFabricClient) BatchDeleteFabricWorkspaceTagsInvoker(request *model.BatchDeleteFabricWorkspaceTagsRequest) *BatchDeleteFabricWorkspaceTagsInvoker {
	requestDef := GenReqDefForBatchDeleteFabricWorkspaceTags()
	return &BatchDeleteFabricWorkspaceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountTagFabricWorkspaces 查询资源实例数量
//
// 查询资源实例数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) CountTagFabricWorkspaces(request *model.CountTagFabricWorkspacesRequest) (*model.CountTagFabricWorkspacesResponse, error) {
	requestDef := GenReqDefForCountTagFabricWorkspaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountTagFabricWorkspacesResponse), nil
	}
}

// CountTagFabricWorkspacesInvoker 查询资源实例数量
func (c *DataArtsFabricClient) CountTagFabricWorkspacesInvoker(request *model.CountTagFabricWorkspacesRequest) *CountTagFabricWorkspacesInvoker {
	requestDef := GenReqDefForCountTagFabricWorkspaces()
	return &CountTagFabricWorkspacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFabricProjectTags 查询项目标签
//
// 查询项目标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) ListFabricProjectTags(request *model.ListFabricProjectTagsRequest) (*model.ListFabricProjectTagsResponse, error) {
	requestDef := GenReqDefForListFabricProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFabricProjectTagsResponse), nil
	}
}

// ListFabricProjectTagsInvoker 查询项目标签
func (c *DataArtsFabricClient) ListFabricProjectTagsInvoker(request *model.ListFabricProjectTagsRequest) *ListFabricProjectTagsInvoker {
	requestDef := GenReqDefForListFabricProjectTags()
	return &ListFabricProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFabricWorkspaceTags 查询资源标签
//
// 查询资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) ListFabricWorkspaceTags(request *model.ListFabricWorkspaceTagsRequest) (*model.ListFabricWorkspaceTagsResponse, error) {
	requestDef := GenReqDefForListFabricWorkspaceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFabricWorkspaceTagsResponse), nil
	}
}

// ListFabricWorkspaceTagsInvoker 查询资源标签
func (c *DataArtsFabricClient) ListFabricWorkspaceTagsInvoker(request *model.ListFabricWorkspaceTagsRequest) *ListFabricWorkspaceTagsInvoker {
	requestDef := GenReqDefForListFabricWorkspaceTags()
	return &ListFabricWorkspaceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTagFabricWorkspaces 查询资源实例列表
//
// 查询资源实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) ListTagFabricWorkspaces(request *model.ListTagFabricWorkspacesRequest) (*model.ListTagFabricWorkspacesResponse, error) {
	requestDef := GenReqDefForListTagFabricWorkspaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagFabricWorkspacesResponse), nil
	}
}

// ListTagFabricWorkspacesInvoker 查询资源实例列表
func (c *DataArtsFabricClient) ListTagFabricWorkspacesInvoker(request *model.ListTagFabricWorkspacesRequest) *ListTagFabricWorkspacesInvoker {
	requestDef := GenReqDefForListTagFabricWorkspaces()
	return &ListTagFabricWorkspacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkspace 创建Workspace
//
// # Create workspace
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) CreateWorkspace(request *model.CreateWorkspaceRequest) (*model.CreateWorkspaceResponse, error) {
	requestDef := GenReqDefForCreateWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkspaceResponse), nil
	}
}

// CreateWorkspaceInvoker 创建Workspace
func (c *DataArtsFabricClient) CreateWorkspaceInvoker(request *model.CreateWorkspaceRequest) *CreateWorkspaceInvoker {
	requestDef := GenReqDefForCreateWorkspace()
	return &CreateWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWorkspace 删除Workspace
//
// 删除Workspace
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) DeleteWorkspace(request *model.DeleteWorkspaceRequest) (*model.DeleteWorkspaceResponse, error) {
	requestDef := GenReqDefForDeleteWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWorkspaceResponse), nil
	}
}

// DeleteWorkspaceInvoker 删除Workspace
func (c *DataArtsFabricClient) DeleteWorkspaceInvoker(request *model.DeleteWorkspaceRequest) *DeleteWorkspaceInvoker {
	requestDef := GenReqDefForDeleteWorkspace()
	return &DeleteWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkspaces 查询Workspace列表
//
// 查询Workspace列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) ListWorkspaces(request *model.ListWorkspacesRequest) (*model.ListWorkspacesResponse, error) {
	requestDef := GenReqDefForListWorkspaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkspacesResponse), nil
	}
}

// ListWorkspacesInvoker 查询Workspace列表
func (c *DataArtsFabricClient) ListWorkspacesInvoker(request *model.ListWorkspacesRequest) *ListWorkspacesInvoker {
	requestDef := GenReqDefForListWorkspaces()
	return &ListWorkspacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkspace 更新Workspace
//
// 更新Workspace
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricClient) UpdateWorkspace(request *model.UpdateWorkspaceRequest) (*model.UpdateWorkspaceResponse, error) {
	requestDef := GenReqDefForUpdateWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkspaceResponse), nil
	}
}

// UpdateWorkspaceInvoker 更新Workspace
func (c *DataArtsFabricClient) UpdateWorkspaceInvoker(request *model.UpdateWorkspaceRequest) *UpdateWorkspaceInvoker {
	requestDef := GenReqDefForUpdateWorkspace()
	return &UpdateWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
