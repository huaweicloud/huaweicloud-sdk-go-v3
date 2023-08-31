package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ces/v2/model"
)

type CesClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCesClient(hcClient *http_client.HcHttpClient) *CesClient {
	return &CesClient{HcClient: hcClient}
}

func CesClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AddAlarmRuleResources 批量增加告警规则资源
//
// 批量增加告警规则资源(资源分组类型的告警规则不支持)，资源分组类型的修改请使用资源分组管理相关接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) AddAlarmRuleResources(request *model.AddAlarmRuleResourcesRequest) (*model.AddAlarmRuleResourcesResponse, error) {
	requestDef := GenReqDefForAddAlarmRuleResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAlarmRuleResourcesResponse), nil
	}
}

// AddAlarmRuleResourcesInvoker 批量增加告警规则资源
func (c *CesClient) AddAlarmRuleResourcesInvoker(request *model.AddAlarmRuleResourcesRequest) *AddAlarmRuleResourcesInvoker {
	requestDef := GenReqDefForAddAlarmRuleResources()
	return &AddAlarmRuleResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateResources 自定义资源分组批量增加关联资源
//
// 给自定义资源分组,即类型为手动添加的资源分组,批量增加关联资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) BatchCreateResources(request *model.BatchCreateResourcesRequest) (*model.BatchCreateResourcesResponse, error) {
	requestDef := GenReqDefForBatchCreateResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateResourcesResponse), nil
	}
}

// BatchCreateResourcesInvoker 自定义资源分组批量增加关联资源
func (c *CesClient) BatchCreateResourcesInvoker(request *model.BatchCreateResourcesRequest) *BatchCreateResourcesInvoker {
	requestDef := GenReqDefForBatchCreateResources()
	return &BatchCreateResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteAlarmRules 批量删除告警规则
//
// 批量删除告警规则V2接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) BatchDeleteAlarmRules(request *model.BatchDeleteAlarmRulesRequest) (*model.BatchDeleteAlarmRulesResponse, error) {
	requestDef := GenReqDefForBatchDeleteAlarmRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAlarmRulesResponse), nil
	}
}

// BatchDeleteAlarmRulesInvoker 批量删除告警规则
func (c *CesClient) BatchDeleteAlarmRulesInvoker(request *model.BatchDeleteAlarmRulesRequest) *BatchDeleteAlarmRulesInvoker {
	requestDef := GenReqDefForBatchDeleteAlarmRules()
	return &BatchDeleteAlarmRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteAlarmTemplates 批量删除自定义告警模板
//
// 批量删除自定义告警模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) BatchDeleteAlarmTemplates(request *model.BatchDeleteAlarmTemplatesRequest) (*model.BatchDeleteAlarmTemplatesResponse, error) {
	requestDef := GenReqDefForBatchDeleteAlarmTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAlarmTemplatesResponse), nil
	}
}

// BatchDeleteAlarmTemplatesInvoker 批量删除自定义告警模板
func (c *CesClient) BatchDeleteAlarmTemplatesInvoker(request *model.BatchDeleteAlarmTemplatesRequest) *BatchDeleteAlarmTemplatesInvoker {
	requestDef := GenReqDefForBatchDeleteAlarmTemplates()
	return &BatchDeleteAlarmTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteResourceGroups 批量删除资源分组
//
// 批量删除资源分组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) BatchDeleteResourceGroups(request *model.BatchDeleteResourceGroupsRequest) (*model.BatchDeleteResourceGroupsResponse, error) {
	requestDef := GenReqDefForBatchDeleteResourceGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteResourceGroupsResponse), nil
	}
}

// BatchDeleteResourceGroupsInvoker 批量删除资源分组
func (c *CesClient) BatchDeleteResourceGroupsInvoker(request *model.BatchDeleteResourceGroupsRequest) *BatchDeleteResourceGroupsInvoker {
	requestDef := GenReqDefForBatchDeleteResourceGroups()
	return &BatchDeleteResourceGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteResources 自定义资源分组批量删除关联资源
//
// 给自定义资源分组,即类型为手动添加的资源分组,批量删除关联资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) BatchDeleteResources(request *model.BatchDeleteResourcesRequest) (*model.BatchDeleteResourcesResponse, error) {
	requestDef := GenReqDefForBatchDeleteResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteResourcesResponse), nil
	}
}

// BatchDeleteResourcesInvoker 自定义资源分组批量删除关联资源
func (c *CesClient) BatchDeleteResourcesInvoker(request *model.BatchDeleteResourcesRequest) *BatchDeleteResourcesInvoker {
	requestDef := GenReqDefForBatchDeleteResources()
	return &BatchDeleteResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchEnableAlarmRules 批量启停告警规则
//
// 批量启停告警规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) BatchEnableAlarmRules(request *model.BatchEnableAlarmRulesRequest) (*model.BatchEnableAlarmRulesResponse, error) {
	requestDef := GenReqDefForBatchEnableAlarmRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchEnableAlarmRulesResponse), nil
	}
}

// BatchEnableAlarmRulesInvoker 批量启停告警规则
func (c *CesClient) BatchEnableAlarmRulesInvoker(request *model.BatchEnableAlarmRulesRequest) *BatchEnableAlarmRulesInvoker {
	requestDef := GenReqDefForBatchEnableAlarmRules()
	return &BatchEnableAlarmRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateWidgets 批量更新监控视图
//
// 批量更新监控视图
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) BatchUpdateWidgets(request *model.BatchUpdateWidgetsRequest) (*model.BatchUpdateWidgetsResponse, error) {
	requestDef := GenReqDefForBatchUpdateWidgets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateWidgetsResponse), nil
	}
}

// BatchUpdateWidgetsInvoker 批量更新监控视图
func (c *CesClient) BatchUpdateWidgetsInvoker(request *model.BatchUpdateWidgetsRequest) *BatchUpdateWidgetsInvoker {
	requestDef := GenReqDefForBatchUpdateWidgets()
	return &BatchUpdateWidgetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAlarmRules 创建告警规则
//
// 创建告警规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) CreateAlarmRules(request *model.CreateAlarmRulesRequest) (*model.CreateAlarmRulesResponse, error) {
	requestDef := GenReqDefForCreateAlarmRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlarmRulesResponse), nil
	}
}

// CreateAlarmRulesInvoker 创建告警规则
func (c *CesClient) CreateAlarmRulesInvoker(request *model.CreateAlarmRulesRequest) *CreateAlarmRulesInvoker {
	requestDef := GenReqDefForCreateAlarmRules()
	return &CreateAlarmRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAlarmTemplate 创建自定义告警模板
//
// 创建自定义告警模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) CreateAlarmTemplate(request *model.CreateAlarmTemplateRequest) (*model.CreateAlarmTemplateResponse, error) {
	requestDef := GenReqDefForCreateAlarmTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlarmTemplateResponse), nil
	}
}

// CreateAlarmTemplateInvoker 创建自定义告警模板
func (c *CesClient) CreateAlarmTemplateInvoker(request *model.CreateAlarmTemplateRequest) *CreateAlarmTemplateInvoker {
	requestDef := GenReqDefForCreateAlarmTemplate()
	return &CreateAlarmTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDashboardWidgets 创建/复制/批量创建监控视图到指定的监控面板
//
// 创建/复制/批量创建监控视图到指定的监控面板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) CreateDashboardWidgets(request *model.CreateDashboardWidgetsRequest) (*model.CreateDashboardWidgetsResponse, error) {
	requestDef := GenReqDefForCreateDashboardWidgets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDashboardWidgetsResponse), nil
	}
}

// CreateDashboardWidgetsInvoker 创建/复制/批量创建监控视图到指定的监控面板
func (c *CesClient) CreateDashboardWidgetsInvoker(request *model.CreateDashboardWidgetsRequest) *CreateDashboardWidgetsInvoker {
	requestDef := GenReqDefForCreateDashboardWidgets()
	return &CreateDashboardWidgetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOneDashboard 创建/复制监控面板
//
// 创建/复制监控面板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) CreateOneDashboard(request *model.CreateOneDashboardRequest) (*model.CreateOneDashboardResponse, error) {
	requestDef := GenReqDefForCreateOneDashboard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOneDashboardResponse), nil
	}
}

// CreateOneDashboardInvoker 创建/复制监控面板
func (c *CesClient) CreateOneDashboardInvoker(request *model.CreateOneDashboardRequest) *CreateOneDashboardInvoker {
	requestDef := GenReqDefForCreateOneDashboard()
	return &CreateOneDashboardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResourceGroup 创建资源分组
//
// 创建资源分组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) CreateResourceGroup(request *model.CreateResourceGroupRequest) (*model.CreateResourceGroupResponse, error) {
	requestDef := GenReqDefForCreateResourceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResourceGroupResponse), nil
	}
}

// CreateResourceGroupInvoker 创建资源分组
func (c *CesClient) CreateResourceGroupInvoker(request *model.CreateResourceGroupRequest) *CreateResourceGroupInvoker {
	requestDef := GenReqDefForCreateResourceGroup()
	return &CreateResourceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAlarmRuleResources 批量删除告警规则资源
//
// 批量删除告警规则资源（资源分组类型的告警规则不支持），资源分组类型的修改请使用资源分组管理相关接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) DeleteAlarmRuleResources(request *model.DeleteAlarmRuleResourcesRequest) (*model.DeleteAlarmRuleResourcesResponse, error) {
	requestDef := GenReqDefForDeleteAlarmRuleResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlarmRuleResourcesResponse), nil
	}
}

// DeleteAlarmRuleResourcesInvoker 批量删除告警规则资源
func (c *CesClient) DeleteAlarmRuleResourcesInvoker(request *model.DeleteAlarmRuleResourcesRequest) *DeleteAlarmRuleResourcesInvoker {
	requestDef := GenReqDefForDeleteAlarmRuleResources()
	return &DeleteAlarmRuleResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDashboards 批量删除监控面板
//
// 批量删除监控面板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) DeleteDashboards(request *model.DeleteDashboardsRequest) (*model.DeleteDashboardsResponse, error) {
	requestDef := GenReqDefForDeleteDashboards()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDashboardsResponse), nil
	}
}

// DeleteDashboardsInvoker 批量删除监控面板
func (c *CesClient) DeleteDashboardsInvoker(request *model.DeleteDashboardsRequest) *DeleteDashboardsInvoker {
	requestDef := GenReqDefForDeleteDashboards()
	return &DeleteDashboardsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteOneWidget 删除指定监控视图
//
// 删除指定监控视图
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) DeleteOneWidget(request *model.DeleteOneWidgetRequest) (*model.DeleteOneWidgetResponse, error) {
	requestDef := GenReqDefForDeleteOneWidget()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteOneWidgetResponse), nil
	}
}

// DeleteOneWidgetInvoker 删除指定监控视图
func (c *CesClient) DeleteOneWidgetInvoker(request *model.DeleteOneWidgetRequest) *DeleteOneWidgetInvoker {
	requestDef := GenReqDefForDeleteOneWidget()
	return &DeleteOneWidgetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAgentDimensionInfo 查询主机监控维度指标信息
//
// 根据ECS/BMS资源ID查询磁盘、挂载点、进程、显卡、RAID控制器维度指标信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListAgentDimensionInfo(request *model.ListAgentDimensionInfoRequest) (*model.ListAgentDimensionInfoResponse, error) {
	requestDef := GenReqDefForListAgentDimensionInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAgentDimensionInfoResponse), nil
	}
}

// ListAgentDimensionInfoInvoker 查询主机监控维度指标信息
func (c *CesClient) ListAgentDimensionInfoInvoker(request *model.ListAgentDimensionInfoRequest) *ListAgentDimensionInfoInvoker {
	requestDef := GenReqDefForListAgentDimensionInfo()
	return &ListAgentDimensionInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmHistories 查询告警记录列表
//
// 查询告警记录列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListAlarmHistories(request *model.ListAlarmHistoriesRequest) (*model.ListAlarmHistoriesResponse, error) {
	requestDef := GenReqDefForListAlarmHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmHistoriesResponse), nil
	}
}

// ListAlarmHistoriesInvoker 查询告警记录列表
func (c *CesClient) ListAlarmHistoriesInvoker(request *model.ListAlarmHistoriesRequest) *ListAlarmHistoriesInvoker {
	requestDef := GenReqDefForListAlarmHistories()
	return &ListAlarmHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmRulePolicies 查询告警规则策略列表
//
// 根据告警规则ID查询策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListAlarmRulePolicies(request *model.ListAlarmRulePoliciesRequest) (*model.ListAlarmRulePoliciesResponse, error) {
	requestDef := GenReqDefForListAlarmRulePolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmRulePoliciesResponse), nil
	}
}

// ListAlarmRulePoliciesInvoker 查询告警规则策略列表
func (c *CesClient) ListAlarmRulePoliciesInvoker(request *model.ListAlarmRulePoliciesRequest) *ListAlarmRulePoliciesInvoker {
	requestDef := GenReqDefForListAlarmRulePolicies()
	return &ListAlarmRulePoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmRuleResources 查询告警规则资源列表
//
// 根据告警规则ID查询告警规则资源列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListAlarmRuleResources(request *model.ListAlarmRuleResourcesRequest) (*model.ListAlarmRuleResourcesResponse, error) {
	requestDef := GenReqDefForListAlarmRuleResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmRuleResourcesResponse), nil
	}
}

// ListAlarmRuleResourcesInvoker 查询告警规则资源列表
func (c *CesClient) ListAlarmRuleResourcesInvoker(request *model.ListAlarmRuleResourcesRequest) *ListAlarmRuleResourcesInvoker {
	requestDef := GenReqDefForListAlarmRuleResources()
	return &ListAlarmRuleResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmRules 查询告警规则列表
//
// 查询告警规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListAlarmRules(request *model.ListAlarmRulesRequest) (*model.ListAlarmRulesResponse, error) {
	requestDef := GenReqDefForListAlarmRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmRulesResponse), nil
	}
}

// ListAlarmRulesInvoker 查询告警规则列表
func (c *CesClient) ListAlarmRulesInvoker(request *model.ListAlarmRulesRequest) *ListAlarmRulesInvoker {
	requestDef := GenReqDefForListAlarmRules()
	return &ListAlarmRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmTemplateAssociationAlarms 查询告警模板关联的告警规则列表
//
// 查询告警模板关联的告警规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListAlarmTemplateAssociationAlarms(request *model.ListAlarmTemplateAssociationAlarmsRequest) (*model.ListAlarmTemplateAssociationAlarmsResponse, error) {
	requestDef := GenReqDefForListAlarmTemplateAssociationAlarms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmTemplateAssociationAlarmsResponse), nil
	}
}

// ListAlarmTemplateAssociationAlarmsInvoker 查询告警模板关联的告警规则列表
func (c *CesClient) ListAlarmTemplateAssociationAlarmsInvoker(request *model.ListAlarmTemplateAssociationAlarmsRequest) *ListAlarmTemplateAssociationAlarmsInvoker {
	requestDef := GenReqDefForListAlarmTemplateAssociationAlarms()
	return &ListAlarmTemplateAssociationAlarmsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmTemplates 查询告警模板列表
//
// 查询告警模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListAlarmTemplates(request *model.ListAlarmTemplatesRequest) (*model.ListAlarmTemplatesResponse, error) {
	requestDef := GenReqDefForListAlarmTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmTemplatesResponse), nil
	}
}

// ListAlarmTemplatesInvoker 查询告警模板列表
func (c *CesClient) ListAlarmTemplatesInvoker(request *model.ListAlarmTemplatesRequest) *ListAlarmTemplatesInvoker {
	requestDef := GenReqDefForListAlarmTemplates()
	return &ListAlarmTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDashboardInfos 查询监控面板列表
//
// 查询监控面板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListDashboardInfos(request *model.ListDashboardInfosRequest) (*model.ListDashboardInfosResponse, error) {
	requestDef := GenReqDefForListDashboardInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDashboardInfosResponse), nil
	}
}

// ListDashboardInfosInvoker 查询监控面板列表
func (c *CesClient) ListDashboardInfosInvoker(request *model.ListDashboardInfosRequest) *ListDashboardInfosInvoker {
	requestDef := GenReqDefForListDashboardInfos()
	return &ListDashboardInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDashboardWidgets 查询指定监控面板下的监控视图列表
//
// 查询指定监控面板下的监控视图列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListDashboardWidgets(request *model.ListDashboardWidgetsRequest) (*model.ListDashboardWidgetsResponse, error) {
	requestDef := GenReqDefForListDashboardWidgets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDashboardWidgetsResponse), nil
	}
}

// ListDashboardWidgetsInvoker 查询指定监控面板下的监控视图列表
func (c *CesClient) ListDashboardWidgetsInvoker(request *model.ListDashboardWidgetsRequest) *ListDashboardWidgetsInvoker {
	requestDef := GenReqDefForListDashboardWidgets()
	return &ListDashboardWidgetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceGroups 查询资源分组列表
//
// 查询资源分组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListResourceGroups(request *model.ListResourceGroupsRequest) (*model.ListResourceGroupsResponse, error) {
	requestDef := GenReqDefForListResourceGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceGroupsResponse), nil
	}
}

// ListResourceGroupsInvoker 查询资源分组列表
func (c *CesClient) ListResourceGroupsInvoker(request *model.ListResourceGroupsRequest) *ListResourceGroupsInvoker {
	requestDef := GenReqDefForListResourceGroups()
	return &ListResourceGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceGroupsServicesResources 查询资源分组下指定服务类别特定维度的资源列表
//
// 查询资源分组下指定服务类别特定维度的资源列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListResourceGroupsServicesResources(request *model.ListResourceGroupsServicesResourcesRequest) (*model.ListResourceGroupsServicesResourcesResponse, error) {
	requestDef := GenReqDefForListResourceGroupsServicesResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceGroupsServicesResourcesResponse), nil
	}
}

// ListResourceGroupsServicesResourcesInvoker 查询资源分组下指定服务类别特定维度的资源列表
func (c *CesClient) ListResourceGroupsServicesResourcesInvoker(request *model.ListResourceGroupsServicesResourcesRequest) *ListResourceGroupsServicesResourcesInvoker {
	requestDef := GenReqDefForListResourceGroupsServicesResources()
	return &ListResourceGroupsServicesResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlarmTemplate 查询告警模板详情
//
// 查询告警模板详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ShowAlarmTemplate(request *model.ShowAlarmTemplateRequest) (*model.ShowAlarmTemplateResponse, error) {
	requestDef := GenReqDefForShowAlarmTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlarmTemplateResponse), nil
	}
}

// ShowAlarmTemplateInvoker 查询告警模板详情
func (c *CesClient) ShowAlarmTemplateInvoker(request *model.ShowAlarmTemplateRequest) *ShowAlarmTemplateInvoker {
	requestDef := GenReqDefForShowAlarmTemplate()
	return &ShowAlarmTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceGroup 查询指定资源分组详情
//
// 查询指定资源分组详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ShowResourceGroup(request *model.ShowResourceGroupRequest) (*model.ShowResourceGroupResponse, error) {
	requestDef := GenReqDefForShowResourceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceGroupResponse), nil
	}
}

// ShowResourceGroupInvoker 查询指定资源分组详情
func (c *CesClient) ShowResourceGroupInvoker(request *model.ShowResourceGroupRequest) *ShowResourceGroupInvoker {
	requestDef := GenReqDefForShowResourceGroup()
	return &ShowResourceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWidget 查询指定监控视图信息
//
// 查询指定监控视图信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ShowWidget(request *model.ShowWidgetRequest) (*model.ShowWidgetResponse, error) {
	requestDef := GenReqDefForShowWidget()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWidgetResponse), nil
	}
}

// ShowWidgetInvoker 查询指定监控视图信息
func (c *CesClient) ShowWidgetInvoker(request *model.ShowWidgetRequest) *ShowWidgetInvoker {
	requestDef := GenReqDefForShowWidget()
	return &ShowWidgetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAlarmRulePolicies 修改告警规则策略(全量修改)
//
// 修改告警规则策略(全量修改)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) UpdateAlarmRulePolicies(request *model.UpdateAlarmRulePoliciesRequest) (*model.UpdateAlarmRulePoliciesResponse, error) {
	requestDef := GenReqDefForUpdateAlarmRulePolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlarmRulePoliciesResponse), nil
	}
}

// UpdateAlarmRulePoliciesInvoker 修改告警规则策略(全量修改)
func (c *CesClient) UpdateAlarmRulePoliciesInvoker(request *model.UpdateAlarmRulePoliciesRequest) *UpdateAlarmRulePoliciesInvoker {
	requestDef := GenReqDefForUpdateAlarmRulePolicies()
	return &UpdateAlarmRulePoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAlarmTemplate 修改自定义告警模板
//
// 修改自定义告警模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) UpdateAlarmTemplate(request *model.UpdateAlarmTemplateRequest) (*model.UpdateAlarmTemplateResponse, error) {
	requestDef := GenReqDefForUpdateAlarmTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlarmTemplateResponse), nil
	}
}

// UpdateAlarmTemplateInvoker 修改自定义告警模板
func (c *CesClient) UpdateAlarmTemplateInvoker(request *model.UpdateAlarmTemplateRequest) *UpdateAlarmTemplateInvoker {
	requestDef := GenReqDefForUpdateAlarmTemplate()
	return &UpdateAlarmTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDashboard 修改监控面板
//
// 修改监控面板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) UpdateDashboard(request *model.UpdateDashboardRequest) (*model.UpdateDashboardResponse, error) {
	requestDef := GenReqDefForUpdateDashboard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDashboardResponse), nil
	}
}

// UpdateDashboardInvoker 修改监控面板
func (c *CesClient) UpdateDashboardInvoker(request *model.UpdateDashboardRequest) *UpdateDashboardInvoker {
	requestDef := GenReqDefForUpdateDashboard()
	return &UpdateDashboardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateResourceGroup 修改资源分组
//
// 修改资源分组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) UpdateResourceGroup(request *model.UpdateResourceGroupRequest) (*model.UpdateResourceGroupResponse, error) {
	requestDef := GenReqDefForUpdateResourceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResourceGroupResponse), nil
	}
}

// UpdateResourceGroupInvoker 修改资源分组
func (c *CesClient) UpdateResourceGroupInvoker(request *model.UpdateResourceGroupRequest) *UpdateResourceGroupInvoker {
	requestDef := GenReqDefForUpdateResourceGroup()
	return &UpdateResourceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
