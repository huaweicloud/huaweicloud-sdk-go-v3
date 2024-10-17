package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dbss/v1/model"
)

type DbssClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDbssClient(hcClient *httpclient.HcHttpClient) *DbssClient {
	return &DbssClient{HcClient: hcClient}
}

func DbssClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AddAuditDatabase 添加自建数据库
//
// 添加自建数据库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) AddAuditDatabase(request *model.AddAuditDatabaseRequest) (*model.AddAuditDatabaseResponse, error) {
	requestDef := GenReqDefForAddAuditDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAuditDatabaseResponse), nil
	}
}

// AddAuditDatabaseInvoker 添加自建数据库
func (c *DbssClient) AddAuditDatabaseInvoker(request *model.AddAuditDatabaseRequest) *AddAuditDatabaseInvoker {
	requestDef := GenReqDefForAddAuditDatabase()
	return &AddAuditDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddRdsDatabase 添加RDS数据库
//
// 添加RDS数据库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) AddRdsDatabase(request *model.AddRdsDatabaseRequest) (*model.AddRdsDatabaseResponse, error) {
	requestDef := GenReqDefForAddRdsDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddRdsDatabaseResponse), nil
	}
}

// AddRdsDatabaseInvoker 添加RDS数据库
func (c *DbssClient) AddRdsDatabaseInvoker(request *model.AddRdsDatabaseRequest) *AddRdsDatabaseInvoker {
	requestDef := GenReqDefForAddRdsDatabase()
	return &AddRdsDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddRdsNoAgentDatabase 添加RDS数据库(V1待下线)
//
// 添加RDS数据库。V1版本已不再维护，待下线。
// 请使用V2版本接口（/v2/{project_id}/{instance_id}/audit/databases/rds）。
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

// AddRdsNoAgentDatabaseInvoker 添加RDS数据库(V1待下线)
func (c *DbssClient) AddRdsNoAgentDatabaseInvoker(request *model.AddRdsNoAgentDatabaseRequest) *AddRdsNoAgentDatabaseInvoker {
	requestDef := GenReqDefForAddRdsNoAgentDatabase()
	return &AddRdsNoAgentDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// DeleteAuditDatabase 删除数据库
//
// 删除数据库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) DeleteAuditDatabase(request *model.DeleteAuditDatabaseRequest) (*model.DeleteAuditDatabaseResponse, error) {
	requestDef := GenReqDefForDeleteAuditDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAuditDatabaseResponse), nil
	}
}

// DeleteAuditDatabaseInvoker 删除数据库
func (c *DbssClient) DeleteAuditDatabaseInvoker(request *model.DeleteAuditDatabaseRequest) *DeleteAuditDatabaseInvoker {
	requestDef := GenReqDefForDeleteAuditDatabase()
	return &DeleteAuditDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstances 删除审计实例
//
// 只有按需计费模式实例没有任务时 或 包周期模式实例状态为故障时，才能执行此操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) DeleteInstances(request *model.DeleteInstancesRequest) (*model.DeleteInstancesResponse, error) {
	requestDef := GenReqDefForDeleteInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstancesResponse), nil
	}
}

// DeleteInstancesInvoker 删除审计实例
func (c *DbssClient) DeleteInstancesInvoker(request *model.DeleteInstancesRequest) *DeleteInstancesInvoker {
	requestDef := GenReqDefForDeleteInstances()
	return &DeleteInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditAlarmLog 查询审计告警信息
//
// 查询审计告警信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditAlarmLog(request *model.ListAuditAlarmLogRequest) (*model.ListAuditAlarmLogResponse, error) {
	requestDef := GenReqDefForListAuditAlarmLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditAlarmLogResponse), nil
	}
}

// ListAuditAlarmLogInvoker 查询审计告警信息
func (c *DbssClient) ListAuditAlarmLogInvoker(request *model.ListAuditAlarmLogRequest) *ListAuditAlarmLogInvoker {
	requestDef := GenReqDefForListAuditAlarmLog()
	return &ListAuditAlarmLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListAuditSqls 查询审计SQL语句
//
// 查询审计SQL语句
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditSqls(request *model.ListAuditSqlsRequest) (*model.ListAuditSqlsResponse, error) {
	requestDef := GenReqDefForListAuditSqls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditSqlsResponse), nil
	}
}

// ListAuditSqlsInvoker 查询审计SQL语句
func (c *DbssClient) ListAuditSqlsInvoker(request *model.ListAuditSqlsRequest) *ListAuditSqlsInvoker {
	requestDef := GenReqDefForListAuditSqls()
	return &ListAuditSqlsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditSummaryInfos 查询审计汇总信息
//
// 查询审计汇总信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditSummaryInfos(request *model.ListAuditSummaryInfosRequest) (*model.ListAuditSummaryInfosResponse, error) {
	requestDef := GenReqDefForListAuditSummaryInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditSummaryInfosResponse), nil
	}
}

// ListAuditSummaryInfosInvoker 查询审计汇总信息
func (c *DbssClient) ListAuditSummaryInfosInvoker(request *model.ListAuditSummaryInfosRequest) *ListAuditSummaryInfosInvoker {
	requestDef := GenReqDefForListAuditSummaryInfos()
	return &ListAuditSummaryInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListEcsSpecification 查询ECS服务器规格信息
//
// 查询ECS服务器规格信息
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

// ListEcsSpecificationInvoker 查询ECS服务器规格信息
func (c *DbssClient) ListEcsSpecificationInvoker(request *model.ListEcsSpecificationRequest) *ListEcsSpecificationInvoker {
	requestDef := GenReqDefForListEcsSpecification()
	return &ListEcsSpecificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRdsDatabases 查询RDS数据库列表
//
// 查询RDS数据库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListRdsDatabases(request *model.ListRdsDatabasesRequest) (*model.ListRdsDatabasesResponse, error) {
	requestDef := GenReqDefForListRdsDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRdsDatabasesResponse), nil
	}
}

// ListRdsDatabasesInvoker 查询RDS数据库列表
func (c *DbssClient) ListRdsDatabasesInvoker(request *model.ListRdsDatabasesRequest) *ListRdsDatabasesInvoker {
	requestDef := GenReqDefForListRdsDatabases()
	return &ListRdsDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// RebootAuditInstance 重启审计实例
//
// 重启审计实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) RebootAuditInstance(request *model.RebootAuditInstanceRequest) (*model.RebootAuditInstanceResponse, error) {
	requestDef := GenReqDefForRebootAuditInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RebootAuditInstanceResponse), nil
	}
}

// RebootAuditInstanceInvoker 重启审计实例
func (c *DbssClient) RebootAuditInstanceInvoker(request *model.RebootAuditInstanceRequest) *RebootAuditInstanceInvoker {
	requestDef := GenReqDefForRebootAuditInstance()
	return &RebootAuditInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// StartAuditInstance 开启审计实例
//
// 开启审计实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) StartAuditInstance(request *model.StartAuditInstanceRequest) (*model.StartAuditInstanceResponse, error) {
	requestDef := GenReqDefForStartAuditInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartAuditInstanceResponse), nil
	}
}

// StartAuditInstanceInvoker 开启审计实例
func (c *DbssClient) StartAuditInstanceInvoker(request *model.StartAuditInstanceRequest) *StartAuditInstanceInvoker {
	requestDef := GenReqDefForStartAuditInstance()
	return &StartAuditInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopAuditInstance 关闭审计实例
//
// 关闭审计实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) StopAuditInstance(request *model.StopAuditInstanceRequest) (*model.StopAuditInstanceResponse, error) {
	requestDef := GenReqDefForStopAuditInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopAuditInstanceResponse), nil
	}
}

// StopAuditInstanceInvoker 关闭审计实例
func (c *DbssClient) StopAuditInstanceInvoker(request *model.StopAuditInstanceRequest) *StopAuditInstanceInvoker {
	requestDef := GenReqDefForStopAuditInstance()
	return &StopAuditInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchAuditDatabase 开启关闭数据库
//
// 开启关闭数据库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SwitchAuditDatabase(request *model.SwitchAuditDatabaseRequest) (*model.SwitchAuditDatabaseResponse, error) {
	requestDef := GenReqDefForSwitchAuditDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchAuditDatabaseResponse), nil
	}
}

// SwitchAuditDatabaseInvoker 开启关闭数据库
func (c *DbssClient) SwitchAuditDatabaseInvoker(request *model.SwitchAuditDatabaseRequest) *SwitchAuditDatabaseInvoker {
	requestDef := GenReqDefForSwitchAuditDatabase()
	return &SwitchAuditDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// UpdateAuditInstance 更新审计实例信息
//
// 更新审计实例信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) UpdateAuditInstance(request *model.UpdateAuditInstanceRequest) (*model.UpdateAuditInstanceResponse, error) {
	requestDef := GenReqDefForUpdateAuditInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAuditInstanceResponse), nil
	}
}

// UpdateAuditInstanceInvoker 更新审计实例信息
func (c *DbssClient) UpdateAuditInstanceInvoker(request *model.UpdateAuditInstanceRequest) *UpdateAuditInstanceInvoker {
	requestDef := GenReqDefForUpdateAuditInstance()
	return &UpdateAuditInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAuditSecurityGroup 修改实例安全组
//
// 修改实例安全组
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

// UpdateAuditSecurityGroupInvoker 修改实例安全组
func (c *DbssClient) UpdateAuditSecurityGroupInvoker(request *model.UpdateAuditSecurityGroupRequest) *UpdateAuditSecurityGroupInvoker {
	requestDef := GenReqDefForUpdateAuditSecurityGroup()
	return &UpdateAuditSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddAuditAgent 添加审计数据库Agent
//
// 添加审计数据库Agent
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) AddAuditAgent(request *model.AddAuditAgentRequest) (*model.AddAuditAgentResponse, error) {
	requestDef := GenReqDefForAddAuditAgent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAuditAgentResponse), nil
	}
}

// AddAuditAgentInvoker 添加审计数据库Agent
func (c *DbssClient) AddAuditAgentInvoker(request *model.AddAuditAgentRequest) *AddAuditAgentInvoker {
	requestDef := GenReqDefForAddAuditAgent()
	return &AddAuditAgentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAuditAgent 删除数据库Agent
//
// 删除数据库Agent
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) DeleteAuditAgent(request *model.DeleteAuditAgentRequest) (*model.DeleteAuditAgentResponse, error) {
	requestDef := GenReqDefForDeleteAuditAgent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAuditAgentResponse), nil
	}
}

// DeleteAuditAgentInvoker 删除数据库Agent
func (c *DbssClient) DeleteAuditAgentInvoker(request *model.DeleteAuditAgentRequest) *DeleteAuditAgentInvoker {
	requestDef := GenReqDefForDeleteAuditAgent()
	return &DeleteAuditAgentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadAuditAgent 下载审计Agent
//
// 下载审计Agent
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) DownloadAuditAgent(request *model.DownloadAuditAgentRequest) (*model.DownloadAuditAgentResponse, error) {
	requestDef := GenReqDefForDownloadAuditAgent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadAuditAgentResponse), nil
	}
}

// DownloadAuditAgentInvoker 下载审计Agent
func (c *DbssClient) DownloadAuditAgentInvoker(request *model.DownloadAuditAgentRequest) *DownloadAuditAgentInvoker {
	requestDef := GenReqDefForDownloadAuditAgent()
	return &DownloadAuditAgentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditAgent 查询数据库Agent列表
//
// 查询数据库Agent列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditAgent(request *model.ListAuditAgentRequest) (*model.ListAuditAgentResponse, error) {
	requestDef := GenReqDefForListAuditAgent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditAgentResponse), nil
	}
}

// ListAuditAgentInvoker 查询数据库Agent列表
func (c *DbssClient) ListAuditAgentInvoker(request *model.ListAuditAgentRequest) *ListAuditAgentInvoker {
	requestDef := GenReqDefForListAuditAgent()
	return &ListAuditAgentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchAgent 开启关闭Agent
//
// 用于开启和关闭Agent审计的功能，当开启后，开始抓取用户的访问信息。
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
