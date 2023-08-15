package v1

import (
	http_client "github.com/dysodeng/huaweicloud-sdk-go-v3/core"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/dbss/v1/model"
)

type DbssClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDbssClient(hcClient *http_client.HcHttpClient) *DbssClient {
	return &DbssClient{HcClient: hcClient}
}

func DbssClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AddRdsNoAgentDatabase 添加RDS免agent数据库
//
// 添加RDS免agent数据库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) AddRdsNoAgentDatabase(request *model.AddRdsNoAgentDatabaseRequest) (*model.AddRdsNoAgentDatabaseResponse, error) {
	requestDef := GenReqDefForAddRdsNoAgentDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddRdsNoAgentDatabaseResponse), nil
	}
}

// AddRdsNoAgentDatabaseInvoker 添加RDS免agent数据库
func (c *DbssClient) AddRdsNoAgentDatabaseInvoker(request *model.AddRdsNoAgentDatabaseRequest) *AddRdsNoAgentDatabaseInvoker {
	requestDef := GenReqDefForAddRdsNoAgentDatabase()
	return &AddRdsNoAgentDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAddResourceTag 批量添加资源标签
//
// 批量添加资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) BatchAddResourceTag(request *model.BatchAddResourceTagRequest) (*model.BatchAddResourceTagResponse, error) {
	requestDef := GenReqDefForBatchAddResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddResourceTagResponse), nil
	}
}

// BatchAddResourceTagInvoker 批量添加资源标签
func (c *DbssClient) BatchAddResourceTagInvoker(request *model.BatchAddResourceTagRequest) *BatchAddResourceTagInvoker {
	requestDef := GenReqDefForBatchAddResourceTag()
	return &BatchAddResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteResourceTag 批量删除资源标签
//
// 批量删除资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) BatchDeleteResourceTag(request *model.BatchDeleteResourceTagRequest) (*model.BatchDeleteResourceTagResponse, error) {
	requestDef := GenReqDefForBatchDeleteResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteResourceTagResponse), nil
	}
}

// BatchDeleteResourceTagInvoker 批量删除资源标签
func (c *DbssClient) BatchDeleteResourceTagInvoker(request *model.BatchDeleteResourceTagRequest) *BatchDeleteResourceTagInvoker {
	requestDef := GenReqDefForBatchDeleteResourceTag()
	return &BatchDeleteResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountResourceInstanceByTag 根据标签查询资源实例数量
//
// 根据标签查询资源实例数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) CountResourceInstanceByTag(request *model.CountResourceInstanceByTagRequest) (*model.CountResourceInstanceByTagResponse, error) {
	requestDef := GenReqDefForCountResourceInstanceByTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountResourceInstanceByTagResponse), nil
	}
}

// CountResourceInstanceByTagInvoker 根据标签查询资源实例数量
func (c *DbssClient) CountResourceInstanceByTagInvoker(request *model.CountResourceInstanceByTagRequest) *CountResourceInstanceByTagInvoker {
	requestDef := GenReqDefForCountResourceInstanceByTag()
	return &CountResourceInstanceByTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstancesPeriodOrder 包年包月计费模式创建审计实例
//
// 包年包月计费模式创建审计实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) CreateInstancesPeriodOrder(request *model.CreateInstancesPeriodOrderRequest) (*model.CreateInstancesPeriodOrderResponse, error) {
	requestDef := GenReqDefForCreateInstancesPeriodOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstancesPeriodOrderResponse), nil
	}
}

// CreateInstancesPeriodOrderInvoker 包年包月计费模式创建审计实例
func (c *DbssClient) CreateInstancesPeriodOrderInvoker(request *model.CreateInstancesPeriodOrderRequest) *CreateInstancesPeriodOrderInvoker {
	requestDef := GenReqDefForCreateInstancesPeriodOrder()
	return &CreateInstancesPeriodOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditDatabases 查询数据库列表
//
// 查询数据库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditDatabases(request *model.ListAuditDatabasesRequest) (*model.ListAuditDatabasesResponse, error) {
	requestDef := GenReqDefForListAuditDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditDatabasesResponse), nil
	}
}

// ListAuditDatabasesInvoker 查询数据库列表
func (c *DbssClient) ListAuditDatabasesInvoker(request *model.ListAuditDatabasesRequest) *ListAuditDatabasesInvoker {
	requestDef := GenReqDefForListAuditDatabases()
	return &ListAuditDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditInstanceJobs 查询实例创建任务信息
//
// 查询实例创建任务信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditInstanceJobs(request *model.ListAuditInstanceJobsRequest) (*model.ListAuditInstanceJobsResponse, error) {
	requestDef := GenReqDefForListAuditInstanceJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditInstanceJobsResponse), nil
	}
}

// ListAuditInstanceJobsInvoker 查询实例创建任务信息
func (c *DbssClient) ListAuditInstanceJobsInvoker(request *model.ListAuditInstanceJobsRequest) *ListAuditInstanceJobsInvoker {
	requestDef := GenReqDefForListAuditInstanceJobs()
	return &ListAuditInstanceJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditInstances 查询审计实例列表
//
// 查询审计实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditInstances(request *model.ListAuditInstancesRequest) (*model.ListAuditInstancesResponse, error) {
	requestDef := GenReqDefForListAuditInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditInstancesResponse), nil
	}
}

// ListAuditInstancesInvoker 查询审计实例列表
func (c *DbssClient) ListAuditInstancesInvoker(request *model.ListAuditInstancesRequest) *ListAuditInstancesInvoker {
	requestDef := GenReqDefForListAuditInstances()
	return &ListAuditInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditOperateLogs 查询用户操作日志信息
//
// 查询用户操作日志信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditOperateLogs(request *model.ListAuditOperateLogsRequest) (*model.ListAuditOperateLogsResponse, error) {
	requestDef := GenReqDefForListAuditOperateLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditOperateLogsResponse), nil
	}
}

// ListAuditOperateLogsInvoker 查询用户操作日志信息
func (c *DbssClient) ListAuditOperateLogsInvoker(request *model.ListAuditOperateLogsRequest) *ListAuditOperateLogsInvoker {
	requestDef := GenReqDefForListAuditOperateLogs()
	return &ListAuditOperateLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditRuleRisks 查询风险规则策略
//
// 查询风险规则策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditRuleRisks(request *model.ListAuditRuleRisksRequest) (*model.ListAuditRuleRisksResponse, error) {
	requestDef := GenReqDefForListAuditRuleRisks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditRuleRisksResponse), nil
	}
}

// ListAuditRuleRisksInvoker 查询风险规则策略
func (c *DbssClient) ListAuditRuleRisksInvoker(request *model.ListAuditRuleRisksRequest) *ListAuditRuleRisksInvoker {
	requestDef := GenReqDefForListAuditRuleRisks()
	return &ListAuditRuleRisksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditRuleScopes 查询审计范围策略列表
//
// 查询审计范围策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditRuleScopes(request *model.ListAuditRuleScopesRequest) (*model.ListAuditRuleScopesResponse, error) {
	requestDef := GenReqDefForListAuditRuleScopes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditRuleScopesResponse), nil
	}
}

// ListAuditRuleScopesInvoker 查询审计范围策略列表
func (c *DbssClient) ListAuditRuleScopesInvoker(request *model.ListAuditRuleScopesRequest) *ListAuditRuleScopesInvoker {
	requestDef := GenReqDefForListAuditRuleScopes()
	return &ListAuditRuleScopesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditSensitiveMasks 查询隐私数据脱敏规则
//
// 查询隐私数据脱敏规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditSensitiveMasks(request *model.ListAuditSensitiveMasksRequest) (*model.ListAuditSensitiveMasksResponse, error) {
	requestDef := GenReqDefForListAuditSensitiveMasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditSensitiveMasksResponse), nil
	}
}

// ListAuditSensitiveMasksInvoker 查询隐私数据脱敏规则
func (c *DbssClient) ListAuditSensitiveMasksInvoker(request *model.ListAuditSensitiveMasksRequest) *ListAuditSensitiveMasksInvoker {
	requestDef := GenReqDefForListAuditSensitiveMasks()
	return &ListAuditSensitiveMasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailabilityZoneInfos 查询可用区信息
//
// 查询可用区信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAvailabilityZoneInfos(request *model.ListAvailabilityZoneInfosRequest) (*model.ListAvailabilityZoneInfosResponse, error) {
	requestDef := GenReqDefForListAvailabilityZoneInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailabilityZoneInfosResponse), nil
	}
}

// ListAvailabilityZoneInfosInvoker 查询可用区信息
func (c *DbssClient) ListAvailabilityZoneInfosInvoker(request *model.ListAvailabilityZoneInfosRequest) *ListAvailabilityZoneInfosInvoker {
	requestDef := GenReqDefForListAvailabilityZoneInfos()
	return &ListAvailabilityZoneInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEcsSpecification 查询ecs服务器规格信息
//
// 查询ecs服务器规格信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListEcsSpecification(request *model.ListEcsSpecificationRequest) (*model.ListEcsSpecificationResponse, error) {
	requestDef := GenReqDefForListEcsSpecification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEcsSpecificationResponse), nil
	}
}

// ListEcsSpecificationInvoker 查询ecs服务器规格信息
func (c *DbssClient) ListEcsSpecificationInvoker(request *model.ListEcsSpecificationRequest) *ListEcsSpecificationInvoker {
	requestDef := GenReqDefForListEcsSpecification()
	return &ListEcsSpecificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectResourceTags 查询项目标签
//
// 查询项目标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListProjectResourceTags(request *model.ListProjectResourceTagsRequest) (*model.ListProjectResourceTagsResponse, error) {
	requestDef := GenReqDefForListProjectResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectResourceTagsResponse), nil
	}
}

// ListProjectResourceTagsInvoker 查询项目标签
func (c *DbssClient) ListProjectResourceTagsInvoker(request *model.ListProjectResourceTagsRequest) *ListProjectResourceTagsInvoker {
	requestDef := GenReqDefForListProjectResourceTags()
	return &ListProjectResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceInstanceByTag 根据标签查询资源实例列表
//
// 根据标签查询资源实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListResourceInstanceByTag(request *model.ListResourceInstanceByTagRequest) (*model.ListResourceInstanceByTagResponse, error) {
	requestDef := GenReqDefForListResourceInstanceByTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceInstanceByTagResponse), nil
	}
}

// ListResourceInstanceByTagInvoker 根据标签查询资源实例列表
func (c *DbssClient) ListResourceInstanceByTagInvoker(request *model.ListResourceInstanceByTagRequest) *ListResourceInstanceByTagInvoker {
	requestDef := GenReqDefForListResourceInstanceByTag()
	return &ListResourceInstanceByTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSqlInjectionRules 查询SQL注入规则策略
//
// 查询SQL注入规则策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListSqlInjectionRules(request *model.ListSqlInjectionRulesRequest) (*model.ListSqlInjectionRulesResponse, error) {
	requestDef := GenReqDefForListSqlInjectionRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSqlInjectionRulesResponse), nil
	}
}

// ListSqlInjectionRulesInvoker 查询SQL注入规则策略
func (c *DbssClient) ListSqlInjectionRulesInvoker(request *model.ListSqlInjectionRulesRequest) *ListSqlInjectionRulesInvoker {
	requestDef := GenReqDefForListSqlInjectionRules()
	return &ListSqlInjectionRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuditQuota 查询账户配额信息
//
// 查询账户配额信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ShowAuditQuota(request *model.ShowAuditQuotaRequest) (*model.ShowAuditQuotaResponse, error) {
	requestDef := GenReqDefForShowAuditQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuditQuotaResponse), nil
	}
}

// ShowAuditQuotaInvoker 查询账户配额信息
func (c *DbssClient) ShowAuditQuotaInvoker(request *model.ShowAuditQuotaRequest) *ShowAuditQuotaInvoker {
	requestDef := GenReqDefForShowAuditQuota()
	return &ShowAuditQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuditRuleRisk 查询指定风险规则策略
//
// 查询指定风险规则策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ShowAuditRuleRisk(request *model.ShowAuditRuleRiskRequest) (*model.ShowAuditRuleRiskResponse, error) {
	requestDef := GenReqDefForShowAuditRuleRisk()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuditRuleRiskResponse), nil
	}
}

// ShowAuditRuleRiskInvoker 查询指定风险规则策略
func (c *DbssClient) ShowAuditRuleRiskInvoker(request *model.ShowAuditRuleRiskRequest) *ShowAuditRuleRiskInvoker {
	requestDef := GenReqDefForShowAuditRuleRisk()
	return &ShowAuditRuleRiskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchAgent 开启关闭Agent
//
// 用于开启和关闭agent的功能，当开启后，开始抓取用户的访问信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SwitchAgent(request *model.SwitchAgentRequest) (*model.SwitchAgentResponse, error) {
	requestDef := GenReqDefForSwitchAgent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchAgentResponse), nil
	}
}

// SwitchAgentInvoker 开启关闭Agent
func (c *DbssClient) SwitchAgentInvoker(request *model.SwitchAgentRequest) *SwitchAgentInvoker {
	requestDef := GenReqDefForSwitchAgent()
	return &SwitchAgentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchRiskRule 开启关闭风险规则
//
// 开启关闭风险规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SwitchRiskRule(request *model.SwitchRiskRuleRequest) (*model.SwitchRiskRuleResponse, error) {
	requestDef := GenReqDefForSwitchRiskRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchRiskRuleResponse), nil
	}
}

// SwitchRiskRuleInvoker 开启关闭风险规则
func (c *DbssClient) SwitchRiskRuleInvoker(request *model.SwitchRiskRuleRequest) *SwitchRiskRuleInvoker {
	requestDef := GenReqDefForSwitchRiskRule()
	return &SwitchRiskRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAuditSecurityGroup 修改安全组
//
// 修改安全组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) UpdateAuditSecurityGroup(request *model.UpdateAuditSecurityGroupRequest) (*model.UpdateAuditSecurityGroupResponse, error) {
	requestDef := GenReqDefForUpdateAuditSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAuditSecurityGroupResponse), nil
	}
}

// UpdateAuditSecurityGroupInvoker 修改安全组
func (c *DbssClient) UpdateAuditSecurityGroupInvoker(request *model.UpdateAuditSecurityGroupRequest) *UpdateAuditSecurityGroupInvoker {
	requestDef := GenReqDefForUpdateAuditSecurityGroup()
	return &UpdateAuditSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
