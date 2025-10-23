package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bcc/v1/model"
)

type BccClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewBccClient(hcClient *httpclient.HcHttpClient) *BccClient {
	return &BccClient{HcClient: hcClient}
}

func BccClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BindResourceLevelCompliance 绑定资源等级合规规则
//
// 资源分级绑定合规规则，会导致属于改等级的资源的合规结果发生变化
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) BindResourceLevelCompliance(request *model.BindResourceLevelComplianceRequest) (*model.BindResourceLevelComplianceResponse, error) {
	requestDef := GenReqDefForBindResourceLevelCompliance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BindResourceLevelComplianceResponse), nil
	}
}

// BindResourceLevelComplianceInvoker 绑定资源等级合规规则
func (c *BccClient) BindResourceLevelComplianceInvoker(request *model.BindResourceLevelComplianceRequest) *BindResourceLevelComplianceInvoker {
	requestDef := GenReqDefForBindResourceLevelCompliance()
	return &BindResourceLevelComplianceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeResourcesLevel 指定资源等级
//
// 手动更改资源的等级，可能会导致资源的合规结果发生变化
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ChangeResourcesLevel(request *model.ChangeResourcesLevelRequest) (*model.ChangeResourcesLevelResponse, error) {
	requestDef := GenReqDefForChangeResourcesLevel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeResourcesLevelResponse), nil
	}
}

// ChangeResourcesLevelInvoker 指定资源等级
func (c *BccClient) ChangeResourcesLevelInvoker(request *model.ChangeResourcesLevelRequest) *ChangeResourcesLevelInvoker {
	requestDef := GenReqDefForChangeResourcesLevel()
	return &ChangeResourcesLevelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateComplianceRule 新建自定义合规规则
//
// 新建自定义合规规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) CreateComplianceRule(request *model.CreateComplianceRuleRequest) (*model.CreateComplianceRuleResponse, error) {
	requestDef := GenReqDefForCreateComplianceRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateComplianceRuleResponse), nil
	}
}

// CreateComplianceRuleInvoker 新建自定义合规规则
func (c *BccClient) CreateComplianceRuleInvoker(request *model.CreateComplianceRuleRequest) *CreateComplianceRuleInvoker {
	requestDef := GenReqDefForCreateComplianceRule()
	return &CreateComplianceRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePolicy 新建策略
//
// 新建策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) CreatePolicy(request *model.CreatePolicyRequest) (*model.CreatePolicyResponse, error) {
	requestDef := GenReqDefForCreatePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePolicyResponse), nil
	}
}

// CreatePolicyInvoker 新建策略
func (c *BccClient) CreatePolicyInvoker(request *model.CreatePolicyRequest) *CreatePolicyInvoker {
	requestDef := GenReqDefForCreatePolicy()
	return &CreatePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateReportSetting 创建报告配置
//
// 创建报告配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) CreateReportSetting(request *model.CreateReportSettingRequest) (*model.CreateReportSettingResponse, error) {
	requestDef := GenReqDefForCreateReportSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateReportSettingResponse), nil
	}
}

// CreateReportSettingInvoker 创建报告配置
func (c *BccClient) CreateReportSettingInvoker(request *model.CreateReportSettingRequest) *CreateReportSettingInvoker {
	requestDef := GenReqDefForCreateReportSetting()
	return &CreateReportSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResourcesLevel 新增资源分级
//
// 新增资源分级
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) CreateResourcesLevel(request *model.CreateResourcesLevelRequest) (*model.CreateResourcesLevelResponse, error) {
	requestDef := GenReqDefForCreateResourcesLevel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResourcesLevelResponse), nil
	}
}

// CreateResourcesLevelInvoker 新增资源分级
func (c *BccClient) CreateResourcesLevelInvoker(request *model.CreateResourcesLevelRequest) *CreateResourcesLevelInvoker {
	requestDef := GenReqDefForCreateResourcesLevel()
	return &CreateResourcesLevelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTemplate 创建模板
//
// 创建模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) CreateTemplate(request *model.CreateTemplateRequest) (*model.CreateTemplateResponse, error) {
	requestDef := GenReqDefForCreateTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTemplateResponse), nil
	}
}

// CreateTemplateInvoker 创建模板
func (c *BccClient) CreateTemplateInvoker(request *model.CreateTemplateRequest) *CreateTemplateInvoker {
	requestDef := GenReqDefForCreateTemplate()
	return &CreateTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteComplianceRule 删除自定义合规规则
//
// 删除自定义合规规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) DeleteComplianceRule(request *model.DeleteComplianceRuleRequest) (*model.DeleteComplianceRuleResponse, error) {
	requestDef := GenReqDefForDeleteComplianceRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteComplianceRuleResponse), nil
	}
}

// DeleteComplianceRuleInvoker 删除自定义合规规则
func (c *BccClient) DeleteComplianceRuleInvoker(request *model.DeleteComplianceRuleRequest) *DeleteComplianceRuleInvoker {
	requestDef := GenReqDefForDeleteComplianceRule()
	return &DeleteComplianceRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePolicy 删除指定策略
//
// 删除指定策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) DeletePolicy(request *model.DeletePolicyRequest) (*model.DeletePolicyResponse, error) {
	requestDef := GenReqDefForDeletePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePolicyResponse), nil
	}
}

// DeletePolicyInvoker 删除指定策略
func (c *BccClient) DeletePolicyInvoker(request *model.DeletePolicyRequest) *DeletePolicyInvoker {
	requestDef := GenReqDefForDeletePolicy()
	return &DeletePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteReport 删除指定的报告
//
// 删除指定的报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) DeleteReport(request *model.DeleteReportRequest) (*model.DeleteReportResponse, error) {
	requestDef := GenReqDefForDeleteReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteReportResponse), nil
	}
}

// DeleteReportInvoker 删除指定的报告
func (c *BccClient) DeleteReportInvoker(request *model.DeleteReportRequest) *DeleteReportInvoker {
	requestDef := GenReqDefForDeleteReport()
	return &DeleteReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteReportSetting 删除报告配置
//
// 删除报告配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) DeleteReportSetting(request *model.DeleteReportSettingRequest) (*model.DeleteReportSettingResponse, error) {
	requestDef := GenReqDefForDeleteReportSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteReportSettingResponse), nil
	}
}

// DeleteReportSettingInvoker 删除报告配置
func (c *BccClient) DeleteReportSettingInvoker(request *model.DeleteReportSettingRequest) *DeleteReportSettingInvoker {
	requestDef := GenReqDefForDeleteReportSetting()
	return &DeleteReportSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTemplate 删除模板
//
// 删除模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) DeleteTemplate(request *model.DeleteTemplateRequest) (*model.DeleteTemplateResponse, error) {
	requestDef := GenReqDefForDeleteTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTemplateResponse), nil
	}
}

// DeleteTemplateInvoker 删除模板
func (c *BccClient) DeleteTemplateInvoker(request *model.DeleteTemplateRequest) *DeleteTemplateInvoker {
	requestDef := GenReqDefForDeleteTemplate()
	return &DeleteTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableDomain 用户授权开启BCC
//
// 用户授权开启BCC
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) EnableDomain(request *model.EnableDomainRequest) (*model.EnableDomainResponse, error) {
	requestDef := GenReqDefForEnableDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableDomainResponse), nil
	}
}

// EnableDomainInvoker 用户授权开启BCC
func (c *BccClient) EnableDomainInvoker(request *model.EnableDomainRequest) *EnableDomainInvoker {
	requestDef := GenReqDefForEnableDomain()
	return &EnableDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmRules 查询告警规则列表
//
// 查询告警规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ListAlarmRules(request *model.ListAlarmRulesRequest) (*model.ListAlarmRulesResponse, error) {
	requestDef := GenReqDefForListAlarmRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmRulesResponse), nil
	}
}

// ListAlarmRulesInvoker 查询告警规则列表
func (c *BccClient) ListAlarmRulesInvoker(request *model.ListAlarmRulesRequest) *ListAlarmRulesInvoker {
	requestDef := GenReqDefForListAlarmRules()
	return &ListAlarmRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarms 查询告警列表
//
// 查询告警列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ListAlarms(request *model.ListAlarmsRequest) (*model.ListAlarmsResponse, error) {
	requestDef := GenReqDefForListAlarms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmsResponse), nil
	}
}

// ListAlarmsInvoker 查询告警列表
func (c *BccClient) ListAlarmsInvoker(request *model.ListAlarmsRequest) *ListAlarmsInvoker {
	requestDef := GenReqDefForListAlarms()
	return &ListAlarmsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListComplianceRule 列举合规规则
//
// 列举合规规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ListComplianceRule(request *model.ListComplianceRuleRequest) (*model.ListComplianceRuleResponse, error) {
	requestDef := GenReqDefForListComplianceRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListComplianceRuleResponse), nil
	}
}

// ListComplianceRuleInvoker 列举合规规则
func (c *BccClient) ListComplianceRuleInvoker(request *model.ListComplianceRuleRequest) *ListComplianceRuleInvoker {
	requestDef := GenReqDefForListComplianceRule()
	return &ListComplianceRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEvents 查询事件数据
//
// 查询事件数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ListEvents(request *model.ListEventsRequest) (*model.ListEventsResponse, error) {
	requestDef := GenReqDefForListEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventsResponse), nil
	}
}

// ListEventsInvoker 查询事件数据
func (c *BccClient) ListEventsInvoker(request *model.ListEventsRequest) *ListEventsInvoker {
	requestDef := GenReqDefForListEvents()
	return &ListEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMetrics 查询监控数据
//
// 查询监控数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ListMetrics(request *model.ListMetricsRequest) (*model.ListMetricsResponse, error) {
	requestDef := GenReqDefForListMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetricsResponse), nil
	}
}

// ListMetricsInvoker 查询监控数据
func (c *BccClient) ListMetricsInvoker(request *model.ListMetricsRequest) *ListMetricsInvoker {
	requestDef := GenReqDefForListMetrics()
	return &ListMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOrganizationPolicy 列举策略
//
// 列举策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ListOrganizationPolicy(request *model.ListOrganizationPolicyRequest) (*model.ListOrganizationPolicyResponse, error) {
	requestDef := GenReqDefForListOrganizationPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrganizationPolicyResponse), nil
	}
}

// ListOrganizationPolicyInvoker 列举策略
func (c *BccClient) ListOrganizationPolicyInvoker(request *model.ListOrganizationPolicyRequest) *ListOrganizationPolicyInvoker {
	requestDef := GenReqDefForListOrganizationPolicy()
	return &ListOrganizationPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicy 列举策略
//
// 列举策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ListPolicy(request *model.ListPolicyRequest) (*model.ListPolicyResponse, error) {
	requestDef := GenReqDefForListPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyResponse), nil
	}
}

// ListPolicyInvoker 列举策略
func (c *BccClient) ListPolicyInvoker(request *model.ListPolicyRequest) *ListPolicyInvoker {
	requestDef := GenReqDefForListPolicy()
	return &ListPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListReportSettings 查询报告配置列表
//
// 查询报告配置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ListReportSettings(request *model.ListReportSettingsRequest) (*model.ListReportSettingsResponse, error) {
	requestDef := GenReqDefForListReportSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListReportSettingsResponse), nil
	}
}

// ListReportSettingsInvoker 查询报告配置列表
func (c *BccClient) ListReportSettingsInvoker(request *model.ListReportSettingsRequest) *ListReportSettingsInvoker {
	requestDef := GenReqDefForListReportSettings()
	return &ListReportSettingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListReports 查询报告列表
//
// 查询报告列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ListReports(request *model.ListReportsRequest) (*model.ListReportsResponse, error) {
	requestDef := GenReqDefForListReports()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListReportsResponse), nil
	}
}

// ListReportsInvoker 查询报告列表
func (c *BccClient) ListReportsInvoker(request *model.ListReportsRequest) *ListReportsInvoker {
	requestDef := GenReqDefForListReports()
	return &ListReportsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceCopies 查询副本列表
//
// 查询副本列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ListResourceCopies(request *model.ListResourceCopiesRequest) (*model.ListResourceCopiesResponse, error) {
	requestDef := GenReqDefForListResourceCopies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceCopiesResponse), nil
	}
}

// ListResourceCopiesInvoker 查询副本列表
func (c *BccClient) ListResourceCopiesInvoker(request *model.ListResourceCopiesRequest) *ListResourceCopiesInvoker {
	requestDef := GenReqDefForListResourceCopies()
	return &ListResourceCopiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResources 查询资源列表
//
// 查询资源列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ListResources(request *model.ListResourcesRequest) (*model.ListResourcesResponse, error) {
	requestDef := GenReqDefForListResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourcesResponse), nil
	}
}

// ListResourcesInvoker 查询资源列表
func (c *BccClient) ListResourcesInvoker(request *model.ListResourcesRequest) *ListResourcesInvoker {
	requestDef := GenReqDefForListResources()
	return &ListResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourcesLevel 列举资源分级
//
// 列举资源分级
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ListResourcesLevel(request *model.ListResourcesLevelRequest) (*model.ListResourcesLevelResponse, error) {
	requestDef := GenReqDefForListResourcesLevel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourcesLevelResponse), nil
	}
}

// ListResourcesLevelInvoker 列举资源分级
func (c *BccClient) ListResourcesLevelInvoker(request *model.ListResourcesLevelRequest) *ListResourcesLevelInvoker {
	requestDef := GenReqDefForListResourcesLevel()
	return &ListResourcesLevelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourcesLevelTags 列举资源分级已指定的标签
//
// 列举资源分级已指定的标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ListResourcesLevelTags(request *model.ListResourcesLevelTagsRequest) (*model.ListResourcesLevelTagsResponse, error) {
	requestDef := GenReqDefForListResourcesLevelTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourcesLevelTagsResponse), nil
	}
}

// ListResourcesLevelTagsInvoker 列举资源分级已指定的标签
func (c *BccClient) ListResourcesLevelTagsInvoker(request *model.ListResourcesLevelTagsRequest) *ListResourcesLevelTagsInvoker {
	requestDef := GenReqDefForListResourcesLevelTags()
	return &ListResourcesLevelTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSupportedRegion 查询支持的region列表
//
// BCC支持的region列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ListSupportedRegion(request *model.ListSupportedRegionRequest) (*model.ListSupportedRegionResponse, error) {
	requestDef := GenReqDefForListSupportedRegion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSupportedRegionResponse), nil
	}
}

// ListSupportedRegionInvoker 查询支持的region列表
func (c *BccClient) ListSupportedRegionInvoker(request *model.ListSupportedRegionRequest) *ListSupportedRegionInvoker {
	requestDef := GenReqDefForListSupportedRegion()
	return &ListSupportedRegionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTasks 查询任务列表
//
// 查询任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ListTasks(request *model.ListTasksRequest) (*model.ListTasksResponse, error) {
	requestDef := GenReqDefForListTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTasksResponse), nil
	}
}

// ListTasksInvoker 查询任务列表
func (c *BccClient) ListTasksInvoker(request *model.ListTasksRequest) *ListTasksInvoker {
	requestDef := GenReqDefForListTasks()
	return &ListTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTemplates 查询模板列表
//
// 查询模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ListTemplates(request *model.ListTemplatesRequest) (*model.ListTemplatesResponse, error) {
	requestDef := GenReqDefForListTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplatesResponse), nil
	}
}

// ListTemplatesInvoker 查询模板列表
func (c *BccClient) ListTemplatesInvoker(request *model.ListTemplatesRequest) *ListTemplatesInvoker {
	requestDef := GenReqDefForListTemplates()
	return &ListTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVault 列举存储库
//
// 列举存储库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ListVault(request *model.ListVaultRequest) (*model.ListVaultResponse, error) {
	requestDef := GenReqDefForListVault()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVaultResponse), nil
	}
}

// ListVaultInvoker 列举存储库
func (c *BccClient) ListVaultInvoker(request *model.ListVaultRequest) *ListVaultInvoker {
	requestDef := GenReqDefForListVault()
	return &ListVaultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVaultTops 查询使用最高的存储库
//
// 查询使用最高的存储库，按使用容量或者按使用率排序返回
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ListVaultTops(request *model.ListVaultTopsRequest) (*model.ListVaultTopsResponse, error) {
	requestDef := GenReqDefForListVaultTops()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVaultTopsResponse), nil
	}
}

// ListVaultTopsInvoker 查询使用最高的存储库
func (c *BccClient) ListVaultTopsInvoker(request *model.ListVaultTopsRequest) *ListVaultTopsInvoker {
	requestDef := GenReqDefForListVaultTops()
	return &ListVaultTopsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyResourceLevel 修改资源分级
//
// 修改资源分级，修改资源等级的应用类型、应用规则、合规规则id时，可能会导致某些资源的归属等级发生变动，或资源的合规结果发生更改
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ModifyResourceLevel(request *model.ModifyResourceLevelRequest) (*model.ModifyResourceLevelResponse, error) {
	requestDef := GenReqDefForModifyResourceLevel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyResourceLevelResponse), nil
	}
}

// ModifyResourceLevelInvoker 修改资源分级
func (c *BccClient) ModifyResourceLevelInvoker(request *model.ModifyResourceLevelRequest) *ModifyResourceLevelInvoker {
	requestDef := GenReqDefForModifyResourceLevel()
	return &ModifyResourceLevelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NotifyResourceChange 资源变化通知
//
// 资源变化通知，当资源的属性变动时，可能会改变资源归属的等级，进而可能导致合规结果的变化
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) NotifyResourceChange(request *model.NotifyResourceChangeRequest) (*model.NotifyResourceChangeResponse, error) {
	requestDef := GenReqDefForNotifyResourceChange()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NotifyResourceChangeResponse), nil
	}
}

// NotifyResourceChangeInvoker 资源变化通知
func (c *BccClient) NotifyResourceChangeInvoker(request *model.NotifyResourceChangeRequest) *NotifyResourceChangeInvoker {
	requestDef := GenReqDefForNotifyResourceChange()
	return &NotifyResourceChangeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveResourceLevel 删除资源分级
//
// 删除资源分级，删除资源等级会导致归属于该等级的资源的合规结果清空
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) RemoveResourceLevel(request *model.RemoveResourceLevelRequest) (*model.RemoveResourceLevelResponse, error) {
	requestDef := GenReqDefForRemoveResourceLevel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveResourceLevelResponse), nil
	}
}

// RemoveResourceLevelInvoker 删除资源分级
func (c *BccClient) RemoveResourceLevelInvoker(request *model.RemoveResourceLevelRequest) *RemoveResourceLevelInvoker {
	requestDef := GenReqDefForRemoveResourceLevel()
	return &RemoveResourceLevelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetProtectionConfiguration 配置资源保护
//
// 配置资源保护，资源保护情况发生变动后，相关资源的合规结果可能会发生变化
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) SetProtectionConfiguration(request *model.SetProtectionConfigurationRequest) (*model.SetProtectionConfigurationResponse, error) {
	requestDef := GenReqDefForSetProtectionConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetProtectionConfigurationResponse), nil
	}
}

// SetProtectionConfigurationInvoker 配置资源保护
func (c *BccClient) SetProtectionConfigurationInvoker(request *model.SetProtectionConfigurationRequest) *SetProtectionConfigurationInvoker {
	requestDef := GenReqDefForSetProtectionConfiguration()
	return &SetProtectionConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlarmSummary 查询告警统计信息
//
// 查询告警统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ShowAlarmSummary(request *model.ShowAlarmSummaryRequest) (*model.ShowAlarmSummaryResponse, error) {
	requestDef := GenReqDefForShowAlarmSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlarmSummaryResponse), nil
	}
}

// ShowAlarmSummaryInvoker 查询告警统计信息
func (c *BccClient) ShowAlarmSummaryInvoker(request *model.ShowAlarmSummaryRequest) *ShowAlarmSummaryInvoker {
	requestDef := GenReqDefForShowAlarmSummary()
	return &ShowAlarmSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowComplianceRule 查看自定义合规规则详情
//
// 查看自定义合规规则详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ShowComplianceRule(request *model.ShowComplianceRuleRequest) (*model.ShowComplianceRuleResponse, error) {
	requestDef := GenReqDefForShowComplianceRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowComplianceRuleResponse), nil
	}
}

// ShowComplianceRuleInvoker 查看自定义合规规则详情
func (c *BccClient) ShowComplianceRuleInvoker(request *model.ShowComplianceRuleRequest) *ShowComplianceRuleInvoker {
	requestDef := GenReqDefForShowComplianceRule()
	return &ShowComplianceRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDomain 查询BCC开启情况
//
// 查询BCC开启情况
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ShowDomain(request *model.ShowDomainRequest) (*model.ShowDomainResponse, error) {
	requestDef := GenReqDefForShowDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainResponse), nil
	}
}

// ShowDomainInvoker 查询BCC开启情况
func (c *BccClient) ShowDomainInvoker(request *model.ShowDomainRequest) *ShowDomainInvoker {
	requestDef := GenReqDefForShowDomain()
	return &ShowDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOrganizationPolicy 查询指定组织策略
//
// 查询指定组织策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ShowOrganizationPolicy(request *model.ShowOrganizationPolicyRequest) (*model.ShowOrganizationPolicyResponse, error) {
	requestDef := GenReqDefForShowOrganizationPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOrganizationPolicyResponse), nil
	}
}

// ShowOrganizationPolicyInvoker 查询指定组织策略
func (c *BccClient) ShowOrganizationPolicyInvoker(request *model.ShowOrganizationPolicyRequest) *ShowOrganizationPolicyInvoker {
	requestDef := GenReqDefForShowOrganizationPolicy()
	return &ShowOrganizationPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPolicy 查询指定策略
//
// 查询指定策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ShowPolicy(request *model.ShowPolicyRequest) (*model.ShowPolicyResponse, error) {
	requestDef := GenReqDefForShowPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPolicyResponse), nil
	}
}

// ShowPolicyInvoker 查询指定策略
func (c *BccClient) ShowPolicyInvoker(request *model.ShowPolicyRequest) *ShowPolicyInvoker {
	requestDef := GenReqDefForShowPolicy()
	return &ShowPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReport 查询报告详情
//
// 查询报告详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ShowReport(request *model.ShowReportRequest) (*model.ShowReportResponse, error) {
	requestDef := GenReqDefForShowReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReportResponse), nil
	}
}

// ShowReportInvoker 查询报告详情
func (c *BccClient) ShowReportInvoker(request *model.ShowReportRequest) *ShowReportInvoker {
	requestDef := GenReqDefForShowReport()
	return &ShowReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReportResourceDetailsData 查询报告详情类数据
//
// 查询指定租户下的指定报告的详情类数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ShowReportResourceDetailsData(request *model.ShowReportResourceDetailsDataRequest) (*model.ShowReportResourceDetailsDataResponse, error) {
	requestDef := GenReqDefForShowReportResourceDetailsData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReportResourceDetailsDataResponse), nil
	}
}

// ShowReportResourceDetailsDataInvoker 查询报告详情类数据
func (c *BccClient) ShowReportResourceDetailsDataInvoker(request *model.ShowReportResourceDetailsDataRequest) *ShowReportResourceDetailsDataInvoker {
	requestDef := GenReqDefForShowReportResourceDetailsData()
	return &ShowReportResourceDetailsDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReportSetting 查询报告配置
//
// 查询报告配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ShowReportSetting(request *model.ShowReportSettingRequest) (*model.ShowReportSettingResponse, error) {
	requestDef := GenReqDefForShowReportSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReportSettingResponse), nil
	}
}

// ShowReportSettingInvoker 查询报告配置
func (c *BccClient) ShowReportSettingInvoker(request *model.ShowReportSettingRequest) *ShowReportSettingInvoker {
	requestDef := GenReqDefForShowReportSetting()
	return &ShowReportSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReportSummaryData 查询报告摘要类数据
//
// 查询指定租户下的指定报告的摘要类数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ShowReportSummaryData(request *model.ShowReportSummaryDataRequest) (*model.ShowReportSummaryDataResponse, error) {
	requestDef := GenReqDefForShowReportSummaryData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReportSummaryDataResponse), nil
	}
}

// ShowReportSummaryDataInvoker 查询报告摘要类数据
func (c *BccClient) ShowReportSummaryDataInvoker(request *model.ShowReportSummaryDataRequest) *ShowReportSummaryDataInvoker {
	requestDef := GenReqDefForShowReportSummaryData()
	return &ShowReportSummaryDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReportTaskDetailsData 查询报告详情类数据
//
// 查询指定租户下的指定报告的详情类数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ShowReportTaskDetailsData(request *model.ShowReportTaskDetailsDataRequest) (*model.ShowReportTaskDetailsDataResponse, error) {
	requestDef := GenReqDefForShowReportTaskDetailsData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReportTaskDetailsDataResponse), nil
	}
}

// ShowReportTaskDetailsDataInvoker 查询报告详情类数据
func (c *BccClient) ShowReportTaskDetailsDataInvoker(request *model.ShowReportTaskDetailsDataRequest) *ShowReportTaskDetailsDataInvoker {
	requestDef := GenReqDefForShowReportTaskDetailsData()
	return &ShowReportTaskDetailsDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceCopiesSummary 查询副本摘要
//
// get copies summary
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ShowResourceCopiesSummary(request *model.ShowResourceCopiesSummaryRequest) (*model.ShowResourceCopiesSummaryResponse, error) {
	requestDef := GenReqDefForShowResourceCopiesSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceCopiesSummaryResponse), nil
	}
}

// ShowResourceCopiesSummaryInvoker 查询副本摘要
func (c *BccClient) ShowResourceCopiesSummaryInvoker(request *model.ShowResourceCopiesSummaryRequest) *ShowResourceCopiesSummaryInvoker {
	requestDef := GenReqDefForShowResourceCopiesSummary()
	return &ShowResourceCopiesSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceDetail 查看资源详情
//
// 查看资源详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ShowResourceDetail(request *model.ShowResourceDetailRequest) (*model.ShowResourceDetailResponse, error) {
	requestDef := GenReqDefForShowResourceDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceDetailResponse), nil
	}
}

// ShowResourceDetailInvoker 查看资源详情
func (c *BccClient) ShowResourceDetailInvoker(request *model.ShowResourceDetailRequest) *ShowResourceDetailInvoker {
	requestDef := GenReqDefForShowResourceDetail()
	return &ShowResourceDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourcesSummary 查询资源统计
//
// 查询资源统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ShowResourcesSummary(request *model.ShowResourcesSummaryRequest) (*model.ShowResourcesSummaryResponse, error) {
	requestDef := GenReqDefForShowResourcesSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourcesSummaryResponse), nil
	}
}

// ShowResourcesSummaryInvoker 查询资源统计
func (c *BccClient) ShowResourcesSummaryInvoker(request *model.ShowResourcesSummaryRequest) *ShowResourcesSummaryInvoker {
	requestDef := GenReqDefForShowResourcesSummary()
	return &ShowResourcesSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTask 查询任务详情
//
// 查询任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ShowTask(request *model.ShowTaskRequest) (*model.ShowTaskResponse, error) {
	requestDef := GenReqDefForShowTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskResponse), nil
	}
}

// ShowTaskInvoker 查询任务详情
func (c *BccClient) ShowTaskInvoker(request *model.ShowTaskRequest) *ShowTaskInvoker {
	requestDef := GenReqDefForShowTask()
	return &ShowTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskStatusSummary 查询任务统计信息
//
// 查询任务统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ShowTaskStatusSummary(request *model.ShowTaskStatusSummaryRequest) (*model.ShowTaskStatusSummaryResponse, error) {
	requestDef := GenReqDefForShowTaskStatusSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskStatusSummaryResponse), nil
	}
}

// ShowTaskStatusSummaryInvoker 查询任务统计信息
func (c *BccClient) ShowTaskStatusSummaryInvoker(request *model.ShowTaskStatusSummaryRequest) *ShowTaskStatusSummaryInvoker {
	requestDef := GenReqDefForShowTaskStatusSummary()
	return &ShowTaskStatusSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskTypeSummary 查询任务统计信息
//
// 查询任务统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ShowTaskTypeSummary(request *model.ShowTaskTypeSummaryRequest) (*model.ShowTaskTypeSummaryResponse, error) {
	requestDef := GenReqDefForShowTaskTypeSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskTypeSummaryResponse), nil
	}
}

// ShowTaskTypeSummaryInvoker 查询任务统计信息
func (c *BccClient) ShowTaskTypeSummaryInvoker(request *model.ShowTaskTypeSummaryRequest) *ShowTaskTypeSummaryInvoker {
	requestDef := GenReqDefForShowTaskTypeSummary()
	return &ShowTaskTypeSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTemplate 查询模板
//
// 查询模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ShowTemplate(request *model.ShowTemplateRequest) (*model.ShowTemplateResponse, error) {
	requestDef := GenReqDefForShowTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateResponse), nil
	}
}

// ShowTemplateInvoker 查询模板
func (c *BccClient) ShowTemplateInvoker(request *model.ShowTemplateRequest) *ShowTemplateInvoker {
	requestDef := GenReqDefForShowTemplate()
	return &ShowTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVault 查看指定存储库
//
// 查看指定存储库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ShowVault(request *model.ShowVaultRequest) (*model.ShowVaultResponse, error) {
	requestDef := GenReqDefForShowVault()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVaultResponse), nil
	}
}

// ShowVaultInvoker 查看指定存储库
func (c *BccClient) ShowVaultInvoker(request *model.ShowVaultRequest) *ShowVaultInvoker {
	requestDef := GenReqDefForShowVault()
	return &ShowVaultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVaultSummary 查看租户的Vault的统计信息
//
// 查看租户Vault的统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) ShowVaultSummary(request *model.ShowVaultSummaryRequest) (*model.ShowVaultSummaryResponse, error) {
	requestDef := GenReqDefForShowVaultSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVaultSummaryResponse), nil
	}
}

// ShowVaultSummaryInvoker 查看租户的Vault的统计信息
func (c *BccClient) ShowVaultSummaryInvoker(request *model.ShowVaultSummaryRequest) *ShowVaultSummaryInvoker {
	requestDef := GenReqDefForShowVaultSummary()
	return &ShowVaultSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnbindResourceLevelCompliance 解绑资源等级合规规则
//
// 资源分级绑定合规规则，修改资源等级的合规规则，会导致属于改等级的资源的合规结果发生变化
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) UnbindResourceLevelCompliance(request *model.UnbindResourceLevelComplianceRequest) (*model.UnbindResourceLevelComplianceResponse, error) {
	requestDef := GenReqDefForUnbindResourceLevelCompliance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnbindResourceLevelComplianceResponse), nil
	}
}

// UnbindResourceLevelComplianceInvoker 解绑资源等级合规规则
func (c *BccClient) UnbindResourceLevelComplianceInvoker(request *model.UnbindResourceLevelComplianceRequest) *UnbindResourceLevelComplianceInvoker {
	requestDef := GenReqDefForUnbindResourceLevelCompliance()
	return &UnbindResourceLevelComplianceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateComplianceRule 更新自定义合规规则
//
// 更新自定义合规规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) UpdateComplianceRule(request *model.UpdateComplianceRuleRequest) (*model.UpdateComplianceRuleResponse, error) {
	requestDef := GenReqDefForUpdateComplianceRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateComplianceRuleResponse), nil
	}
}

// UpdateComplianceRuleInvoker 更新自定义合规规则
func (c *BccClient) UpdateComplianceRuleInvoker(request *model.UpdateComplianceRuleRequest) *UpdateComplianceRuleInvoker {
	requestDef := GenReqDefForUpdateComplianceRule()
	return &UpdateComplianceRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePolicy 更新指定策略
//
// 更新指定策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) UpdatePolicy(request *model.UpdatePolicyRequest) (*model.UpdatePolicyResponse, error) {
	requestDef := GenReqDefForUpdatePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePolicyResponse), nil
	}
}

// UpdatePolicyInvoker 更新指定策略
func (c *BccClient) UpdatePolicyInvoker(request *model.UpdatePolicyRequest) *UpdatePolicyInvoker {
	requestDef := GenReqDefForUpdatePolicy()
	return &UpdatePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateReportSetting 修改报告配置
//
// 更新报告配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) UpdateReportSetting(request *model.UpdateReportSettingRequest) (*model.UpdateReportSettingResponse, error) {
	requestDef := GenReqDefForUpdateReportSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateReportSettingResponse), nil
	}
}

// UpdateReportSettingInvoker 修改报告配置
func (c *BccClient) UpdateReportSettingInvoker(request *model.UpdateReportSettingRequest) *UpdateReportSettingInvoker {
	requestDef := GenReqDefForUpdateReportSetting()
	return &UpdateReportSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTemplate 修改模板
//
// 修改模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BccClient) UpdateTemplate(request *model.UpdateTemplateRequest) (*model.UpdateTemplateResponse, error) {
	requestDef := GenReqDefForUpdateTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTemplateResponse), nil
	}
}

// UpdateTemplateInvoker 修改模板
func (c *BccClient) UpdateTemplateInvoker(request *model.UpdateTemplateRequest) *UpdateTemplateInvoker {
	requestDef := GenReqDefForUpdateTemplate()
	return &UpdateTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
