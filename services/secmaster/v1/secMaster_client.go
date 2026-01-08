package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/secmaster/v1/model"
)

type SecMasterClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewSecMasterClient(hcClient *httpclient.HcHttpClient) *SecMasterClient {
	return &SecMasterClient{HcClient: hcClient}
}

func SecMasterClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BatchCreateDataobjectRelations 批量关联数据对象
//
// 批量关联数据对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) BatchCreateDataobjectRelations(request *model.BatchCreateDataobjectRelationsRequest) (*model.BatchCreateDataobjectRelationsResponse, error) {
	requestDef := GenReqDefForBatchCreateDataobjectRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateDataobjectRelationsResponse), nil
	}
}

// BatchCreateDataobjectRelationsInvoker 批量关联数据对象
func (c *SecMasterClient) BatchCreateDataobjectRelationsInvoker(request *model.BatchCreateDataobjectRelationsRequest) *BatchCreateDataobjectRelationsInvoker {
	requestDef := GenReqDefForBatchCreateDataobjectRelations()
	return &BatchCreateDataobjectRelationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateDatapanelObjects 批量创建数据对象
//
// 数据面批量创建数据类纳管的数据对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) BatchCreateDatapanelObjects(request *model.BatchCreateDatapanelObjectsRequest) (*model.BatchCreateDatapanelObjectsResponse, error) {
	requestDef := GenReqDefForBatchCreateDatapanelObjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateDatapanelObjectsResponse), nil
	}
}

// BatchCreateDatapanelObjectsInvoker 批量创建数据对象
func (c *SecMasterClient) BatchCreateDatapanelObjectsInvoker(request *model.BatchCreateDatapanelObjectsRequest) *BatchCreateDatapanelObjectsInvoker {
	requestDef := GenReqDefForBatchCreateDatapanelObjects()
	return &BatchCreateDatapanelObjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSearchMetricHits 批量查询指标结果
//
// 批量查询指标结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) BatchSearchMetricHits(request *model.BatchSearchMetricHitsRequest) (*model.BatchSearchMetricHitsResponse, error) {
	requestDef := GenReqDefForBatchSearchMetricHits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSearchMetricHitsResponse), nil
	}
}

// BatchSearchMetricHitsInvoker 批量查询指标结果
func (c *SecMasterClient) BatchSearchMetricHitsInvoker(request *model.BatchSearchMetricHitsRequest) *BatchSearchMetricHitsInvoker {
	requestDef := GenReqDefForBatchSearchMetricHits()
	return &BatchSearchMetricHitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchTagResources 批量添加资源标签
//
// 为指定实例批量添加标签
// 标签管理服务需要使用该接口批量管理实例的标签。
// 一个资源上最多有20个标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) BatchTagResources(request *model.BatchTagResourcesRequest) (*model.BatchTagResourcesResponse, error) {
	requestDef := GenReqDefForBatchTagResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchTagResourcesResponse), nil
	}
}

// BatchTagResourcesInvoker 批量添加资源标签
func (c *SecMasterClient) BatchTagResourcesInvoker(request *model.BatchTagResourcesRequest) *BatchTagResourcesInvoker {
	requestDef := GenReqDefForBatchTagResources()
	return &BatchTagResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUntagResources 批量删除资源标签
//
// 为指定实例批量删除标签
// 标签管理服务需要使用该接口批量管理实例的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) BatchUntagResources(request *model.BatchUntagResourcesRequest) (*model.BatchUntagResourcesResponse, error) {
	requestDef := GenReqDefForBatchUntagResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUntagResourcesResponse), nil
	}
}

// BatchUntagResourcesInvoker 批量删除资源标签
func (c *SecMasterClient) BatchUntagResourcesInvoker(request *model.BatchUntagResourcesRequest) *BatchUntagResourcesInvoker {
	requestDef := GenReqDefForBatchUntagResources()
	return &BatchUntagResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateCatalogue 批量修改目录
//
// 批量修改自定义目录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) BatchUpdateCatalogue(request *model.BatchUpdateCatalogueRequest) (*model.BatchUpdateCatalogueResponse, error) {
	requestDef := GenReqDefForBatchUpdateCatalogue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateCatalogueResponse), nil
	}
}

// BatchUpdateCatalogueInvoker 批量修改目录
func (c *SecMasterClient) BatchUpdateCatalogueInvoker(request *model.BatchUpdateCatalogueRequest) *BatchUpdateCatalogueInvoker {
	requestDef := GenReqDefForBatchUpdateCatalogue()
	return &BatchUpdateCatalogueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeAlert 更新告警
//
// 编辑告警，根据实际修改的属性更新，未修改的列不更新
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ChangeAlert(request *model.ChangeAlertRequest) (*model.ChangeAlertResponse, error) {
	requestDef := GenReqDefForChangeAlert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeAlertResponse), nil
	}
}

// ChangeAlertInvoker 更新告警
func (c *SecMasterClient) ChangeAlertInvoker(request *model.ChangeAlertRequest) *ChangeAlertInvoker {
	requestDef := GenReqDefForChangeAlert()
	return &ChangeAlertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeAlerts 批量更新告警
//
// 批量更新告警，根据实际修改的属性更新，未修改的列不更新
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ChangeAlerts(request *model.ChangeAlertsRequest) (*model.ChangeAlertsResponse, error) {
	requestDef := GenReqDefForChangeAlerts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeAlertsResponse), nil
	}
}

// ChangeAlertsInvoker 批量更新告警
func (c *SecMasterClient) ChangeAlertsInvoker(request *model.ChangeAlertsRequest) *ChangeAlertsInvoker {
	requestDef := GenReqDefForChangeAlerts()
	return &ChangeAlertsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeIncident 更新事件
//
// 编辑事件，根据实际修改的属性更新，未修改的列不更新
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ChangeIncident(request *model.ChangeIncidentRequest) (*model.ChangeIncidentResponse, error) {
	requestDef := GenReqDefForChangeIncident()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeIncidentResponse), nil
	}
}

// ChangeIncidentInvoker 更新事件
func (c *SecMasterClient) ChangeIncidentInvoker(request *model.ChangeIncidentRequest) *ChangeIncidentInvoker {
	requestDef := GenReqDefForChangeIncident()
	return &ChangeIncidentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeIncidents 批量更新事件
//
// 更新事件，根据实际修改的属性更新，未修改的列不更新
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ChangeIncidents(request *model.ChangeIncidentsRequest) (*model.ChangeIncidentsResponse, error) {
	requestDef := GenReqDefForChangeIncidents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeIncidentsResponse), nil
	}
}

// ChangeIncidentsInvoker 批量更新事件
func (c *SecMasterClient) ChangeIncidentsInvoker(request *model.ChangeIncidentsRequest) *ChangeIncidentsInvoker {
	requestDef := GenReqDefForChangeIncidents()
	return &ChangeIncidentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangePlaybookInstance 操作剧本实例
//
// 操作剧本实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ChangePlaybookInstance(request *model.ChangePlaybookInstanceRequest) (*model.ChangePlaybookInstanceResponse, error) {
	requestDef := GenReqDefForChangePlaybookInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangePlaybookInstanceResponse), nil
	}
}

// ChangePlaybookInstanceInvoker 操作剧本实例
func (c *SecMasterClient) ChangePlaybookInstanceInvoker(request *model.ChangePlaybookInstanceRequest) *ChangePlaybookInstanceInvoker {
	requestDef := GenReqDefForChangePlaybookInstance()
	return &ChangePlaybookInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeResource 更新资产信息
//
// 编辑资产，根据实际修改的属性更新，未修改的列不更新
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ChangeResource(request *model.ChangeResourceRequest) (*model.ChangeResourceResponse, error) {
	requestDef := GenReqDefForChangeResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeResourceResponse), nil
	}
}

// ChangeResourceInvoker 更新资产信息
func (c *SecMasterClient) ChangeResourceInvoker(request *model.ChangeResourceRequest) *ChangeResourceInvoker {
	requestDef := GenReqDefForChangeResource()
	return &ChangeResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyMapping 复制映射
//
// 复制映射
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CopyMapping(request *model.CopyMappingRequest) (*model.CopyMappingResponse, error) {
	requestDef := GenReqDefForCopyMapping()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyMappingResponse), nil
	}
}

// CopyMappingInvoker 复制映射
func (c *SecMasterClient) CopyMappingInvoker(request *model.CopyMappingRequest) *CopyMappingInvoker {
	requestDef := GenReqDefForCopyMapping()
	return &CopyMappingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyPlaybookVersion 克隆剧本及版本
//
// 克隆剧本及版本（待下线）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CopyPlaybookVersion(request *model.CopyPlaybookVersionRequest) (*model.CopyPlaybookVersionResponse, error) {
	requestDef := GenReqDefForCopyPlaybookVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyPlaybookVersionResponse), nil
	}
}

// CopyPlaybookVersionInvoker 克隆剧本及版本
func (c *SecMasterClient) CopyPlaybookVersionInvoker(request *model.CopyPlaybookVersionRequest) *CopyPlaybookVersionInvoker {
	requestDef := GenReqDefForCopyPlaybookVersion()
	return &CopyPlaybookVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountResourceInstance 查询资源实例数量
//
// 使用标签过滤实例，查询资源实例数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CountResourceInstance(request *model.CountResourceInstanceRequest) (*model.CountResourceInstanceResponse, error) {
	requestDef := GenReqDefForCountResourceInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountResourceInstanceResponse), nil
	}
}

// CountResourceInstanceInvoker 查询资源实例数量
func (c *SecMasterClient) CountResourceInstanceInvoker(request *model.CountResourceInstanceRequest) *CountResourceInstanceInvoker {
	requestDef := GenReqDefForCountResourceInstance()
	return &CountResourceInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAlert 创建告警
//
// 创建告警
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateAlert(request *model.CreateAlertRequest) (*model.CreateAlertResponse, error) {
	requestDef := GenReqDefForCreateAlert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlertResponse), nil
	}
}

// CreateAlertInvoker 创建告警
func (c *SecMasterClient) CreateAlertInvoker(request *model.CreateAlertRequest) *CreateAlertInvoker {
	requestDef := GenReqDefForCreateAlert()
	return &CreateAlertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAlertRule 创建告警规则
//
// 创建告警规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateAlertRule(request *model.CreateAlertRuleRequest) (*model.CreateAlertRuleResponse, error) {
	requestDef := GenReqDefForCreateAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlertRuleResponse), nil
	}
}

// CreateAlertRuleInvoker 创建告警规则
func (c *SecMasterClient) CreateAlertRuleInvoker(request *model.CreateAlertRuleRequest) *CreateAlertRuleInvoker {
	requestDef := GenReqDefForCreateAlertRule()
	return &CreateAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAlertRuleSimulation 模拟告警规则
//
// 模拟告警规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateAlertRuleSimulation(request *model.CreateAlertRuleSimulationRequest) (*model.CreateAlertRuleSimulationResponse, error) {
	requestDef := GenReqDefForCreateAlertRuleSimulation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlertRuleSimulationResponse), nil
	}
}

// CreateAlertRuleSimulationInvoker 模拟告警规则
func (c *SecMasterClient) CreateAlertRuleSimulationInvoker(request *model.CreateAlertRuleSimulationRequest) *CreateAlertRuleSimulationInvoker {
	requestDef := GenReqDefForCreateAlertRuleSimulation()
	return &CreateAlertRuleSimulationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAopWorkflow 创建流程
//
// 创建流程
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateAopWorkflow(request *model.CreateAopWorkflowRequest) (*model.CreateAopWorkflowResponse, error) {
	requestDef := GenReqDefForCreateAopWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAopWorkflowResponse), nil
	}
}

// CreateAopWorkflowInvoker 创建流程
func (c *SecMasterClient) CreateAopWorkflowInvoker(request *model.CreateAopWorkflowRequest) *CreateAopWorkflowInvoker {
	requestDef := GenReqDefForCreateAopWorkflow()
	return &CreateAopWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAopWorkflowVersion 创建流程版本
//
// 创建流程版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateAopWorkflowVersion(request *model.CreateAopWorkflowVersionRequest) (*model.CreateAopWorkflowVersionResponse, error) {
	requestDef := GenReqDefForCreateAopWorkflowVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAopWorkflowVersionResponse), nil
	}
}

// CreateAopWorkflowVersionInvoker 创建流程版本
func (c *SecMasterClient) CreateAopWorkflowVersionInvoker(request *model.CreateAopWorkflowVersionRequest) *CreateAopWorkflowVersionInvoker {
	requestDef := GenReqDefForCreateAopWorkflowVersion()
	return &CreateAopWorkflowVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAopWorkflowVersionApprovel 审核流程版本的发布
//
// 审核流程版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateAopWorkflowVersionApprovel(request *model.CreateAopWorkflowVersionApprovelRequest) (*model.CreateAopWorkflowVersionApprovelResponse, error) {
	requestDef := GenReqDefForCreateAopWorkflowVersionApprovel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAopWorkflowVersionApprovelResponse), nil
	}
}

// CreateAopWorkflowVersionApprovelInvoker 审核流程版本的发布
func (c *SecMasterClient) CreateAopWorkflowVersionApprovelInvoker(request *model.CreateAopWorkflowVersionApprovelRequest) *CreateAopWorkflowVersionApprovelInvoker {
	requestDef := GenReqDefForCreateAopWorkflowVersionApprovel()
	return &CreateAopWorkflowVersionApprovelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBatchOrderAlerts 告警转事件
//
// 告警转事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateBatchOrderAlerts(request *model.CreateBatchOrderAlertsRequest) (*model.CreateBatchOrderAlertsResponse, error) {
	requestDef := GenReqDefForCreateBatchOrderAlerts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBatchOrderAlertsResponse), nil
	}
}

// CreateBatchOrderAlertsInvoker 告警转事件
func (c *SecMasterClient) CreateBatchOrderAlertsInvoker(request *model.CreateBatchOrderAlertsRequest) *CreateBatchOrderAlertsInvoker {
	requestDef := GenReqDefForCreateBatchOrderAlerts()
	return &CreateBatchOrderAlertsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCatalogue 创建自定义目录
//
// 新增自定义目录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateCatalogue(request *model.CreateCatalogueRequest) (*model.CreateCatalogueResponse, error) {
	requestDef := GenReqDefForCreateCatalogue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCatalogueResponse), nil
	}
}

// CreateCatalogueInvoker 创建自定义目录
func (c *SecMasterClient) CreateCatalogueInvoker(request *model.CreateCatalogueRequest) *CreateCatalogueInvoker {
	requestDef := GenReqDefForCreateCatalogue()
	return &CreateCatalogueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateClassifier 新增分类
//
// 新增分类
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateClassifier(request *model.CreateClassifierRequest) (*model.CreateClassifierResponse, error) {
	requestDef := GenReqDefForCreateClassifier()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClassifierResponse), nil
	}
}

// CreateClassifierInvoker 新增分类
func (c *SecMasterClient) CreateClassifierInvoker(request *model.CreateClassifierRequest) *CreateClassifierInvoker {
	requestDef := GenReqDefForCreateClassifier()
	return &CreateClassifierInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCollectConfig 保存云服务采集配置
//
// 保存云服务采集配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateCollectConfig(request *model.CreateCollectConfigRequest) (*model.CreateCollectConfigResponse, error) {
	requestDef := GenReqDefForCreateCollectConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCollectConfigResponse), nil
	}
}

// CreateCollectConfigInvoker 保存云服务采集配置
func (c *SecMasterClient) CreateCollectConfigInvoker(request *model.CreateCollectConfigRequest) *CreateCollectConfigInvoker {
	requestDef := GenReqDefForCreateCollectConfig()
	return &CreateCollectConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCollectorChannel 创建采集通道
//
// 创建采集通道
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateCollectorChannel(request *model.CreateCollectorChannelRequest) (*model.CreateCollectorChannelResponse, error) {
	requestDef := GenReqDefForCreateCollectorChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCollectorChannelResponse), nil
	}
}

// CreateCollectorChannelInvoker 创建采集通道
func (c *SecMasterClient) CreateCollectorChannelInvoker(request *model.CreateCollectorChannelRequest) *CreateCollectorChannelInvoker {
	requestDef := GenReqDefForCreateCollectorChannel()
	return &CreateCollectorChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCollectorChannelGroup 创建采集通道分组
//
// 创建采集通道分组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateCollectorChannelGroup(request *model.CreateCollectorChannelGroupRequest) (*model.CreateCollectorChannelGroupResponse, error) {
	requestDef := GenReqDefForCreateCollectorChannelGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCollectorChannelGroupResponse), nil
	}
}

// CreateCollectorChannelGroupInvoker 创建采集通道分组
func (c *SecMasterClient) CreateCollectorChannelGroupInvoker(request *model.CreateCollectorChannelGroupRequest) *CreateCollectorChannelGroupInvoker {
	requestDef := GenReqDefForCreateCollectorChannelGroup()
	return &CreateCollectorChannelGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCollectorChannelOperation 执行采集通道操作
//
// 执行采集通道操作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateCollectorChannelOperation(request *model.CreateCollectorChannelOperationRequest) (*model.CreateCollectorChannelOperationResponse, error) {
	requestDef := GenReqDefForCreateCollectorChannelOperation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCollectorChannelOperationResponse), nil
	}
}

// CreateCollectorChannelOperationInvoker 执行采集通道操作
func (c *SecMasterClient) CreateCollectorChannelOperationInvoker(request *model.CreateCollectorChannelOperationRequest) *CreateCollectorChannelOperationInvoker {
	requestDef := GenReqDefForCreateCollectorChannelOperation()
	return &CreateCollectorChannelOperationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCollectorConnection 创建采集连接
//
// 创建采集连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateCollectorConnection(request *model.CreateCollectorConnectionRequest) (*model.CreateCollectorConnectionResponse, error) {
	requestDef := GenReqDefForCreateCollectorConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCollectorConnectionResponse), nil
	}
}

// CreateCollectorConnectionInvoker 创建采集连接
func (c *SecMasterClient) CreateCollectorConnectionInvoker(request *model.CreateCollectorConnectionRequest) *CreateCollectorConnectionInvoker {
	requestDef := GenReqDefForCreateCollectorConnection()
	return &CreateCollectorConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCollectorFiles 安装采集通道文件
//
// 安装采集通道文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateCollectorFiles(request *model.CreateCollectorFilesRequest) (*model.CreateCollectorFilesResponse, error) {
	requestDef := GenReqDefForCreateCollectorFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCollectorFilesResponse), nil
	}
}

// CreateCollectorFilesInvoker 安装采集通道文件
func (c *SecMasterClient) CreateCollectorFilesInvoker(request *model.CreateCollectorFilesRequest) *CreateCollectorFilesInvoker {
	requestDef := GenReqDefForCreateCollectorFiles()
	return &CreateCollectorFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCollectorParser 创建采集解析器
//
// 创建采集解析器
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateCollectorParser(request *model.CreateCollectorParserRequest) (*model.CreateCollectorParserResponse, error) {
	requestDef := GenReqDefForCreateCollectorParser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCollectorParserResponse), nil
	}
}

// CreateCollectorParserInvoker 创建采集解析器
func (c *SecMasterClient) CreateCollectorParserInvoker(request *model.CreateCollectorParserRequest) *CreateCollectorParserInvoker {
	requestDef := GenReqDefForCreateCollectorParser()
	return &CreateCollectorParserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateComponentTemplate 创建插件配置模板
//
// 创建在配置流程时的插件配置模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateComponentTemplate(request *model.CreateComponentTemplateRequest) (*model.CreateComponentTemplateResponse, error) {
	requestDef := GenReqDefForCreateComponentTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateComponentTemplateResponse), nil
	}
}

// CreateComponentTemplateInvoker 创建插件配置模板
func (c *SecMasterClient) CreateComponentTemplateInvoker(request *model.CreateComponentTemplateRequest) *CreateComponentTemplateInvoker {
	requestDef := GenReqDefForCreateComponentTemplate()
	return &CreateComponentTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConfigurationApplication 创建配置应用
//
// 创建配置应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateConfigurationApplication(request *model.CreateConfigurationApplicationRequest) (*model.CreateConfigurationApplicationResponse, error) {
	requestDef := GenReqDefForCreateConfigurationApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConfigurationApplicationResponse), nil
	}
}

// CreateConfigurationApplicationInvoker 创建配置应用
func (c *SecMasterClient) CreateConfigurationApplicationInvoker(request *model.CreateConfigurationApplicationRequest) *CreateConfigurationApplicationInvoker {
	requestDef := GenReqDefForCreateConfigurationApplication()
	return &CreateConfigurationApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConfigurationDictionaries 创建字典
//
// 创建字典数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateConfigurationDictionaries(request *model.CreateConfigurationDictionariesRequest) (*model.CreateConfigurationDictionariesResponse, error) {
	requestDef := GenReqDefForCreateConfigurationDictionaries()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConfigurationDictionariesResponse), nil
	}
}

// CreateConfigurationDictionariesInvoker 创建字典
func (c *SecMasterClient) CreateConfigurationDictionariesInvoker(request *model.CreateConfigurationDictionariesRequest) *CreateConfigurationDictionariesInvoker {
	requestDef := GenReqDefForCreateConfigurationDictionaries()
	return &CreateConfigurationDictionariesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConnection 新建操作连接
//
// 新建操作连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateConnection(request *model.CreateConnectionRequest) (*model.CreateConnectionResponse, error) {
	requestDef := GenReqDefForCreateConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConnectionResponse), nil
	}
}

// CreateConnectionInvoker 新建操作连接
func (c *SecMasterClient) CreateConnectionInvoker(request *model.CreateConnectionRequest) *CreateConnectionInvoker {
	requestDef := GenReqDefForCreateConnection()
	return &CreateConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDataclassType 数据类类型新增
//
// 新增数据类类型
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateDataclassType(request *model.CreateDataclassTypeRequest) (*model.CreateDataclassTypeResponse, error) {
	requestDef := GenReqDefForCreateDataclassType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDataclassTypeResponse), nil
	}
}

// CreateDataclassTypeInvoker 数据类类型新增
func (c *SecMasterClient) CreateDataclassTypeInvoker(request *model.CreateDataclassTypeRequest) *CreateDataclassTypeInvoker {
	requestDef := GenReqDefForCreateDataclassType()
	return &CreateDataclassTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDataobject 创建数据对象
//
// 创建数据对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateDataobject(request *model.CreateDataobjectRequest) (*model.CreateDataobjectResponse, error) {
	requestDef := GenReqDefForCreateDataobject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDataobjectResponse), nil
	}
}

// CreateDataobjectInvoker 创建数据对象
func (c *SecMasterClient) CreateDataobjectInvoker(request *model.CreateDataobjectRequest) *CreateDataobjectInvoker {
	requestDef := GenReqDefForCreateDataobject()
	return &CreateDataobjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDataobjectRelation 关联数据对象
//
// 关联数据对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateDataobjectRelation(request *model.CreateDataobjectRelationRequest) (*model.CreateDataobjectRelationResponse, error) {
	requestDef := GenReqDefForCreateDataobjectRelation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDataobjectRelationResponse), nil
	}
}

// CreateDataobjectRelationInvoker 关联数据对象
func (c *SecMasterClient) CreateDataobjectRelationInvoker(request *model.CreateDataobjectRelationRequest) *CreateDataobjectRelationInvoker {
	requestDef := GenReqDefForCreateDataobjectRelation()
	return &CreateDataobjectRelationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDataspace 创建数据空间
//
// 创建数据空间
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateDataspace(request *model.CreateDataspaceRequest) (*model.CreateDataspaceResponse, error) {
	requestDef := GenReqDefForCreateDataspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDataspaceResponse), nil
	}
}

// CreateDataspaceInvoker 创建数据空间
func (c *SecMasterClient) CreateDataspaceInvoker(request *model.CreateDataspaceRequest) *CreateDataspaceInvoker {
	requestDef := GenReqDefForCreateDataspace()
	return &CreateDataspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateIncident 创建事件
//
// 创建事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateIncident(request *model.CreateIncidentRequest) (*model.CreateIncidentResponse, error) {
	requestDef := GenReqDefForCreateIncident()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIncidentResponse), nil
	}
}

// CreateIncidentInvoker 创建事件
func (c *SecMasterClient) CreateIncidentInvoker(request *model.CreateIncidentRequest) *CreateIncidentInvoker {
	requestDef := GenReqDefForCreateIncident()
	return &CreateIncidentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateIndicator 创建威胁情报
//
// 创建威胁情报
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateIndicator(request *model.CreateIndicatorRequest) (*model.CreateIndicatorResponse, error) {
	requestDef := GenReqDefForCreateIndicator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIndicatorResponse), nil
	}
}

// CreateIndicatorInvoker 创建威胁情报
func (c *SecMasterClient) CreateIndicatorInvoker(request *model.CreateIndicatorRequest) *CreateIndicatorInvoker {
	requestDef := GenReqDefForCreateIndicator()
	return &CreateIndicatorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLayout 创建布局
//
// 创建布局
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateLayout(request *model.CreateLayoutRequest) (*model.CreateLayoutResponse, error) {
	requestDef := GenReqDefForCreateLayout()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLayoutResponse), nil
	}
}

// CreateLayoutInvoker 创建布局
func (c *SecMasterClient) CreateLayoutInvoker(request *model.CreateLayoutRequest) *CreateLayoutInvoker {
	requestDef := GenReqDefForCreateLayout()
	return &CreateLayoutInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLayoutWizard 新建布局页面
//
// 创建页面
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateLayoutWizard(request *model.CreateLayoutWizardRequest) (*model.CreateLayoutWizardResponse, error) {
	requestDef := GenReqDefForCreateLayoutWizard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLayoutWizardResponse), nil
	}
}

// CreateLayoutWizardInvoker 新建布局页面
func (c *SecMasterClient) CreateLayoutWizardInvoker(request *model.CreateLayoutWizardRequest) *CreateLayoutWizardInvoker {
	requestDef := GenReqDefForCreateLayoutWizard()
	return &CreateLayoutWizardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMapper 新增映射
//
// 新增映射
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateMapper(request *model.CreateMapperRequest) (*model.CreateMapperResponse, error) {
	requestDef := GenReqDefForCreateMapper()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMapperResponse), nil
	}
}

// CreateMapperInvoker 新增映射
func (c *SecMasterClient) CreateMapperInvoker(request *model.CreateMapperRequest) *CreateMapperInvoker {
	requestDef := GenReqDefForCreateMapper()
	return &CreateMapperInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMetric 创建指标
//
// 创建指标
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateMetric(request *model.CreateMetricRequest) (*model.CreateMetricResponse, error) {
	requestDef := GenReqDefForCreateMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMetricResponse), nil
	}
}

// CreateMetricInvoker 创建指标
func (c *SecMasterClient) CreateMetricInvoker(request *model.CreateMetricRequest) *CreateMetricInvoker {
	requestDef := GenReqDefForCreateMetric()
	return &CreateMetricInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateModule 新建模块
//
// 创建模块.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateModule(request *model.CreateModuleRequest) (*model.CreateModuleResponse, error) {
	requestDef := GenReqDefForCreateModule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateModuleResponse), nil
	}
}

// CreateModuleInvoker 新建模块
func (c *SecMasterClient) CreateModuleInvoker(request *model.CreateModuleRequest) *CreateModuleInvoker {
	requestDef := GenReqDefForCreateModule()
	return &CreateModuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNote 创建评论
//
// 创建评论
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateNote(request *model.CreateNoteRequest) (*model.CreateNoteResponse, error) {
	requestDef := GenReqDefForCreateNote()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNoteResponse), nil
	}
}

// CreateNoteInvoker 创建评论
func (c *SecMasterClient) CreateNoteInvoker(request *model.CreateNoteRequest) *CreateNoteInvoker {
	requestDef := GenReqDefForCreateNote()
	return &CreateNoteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePipe 创建数据管道
//
// 创建数据管道
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreatePipe(request *model.CreatePipeRequest) (*model.CreatePipeResponse, error) {
	requestDef := GenReqDefForCreatePipe()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePipeResponse), nil
	}
}

// CreatePipeInvoker 创建数据管道
func (c *SecMasterClient) CreatePipeInvoker(request *model.CreatePipeRequest) *CreatePipeInvoker {
	requestDef := GenReqDefForCreatePipe()
	return &CreatePipeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePipeConsumption 开启实时消费
//
// 开启实时消费
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreatePipeConsumption(request *model.CreatePipeConsumptionRequest) (*model.CreatePipeConsumptionResponse, error) {
	requestDef := GenReqDefForCreatePipeConsumption()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePipeConsumptionResponse), nil
	}
}

// CreatePipeConsumptionInvoker 开启实时消费
func (c *SecMasterClient) CreatePipeConsumptionInvoker(request *model.CreatePipeConsumptionRequest) *CreatePipeConsumptionInvoker {
	requestDef := GenReqDefForCreatePipeConsumption()
	return &CreatePipeConsumptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePlaybook 创建剧本
//
// 创建剧本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreatePlaybook(request *model.CreatePlaybookRequest) (*model.CreatePlaybookResponse, error) {
	requestDef := GenReqDefForCreatePlaybook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePlaybookResponse), nil
	}
}

// CreatePlaybookInvoker 创建剧本
func (c *SecMasterClient) CreatePlaybookInvoker(request *model.CreatePlaybookRequest) *CreatePlaybookInvoker {
	requestDef := GenReqDefForCreatePlaybook()
	return &CreatePlaybookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePlaybookAction 创建剧本动作
//
// 创建剧本动作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreatePlaybookAction(request *model.CreatePlaybookActionRequest) (*model.CreatePlaybookActionResponse, error) {
	requestDef := GenReqDefForCreatePlaybookAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePlaybookActionResponse), nil
	}
}

// CreatePlaybookActionInvoker 创建剧本动作
func (c *SecMasterClient) CreatePlaybookActionInvoker(request *model.CreatePlaybookActionRequest) *CreatePlaybookActionInvoker {
	requestDef := GenReqDefForCreatePlaybookAction()
	return &CreatePlaybookActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePlaybookApprove 审核剧本
//
// 审核剧本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreatePlaybookApprove(request *model.CreatePlaybookApproveRequest) (*model.CreatePlaybookApproveResponse, error) {
	requestDef := GenReqDefForCreatePlaybookApprove()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePlaybookApproveResponse), nil
	}
}

// CreatePlaybookApproveInvoker 审核剧本
func (c *SecMasterClient) CreatePlaybookApproveInvoker(request *model.CreatePlaybookApproveRequest) *CreatePlaybookApproveInvoker {
	requestDef := GenReqDefForCreatePlaybookApprove()
	return &CreatePlaybookApproveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePlaybookRule 创建剧本规则
//
// 创建剧本规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreatePlaybookRule(request *model.CreatePlaybookRuleRequest) (*model.CreatePlaybookRuleResponse, error) {
	requestDef := GenReqDefForCreatePlaybookRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePlaybookRuleResponse), nil
	}
}

// CreatePlaybookRuleInvoker 创建剧本规则
func (c *SecMasterClient) CreatePlaybookRuleInvoker(request *model.CreatePlaybookRuleRequest) *CreatePlaybookRuleInvoker {
	requestDef := GenReqDefForCreatePlaybookRule()
	return &CreatePlaybookRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePlaybookVersion 创建剧本版本
//
// 创建剧本版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreatePlaybookVersion(request *model.CreatePlaybookVersionRequest) (*model.CreatePlaybookVersionResponse, error) {
	requestDef := GenReqDefForCreatePlaybookVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePlaybookVersionResponse), nil
	}
}

// CreatePlaybookVersionInvoker 创建剧本版本
func (c *SecMasterClient) CreatePlaybookVersionInvoker(request *model.CreatePlaybookVersionRequest) *CreatePlaybookVersionInvoker {
	requestDef := GenReqDefForCreatePlaybookVersion()
	return &CreatePlaybookVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePreProcessRules 创建预处理规则
//
// 创建预处理规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreatePreProcessRules(request *model.CreatePreProcessRulesRequest) (*model.CreatePreProcessRulesResponse, error) {
	requestDef := GenReqDefForCreatePreProcessRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePreProcessRulesResponse), nil
	}
}

// CreatePreProcessRulesInvoker 创建预处理规则
func (c *SecMasterClient) CreatePreProcessRulesInvoker(request *model.CreatePreProcessRulesRequest) *CreatePreProcessRulesInvoker {
	requestDef := GenReqDefForCreatePreProcessRules()
	return &CreatePreProcessRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateReport 创建安全报告
//
// 创建安全报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateReport(request *model.CreateReportRequest) (*model.CreateReportResponse, error) {
	requestDef := GenReqDefForCreateReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateReportResponse), nil
	}
}

// CreateReportInvoker 创建安全报告
func (c *SecMasterClient) CreateReportInvoker(request *model.CreateReportRequest) *CreateReportInvoker {
	requestDef := GenReqDefForCreateReport()
	return &CreateReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResourceConfig 创建云日志资源
//
// 创建云日志资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateResourceConfig(request *model.CreateResourceConfigRequest) (*model.CreateResourceConfigResponse, error) {
	requestDef := GenReqDefForCreateResourceConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResourceConfigResponse), nil
	}
}

// CreateResourceConfigInvoker 创建云日志资源
func (c *SecMasterClient) CreateResourceConfigInvoker(request *model.CreateResourceConfigRequest) *CreateResourceConfigInvoker {
	requestDef := GenReqDefForCreateResourceConfig()
	return &CreateResourceConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRetryPolicy 创建重试应急策略
//
// 创建重试应急策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateRetryPolicy(request *model.CreateRetryPolicyRequest) (*model.CreateRetryPolicyResponse, error) {
	requestDef := GenReqDefForCreateRetryPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRetryPolicyResponse), nil
	}
}

// CreateRetryPolicyInvoker 创建重试应急策略
func (c *SecMasterClient) CreateRetryPolicyInvoker(request *model.CreateRetryPolicyRequest) *CreateRetryPolicyInvoker {
	requestDef := GenReqDefForCreateRetryPolicy()
	return &CreateRetryPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSearchCondition 创建检索条件
//
// 创建检索条件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateSearchCondition(request *model.CreateSearchConditionRequest) (*model.CreateSearchConditionResponse, error) {
	requestDef := GenReqDefForCreateSearchCondition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSearchConditionResponse), nil
	}
}

// CreateSearchConditionInvoker 创建检索条件
func (c *SecMasterClient) CreateSearchConditionInvoker(request *model.CreateSearchConditionRequest) *CreateSearchConditionInvoker {
	requestDef := GenReqDefForCreateSearchCondition()
	return &CreateSearchConditionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateServiceAgency 创建委托并授权
//
// 根据body体中的角色和作用范围，创建委托，并将策略赋予给委托
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateServiceAgency(request *model.CreateServiceAgencyRequest) (*model.CreateServiceAgencyResponse, error) {
	requestDef := GenReqDefForCreateServiceAgency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateServiceAgencyResponse), nil
	}
}

// CreateServiceAgencyInvoker 创建委托并授权
func (c *SecMasterClient) CreateServiceAgencyInvoker(request *model.CreateServiceAgencyRequest) *CreateServiceAgencyInvoker {
	requestDef := GenReqDefForCreateServiceAgency()
	return &CreateServiceAgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateShipper 创建数据投递
//
// 创建数据投递
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateShipper(request *model.CreateShipperRequest) (*model.CreateShipperResponse, error) {
	requestDef := GenReqDefForCreateShipper()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateShipperResponse), nil
	}
}

// CreateShipperInvoker 创建数据投递
func (c *SecMasterClient) CreateShipperInvoker(request *model.CreateShipperRequest) *CreateShipperInvoker {
	requestDef := GenReqDefForCreateShipper()
	return &CreateShipperInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateShipperDelegateAuth 创建委托权限
//
// 创建委托权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateShipperDelegateAuth(request *model.CreateShipperDelegateAuthRequest) (*model.CreateShipperDelegateAuthResponse, error) {
	requestDef := GenReqDefForCreateShipperDelegateAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateShipperDelegateAuthResponse), nil
	}
}

// CreateShipperDelegateAuthInvoker 创建委托权限
func (c *SecMasterClient) CreateShipperDelegateAuthInvoker(request *model.CreateShipperDelegateAuthRequest) *CreateShipperDelegateAuthInvoker {
	requestDef := GenReqDefForCreateShipperDelegateAuth()
	return &CreateShipperDelegateAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkflowInstance 创建流程实例
//
// 创建流程实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateWorkflowInstance(request *model.CreateWorkflowInstanceRequest) (*model.CreateWorkflowInstanceResponse, error) {
	requestDef := GenReqDefForCreateWorkflowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkflowInstanceResponse), nil
	}
}

// CreateWorkflowInstanceInvoker 创建流程实例
func (c *SecMasterClient) CreateWorkflowInstanceInvoker(request *model.CreateWorkflowInstanceRequest) *CreateWorkflowInstanceInvoker {
	requestDef := GenReqDefForCreateWorkflowInstance()
	return &CreateWorkflowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkspace 新建工作空间
//
// 在使用安全云脑的基线检查、告警管理、安全分析、安全编排等功能前，需要创建工作空间，它可以将资源划分为各个不同的工作场景，避免资源冗余查找不便，影响日常使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateWorkspace(request *model.CreateWorkspaceRequest) (*model.CreateWorkspaceResponse, error) {
	requestDef := GenReqDefForCreateWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkspaceResponse), nil
	}
}

// CreateWorkspaceInvoker 新建工作空间
func (c *SecMasterClient) CreateWorkspaceInvoker(request *model.CreateWorkspaceRequest) *CreateWorkspaceInvoker {
	requestDef := GenReqDefForCreateWorkspace()
	return &CreateWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAlert 删除告警
//
// 删除告警
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteAlert(request *model.DeleteAlertRequest) (*model.DeleteAlertResponse, error) {
	requestDef := GenReqDefForDeleteAlert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlertResponse), nil
	}
}

// DeleteAlertInvoker 删除告警
func (c *SecMasterClient) DeleteAlertInvoker(request *model.DeleteAlertRequest) *DeleteAlertInvoker {
	requestDef := GenReqDefForDeleteAlert()
	return &DeleteAlertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAlertRule 删除告警规则
//
// 删除告警规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteAlertRule(request *model.DeleteAlertRuleRequest) (*model.DeleteAlertRuleResponse, error) {
	requestDef := GenReqDefForDeleteAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlertRuleResponse), nil
	}
}

// DeleteAlertRuleInvoker 删除告警规则
func (c *SecMasterClient) DeleteAlertRuleInvoker(request *model.DeleteAlertRuleRequest) *DeleteAlertRuleInvoker {
	requestDef := GenReqDefForDeleteAlertRule()
	return &DeleteAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAlerts 批量删除告警
//
// 批量删除告警
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteAlerts(request *model.DeleteAlertsRequest) (*model.DeleteAlertsResponse, error) {
	requestDef := GenReqDefForDeleteAlerts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlertsResponse), nil
	}
}

// DeleteAlertsInvoker 批量删除告警
func (c *SecMasterClient) DeleteAlertsInvoker(request *model.DeleteAlertsRequest) *DeleteAlertsInvoker {
	requestDef := GenReqDefForDeleteAlerts()
	return &DeleteAlertsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAopWorkflow 删除流程
//
// 删除流程
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteAopWorkflow(request *model.DeleteAopWorkflowRequest) (*model.DeleteAopWorkflowResponse, error) {
	requestDef := GenReqDefForDeleteAopWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAopWorkflowResponse), nil
	}
}

// DeleteAopWorkflowInvoker 删除流程
func (c *SecMasterClient) DeleteAopWorkflowInvoker(request *model.DeleteAopWorkflowRequest) *DeleteAopWorkflowInvoker {
	requestDef := GenReqDefForDeleteAopWorkflow()
	return &DeleteAopWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAopWorkflowVersion 删除流程版本
//
// 删除流程版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteAopWorkflowVersion(request *model.DeleteAopWorkflowVersionRequest) (*model.DeleteAopWorkflowVersionResponse, error) {
	requestDef := GenReqDefForDeleteAopWorkflowVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAopWorkflowVersionResponse), nil
	}
}

// DeleteAopWorkflowVersionInvoker 删除流程版本
func (c *SecMasterClient) DeleteAopWorkflowVersionInvoker(request *model.DeleteAopWorkflowVersionRequest) *DeleteAopWorkflowVersionInvoker {
	requestDef := GenReqDefForDeleteAopWorkflowVersion()
	return &DeleteAopWorkflowVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCatalogue 批量删除目录
//
// 批量删除目录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteCatalogue(request *model.DeleteCatalogueRequest) (*model.DeleteCatalogueResponse, error) {
	requestDef := GenReqDefForDeleteCatalogue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCatalogueResponse), nil
	}
}

// DeleteCatalogueInvoker 批量删除目录
func (c *SecMasterClient) DeleteCatalogueInvoker(request *model.DeleteCatalogueRequest) *DeleteCatalogueInvoker {
	requestDef := GenReqDefForDeleteCatalogue()
	return &DeleteCatalogueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCollectorChannel 删除采集通道
//
// 删除采集通道
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteCollectorChannel(request *model.DeleteCollectorChannelRequest) (*model.DeleteCollectorChannelResponse, error) {
	requestDef := GenReqDefForDeleteCollectorChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCollectorChannelResponse), nil
	}
}

// DeleteCollectorChannelInvoker 删除采集通道
func (c *SecMasterClient) DeleteCollectorChannelInvoker(request *model.DeleteCollectorChannelRequest) *DeleteCollectorChannelInvoker {
	requestDef := GenReqDefForDeleteCollectorChannel()
	return &DeleteCollectorChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCollectorChannelGroup 删除采集通道分组
//
// 删除采集通道分组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteCollectorChannelGroup(request *model.DeleteCollectorChannelGroupRequest) (*model.DeleteCollectorChannelGroupResponse, error) {
	requestDef := GenReqDefForDeleteCollectorChannelGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCollectorChannelGroupResponse), nil
	}
}

// DeleteCollectorChannelGroupInvoker 删除采集通道分组
func (c *SecMasterClient) DeleteCollectorChannelGroupInvoker(request *model.DeleteCollectorChannelGroupRequest) *DeleteCollectorChannelGroupInvoker {
	requestDef := GenReqDefForDeleteCollectorChannelGroup()
	return &DeleteCollectorChannelGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCollectorConnection 删除采集连接
//
// 删除采集连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteCollectorConnection(request *model.DeleteCollectorConnectionRequest) (*model.DeleteCollectorConnectionResponse, error) {
	requestDef := GenReqDefForDeleteCollectorConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCollectorConnectionResponse), nil
	}
}

// DeleteCollectorConnectionInvoker 删除采集连接
func (c *SecMasterClient) DeleteCollectorConnectionInvoker(request *model.DeleteCollectorConnectionRequest) *DeleteCollectorConnectionInvoker {
	requestDef := GenReqDefForDeleteCollectorConnection()
	return &DeleteCollectorConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCollectorParser 删除采集解析器
//
// 删除采集解析器
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteCollectorParser(request *model.DeleteCollectorParserRequest) (*model.DeleteCollectorParserResponse, error) {
	requestDef := GenReqDefForDeleteCollectorParser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCollectorParserResponse), nil
	}
}

// DeleteCollectorParserInvoker 删除采集解析器
func (c *SecMasterClient) DeleteCollectorParserInvoker(request *model.DeleteCollectorParserRequest) *DeleteCollectorParserInvoker {
	requestDef := GenReqDefForDeleteCollectorParser()
	return &DeleteCollectorParserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteComponentTemplate 删除插件配置模板
//
// 删除某个在配置流程时的插件配置模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteComponentTemplate(request *model.DeleteComponentTemplateRequest) (*model.DeleteComponentTemplateResponse, error) {
	requestDef := GenReqDefForDeleteComponentTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteComponentTemplateResponse), nil
	}
}

// DeleteComponentTemplateInvoker 删除插件配置模板
func (c *SecMasterClient) DeleteComponentTemplateInvoker(request *model.DeleteComponentTemplateRequest) *DeleteComponentTemplateInvoker {
	requestDef := GenReqDefForDeleteComponentTemplate()
	return &DeleteComponentTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConfigurationDictionaries 删除字典
//
// 删除字典数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteConfigurationDictionaries(request *model.DeleteConfigurationDictionariesRequest) (*model.DeleteConfigurationDictionariesResponse, error) {
	requestDef := GenReqDefForDeleteConfigurationDictionaries()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConfigurationDictionariesResponse), nil
	}
}

// DeleteConfigurationDictionariesInvoker 删除字典
func (c *SecMasterClient) DeleteConfigurationDictionariesInvoker(request *model.DeleteConfigurationDictionariesRequest) *DeleteConfigurationDictionariesInvoker {
	requestDef := GenReqDefForDeleteConfigurationDictionaries()
	return &DeleteConfigurationDictionariesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConnection 删除操作连接
//
// 删除操作连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteConnection(request *model.DeleteConnectionRequest) (*model.DeleteConnectionResponse, error) {
	requestDef := GenReqDefForDeleteConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConnectionResponse), nil
	}
}

// DeleteConnectionInvoker 删除操作连接
func (c *SecMasterClient) DeleteConnectionInvoker(request *model.DeleteConnectionRequest) *DeleteConnectionInvoker {
	requestDef := GenReqDefForDeleteConnection()
	return &DeleteConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDataclassType 数据类类型删除
//
// 删除数据类类型
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteDataclassType(request *model.DeleteDataclassTypeRequest) (*model.DeleteDataclassTypeResponse, error) {
	requestDef := GenReqDefForDeleteDataclassType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDataclassTypeResponse), nil
	}
}

// DeleteDataclassTypeInvoker 数据类类型删除
func (c *SecMasterClient) DeleteDataclassTypeInvoker(request *model.DeleteDataclassTypeRequest) *DeleteDataclassTypeInvoker {
	requestDef := GenReqDefForDeleteDataclassType()
	return &DeleteDataclassTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDataobjectRelation 取消关联数据对象
//
// 取消关联数据对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteDataobjectRelation(request *model.DeleteDataobjectRelationRequest) (*model.DeleteDataobjectRelationResponse, error) {
	requestDef := GenReqDefForDeleteDataobjectRelation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDataobjectRelationResponse), nil
	}
}

// DeleteDataobjectRelationInvoker 取消关联数据对象
func (c *SecMasterClient) DeleteDataobjectRelationInvoker(request *model.DeleteDataobjectRelationRequest) *DeleteDataobjectRelationInvoker {
	requestDef := GenReqDefForDeleteDataobjectRelation()
	return &DeleteDataobjectRelationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDataobjects 批量删除数据对象
//
// 批量删除数据对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteDataobjects(request *model.DeleteDataobjectsRequest) (*model.DeleteDataobjectsResponse, error) {
	requestDef := GenReqDefForDeleteDataobjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDataobjectsResponse), nil
	}
}

// DeleteDataobjectsInvoker 批量删除数据对象
func (c *SecMasterClient) DeleteDataobjectsInvoker(request *model.DeleteDataobjectsRequest) *DeleteDataobjectsInvoker {
	requestDef := GenReqDefForDeleteDataobjects()
	return &DeleteDataobjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDataspace 删除数据空间
//
// 删除数据空间
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteDataspace(request *model.DeleteDataspaceRequest) (*model.DeleteDataspaceResponse, error) {
	requestDef := GenReqDefForDeleteDataspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDataspaceResponse), nil
	}
}

// DeleteDataspaceInvoker 删除数据空间
func (c *SecMasterClient) DeleteDataspaceInvoker(request *model.DeleteDataspaceRequest) *DeleteDataspaceInvoker {
	requestDef := GenReqDefForDeleteDataspace()
	return &DeleteDataspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteIncident 删除事件
//
// 删除事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteIncident(request *model.DeleteIncidentRequest) (*model.DeleteIncidentResponse, error) {
	requestDef := GenReqDefForDeleteIncident()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIncidentResponse), nil
	}
}

// DeleteIncidentInvoker 删除事件
func (c *SecMasterClient) DeleteIncidentInvoker(request *model.DeleteIncidentRequest) *DeleteIncidentInvoker {
	requestDef := GenReqDefForDeleteIncident()
	return &DeleteIncidentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteIncidents 批量删除事件
//
// 批量删除事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteIncidents(request *model.DeleteIncidentsRequest) (*model.DeleteIncidentsResponse, error) {
	requestDef := GenReqDefForDeleteIncidents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIncidentsResponse), nil
	}
}

// DeleteIncidentsInvoker 批量删除事件
func (c *SecMasterClient) DeleteIncidentsInvoker(request *model.DeleteIncidentsRequest) *DeleteIncidentsInvoker {
	requestDef := GenReqDefForDeleteIncidents()
	return &DeleteIncidentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteIndicator 删除威胁情报
//
// 删除威胁情报
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteIndicator(request *model.DeleteIndicatorRequest) (*model.DeleteIndicatorResponse, error) {
	requestDef := GenReqDefForDeleteIndicator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIndicatorResponse), nil
	}
}

// DeleteIndicatorInvoker 删除威胁情报
func (c *SecMasterClient) DeleteIndicatorInvoker(request *model.DeleteIndicatorRequest) *DeleteIndicatorInvoker {
	requestDef := GenReqDefForDeleteIndicator()
	return &DeleteIndicatorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLayoutWizard 删除布局页面
//
// 删除页面
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteLayoutWizard(request *model.DeleteLayoutWizardRequest) (*model.DeleteLayoutWizardResponse, error) {
	requestDef := GenReqDefForDeleteLayoutWizard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLayoutWizardResponse), nil
	}
}

// DeleteLayoutWizardInvoker 删除布局页面
func (c *SecMasterClient) DeleteLayoutWizardInvoker(request *model.DeleteLayoutWizardRequest) *DeleteLayoutWizardInvoker {
	requestDef := GenReqDefForDeleteLayoutWizard()
	return &DeleteLayoutWizardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLayouts 删除布局
//
// 删除布局
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteLayouts(request *model.DeleteLayoutsRequest) (*model.DeleteLayoutsResponse, error) {
	requestDef := GenReqDefForDeleteLayouts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLayoutsResponse), nil
	}
}

// DeleteLayoutsInvoker 删除布局
func (c *SecMasterClient) DeleteLayoutsInvoker(request *model.DeleteLayoutsRequest) *DeleteLayoutsInvoker {
	requestDef := GenReqDefForDeleteLayouts()
	return &DeleteLayoutsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLtsCloudLog 删除lts日志订阅
//
// 删除云日志资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteLtsCloudLog(request *model.DeleteLtsCloudLogRequest) (*model.DeleteLtsCloudLogResponse, error) {
	requestDef := GenReqDefForDeleteLtsCloudLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLtsCloudLogResponse), nil
	}
}

// DeleteLtsCloudLogInvoker 删除lts日志订阅
func (c *SecMasterClient) DeleteLtsCloudLogInvoker(request *model.DeleteLtsCloudLogRequest) *DeleteLtsCloudLogInvoker {
	requestDef := GenReqDefForDeleteLtsCloudLog()
	return &DeleteLtsCloudLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMappingInfo 删除映射信息
//
// 删除映射信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteMappingInfo(request *model.DeleteMappingInfoRequest) (*model.DeleteMappingInfoResponse, error) {
	requestDef := GenReqDefForDeleteMappingInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMappingInfoResponse), nil
	}
}

// DeleteMappingInfoInvoker 删除映射信息
func (c *SecMasterClient) DeleteMappingInfoInvoker(request *model.DeleteMappingInfoRequest) *DeleteMappingInfoInvoker {
	requestDef := GenReqDefForDeleteMappingInfo()
	return &DeleteMappingInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMetrics 删除指标
//
// 删除指标
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteMetrics(request *model.DeleteMetricsRequest) (*model.DeleteMetricsResponse, error) {
	requestDef := GenReqDefForDeleteMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMetricsResponse), nil
	}
}

// DeleteMetricsInvoker 删除指标
func (c *SecMasterClient) DeleteMetricsInvoker(request *model.DeleteMetricsRequest) *DeleteMetricsInvoker {
	requestDef := GenReqDefForDeleteMetrics()
	return &DeleteMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteModule 删除模块
//
// 删除模块
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteModule(request *model.DeleteModuleRequest) (*model.DeleteModuleResponse, error) {
	requestDef := GenReqDefForDeleteModule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteModuleResponse), nil
	}
}

// DeleteModuleInvoker 删除模块
func (c *SecMasterClient) DeleteModuleInvoker(request *model.DeleteModuleRequest) *DeleteModuleInvoker {
	requestDef := GenReqDefForDeleteModule()
	return &DeleteModuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNode 通过节点id删除节点
//
// 删除节点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteNode(request *model.DeleteNodeRequest) (*model.DeleteNodeResponse, error) {
	requestDef := GenReqDefForDeleteNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNodeResponse), nil
	}
}

// DeleteNodeInvoker 通过节点id删除节点
func (c *SecMasterClient) DeleteNodeInvoker(request *model.DeleteNodeRequest) *DeleteNodeInvoker {
	requestDef := GenReqDefForDeleteNode()
	return &DeleteNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNotes 删除评论
//
// 删除评论
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteNotes(request *model.DeleteNotesRequest) (*model.DeleteNotesResponse, error) {
	requestDef := GenReqDefForDeleteNotes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNotesResponse), nil
	}
}

// DeleteNotesInvoker 删除评论
func (c *SecMasterClient) DeleteNotesInvoker(request *model.DeleteNotesRequest) *DeleteNotesInvoker {
	requestDef := GenReqDefForDeleteNotes()
	return &DeleteNotesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePipe 删除数据管道
//
// 删除数据管道
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeletePipe(request *model.DeletePipeRequest) (*model.DeletePipeResponse, error) {
	requestDef := GenReqDefForDeletePipe()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePipeResponse), nil
	}
}

// DeletePipeInvoker 删除数据管道
func (c *SecMasterClient) DeletePipeInvoker(request *model.DeletePipeRequest) *DeletePipeInvoker {
	requestDef := GenReqDefForDeletePipe()
	return &DeletePipeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePipeConsumption 关闭实时消费
//
// 关闭实时消费
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeletePipeConsumption(request *model.DeletePipeConsumptionRequest) (*model.DeletePipeConsumptionResponse, error) {
	requestDef := GenReqDefForDeletePipeConsumption()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePipeConsumptionResponse), nil
	}
}

// DeletePipeConsumptionInvoker 关闭实时消费
func (c *SecMasterClient) DeletePipeConsumptionInvoker(request *model.DeletePipeConsumptionRequest) *DeletePipeConsumptionInvoker {
	requestDef := GenReqDefForDeletePipeConsumption()
	return &DeletePipeConsumptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePlaybook 删除剧本
//
// 删除剧本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeletePlaybook(request *model.DeletePlaybookRequest) (*model.DeletePlaybookResponse, error) {
	requestDef := GenReqDefForDeletePlaybook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePlaybookResponse), nil
	}
}

// DeletePlaybookInvoker 删除剧本
func (c *SecMasterClient) DeletePlaybookInvoker(request *model.DeletePlaybookRequest) *DeletePlaybookInvoker {
	requestDef := GenReqDefForDeletePlaybook()
	return &DeletePlaybookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePlaybookAction 删除剧本动作
//
// 删除剧本动作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeletePlaybookAction(request *model.DeletePlaybookActionRequest) (*model.DeletePlaybookActionResponse, error) {
	requestDef := GenReqDefForDeletePlaybookAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePlaybookActionResponse), nil
	}
}

// DeletePlaybookActionInvoker 删除剧本动作
func (c *SecMasterClient) DeletePlaybookActionInvoker(request *model.DeletePlaybookActionRequest) *DeletePlaybookActionInvoker {
	requestDef := GenReqDefForDeletePlaybookAction()
	return &DeletePlaybookActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePlaybookRule 删除剧本规则
//
// 删除剧本规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeletePlaybookRule(request *model.DeletePlaybookRuleRequest) (*model.DeletePlaybookRuleResponse, error) {
	requestDef := GenReqDefForDeletePlaybookRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePlaybookRuleResponse), nil
	}
}

// DeletePlaybookRuleInvoker 删除剧本规则
func (c *SecMasterClient) DeletePlaybookRuleInvoker(request *model.DeletePlaybookRuleRequest) *DeletePlaybookRuleInvoker {
	requestDef := GenReqDefForDeletePlaybookRule()
	return &DeletePlaybookRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePlaybookVersion 删除剧本版本
//
// 删除剧本版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeletePlaybookVersion(request *model.DeletePlaybookVersionRequest) (*model.DeletePlaybookVersionResponse, error) {
	requestDef := GenReqDefForDeletePlaybookVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePlaybookVersionResponse), nil
	}
}

// DeletePlaybookVersionInvoker 删除剧本版本
func (c *SecMasterClient) DeletePlaybookVersionInvoker(request *model.DeletePlaybookVersionRequest) *DeletePlaybookVersionInvoker {
	requestDef := GenReqDefForDeletePlaybookVersion()
	return &DeletePlaybookVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePolicies 批量删除应急策略
//
// 批量删除应急策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeletePolicies(request *model.DeletePoliciesRequest) (*model.DeletePoliciesResponse, error) {
	requestDef := GenReqDefForDeletePolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePoliciesResponse), nil
	}
}

// DeletePoliciesInvoker 批量删除应急策略
func (c *SecMasterClient) DeletePoliciesInvoker(request *model.DeletePoliciesRequest) *DeletePoliciesInvoker {
	requestDef := GenReqDefForDeletePolicies()
	return &DeletePoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteReport 删除报告
//
// 删除报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteReport(request *model.DeleteReportRequest) (*model.DeleteReportResponse, error) {
	requestDef := GenReqDefForDeleteReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteReportResponse), nil
	}
}

// DeleteReportInvoker 删除报告
func (c *SecMasterClient) DeleteReportInvoker(request *model.DeleteReportRequest) *DeleteReportInvoker {
	requestDef := GenReqDefForDeleteReport()
	return &DeleteReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResource 删除资产
//
// 删除资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteResource(request *model.DeleteResourceRequest) (*model.DeleteResourceResponse, error) {
	requestDef := GenReqDefForDeleteResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceResponse), nil
	}
}

// DeleteResourceInvoker 删除资产
func (c *SecMasterClient) DeleteResourceInvoker(request *model.DeleteResourceRequest) *DeleteResourceInvoker {
	requestDef := GenReqDefForDeleteResource()
	return &DeleteResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSearchCondition 删除检索条件
//
// 删除检索条件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteSearchCondition(request *model.DeleteSearchConditionRequest) (*model.DeleteSearchConditionResponse, error) {
	requestDef := GenReqDefForDeleteSearchCondition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSearchConditionResponse), nil
	}
}

// DeleteSearchConditionInvoker 删除检索条件
func (c *SecMasterClient) DeleteSearchConditionInvoker(request *model.DeleteSearchConditionRequest) *DeleteSearchConditionInvoker {
	requestDef := GenReqDefForDeleteSearchCondition()
	return &DeleteSearchConditionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteShipper 删除投递信息
//
// 删除投递信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteShipper(request *model.DeleteShipperRequest) (*model.DeleteShipperResponse, error) {
	requestDef := GenReqDefForDeleteShipper()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteShipperResponse), nil
	}
}

// DeleteShipperInvoker 删除投递信息
func (c *SecMasterClient) DeleteShipperInvoker(request *model.DeleteShipperRequest) *DeleteShipperInvoker {
	requestDef := GenReqDefForDeleteShipper()
	return &DeleteShipperInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSingleMapper 删除单个映射
//
// 删除单个映射
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteSingleMapper(request *model.DeleteSingleMapperRequest) (*model.DeleteSingleMapperResponse, error) {
	requestDef := GenReqDefForDeleteSingleMapper()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSingleMapperResponse), nil
	}
}

// DeleteSingleMapperInvoker 删除单个映射
func (c *SecMasterClient) DeleteSingleMapperInvoker(request *model.DeleteSingleMapperRequest) *DeleteSingleMapperInvoker {
	requestDef := GenReqDefForDeleteSingleMapper()
	return &DeleteSingleMapperInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTags 删除资源标签
//
// 为指定实例批量删除标签
// 标签管理服务需要使用该接口批量管理实例的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteTags(request *model.DeleteTagsRequest) (*model.DeleteTagsResponse, error) {
	requestDef := GenReqDefForDeleteTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagsResponse), nil
	}
}

// DeleteTagsInvoker 删除资源标签
func (c *SecMasterClient) DeleteTagsInvoker(request *model.DeleteTagsRequest) *DeleteTagsInvoker {
	requestDef := GenReqDefForDeleteTags()
	return &DeleteTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWorkspace 删除工作空间
//
// 删除工作空间
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteWorkspace(request *model.DeleteWorkspaceRequest) (*model.DeleteWorkspaceResponse, error) {
	requestDef := GenReqDefForDeleteWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWorkspaceResponse), nil
	}
}

// DeleteWorkspaceInvoker 删除工作空间
func (c *SecMasterClient) DeleteWorkspaceInvoker(request *model.DeleteWorkspaceRequest) *DeleteWorkspaceInvoker {
	requestDef := GenReqDefForDeleteWorkspace()
	return &DeleteWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableAlertRule 停用告警规则
//
// 停用告警规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DisableAlertRule(request *model.DisableAlertRuleRequest) (*model.DisableAlertRuleResponse, error) {
	requestDef := GenReqDefForDisableAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableAlertRuleResponse), nil
	}
}

// DisableAlertRuleInvoker 停用告警规则
func (c *SecMasterClient) DisableAlertRuleInvoker(request *model.DisableAlertRuleRequest) *DisableAlertRuleInvoker {
	requestDef := GenReqDefForDisableAlertRule()
	return &DisableAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadAlertTemplate 下载告警导入模板
//
// 下载告警导入模板，模板根据标准的dataclass字段属性来组织，下载时需要实现默认的告警样例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DownloadAlertTemplate(request *model.DownloadAlertTemplateRequest) (*model.DownloadAlertTemplateResponse, error) {
	requestDef := GenReqDefForDownloadAlertTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadAlertTemplateResponse), nil
	}
}

// DownloadAlertTemplateInvoker 下载告警导入模板
func (c *SecMasterClient) DownloadAlertTemplateInvoker(request *model.DownloadAlertTemplateRequest) *DownloadAlertTemplateInvoker {
	requestDef := GenReqDefForDownloadAlertTemplate()
	return &DownloadAlertTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadAttachment 下载附件
//
// 从OBS下载附件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DownloadAttachment(request *model.DownloadAttachmentRequest) (*model.DownloadAttachmentResponse, error) {
	requestDef := GenReqDefForDownloadAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadAttachmentResponse), nil
	}
}

// DownloadAttachmentInvoker 下载附件
func (c *SecMasterClient) DownloadAttachmentInvoker(request *model.DownloadAttachmentRequest) *DownloadAttachmentInvoker {
	requestDef := GenReqDefForDownloadAttachment()
	return &DownloadAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadIncidentTemplate 下载事件导入模板
//
// 下载事件导入模板，模板根据标准的dataclass字段属性来组织，下载时需要实现默认的事件样例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DownloadIncidentTemplate(request *model.DownloadIncidentTemplateRequest) (*model.DownloadIncidentTemplateResponse, error) {
	requestDef := GenReqDefForDownloadIncidentTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadIncidentTemplateResponse), nil
	}
}

// DownloadIncidentTemplateInvoker 下载事件导入模板
func (c *SecMasterClient) DownloadIncidentTemplateInvoker(request *model.DownloadIncidentTemplateRequest) *DownloadIncidentTemplateInvoker {
	requestDef := GenReqDefForDownloadIncidentTemplate()
	return &DownloadIncidentTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadIndicatorTemplate 下载情报模板
//
// 下载情报导入模板，模板根据标准的dataclass字段属性来组织，下载时需要实现默认的情报样例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DownloadIndicatorTemplate(request *model.DownloadIndicatorTemplateRequest) (*model.DownloadIndicatorTemplateResponse, error) {
	requestDef := GenReqDefForDownloadIndicatorTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadIndicatorTemplateResponse), nil
	}
}

// DownloadIndicatorTemplateInvoker 下载情报模板
func (c *SecMasterClient) DownloadIndicatorTemplateInvoker(request *model.DownloadIndicatorTemplateRequest) *DownloadIndicatorTemplateInvoker {
	requestDef := GenReqDefForDownloadIndicatorTemplate()
	return &DownloadIndicatorTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadResourceTemplate 下载资产导入模板
//
// 下载资产导入模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DownloadResourceTemplate(request *model.DownloadResourceTemplateRequest) (*model.DownloadResourceTemplateResponse, error) {
	requestDef := GenReqDefForDownloadResourceTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadResourceTemplateResponse), nil
	}
}

// DownloadResourceTemplateInvoker 下载资产导入模板
func (c *SecMasterClient) DownloadResourceTemplateInvoker(request *model.DownloadResourceTemplateRequest) *DownloadResourceTemplateInvoker {
	requestDef := GenReqDefForDownloadResourceTemplate()
	return &DownloadResourceTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadVulnerabilityTemplate 下载漏洞导入模板
//
// 下载漏洞导入模板，模板根据标准的dataclass字段属性来组织
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DownloadVulnerabilityTemplate(request *model.DownloadVulnerabilityTemplateRequest) (*model.DownloadVulnerabilityTemplateResponse, error) {
	requestDef := GenReqDefForDownloadVulnerabilityTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadVulnerabilityTemplateResponse), nil
	}
}

// DownloadVulnerabilityTemplateInvoker 下载漏洞导入模板
func (c *SecMasterClient) DownloadVulnerabilityTemplateInvoker(request *model.DownloadVulnerabilityTemplateRequest) *DownloadVulnerabilityTemplateInvoker {
	requestDef := GenReqDefForDownloadVulnerabilityTemplate()
	return &DownloadVulnerabilityTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableAlertRule 启用告警规则
//
// 启用告警规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) EnableAlertRule(request *model.EnableAlertRuleRequest) (*model.EnableAlertRuleResponse, error) {
	requestDef := GenReqDefForEnableAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableAlertRuleResponse), nil
	}
}

// EnableAlertRuleInvoker 启用告警规则
func (c *SecMasterClient) EnableAlertRuleInvoker(request *model.EnableAlertRuleRequest) *EnableAlertRuleInvoker {
	requestDef := GenReqDefForEnableAlertRule()
	return &EnableAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableDataclassType 数据类类型启用/禁用
//
// 禁用/启用数据类类型
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) EnableDataclassType(request *model.EnableDataclassTypeRequest) (*model.EnableDataclassTypeResponse, error) {
	requestDef := GenReqDefForEnableDataclassType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableDataclassTypeResponse), nil
	}
}

// EnableDataclassTypeInvoker 数据类类型启用/禁用
func (c *SecMasterClient) EnableDataclassTypeInvoker(request *model.EnableDataclassTypeRequest) *EnableDataclassTypeInvoker {
	requestDef := GenReqDefForEnableDataclassType()
	return &EnableDataclassTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteReportAction 操作安全报告
//
// 操作安全报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ExecuteReportAction(request *model.ExecuteReportActionRequest) (*model.ExecuteReportActionResponse, error) {
	requestDef := GenReqDefForExecuteReportAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteReportActionResponse), nil
	}
}

// ExecuteReportActionInvoker 操作安全报告
func (c *SecMasterClient) ExecuteReportActionInvoker(request *model.ExecuteReportActionRequest) *ExecuteReportActionInvoker {
	requestDef := GenReqDefForExecuteReportAction()
	return &ExecuteReportActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportAlerts 导出告警
//
// 导出告警,若字段是object类型，则将整个子对象的内容导出
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ExportAlerts(request *model.ExportAlertsRequest) (*model.ExportAlertsResponse, error) {
	requestDef := GenReqDefForExportAlerts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportAlertsResponse), nil
	}
}

// ExportAlertsInvoker 导出告警
func (c *SecMasterClient) ExportAlertsInvoker(request *model.ExportAlertsRequest) *ExportAlertsInvoker {
	requestDef := GenReqDefForExportAlerts()
	return &ExportAlertsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportAopworkflow 导出流程
//
// 导出流程信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ExportAopworkflow(request *model.ExportAopworkflowRequest) (*model.ExportAopworkflowResponse, error) {
	requestDef := GenReqDefForExportAopworkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportAopworkflowResponse), nil
	}
}

// ExportAopworkflowInvoker 导出流程
func (c *SecMasterClient) ExportAopworkflowInvoker(request *model.ExportAopworkflowRequest) *ExportAopworkflowInvoker {
	requestDef := GenReqDefForExportAopworkflow()
	return &ExportAopworkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportCollectorParser 导出采集解析器
//
// 导出采集解析器
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ExportCollectorParser(request *model.ExportCollectorParserRequest) (*model.ExportCollectorParserResponse, error) {
	requestDef := GenReqDefForExportCollectorParser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportCollectorParserResponse), nil
	}
}

// ExportCollectorParserInvoker 导出采集解析器
func (c *SecMasterClient) ExportCollectorParserInvoker(request *model.ExportCollectorParserRequest) *ExportCollectorParserInvoker {
	requestDef := GenReqDefForExportCollectorParser()
	return &ExportCollectorParserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportDataobjects 导出数据对象
//
// 导出数据对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ExportDataobjects(request *model.ExportDataobjectsRequest) (*model.ExportDataobjectsResponse, error) {
	requestDef := GenReqDefForExportDataobjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportDataobjectsResponse), nil
	}
}

// ExportDataobjectsInvoker 导出数据对象
func (c *SecMasterClient) ExportDataobjectsInvoker(request *model.ExportDataobjectsRequest) *ExportDataobjectsInvoker {
	requestDef := GenReqDefForExportDataobjects()
	return &ExportDataobjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportIncidents 导出事件
//
// 导出事件列表,若字段是object类型，需要将整个子对象的内容导出
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ExportIncidents(request *model.ExportIncidentsRequest) (*model.ExportIncidentsResponse, error) {
	requestDef := GenReqDefForExportIncidents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportIncidentsResponse), nil
	}
}

// ExportIncidentsInvoker 导出事件
func (c *SecMasterClient) ExportIncidentsInvoker(request *model.ExportIncidentsRequest) *ExportIncidentsInvoker {
	requestDef := GenReqDefForExportIncidents()
	return &ExportIncidentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportIndicators 导出情报
//
// 导出情报,若字段是object类型，则将整个子对象的内容导出
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ExportIndicators(request *model.ExportIndicatorsRequest) (*model.ExportIndicatorsResponse, error) {
	requestDef := GenReqDefForExportIndicators()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportIndicatorsResponse), nil
	}
}

// ExportIndicatorsInvoker 导出情报
func (c *SecMasterClient) ExportIndicatorsInvoker(request *model.ExportIndicatorsRequest) *ExportIndicatorsInvoker {
	requestDef := GenReqDefForExportIndicators()
	return &ExportIndicatorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportPlaybook 导出剧本信息
//
// 导出剧本信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ExportPlaybook(request *model.ExportPlaybookRequest) (*model.ExportPlaybookResponse, error) {
	requestDef := GenReqDefForExportPlaybook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportPlaybookResponse), nil
	}
}

// ExportPlaybookInvoker 导出剧本信息
func (c *SecMasterClient) ExportPlaybookInvoker(request *model.ExportPlaybookRequest) *ExportPlaybookInvoker {
	requestDef := GenReqDefForExportPlaybook()
	return &ExportPlaybookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportResources 导出资产列表
//
// 导出资产列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ExportResources(request *model.ExportResourcesRequest) (*model.ExportResourcesResponse, error) {
	requestDef := GenReqDefForExportResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportResourcesResponse), nil
	}
}

// ExportResourcesInvoker 导出资产列表
func (c *SecMasterClient) ExportResourcesInvoker(request *model.ExportResourcesRequest) *ExportResourcesInvoker {
	requestDef := GenReqDefForExportResources()
	return &ExportResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportVulnerabilities 导出漏洞
//
// 导出详细漏洞信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ExportVulnerabilities(request *model.ExportVulnerabilitiesRequest) (*model.ExportVulnerabilitiesResponse, error) {
	requestDef := GenReqDefForExportVulnerabilities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportVulnerabilitiesResponse), nil
	}
}

// ExportVulnerabilitiesInvoker 导出漏洞
func (c *SecMasterClient) ExportVulnerabilitiesInvoker(request *model.ExportVulnerabilitiesRequest) *ExportVulnerabilitiesInvoker {
	requestDef := GenReqDefForExportVulnerabilities()
	return &ExportVulnerabilitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// HandleShipperAuthorization 授权处理
//
// 授权处理
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) HandleShipperAuthorization(request *model.HandleShipperAuthorizationRequest) (*model.HandleShipperAuthorizationResponse, error) {
	requestDef := GenReqDefForHandleShipperAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.HandleShipperAuthorizationResponse), nil
	}
}

// HandleShipperAuthorizationInvoker 授权处理
func (c *SecMasterClient) HandleShipperAuthorizationInvoker(request *model.HandleShipperAuthorizationRequest) *HandleShipperAuthorizationInvoker {
	requestDef := GenReqDefForHandleShipperAuthorization()
	return &HandleShipperAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportAlerts 导入告警
//
// 批量导入告警，根据template下载的模板填写告警，告警的所有字段根据是否必填来限制，参照:安全事件数据对象定义中的类型和说明来确定, 后台实现时需要校验
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ImportAlerts(request *model.ImportAlertsRequest) (*model.ImportAlertsResponse, error) {
	requestDef := GenReqDefForImportAlerts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportAlertsResponse), nil
	}
}

// ImportAlertsInvoker 导入告警
func (c *SecMasterClient) ImportAlertsInvoker(request *model.ImportAlertsRequest) *ImportAlertsInvoker {
	requestDef := GenReqDefForImportAlerts()
	return &ImportAlertsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportAopworkflow 导入流程信息
//
// 导入流程信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ImportAopworkflow(request *model.ImportAopworkflowRequest) (*model.ImportAopworkflowResponse, error) {
	requestDef := GenReqDefForImportAopworkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportAopworkflowResponse), nil
	}
}

// ImportAopworkflowInvoker 导入流程信息
func (c *SecMasterClient) ImportAopworkflowInvoker(request *model.ImportAopworkflowRequest) *ImportAopworkflowInvoker {
	requestDef := GenReqDefForImportAopworkflow()
	return &ImportAopworkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportCollectorParser 导入采集解析器
//
// 导入采集解析器
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ImportCollectorParser(request *model.ImportCollectorParserRequest) (*model.ImportCollectorParserResponse, error) {
	requestDef := GenReqDefForImportCollectorParser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportCollectorParserResponse), nil
	}
}

// ImportCollectorParserInvoker 导入采集解析器
func (c *SecMasterClient) ImportCollectorParserInvoker(request *model.ImportCollectorParserRequest) *ImportCollectorParserInvoker {
	requestDef := GenReqDefForImportCollectorParser()
	return &ImportCollectorParserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportDataobjects 导入数据对象
//
// 导入数据对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ImportDataobjects(request *model.ImportDataobjectsRequest) (*model.ImportDataobjectsResponse, error) {
	requestDef := GenReqDefForImportDataobjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportDataobjectsResponse), nil
	}
}

// ImportDataobjectsInvoker 导入数据对象
func (c *SecMasterClient) ImportDataobjectsInvoker(request *model.ImportDataobjectsRequest) *ImportDataobjectsInvoker {
	requestDef := GenReqDefForImportDataobjects()
	return &ImportDataobjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportIncidents 导入事件
//
// 导入事件，根据template下载的模板填写事件，事件的所有字段根据是否必填来限制，参照:安全事件数据对象定义中的类型和说明来确定, 后台实现时需要校验
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ImportIncidents(request *model.ImportIncidentsRequest) (*model.ImportIncidentsResponse, error) {
	requestDef := GenReqDefForImportIncidents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportIncidentsResponse), nil
	}
}

// ImportIncidentsInvoker 导入事件
func (c *SecMasterClient) ImportIncidentsInvoker(request *model.ImportIncidentsRequest) *ImportIncidentsInvoker {
	requestDef := GenReqDefForImportIncidents()
	return &ImportIncidentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportIndicators 导入情报
//
// 批量导入情报，根据template下载的模板填写情报，告警的所有字段根据是否必填来限制，参照:安全事件数据对象定义中的类型和说明来确定, 后台实现时需要校验
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ImportIndicators(request *model.ImportIndicatorsRequest) (*model.ImportIndicatorsResponse, error) {
	requestDef := GenReqDefForImportIndicators()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportIndicatorsResponse), nil
	}
}

// ImportIndicatorsInvoker 导入情报
func (c *SecMasterClient) ImportIndicatorsInvoker(request *model.ImportIndicatorsRequest) *ImportIndicatorsInvoker {
	requestDef := GenReqDefForImportIndicators()
	return &ImportIndicatorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportPlaybook 导入剧本信息
//
// 导入剧本信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ImportPlaybook(request *model.ImportPlaybookRequest) (*model.ImportPlaybookResponse, error) {
	requestDef := GenReqDefForImportPlaybook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportPlaybookResponse), nil
	}
}

// ImportPlaybookInvoker 导入剧本信息
func (c *SecMasterClient) ImportPlaybookInvoker(request *model.ImportPlaybookRequest) *ImportPlaybookInvoker {
	requestDef := GenReqDefForImportPlaybook()
	return &ImportPlaybookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportResource 导入资产
//
// 导入资产，根据下载的模板填写资产，资产的所有字段根据是否必填来限制，参照:资产对象定义中的类型和说明来确定, 后台实现时需要校验
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ImportResource(request *model.ImportResourceRequest) (*model.ImportResourceResponse, error) {
	requestDef := GenReqDefForImportResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportResourceResponse), nil
	}
}

// ImportResourceInvoker 导入资产
func (c *SecMasterClient) ImportResourceInvoker(request *model.ImportResourceRequest) *ImportResourceInvoker {
	requestDef := GenReqDefForImportResource()
	return &ImportResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportVulnerabilities 导入漏洞
//
// 批量导入漏洞，根据template下载的模板填写漏洞，漏洞的所有字段根据是否必填来限制，参照:安全事件数据对象定义中的类型和说明来确定, 后台实现时需要校验
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ImportVulnerabilities(request *model.ImportVulnerabilitiesRequest) (*model.ImportVulnerabilitiesResponse, error) {
	requestDef := GenReqDefForImportVulnerabilities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportVulnerabilitiesResponse), nil
	}
}

// ImportVulnerabilitiesInvoker 导入漏洞
func (c *SecMasterClient) ImportVulnerabilitiesInvoker(request *model.ImportVulnerabilitiesRequest) *ImportVulnerabilitiesInvoker {
	requestDef := GenReqDefForImportVulnerabilities()
	return &ImportVulnerabilitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlertRuleMetrics 告警规则总览
//
// 告警规则总览
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListAlertRuleMetrics(request *model.ListAlertRuleMetricsRequest) (*model.ListAlertRuleMetricsResponse, error) {
	requestDef := GenReqDefForListAlertRuleMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlertRuleMetricsResponse), nil
	}
}

// ListAlertRuleMetricsInvoker 告警规则总览
func (c *SecMasterClient) ListAlertRuleMetricsInvoker(request *model.ListAlertRuleMetricsRequest) *ListAlertRuleMetricsInvoker {
	requestDef := GenReqDefForListAlertRuleMetrics()
	return &ListAlertRuleMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlertRuleTemplateMetrics 告警模板总览
//
// 告警模板总览
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListAlertRuleTemplateMetrics(request *model.ListAlertRuleTemplateMetricsRequest) (*model.ListAlertRuleTemplateMetricsResponse, error) {
	requestDef := GenReqDefForListAlertRuleTemplateMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlertRuleTemplateMetricsResponse), nil
	}
}

// ListAlertRuleTemplateMetricsInvoker 告警模板总览
func (c *SecMasterClient) ListAlertRuleTemplateMetricsInvoker(request *model.ListAlertRuleTemplateMetricsRequest) *ListAlertRuleTemplateMetricsInvoker {
	requestDef := GenReqDefForListAlertRuleTemplateMetrics()
	return &ListAlertRuleTemplateMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlertRuleTemplates 列出告警规则模板
//
// 列出告警规则模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListAlertRuleTemplates(request *model.ListAlertRuleTemplatesRequest) (*model.ListAlertRuleTemplatesResponse, error) {
	requestDef := GenReqDefForListAlertRuleTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlertRuleTemplatesResponse), nil
	}
}

// ListAlertRuleTemplatesInvoker 列出告警规则模板
func (c *SecMasterClient) ListAlertRuleTemplatesInvoker(request *model.ListAlertRuleTemplatesRequest) *ListAlertRuleTemplatesInvoker {
	requestDef := GenReqDefForListAlertRuleTemplates()
	return &ListAlertRuleTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlertRules 列出告警规则
//
// 列出告警规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListAlertRules(request *model.ListAlertRulesRequest) (*model.ListAlertRulesResponse, error) {
	requestDef := GenReqDefForListAlertRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlertRulesResponse), nil
	}
}

// ListAlertRulesInvoker 列出告警规则
func (c *SecMasterClient) ListAlertRulesInvoker(request *model.ListAlertRulesRequest) *ListAlertRulesInvoker {
	requestDef := GenReqDefForListAlertRules()
	return &ListAlertRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlerts 搜索告警列表
//
// 搜索告警列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListAlerts(request *model.ListAlertsRequest) (*model.ListAlertsResponse, error) {
	requestDef := GenReqDefForListAlerts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlertsResponse), nil
	}
}

// ListAlertsInvoker 搜索告警列表
func (c *SecMasterClient) ListAlertsInvoker(request *model.ListAlertsRequest) *ListAlertsInvoker {
	requestDef := GenReqDefForListAlerts()
	return &ListAlertsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllSecondCatalogue 获取目录全量列表
//
// 获取目录全量列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListAllSecondCatalogue(request *model.ListAllSecondCatalogueRequest) (*model.ListAllSecondCatalogueResponse, error) {
	requestDef := GenReqDefForListAllSecondCatalogue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllSecondCatalogueResponse), nil
	}
}

// ListAllSecondCatalogueInvoker 获取目录全量列表
func (c *SecMasterClient) ListAllSecondCatalogueInvoker(request *model.ListAllSecondCatalogueRequest) *ListAllSecondCatalogueInvoker {
	requestDef := GenReqDefForListAllSecondCatalogue()
	return &ListAllSecondCatalogueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAopWorkflowInstance 查询流程实例的列表
//
// 查询流程实例的列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListAopWorkflowInstance(request *model.ListAopWorkflowInstanceRequest) (*model.ListAopWorkflowInstanceResponse, error) {
	requestDef := GenReqDefForListAopWorkflowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAopWorkflowInstanceResponse), nil
	}
}

// ListAopWorkflowInstanceInvoker 查询流程实例的列表
func (c *SecMasterClient) ListAopWorkflowInstanceInvoker(request *model.ListAopWorkflowInstanceRequest) *ListAopWorkflowInstanceInvoker {
	requestDef := GenReqDefForListAopWorkflowInstance()
	return &ListAopWorkflowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAopWorkflowVersions 查询流程版本列表
//
// 查询流程版本列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListAopWorkflowVersions(request *model.ListAopWorkflowVersionsRequest) (*model.ListAopWorkflowVersionsResponse, error) {
	requestDef := GenReqDefForListAopWorkflowVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAopWorkflowVersionsResponse), nil
	}
}

// ListAopWorkflowVersionsInvoker 查询流程版本列表
func (c *SecMasterClient) ListAopWorkflowVersionsInvoker(request *model.ListAopWorkflowVersionsRequest) *ListAopWorkflowVersionsInvoker {
	requestDef := GenReqDefForListAopWorkflowVersions()
	return &ListAopWorkflowVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCatalogue 目录列表查询
//
// 目录列表查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListCatalogue(request *model.ListCatalogueRequest) (*model.ListCatalogueResponse, error) {
	requestDef := GenReqDefForListCatalogue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCatalogueResponse), nil
	}
}

// ListCatalogueInvoker 目录列表查询
func (c *SecMasterClient) ListCatalogueInvoker(request *model.ListCatalogueRequest) *ListCatalogueInvoker {
	requestDef := GenReqDefForListCatalogue()
	return &ListCatalogueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCloudLogAlias 列出管道alias
//
// 列出云日志资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListCloudLogAlias(request *model.ListCloudLogAliasRequest) (*model.ListCloudLogAliasResponse, error) {
	requestDef := GenReqDefForListCloudLogAlias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudLogAliasResponse), nil
	}
}

// ListCloudLogAliasInvoker 列出管道alias
func (c *SecMasterClient) ListCloudLogAliasInvoker(request *model.ListCloudLogAliasRequest) *ListCloudLogAliasInvoker {
	requestDef := GenReqDefForListCloudLogAlias()
	return &ListCloudLogAliasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCloudLogManages 列出行管账户信息
//
// 列出云日志资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListCloudLogManages(request *model.ListCloudLogManagesRequest) (*model.ListCloudLogManagesResponse, error) {
	requestDef := GenReqDefForListCloudLogManages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudLogManagesResponse), nil
	}
}

// ListCloudLogManagesInvoker 列出行管账户信息
func (c *SecMasterClient) ListCloudLogManagesInvoker(request *model.ListCloudLogManagesRequest) *ListCloudLogManagesInvoker {
	requestDef := GenReqDefForListCloudLogManages()
	return &ListCloudLogManagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCloudLogPlatform 列出平台行管账户信息
//
// 列出云日志资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListCloudLogPlatform(request *model.ListCloudLogPlatformRequest) (*model.ListCloudLogPlatformResponse, error) {
	requestDef := GenReqDefForListCloudLogPlatform()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudLogPlatformResponse), nil
	}
}

// ListCloudLogPlatformInvoker 列出平台行管账户信息
func (c *SecMasterClient) ListCloudLogPlatformInvoker(request *model.ListCloudLogPlatformRequest) *ListCloudLogPlatformInvoker {
	requestDef := GenReqDefForListCloudLogPlatform()
	return &ListCloudLogPlatformInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCloudLogResources 列出云日志资源
//
// 列出云日志资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListCloudLogResources(request *model.ListCloudLogResourcesRequest) (*model.ListCloudLogResourcesResponse, error) {
	requestDef := GenReqDefForListCloudLogResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudLogResourcesResponse), nil
	}
}

// ListCloudLogResourcesInvoker 列出云日志资源
func (c *SecMasterClient) ListCloudLogResourcesInvoker(request *model.ListCloudLogResourcesRequest) *ListCloudLogResourcesInvoker {
	requestDef := GenReqDefForListCloudLogResources()
	return &ListCloudLogResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCollectConfig 获取云服务采集配置
//
// 获取云服务采集配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListCollectConfig(request *model.ListCollectConfigRequest) (*model.ListCollectConfigResponse, error) {
	requestDef := GenReqDefForListCollectConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCollectConfigResponse), nil
	}
}

// ListCollectConfigInvoker 获取云服务采集配置
func (c *SecMasterClient) ListCollectConfigInvoker(request *model.ListCollectConfigRequest) *ListCollectConfigInvoker {
	requestDef := GenReqDefForListCollectConfig()
	return &ListCollectConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCollectorChannelGroup 列出采集通道分组
//
// 列出采集通道分组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListCollectorChannelGroup(request *model.ListCollectorChannelGroupRequest) (*model.ListCollectorChannelGroupResponse, error) {
	requestDef := GenReqDefForListCollectorChannelGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCollectorChannelGroupResponse), nil
	}
}

// ListCollectorChannelGroupInvoker 列出采集通道分组
func (c *SecMasterClient) ListCollectorChannelGroupInvoker(request *model.ListCollectorChannelGroupRequest) *ListCollectorChannelGroupInvoker {
	requestDef := GenReqDefForListCollectorChannelGroup()
	return &ListCollectorChannelGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCollectorChannels 列出采集通道列表
//
// 列出采集通道列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListCollectorChannels(request *model.ListCollectorChannelsRequest) (*model.ListCollectorChannelsResponse, error) {
	requestDef := GenReqDefForListCollectorChannels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCollectorChannelsResponse), nil
	}
}

// ListCollectorChannelsInvoker 列出采集通道列表
func (c *SecMasterClient) ListCollectorChannelsInvoker(request *model.ListCollectorChannelsRequest) *ListCollectorChannelsInvoker {
	requestDef := GenReqDefForListCollectorChannels()
	return &ListCollectorChannelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCollectorConnections 列表采集连接
//
// 列表采集连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListCollectorConnections(request *model.ListCollectorConnectionsRequest) (*model.ListCollectorConnectionsResponse, error) {
	requestDef := GenReqDefForListCollectorConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCollectorConnectionsResponse), nil
	}
}

// ListCollectorConnectionsInvoker 列表采集连接
func (c *SecMasterClient) ListCollectorConnectionsInvoker(request *model.ListCollectorConnectionsRequest) *ListCollectorConnectionsInvoker {
	requestDef := GenReqDefForListCollectorConnections()
	return &ListCollectorConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCollectorInstances 列出采集通道实例
//
// 列出采集通道实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListCollectorInstances(request *model.ListCollectorInstancesRequest) (*model.ListCollectorInstancesResponse, error) {
	requestDef := GenReqDefForListCollectorInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCollectorInstancesResponse), nil
	}
}

// ListCollectorInstancesInvoker 列出采集通道实例
func (c *SecMasterClient) ListCollectorInstancesInvoker(request *model.ListCollectorInstancesRequest) *ListCollectorInstancesInvoker {
	requestDef := GenReqDefForListCollectorInstances()
	return &ListCollectorInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCollectorModuleRestrictions 列出采集模块规则
//
// 列出采集模块规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListCollectorModuleRestrictions(request *model.ListCollectorModuleRestrictionsRequest) (*model.ListCollectorModuleRestrictionsResponse, error) {
	requestDef := GenReqDefForListCollectorModuleRestrictions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCollectorModuleRestrictionsResponse), nil
	}
}

// ListCollectorModuleRestrictionsInvoker 列出采集模块规则
func (c *SecMasterClient) ListCollectorModuleRestrictionsInvoker(request *model.ListCollectorModuleRestrictionsRequest) *ListCollectorModuleRestrictionsInvoker {
	requestDef := GenReqDefForListCollectorModuleRestrictions()
	return &ListCollectorModuleRestrictionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCollectorModuleTemplate 列出采集模块
//
// 列出采集模块
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListCollectorModuleTemplate(request *model.ListCollectorModuleTemplateRequest) (*model.ListCollectorModuleTemplateResponse, error) {
	requestDef := GenReqDefForListCollectorModuleTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCollectorModuleTemplateResponse), nil
	}
}

// ListCollectorModuleTemplateInvoker 列出采集模块
func (c *SecMasterClient) ListCollectorModuleTemplateInvoker(request *model.ListCollectorModuleTemplateRequest) *ListCollectorModuleTemplateInvoker {
	requestDef := GenReqDefForListCollectorModuleTemplate()
	return &ListCollectorModuleTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCollectorNode 列出采集节点
//
// 列出采集节点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListCollectorNode(request *model.ListCollectorNodeRequest) (*model.ListCollectorNodeResponse, error) {
	requestDef := GenReqDefForListCollectorNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCollectorNodeResponse), nil
	}
}

// ListCollectorNodeInvoker 列出采集节点
func (c *SecMasterClient) ListCollectorNodeInvoker(request *model.ListCollectorNodeRequest) *ListCollectorNodeInvoker {
	requestDef := GenReqDefForListCollectorNode()
	return &ListCollectorNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCollectorParserTemplate 列出采集解析器模板
//
// 列出采集解析器模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListCollectorParserTemplate(request *model.ListCollectorParserTemplateRequest) (*model.ListCollectorParserTemplateResponse, error) {
	requestDef := GenReqDefForListCollectorParserTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCollectorParserTemplateResponse), nil
	}
}

// ListCollectorParserTemplateInvoker 列出采集解析器模板
func (c *SecMasterClient) ListCollectorParserTemplateInvoker(request *model.ListCollectorParserTemplateRequest) *ListCollectorParserTemplateInvoker {
	requestDef := GenReqDefForListCollectorParserTemplate()
	return &ListCollectorParserTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCollectorParsers 列出采集解析器
//
// 列出采集解析器
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListCollectorParsers(request *model.ListCollectorParsersRequest) (*model.ListCollectorParsersResponse, error) {
	requestDef := GenReqDefForListCollectorParsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCollectorParsersResponse), nil
	}
}

// ListCollectorParsersInvoker 列出采集解析器
func (c *SecMasterClient) ListCollectorParsersInvoker(request *model.ListCollectorParsersRequest) *ListCollectorParsersInvoker {
	requestDef := GenReqDefForListCollectorParsers()
	return &ListCollectorParsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListComponentActions 查询插件Action列表
//
// 查询插件的函数、连接器和管理器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListComponentActions(request *model.ListComponentActionsRequest) (*model.ListComponentActionsResponse, error) {
	requestDef := GenReqDefForListComponentActions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListComponentActionsResponse), nil
	}
}

// ListComponentActionsInvoker 查询插件Action列表
func (c *SecMasterClient) ListComponentActionsInvoker(request *model.ListComponentActionsRequest) *ListComponentActionsInvoker {
	requestDef := GenReqDefForListComponentActions()
	return &ListComponentActionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListComponentConfiguration 列出组件配置
//
// 列出组件配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListComponentConfiguration(request *model.ListComponentConfigurationRequest) (*model.ListComponentConfigurationResponse, error) {
	requestDef := GenReqDefForListComponentConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListComponentConfigurationResponse), nil
	}
}

// ListComponentConfigurationInvoker 列出组件配置
func (c *SecMasterClient) ListComponentConfigurationInvoker(request *model.ListComponentConfigurationRequest) *ListComponentConfigurationInvoker {
	requestDef := GenReqDefForListComponentConfiguration()
	return &ListComponentConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListComponentTemplate 列出组件模板
//
// 列出组件模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListComponentTemplate(request *model.ListComponentTemplateRequest) (*model.ListComponentTemplateResponse, error) {
	requestDef := GenReqDefForListComponentTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListComponentTemplateResponse), nil
	}
}

// ListComponentTemplateInvoker 列出组件模板
func (c *SecMasterClient) ListComponentTemplateInvoker(request *model.ListComponentTemplateRequest) *ListComponentTemplateInvoker {
	requestDef := GenReqDefForListComponentTemplate()
	return &ListComponentTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListComponentTemplates 查询插件配置模板列表
//
// 查询在配置流程时的插件配置模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListComponentTemplates(request *model.ListComponentTemplatesRequest) (*model.ListComponentTemplatesResponse, error) {
	requestDef := GenReqDefForListComponentTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListComponentTemplatesResponse), nil
	}
}

// ListComponentTemplatesInvoker 查询插件配置模板列表
func (c *SecMasterClient) ListComponentTemplatesInvoker(request *model.ListComponentTemplatesRequest) *ListComponentTemplatesInvoker {
	requestDef := GenReqDefForListComponentTemplates()
	return &ListComponentTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListComponents 查询插件定义列表
//
// 查询插件定义列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListComponents(request *model.ListComponentsRequest) (*model.ListComponentsResponse, error) {
	requestDef := GenReqDefForListComponents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListComponentsResponse), nil
	}
}

// ListComponentsInvoker 查询插件定义列表
func (c *SecMasterClient) ListComponentsInvoker(request *model.ListComponentsRequest) *ListComponentsInvoker {
	requestDef := GenReqDefForListComponents()
	return &ListComponentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigurationDictionaries 获取字典信息
//
// 获取字典信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListConfigurationDictionaries(request *model.ListConfigurationDictionariesRequest) (*model.ListConfigurationDictionariesResponse, error) {
	requestDef := GenReqDefForListConfigurationDictionaries()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigurationDictionariesResponse), nil
	}
}

// ListConfigurationDictionariesInvoker 获取字典信息
func (c *SecMasterClient) ListConfigurationDictionariesInvoker(request *model.ListConfigurationDictionariesRequest) *ListConfigurationDictionariesInvoker {
	requestDef := GenReqDefForListConfigurationDictionaries()
	return &ListConfigurationDictionariesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConnections 操作连接列表查询接口
//
// 操作连接列表查询接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListConnections(request *model.ListConnectionsRequest) (*model.ListConnectionsResponse, error) {
	requestDef := GenReqDefForListConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConnectionsResponse), nil
	}
}

// ListConnectionsInvoker 操作连接列表查询接口
func (c *SecMasterClient) ListConnectionsInvoker(request *model.ListConnectionsRequest) *ListConnectionsInvoker {
	requestDef := GenReqDefForListConnections()
	return &ListConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataclass 查询数据类列表
//
// 查询数据类列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListDataclass(request *model.ListDataclassRequest) (*model.ListDataclassResponse, error) {
	requestDef := GenReqDefForListDataclass()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataclassResponse), nil
	}
}

// ListDataclassInvoker 查询数据类列表
func (c *SecMasterClient) ListDataclassInvoker(request *model.ListDataclassRequest) *ListDataclassInvoker {
	requestDef := GenReqDefForListDataclass()
	return &ListDataclassInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataclassFields 查询字段列表
//
// 查询字段列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListDataclassFields(request *model.ListDataclassFieldsRequest) (*model.ListDataclassFieldsResponse, error) {
	requestDef := GenReqDefForListDataclassFields()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataclassFieldsResponse), nil
	}
}

// ListDataclassFieldsInvoker 查询字段列表
func (c *SecMasterClient) ListDataclassFieldsInvoker(request *model.ListDataclassFieldsRequest) *ListDataclassFieldsInvoker {
	requestDef := GenReqDefForListDataclassFields()
	return &ListDataclassFieldsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataobjectRelations 查询关联数据对象列表
//
// 查询关联数据对象列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListDataobjectRelations(request *model.ListDataobjectRelationsRequest) (*model.ListDataobjectRelationsResponse, error) {
	requestDef := GenReqDefForListDataobjectRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataobjectRelationsResponse), nil
	}
}

// ListDataobjectRelationsInvoker 查询关联数据对象列表
func (c *SecMasterClient) ListDataobjectRelationsInvoker(request *model.ListDataobjectRelationsRequest) *ListDataobjectRelationsInvoker {
	requestDef := GenReqDefForListDataobjectRelations()
	return &ListDataobjectRelationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataobjects 列出所有数据对象
//
// 列出所有与数据对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListDataobjects(request *model.ListDataobjectsRequest) (*model.ListDataobjectsResponse, error) {
	requestDef := GenReqDefForListDataobjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataobjectsResponse), nil
	}
}

// ListDataobjectsInvoker 列出所有数据对象
func (c *SecMasterClient) ListDataobjectsInvoker(request *model.ListDataobjectsRequest) *ListDataobjectsInvoker {
	requestDef := GenReqDefForListDataobjects()
	return &ListDataobjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatapanelObjects 查询数据对象列表
//
// 数据面查询数据类纳管的数据对象列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListDatapanelObjects(request *model.ListDatapanelObjectsRequest) (*model.ListDatapanelObjectsResponse, error) {
	requestDef := GenReqDefForListDatapanelObjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatapanelObjectsResponse), nil
	}
}

// ListDatapanelObjectsInvoker 查询数据对象列表
func (c *SecMasterClient) ListDatapanelObjectsInvoker(request *model.ListDatapanelObjectsRequest) *ListDatapanelObjectsInvoker {
	requestDef := GenReqDefForListDatapanelObjects()
	return &ListDatapanelObjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataspaces 获取数据空间列表
//
// 获取数据空间列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListDataspaces(request *model.ListDataspacesRequest) (*model.ListDataspacesResponse, error) {
	requestDef := GenReqDefForListDataspaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataspacesResponse), nil
	}
}

// ListDataspacesInvoker 获取数据空间列表
func (c *SecMasterClient) ListDataspacesInvoker(request *model.ListDataspacesRequest) *ListDataspacesInvoker {
	requestDef := GenReqDefForListDataspaces()
	return &ListDataspacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHistoryComponentConfiguration 获取组件历史版本的配置数据
//
// 获取组件历史版本的配置数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListHistoryComponentConfiguration(request *model.ListHistoryComponentConfigurationRequest) (*model.ListHistoryComponentConfigurationResponse, error) {
	requestDef := GenReqDefForListHistoryComponentConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHistoryComponentConfigurationResponse), nil
	}
}

// ListHistoryComponentConfigurationInvoker 获取组件历史版本的配置数据
func (c *SecMasterClient) ListHistoryComponentConfigurationInvoker(request *model.ListHistoryComponentConfigurationRequest) *ListHistoryComponentConfigurationInvoker {
	requestDef := GenReqDefForListHistoryComponentConfiguration()
	return &ListHistoryComponentConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIncidents 搜索事件列表
//
// 搜索事件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListIncidents(request *model.ListIncidentsRequest) (*model.ListIncidentsResponse, error) {
	requestDef := GenReqDefForListIncidents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIncidentsResponse), nil
	}
}

// ListIncidentsInvoker 搜索事件列表
func (c *SecMasterClient) ListIncidentsInvoker(request *model.ListIncidentsRequest) *ListIncidentsInvoker {
	requestDef := GenReqDefForListIncidents()
	return &ListIncidentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIndicators 查询威胁情报列表
//
// 查询威胁情报列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListIndicators(request *model.ListIndicatorsRequest) (*model.ListIndicatorsResponse, error) {
	requestDef := GenReqDefForListIndicators()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIndicatorsResponse), nil
	}
}

// ListIndicatorsInvoker 查询威胁情报列表
func (c *SecMasterClient) ListIndicatorsInvoker(request *model.ListIndicatorsRequest) *ListIndicatorsInvoker {
	requestDef := GenReqDefForListIndicators()
	return &ListIndicatorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstallation 查询安装脚本列表
//
// 查询安装脚本列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListInstallation(request *model.ListInstallationRequest) (*model.ListInstallationResponse, error) {
	requestDef := GenReqDefForListInstallation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstallationResponse), nil
	}
}

// ListInstallationInvoker 查询安装脚本列表
func (c *SecMasterClient) ListInstallationInvoker(request *model.ListInstallationRequest) *ListInstallationInvoker {
	requestDef := GenReqDefForListInstallation()
	return &ListInstallationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIsapComponents 列出组件
//
// 列出组件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListIsapComponents(request *model.ListIsapComponentsRequest) (*model.ListIsapComponentsResponse, error) {
	requestDef := GenReqDefForListIsapComponents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIsapComponentsResponse), nil
	}
}

// ListIsapComponentsInvoker 列出组件
func (c *SecMasterClient) ListIsapComponentsInvoker(request *model.ListIsapComponentsRequest) *ListIsapComponentsInvoker {
	requestDef := GenReqDefForListIsapComponents()
	return &ListIsapComponentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLayout 查询布局列表
//
// 查询布局列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListLayout(request *model.ListLayoutRequest) (*model.ListLayoutResponse, error) {
	requestDef := GenReqDefForListLayout()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLayoutResponse), nil
	}
}

// ListLayoutInvoker 查询布局列表
func (c *SecMasterClient) ListLayoutInvoker(request *model.ListLayoutRequest) *ListLayoutInvoker {
	requestDef := GenReqDefForListLayout()
	return &ListLayoutInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLayoutWizards 查询布局页面列表
//
// 查询所有页面
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListLayoutWizards(request *model.ListLayoutWizardsRequest) (*model.ListLayoutWizardsResponse, error) {
	requestDef := GenReqDefForListLayoutWizards()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLayoutWizardsResponse), nil
	}
}

// ListLayoutWizardsInvoker 查询布局页面列表
func (c *SecMasterClient) ListLayoutWizardsInvoker(request *model.ListLayoutWizardsRequest) *ListLayoutWizardsInvoker {
	requestDef := GenReqDefForListLayoutWizards()
	return &ListLayoutWizardsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMetrics 获取指标列表
//
// 获取指标列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListMetrics(request *model.ListMetricsRequest) (*model.ListMetricsResponse, error) {
	requestDef := GenReqDefForListMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetricsResponse), nil
	}
}

// ListMetricsInvoker 获取指标列表
func (c *SecMasterClient) ListMetricsInvoker(request *model.ListMetricsRequest) *ListMetricsInvoker {
	requestDef := GenReqDefForListMetrics()
	return &ListMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListModules 查询模块列表
//
// 查询所有模块
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListModules(request *model.ListModulesRequest) (*model.ListModulesResponse, error) {
	requestDef := GenReqDefForListModules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListModulesResponse), nil
	}
}

// ListModulesInvoker 查询模块列表
func (c *SecMasterClient) ListModulesInvoker(request *model.ListModulesRequest) *ListModulesInvoker {
	requestDef := GenReqDefForListModules()
	return &ListModulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNodes 通过节点id查询组件信息
//
// 查询节点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListNodes(request *model.ListNodesRequest) (*model.ListNodesResponse, error) {
	requestDef := GenReqDefForListNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNodesResponse), nil
	}
}

// ListNodesInvoker 通过节点id查询组件信息
func (c *SecMasterClient) ListNodesInvoker(request *model.ListNodesRequest) *ListNodesInvoker {
	requestDef := GenReqDefForListNodes()
	return &ListNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNotes 搜索评论列表
//
// 搜索评论列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListNotes(request *model.ListNotesRequest) (*model.ListNotesResponse, error) {
	requestDef := GenReqDefForListNotes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotesResponse), nil
	}
}

// ListNotesInvoker 搜索评论列表
func (c *SecMasterClient) ListNotesInvoker(request *model.ListNotesRequest) *ListNotesInvoker {
	requestDef := GenReqDefForListNotes()
	return &ListNotesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPipes 获取数据管道列表
//
// 获取数据管道列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListPipes(request *model.ListPipesRequest) (*model.ListPipesResponse, error) {
	requestDef := GenReqDefForListPipes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPipesResponse), nil
	}
}

// ListPipesInvoker 获取数据管道列表
func (c *SecMasterClient) ListPipesInvoker(request *model.ListPipesRequest) *ListPipesInvoker {
	requestDef := GenReqDefForListPipes()
	return &ListPipesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlaybookActions 查询剧本动作
//
// 查询剧本动作列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListPlaybookActions(request *model.ListPlaybookActionsRequest) (*model.ListPlaybookActionsResponse, error) {
	requestDef := GenReqDefForListPlaybookActions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlaybookActionsResponse), nil
	}
}

// ListPlaybookActionsInvoker 查询剧本动作
func (c *SecMasterClient) ListPlaybookActionsInvoker(request *model.ListPlaybookActionsRequest) *ListPlaybookActionsInvoker {
	requestDef := GenReqDefForListPlaybookActions()
	return &ListPlaybookActionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlaybookApproves 查询剧本审核结果
//
// 查询剧本审核结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListPlaybookApproves(request *model.ListPlaybookApprovesRequest) (*model.ListPlaybookApprovesResponse, error) {
	requestDef := GenReqDefForListPlaybookApproves()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlaybookApprovesResponse), nil
	}
}

// ListPlaybookApprovesInvoker 查询剧本审核结果
func (c *SecMasterClient) ListPlaybookApprovesInvoker(request *model.ListPlaybookApprovesRequest) *ListPlaybookApprovesInvoker {
	requestDef := GenReqDefForListPlaybookApproves()
	return &ListPlaybookApprovesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlaybookAuditLogs 查询剧本实例审计日志
//
// 查询剧本实例审计日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListPlaybookAuditLogs(request *model.ListPlaybookAuditLogsRequest) (*model.ListPlaybookAuditLogsResponse, error) {
	requestDef := GenReqDefForListPlaybookAuditLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlaybookAuditLogsResponse), nil
	}
}

// ListPlaybookAuditLogsInvoker 查询剧本实例审计日志
func (c *SecMasterClient) ListPlaybookAuditLogsInvoker(request *model.ListPlaybookAuditLogsRequest) *ListPlaybookAuditLogsInvoker {
	requestDef := GenReqDefForListPlaybookAuditLogs()
	return &ListPlaybookAuditLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlaybookInstances 查询剧本实例列表
//
// 查询剧本实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListPlaybookInstances(request *model.ListPlaybookInstancesRequest) (*model.ListPlaybookInstancesResponse, error) {
	requestDef := GenReqDefForListPlaybookInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlaybookInstancesResponse), nil
	}
}

// ListPlaybookInstancesInvoker 查询剧本实例列表
func (c *SecMasterClient) ListPlaybookInstancesInvoker(request *model.ListPlaybookInstancesRequest) *ListPlaybookInstancesInvoker {
	requestDef := GenReqDefForListPlaybookInstances()
	return &ListPlaybookInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlaybookVersions 查询剧本版本列表
//
// 查询剧本版本列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListPlaybookVersions(request *model.ListPlaybookVersionsRequest) (*model.ListPlaybookVersionsResponse, error) {
	requestDef := GenReqDefForListPlaybookVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlaybookVersionsResponse), nil
	}
}

// ListPlaybookVersionsInvoker 查询剧本版本列表
func (c *SecMasterClient) ListPlaybookVersionsInvoker(request *model.ListPlaybookVersionsRequest) *ListPlaybookVersionsInvoker {
	requestDef := GenReqDefForListPlaybookVersions()
	return &ListPlaybookVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlaybooks 查询剧本列表
//
// 查询剧本列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListPlaybooks(request *model.ListPlaybooksRequest) (*model.ListPlaybooksResponse, error) {
	requestDef := GenReqDefForListPlaybooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlaybooksResponse), nil
	}
}

// ListPlaybooksInvoker 查询剧本列表
func (c *SecMasterClient) ListPlaybooksInvoker(request *model.ListPlaybooksRequest) *ListPlaybooksInvoker {
	requestDef := GenReqDefForListPlaybooks()
	return &ListPlaybooksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectTag 查询项目标签
//
// 查询租户在指定Project中实例类型的所有资源标签集合
// 标签管理服务需要能够列出当前租户全部已使用的资源标签集合，为各服务Console打资源标签和过滤实例时提供标签联想功能
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListProjectTag(request *model.ListProjectTagRequest) (*model.ListProjectTagResponse, error) {
	requestDef := GenReqDefForListProjectTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagResponse), nil
	}
}

// ListProjectTagInvoker 查询项目标签
func (c *SecMasterClient) ListProjectTagInvoker(request *model.ListProjectTagRequest) *ListProjectTagInvoker {
	requestDef := GenReqDefForListProjectTag()
	return &ListProjectTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRecipientsStatus 查询收件人状态
//
// 查询收件人状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListRecipientsStatus(request *model.ListRecipientsStatusRequest) (*model.ListRecipientsStatusResponse, error) {
	requestDef := GenReqDefForListRecipientsStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecipientsStatusResponse), nil
	}
}

// ListRecipientsStatusInvoker 查询收件人状态
func (c *SecMasterClient) ListRecipientsStatusInvoker(request *model.ListRecipientsStatusRequest) *ListRecipientsStatusInvoker {
	requestDef := GenReqDefForListRecipientsStatus()
	return &ListRecipientsStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRegions 查询区域列表
//
// 查询区域列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListRegions(request *model.ListRegionsRequest) (*model.ListRegionsResponse, error) {
	requestDef := GenReqDefForListRegions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRegionsResponse), nil
	}
}

// ListRegionsInvoker 查询区域列表
func (c *SecMasterClient) ListRegionsInvoker(request *model.ListRegionsRequest) *ListRegionsInvoker {
	requestDef := GenReqDefForListRegions()
	return &ListRegionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListReports 获取报告列表
//
// 获取报告列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListReports(request *model.ListReportsRequest) (*model.ListReportsResponse, error) {
	requestDef := GenReqDefForListReports()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListReportsResponse), nil
	}
}

// ListReportsInvoker 获取报告列表
func (c *SecMasterClient) ListReportsInvoker(request *model.ListReportsRequest) *ListReportsInvoker {
	requestDef := GenReqDefForListReports()
	return &ListReportsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceInstance 查询资源实例列表
//
// 使用标签过滤实例，查询资源实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListResourceInstance(request *model.ListResourceInstanceRequest) (*model.ListResourceInstanceResponse, error) {
	requestDef := GenReqDefForListResourceInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceInstanceResponse), nil
	}
}

// ListResourceInstanceInvoker 查询资源实例列表
func (c *SecMasterClient) ListResourceInstanceInvoker(request *model.ListResourceInstanceRequest) *ListResourceInstanceInvoker {
	requestDef := GenReqDefForListResourceInstance()
	return &ListResourceInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceTag 查询资源标签
//
// 查询指定实例的标签信息
// 标签管理服务需要使用该接口查询指定实例的全部标签数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListResourceTag(request *model.ListResourceTagRequest) (*model.ListResourceTagResponse, error) {
	requestDef := GenReqDefForListResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceTagResponse), nil
	}
}

// ListResourceTagInvoker 查询资源标签
func (c *SecMasterClient) ListResourceTagInvoker(request *model.ListResourceTagRequest) *ListResourceTagInvoker {
	requestDef := GenReqDefForListResourceTag()
	return &ListResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResources 搜索资产列表
//
// 搜索资产列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListResources(request *model.ListResourcesRequest) (*model.ListResourcesResponse, error) {
	requestDef := GenReqDefForListResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourcesResponse), nil
	}
}

// ListResourcesInvoker 搜索资产列表
func (c *SecMasterClient) ListResourcesInvoker(request *model.ListResourcesRequest) *ListResourcesInvoker {
	requestDef := GenReqDefForListResources()
	return &ListResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRunningNodes 列出运行中节点
//
// 列出运行中节点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListRunningNodes(request *model.ListRunningNodesRequest) (*model.ListRunningNodesResponse, error) {
	requestDef := GenReqDefForListRunningNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRunningNodesResponse), nil
	}
}

// ListRunningNodesInvoker 列出运行中节点
func (c *SecMasterClient) ListRunningNodesInvoker(request *model.ListRunningNodesRequest) *ListRunningNodesInvoker {
	requestDef := GenReqDefForListRunningNodes()
	return &ListRunningNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSearchConditions 获取检索条件列表
//
// 获取检索条件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListSearchConditions(request *model.ListSearchConditionsRequest) (*model.ListSearchConditionsResponse, error) {
	requestDef := GenReqDefForListSearchConditions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSearchConditionsResponse), nil
	}
}

// ListSearchConditionsInvoker 获取检索条件列表
func (c *SecMasterClient) ListSearchConditionsInvoker(request *model.ListSearchConditionsRequest) *ListSearchConditionsInvoker {
	requestDef := GenReqDefForListSearchConditions()
	return &ListSearchConditionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSearchHistograms 获取数据分布直方图
//
// 获取数据分布直方图
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListSearchHistograms(request *model.ListSearchHistogramsRequest) (*model.ListSearchHistogramsResponse, error) {
	requestDef := GenReqDefForListSearchHistograms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSearchHistogramsResponse), nil
	}
}

// ListSearchHistogramsInvoker 获取数据分布直方图
func (c *SecMasterClient) ListSearchHistogramsInvoker(request *model.ListSearchHistogramsRequest) *ListSearchHistogramsInvoker {
	requestDef := GenReqDefForListSearchHistograms()
	return &ListSearchHistogramsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSearchLogs 获取检索数据
//
// 获取检索数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListSearchLogs(request *model.ListSearchLogsRequest) (*model.ListSearchLogsResponse, error) {
	requestDef := GenReqDefForListSearchLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSearchLogsResponse), nil
	}
}

// ListSearchLogsInvoker 获取检索数据
func (c *SecMasterClient) ListSearchLogsInvoker(request *model.ListSearchLogsRequest) *ListSearchLogsInvoker {
	requestDef := GenReqDefForListSearchLogs()
	return &ListSearchLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServiceAgency 查看委托信息
//
// check service linked agency
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListServiceAgency(request *model.ListServiceAgencyRequest) (*model.ListServiceAgencyResponse, error) {
	requestDef := GenReqDefForListServiceAgency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceAgencyResponse), nil
	}
}

// ListServiceAgencyInvoker 查看委托信息
func (c *SecMasterClient) ListServiceAgencyInvoker(request *model.ListServiceAgencyRequest) *ListServiceAgencyInvoker {
	requestDef := GenReqDefForListServiceAgency()
	return &ListServiceAgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListShipperAuthorizations 获取投递授权信息列表
//
// 获取投递授权信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListShipperAuthorizations(request *model.ListShipperAuthorizationsRequest) (*model.ListShipperAuthorizationsResponse, error) {
	requestDef := GenReqDefForListShipperAuthorizations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListShipperAuthorizationsResponse), nil
	}
}

// ListShipperAuthorizationsInvoker 获取投递授权信息列表
func (c *SecMasterClient) ListShipperAuthorizationsInvoker(request *model.ListShipperAuthorizationsRequest) *ListShipperAuthorizationsInvoker {
	requestDef := GenReqDefForListShipperAuthorizations()
	return &ListShipperAuthorizationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListShippers 查询投递信息
//
// 查询投递信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListShippers(request *model.ListShippersRequest) (*model.ListShippersResponse, error) {
	requestDef := GenReqDefForListShippers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListShippersResponse), nil
	}
}

// ListShippersInvoker 查询投递信息
func (c *SecMasterClient) ListShippersInvoker(request *model.ListShippersRequest) *ListShippersInvoker {
	requestDef := GenReqDefForListShippers()
	return &ListShippersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubscriptionProduct 查询当前站点支持的商品清单
//
// 查询当前站点SecMaster支持的商品清单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListSubscriptionProduct(request *model.ListSubscriptionProductRequest) (*model.ListSubscriptionProductResponse, error) {
	requestDef := GenReqDefForListSubscriptionProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubscriptionProductResponse), nil
	}
}

// ListSubscriptionProductInvoker 查询当前站点支持的商品清单
func (c *SecMasterClient) ListSubscriptionProductInvoker(request *model.ListSubscriptionProductRequest) *ListSubscriptionProductInvoker {
	requestDef := GenReqDefForListSubscriptionProduct()
	return &ListSubscriptionProductInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVpcEndpointService 列出VPC终端节点服务
//
// 列出VPC终端节点服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListVpcEndpointService(request *model.ListVpcEndpointServiceRequest) (*model.ListVpcEndpointServiceResponse, error) {
	requestDef := GenReqDefForListVpcEndpointService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcEndpointServiceResponse), nil
	}
}

// ListVpcEndpointServiceInvoker 列出VPC终端节点服务
func (c *SecMasterClient) ListVpcEndpointServiceInvoker(request *model.ListVpcEndpointServiceRequest) *ListVpcEndpointServiceInvoker {
	requestDef := GenReqDefForListVpcEndpointService()
	return &ListVpcEndpointServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVulnerabilities 查询漏洞列表
//
// 查询漏洞列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListVulnerabilities(request *model.ListVulnerabilitiesRequest) (*model.ListVulnerabilitiesResponse, error) {
	requestDef := GenReqDefForListVulnerabilities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVulnerabilitiesResponse), nil
	}
}

// ListVulnerabilitiesInvoker 查询漏洞列表
func (c *SecMasterClient) ListVulnerabilitiesInvoker(request *model.ListVulnerabilitiesRequest) *ListVulnerabilitiesInvoker {
	requestDef := GenReqDefForListVulnerabilities()
	return &ListVulnerabilitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkflows 查询流程列表
//
// 查询流程列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListWorkflows(request *model.ListWorkflowsRequest) (*model.ListWorkflowsResponse, error) {
	requestDef := GenReqDefForListWorkflows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkflowsResponse), nil
	}
}

// ListWorkflowsInvoker 查询流程列表
func (c *SecMasterClient) ListWorkflowsInvoker(request *model.ListWorkflowsRequest) *ListWorkflowsInvoker {
	requestDef := GenReqDefForListWorkflows()
	return &ListWorkflowsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkspaces 查询工作空间列表
//
// 可通过工作空间名称、工作空间描述、创建时间等条件对租户的工作空间进行筛选。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListWorkspaces(request *model.ListWorkspacesRequest) (*model.ListWorkspacesResponse, error) {
	requestDef := GenReqDefForListWorkspaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkspacesResponse), nil
	}
}

// ListWorkspacesInvoker 查询工作空间列表
func (c *SecMasterClient) ListWorkspacesInvoker(request *model.ListWorkspacesRequest) *ListWorkspacesInvoker {
	requestDef := GenReqDefForListWorkspaces()
	return &ListWorkspacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PauseShipper 投递挂起
//
// 投递挂起
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) PauseShipper(request *model.PauseShipperRequest) (*model.PauseShipperResponse, error) {
	requestDef := GenReqDefForPauseShipper()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PauseShipperResponse), nil
	}
}

// PauseShipperInvoker 投递挂起
func (c *SecMasterClient) PauseShipperInvoker(request *model.PauseShipperRequest) *PauseShipperInvoker {
	requestDef := GenReqDefForPauseShipper()
	return &PauseShipperInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResumeShipper 启动投递
//
// 启动投递
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ResumeShipper(request *model.ResumeShipperRequest) (*model.ResumeShipperResponse, error) {
	requestDef := GenReqDefForResumeShipper()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResumeShipperResponse), nil
	}
}

// ResumeShipperInvoker 启动投递
func (c *SecMasterClient) ResumeShipperInvoker(request *model.ResumeShipperRequest) *ResumeShipperInvoker {
	requestDef := GenReqDefForResumeShipper()
	return &ResumeShipperInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetryShipper 重新投递
//
// 重新投递
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) RetryShipper(request *model.RetryShipperRequest) (*model.RetryShipperResponse, error) {
	requestDef := GenReqDefForRetryShipper()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryShipperResponse), nil
	}
}

// RetryShipperInvoker 重新投递
func (c *SecMasterClient) RetryShipperInvoker(request *model.RetryShipperRequest) *RetryShipperInvoker {
	requestDef := GenReqDefForRetryShipper()
	return &RetryShipperInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetryShipperAuthorization 重新授权
//
// 重新授权
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) RetryShipperAuthorization(request *model.RetryShipperAuthorizationRequest) (*model.RetryShipperAuthorizationResponse, error) {
	requestDef := GenReqDefForRetryShipperAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryShipperAuthorizationResponse), nil
	}
}

// RetryShipperAuthorizationInvoker 重新授权
func (c *SecMasterClient) RetryShipperAuthorizationInvoker(request *model.RetryShipperAuthorizationRequest) *RetryShipperAuthorizationInvoker {
	requestDef := GenReqDefForRetryShipperAuthorization()
	return &RetryShipperAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchPolicy 查询策略视图列表
//
// 查询策略视图列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) SearchPolicy(request *model.SearchPolicyRequest) (*model.SearchPolicyResponse, error) {
	requestDef := GenReqDefForSearchPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchPolicyResponse), nil
	}
}

// SearchPolicyInvoker 查询策略视图列表
func (c *SecMasterClient) SearchPolicyInvoker(request *model.SearchPolicyRequest) *SearchPolicyInvoker {
	requestDef := GenReqDefForSearchPolicy()
	return &SearchPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchPolicyRecord 查询任务视图列表
//
// 查询任务视图列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) SearchPolicyRecord(request *model.SearchPolicyRecordRequest) (*model.SearchPolicyRecordResponse, error) {
	requestDef := GenReqDefForSearchPolicyRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchPolicyRecordResponse), nil
	}
}

// SearchPolicyRecordInvoker 查询任务视图列表
func (c *SecMasterClient) SearchPolicyRecordInvoker(request *model.SearchPolicyRecordRequest) *SearchPolicyRecordInvoker {
	requestDef := GenReqDefForSearchPolicyRecord()
	return &SearchPolicyRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchPolicyRecordByPolicyId 根据策略ID查询历史记录列表
//
// 根据策略ID查询历史记录列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) SearchPolicyRecordByPolicyId(request *model.SearchPolicyRecordByPolicyIdRequest) (*model.SearchPolicyRecordByPolicyIdResponse, error) {
	requestDef := GenReqDefForSearchPolicyRecordByPolicyId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchPolicyRecordByPolicyIdResponse), nil
	}
}

// SearchPolicyRecordByPolicyIdInvoker 根据策略ID查询历史记录列表
func (c *SecMasterClient) SearchPolicyRecordByPolicyIdInvoker(request *model.SearchPolicyRecordByPolicyIdRequest) *SearchPolicyRecordByPolicyIdInvoker {
	requestDef := GenReqDefForSearchPolicyRecordByPolicyId()
	return &SearchPolicyRecordByPolicyIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlert 获取告警详情
//
// 获取告警详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowAlert(request *model.ShowAlertRequest) (*model.ShowAlertResponse, error) {
	requestDef := GenReqDefForShowAlert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlertResponse), nil
	}
}

// ShowAlertInvoker 获取告警详情
func (c *SecMasterClient) ShowAlertInvoker(request *model.ShowAlertRequest) *ShowAlertInvoker {
	requestDef := GenReqDefForShowAlert()
	return &ShowAlertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlertRule 查看告警规则
//
// 查看告警规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowAlertRule(request *model.ShowAlertRuleRequest) (*model.ShowAlertRuleResponse, error) {
	requestDef := GenReqDefForShowAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlertRuleResponse), nil
	}
}

// ShowAlertRuleInvoker 查看告警规则
func (c *SecMasterClient) ShowAlertRuleInvoker(request *model.ShowAlertRuleRequest) *ShowAlertRuleInvoker {
	requestDef := GenReqDefForShowAlertRule()
	return &ShowAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlertRuleTemplate 查看告警规则模板
//
// 查看告警规则模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowAlertRuleTemplate(request *model.ShowAlertRuleTemplateRequest) (*model.ShowAlertRuleTemplateResponse, error) {
	requestDef := GenReqDefForShowAlertRuleTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlertRuleTemplateResponse), nil
	}
}

// ShowAlertRuleTemplateInvoker 查看告警规则模板
func (c *SecMasterClient) ShowAlertRuleTemplateInvoker(request *model.ShowAlertRuleTemplateRequest) *ShowAlertRuleTemplateInvoker {
	requestDef := GenReqDefForShowAlertRuleTemplate()
	return &ShowAlertRuleTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAopWorkflow 查询流程详情
//
// 查询流程详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowAopWorkflow(request *model.ShowAopWorkflowRequest) (*model.ShowAopWorkflowResponse, error) {
	requestDef := GenReqDefForShowAopWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAopWorkflowResponse), nil
	}
}

// ShowAopWorkflowInvoker 查询流程详情
func (c *SecMasterClient) ShowAopWorkflowInvoker(request *model.ShowAopWorkflowRequest) *ShowAopWorkflowInvoker {
	requestDef := GenReqDefForShowAopWorkflow()
	return &ShowAopWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAopWorkflowInstance 查询流程实例的详情
//
// 查询流程实例的详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowAopWorkflowInstance(request *model.ShowAopWorkflowInstanceRequest) (*model.ShowAopWorkflowInstanceResponse, error) {
	requestDef := GenReqDefForShowAopWorkflowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAopWorkflowInstanceResponse), nil
	}
}

// ShowAopWorkflowInstanceInvoker 查询流程实例的详情
func (c *SecMasterClient) ShowAopWorkflowInstanceInvoker(request *model.ShowAopWorkflowInstanceRequest) *ShowAopWorkflowInstanceInvoker {
	requestDef := GenReqDefForShowAopWorkflowInstance()
	return &ShowAopWorkflowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAopWorkflowVersion 查询流程版本的详细信息
//
// 查询流程版本的详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowAopWorkflowVersion(request *model.ShowAopWorkflowVersionRequest) (*model.ShowAopWorkflowVersionResponse, error) {
	requestDef := GenReqDefForShowAopWorkflowVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAopWorkflowVersionResponse), nil
	}
}

// ShowAopWorkflowVersionInvoker 查询流程版本的详细信息
func (c *SecMasterClient) ShowAopWorkflowVersionInvoker(request *model.ShowAopWorkflowVersionRequest) *ShowAopWorkflowVersionInvoker {
	requestDef := GenReqDefForShowAopWorkflowVersion()
	return &ShowAopWorkflowVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAttachment 获取附件详情
//
// 获取附件详情信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowAttachment(request *model.ShowAttachmentRequest) (*model.ShowAttachmentResponse, error) {
	requestDef := GenReqDefForShowAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAttachmentResponse), nil
	}
}

// ShowAttachmentInvoker 获取附件详情
func (c *SecMasterClient) ShowAttachmentInvoker(request *model.ShowAttachmentRequest) *ShowAttachmentInvoker {
	requestDef := GenReqDefForShowAttachment()
	return &ShowAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClassifierInfo 查询分类详情
//
// 查询分类详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowClassifierInfo(request *model.ShowClassifierInfoRequest) (*model.ShowClassifierInfoResponse, error) {
	requestDef := GenReqDefForShowClassifierInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClassifierInfoResponse), nil
	}
}

// ShowClassifierInfoInvoker 查询分类详情
func (c *SecMasterClient) ShowClassifierInfoInvoker(request *model.ShowClassifierInfoRequest) *ShowClassifierInfoInvoker {
	requestDef := GenReqDefForShowClassifierInfo()
	return &ShowClassifierInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCloudLogTenantWhitelist 获取是否开启行管
//
// 获取是否开启行管
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowCloudLogTenantWhitelist(request *model.ShowCloudLogTenantWhitelistRequest) (*model.ShowCloudLogTenantWhitelistResponse, error) {
	requestDef := GenReqDefForShowCloudLogTenantWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCloudLogTenantWhitelistResponse), nil
	}
}

// ShowCloudLogTenantWhitelistInvoker 获取是否开启行管
func (c *SecMasterClient) ShowCloudLogTenantWhitelistInvoker(request *model.ShowCloudLogTenantWhitelistRequest) *ShowCloudLogTenantWhitelistInvoker {
	requestDef := GenReqDefForShowCloudLogTenantWhitelist()
	return &ShowCloudLogTenantWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCollectorChannel 展示采集通道
//
// 展示采集通道
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowCollectorChannel(request *model.ShowCollectorChannelRequest) (*model.ShowCollectorChannelResponse, error) {
	requestDef := GenReqDefForShowCollectorChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCollectorChannelResponse), nil
	}
}

// ShowCollectorChannelInvoker 展示采集通道
func (c *SecMasterClient) ShowCollectorChannelInvoker(request *model.ShowCollectorChannelRequest) *ShowCollectorChannelInvoker {
	requestDef := GenReqDefForShowCollectorChannel()
	return &ShowCollectorChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCollectorConnection 展示采集连接详情
//
// 展示采集连接详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowCollectorConnection(request *model.ShowCollectorConnectionRequest) (*model.ShowCollectorConnectionResponse, error) {
	requestDef := GenReqDefForShowCollectorConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCollectorConnectionResponse), nil
	}
}

// ShowCollectorConnectionInvoker 展示采集连接详情
func (c *SecMasterClient) ShowCollectorConnectionInvoker(request *model.ShowCollectorConnectionRequest) *ShowCollectorConnectionInvoker {
	requestDef := GenReqDefForShowCollectorConnection()
	return &ShowCollectorConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCollectorParser 获取采集解析器详情
//
// 获取采集解析器详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowCollectorParser(request *model.ShowCollectorParserRequest) (*model.ShowCollectorParserResponse, error) {
	requestDef := GenReqDefForShowCollectorParser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCollectorParserResponse), nil
	}
}

// ShowCollectorParserInvoker 获取采集解析器详情
func (c *SecMasterClient) ShowCollectorParserInvoker(request *model.ShowCollectorParserRequest) *ShowCollectorParserInvoker {
	requestDef := GenReqDefForShowCollectorParser()
	return &ShowCollectorParserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowComponent 查询插件详细信息
//
// 查询插件详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowComponent(request *model.ShowComponentRequest) (*model.ShowComponentResponse, error) {
	requestDef := GenReqDefForShowComponent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowComponentResponse), nil
	}
}

// ShowComponentInvoker 查询插件详细信息
func (c *SecMasterClient) ShowComponentInvoker(request *model.ShowComponentRequest) *ShowComponentInvoker {
	requestDef := GenReqDefForShowComponent()
	return &ShowComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowComponentAction 查询插件Action详情
//
// 查询插件的action详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowComponentAction(request *model.ShowComponentActionRequest) (*model.ShowComponentActionResponse, error) {
	requestDef := GenReqDefForShowComponentAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowComponentActionResponse), nil
	}
}

// ShowComponentActionInvoker 查询插件Action详情
func (c *SecMasterClient) ShowComponentActionInvoker(request *model.ShowComponentActionRequest) *ShowComponentActionInvoker {
	requestDef := GenReqDefForShowComponentAction()
	return &ShowComponentActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowComponentTemplate 查询插件配置模板详情
//
// 查询在配置流程时的插件配置模板详情信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowComponentTemplate(request *model.ShowComponentTemplateRequest) (*model.ShowComponentTemplateResponse, error) {
	requestDef := GenReqDefForShowComponentTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowComponentTemplateResponse), nil
	}
}

// ShowComponentTemplateInvoker 查询插件配置模板详情
func (c *SecMasterClient) ShowComponentTemplateInvoker(request *model.ShowComponentTemplateRequest) *ShowComponentTemplateInvoker {
	requestDef := GenReqDefForShowComponentTemplate()
	return &ShowComponentTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConnection 查询操作连接详情
//
// 查询操作连接详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowConnection(request *model.ShowConnectionRequest) (*model.ShowConnectionResponse, error) {
	requestDef := GenReqDefForShowConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConnectionResponse), nil
	}
}

// ShowConnectionInvoker 查询操作连接详情
func (c *SecMasterClient) ShowConnectionInvoker(request *model.ShowConnectionRequest) *ShowConnectionInvoker {
	requestDef := GenReqDefForShowConnection()
	return &ShowConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataClassInfo 展示数据类详情
//
// 展示数据类详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowDataClassInfo(request *model.ShowDataClassInfoRequest) (*model.ShowDataClassInfoResponse, error) {
	requestDef := GenReqDefForShowDataClassInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataClassInfoResponse), nil
	}
}

// ShowDataClassInfoInvoker 展示数据类详情
func (c *SecMasterClient) ShowDataClassInfoInvoker(request *model.ShowDataClassInfoRequest) *ShowDataClassInfoInvoker {
	requestDef := GenReqDefForShowDataClassInfo()
	return &ShowDataClassInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataobject 查询数据对象
//
// 查询数据对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowDataobject(request *model.ShowDataobjectRequest) (*model.ShowDataobjectResponse, error) {
	requestDef := GenReqDefForShowDataobject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataobjectResponse), nil
	}
}

// ShowDataobjectInvoker 查询数据对象
func (c *SecMasterClient) ShowDataobjectInvoker(request *model.ShowDataobjectRequest) *ShowDataobjectInvoker {
	requestDef := GenReqDefForShowDataobject()
	return &ShowDataobjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDatapanelObject 查询数据面对象
//
// 数据面查询数据类纳管的数据对象详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowDatapanelObject(request *model.ShowDatapanelObjectRequest) (*model.ShowDatapanelObjectResponse, error) {
	requestDef := GenReqDefForShowDatapanelObject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDatapanelObjectResponse), nil
	}
}

// ShowDatapanelObjectInvoker 查询数据面对象
func (c *SecMasterClient) ShowDatapanelObjectInvoker(request *model.ShowDatapanelObjectRequest) *ShowDatapanelObjectInvoker {
	requestDef := GenReqDefForShowDatapanelObject()
	return &ShowDatapanelObjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataspace 获取数据空间详情
//
// 获取数据空间详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowDataspace(request *model.ShowDataspaceRequest) (*model.ShowDataspaceResponse, error) {
	requestDef := GenReqDefForShowDataspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataspaceResponse), nil
	}
}

// ShowDataspaceInvoker 获取数据空间详情
func (c *SecMasterClient) ShowDataspaceInvoker(request *model.ShowDataspaceRequest) *ShowDataspaceInvoker {
	requestDef := GenReqDefForShowDataspace()
	return &ShowDataspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIncident 获取事件详情
//
// 获取事件详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowIncident(request *model.ShowIncidentRequest) (*model.ShowIncidentResponse, error) {
	requestDef := GenReqDefForShowIncident()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIncidentResponse), nil
	}
}

// ShowIncidentInvoker 获取事件详情
func (c *SecMasterClient) ShowIncidentInvoker(request *model.ShowIncidentRequest) *ShowIncidentInvoker {
	requestDef := GenReqDefForShowIncident()
	return &ShowIncidentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIndicatorDetail 查询威胁情报详情
//
// 查询威胁情报详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowIndicatorDetail(request *model.ShowIndicatorDetailRequest) (*model.ShowIndicatorDetailResponse, error) {
	requestDef := GenReqDefForShowIndicatorDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIndicatorDetailResponse), nil
	}
}

// ShowIndicatorDetailInvoker 查询威胁情报详情
func (c *SecMasterClient) ShowIndicatorDetailInvoker(request *model.ShowIndicatorDetailRequest) *ShowIndicatorDetailInvoker {
	requestDef := GenReqDefForShowIndicatorDetail()
	return &ShowIndicatorDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIsapComponent 展示组件详情
//
// 展示组件详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowIsapComponent(request *model.ShowIsapComponentRequest) (*model.ShowIsapComponentResponse, error) {
	requestDef := GenReqDefForShowIsapComponent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIsapComponentResponse), nil
	}
}

// ShowIsapComponentInvoker 展示组件详情
func (c *SecMasterClient) ShowIsapComponentInvoker(request *model.ShowIsapComponentRequest) *ShowIsapComponentInvoker {
	requestDef := GenReqDefForShowIsapComponent()
	return &ShowIsapComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLayout 查询布局详情
//
// 查询布局详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowLayout(request *model.ShowLayoutRequest) (*model.ShowLayoutResponse, error) {
	requestDef := GenReqDefForShowLayout()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLayoutResponse), nil
	}
}

// ShowLayoutInvoker 查询布局详情
func (c *SecMasterClient) ShowLayoutInvoker(request *model.ShowLayoutRequest) *ShowLayoutInvoker {
	requestDef := GenReqDefForShowLayout()
	return &ShowLayoutInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLayoutWizard 查询布局页面
//
// 查询布局页面
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowLayoutWizard(request *model.ShowLayoutWizardRequest) (*model.ShowLayoutWizardResponse, error) {
	requestDef := GenReqDefForShowLayoutWizard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLayoutWizardResponse), nil
	}
}

// ShowLayoutWizardInvoker 查询布局页面
func (c *SecMasterClient) ShowLayoutWizardInvoker(request *model.ShowLayoutWizardRequest) *ShowLayoutWizardInvoker {
	requestDef := GenReqDefForShowLayoutWizard()
	return &ShowLayoutWizardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLayoutWizardByField 查询布局页面详情
//
// 查询布局页面详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowLayoutWizardByField(request *model.ShowLayoutWizardByFieldRequest) (*model.ShowLayoutWizardByFieldResponse, error) {
	requestDef := GenReqDefForShowLayoutWizardByField()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLayoutWizardByFieldResponse), nil
	}
}

// ShowLayoutWizardByFieldInvoker 查询布局页面详情
func (c *SecMasterClient) ShowLayoutWizardByFieldInvoker(request *model.ShowLayoutWizardByFieldRequest) *ShowLayoutWizardByFieldInvoker {
	requestDef := GenReqDefForShowLayoutWizardByField()
	return &ShowLayoutWizardByFieldInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMapperDetail 查询映射详情
//
// 查询映射详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowMapperDetail(request *model.ShowMapperDetailRequest) (*model.ShowMapperDetailResponse, error) {
	requestDef := GenReqDefForShowMapperDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMapperDetailResponse), nil
	}
}

// ShowMapperDetailInvoker 查询映射详情
func (c *SecMasterClient) ShowMapperDetailInvoker(request *model.ShowMapperDetailRequest) *ShowMapperDetailInvoker {
	requestDef := GenReqDefForShowMapperDetail()
	return &ShowMapperDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMapperList 查询映射列表
//
// 查询映射列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowMapperList(request *model.ShowMapperListRequest) (*model.ShowMapperListResponse, error) {
	requestDef := GenReqDefForShowMapperList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMapperListResponse), nil
	}
}

// ShowMapperListInvoker 查询映射列表
func (c *SecMasterClient) ShowMapperListInvoker(request *model.ShowMapperListRequest) *ShowMapperListInvoker {
	requestDef := GenReqDefForShowMapperList()
	return &ShowMapperListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMappingFunction 查询分类映射函数
//
// 查询分类映射函数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowMappingFunction(request *model.ShowMappingFunctionRequest) (*model.ShowMappingFunctionResponse, error) {
	requestDef := GenReqDefForShowMappingFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMappingFunctionResponse), nil
	}
}

// ShowMappingFunctionInvoker 查询分类映射函数
func (c *SecMasterClient) ShowMappingFunctionInvoker(request *model.ShowMappingFunctionRequest) *ShowMappingFunctionInvoker {
	requestDef := GenReqDefForShowMappingFunction()
	return &ShowMappingFunctionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMappingInfoList 查询分类映射列表
//
// 查询分类映射列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowMappingInfoList(request *model.ShowMappingInfoListRequest) (*model.ShowMappingInfoListResponse, error) {
	requestDef := GenReqDefForShowMappingInfoList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMappingInfoListResponse), nil
	}
}

// ShowMappingInfoListInvoker 查询分类映射列表
func (c *SecMasterClient) ShowMappingInfoListInvoker(request *model.ShowMappingInfoListRequest) *ShowMappingInfoListInvoker {
	requestDef := GenReqDefForShowMappingInfoList()
	return &ShowMappingInfoListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMetricMetaData 获取指标元数据
//
// 获取指标元数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowMetricMetaData(request *model.ShowMetricMetaDataRequest) (*model.ShowMetricMetaDataResponse, error) {
	requestDef := GenReqDefForShowMetricMetaData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMetricMetaDataResponse), nil
	}
}

// ShowMetricMetaDataInvoker 获取指标元数据
func (c *SecMasterClient) ShowMetricMetaDataInvoker(request *model.ShowMetricMetaDataRequest) *ShowMetricMetaDataInvoker {
	requestDef := GenReqDefForShowMetricMetaData()
	return &ShowMetricMetaDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowModule 查询模块详情
//
// 查询模块详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowModule(request *model.ShowModuleRequest) (*model.ShowModuleResponse, error) {
	requestDef := GenReqDefForShowModule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowModuleResponse), nil
	}
}

// ShowModuleInvoker 查询模块详情
func (c *SecMasterClient) ShowModuleInvoker(request *model.ShowModuleRequest) *ShowModuleInvoker {
	requestDef := GenReqDefForShowModule()
	return &ShowModuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMoniterMetricStats 获取指标统计数据
//
// 获取指标统计数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowMoniterMetricStats(request *model.ShowMoniterMetricStatsRequest) (*model.ShowMoniterMetricStatsResponse, error) {
	requestDef := GenReqDefForShowMoniterMetricStats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMoniterMetricStatsResponse), nil
	}
}

// ShowMoniterMetricStatsInvoker 获取指标统计数据
func (c *SecMasterClient) ShowMoniterMetricStatsInvoker(request *model.ShowMoniterMetricStatsRequest) *ShowMoniterMetricStatsInvoker {
	requestDef := GenReqDefForShowMoniterMetricStats()
	return &ShowMoniterMetricStatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPipe 获取数据管道详情
//
// 获取数据管道详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPipe(request *model.ShowPipeRequest) (*model.ShowPipeResponse, error) {
	requestDef := GenReqDefForShowPipe()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPipeResponse), nil
	}
}

// ShowPipeInvoker 获取数据管道详情
func (c *SecMasterClient) ShowPipeInvoker(request *model.ShowPipeRequest) *ShowPipeInvoker {
	requestDef := GenReqDefForShowPipe()
	return &ShowPipeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPipeConsumption 获取实时消费配置
//
// 获取实时消费配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPipeConsumption(request *model.ShowPipeConsumptionRequest) (*model.ShowPipeConsumptionResponse, error) {
	requestDef := GenReqDefForShowPipeConsumption()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPipeConsumptionResponse), nil
	}
}

// ShowPipeConsumptionInvoker 获取实时消费配置
func (c *SecMasterClient) ShowPipeConsumptionInvoker(request *model.ShowPipeConsumptionRequest) *ShowPipeConsumptionInvoker {
	requestDef := GenReqDefForShowPipeConsumption()
	return &ShowPipeConsumptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPipeIndex 获取索引信息
//
// 获取索引信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPipeIndex(request *model.ShowPipeIndexRequest) (*model.ShowPipeIndexResponse, error) {
	requestDef := GenReqDefForShowPipeIndex()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPipeIndexResponse), nil
	}
}

// ShowPipeIndexInvoker 获取索引信息
func (c *SecMasterClient) ShowPipeIndexInvoker(request *model.ShowPipeIndexRequest) *ShowPipeIndexInvoker {
	requestDef := GenReqDefForShowPipeIndex()
	return &ShowPipeIndexInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlatformManaged 获取行管信息
//
// 获取行管信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPlatformManaged(request *model.ShowPlatformManagedRequest) (*model.ShowPlatformManagedResponse, error) {
	requestDef := GenReqDefForShowPlatformManaged()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlatformManagedResponse), nil
	}
}

// ShowPlatformManagedInvoker 获取行管信息
func (c *SecMasterClient) ShowPlatformManagedInvoker(request *model.ShowPlatformManagedRequest) *ShowPlatformManagedInvoker {
	requestDef := GenReqDefForShowPlatformManaged()
	return &ShowPlatformManagedInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybook 查询剧本详情
//
// 查询剧本详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPlaybook(request *model.ShowPlaybookRequest) (*model.ShowPlaybookResponse, error) {
	requestDef := GenReqDefForShowPlaybook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookResponse), nil
	}
}

// ShowPlaybookInvoker 查询剧本详情
func (c *SecMasterClient) ShowPlaybookInvoker(request *model.ShowPlaybookRequest) *ShowPlaybookInvoker {
	requestDef := GenReqDefForShowPlaybook()
	return &ShowPlaybookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybookInstance 查询剧本实例详情
//
// 查询剧本实例详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPlaybookInstance(request *model.ShowPlaybookInstanceRequest) (*model.ShowPlaybookInstanceResponse, error) {
	requestDef := GenReqDefForShowPlaybookInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookInstanceResponse), nil
	}
}

// ShowPlaybookInstanceInvoker 查询剧本实例详情
func (c *SecMasterClient) ShowPlaybookInstanceInvoker(request *model.ShowPlaybookInstanceRequest) *ShowPlaybookInstanceInvoker {
	requestDef := GenReqDefForShowPlaybookInstance()
	return &ShowPlaybookInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybookMonitors 剧本运行监控
//
// 剧本运行监控（待下线）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPlaybookMonitors(request *model.ShowPlaybookMonitorsRequest) (*model.ShowPlaybookMonitorsResponse, error) {
	requestDef := GenReqDefForShowPlaybookMonitors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookMonitorsResponse), nil
	}
}

// ShowPlaybookMonitorsInvoker 剧本运行监控
func (c *SecMasterClient) ShowPlaybookMonitorsInvoker(request *model.ShowPlaybookMonitorsRequest) *ShowPlaybookMonitorsInvoker {
	requestDef := GenReqDefForShowPlaybookMonitors()
	return &ShowPlaybookMonitorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybookRule 查询剧本规则详情
//
// 查询剧本规则详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPlaybookRule(request *model.ShowPlaybookRuleRequest) (*model.ShowPlaybookRuleResponse, error) {
	requestDef := GenReqDefForShowPlaybookRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookRuleResponse), nil
	}
}

// ShowPlaybookRuleInvoker 查询剧本规则详情
func (c *SecMasterClient) ShowPlaybookRuleInvoker(request *model.ShowPlaybookRuleRequest) *ShowPlaybookRuleInvoker {
	requestDef := GenReqDefForShowPlaybookRule()
	return &ShowPlaybookRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybookStatistics 剧本数据统计
//
// 剧本统计数据（待下线）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPlaybookStatistics(request *model.ShowPlaybookStatisticsRequest) (*model.ShowPlaybookStatisticsResponse, error) {
	requestDef := GenReqDefForShowPlaybookStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookStatisticsResponse), nil
	}
}

// ShowPlaybookStatisticsInvoker 剧本数据统计
func (c *SecMasterClient) ShowPlaybookStatisticsInvoker(request *model.ShowPlaybookStatisticsRequest) *ShowPlaybookStatisticsInvoker {
	requestDef := GenReqDefForShowPlaybookStatistics()
	return &ShowPlaybookStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybookTopology 查询剧本拓扑关系
//
// 查询剧本拓扑关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPlaybookTopology(request *model.ShowPlaybookTopologyRequest) (*model.ShowPlaybookTopologyResponse, error) {
	requestDef := GenReqDefForShowPlaybookTopology()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookTopologyResponse), nil
	}
}

// ShowPlaybookTopologyInvoker 查询剧本拓扑关系
func (c *SecMasterClient) ShowPlaybookTopologyInvoker(request *model.ShowPlaybookTopologyRequest) *ShowPlaybookTopologyInvoker {
	requestDef := GenReqDefForShowPlaybookTopology()
	return &ShowPlaybookTopologyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybookVersion 查询剧本版本详情
//
// 查询剧本版本详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPlaybookVersion(request *model.ShowPlaybookVersionRequest) (*model.ShowPlaybookVersionResponse, error) {
	requestDef := GenReqDefForShowPlaybookVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookVersionResponse), nil
	}
}

// ShowPlaybookVersionInvoker 查询剧本版本详情
func (c *SecMasterClient) ShowPlaybookVersionInvoker(request *model.ShowPlaybookVersionRequest) *ShowPlaybookVersionInvoker {
	requestDef := GenReqDefForShowPlaybookVersion()
	return &ShowPlaybookVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPolicy 查询策略视图对象
//
// 查询策略视图对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPolicy(request *model.ShowPolicyRequest) (*model.ShowPolicyResponse, error) {
	requestDef := GenReqDefForShowPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPolicyResponse), nil
	}
}

// ShowPolicyInvoker 查询策略视图对象
func (c *SecMasterClient) ShowPolicyInvoker(request *model.ShowPolicyRequest) *ShowPolicyInvoker {
	requestDef := GenReqDefForShowPolicy()
	return &ShowPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPreProcessRulesList 查询预处理规则列表
//
// 查询预处理规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPreProcessRulesList(request *model.ShowPreProcessRulesListRequest) (*model.ShowPreProcessRulesListResponse, error) {
	requestDef := GenReqDefForShowPreProcessRulesList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPreProcessRulesListResponse), nil
	}
}

// ShowPreProcessRulesListInvoker 查询预处理规则列表
func (c *SecMasterClient) ShowPreProcessRulesListInvoker(request *model.ShowPreProcessRulesListRequest) *ShowPreProcessRulesListInvoker {
	requestDef := GenReqDefForShowPreProcessRulesList()
	return &ShowPreProcessRulesListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReportInfo 获取报告详情
//
// 获取报告详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowReportInfo(request *model.ShowReportInfoRequest) (*model.ShowReportInfoResponse, error) {
	requestDef := GenReqDefForShowReportInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReportInfoResponse), nil
	}
}

// ShowReportInfoInvoker 获取报告详情
func (c *SecMasterClient) ShowReportInfoInvoker(request *model.ShowReportInfoRequest) *ShowReportInfoInvoker {
	requestDef := GenReqDefForShowReportInfo()
	return &ShowReportInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResource 获取资产详情
//
// 获取资产详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowResource(request *model.ShowResourceRequest) (*model.ShowResourceResponse, error) {
	requestDef := GenReqDefForShowResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceResponse), nil
	}
}

// ShowResourceInvoker 获取资产详情
func (c *SecMasterClient) ShowResourceInvoker(request *model.ShowResourceRequest) *ShowResourceInvoker {
	requestDef := GenReqDefForShowResource()
	return &ShowResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSearchCondition 获取检索条件详情
//
// 获取检索条件详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowSearchCondition(request *model.ShowSearchConditionRequest) (*model.ShowSearchConditionResponse, error) {
	requestDef := GenReqDefForShowSearchCondition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSearchConditionResponse), nil
	}
}

// ShowSearchConditionInvoker 获取检索条件详情
func (c *SecMasterClient) ShowSearchConditionInvoker(request *model.ShowSearchConditionRequest) *ShowSearchConditionInvoker {
	requestDef := GenReqDefForShowSearchCondition()
	return &ShowSearchConditionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowShipper 投递规则详情
//
// 投递规则详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowShipper(request *model.ShowShipperRequest) (*model.ShowShipperResponse, error) {
	requestDef := GenReqDefForShowShipper()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowShipperResponse), nil
	}
}

// ShowShipperInvoker 投递规则详情
func (c *SecMasterClient) ShowShipperInvoker(request *model.ShowShipperRequest) *ShowShipperInvoker {
	requestDef := GenReqDefForShowShipper()
	return &ShowShipperInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowShipperDelegateAuth 获取委托权限
//
// 获取委托权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowShipperDelegateAuth(request *model.ShowShipperDelegateAuthRequest) (*model.ShowShipperDelegateAuthResponse, error) {
	requestDef := GenReqDefForShowShipperDelegateAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowShipperDelegateAuthResponse), nil
	}
}

// ShowShipperDelegateAuthInvoker 获取委托权限
func (c *SecMasterClient) ShowShipperDelegateAuthInvoker(request *model.ShowShipperDelegateAuthRequest) *ShowShipperDelegateAuthInvoker {
	requestDef := GenReqDefForShowShipperDelegateAuth()
	return &ShowShipperDelegateAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowShipperParam 参数查询
//
// 参数查询接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowShipperParam(request *model.ShowShipperParamRequest) (*model.ShowShipperParamResponse, error) {
	requestDef := GenReqDefForShowShipperParam()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowShipperParamResponse), nil
	}
}

// ShowShipperParamInvoker 参数查询
func (c *SecMasterClient) ShowShipperParamInvoker(request *model.ShowShipperParamRequest) *ShowShipperParamInvoker {
	requestDef := GenReqDefForShowShipperParam()
	return &ShowShipperParamInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVulnerability 获取漏洞详情
//
// 获取漏洞详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowVulnerability(request *model.ShowVulnerabilityRequest) (*model.ShowVulnerabilityResponse, error) {
	requestDef := GenReqDefForShowVulnerability()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVulnerabilityResponse), nil
	}
}

// ShowVulnerabilityInvoker 获取漏洞详情
func (c *SecMasterClient) ShowVulnerabilityInvoker(request *model.ShowVulnerabilityRequest) *ShowVulnerabilityInvoker {
	requestDef := GenReqDefForShowVulnerability()
	return &ShowVulnerabilityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkspace 查询工作空间详情
//
// 查询工作空间名称、描述等详情信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowWorkspace(request *model.ShowWorkspaceRequest) (*model.ShowWorkspaceResponse, error) {
	requestDef := GenReqDefForShowWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkspaceResponse), nil
	}
}

// ShowWorkspaceInvoker 查询工作空间详情
func (c *SecMasterClient) ShowWorkspaceInvoker(request *model.ShowWorkspaceRequest) *ShowWorkspaceInvoker {
	requestDef := GenReqDefForShowWorkspace()
	return &ShowWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAlertRule 更新告警规则
//
// 更新告警规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateAlertRule(request *model.UpdateAlertRuleRequest) (*model.UpdateAlertRuleResponse, error) {
	requestDef := GenReqDefForUpdateAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlertRuleResponse), nil
	}
}

// UpdateAlertRuleInvoker 更新告警规则
func (c *SecMasterClient) UpdateAlertRuleInvoker(request *model.UpdateAlertRuleRequest) *UpdateAlertRuleInvoker {
	requestDef := GenReqDefForUpdateAlertRule()
	return &UpdateAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAopWorkflow 更新流程
//
// 更新流程
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateAopWorkflow(request *model.UpdateAopWorkflowRequest) (*model.UpdateAopWorkflowResponse, error) {
	requestDef := GenReqDefForUpdateAopWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAopWorkflowResponse), nil
	}
}

// UpdateAopWorkflowInvoker 更新流程
func (c *SecMasterClient) UpdateAopWorkflowInvoker(request *model.UpdateAopWorkflowRequest) *UpdateAopWorkflowInvoker {
	requestDef := GenReqDefForUpdateAopWorkflow()
	return &UpdateAopWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAopWorkflowVersion 更新流程版本信息
//
// 更新流程版本信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateAopWorkflowVersion(request *model.UpdateAopWorkflowVersionRequest) (*model.UpdateAopWorkflowVersionResponse, error) {
	requestDef := GenReqDefForUpdateAopWorkflowVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAopWorkflowVersionResponse), nil
	}
}

// UpdateAopWorkflowVersionInvoker 更新流程版本信息
func (c *SecMasterClient) UpdateAopWorkflowVersionInvoker(request *model.UpdateAopWorkflowVersionRequest) *UpdateAopWorkflowVersionInvoker {
	requestDef := GenReqDefForUpdateAopWorkflowVersion()
	return &UpdateAopWorkflowVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCatalogue 修改目录
//
// 修改自定义目录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateCatalogue(request *model.UpdateCatalogueRequest) (*model.UpdateCatalogueResponse, error) {
	requestDef := GenReqDefForUpdateCatalogue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCatalogueResponse), nil
	}
}

// UpdateCatalogueInvoker 修改目录
func (c *SecMasterClient) UpdateCatalogueInvoker(request *model.UpdateCatalogueRequest) *UpdateCatalogueInvoker {
	requestDef := GenReqDefForUpdateCatalogue()
	return &UpdateCatalogueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateClassifier 修改分类
//
// 修改分类
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateClassifier(request *model.UpdateClassifierRequest) (*model.UpdateClassifierResponse, error) {
	requestDef := GenReqDefForUpdateClassifier()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClassifierResponse), nil
	}
}

// UpdateClassifierInvoker 修改分类
func (c *SecMasterClient) UpdateClassifierInvoker(request *model.UpdateClassifierRequest) *UpdateClassifierInvoker {
	requestDef := GenReqDefForUpdateClassifier()
	return &UpdateClassifierInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCollectorChannel 更改采集通道
//
// 更改采集通道
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateCollectorChannel(request *model.UpdateCollectorChannelRequest) (*model.UpdateCollectorChannelResponse, error) {
	requestDef := GenReqDefForUpdateCollectorChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCollectorChannelResponse), nil
	}
}

// UpdateCollectorChannelInvoker 更改采集通道
func (c *SecMasterClient) UpdateCollectorChannelInvoker(request *model.UpdateCollectorChannelRequest) *UpdateCollectorChannelInvoker {
	requestDef := GenReqDefForUpdateCollectorChannel()
	return &UpdateCollectorChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCollectorChannelGroup 更新采集通道分组
//
// 更新采集通道分组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateCollectorChannelGroup(request *model.UpdateCollectorChannelGroupRequest) (*model.UpdateCollectorChannelGroupResponse, error) {
	requestDef := GenReqDefForUpdateCollectorChannelGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCollectorChannelGroupResponse), nil
	}
}

// UpdateCollectorChannelGroupInvoker 更新采集通道分组
func (c *SecMasterClient) UpdateCollectorChannelGroupInvoker(request *model.UpdateCollectorChannelGroupRequest) *UpdateCollectorChannelGroupInvoker {
	requestDef := GenReqDefForUpdateCollectorChannelGroup()
	return &UpdateCollectorChannelGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCollectorConnection 更新采集连接
//
// 更新采集连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateCollectorConnection(request *model.UpdateCollectorConnectionRequest) (*model.UpdateCollectorConnectionResponse, error) {
	requestDef := GenReqDefForUpdateCollectorConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCollectorConnectionResponse), nil
	}
}

// UpdateCollectorConnectionInvoker 更新采集连接
func (c *SecMasterClient) UpdateCollectorConnectionInvoker(request *model.UpdateCollectorConnectionRequest) *UpdateCollectorConnectionInvoker {
	requestDef := GenReqDefForUpdateCollectorConnection()
	return &UpdateCollectorConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateComponentConfiguration 点击保存按钮
//
// 点击保存按钮
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateComponentConfiguration(request *model.UpdateComponentConfigurationRequest) (*model.UpdateComponentConfigurationResponse, error) {
	requestDef := GenReqDefForUpdateComponentConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateComponentConfigurationResponse), nil
	}
}

// UpdateComponentConfigurationInvoker 点击保存按钮
func (c *SecMasterClient) UpdateComponentConfigurationInvoker(request *model.UpdateComponentConfigurationRequest) *UpdateComponentConfigurationInvoker {
	requestDef := GenReqDefForUpdateComponentConfiguration()
	return &UpdateComponentConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateComponentTemplate 更新插件配置模板
//
// 更新某个在配置流程时的插件配置模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateComponentTemplate(request *model.UpdateComponentTemplateRequest) (*model.UpdateComponentTemplateResponse, error) {
	requestDef := GenReqDefForUpdateComponentTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateComponentTemplateResponse), nil
	}
}

// UpdateComponentTemplateInvoker 更新插件配置模板
func (c *SecMasterClient) UpdateComponentTemplateInvoker(request *model.UpdateComponentTemplateRequest) *UpdateComponentTemplateInvoker {
	requestDef := GenReqDefForUpdateComponentTemplate()
	return &UpdateComponentTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateConfigurationDictionaries 更新字典
//
// 更新字典数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateConfigurationDictionaries(request *model.UpdateConfigurationDictionariesRequest) (*model.UpdateConfigurationDictionariesResponse, error) {
	requestDef := GenReqDefForUpdateConfigurationDictionaries()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConfigurationDictionariesResponse), nil
	}
}

// UpdateConfigurationDictionariesInvoker 更新字典
func (c *SecMasterClient) UpdateConfigurationDictionariesInvoker(request *model.UpdateConfigurationDictionariesRequest) *UpdateConfigurationDictionariesInvoker {
	requestDef := GenReqDefForUpdateConfigurationDictionaries()
	return &UpdateConfigurationDictionariesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateConnection 更新操作连接
//
// 更新操作连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateConnection(request *model.UpdateConnectionRequest) (*model.UpdateConnectionResponse, error) {
	requestDef := GenReqDefForUpdateConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConnectionResponse), nil
	}
}

// UpdateConnectionInvoker 更新操作连接
func (c *SecMasterClient) UpdateConnectionInvoker(request *model.UpdateConnectionRequest) *UpdateConnectionInvoker {
	requestDef := GenReqDefForUpdateConnection()
	return &UpdateConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDataobject 编辑数据对象
//
// 编辑数据对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateDataobject(request *model.UpdateDataobjectRequest) (*model.UpdateDataobjectResponse, error) {
	requestDef := GenReqDefForUpdateDataobject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDataobjectResponse), nil
	}
}

// UpdateDataobjectInvoker 编辑数据对象
func (c *SecMasterClient) UpdateDataobjectInvoker(request *model.UpdateDataobjectRequest) *UpdateDataobjectInvoker {
	requestDef := GenReqDefForUpdateDataobject()
	return &UpdateDataobjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDataspace 修改数据空间
//
// 修改数据空间
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateDataspace(request *model.UpdateDataspaceRequest) (*model.UpdateDataspaceResponse, error) {
	requestDef := GenReqDefForUpdateDataspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDataspaceResponse), nil
	}
}

// UpdateDataspaceInvoker 修改数据空间
func (c *SecMasterClient) UpdateDataspaceInvoker(request *model.UpdateDataspaceRequest) *UpdateDataspaceInvoker {
	requestDef := GenReqDefForUpdateDataspace()
	return &UpdateDataspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIndicator 更新威胁情报
//
// 更新威胁情报
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateIndicator(request *model.UpdateIndicatorRequest) (*model.UpdateIndicatorResponse, error) {
	requestDef := GenReqDefForUpdateIndicator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIndicatorResponse), nil
	}
}

// UpdateIndicatorInvoker 更新威胁情报
func (c *SecMasterClient) UpdateIndicatorInvoker(request *model.UpdateIndicatorRequest) *UpdateIndicatorInvoker {
	requestDef := GenReqDefForUpdateIndicator()
	return &UpdateIndicatorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLayout 修改布局
//
// 修改布局
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateLayout(request *model.UpdateLayoutRequest) (*model.UpdateLayoutResponse, error) {
	requestDef := GenReqDefForUpdateLayout()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLayoutResponse), nil
	}
}

// UpdateLayoutInvoker 修改布局
func (c *SecMasterClient) UpdateLayoutInvoker(request *model.UpdateLayoutRequest) *UpdateLayoutInvoker {
	requestDef := GenReqDefForUpdateLayout()
	return &UpdateLayoutInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLayoutWizards 批量更新布局页面
//
// 更新页面
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateLayoutWizards(request *model.UpdateLayoutWizardsRequest) (*model.UpdateLayoutWizardsResponse, error) {
	requestDef := GenReqDefForUpdateLayoutWizards()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLayoutWizardsResponse), nil
	}
}

// UpdateLayoutWizardsInvoker 批量更新布局页面
func (c *SecMasterClient) UpdateLayoutWizardsInvoker(request *model.UpdateLayoutWizardsRequest) *UpdateLayoutWizardsInvoker {
	requestDef := GenReqDefForUpdateLayoutWizards()
	return &UpdateLayoutWizardsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMappingInfoStatus 修分类映射启用禁用状态
//
// 修分类映射启用禁用状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateMappingInfoStatus(request *model.UpdateMappingInfoStatusRequest) (*model.UpdateMappingInfoStatusResponse, error) {
	requestDef := GenReqDefForUpdateMappingInfoStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMappingInfoStatusResponse), nil
	}
}

// UpdateMappingInfoStatusInvoker 修分类映射启用禁用状态
func (c *SecMasterClient) UpdateMappingInfoStatusInvoker(request *model.UpdateMappingInfoStatusRequest) *UpdateMappingInfoStatusInvoker {
	requestDef := GenReqDefForUpdateMappingInfoStatus()
	return &UpdateMappingInfoStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMetrics 更新指标定义
//
// 更新指标定义
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateMetrics(request *model.UpdateMetricsRequest) (*model.UpdateMetricsResponse, error) {
	requestDef := GenReqDefForUpdateMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMetricsResponse), nil
	}
}

// UpdateMetricsInvoker 更新指标定义
func (c *SecMasterClient) UpdateMetricsInvoker(request *model.UpdateMetricsRequest) *UpdateMetricsInvoker {
	requestDef := GenReqDefForUpdateMetrics()
	return &UpdateMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateModule 更新布局模块
//
// 更新模块
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateModule(request *model.UpdateModuleRequest) (*model.UpdateModuleResponse, error) {
	requestDef := GenReqDefForUpdateModule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateModuleResponse), nil
	}
}

// UpdateModuleInvoker 更新布局模块
func (c *SecMasterClient) UpdateModuleInvoker(request *model.UpdateModuleRequest) *UpdateModuleInvoker {
	requestDef := GenReqDefForUpdateModule()
	return &UpdateModuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNode 更新节点补充信息
//
// 更新节点信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateNode(request *model.UpdateNodeRequest) (*model.UpdateNodeResponse, error) {
	requestDef := GenReqDefForUpdateNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNodeResponse), nil
	}
}

// UpdateNodeInvoker 更新节点补充信息
func (c *SecMasterClient) UpdateNodeInvoker(request *model.UpdateNodeRequest) *UpdateNodeInvoker {
	requestDef := GenReqDefForUpdateNode()
	return &UpdateNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePipe 修改数据管道
//
// 修改数据管道
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdatePipe(request *model.UpdatePipeRequest) (*model.UpdatePipeResponse, error) {
	requestDef := GenReqDefForUpdatePipe()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePipeResponse), nil
	}
}

// UpdatePipeInvoker 修改数据管道
func (c *SecMasterClient) UpdatePipeInvoker(request *model.UpdatePipeRequest) *UpdatePipeInvoker {
	requestDef := GenReqDefForUpdatePipe()
	return &UpdatePipeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePipeIndex 修改索引
//
// 修改索引
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdatePipeIndex(request *model.UpdatePipeIndexRequest) (*model.UpdatePipeIndexResponse, error) {
	requestDef := GenReqDefForUpdatePipeIndex()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePipeIndexResponse), nil
	}
}

// UpdatePipeIndexInvoker 修改索引
func (c *SecMasterClient) UpdatePipeIndexInvoker(request *model.UpdatePipeIndexRequest) *UpdatePipeIndexInvoker {
	requestDef := GenReqDefForUpdatePipeIndex()
	return &UpdatePipeIndexInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePlaybook 修改剧本
//
// 修改剧本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdatePlaybook(request *model.UpdatePlaybookRequest) (*model.UpdatePlaybookResponse, error) {
	requestDef := GenReqDefForUpdatePlaybook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePlaybookResponse), nil
	}
}

// UpdatePlaybookInvoker 修改剧本
func (c *SecMasterClient) UpdatePlaybookInvoker(request *model.UpdatePlaybookRequest) *UpdatePlaybookInvoker {
	requestDef := GenReqDefForUpdatePlaybook()
	return &UpdatePlaybookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePlaybookAction 更新剧本动作
//
// 更新剧本动作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdatePlaybookAction(request *model.UpdatePlaybookActionRequest) (*model.UpdatePlaybookActionResponse, error) {
	requestDef := GenReqDefForUpdatePlaybookAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePlaybookActionResponse), nil
	}
}

// UpdatePlaybookActionInvoker 更新剧本动作
func (c *SecMasterClient) UpdatePlaybookActionInvoker(request *model.UpdatePlaybookActionRequest) *UpdatePlaybookActionInvoker {
	requestDef := GenReqDefForUpdatePlaybookAction()
	return &UpdatePlaybookActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePlaybookRule 更新剧本规则
//
// 更新剧本规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdatePlaybookRule(request *model.UpdatePlaybookRuleRequest) (*model.UpdatePlaybookRuleResponse, error) {
	requestDef := GenReqDefForUpdatePlaybookRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePlaybookRuleResponse), nil
	}
}

// UpdatePlaybookRuleInvoker 更新剧本规则
func (c *SecMasterClient) UpdatePlaybookRuleInvoker(request *model.UpdatePlaybookRuleRequest) *UpdatePlaybookRuleInvoker {
	requestDef := GenReqDefForUpdatePlaybookRule()
	return &UpdatePlaybookRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePlaybookVersion 更新剧本版本
//
// 更新剧本版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdatePlaybookVersion(request *model.UpdatePlaybookVersionRequest) (*model.UpdatePlaybookVersionResponse, error) {
	requestDef := GenReqDefForUpdatePlaybookVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePlaybookVersionResponse), nil
	}
}

// UpdatePlaybookVersionInvoker 更新剧本版本
func (c *SecMasterClient) UpdatePlaybookVersionInvoker(request *model.UpdatePlaybookVersionRequest) *UpdatePlaybookVersionInvoker {
	requestDef := GenReqDefForUpdatePlaybookVersion()
	return &UpdatePlaybookVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRecipientsEmailStatus 更新收件人邮箱状态
//
// 更新收件人邮箱状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateRecipientsEmailStatus(request *model.UpdateRecipientsEmailStatusRequest) (*model.UpdateRecipientsEmailStatusResponse, error) {
	requestDef := GenReqDefForUpdateRecipientsEmailStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRecipientsEmailStatusResponse), nil
	}
}

// UpdateRecipientsEmailStatusInvoker 更新收件人邮箱状态
func (c *SecMasterClient) UpdateRecipientsEmailStatusInvoker(request *model.UpdateRecipientsEmailStatusRequest) *UpdateRecipientsEmailStatusInvoker {
	requestDef := GenReqDefForUpdateRecipientsEmailStatus()
	return &UpdateRecipientsEmailStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateReport 更新报告
//
// 更新报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateReport(request *model.UpdateReportRequest) (*model.UpdateReportResponse, error) {
	requestDef := GenReqDefForUpdateReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateReportResponse), nil
	}
}

// UpdateReportInvoker 更新报告
func (c *SecMasterClient) UpdateReportInvoker(request *model.UpdateReportRequest) *UpdateReportInvoker {
	requestDef := GenReqDefForUpdateReport()
	return &UpdateReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSearchCondition 修改检索条件
//
// 修改检索条件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateSearchCondition(request *model.UpdateSearchConditionRequest) (*model.UpdateSearchConditionResponse, error) {
	requestDef := GenReqDefForUpdateSearchCondition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSearchConditionResponse), nil
	}
}

// UpdateSearchConditionInvoker 修改检索条件
func (c *SecMasterClient) UpdateSearchConditionInvoker(request *model.UpdateSearchConditionRequest) *UpdateSearchConditionInvoker {
	requestDef := GenReqDefForUpdateSearchCondition()
	return &UpdateSearchConditionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTagValue 更新标签值
//
// 更新标签值
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateTagValue(request *model.UpdateTagValueRequest) (*model.UpdateTagValueResponse, error) {
	requestDef := GenReqDefForUpdateTagValue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTagValueResponse), nil
	}
}

// UpdateTagValueInvoker 更新标签值
func (c *SecMasterClient) UpdateTagValueInvoker(request *model.UpdateTagValueRequest) *UpdateTagValueInvoker {
	requestDef := GenReqDefForUpdateTagValue()
	return &UpdateTagValueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVpcEndpointService 更新VPC终端节点服务
//
// 更新VPC终端节点服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateVpcEndpointService(request *model.UpdateVpcEndpointServiceRequest) (*model.UpdateVpcEndpointServiceResponse, error) {
	requestDef := GenReqDefForUpdateVpcEndpointService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpcEndpointServiceResponse), nil
	}
}

// UpdateVpcEndpointServiceInvoker 更新VPC终端节点服务
func (c *SecMasterClient) UpdateVpcEndpointServiceInvoker(request *model.UpdateVpcEndpointServiceRequest) *UpdateVpcEndpointServiceInvoker {
	requestDef := GenReqDefForUpdateVpcEndpointService()
	return &UpdateVpcEndpointServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkspace 更新工作空间
//
// 更新工作空间名称、描述等信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateWorkspace(request *model.UpdateWorkspaceRequest) (*model.UpdateWorkspaceResponse, error) {
	requestDef := GenReqDefForUpdateWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkspaceResponse), nil
	}
}

// UpdateWorkspaceInvoker 更新工作空间
func (c *SecMasterClient) UpdateWorkspaceInvoker(request *model.UpdateWorkspaceRequest) *UpdateWorkspaceInvoker {
	requestDef := GenReqDefForUpdateWorkspace()
	return &UpdateWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadAttachment 上传附件
//
// 上传附件至OBS
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UploadAttachment(request *model.UploadAttachmentRequest) (*model.UploadAttachmentResponse, error) {
	requestDef := GenReqDefForUploadAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadAttachmentResponse), nil
	}
}

// UploadAttachmentInvoker 上传附件
func (c *SecMasterClient) UploadAttachmentInvoker(request *model.UploadAttachmentRequest) *UploadAttachmentInvoker {
	requestDef := GenReqDefForUploadAttachment()
	return &UploadAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ValidateAopWorkflowVersion 校验流程版本
//
// 校验流程版本.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ValidateAopWorkflowVersion(request *model.ValidateAopWorkflowVersionRequest) (*model.ValidateAopWorkflowVersionResponse, error) {
	requestDef := GenReqDefForValidateAopWorkflowVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateAopWorkflowVersionResponse), nil
	}
}

// ValidateAopWorkflowVersionInvoker 校验流程版本
func (c *SecMasterClient) ValidateAopWorkflowVersionInvoker(request *model.ValidateAopWorkflowVersionRequest) *ValidateAopWorkflowVersionInvoker {
	requestDef := GenReqDefForValidateAopWorkflowVersion()
	return &ValidateAopWorkflowVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
