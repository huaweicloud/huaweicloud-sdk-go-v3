package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/sa/v2/model"
)

type SaClient struct {
	HcClient *http_client.HcHttpClient
}

func NewSaClient(hcClient *http_client.HcHttpClient) *SaClient {
	return &SaClient{HcClient: hcClient}
}

func SaClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateAlertRule 创建告警规则（仅支持华东-上海一使用）
//
// # Create alert rule
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) CreateAlertRule(request *model.CreateAlertRuleRequest) (*model.CreateAlertRuleResponse, error) {
	requestDef := GenReqDefForCreateAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlertRuleResponse), nil
	}
}

// CreateAlertRuleInvoker 创建告警规则（仅支持华东-上海一使用）
func (c *SaClient) CreateAlertRuleInvoker(request *model.CreateAlertRuleRequest) *CreateAlertRuleInvoker {
	requestDef := GenReqDefForCreateAlertRule()
	return &CreateAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAlertRuleSimulation 模拟告警规则（仅支持华东-上海一使用）
//
// # Simulate alert rule
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) CreateAlertRuleSimulation(request *model.CreateAlertRuleSimulationRequest) (*model.CreateAlertRuleSimulationResponse, error) {
	requestDef := GenReqDefForCreateAlertRuleSimulation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlertRuleSimulationResponse), nil
	}
}

// CreateAlertRuleSimulationInvoker 模拟告警规则（仅支持华东-上海一使用）
func (c *SaClient) CreateAlertRuleSimulationInvoker(request *model.CreateAlertRuleSimulationRequest) *CreateAlertRuleSimulationInvoker {
	requestDef := GenReqDefForCreateAlertRuleSimulation()
	return &CreateAlertRuleSimulationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAlertRule 删除告警规则（仅支持华东-上海一使用）
//
// # Delete alert rule
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) DeleteAlertRule(request *model.DeleteAlertRuleRequest) (*model.DeleteAlertRuleResponse, error) {
	requestDef := GenReqDefForDeleteAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlertRuleResponse), nil
	}
}

// DeleteAlertRuleInvoker 删除告警规则（仅支持华东-上海一使用）
func (c *SaClient) DeleteAlertRuleInvoker(request *model.DeleteAlertRuleRequest) *DeleteAlertRuleInvoker {
	requestDef := GenReqDefForDeleteAlertRule()
	return &DeleteAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableAlertRule 停用告警规则（仅支持华东-上海一使用）
//
// # Disable alert rule
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) DisableAlertRule(request *model.DisableAlertRuleRequest) (*model.DisableAlertRuleResponse, error) {
	requestDef := GenReqDefForDisableAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableAlertRuleResponse), nil
	}
}

// DisableAlertRuleInvoker 停用告警规则（仅支持华东-上海一使用）
func (c *SaClient) DisableAlertRuleInvoker(request *model.DisableAlertRuleRequest) *DisableAlertRuleInvoker {
	requestDef := GenReqDefForDisableAlertRule()
	return &DisableAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableAlertRule 启用告警规则（仅支持华东-上海一使用）
//
// # Enable alert rule
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) EnableAlertRule(request *model.EnableAlertRuleRequest) (*model.EnableAlertRuleResponse, error) {
	requestDef := GenReqDefForEnableAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableAlertRuleResponse), nil
	}
}

// EnableAlertRuleInvoker 启用告警规则（仅支持华东-上海一使用）
func (c *SaClient) EnableAlertRuleInvoker(request *model.EnableAlertRuleRequest) *EnableAlertRuleInvoker {
	requestDef := GenReqDefForEnableAlertRule()
	return &EnableAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlertRuleMetrics 告警规则总览（仅支持华东-上海一使用）
//
// # List alert rule metrics
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ListAlertRuleMetrics(request *model.ListAlertRuleMetricsRequest) (*model.ListAlertRuleMetricsResponse, error) {
	requestDef := GenReqDefForListAlertRuleMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlertRuleMetricsResponse), nil
	}
}

// ListAlertRuleMetricsInvoker 告警规则总览（仅支持华东-上海一使用）
func (c *SaClient) ListAlertRuleMetricsInvoker(request *model.ListAlertRuleMetricsRequest) *ListAlertRuleMetricsInvoker {
	requestDef := GenReqDefForListAlertRuleMetrics()
	return &ListAlertRuleMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlertRuleTemplates 列出告警规则模板（仅支持华东-上海一使用）
//
// # List alert rule templates
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ListAlertRuleTemplates(request *model.ListAlertRuleTemplatesRequest) (*model.ListAlertRuleTemplatesResponse, error) {
	requestDef := GenReqDefForListAlertRuleTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlertRuleTemplatesResponse), nil
	}
}

// ListAlertRuleTemplatesInvoker 列出告警规则模板（仅支持华东-上海一使用）
func (c *SaClient) ListAlertRuleTemplatesInvoker(request *model.ListAlertRuleTemplatesRequest) *ListAlertRuleTemplatesInvoker {
	requestDef := GenReqDefForListAlertRuleTemplates()
	return &ListAlertRuleTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlertRules 列出告警规则（仅支持华东-上海一使用）
//
// # List alert rules
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ListAlertRules(request *model.ListAlertRulesRequest) (*model.ListAlertRulesResponse, error) {
	requestDef := GenReqDefForListAlertRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlertRulesResponse), nil
	}
}

// ListAlertRulesInvoker 列出告警规则（仅支持华东-上海一使用）
func (c *SaClient) ListAlertRulesInvoker(request *model.ListAlertRulesRequest) *ListAlertRulesInvoker {
	requestDef := GenReqDefForListAlertRules()
	return &ListAlertRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlertRule 查看告警规则（仅支持华东-上海一使用）
//
// # Get alert rule
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ShowAlertRule(request *model.ShowAlertRuleRequest) (*model.ShowAlertRuleResponse, error) {
	requestDef := GenReqDefForShowAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlertRuleResponse), nil
	}
}

// ShowAlertRuleInvoker 查看告警规则（仅支持华东-上海一使用）
func (c *SaClient) ShowAlertRuleInvoker(request *model.ShowAlertRuleRequest) *ShowAlertRuleInvoker {
	requestDef := GenReqDefForShowAlertRule()
	return &ShowAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlertRuleTemplate 查看告警规则模板（仅支持华东-上海一使用）
//
// # List alert rule templates
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ShowAlertRuleTemplate(request *model.ShowAlertRuleTemplateRequest) (*model.ShowAlertRuleTemplateResponse, error) {
	requestDef := GenReqDefForShowAlertRuleTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlertRuleTemplateResponse), nil
	}
}

// ShowAlertRuleTemplateInvoker 查看告警规则模板（仅支持华东-上海一使用）
func (c *SaClient) ShowAlertRuleTemplateInvoker(request *model.ShowAlertRuleTemplateRequest) *ShowAlertRuleTemplateInvoker {
	requestDef := GenReqDefForShowAlertRuleTemplate()
	return &ShowAlertRuleTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAlertRule 更新告警规则（仅支持华东-上海一使用）
//
// # Update alert rule
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) UpdateAlertRule(request *model.UpdateAlertRuleRequest) (*model.UpdateAlertRuleResponse, error) {
	requestDef := GenReqDefForUpdateAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlertRuleResponse), nil
	}
}

// UpdateAlertRuleInvoker 更新告警规则（仅支持华东-上海一使用）
func (c *SaClient) UpdateAlertRuleInvoker(request *model.UpdateAlertRuleRequest) *UpdateAlertRuleInvoker {
	requestDef := GenReqDefForUpdateAlertRule()
	return &UpdateAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeAlert 更新告警（仅支持华东-上海一使用）
//
// 编辑告警，根据实际修改的属性更新，未修改的列不更新
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ChangeAlert(request *model.ChangeAlertRequest) (*model.ChangeAlertResponse, error) {
	requestDef := GenReqDefForChangeAlert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeAlertResponse), nil
	}
}

// ChangeAlertInvoker 更新告警（仅支持华东-上海一使用）
func (c *SaClient) ChangeAlertInvoker(request *model.ChangeAlertRequest) *ChangeAlertInvoker {
	requestDef := GenReqDefForChangeAlert()
	return &ChangeAlertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAlert 创建告警（仅支持华东-上海一使用）
//
// 创建告警
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) CreateAlert(request *model.CreateAlertRequest) (*model.CreateAlertResponse, error) {
	requestDef := GenReqDefForCreateAlert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlertResponse), nil
	}
}

// CreateAlertInvoker 创建告警（仅支持华东-上海一使用）
func (c *SaClient) CreateAlertInvoker(request *model.CreateAlertRequest) *CreateAlertInvoker {
	requestDef := GenReqDefForCreateAlert()
	return &CreateAlertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBatchOrderAlerts 告警转事件（仅支持华东-上海一使用）
//
// 告警转事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) CreateBatchOrderAlerts(request *model.CreateBatchOrderAlertsRequest) (*model.CreateBatchOrderAlertsResponse, error) {
	requestDef := GenReqDefForCreateBatchOrderAlerts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBatchOrderAlertsResponse), nil
	}
}

// CreateBatchOrderAlertsInvoker 告警转事件（仅支持华东-上海一使用）
func (c *SaClient) CreateBatchOrderAlertsInvoker(request *model.CreateBatchOrderAlertsRequest) *CreateBatchOrderAlertsInvoker {
	requestDef := GenReqDefForCreateBatchOrderAlerts()
	return &CreateBatchOrderAlertsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAlert 删除告警（仅支持华东-上海一使用）
//
// 删除告警
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) DeleteAlert(request *model.DeleteAlertRequest) (*model.DeleteAlertResponse, error) {
	requestDef := GenReqDefForDeleteAlert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlertResponse), nil
	}
}

// DeleteAlertInvoker 删除告警（仅支持华东-上海一使用）
func (c *SaClient) DeleteAlertInvoker(request *model.DeleteAlertRequest) *DeleteAlertInvoker {
	requestDef := GenReqDefForDeleteAlert()
	return &DeleteAlertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlerts 搜索告警列表（仅支持华东-上海一使用）
//
// 搜索告警列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ListAlerts(request *model.ListAlertsRequest) (*model.ListAlertsResponse, error) {
	requestDef := GenReqDefForListAlerts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlertsResponse), nil
	}
}

// ListAlertsInvoker 搜索告警列表（仅支持华东-上海一使用）
func (c *SaClient) ListAlertsInvoker(request *model.ListAlertsRequest) *ListAlertsInvoker {
	requestDef := GenReqDefForListAlerts()
	return &ListAlertsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlert 获取告警详情（仅支持华东-上海一使用）
//
// 获取告警详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ShowAlert(request *model.ShowAlertRequest) (*model.ShowAlertResponse, error) {
	requestDef := GenReqDefForShowAlert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlertResponse), nil
	}
}

// ShowAlertInvoker 获取告警详情（仅支持华东-上海一使用）
func (c *SaClient) ShowAlertInvoker(request *model.ShowAlertRequest) *ShowAlertInvoker {
	requestDef := GenReqDefForShowAlert()
	return &ShowAlertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeIncident 更新事件（仅支持华东-上海一使用）
//
// 编辑事件，根据实际修改的属性更新，未修改的列不更新
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ChangeIncident(request *model.ChangeIncidentRequest) (*model.ChangeIncidentResponse, error) {
	requestDef := GenReqDefForChangeIncident()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeIncidentResponse), nil
	}
}

// ChangeIncidentInvoker 更新事件（仅支持华东-上海一使用）
func (c *SaClient) ChangeIncidentInvoker(request *model.ChangeIncidentRequest) *ChangeIncidentInvoker {
	requestDef := GenReqDefForChangeIncident()
	return &ChangeIncidentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateIncident 创建事件（仅支持华东-上海一使用）
//
// 创建事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) CreateIncident(request *model.CreateIncidentRequest) (*model.CreateIncidentResponse, error) {
	requestDef := GenReqDefForCreateIncident()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIncidentResponse), nil
	}
}

// CreateIncidentInvoker 创建事件（仅支持华东-上海一使用）
func (c *SaClient) CreateIncidentInvoker(request *model.CreateIncidentRequest) *CreateIncidentInvoker {
	requestDef := GenReqDefForCreateIncident()
	return &CreateIncidentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteIncident 创建事件（仅支持华东-上海一使用）
//
// 创建事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) DeleteIncident(request *model.DeleteIncidentRequest) (*model.DeleteIncidentResponse, error) {
	requestDef := GenReqDefForDeleteIncident()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIncidentResponse), nil
	}
}

// DeleteIncidentInvoker 创建事件（仅支持华东-上海一使用）
func (c *SaClient) DeleteIncidentInvoker(request *model.DeleteIncidentRequest) *DeleteIncidentInvoker {
	requestDef := GenReqDefForDeleteIncident()
	return &DeleteIncidentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIncidentTypes 获取事件的类型列表（仅支持华东-上海一使用）
//
// 获取事件的类型列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ListIncidentTypes(request *model.ListIncidentTypesRequest) (*model.ListIncidentTypesResponse, error) {
	requestDef := GenReqDefForListIncidentTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIncidentTypesResponse), nil
	}
}

// ListIncidentTypesInvoker 获取事件的类型列表（仅支持华东-上海一使用）
func (c *SaClient) ListIncidentTypesInvoker(request *model.ListIncidentTypesRequest) *ListIncidentTypesInvoker {
	requestDef := GenReqDefForListIncidentTypes()
	return &ListIncidentTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIncidents 搜索事件列表（仅支持华东-上海一使用）
//
// 搜索事件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ListIncidents(request *model.ListIncidentsRequest) (*model.ListIncidentsResponse, error) {
	requestDef := GenReqDefForListIncidents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIncidentsResponse), nil
	}
}

// ListIncidentsInvoker 搜索事件列表（仅支持华东-上海一使用）
func (c *SaClient) ListIncidentsInvoker(request *model.ListIncidentsRequest) *ListIncidentsInvoker {
	requestDef := GenReqDefForListIncidents()
	return &ListIncidentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIncident 获取事件详情（仅支持华东-上海一使用）
//
// 获取事件详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ShowIncident(request *model.ShowIncidentRequest) (*model.ShowIncidentResponse, error) {
	requestDef := GenReqDefForShowIncident()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIncidentResponse), nil
	}
}

// ShowIncidentInvoker 获取事件详情（仅支持华东-上海一使用）
func (c *SaClient) ShowIncidentInvoker(request *model.ShowIncidentRequest) *ShowIncidentInvoker {
	requestDef := GenReqDefForShowIncident()
	return &ShowIncidentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateIndicator 创建指标
//
// 创建指标（仅支持华东-上海一使用）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) CreateIndicator(request *model.CreateIndicatorRequest) (*model.CreateIndicatorResponse, error) {
	requestDef := GenReqDefForCreateIndicator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIndicatorResponse), nil
	}
}

// CreateIndicatorInvoker 创建指标
func (c *SaClient) CreateIndicatorInvoker(request *model.CreateIndicatorRequest) *CreateIndicatorInvoker {
	requestDef := GenReqDefForCreateIndicator()
	return &CreateIndicatorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteIndicator 删除指标
//
// 删除指标（仅支持华东-上海一使用）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) DeleteIndicator(request *model.DeleteIndicatorRequest) (*model.DeleteIndicatorResponse, error) {
	requestDef := GenReqDefForDeleteIndicator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIndicatorResponse), nil
	}
}

// DeleteIndicatorInvoker 删除指标
func (c *SaClient) DeleteIndicatorInvoker(request *model.DeleteIndicatorRequest) *DeleteIndicatorInvoker {
	requestDef := GenReqDefForDeleteIndicator()
	return &DeleteIndicatorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIndicators 查询指标列表（仅支持华东-上海一使用）
//
// # List all indicators
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ListIndicators(request *model.ListIndicatorsRequest) (*model.ListIndicatorsResponse, error) {
	requestDef := GenReqDefForListIndicators()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIndicatorsResponse), nil
	}
}

// ListIndicatorsInvoker 查询指标列表（仅支持华东-上海一使用）
func (c *SaClient) ListIndicatorsInvoker(request *model.ListIndicatorsRequest) *ListIndicatorsInvoker {
	requestDef := GenReqDefForListIndicators()
	return &ListIndicatorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIndicatorDetail 查询指标详情
//
// 查询指标详情（仅支持华东-上海一使用）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ShowIndicatorDetail(request *model.ShowIndicatorDetailRequest) (*model.ShowIndicatorDetailResponse, error) {
	requestDef := GenReqDefForShowIndicatorDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIndicatorDetailResponse), nil
	}
}

// ShowIndicatorDetailInvoker 查询指标详情
func (c *SaClient) ShowIndicatorDetailInvoker(request *model.ShowIndicatorDetailRequest) *ShowIndicatorDetailInvoker {
	requestDef := GenReqDefForShowIndicatorDetail()
	return &ShowIndicatorDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIndicator 更新指标
//
// 更新指标（仅支持华东-上海一使用）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) UpdateIndicator(request *model.UpdateIndicatorRequest) (*model.UpdateIndicatorResponse, error) {
	requestDef := GenReqDefForUpdateIndicator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIndicatorResponse), nil
	}
}

// UpdateIndicatorInvoker 更新指标
func (c *SaClient) UpdateIndicatorInvoker(request *model.UpdateIndicatorRequest) *UpdateIndicatorInvoker {
	requestDef := GenReqDefForUpdateIndicator()
	return &UpdateIndicatorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePlaybook 创建剧本（仅支持华东-上海一使用）
//
// Create playbook.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) CreatePlaybook(request *model.CreatePlaybookRequest) (*model.CreatePlaybookResponse, error) {
	requestDef := GenReqDefForCreatePlaybook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePlaybookResponse), nil
	}
}

// CreatePlaybookInvoker 创建剧本（仅支持华东-上海一使用）
func (c *SaClient) CreatePlaybookInvoker(request *model.CreatePlaybookRequest) *CreatePlaybookInvoker {
	requestDef := GenReqDefForCreatePlaybook()
	return &CreatePlaybookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePlaybook 删除剧本（仅支持华东-上海一使用）
//
// Delete playbook.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) DeletePlaybook(request *model.DeletePlaybookRequest) (*model.DeletePlaybookResponse, error) {
	requestDef := GenReqDefForDeletePlaybook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePlaybookResponse), nil
	}
}

// DeletePlaybookInvoker 删除剧本（仅支持华东-上海一使用）
func (c *SaClient) DeletePlaybookInvoker(request *model.DeletePlaybookRequest) *DeletePlaybookInvoker {
	requestDef := GenReqDefForDeletePlaybook()
	return &DeletePlaybookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlaybooks 查询剧本列表（仅支持华东-上海一使用）
//
// List all playbooks.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ListPlaybooks(request *model.ListPlaybooksRequest) (*model.ListPlaybooksResponse, error) {
	requestDef := GenReqDefForListPlaybooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlaybooksResponse), nil
	}
}

// ListPlaybooksInvoker 查询剧本列表（仅支持华东-上海一使用）
func (c *SaClient) ListPlaybooksInvoker(request *model.ListPlaybooksRequest) *ListPlaybooksInvoker {
	requestDef := GenReqDefForListPlaybooks()
	return &ListPlaybooksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybook 查询剧本详情（仅支持华东-上海一使用）
//
// # Show playbook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ShowPlaybook(request *model.ShowPlaybookRequest) (*model.ShowPlaybookResponse, error) {
	requestDef := GenReqDefForShowPlaybook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookResponse), nil
	}
}

// ShowPlaybookInvoker 查询剧本详情（仅支持华东-上海一使用）
func (c *SaClient) ShowPlaybookInvoker(request *model.ShowPlaybookRequest) *ShowPlaybookInvoker {
	requestDef := GenReqDefForShowPlaybook()
	return &ShowPlaybookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybookMonitors 剧本运行监控（仅支持华东-上海一使用）
//
// 剧本运行监控
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ShowPlaybookMonitors(request *model.ShowPlaybookMonitorsRequest) (*model.ShowPlaybookMonitorsResponse, error) {
	requestDef := GenReqDefForShowPlaybookMonitors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookMonitorsResponse), nil
	}
}

// ShowPlaybookMonitorsInvoker 剧本运行监控（仅支持华东-上海一使用）
func (c *SaClient) ShowPlaybookMonitorsInvoker(request *model.ShowPlaybookMonitorsRequest) *ShowPlaybookMonitorsInvoker {
	requestDef := GenReqDefForShowPlaybookMonitors()
	return &ShowPlaybookMonitorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybookStatistics 剧本数据统计（仅支持华东-上海一使用）
//
// 剧本统计数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ShowPlaybookStatistics(request *model.ShowPlaybookStatisticsRequest) (*model.ShowPlaybookStatisticsResponse, error) {
	requestDef := GenReqDefForShowPlaybookStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookStatisticsResponse), nil
	}
}

// ShowPlaybookStatisticsInvoker 剧本数据统计（仅支持华东-上海一使用）
func (c *SaClient) ShowPlaybookStatisticsInvoker(request *model.ShowPlaybookStatisticsRequest) *ShowPlaybookStatisticsInvoker {
	requestDef := GenReqDefForShowPlaybookStatistics()
	return &ShowPlaybookStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePlaybook 修改剧本（仅支持华东-上海一使用）
//
// Update playbook.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) UpdatePlaybook(request *model.UpdatePlaybookRequest) (*model.UpdatePlaybookResponse, error) {
	requestDef := GenReqDefForUpdatePlaybook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePlaybookResponse), nil
	}
}

// UpdatePlaybookInvoker 修改剧本（仅支持华东-上海一使用）
func (c *SaClient) UpdatePlaybookInvoker(request *model.UpdatePlaybookRequest) *UpdatePlaybookInvoker {
	requestDef := GenReqDefForUpdatePlaybook()
	return &UpdatePlaybookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePlaybookAction 创建剧本动作（仅支持华东-上海一使用）
//
// Create action.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) CreatePlaybookAction(request *model.CreatePlaybookActionRequest) (*model.CreatePlaybookActionResponse, error) {
	requestDef := GenReqDefForCreatePlaybookAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePlaybookActionResponse), nil
	}
}

// CreatePlaybookActionInvoker 创建剧本动作（仅支持华东-上海一使用）
func (c *SaClient) CreatePlaybookActionInvoker(request *model.CreatePlaybookActionRequest) *CreatePlaybookActionInvoker {
	requestDef := GenReqDefForCreatePlaybookAction()
	return &CreatePlaybookActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePlaybookAction 删除剧本动作（仅支持华东-上海一使用）
//
// Delete action.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) DeletePlaybookAction(request *model.DeletePlaybookActionRequest) (*model.DeletePlaybookActionResponse, error) {
	requestDef := GenReqDefForDeletePlaybookAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePlaybookActionResponse), nil
	}
}

// DeletePlaybookActionInvoker 删除剧本动作（仅支持华东-上海一使用）
func (c *SaClient) DeletePlaybookActionInvoker(request *model.DeletePlaybookActionRequest) *DeletePlaybookActionInvoker {
	requestDef := GenReqDefForDeletePlaybookAction()
	return &DeletePlaybookActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlaybookActions 查询剧本动作（仅支持华东-上海一使用）
//
// List all actions.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ListPlaybookActions(request *model.ListPlaybookActionsRequest) (*model.ListPlaybookActionsResponse, error) {
	requestDef := GenReqDefForListPlaybookActions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlaybookActionsResponse), nil
	}
}

// ListPlaybookActionsInvoker 查询剧本动作（仅支持华东-上海一使用）
func (c *SaClient) ListPlaybookActionsInvoker(request *model.ListPlaybookActionsRequest) *ListPlaybookActionsInvoker {
	requestDef := GenReqDefForListPlaybookActions()
	return &ListPlaybookActionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePlaybookAction 更新剧本动作（仅支持华东-上海一使用）
//
// Update action.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) UpdatePlaybookAction(request *model.UpdatePlaybookActionRequest) (*model.UpdatePlaybookActionResponse, error) {
	requestDef := GenReqDefForUpdatePlaybookAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePlaybookActionResponse), nil
	}
}

// UpdatePlaybookActionInvoker 更新剧本动作（仅支持华东-上海一使用）
func (c *SaClient) UpdatePlaybookActionInvoker(request *model.UpdatePlaybookActionRequest) *UpdatePlaybookActionInvoker {
	requestDef := GenReqDefForUpdatePlaybookAction()
	return &UpdatePlaybookActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePlaybookApprove 审核剧本（仅支持华东-上海一使用）
//
// Create playbook approve.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) CreatePlaybookApprove(request *model.CreatePlaybookApproveRequest) (*model.CreatePlaybookApproveResponse, error) {
	requestDef := GenReqDefForCreatePlaybookApprove()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePlaybookApproveResponse), nil
	}
}

// CreatePlaybookApproveInvoker 审核剧本（仅支持华东-上海一使用）
func (c *SaClient) CreatePlaybookApproveInvoker(request *model.CreatePlaybookApproveRequest) *CreatePlaybookApproveInvoker {
	requestDef := GenReqDefForCreatePlaybookApprove()
	return &CreatePlaybookApproveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlaybookApproves 查询剧本审核结果（仅支持华东-上海一使用）
//
// List approves.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ListPlaybookApproves(request *model.ListPlaybookApprovesRequest) (*model.ListPlaybookApprovesResponse, error) {
	requestDef := GenReqDefForListPlaybookApproves()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlaybookApprovesResponse), nil
	}
}

// ListPlaybookApprovesInvoker 查询剧本审核结果（仅支持华东-上海一使用）
func (c *SaClient) ListPlaybookApprovesInvoker(request *model.ListPlaybookApprovesRequest) *ListPlaybookApprovesInvoker {
	requestDef := GenReqDefForListPlaybookApproves()
	return &ListPlaybookApprovesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangePlaybookInstance 操作剧本实例（仅支持华东-上海一使用）
//
// Operation Playbook Instance.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ChangePlaybookInstance(request *model.ChangePlaybookInstanceRequest) (*model.ChangePlaybookInstanceResponse, error) {
	requestDef := GenReqDefForChangePlaybookInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangePlaybookInstanceResponse), nil
	}
}

// ChangePlaybookInstanceInvoker 操作剧本实例（仅支持华东-上海一使用）
func (c *SaClient) ChangePlaybookInstanceInvoker(request *model.ChangePlaybookInstanceRequest) *ChangePlaybookInstanceInvoker {
	requestDef := GenReqDefForChangePlaybookInstance()
	return &ChangePlaybookInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlaybookAuditLogs 查询剧本实例审计日志（仅支持华东-上海一使用）
//
// List audit logs.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ListPlaybookAuditLogs(request *model.ListPlaybookAuditLogsRequest) (*model.ListPlaybookAuditLogsResponse, error) {
	requestDef := GenReqDefForListPlaybookAuditLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlaybookAuditLogsResponse), nil
	}
}

// ListPlaybookAuditLogsInvoker 查询剧本实例审计日志（仅支持华东-上海一使用）
func (c *SaClient) ListPlaybookAuditLogsInvoker(request *model.ListPlaybookAuditLogsRequest) *ListPlaybookAuditLogsInvoker {
	requestDef := GenReqDefForListPlaybookAuditLogs()
	return &ListPlaybookAuditLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlaybookInstances 查询剧本实例列表（仅支持华东-上海一使用）
//
// # List playbook instances
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ListPlaybookInstances(request *model.ListPlaybookInstancesRequest) (*model.ListPlaybookInstancesResponse, error) {
	requestDef := GenReqDefForListPlaybookInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlaybookInstancesResponse), nil
	}
}

// ListPlaybookInstancesInvoker 查询剧本实例列表（仅支持华东-上海一使用）
func (c *SaClient) ListPlaybookInstancesInvoker(request *model.ListPlaybookInstancesRequest) *ListPlaybookInstancesInvoker {
	requestDef := GenReqDefForListPlaybookInstances()
	return &ListPlaybookInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybookInstance 查询剧本实例详情（仅支持华东-上海一使用）
//
// # Show playbook instance
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ShowPlaybookInstance(request *model.ShowPlaybookInstanceRequest) (*model.ShowPlaybookInstanceResponse, error) {
	requestDef := GenReqDefForShowPlaybookInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookInstanceResponse), nil
	}
}

// ShowPlaybookInstanceInvoker 查询剧本实例详情（仅支持华东-上海一使用）
func (c *SaClient) ShowPlaybookInstanceInvoker(request *model.ShowPlaybookInstanceRequest) *ShowPlaybookInstanceInvoker {
	requestDef := GenReqDefForShowPlaybookInstance()
	return &ShowPlaybookInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybookTopology 查询剧本拓扑关系（仅支持华东-上海一使用）
//
// # Show playbook Topology
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ShowPlaybookTopology(request *model.ShowPlaybookTopologyRequest) (*model.ShowPlaybookTopologyResponse, error) {
	requestDef := GenReqDefForShowPlaybookTopology()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookTopologyResponse), nil
	}
}

// ShowPlaybookTopologyInvoker 查询剧本拓扑关系（仅支持华东-上海一使用）
func (c *SaClient) ShowPlaybookTopologyInvoker(request *model.ShowPlaybookTopologyRequest) *ShowPlaybookTopologyInvoker {
	requestDef := GenReqDefForShowPlaybookTopology()
	return &ShowPlaybookTopologyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePlaybookRule 创建剧本规则（仅支持华东-上海一使用）
//
// Create rule.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) CreatePlaybookRule(request *model.CreatePlaybookRuleRequest) (*model.CreatePlaybookRuleResponse, error) {
	requestDef := GenReqDefForCreatePlaybookRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePlaybookRuleResponse), nil
	}
}

// CreatePlaybookRuleInvoker 创建剧本规则（仅支持华东-上海一使用）
func (c *SaClient) CreatePlaybookRuleInvoker(request *model.CreatePlaybookRuleRequest) *CreatePlaybookRuleInvoker {
	requestDef := GenReqDefForCreatePlaybookRule()
	return &CreatePlaybookRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePlaybookRule 删除剧本规则（仅支持华东-上海一使用）
//
// Delete rule.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) DeletePlaybookRule(request *model.DeletePlaybookRuleRequest) (*model.DeletePlaybookRuleResponse, error) {
	requestDef := GenReqDefForDeletePlaybookRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePlaybookRuleResponse), nil
	}
}

// DeletePlaybookRuleInvoker 删除剧本规则（仅支持华东-上海一使用）
func (c *SaClient) DeletePlaybookRuleInvoker(request *model.DeletePlaybookRuleRequest) *DeletePlaybookRuleInvoker {
	requestDef := GenReqDefForDeletePlaybookRule()
	return &DeletePlaybookRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybookRule 查询剧本规则详情（仅支持华东-上海一使用）
//
// Show rule formation.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ShowPlaybookRule(request *model.ShowPlaybookRuleRequest) (*model.ShowPlaybookRuleResponse, error) {
	requestDef := GenReqDefForShowPlaybookRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookRuleResponse), nil
	}
}

// ShowPlaybookRuleInvoker 查询剧本规则详情（仅支持华东-上海一使用）
func (c *SaClient) ShowPlaybookRuleInvoker(request *model.ShowPlaybookRuleRequest) *ShowPlaybookRuleInvoker {
	requestDef := GenReqDefForShowPlaybookRule()
	return &ShowPlaybookRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePlaybookRule 更新剧本规则（仅支持华东-上海一使用）
//
// Update rule.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) UpdatePlaybookRule(request *model.UpdatePlaybookRuleRequest) (*model.UpdatePlaybookRuleResponse, error) {
	requestDef := GenReqDefForUpdatePlaybookRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePlaybookRuleResponse), nil
	}
}

// UpdatePlaybookRuleInvoker 更新剧本规则（仅支持华东-上海一使用）
func (c *SaClient) UpdatePlaybookRuleInvoker(request *model.UpdatePlaybookRuleRequest) *UpdatePlaybookRuleInvoker {
	requestDef := GenReqDefForUpdatePlaybookRule()
	return &UpdatePlaybookRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyPlaybookVersion 克隆剧本及版本（仅支持华东-上海一使用）
//
// # Copy Playbook and version to a new one
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) CopyPlaybookVersion(request *model.CopyPlaybookVersionRequest) (*model.CopyPlaybookVersionResponse, error) {
	requestDef := GenReqDefForCopyPlaybookVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyPlaybookVersionResponse), nil
	}
}

// CopyPlaybookVersionInvoker 克隆剧本及版本（仅支持华东-上海一使用）
func (c *SaClient) CopyPlaybookVersionInvoker(request *model.CopyPlaybookVersionRequest) *CopyPlaybookVersionInvoker {
	requestDef := GenReqDefForCopyPlaybookVersion()
	return &CopyPlaybookVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePlaybookVersion 创建剧本版本（仅支持华东-上海一使用）
//
// Create playbook version.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) CreatePlaybookVersion(request *model.CreatePlaybookVersionRequest) (*model.CreatePlaybookVersionResponse, error) {
	requestDef := GenReqDefForCreatePlaybookVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePlaybookVersionResponse), nil
	}
}

// CreatePlaybookVersionInvoker 创建剧本版本（仅支持华东-上海一使用）
func (c *SaClient) CreatePlaybookVersionInvoker(request *model.CreatePlaybookVersionRequest) *CreatePlaybookVersionInvoker {
	requestDef := GenReqDefForCreatePlaybookVersion()
	return &CreatePlaybookVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePlaybookVersion 删除剧本版本（仅支持华东-上海一使用）
//
// Delete playbook version.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) DeletePlaybookVersion(request *model.DeletePlaybookVersionRequest) (*model.DeletePlaybookVersionResponse, error) {
	requestDef := GenReqDefForDeletePlaybookVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePlaybookVersionResponse), nil
	}
}

// DeletePlaybookVersionInvoker 删除剧本版本（仅支持华东-上海一使用）
func (c *SaClient) DeletePlaybookVersionInvoker(request *model.DeletePlaybookVersionRequest) *DeletePlaybookVersionInvoker {
	requestDef := GenReqDefForDeletePlaybookVersion()
	return &DeletePlaybookVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlaybookVersions 查询剧本版本列表（仅支持华东-上海一使用）
//
// List all versions of playbook.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ListPlaybookVersions(request *model.ListPlaybookVersionsRequest) (*model.ListPlaybookVersionsResponse, error) {
	requestDef := GenReqDefForListPlaybookVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlaybookVersionsResponse), nil
	}
}

// ListPlaybookVersionsInvoker 查询剧本版本列表（仅支持华东-上海一使用）
func (c *SaClient) ListPlaybookVersionsInvoker(request *model.ListPlaybookVersionsRequest) *ListPlaybookVersionsInvoker {
	requestDef := GenReqDefForListPlaybookVersions()
	return &ListPlaybookVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybookVersion 查询剧本版本详情（仅支持华东-上海一使用）
//
// # Show playbook version version
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ShowPlaybookVersion(request *model.ShowPlaybookVersionRequest) (*model.ShowPlaybookVersionResponse, error) {
	requestDef := GenReqDefForShowPlaybookVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookVersionResponse), nil
	}
}

// ShowPlaybookVersionInvoker 查询剧本版本详情（仅支持华东-上海一使用）
func (c *SaClient) ShowPlaybookVersionInvoker(request *model.ShowPlaybookVersionRequest) *ShowPlaybookVersionInvoker {
	requestDef := GenReqDefForShowPlaybookVersion()
	return &ShowPlaybookVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePlaybookVersion 更新剧本版本（仅支持华东-上海一使用）
//
// Update playbook version.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) UpdatePlaybookVersion(request *model.UpdatePlaybookVersionRequest) (*model.UpdatePlaybookVersionResponse, error) {
	requestDef := GenReqDefForUpdatePlaybookVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePlaybookVersionResponse), nil
	}
}

// UpdatePlaybookVersionInvoker 更新剧本版本（仅支持华东-上海一使用）
func (c *SaClient) UpdatePlaybookVersionInvoker(request *model.UpdatePlaybookVersionRequest) *UpdatePlaybookVersionInvoker {
	requestDef := GenReqDefForUpdatePlaybookVersion()
	return &UpdatePlaybookVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDataobjectRelation 创建Dataobject关系
//
// # Create Dataobject Relation
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) CreateDataobjectRelation(request *model.CreateDataobjectRelationRequest) (*model.CreateDataobjectRelationResponse, error) {
	requestDef := GenReqDefForCreateDataobjectRelation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDataobjectRelationResponse), nil
	}
}

// CreateDataobjectRelationInvoker 创建Dataobject关系
func (c *SaClient) CreateDataobjectRelationInvoker(request *model.CreateDataobjectRelationRequest) *CreateDataobjectRelationInvoker {
	requestDef := GenReqDefForCreateDataobjectRelation()
	return &CreateDataobjectRelationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDataobjectRelation 删除Dataobject关系
//
// # Delete Dataobject Relation
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) DeleteDataobjectRelation(request *model.DeleteDataobjectRelationRequest) (*model.DeleteDataobjectRelationResponse, error) {
	requestDef := GenReqDefForDeleteDataobjectRelation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDataobjectRelationResponse), nil
	}
}

// DeleteDataobjectRelationInvoker 删除Dataobject关系
func (c *SaClient) DeleteDataobjectRelationInvoker(request *model.DeleteDataobjectRelationRequest) *DeleteDataobjectRelationInvoker {
	requestDef := GenReqDefForDeleteDataobjectRelation()
	return &DeleteDataobjectRelationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataobjectRelation 查询Dataobject关系列表
//
// # List Dataobject Relation
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SaClient) ListDataobjectRelation(request *model.ListDataobjectRelationRequest) (*model.ListDataobjectRelationResponse, error) {
	requestDef := GenReqDefForListDataobjectRelation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataobjectRelationResponse), nil
	}
}

// ListDataobjectRelationInvoker 查询Dataobject关系列表
func (c *SaClient) ListDataobjectRelationInvoker(request *model.ListDataobjectRelationRequest) *ListDataobjectRelationInvoker {
	requestDef := GenReqDefForListDataobjectRelation()
	return &ListDataobjectRelationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
