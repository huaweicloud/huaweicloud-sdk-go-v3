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

// Deprecated: This function is deprecated and will be removed in the future versions.
// AddAuditAgent 添加审计数据库Agent[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// AddAuditAgentInvoker 添加审计数据库Agent[待下线]
func (c *DbssClient) AddAuditAgentInvoker(request *model.AddAuditAgentRequest) *AddAuditAgentInvoker {
	requestDef := GenReqDefForAddAuditAgent()
	return &AddAuditAgentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// AddAuditDatabase 添加自建数据库[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// AddAuditDatabaseInvoker 添加自建数据库[待下线]
func (c *DbssClient) AddAuditDatabaseInvoker(request *model.AddAuditDatabaseRequest) *AddAuditDatabaseInvoker {
	requestDef := GenReqDefForAddAuditDatabase()
	return &AddAuditDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddAuditDatabaseNew 添加自建数据库
//
// 添加自建数据库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) AddAuditDatabaseNew(request *model.AddAuditDatabaseNewRequest) (*model.AddAuditDatabaseNewResponse, error) {
	requestDef := GenReqDefForAddAuditDatabaseNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAuditDatabaseNewResponse), nil
	}
}

// AddAuditDatabaseNewInvoker 添加自建数据库
func (c *DbssClient) AddAuditDatabaseNewInvoker(request *model.AddAuditDatabaseNewRequest) *AddAuditDatabaseNewInvoker {
	requestDef := GenReqDefForAddAuditDatabaseNew()
	return &AddAuditDatabaseNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// AddRdsDatabase 添加RDS数据库[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// AddRdsDatabaseInvoker 添加RDS数据库[待下线]
func (c *DbssClient) AddRdsDatabaseInvoker(request *model.AddRdsDatabaseRequest) *AddRdsDatabaseInvoker {
	requestDef := GenReqDefForAddRdsDatabase()
	return &AddRdsDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddRdsDatabaseNew 添加RDS数据库
//
// 添加RDS数据库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) AddRdsDatabaseNew(request *model.AddRdsDatabaseNewRequest) (*model.AddRdsDatabaseNewResponse, error) {
	requestDef := GenReqDefForAddRdsDatabaseNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddRdsDatabaseNewResponse), nil
	}
}

// AddRdsDatabaseNewInvoker 添加RDS数据库
func (c *DbssClient) AddRdsDatabaseNewInvoker(request *model.AddRdsDatabaseNewRequest) *AddRdsDatabaseNewInvoker {
	requestDef := GenReqDefForAddRdsDatabaseNew()
	return &AddRdsDatabaseNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// AddRdsNoAgentDatabase 添加RDS数据库[待下线]
//
// 添加RDS数据库。V1版本已不再维护，待下线
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// AddRdsNoAgentDatabaseInvoker 添加RDS数据库[待下线]
func (c *DbssClient) AddRdsNoAgentDatabaseInvoker(request *model.AddRdsNoAgentDatabaseRequest) *AddRdsNoAgentDatabaseInvoker {
	requestDef := GenReqDefForAddRdsNoAgentDatabase()
	return &AddRdsNoAgentDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteAuditScope 审计范围规则操作-删除策略
//
// 审计范围规则操作-删除策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) BatchDeleteAuditScope(request *model.BatchDeleteAuditScopeRequest) (*model.BatchDeleteAuditScopeResponse, error) {
	requestDef := GenReqDefForBatchDeleteAuditScope()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAuditScopeResponse), nil
	}
}

// BatchDeleteAuditScopeInvoker 审计范围规则操作-删除策略
func (c *DbssClient) BatchDeleteAuditScopeInvoker(request *model.BatchDeleteAuditScopeRequest) *BatchDeleteAuditScopeInvoker {
	requestDef := GenReqDefForBatchDeleteAuditScope()
	return &BatchDeleteAuditScopeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSetAuditAlarmLogStatus 批量标记
//
// 批量标记
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) BatchSetAuditAlarmLogStatus(request *model.BatchSetAuditAlarmLogStatusRequest) (*model.BatchSetAuditAlarmLogStatusResponse, error) {
	requestDef := GenReqDefForBatchSetAuditAlarmLogStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetAuditAlarmLogStatusResponse), nil
	}
}

// BatchSetAuditAlarmLogStatusInvoker 批量标记
func (c *DbssClient) BatchSetAuditAlarmLogStatusInvoker(request *model.BatchSetAuditAlarmLogStatusRequest) *BatchSetAuditAlarmLogStatusInvoker {
	requestDef := GenReqDefForBatchSetAuditAlarmLogStatus()
	return &BatchSetAuditAlarmLogStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BindDbEncryptEip 绑定数据库加密实例的EIP
//
// 绑定数据库加密实例的EIP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) BindDbEncryptEip(request *model.BindDbEncryptEipRequest) (*model.BindDbEncryptEipResponse, error) {
	requestDef := GenReqDefForBindDbEncryptEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BindDbEncryptEipResponse), nil
	}
}

// BindDbEncryptEipInvoker 绑定数据库加密实例的EIP
func (c *DbssClient) BindDbEncryptEipInvoker(request *model.BindDbEncryptEipRequest) *BindDbEncryptEipInvoker {
	requestDef := GenReqDefForBindDbEncryptEip()
	return &BindDbEncryptEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BindDbOmEip 绑定数据库运维实例的EIP
//
// 绑定数据库运维实例的EIP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) BindDbOmEip(request *model.BindDbOmEipRequest) (*model.BindDbOmEipResponse, error) {
	requestDef := GenReqDefForBindDbOmEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BindDbOmEipResponse), nil
	}
}

// BindDbOmEipInvoker 绑定数据库运维实例的EIP
func (c *DbssClient) BindDbOmEipInvoker(request *model.BindDbOmEipRequest) *BindDbOmEipInvoker {
	requestDef := GenReqDefForBindDbOmEip()
	return &BindDbOmEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeDbEncryptSecurityGroup 更改数据库加密实例的安全组
//
// 更改数据库加密实例的安全组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ChangeDbEncryptSecurityGroup(request *model.ChangeDbEncryptSecurityGroupRequest) (*model.ChangeDbEncryptSecurityGroupResponse, error) {
	requestDef := GenReqDefForChangeDbEncryptSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeDbEncryptSecurityGroupResponse), nil
	}
}

// ChangeDbEncryptSecurityGroupInvoker 更改数据库加密实例的安全组
func (c *DbssClient) ChangeDbEncryptSecurityGroupInvoker(request *model.ChangeDbEncryptSecurityGroupRequest) *ChangeDbEncryptSecurityGroupInvoker {
	requestDef := GenReqDefForChangeDbEncryptSecurityGroup()
	return &ChangeDbEncryptSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeDbOmSecurityGroup 更改数据库运维实例的安全组
//
// 更改数据库运维实例的安全组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ChangeDbOmSecurityGroup(request *model.ChangeDbOmSecurityGroupRequest) (*model.ChangeDbOmSecurityGroupResponse, error) {
	requestDef := GenReqDefForChangeDbOmSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeDbOmSecurityGroupResponse), nil
	}
}

// ChangeDbOmSecurityGroupInvoker 更改数据库运维实例的安全组
func (c *DbssClient) ChangeDbOmSecurityGroupInvoker(request *model.ChangeDbOmSecurityGroupRequest) *ChangeDbOmSecurityGroupInvoker {
	requestDef := GenReqDefForChangeDbOmSecurityGroup()
	return &ChangeDbOmSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ConfirmUpgradeAudit 触发审计实例升级
//
// 触发审计实例升级
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ConfirmUpgradeAudit(request *model.ConfirmUpgradeAuditRequest) (*model.ConfirmUpgradeAuditResponse, error) {
	requestDef := GenReqDefForConfirmUpgradeAudit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ConfirmUpgradeAuditResponse), nil
	}
}

// ConfirmUpgradeAuditInvoker 触发审计实例升级
func (c *DbssClient) ConfirmUpgradeAuditInvoker(request *model.ConfirmUpgradeAuditRequest) *ConfirmUpgradeAuditInvoker {
	requestDef := GenReqDefForConfirmUpgradeAudit()
	return &ConfirmUpgradeAuditInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountDbAccountSession 查询数据库用户会话分布
//
// 查询数据库用户会话分布
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) CountDbAccountSession(request *model.CountDbAccountSessionRequest) (*model.CountDbAccountSessionResponse, error) {
	requestDef := GenReqDefForCountDbAccountSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountDbAccountSessionResponse), nil
	}
}

// CountDbAccountSessionInvoker 查询数据库用户会话分布
func (c *DbssClient) CountDbAccountSessionInvoker(request *model.CountDbAccountSessionRequest) *CountDbAccountSessionInvoker {
	requestDef := GenReqDefForCountDbAccountSession()
	return &CountDbAccountSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountDbClientSession 查询客户端会话分布
//
// 查询客户端会话分布
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) CountDbClientSession(request *model.CountDbClientSessionRequest) (*model.CountDbClientSessionResponse, error) {
	requestDef := GenReqDefForCountDbClientSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountDbClientSessionResponse), nil
	}
}

// CountDbClientSessionInvoker 查询客户端会话分布
func (c *DbssClient) CountDbClientSessionInvoker(request *model.CountDbClientSessionRequest) *CountDbClientSessionInvoker {
	requestDef := GenReqDefForCountDbClientSession()
	return &CountDbClientSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountInjectionStatistics 获取指定时间段内的sql注入分布统计
//
// 获取指定时间段内的sql注入分布统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) CountInjectionStatistics(request *model.CountInjectionStatisticsRequest) (*model.CountInjectionStatisticsResponse, error) {
	requestDef := GenReqDefForCountInjectionStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountInjectionStatisticsResponse), nil
	}
}

// CountInjectionStatisticsInvoker 获取指定时间段内的sql注入分布统计
func (c *DbssClient) CountInjectionStatisticsInvoker(request *model.CountInjectionStatisticsRequest) *CountInjectionStatisticsInvoker {
	requestDef := GenReqDefForCountInjectionStatistics()
	return &CountInjectionStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountOperationStatistics 获取指定时间段内的风险操作数量分布统计
//
// 获取指定时间段内的风险操作数量分布统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) CountOperationStatistics(request *model.CountOperationStatisticsRequest) (*model.CountOperationStatisticsResponse, error) {
	requestDef := GenReqDefForCountOperationStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountOperationStatisticsResponse), nil
	}
}

// CountOperationStatisticsInvoker 获取指定时间段内的风险操作数量分布统计
func (c *DbssClient) CountOperationStatisticsInvoker(request *model.CountOperationStatisticsRequest) *CountOperationStatisticsInvoker {
	requestDef := GenReqDefForCountOperationStatistics()
	return &CountOperationStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountRiskTrendStatistics 获取指定时间段内的风险分布统计
//
// 获取指定时间段内的风险分布统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) CountRiskTrendStatistics(request *model.CountRiskTrendStatisticsRequest) (*model.CountRiskTrendStatisticsResponse, error) {
	requestDef := GenReqDefForCountRiskTrendStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountRiskTrendStatisticsResponse), nil
	}
}

// CountRiskTrendStatisticsInvoker 获取指定时间段内的风险分布统计
func (c *DbssClient) CountRiskTrendStatisticsInvoker(request *model.CountRiskTrendStatisticsRequest) *CountRiskTrendStatisticsInvoker {
	requestDef := GenReqDefForCountRiskTrendStatistics()
	return &CountRiskTrendStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountSessionStatistics 获取指定时间段内的查询会话统计
//
// 获取指定时间段内的查询会话统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) CountSessionStatistics(request *model.CountSessionStatisticsRequest) (*model.CountSessionStatisticsResponse, error) {
	requestDef := GenReqDefForCountSessionStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountSessionStatisticsResponse), nil
	}
}

// CountSessionStatisticsInvoker 获取指定时间段内的查询会话统计
func (c *DbssClient) CountSessionStatisticsInvoker(request *model.CountSessionStatisticsRequest) *CountSessionStatisticsInvoker {
	requestDef := GenReqDefForCountSessionStatistics()
	return &CountSessionStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountSqlStatistics 获取指定时间段内的SQL分布统计
//
// 获取指定时间段内的SQL分布统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) CountSqlStatistics(request *model.CountSqlStatisticsRequest) (*model.CountSqlStatisticsResponse, error) {
	requestDef := GenReqDefForCountSqlStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountSqlStatisticsResponse), nil
	}
}

// CountSqlStatisticsInvoker 获取指定时间段内的SQL分布统计
func (c *DbssClient) CountSqlStatisticsInvoker(request *model.CountSqlStatisticsRequest) *CountSqlStatisticsInvoker {
	requestDef := GenReqDefForCountSqlStatistics()
	return &CountSqlStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountSqlTrendStatistics 获取指定时间段内的sql数量分布统计
//
// 获取指定时间段内的sql数量分布统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) CountSqlTrendStatistics(request *model.CountSqlTrendStatisticsRequest) (*model.CountSqlTrendStatisticsResponse, error) {
	requestDef := GenReqDefForCountSqlTrendStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountSqlTrendStatisticsResponse), nil
	}
}

// CountSqlTrendStatisticsInvoker 获取指定时间段内的sql数量分布统计
func (c *DbssClient) CountSqlTrendStatisticsInvoker(request *model.CountSqlTrendStatisticsRequest) *CountSqlTrendStatisticsInvoker {
	requestDef := GenReqDefForCountSqlTrendStatistics()
	return &CountSqlTrendStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAuditRiskRule 添加风险规则
//
// 添加风险规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) CreateAuditRiskRule(request *model.CreateAuditRiskRuleRequest) (*model.CreateAuditRiskRuleResponse, error) {
	requestDef := GenReqDefForCreateAuditRiskRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuditRiskRuleResponse), nil
	}
}

// CreateAuditRiskRuleInvoker 添加风险规则
func (c *DbssClient) CreateAuditRiskRuleInvoker(request *model.CreateAuditRiskRuleRequest) *CreateAuditRiskRuleInvoker {
	requestDef := GenReqDefForCreateAuditRiskRule()
	return &CreateAuditRiskRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAuditScopeRule 添加审计范围策略
//
// 添加审计范围策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) CreateAuditScopeRule(request *model.CreateAuditScopeRuleRequest) (*model.CreateAuditScopeRuleResponse, error) {
	requestDef := GenReqDefForCreateAuditScopeRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuditScopeRuleResponse), nil
	}
}

// CreateAuditScopeRuleInvoker 添加审计范围策略
func (c *DbssClient) CreateAuditScopeRuleInvoker(request *model.CreateAuditScopeRuleRequest) *CreateAuditScopeRuleInvoker {
	requestDef := GenReqDefForCreateAuditScopeRule()
	return &CreateAuditScopeRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDbEncryptInstancePeriod 按包周期方式购买数据库加密实例
//
// 按包周期方式购买数据库加密实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) CreateDbEncryptInstancePeriod(request *model.CreateDbEncryptInstancePeriodRequest) (*model.CreateDbEncryptInstancePeriodResponse, error) {
	requestDef := GenReqDefForCreateDbEncryptInstancePeriod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDbEncryptInstancePeriodResponse), nil
	}
}

// CreateDbEncryptInstancePeriodInvoker 按包周期方式购买数据库加密实例
func (c *DbssClient) CreateDbEncryptInstancePeriodInvoker(request *model.CreateDbEncryptInstancePeriodRequest) *CreateDbEncryptInstancePeriodInvoker {
	requestDef := GenReqDefForCreateDbEncryptInstancePeriod()
	return &CreateDbEncryptInstancePeriodInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDbOmInstancePeriod 包周期购买数据库运维实例
//
// 包周期购买数据库运维实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) CreateDbOmInstancePeriod(request *model.CreateDbOmInstancePeriodRequest) (*model.CreateDbOmInstancePeriodResponse, error) {
	requestDef := GenReqDefForCreateDbOmInstancePeriod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDbOmInstancePeriodResponse), nil
	}
}

// CreateDbOmInstancePeriodInvoker 包周期购买数据库运维实例
func (c *DbssClient) CreateDbOmInstancePeriodInvoker(request *model.CreateDbOmInstancePeriodRequest) *CreateDbOmInstancePeriodInvoker {
	requestDef := GenReqDefForCreateDbOmInstancePeriod()
	return &CreateDbOmInstancePeriodInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// CreateInstancesPeriodOrder 包年包月计费模式创建审计实例[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// CreateInstancesPeriodOrderInvoker 包年包月计费模式创建审计实例[待下线]
func (c *DbssClient) CreateInstancesPeriodOrderInvoker(request *model.CreateInstancesPeriodOrderRequest) *CreateInstancesPeriodOrderInvoker {
	requestDef := GenReqDefForCreateInstancesPeriodOrder()
	return &CreateInstancesPeriodOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstancesPeriodOrderNew 包年包月计费模式创建审计实例
//
// 包年包月计费模式创建审计实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) CreateInstancesPeriodOrderNew(request *model.CreateInstancesPeriodOrderNewRequest) (*model.CreateInstancesPeriodOrderNewResponse, error) {
	requestDef := GenReqDefForCreateInstancesPeriodOrderNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstancesPeriodOrderNewResponse), nil
	}
}

// CreateInstancesPeriodOrderNewInvoker 包年包月计费模式创建审计实例
func (c *DbssClient) CreateInstancesPeriodOrderNewInvoker(request *model.CreateInstancesPeriodOrderNewRequest) *CreateInstancesPeriodOrderNewInvoker {
	requestDef := GenReqDefForCreateInstancesPeriodOrderNew()
	return &CreateInstancesPeriodOrderNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateReport 立即生成报表
//
// 立即生成报表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) CreateReport(request *model.CreateReportRequest) (*model.CreateReportResponse, error) {
	requestDef := GenReqDefForCreateReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateReportResponse), nil
	}
}

// CreateReportInvoker 立即生成报表
func (c *DbssClient) CreateReportInvoker(request *model.CreateReportRequest) *CreateReportInvoker {
	requestDef := GenReqDefForCreateReport()
	return &CreateReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSensitiveMaskRule 增加隐私数据脱敏规则
//
// 增加隐私数据脱敏规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) CreateSensitiveMaskRule(request *model.CreateSensitiveMaskRuleRequest) (*model.CreateSensitiveMaskRuleResponse, error) {
	requestDef := GenReqDefForCreateSensitiveMaskRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSensitiveMaskRuleResponse), nil
	}
}

// CreateSensitiveMaskRuleInvoker 增加隐私数据脱敏规则
func (c *DbssClient) CreateSensitiveMaskRuleInvoker(request *model.CreateSensitiveMaskRuleRequest) *CreateSensitiveMaskRuleInvoker {
	requestDef := GenReqDefForCreateSensitiveMaskRule()
	return &CreateSensitiveMaskRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteAuditAgent 删除数据库Agent[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteAuditAgentInvoker 删除数据库Agent[待下线]
func (c *DbssClient) DeleteAuditAgentInvoker(request *model.DeleteAuditAgentRequest) *DeleteAuditAgentInvoker {
	requestDef := GenReqDefForDeleteAuditAgent()
	return &DeleteAuditAgentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAuditAlarmLog 删除告警监控记录
//
// 删除告警监控记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) DeleteAuditAlarmLog(request *model.DeleteAuditAlarmLogRequest) (*model.DeleteAuditAlarmLogResponse, error) {
	requestDef := GenReqDefForDeleteAuditAlarmLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAuditAlarmLogResponse), nil
	}
}

// DeleteAuditAlarmLogInvoker 删除告警监控记录
func (c *DbssClient) DeleteAuditAlarmLogInvoker(request *model.DeleteAuditAlarmLogRequest) *DeleteAuditAlarmLogInvoker {
	requestDef := GenReqDefForDeleteAuditAlarmLog()
	return &DeleteAuditAlarmLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAuditBackupTask 删除指定备份任务
//
// 删除指定备份任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) DeleteAuditBackupTask(request *model.DeleteAuditBackupTaskRequest) (*model.DeleteAuditBackupTaskResponse, error) {
	requestDef := GenReqDefForDeleteAuditBackupTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAuditBackupTaskResponse), nil
	}
}

// DeleteAuditBackupTaskInvoker 删除指定备份任务
func (c *DbssClient) DeleteAuditBackupTaskInvoker(request *model.DeleteAuditBackupTaskRequest) *DeleteAuditBackupTaskInvoker {
	requestDef := GenReqDefForDeleteAuditBackupTask()
	return &DeleteAuditBackupTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteAuditDatabase 删除数据库[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteAuditDatabaseInvoker 删除数据库[待下线]
func (c *DbssClient) DeleteAuditDatabaseInvoker(request *model.DeleteAuditDatabaseRequest) *DeleteAuditDatabaseInvoker {
	requestDef := GenReqDefForDeleteAuditDatabase()
	return &DeleteAuditDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAuditDatabaseNew 删除数据库
//
// 删除数据库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) DeleteAuditDatabaseNew(request *model.DeleteAuditDatabaseNewRequest) (*model.DeleteAuditDatabaseNewResponse, error) {
	requestDef := GenReqDefForDeleteAuditDatabaseNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAuditDatabaseNewResponse), nil
	}
}

// DeleteAuditDatabaseNewInvoker 删除数据库
func (c *DbssClient) DeleteAuditDatabaseNewInvoker(request *model.DeleteAuditDatabaseNewRequest) *DeleteAuditDatabaseNewInvoker {
	requestDef := GenReqDefForDeleteAuditDatabaseNew()
	return &DeleteAuditDatabaseNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAuditReport 删除报表
//
// 删除报表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) DeleteAuditReport(request *model.DeleteAuditReportRequest) (*model.DeleteAuditReportResponse, error) {
	requestDef := GenReqDefForDeleteAuditReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAuditReportResponse), nil
	}
}

// DeleteAuditReportInvoker 删除报表
func (c *DbssClient) DeleteAuditReportInvoker(request *model.DeleteAuditReportRequest) *DeleteAuditReportInvoker {
	requestDef := GenReqDefForDeleteAuditReport()
	return &DeleteAuditReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAuditRuleRisk 删除风险策略
//
// 删除风险策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) DeleteAuditRuleRisk(request *model.DeleteAuditRuleRiskRequest) (*model.DeleteAuditRuleRiskResponse, error) {
	requestDef := GenReqDefForDeleteAuditRuleRisk()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAuditRuleRiskResponse), nil
	}
}

// DeleteAuditRuleRiskInvoker 删除风险策略
func (c *DbssClient) DeleteAuditRuleRiskInvoker(request *model.DeleteAuditRuleRiskRequest) *DeleteAuditRuleRiskInvoker {
	requestDef := GenReqDefForDeleteAuditRuleRisk()
	return &DeleteAuditRuleRiskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAuditScope 删除审计范围策略
//
// 删除审计范围策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) DeleteAuditScope(request *model.DeleteAuditScopeRequest) (*model.DeleteAuditScopeResponse, error) {
	requestDef := GenReqDefForDeleteAuditScope()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAuditScopeResponse), nil
	}
}

// DeleteAuditScopeInvoker 删除审计范围策略
func (c *DbssClient) DeleteAuditScopeInvoker(request *model.DeleteAuditScopeRequest) *DeleteAuditScopeInvoker {
	requestDef := GenReqDefForDeleteAuditScope()
	return &DeleteAuditScopeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDbEncryptInstance 删除数据库加密实例
//
// 删除数据库加密实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) DeleteDbEncryptInstance(request *model.DeleteDbEncryptInstanceRequest) (*model.DeleteDbEncryptInstanceResponse, error) {
	requestDef := GenReqDefForDeleteDbEncryptInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDbEncryptInstanceResponse), nil
	}
}

// DeleteDbEncryptInstanceInvoker 删除数据库加密实例
func (c *DbssClient) DeleteDbEncryptInstanceInvoker(request *model.DeleteDbEncryptInstanceRequest) *DeleteDbEncryptInstanceInvoker {
	requestDef := GenReqDefForDeleteDbEncryptInstance()
	return &DeleteDbEncryptInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDbOmInstance 删除数据库运维实例
//
// 删除数据库运维实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) DeleteDbOmInstance(request *model.DeleteDbOmInstanceRequest) (*model.DeleteDbOmInstanceResponse, error) {
	requestDef := GenReqDefForDeleteDbOmInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDbOmInstanceResponse), nil
	}
}

// DeleteDbOmInstanceInvoker 删除数据库运维实例
func (c *DbssClient) DeleteDbOmInstanceInvoker(request *model.DeleteDbOmInstanceRequest) *DeleteDbOmInstanceInvoker {
	requestDef := GenReqDefForDeleteDbOmInstance()
	return &DeleteDbOmInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteInstances 删除审计实例[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteInstancesInvoker 删除审计实例[待下线]
func (c *DbssClient) DeleteInstancesInvoker(request *model.DeleteInstancesRequest) *DeleteInstancesInvoker {
	requestDef := GenReqDefForDeleteInstances()
	return &DeleteInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstancesNew 删除审计实例
//
// 只有按需计费模式实例没有任务时 或 包周期模式实例状态为故障时，才能执行此操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) DeleteInstancesNew(request *model.DeleteInstancesNewRequest) (*model.DeleteInstancesNewResponse, error) {
	requestDef := GenReqDefForDeleteInstancesNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstancesNewResponse), nil
	}
}

// DeleteInstancesNewInvoker 删除审计实例
func (c *DbssClient) DeleteInstancesNewInvoker(request *model.DeleteInstancesNewRequest) *DeleteInstancesNewInvoker {
	requestDef := GenReqDefForDeleteInstancesNew()
	return &DeleteInstancesNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSensitiveRules 删除隐私数据脱敏规则
//
// 删除隐私数据脱敏规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) DeleteSensitiveRules(request *model.DeleteSensitiveRulesRequest) (*model.DeleteSensitiveRulesResponse, error) {
	requestDef := GenReqDefForDeleteSensitiveRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSensitiveRulesResponse), nil
	}
}

// DeleteSensitiveRulesInvoker 删除隐私数据脱敏规则
func (c *DbssClient) DeleteSensitiveRulesInvoker(request *model.DeleteSensitiveRulesRequest) *DeleteSensitiveRulesInvoker {
	requestDef := GenReqDefForDeleteSensitiveRules()
	return &DeleteSensitiveRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// DownloadAuditAgent 下载审计Agent[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// DownloadAuditAgentInvoker 下载审计Agent[待下线]
func (c *DbssClient) DownloadAuditAgentInvoker(request *model.DownloadAuditAgentRequest) *DownloadAuditAgentInvoker {
	requestDef := GenReqDefForDownloadAuditAgent()
	return &DownloadAuditAgentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadAuditReport 下载报表
//
// 下载报表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) DownloadAuditReport(request *model.DownloadAuditReportRequest) (*model.DownloadAuditReportResponse, error) {
	requestDef := GenReqDefForDownloadAuditReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadAuditReportResponse), nil
	}
}

// DownloadAuditReportInvoker 下载报表
func (c *DbssClient) DownloadAuditReportInvoker(request *model.DownloadAuditReportRequest) *DownloadAuditReportInvoker {
	requestDef := GenReqDefForDownloadAuditReport()
	return &DownloadAuditReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAlarmTopicConfigInfo 获取实例告警配置[待下线]
//
// 获取实例告警配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAlarmTopicConfigInfo(request *model.ListAlarmTopicConfigInfoRequest) (*model.ListAlarmTopicConfigInfoResponse, error) {
	requestDef := GenReqDefForListAlarmTopicConfigInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmTopicConfigInfoResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAlarmTopicConfigInfoInvoker 获取实例告警配置[待下线]
func (c *DbssClient) ListAlarmTopicConfigInfoInvoker(request *model.ListAlarmTopicConfigInfoRequest) *ListAlarmTopicConfigInfoInvoker {
	requestDef := GenReqDefForListAlarmTopicConfigInfo()
	return &ListAlarmTopicConfigInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmTopicConfigInfoNew 获取实例告警配置
//
// 获取实例告警配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAlarmTopicConfigInfoNew(request *model.ListAlarmTopicConfigInfoNewRequest) (*model.ListAlarmTopicConfigInfoNewResponse, error) {
	requestDef := GenReqDefForListAlarmTopicConfigInfoNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmTopicConfigInfoNewResponse), nil
	}
}

// ListAlarmTopicConfigInfoNewInvoker 获取实例告警配置
func (c *DbssClient) ListAlarmTopicConfigInfoNewInvoker(request *model.ListAlarmTopicConfigInfoNewRequest) *ListAlarmTopicConfigInfoNewInvoker {
	requestDef := GenReqDefForListAlarmTopicConfigInfoNew()
	return &ListAlarmTopicConfigInfoNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditAgent 查询数据库Agent列表[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditAgentInvoker 查询数据库Agent列表[待下线]
func (c *DbssClient) ListAuditAgentInvoker(request *model.ListAuditAgentRequest) *ListAuditAgentInvoker {
	requestDef := GenReqDefForListAuditAgent()
	return &ListAuditAgentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditAlarmLog 查询审计告警信息[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditAlarmLogInvoker 查询审计告警信息[待下线]
func (c *DbssClient) ListAuditAlarmLogInvoker(request *model.ListAuditAlarmLogRequest) *ListAuditAlarmLogInvoker {
	requestDef := GenReqDefForListAuditAlarmLog()
	return &ListAuditAlarmLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditAlarmLogNew 查询审计告警信息
//
// 查询审计告警信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditAlarmLogNew(request *model.ListAuditAlarmLogNewRequest) (*model.ListAuditAlarmLogNewResponse, error) {
	requestDef := GenReqDefForListAuditAlarmLogNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditAlarmLogNewResponse), nil
	}
}

// ListAuditAlarmLogNewInvoker 查询审计告警信息
func (c *DbssClient) ListAuditAlarmLogNewInvoker(request *model.ListAuditAlarmLogNewRequest) *ListAuditAlarmLogNewInvoker {
	requestDef := GenReqDefForListAuditAlarmLogNew()
	return &ListAuditAlarmLogNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditBackupInfo 获取所有备份信息
//
// 获取所有备份信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditBackupInfo(request *model.ListAuditBackupInfoRequest) (*model.ListAuditBackupInfoResponse, error) {
	requestDef := GenReqDefForListAuditBackupInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditBackupInfoResponse), nil
	}
}

// ListAuditBackupInfoInvoker 获取所有备份信息
func (c *DbssClient) ListAuditBackupInfoInvoker(request *model.ListAuditBackupInfoRequest) *ListAuditBackupInfoInvoker {
	requestDef := GenReqDefForListAuditBackupInfo()
	return &ListAuditBackupInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditBackupRiskTemplates 获取风险导出配置列表
//
// 获取风险导出配置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditBackupRiskTemplates(request *model.ListAuditBackupRiskTemplatesRequest) (*model.ListAuditBackupRiskTemplatesResponse, error) {
	requestDef := GenReqDefForListAuditBackupRiskTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditBackupRiskTemplatesResponse), nil
	}
}

// ListAuditBackupRiskTemplatesInvoker 获取风险导出配置列表
func (c *DbssClient) ListAuditBackupRiskTemplatesInvoker(request *model.ListAuditBackupRiskTemplatesRequest) *ListAuditBackupRiskTemplatesInvoker {
	requestDef := GenReqDefForListAuditBackupRiskTemplates()
	return &ListAuditBackupRiskTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditDatabases 查询数据库列表[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditDatabasesInvoker 查询数据库列表[待下线]
func (c *DbssClient) ListAuditDatabasesInvoker(request *model.ListAuditDatabasesRequest) *ListAuditDatabasesInvoker {
	requestDef := GenReqDefForListAuditDatabases()
	return &ListAuditDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditDatabasesNew 查询数据库列表
//
// 查询数据库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditDatabasesNew(request *model.ListAuditDatabasesNewRequest) (*model.ListAuditDatabasesNewResponse, error) {
	requestDef := GenReqDefForListAuditDatabasesNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditDatabasesNewResponse), nil
	}
}

// ListAuditDatabasesNewInvoker 查询数据库列表
func (c *DbssClient) ListAuditDatabasesNewInvoker(request *model.ListAuditDatabasesNewRequest) *ListAuditDatabasesNewInvoker {
	requestDef := GenReqDefForListAuditDatabasesNew()
	return &ListAuditDatabasesNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditInstanceJobs 查询实例创建任务信息[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditInstanceJobsInvoker 查询实例创建任务信息[待下线]
func (c *DbssClient) ListAuditInstanceJobsInvoker(request *model.ListAuditInstanceJobsRequest) *ListAuditInstanceJobsInvoker {
	requestDef := GenReqDefForListAuditInstanceJobs()
	return &ListAuditInstanceJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditInstanceJobsNew 查询实例创建任务信息
//
// 查询实例创建任务信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditInstanceJobsNew(request *model.ListAuditInstanceJobsNewRequest) (*model.ListAuditInstanceJobsNewResponse, error) {
	requestDef := GenReqDefForListAuditInstanceJobsNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditInstanceJobsNewResponse), nil
	}
}

// ListAuditInstanceJobsNewInvoker 查询实例创建任务信息
func (c *DbssClient) ListAuditInstanceJobsNewInvoker(request *model.ListAuditInstanceJobsNewRequest) *ListAuditInstanceJobsNewInvoker {
	requestDef := GenReqDefForListAuditInstanceJobsNew()
	return &ListAuditInstanceJobsNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditInstances 查询审计实例列表[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditInstancesInvoker 查询审计实例列表[待下线]
func (c *DbssClient) ListAuditInstancesInvoker(request *model.ListAuditInstancesRequest) *ListAuditInstancesInvoker {
	requestDef := GenReqDefForListAuditInstances()
	return &ListAuditInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditInstancesNew 查询审计实例列表
//
// 查询审计实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditInstancesNew(request *model.ListAuditInstancesNewRequest) (*model.ListAuditInstancesNewResponse, error) {
	requestDef := GenReqDefForListAuditInstancesNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditInstancesNewResponse), nil
	}
}

// ListAuditInstancesNewInvoker 查询审计实例列表
func (c *DbssClient) ListAuditInstancesNewInvoker(request *model.ListAuditInstancesNewRequest) *ListAuditInstancesNewInvoker {
	requestDef := GenReqDefForListAuditInstancesNew()
	return &ListAuditInstancesNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditObsBuckets 备份或风险导出时，选择obs桶
//
// 备份或风险导出时，选择obs桶
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditObsBuckets(request *model.ListAuditObsBucketsRequest) (*model.ListAuditObsBucketsResponse, error) {
	requestDef := GenReqDefForListAuditObsBuckets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditObsBucketsResponse), nil
	}
}

// ListAuditObsBucketsInvoker 备份或风险导出时，选择obs桶
func (c *DbssClient) ListAuditObsBucketsInvoker(request *model.ListAuditObsBucketsRequest) *ListAuditObsBucketsInvoker {
	requestDef := GenReqDefForListAuditObsBuckets()
	return &ListAuditObsBucketsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditOperateLogs 查询用户操作日志信息[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditOperateLogsInvoker 查询用户操作日志信息[待下线]
func (c *DbssClient) ListAuditOperateLogsInvoker(request *model.ListAuditOperateLogsRequest) *ListAuditOperateLogsInvoker {
	requestDef := GenReqDefForListAuditOperateLogs()
	return &ListAuditOperateLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditOperateLogsNew 查询用户操作日志信息
//
// 查询用户操作日志信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditOperateLogsNew(request *model.ListAuditOperateLogsNewRequest) (*model.ListAuditOperateLogsNewResponse, error) {
	requestDef := GenReqDefForListAuditOperateLogsNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditOperateLogsNewResponse), nil
	}
}

// ListAuditOperateLogsNewInvoker 查询用户操作日志信息
func (c *DbssClient) ListAuditOperateLogsNewInvoker(request *model.ListAuditOperateLogsNewRequest) *ListAuditOperateLogsNewInvoker {
	requestDef := GenReqDefForListAuditOperateLogsNew()
	return &ListAuditOperateLogsNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditReportTemplates 获取报表模板
//
// 获取报表模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditReportTemplates(request *model.ListAuditReportTemplatesRequest) (*model.ListAuditReportTemplatesResponse, error) {
	requestDef := GenReqDefForListAuditReportTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditReportTemplatesResponse), nil
	}
}

// ListAuditReportTemplatesInvoker 获取报表模板
func (c *DbssClient) ListAuditReportTemplatesInvoker(request *model.ListAuditReportTemplatesRequest) *ListAuditReportTemplatesInvoker {
	requestDef := GenReqDefForListAuditReportTemplates()
	return &ListAuditReportTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditReports 查询报表
//
// 查询报表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditReports(request *model.ListAuditReportsRequest) (*model.ListAuditReportsResponse, error) {
	requestDef := GenReqDefForListAuditReports()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditReportsResponse), nil
	}
}

// ListAuditReportsInvoker 查询报表
func (c *DbssClient) ListAuditReportsInvoker(request *model.ListAuditReportsRequest) *ListAuditReportsInvoker {
	requestDef := GenReqDefForListAuditReports()
	return &ListAuditReportsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditRuleRisks 查询风险规则策略[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditRuleRisksInvoker 查询风险规则策略[待下线]
func (c *DbssClient) ListAuditRuleRisksInvoker(request *model.ListAuditRuleRisksRequest) *ListAuditRuleRisksInvoker {
	requestDef := GenReqDefForListAuditRuleRisks()
	return &ListAuditRuleRisksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditRuleRisksNew 查询风险规则策略
//
// 查询风险规则策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditRuleRisksNew(request *model.ListAuditRuleRisksNewRequest) (*model.ListAuditRuleRisksNewResponse, error) {
	requestDef := GenReqDefForListAuditRuleRisksNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditRuleRisksNewResponse), nil
	}
}

// ListAuditRuleRisksNewInvoker 查询风险规则策略
func (c *DbssClient) ListAuditRuleRisksNewInvoker(request *model.ListAuditRuleRisksNewRequest) *ListAuditRuleRisksNewInvoker {
	requestDef := GenReqDefForListAuditRuleRisksNew()
	return &ListAuditRuleRisksNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditRuleScopes 查询审计范围策略列表[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditRuleScopesInvoker 查询审计范围策略列表[待下线]
func (c *DbssClient) ListAuditRuleScopesInvoker(request *model.ListAuditRuleScopesRequest) *ListAuditRuleScopesInvoker {
	requestDef := GenReqDefForListAuditRuleScopes()
	return &ListAuditRuleScopesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditRuleScopesNew 查询审计范围策略列表
//
// 查询审计范围策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditRuleScopesNew(request *model.ListAuditRuleScopesNewRequest) (*model.ListAuditRuleScopesNewResponse, error) {
	requestDef := GenReqDefForListAuditRuleScopesNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditRuleScopesNewResponse), nil
	}
}

// ListAuditRuleScopesNewInvoker 查询审计范围策略列表
func (c *DbssClient) ListAuditRuleScopesNewInvoker(request *model.ListAuditRuleScopesNewRequest) *ListAuditRuleScopesNewInvoker {
	requestDef := GenReqDefForListAuditRuleScopesNew()
	return &ListAuditRuleScopesNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditSensitiveMasks 查询隐私数据脱敏规则[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditSensitiveMasksInvoker 查询隐私数据脱敏规则[待下线]
func (c *DbssClient) ListAuditSensitiveMasksInvoker(request *model.ListAuditSensitiveMasksRequest) *ListAuditSensitiveMasksInvoker {
	requestDef := GenReqDefForListAuditSensitiveMasks()
	return &ListAuditSensitiveMasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditSensitiveMasksNew 查询隐私数据脱敏规则
//
// 查询隐私数据脱敏规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditSensitiveMasksNew(request *model.ListAuditSensitiveMasksNewRequest) (*model.ListAuditSensitiveMasksNewResponse, error) {
	requestDef := GenReqDefForListAuditSensitiveMasksNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditSensitiveMasksNewResponse), nil
	}
}

// ListAuditSensitiveMasksNewInvoker 查询隐私数据脱敏规则
func (c *DbssClient) ListAuditSensitiveMasksNewInvoker(request *model.ListAuditSensitiveMasksNewRequest) *ListAuditSensitiveMasksNewInvoker {
	requestDef := GenReqDefForListAuditSensitiveMasksNew()
	return &ListAuditSensitiveMasksNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditSqls 查询审计SQL语句[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuditSqlsInvoker 查询审计SQL语句[待下线]
func (c *DbssClient) ListAuditSqlsInvoker(request *model.ListAuditSqlsRequest) *ListAuditSqlsInvoker {
	requestDef := GenReqDefForListAuditSqls()
	return &ListAuditSqlsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditSqlsNew 查询审计SQL语句
//
// 查询审计SQL语句
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditSqlsNew(request *model.ListAuditSqlsNewRequest) (*model.ListAuditSqlsNewResponse, error) {
	requestDef := GenReqDefForListAuditSqlsNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditSqlsNewResponse), nil
	}
}

// ListAuditSqlsNewInvoker 查询审计SQL语句
func (c *DbssClient) ListAuditSqlsNewInvoker(request *model.ListAuditSqlsNewRequest) *ListAuditSqlsNewInvoker {
	requestDef := GenReqDefForListAuditSqlsNew()
	return &ListAuditSqlsNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListAuditTrendHistory 查询趋势历史数据
//
// 查询趋势历史数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditTrendHistory(request *model.ListAuditTrendHistoryRequest) (*model.ListAuditTrendHistoryResponse, error) {
	requestDef := GenReqDefForListAuditTrendHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditTrendHistoryResponse), nil
	}
}

// ListAuditTrendHistoryInvoker 查询趋势历史数据
func (c *DbssClient) ListAuditTrendHistoryInvoker(request *model.ListAuditTrendHistoryRequest) *ListAuditTrendHistoryInvoker {
	requestDef := GenReqDefForListAuditTrendHistory()
	return &ListAuditTrendHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAvailabilityZoneInfos 查询可用区信息[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAvailabilityZoneInfosInvoker 查询可用区信息[待下线]
func (c *DbssClient) ListAvailabilityZoneInfosInvoker(request *model.ListAvailabilityZoneInfosRequest) *ListAvailabilityZoneInfosInvoker {
	requestDef := GenReqDefForListAvailabilityZoneInfos()
	return &ListAvailabilityZoneInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailabilityZoneInfosNew 查询可用区信息
//
// 查询可用区信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAvailabilityZoneInfosNew(request *model.ListAvailabilityZoneInfosNewRequest) (*model.ListAvailabilityZoneInfosNewResponse, error) {
	requestDef := GenReqDefForListAvailabilityZoneInfosNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailabilityZoneInfosNewResponse), nil
	}
}

// ListAvailabilityZoneInfosNewInvoker 查询可用区信息
func (c *DbssClient) ListAvailabilityZoneInfosNewInvoker(request *model.ListAvailabilityZoneInfosNewRequest) *ListAvailabilityZoneInfosNewInvoker {
	requestDef := GenReqDefForListAvailabilityZoneInfosNew()
	return &ListAvailabilityZoneInfosNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDbEncryptInstances 列举数据库加密实例
//
// 列举数据库加密实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListDbEncryptInstances(request *model.ListDbEncryptInstancesRequest) (*model.ListDbEncryptInstancesResponse, error) {
	requestDef := GenReqDefForListDbEncryptInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDbEncryptInstancesResponse), nil
	}
}

// ListDbEncryptInstancesInvoker 列举数据库加密实例
func (c *DbssClient) ListDbEncryptInstancesInvoker(request *model.ListDbEncryptInstancesRequest) *ListDbEncryptInstancesInvoker {
	requestDef := GenReqDefForListDbEncryptInstances()
	return &ListDbEncryptInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListEcsSpecification 查询ECS服务器规格信息[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListEcsSpecificationInvoker 查询ECS服务器规格信息[待下线]
func (c *DbssClient) ListEcsSpecificationInvoker(request *model.ListEcsSpecificationRequest) *ListEcsSpecificationInvoker {
	requestDef := GenReqDefForListEcsSpecification()
	return &ListEcsSpecificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEcsSpecificationNew 查询ECS服务器规格信息
//
// 查询ECS服务器规格信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListEcsSpecificationNew(request *model.ListEcsSpecificationNewRequest) (*model.ListEcsSpecificationNewResponse, error) {
	requestDef := GenReqDefForListEcsSpecificationNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEcsSpecificationNewResponse), nil
	}
}

// ListEcsSpecificationNewInvoker 查询ECS服务器规格信息
func (c *DbssClient) ListEcsSpecificationNewInvoker(request *model.ListEcsSpecificationNewRequest) *ListEcsSpecificationNewInvoker {
	requestDef := GenReqDefForListEcsSpecificationNew()
	return &ListEcsSpecificationNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 获取数据库运维集群实例
//
// 获取数据库运维集群实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 获取数据库运维集群实例
func (c *DbssClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListSqlInjectionRules 查询SQL注入规则策略[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListSqlInjectionRulesInvoker 查询SQL注入规则策略[待下线]
func (c *DbssClient) ListSqlInjectionRulesInvoker(request *model.ListSqlInjectionRulesRequest) *ListSqlInjectionRulesInvoker {
	requestDef := GenReqDefForListSqlInjectionRules()
	return &ListSqlInjectionRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// RebootAuditInstance 重启审计实例[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// RebootAuditInstanceInvoker 重启审计实例[待下线]
func (c *DbssClient) RebootAuditInstanceInvoker(request *model.RebootAuditInstanceRequest) *RebootAuditInstanceInvoker {
	requestDef := GenReqDefForRebootAuditInstance()
	return &RebootAuditInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RebootAuditInstanceNew 重启审计实例
//
// 重启审计实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) RebootAuditInstanceNew(request *model.RebootAuditInstanceNewRequest) (*model.RebootAuditInstanceNewResponse, error) {
	requestDef := GenReqDefForRebootAuditInstanceNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RebootAuditInstanceNewResponse), nil
	}
}

// RebootAuditInstanceNewInvoker 重启审计实例
func (c *DbssClient) RebootAuditInstanceNewInvoker(request *model.RebootAuditInstanceNewRequest) *RebootAuditInstanceNewInvoker {
	requestDef := GenReqDefForRebootAuditInstanceNew()
	return &RebootAuditInstanceNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RebootDbEncryptInstance 重启数据库加密实例
//
// 重启数据库加密实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) RebootDbEncryptInstance(request *model.RebootDbEncryptInstanceRequest) (*model.RebootDbEncryptInstanceResponse, error) {
	requestDef := GenReqDefForRebootDbEncryptInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RebootDbEncryptInstanceResponse), nil
	}
}

// RebootDbEncryptInstanceInvoker 重启数据库加密实例
func (c *DbssClient) RebootDbEncryptInstanceInvoker(request *model.RebootDbEncryptInstanceRequest) *RebootDbEncryptInstanceInvoker {
	requestDef := GenReqDefForRebootDbEncryptInstance()
	return &RebootDbEncryptInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RebootDbOmInstance 重启数据库运维实例
//
// 重启数据库运维实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) RebootDbOmInstance(request *model.RebootDbOmInstanceRequest) (*model.RebootDbOmInstanceResponse, error) {
	requestDef := GenReqDefForRebootDbOmInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RebootDbOmInstanceResponse), nil
	}
}

// RebootDbOmInstanceInvoker 重启数据库运维实例
func (c *DbssClient) RebootDbOmInstanceInvoker(request *model.RebootDbOmInstanceRequest) *RebootDbOmInstanceInvoker {
	requestDef := GenReqDefForRebootDbOmInstance()
	return &RebootDbOmInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetDbEncryptPassword 重置数据库加密实例的密码
//
// 重置数据库加密实例的密码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ResetDbEncryptPassword(request *model.ResetDbEncryptPasswordRequest) (*model.ResetDbEncryptPasswordResponse, error) {
	requestDef := GenReqDefForResetDbEncryptPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetDbEncryptPasswordResponse), nil
	}
}

// ResetDbEncryptPasswordInvoker 重置数据库加密实例的密码
func (c *DbssClient) ResetDbEncryptPasswordInvoker(request *model.ResetDbEncryptPasswordRequest) *ResetDbEncryptPasswordInvoker {
	requestDef := GenReqDefForResetDbEncryptPassword()
	return &ResetDbEncryptPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetDbOmPassword 重置数据库运维实例的密码
//
// 重置数据库运维实例的密码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ResetDbOmPassword(request *model.ResetDbOmPasswordRequest) (*model.ResetDbOmPasswordResponse, error) {
	requestDef := GenReqDefForResetDbOmPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetDbOmPasswordResponse), nil
	}
}

// ResetDbOmPasswordInvoker 重置数据库运维实例的密码
func (c *DbssClient) ResetDbOmPasswordInvoker(request *model.ResetDbOmPasswordRequest) *ResetDbOmPasswordInvoker {
	requestDef := GenReqDefForResetDbOmPassword()
	return &ResetDbOmPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreAuditBackup 恢复备份信息
//
// 恢复备份信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) RestoreAuditBackup(request *model.RestoreAuditBackupRequest) (*model.RestoreAuditBackupResponse, error) {
	requestDef := GenReqDefForRestoreAuditBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreAuditBackupResponse), nil
	}
}

// RestoreAuditBackupInvoker 恢复备份信息
func (c *DbssClient) RestoreAuditBackupInvoker(request *model.RestoreAuditBackupRequest) *RestoreAuditBackupInvoker {
	requestDef := GenReqDefForRestoreAuditBackup()
	return &RestoreAuditBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetryAuditBackupTask 重试备份任务
//
// 重试备份任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) RetryAuditBackupTask(request *model.RetryAuditBackupTaskRequest) (*model.RetryAuditBackupTaskResponse, error) {
	requestDef := GenReqDefForRetryAuditBackupTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryAuditBackupTaskResponse), nil
	}
}

// RetryAuditBackupTaskInvoker 重试备份任务
func (c *DbssClient) RetryAuditBackupTaskInvoker(request *model.RetryAuditBackupTaskRequest) *RetryAuditBackupTaskInvoker {
	requestDef := GenReqDefForRetryAuditBackupTask()
	return &RetryAuditBackupTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// SetAlarmTopicConfigInfo 设置实例告警配置[待下线]
//
// 设置实例告警配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SetAlarmTopicConfigInfo(request *model.SetAlarmTopicConfigInfoRequest) (*model.SetAlarmTopicConfigInfoResponse, error) {
	requestDef := GenReqDefForSetAlarmTopicConfigInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetAlarmTopicConfigInfoResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// SetAlarmTopicConfigInfoInvoker 设置实例告警配置[待下线]
func (c *DbssClient) SetAlarmTopicConfigInfoInvoker(request *model.SetAlarmTopicConfigInfoRequest) *SetAlarmTopicConfigInfoInvoker {
	requestDef := GenReqDefForSetAlarmTopicConfigInfo()
	return &SetAlarmTopicConfigInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetAlarmTopicConfigInfoNew 设置实例告警配置
//
// 设置实例告警配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SetAlarmTopicConfigInfoNew(request *model.SetAlarmTopicConfigInfoNewRequest) (*model.SetAlarmTopicConfigInfoNewResponse, error) {
	requestDef := GenReqDefForSetAlarmTopicConfigInfoNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetAlarmTopicConfigInfoNewResponse), nil
	}
}

// SetAlarmTopicConfigInfoNewInvoker 设置实例告警配置
func (c *DbssClient) SetAlarmTopicConfigInfoNewInvoker(request *model.SetAlarmTopicConfigInfoNewRequest) *SetAlarmTopicConfigInfoNewInvoker {
	requestDef := GenReqDefForSetAlarmTopicConfigInfoNew()
	return &SetAlarmTopicConfigInfoNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetAuditAlarmLogStatus 标记监控告警
//
// 标记监控告警
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SetAuditAlarmLogStatus(request *model.SetAuditAlarmLogStatusRequest) (*model.SetAuditAlarmLogStatusResponse, error) {
	requestDef := GenReqDefForSetAuditAlarmLogStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetAuditAlarmLogStatusResponse), nil
	}
}

// SetAuditAlarmLogStatusInvoker 标记监控告警
func (c *DbssClient) SetAuditAlarmLogStatusInvoker(request *model.SetAuditAlarmLogStatusRequest) *SetAuditAlarmLogStatusInvoker {
	requestDef := GenReqDefForSetAuditAlarmLogStatus()
	return &SetAuditAlarmLogStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetAuditAutoBackupTemplate 获取备份配置信息
//
// 设置备份配置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SetAuditAutoBackupTemplate(request *model.SetAuditAutoBackupTemplateRequest) (*model.SetAuditAutoBackupTemplateResponse, error) {
	requestDef := GenReqDefForSetAuditAutoBackupTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetAuditAutoBackupTemplateResponse), nil
	}
}

// SetAuditAutoBackupTemplateInvoker 获取备份配置信息
func (c *DbssClient) SetAuditAutoBackupTemplateInvoker(request *model.SetAuditAutoBackupTemplateRequest) *SetAuditAutoBackupTemplateInvoker {
	requestDef := GenReqDefForSetAuditAutoBackupTemplate()
	return &SetAuditAutoBackupTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetAuditBackupRiskSwitch 风险导出开关(按DomainId)
//
// 风险导出开关(按DomainId)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SetAuditBackupRiskSwitch(request *model.SetAuditBackupRiskSwitchRequest) (*model.SetAuditBackupRiskSwitchResponse, error) {
	requestDef := GenReqDefForSetAuditBackupRiskSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetAuditBackupRiskSwitchResponse), nil
	}
}

// SetAuditBackupRiskSwitchInvoker 风险导出开关(按DomainId)
func (c *DbssClient) SetAuditBackupRiskSwitchInvoker(request *model.SetAuditBackupRiskSwitchRequest) *SetAuditBackupRiskSwitchInvoker {
	requestDef := GenReqDefForSetAuditBackupRiskSwitch()
	return &SetAuditBackupRiskSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetAuditBackupSwitch 开启/关闭备份
//
// 开启/关闭备份
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SetAuditBackupSwitch(request *model.SetAuditBackupSwitchRequest) (*model.SetAuditBackupSwitchResponse, error) {
	requestDef := GenReqDefForSetAuditBackupSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetAuditBackupSwitchResponse), nil
	}
}

// SetAuditBackupSwitchInvoker 开启/关闭备份
func (c *DbssClient) SetAuditBackupSwitchInvoker(request *model.SetAuditBackupSwitchRequest) *SetAuditBackupSwitchInvoker {
	requestDef := GenReqDefForSetAuditBackupSwitch()
	return &SetAuditBackupSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetAuditInstanceTimeZone 配置审计实例时区信息
//
// 配置审计实例时区信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SetAuditInstanceTimeZone(request *model.SetAuditInstanceTimeZoneRequest) (*model.SetAuditInstanceTimeZoneResponse, error) {
	requestDef := GenReqDefForSetAuditInstanceTimeZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetAuditInstanceTimeZoneResponse), nil
	}
}

// SetAuditInstanceTimeZoneInvoker 配置审计实例时区信息
func (c *DbssClient) SetAuditInstanceTimeZoneInvoker(request *model.SetAuditInstanceTimeZoneRequest) *SetAuditInstanceTimeZoneInvoker {
	requestDef := GenReqDefForSetAuditInstanceTimeZone()
	return &SetAuditInstanceTimeZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetAuditScopeRuleSwitch 启用禁用策略
//
// 启用禁用策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SetAuditScopeRuleSwitch(request *model.SetAuditScopeRuleSwitchRequest) (*model.SetAuditScopeRuleSwitchResponse, error) {
	requestDef := GenReqDefForSetAuditScopeRuleSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetAuditScopeRuleSwitchResponse), nil
	}
}

// SetAuditScopeRuleSwitchInvoker 启用禁用策略
func (c *DbssClient) SetAuditScopeRuleSwitchInvoker(request *model.SetAuditScopeRuleSwitchRequest) *SetAuditScopeRuleSwitchInvoker {
	requestDef := GenReqDefForSetAuditScopeRuleSwitch()
	return &SetAuditScopeRuleSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetRiskOperationPolicy 编辑风险操作策略
//
// 编辑风险操作策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SetRiskOperationPolicy(request *model.SetRiskOperationPolicyRequest) (*model.SetRiskOperationPolicyResponse, error) {
	requestDef := GenReqDefForSetRiskOperationPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetRiskOperationPolicyResponse), nil
	}
}

// SetRiskOperationPolicyInvoker 编辑风险操作策略
func (c *DbssClient) SetRiskOperationPolicyInvoker(request *model.SetRiskOperationPolicyRequest) *SetRiskOperationPolicyInvoker {
	requestDef := GenReqDefForSetRiskOperationPolicy()
	return &SetRiskOperationPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetRiskRuleRank 审计实例风险规则排序
//
// 审计实例风险规则排序
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SetRiskRuleRank(request *model.SetRiskRuleRankRequest) (*model.SetRiskRuleRankResponse, error) {
	requestDef := GenReqDefForSetRiskRuleRank()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetRiskRuleRankResponse), nil
	}
}

// SetRiskRuleRankInvoker 审计实例风险规则排序
func (c *DbssClient) SetRiskRuleRankInvoker(request *model.SetRiskRuleRankRequest) *SetRiskRuleRankInvoker {
	requestDef := GenReqDefForSetRiskRuleRank()
	return &SetRiskRuleRankInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetSensitiveMaskRuleSwitch 启用禁用单条隐私数据脱敏规则
//
// 启用禁用隐私数据脱敏规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SetSensitiveMaskRuleSwitch(request *model.SetSensitiveMaskRuleSwitchRequest) (*model.SetSensitiveMaskRuleSwitchResponse, error) {
	requestDef := GenReqDefForSetSensitiveMaskRuleSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetSensitiveMaskRuleSwitchResponse), nil
	}
}

// SetSensitiveMaskRuleSwitchInvoker 启用禁用单条隐私数据脱敏规则
func (c *DbssClient) SetSensitiveMaskRuleSwitchInvoker(request *model.SetSensitiveMaskRuleSwitchRequest) *SetSensitiveMaskRuleSwitchInvoker {
	requestDef := GenReqDefForSetSensitiveMaskRuleSwitch()
	return &SetSensitiveMaskRuleSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetSensitiveResultSwitch 开启关闭结果集存储
//
// 开启关闭结果集存储
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SetSensitiveResultSwitch(request *model.SetSensitiveResultSwitchRequest) (*model.SetSensitiveResultSwitchResponse, error) {
	requestDef := GenReqDefForSetSensitiveResultSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetSensitiveResultSwitchResponse), nil
	}
}

// SetSensitiveResultSwitchInvoker 开启关闭结果集存储
func (c *DbssClient) SetSensitiveResultSwitchInvoker(request *model.SetSensitiveResultSwitchRequest) *SetSensitiveResultSwitchInvoker {
	requestDef := GenReqDefForSetSensitiveResultSwitch()
	return &SetSensitiveResultSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetSensitiveSwitch 开启关闭隐私数据脱敏功能
//
// 开启关闭隐私数据脱敏
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SetSensitiveSwitch(request *model.SetSensitiveSwitchRequest) (*model.SetSensitiveSwitchResponse, error) {
	requestDef := GenReqDefForSetSensitiveSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetSensitiveSwitchResponse), nil
	}
}

// SetSensitiveSwitchInvoker 开启关闭隐私数据脱敏功能
func (c *DbssClient) SetSensitiveSwitchInvoker(request *model.SetSensitiveSwitchRequest) *SetSensitiveSwitchInvoker {
	requestDef := GenReqDefForSetSensitiveSwitch()
	return &SetSensitiveSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuditBackRiskTemplate 获取单个风险导出配置
//
// 获取单个风险导出配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ShowAuditBackRiskTemplate(request *model.ShowAuditBackRiskTemplateRequest) (*model.ShowAuditBackRiskTemplateResponse, error) {
	requestDef := GenReqDefForShowAuditBackRiskTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuditBackRiskTemplateResponse), nil
	}
}

// ShowAuditBackRiskTemplateInvoker 获取单个风险导出配置
func (c *DbssClient) ShowAuditBackRiskTemplateInvoker(request *model.ShowAuditBackRiskTemplateRequest) *ShowAuditBackRiskTemplateInvoker {
	requestDef := GenReqDefForShowAuditBackRiskTemplate()
	return &ShowAuditBackRiskTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuditBackupStatus 获取备份状态
//
// 获取备份状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ShowAuditBackupStatus(request *model.ShowAuditBackupStatusRequest) (*model.ShowAuditBackupStatusResponse, error) {
	requestDef := GenReqDefForShowAuditBackupStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuditBackupStatusResponse), nil
	}
}

// ShowAuditBackupStatusInvoker 获取备份状态
func (c *DbssClient) ShowAuditBackupStatusInvoker(request *model.ShowAuditBackupStatusRequest) *ShowAuditBackupStatusInvoker {
	requestDef := GenReqDefForShowAuditBackupStatus()
	return &ShowAuditBackupStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowAuditQuota 查询账户配额信息[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowAuditQuotaInvoker 查询账户配额信息[待下线]
func (c *DbssClient) ShowAuditQuotaInvoker(request *model.ShowAuditQuotaRequest) *ShowAuditQuotaInvoker {
	requestDef := GenReqDefForShowAuditQuota()
	return &ShowAuditQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuditQuotaNew 查询账户配额信息
//
// 查询账户配额信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ShowAuditQuotaNew(request *model.ShowAuditQuotaNewRequest) (*model.ShowAuditQuotaNewResponse, error) {
	requestDef := GenReqDefForShowAuditQuotaNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuditQuotaNewResponse), nil
	}
}

// ShowAuditQuotaNewInvoker 查询账户配额信息
func (c *DbssClient) ShowAuditQuotaNewInvoker(request *model.ShowAuditQuotaNewRequest) *ShowAuditQuotaNewInvoker {
	requestDef := GenReqDefForShowAuditQuotaNew()
	return &ShowAuditQuotaNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowAuditRuleRisk 查询指定风险规则策略[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowAuditRuleRiskInvoker 查询指定风险规则策略[待下线]
func (c *DbssClient) ShowAuditRuleRiskInvoker(request *model.ShowAuditRuleRiskRequest) *ShowAuditRuleRiskInvoker {
	requestDef := GenReqDefForShowAuditRuleRisk()
	return &ShowAuditRuleRiskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuditRuleRiskNew 查询指定风险规则策略
//
// 查询指定风险规则策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ShowAuditRuleRiskNew(request *model.ShowAuditRuleRiskNewRequest) (*model.ShowAuditRuleRiskNewResponse, error) {
	requestDef := GenReqDefForShowAuditRuleRiskNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuditRuleRiskNewResponse), nil
	}
}

// ShowAuditRuleRiskNewInvoker 查询指定风险规则策略
func (c *DbssClient) ShowAuditRuleRiskNewInvoker(request *model.ShowAuditRuleRiskNewRequest) *ShowAuditRuleRiskNewInvoker {
	requestDef := GenReqDefForShowAuditRuleRiskNew()
	return &ShowAuditRuleRiskNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuditStatistics 获取租户下所有实例的风险汇总信息
//
// 获取租户下所有实例的风险汇总信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ShowAuditStatistics(request *model.ShowAuditStatisticsRequest) (*model.ShowAuditStatisticsResponse, error) {
	requestDef := GenReqDefForShowAuditStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuditStatisticsResponse), nil
	}
}

// ShowAuditStatisticsInvoker 获取租户下所有实例的风险汇总信息
func (c *DbssClient) ShowAuditStatisticsInvoker(request *model.ShowAuditStatisticsRequest) *ShowAuditStatisticsInvoker {
	requestDef := GenReqDefForShowAuditStatistics()
	return &ShowAuditStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuditTaskStatus 获取租户审计信息汇总任务状态
//
// 获取租户审计信息汇总任务状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ShowAuditTaskStatus(request *model.ShowAuditTaskStatusRequest) (*model.ShowAuditTaskStatusResponse, error) {
	requestDef := GenReqDefForShowAuditTaskStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuditTaskStatusResponse), nil
	}
}

// ShowAuditTaskStatusInvoker 获取租户审计信息汇总任务状态
func (c *DbssClient) ShowAuditTaskStatusInvoker(request *model.ShowAuditTaskStatusRequest) *ShowAuditTaskStatusInvoker {
	requestDef := GenReqDefForShowAuditTaskStatus()
	return &ShowAuditTaskStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuditTopicReportSchedulerConfig 获取报表的计划任务配置信息（topic方式）
//
// 获取报表的计划任务配置信息（topic方式）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ShowAuditTopicReportSchedulerConfig(request *model.ShowAuditTopicReportSchedulerConfigRequest) (*model.ShowAuditTopicReportSchedulerConfigResponse, error) {
	requestDef := GenReqDefForShowAuditTopicReportSchedulerConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuditTopicReportSchedulerConfigResponse), nil
	}
}

// ShowAuditTopicReportSchedulerConfigInvoker 获取报表的计划任务配置信息（topic方式）
func (c *DbssClient) ShowAuditTopicReportSchedulerConfigInvoker(request *model.ShowAuditTopicReportSchedulerConfigRequest) *ShowAuditTopicReportSchedulerConfigInvoker {
	requestDef := GenReqDefForShowAuditTopicReportSchedulerConfig()
	return &ShowAuditTopicReportSchedulerConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuditUpgradeStatus 查询租户的实例是否可升级
//
// 查询租户的实例是否可升级
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ShowAuditUpgradeStatus(request *model.ShowAuditUpgradeStatusRequest) (*model.ShowAuditUpgradeStatusResponse, error) {
	requestDef := GenReqDefForShowAuditUpgradeStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuditUpgradeStatusResponse), nil
	}
}

// ShowAuditUpgradeStatusInvoker 查询租户的实例是否可升级
func (c *DbssClient) ShowAuditUpgradeStatusInvoker(request *model.ShowAuditUpgradeStatusRequest) *ShowAuditUpgradeStatusInvoker {
	requestDef := GenReqDefForShowAuditUpgradeStatus()
	return &ShowAuditUpgradeStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBackupRiskBucketPath 获取风险导出桶名和路径
//
// 获取风险导出桶名和路径
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ShowBackupRiskBucketPath(request *model.ShowBackupRiskBucketPathRequest) (*model.ShowBackupRiskBucketPathResponse, error) {
	requestDef := GenReqDefForShowBackupRiskBucketPath()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackupRiskBucketPathResponse), nil
	}
}

// ShowBackupRiskBucketPathInvoker 获取风险导出桶名和路径
func (c *DbssClient) ShowBackupRiskBucketPathInvoker(request *model.ShowBackupRiskBucketPathRequest) *ShowBackupRiskBucketPathInvoker {
	requestDef := GenReqDefForShowBackupRiskBucketPath()
	return &ShowBackupRiskBucketPathInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceMonitorInfo 获取实例监控数据
//
// 获取实例监控数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ShowInstanceMonitorInfo(request *model.ShowInstanceMonitorInfoRequest) (*model.ShowInstanceMonitorInfoResponse, error) {
	requestDef := GenReqDefForShowInstanceMonitorInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceMonitorInfoResponse), nil
	}
}

// ShowInstanceMonitorInfoInvoker 获取实例监控数据
func (c *DbssClient) ShowInstanceMonitorInfoInvoker(request *model.ShowInstanceMonitorInfoRequest) *ShowInstanceMonitorInfoInvoker {
	requestDef := GenReqDefForShowInstanceMonitorInfo()
	return &ShowInstanceMonitorInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceQuota 查询加密/运维增强配额
//
// 查询加密/运维增强配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ShowInstanceQuota(request *model.ShowInstanceQuotaRequest) (*model.ShowInstanceQuotaResponse, error) {
	requestDef := GenReqDefForShowInstanceQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceQuotaResponse), nil
	}
}

// ShowInstanceQuotaInvoker 查询加密/运维增强配额
func (c *DbssClient) ShowInstanceQuotaInvoker(request *model.ShowInstanceQuotaRequest) *ShowInstanceQuotaInvoker {
	requestDef := GenReqDefForShowInstanceQuota()
	return &ShowInstanceQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSensitiveMaskSwitch 获取隐私数据脱敏开关
//
// 获取隐私数据脱敏开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ShowSensitiveMaskSwitch(request *model.ShowSensitiveMaskSwitchRequest) (*model.ShowSensitiveMaskSwitchResponse, error) {
	requestDef := GenReqDefForShowSensitiveMaskSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSensitiveMaskSwitchResponse), nil
	}
}

// ShowSensitiveMaskSwitchInvoker 获取隐私数据脱敏开关
func (c *DbssClient) ShowSensitiveMaskSwitchInvoker(request *model.ShowSensitiveMaskSwitchRequest) *ShowSensitiveMaskSwitchInvoker {
	requestDef := GenReqDefForShowSensitiveMaskSwitch()
	return &ShowSensitiveMaskSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSensitiveResultSwitch 获取隐私数据结果集开关
//
// 获取隐私数据结果集开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ShowSensitiveResultSwitch(request *model.ShowSensitiveResultSwitchRequest) (*model.ShowSensitiveResultSwitchResponse, error) {
	requestDef := GenReqDefForShowSensitiveResultSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSensitiveResultSwitchResponse), nil
	}
}

// ShowSensitiveResultSwitchInvoker 获取隐私数据结果集开关
func (c *DbssClient) ShowSensitiveResultSwitchInvoker(request *model.ShowSensitiveResultSwitchRequest) *ShowSensitiveResultSwitchInvoker {
	requestDef := GenReqDefForShowSensitiveResultSwitch()
	return &ShowSensitiveResultSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServerVersion 获取管理侧版本信息
//
// 获取管理侧版本信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ShowServerVersion(request *model.ShowServerVersionRequest) (*model.ShowServerVersionResponse, error) {
	requestDef := GenReqDefForShowServerVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerVersionResponse), nil
	}
}

// ShowServerVersionInvoker 获取管理侧版本信息
func (c *DbssClient) ShowServerVersionInvoker(request *model.ShowServerVersionRequest) *ShowServerVersionInvoker {
	requestDef := GenReqDefForShowServerVersion()
	return &ShowServerVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSqlDetail 获取指定SQL语句的详细信息
//
// 获取指定SQL语句的详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ShowSqlDetail(request *model.ShowSqlDetailRequest) (*model.ShowSqlDetailResponse, error) {
	requestDef := GenReqDefForShowSqlDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlDetailResponse), nil
	}
}

// ShowSqlDetailInvoker 获取指定SQL语句的详细信息
func (c *DbssClient) ShowSqlDetailInvoker(request *model.ShowSqlDetailRequest) *ShowSqlDetailInvoker {
	requestDef := GenReqDefForShowSqlDetail()
	return &ShowSqlDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// StartAuditInstance 开启审计实例[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// StartAuditInstanceInvoker 开启审计实例[待下线]
func (c *DbssClient) StartAuditInstanceInvoker(request *model.StartAuditInstanceRequest) *StartAuditInstanceInvoker {
	requestDef := GenReqDefForStartAuditInstance()
	return &StartAuditInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartAuditInstanceNew 开启审计实例
//
// 开启审计实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) StartAuditInstanceNew(request *model.StartAuditInstanceNewRequest) (*model.StartAuditInstanceNewResponse, error) {
	requestDef := GenReqDefForStartAuditInstanceNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartAuditInstanceNewResponse), nil
	}
}

// StartAuditInstanceNewInvoker 开启审计实例
func (c *DbssClient) StartAuditInstanceNewInvoker(request *model.StartAuditInstanceNewRequest) *StartAuditInstanceNewInvoker {
	requestDef := GenReqDefForStartAuditInstanceNew()
	return &StartAuditInstanceNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartDbEncryptInstance 启动数据库加密实例
//
// 启动数据库加密实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) StartDbEncryptInstance(request *model.StartDbEncryptInstanceRequest) (*model.StartDbEncryptInstanceResponse, error) {
	requestDef := GenReqDefForStartDbEncryptInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartDbEncryptInstanceResponse), nil
	}
}

// StartDbEncryptInstanceInvoker 启动数据库加密实例
func (c *DbssClient) StartDbEncryptInstanceInvoker(request *model.StartDbEncryptInstanceRequest) *StartDbEncryptInstanceInvoker {
	requestDef := GenReqDefForStartDbEncryptInstance()
	return &StartDbEncryptInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartDbOmInstance 启动数据库运维实例
//
// 启动数据库运维实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) StartDbOmInstance(request *model.StartDbOmInstanceRequest) (*model.StartDbOmInstanceResponse, error) {
	requestDef := GenReqDefForStartDbOmInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartDbOmInstanceResponse), nil
	}
}

// StartDbOmInstanceInvoker 启动数据库运维实例
func (c *DbssClient) StartDbOmInstanceInvoker(request *model.StartDbOmInstanceRequest) *StartDbOmInstanceInvoker {
	requestDef := GenReqDefForStartDbOmInstance()
	return &StartDbOmInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// StopAuditInstance 关闭审计实例[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// StopAuditInstanceInvoker 关闭审计实例[待下线]
func (c *DbssClient) StopAuditInstanceInvoker(request *model.StopAuditInstanceRequest) *StopAuditInstanceInvoker {
	requestDef := GenReqDefForStopAuditInstance()
	return &StopAuditInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopAuditInstanceNew 关闭审计实例
//
// 关闭审计实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) StopAuditInstanceNew(request *model.StopAuditInstanceNewRequest) (*model.StopAuditInstanceNewResponse, error) {
	requestDef := GenReqDefForStopAuditInstanceNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopAuditInstanceNewResponse), nil
	}
}

// StopAuditInstanceNewInvoker 关闭审计实例
func (c *DbssClient) StopAuditInstanceNewInvoker(request *model.StopAuditInstanceNewRequest) *StopAuditInstanceNewInvoker {
	requestDef := GenReqDefForStopAuditInstanceNew()
	return &StopAuditInstanceNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopDbEncryptInstance 停止数据库加密实例
//
// 停止数据库加密实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) StopDbEncryptInstance(request *model.StopDbEncryptInstanceRequest) (*model.StopDbEncryptInstanceResponse, error) {
	requestDef := GenReqDefForStopDbEncryptInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopDbEncryptInstanceResponse), nil
	}
}

// StopDbEncryptInstanceInvoker 停止数据库加密实例
func (c *DbssClient) StopDbEncryptInstanceInvoker(request *model.StopDbEncryptInstanceRequest) *StopDbEncryptInstanceInvoker {
	requestDef := GenReqDefForStopDbEncryptInstance()
	return &StopDbEncryptInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopDbOmInstance 停止数据库运维实例
//
// 停止数据库运维实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) StopDbOmInstance(request *model.StopDbOmInstanceRequest) (*model.StopDbOmInstanceResponse, error) {
	requestDef := GenReqDefForStopDbOmInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopDbOmInstanceResponse), nil
	}
}

// StopDbOmInstanceInvoker 停止数据库运维实例
func (c *DbssClient) StopDbOmInstanceInvoker(request *model.StopDbOmInstanceRequest) *StopDbOmInstanceInvoker {
	requestDef := GenReqDefForStopDbOmInstance()
	return &StopDbOmInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// SwitchAgent 开启关闭Agent[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// SwitchAgentInvoker 开启关闭Agent[待下线]
func (c *DbssClient) SwitchAgentInvoker(request *model.SwitchAgentRequest) *SwitchAgentInvoker {
	requestDef := GenReqDefForSwitchAgent()
	return &SwitchAgentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// SwitchAuditDatabase 开启关闭数据库[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// SwitchAuditDatabaseInvoker 开启关闭数据库[待下线]
func (c *DbssClient) SwitchAuditDatabaseInvoker(request *model.SwitchAuditDatabaseRequest) *SwitchAuditDatabaseInvoker {
	requestDef := GenReqDefForSwitchAuditDatabase()
	return &SwitchAuditDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchAuditDatabaseNew 开启关闭数据库
//
// 开启关闭数据库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SwitchAuditDatabaseNew(request *model.SwitchAuditDatabaseNewRequest) (*model.SwitchAuditDatabaseNewResponse, error) {
	requestDef := GenReqDefForSwitchAuditDatabaseNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchAuditDatabaseNewResponse), nil
	}
}

// SwitchAuditDatabaseNewInvoker 开启关闭数据库
func (c *DbssClient) SwitchAuditDatabaseNewInvoker(request *model.SwitchAuditDatabaseNewRequest) *SwitchAuditDatabaseNewInvoker {
	requestDef := GenReqDefForSwitchAuditDatabaseNew()
	return &SwitchAuditDatabaseNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// SwitchRiskRule 开启关闭风险规则[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// SwitchRiskRuleInvoker 开启关闭风险规则[待下线]
func (c *DbssClient) SwitchRiskRuleInvoker(request *model.SwitchRiskRuleRequest) *SwitchRiskRuleInvoker {
	requestDef := GenReqDefForSwitchRiskRule()
	return &SwitchRiskRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchRiskRuleNew 开启/关闭风险规则
//
// 开启/关闭风险规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SwitchRiskRuleNew(request *model.SwitchRiskRuleNewRequest) (*model.SwitchRiskRuleNewResponse, error) {
	requestDef := GenReqDefForSwitchRiskRuleNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchRiskRuleNewResponse), nil
	}
}

// SwitchRiskRuleNewInvoker 开启/关闭风险规则
func (c *DbssClient) SwitchRiskRuleNewInvoker(request *model.SwitchRiskRuleNewRequest) *SwitchRiskRuleNewInvoker {
	requestDef := GenReqDefForSwitchRiskRuleNew()
	return &SwitchRiskRuleNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnbindDbEncryptEip 解绑数据库加密实例的EIP
//
// 解绑数据库加密实例的EIP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) UnbindDbEncryptEip(request *model.UnbindDbEncryptEipRequest) (*model.UnbindDbEncryptEipResponse, error) {
	requestDef := GenReqDefForUnbindDbEncryptEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnbindDbEncryptEipResponse), nil
	}
}

// UnbindDbEncryptEipInvoker 解绑数据库加密实例的EIP
func (c *DbssClient) UnbindDbEncryptEipInvoker(request *model.UnbindDbEncryptEipRequest) *UnbindDbEncryptEipInvoker {
	requestDef := GenReqDefForUnbindDbEncryptEip()
	return &UnbindDbEncryptEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnbindDbOmEip 解绑数据库运维实例的EIP
//
// 解绑数据库运维实例的EIP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) UnbindDbOmEip(request *model.UnbindDbOmEipRequest) (*model.UnbindDbOmEipResponse, error) {
	requestDef := GenReqDefForUnbindDbOmEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnbindDbOmEipResponse), nil
	}
}

// UnbindDbOmEipInvoker 解绑数据库运维实例的EIP
func (c *DbssClient) UnbindDbOmEipInvoker(request *model.UnbindDbOmEipRequest) *UnbindDbOmEipInvoker {
	requestDef := GenReqDefForUnbindDbOmEip()
	return &UnbindDbOmEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UpdateAuditInstance 更新审计实例信息[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// UpdateAuditInstanceInvoker 更新审计实例信息[待下线]
func (c *DbssClient) UpdateAuditInstanceInvoker(request *model.UpdateAuditInstanceRequest) *UpdateAuditInstanceInvoker {
	requestDef := GenReqDefForUpdateAuditInstance()
	return &UpdateAuditInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAuditInstanceNew 更新审计实例信息
//
// 更新审计实例信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) UpdateAuditInstanceNew(request *model.UpdateAuditInstanceNewRequest) (*model.UpdateAuditInstanceNewResponse, error) {
	requestDef := GenReqDefForUpdateAuditInstanceNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAuditInstanceNewResponse), nil
	}
}

// UpdateAuditInstanceNewInvoker 更新审计实例信息
func (c *DbssClient) UpdateAuditInstanceNewInvoker(request *model.UpdateAuditInstanceNewRequest) *UpdateAuditInstanceNewInvoker {
	requestDef := GenReqDefForUpdateAuditInstanceNew()
	return &UpdateAuditInstanceNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAuditRiskBucketPath 修改风险导出桶名和路径(按DomainId)
//
// 修改风险导出桶名和路径(按DomainId)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) UpdateAuditRiskBucketPath(request *model.UpdateAuditRiskBucketPathRequest) (*model.UpdateAuditRiskBucketPathResponse, error) {
	requestDef := GenReqDefForUpdateAuditRiskBucketPath()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAuditRiskBucketPathResponse), nil
	}
}

// UpdateAuditRiskBucketPathInvoker 修改风险导出桶名和路径(按DomainId)
func (c *DbssClient) UpdateAuditRiskBucketPathInvoker(request *model.UpdateAuditRiskBucketPathRequest) *UpdateAuditRiskBucketPathInvoker {
	requestDef := GenReqDefForUpdateAuditRiskBucketPath()
	return &UpdateAuditRiskBucketPathInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAuditScopeRule 编辑审计范围规则
//
// 编辑审计范围规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) UpdateAuditScopeRule(request *model.UpdateAuditScopeRuleRequest) (*model.UpdateAuditScopeRuleResponse, error) {
	requestDef := GenReqDefForUpdateAuditScopeRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAuditScopeRuleResponse), nil
	}
}

// UpdateAuditScopeRuleInvoker 编辑审计范围规则
func (c *DbssClient) UpdateAuditScopeRuleInvoker(request *model.UpdateAuditScopeRuleRequest) *UpdateAuditScopeRuleInvoker {
	requestDef := GenReqDefForUpdateAuditScopeRule()
	return &UpdateAuditScopeRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UpdateAuditSecurityGroup 修改实例安全组[待下线]
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// UpdateAuditSecurityGroupInvoker 修改实例安全组[待下线]
func (c *DbssClient) UpdateAuditSecurityGroupInvoker(request *model.UpdateAuditSecurityGroupRequest) *UpdateAuditSecurityGroupInvoker {
	requestDef := GenReqDefForUpdateAuditSecurityGroup()
	return &UpdateAuditSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAuditSecurityGroupNew 修改实例安全组
//
// 修改实例安全组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) UpdateAuditSecurityGroupNew(request *model.UpdateAuditSecurityGroupNewRequest) (*model.UpdateAuditSecurityGroupNewResponse, error) {
	requestDef := GenReqDefForUpdateAuditSecurityGroupNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAuditSecurityGroupNewResponse), nil
	}
}

// UpdateAuditSecurityGroupNewInvoker 修改实例安全组
func (c *DbssClient) UpdateAuditSecurityGroupNewInvoker(request *model.UpdateAuditSecurityGroupNewRequest) *UpdateAuditSecurityGroupNewInvoker {
	requestDef := GenReqDefForUpdateAuditSecurityGroupNew()
	return &UpdateAuditSecurityGroupNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAuditTaskStatus 更新租户审计信息汇总任务状态
//
// 更新租户审计信息汇总任务状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) UpdateAuditTaskStatus(request *model.UpdateAuditTaskStatusRequest) (*model.UpdateAuditTaskStatusResponse, error) {
	requestDef := GenReqDefForUpdateAuditTaskStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAuditTaskStatusResponse), nil
	}
}

// UpdateAuditTaskStatusInvoker 更新租户审计信息汇总任务状态
func (c *DbssClient) UpdateAuditTaskStatusInvoker(request *model.UpdateAuditTaskStatusRequest) *UpdateAuditTaskStatusInvoker {
	requestDef := GenReqDefForUpdateAuditTaskStatus()
	return &UpdateAuditTaskStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAuditTopicReportSchedulerConfig 更改报表的计划任务配置信息（topic方式）
//
// 更改报表的计划任务配置信息（topic方式）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) UpdateAuditTopicReportSchedulerConfig(request *model.UpdateAuditTopicReportSchedulerConfigRequest) (*model.UpdateAuditTopicReportSchedulerConfigResponse, error) {
	requestDef := GenReqDefForUpdateAuditTopicReportSchedulerConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAuditTopicReportSchedulerConfigResponse), nil
	}
}

// UpdateAuditTopicReportSchedulerConfigInvoker 更改报表的计划任务配置信息（topic方式）
func (c *DbssClient) UpdateAuditTopicReportSchedulerConfigInvoker(request *model.UpdateAuditTopicReportSchedulerConfigRequest) *UpdateAuditTopicReportSchedulerConfigInvoker {
	requestDef := GenReqDefForUpdateAuditTopicReportSchedulerConfig()
	return &UpdateAuditTopicReportSchedulerConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDbEncryptInstanceName 更改数据库加密实例的名称
//
// 更改数据库加密实例的名称
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) UpdateDbEncryptInstanceName(request *model.UpdateDbEncryptInstanceNameRequest) (*model.UpdateDbEncryptInstanceNameResponse, error) {
	requestDef := GenReqDefForUpdateDbEncryptInstanceName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDbEncryptInstanceNameResponse), nil
	}
}

// UpdateDbEncryptInstanceNameInvoker 更改数据库加密实例的名称
func (c *DbssClient) UpdateDbEncryptInstanceNameInvoker(request *model.UpdateDbEncryptInstanceNameRequest) *UpdateDbEncryptInstanceNameInvoker {
	requestDef := GenReqDefForUpdateDbEncryptInstanceName()
	return &UpdateDbEncryptInstanceNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDbOmInstanceName 更改数据库运维实例的名称
//
// 更改数据库运维实例的名称
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) UpdateDbOmInstanceName(request *model.UpdateDbOmInstanceNameRequest) (*model.UpdateDbOmInstanceNameResponse, error) {
	requestDef := GenReqDefForUpdateDbOmInstanceName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDbOmInstanceNameResponse), nil
	}
}

// UpdateDbOmInstanceNameInvoker 更改数据库运维实例的名称
func (c *DbssClient) UpdateDbOmInstanceNameInvoker(request *model.UpdateDbOmInstanceNameRequest) *UpdateDbOmInstanceNameInvoker {
	requestDef := GenReqDefForUpdateDbOmInstanceName()
	return &UpdateDbOmInstanceNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSensitiveMaskRule 修改编辑隐私数据脱敏规则
//
// 修改编辑隐私数据脱敏规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) UpdateSensitiveMaskRule(request *model.UpdateSensitiveMaskRuleRequest) (*model.UpdateSensitiveMaskRuleResponse, error) {
	requestDef := GenReqDefForUpdateSensitiveMaskRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSensitiveMaskRuleResponse), nil
	}
}

// UpdateSensitiveMaskRuleInvoker 修改编辑隐私数据脱敏规则
func (c *DbssClient) UpdateSensitiveMaskRuleInvoker(request *model.UpdateSensitiveMaskRuleRequest) *UpdateSensitiveMaskRuleInvoker {
	requestDef := GenReqDefForUpdateSensitiveMaskRule()
	return &UpdateSensitiveMaskRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadAuditDbConfig 往OBS导出审计数据库配置
//
// 往OBS导出审计数据库配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) UploadAuditDbConfig(request *model.UploadAuditDbConfigRequest) (*model.UploadAuditDbConfigResponse, error) {
	requestDef := GenReqDefForUploadAuditDbConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadAuditDbConfigResponse), nil
	}
}

// UploadAuditDbConfigInvoker 往OBS导出审计数据库配置
func (c *DbssClient) UploadAuditDbConfigInvoker(request *model.UploadAuditDbConfigRequest) *UploadAuditDbConfigInvoker {
	requestDef := GenReqDefForUploadAuditDbConfig()
	return &UploadAuditDbConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddAuditAgentNew 添加审计数据库Agent
//
// 添加审计数据库Agent
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) AddAuditAgentNew(request *model.AddAuditAgentNewRequest) (*model.AddAuditAgentNewResponse, error) {
	requestDef := GenReqDefForAddAuditAgentNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAuditAgentNewResponse), nil
	}
}

// AddAuditAgentNewInvoker 添加审计数据库Agent
func (c *DbssClient) AddAuditAgentNewInvoker(request *model.AddAuditAgentNewRequest) *AddAuditAgentNewInvoker {
	requestDef := GenReqDefForAddAuditAgentNew()
	return &AddAuditAgentNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAuditDbAgent 指定agent_id方式添加agent
//
// 指定agent_id方式添加agent
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) CreateAuditDbAgent(request *model.CreateAuditDbAgentRequest) (*model.CreateAuditDbAgentResponse, error) {
	requestDef := GenReqDefForCreateAuditDbAgent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuditDbAgentResponse), nil
	}
}

// CreateAuditDbAgentInvoker 指定agent_id方式添加agent
func (c *DbssClient) CreateAuditDbAgentInvoker(request *model.CreateAuditDbAgentRequest) *CreateAuditDbAgentInvoker {
	requestDef := GenReqDefForCreateAuditDbAgent()
	return &CreateAuditDbAgentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAuditAgentNew 删除数据库Agent
//
// 删除数据库Agent
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) DeleteAuditAgentNew(request *model.DeleteAuditAgentNewRequest) (*model.DeleteAuditAgentNewResponse, error) {
	requestDef := GenReqDefForDeleteAuditAgentNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAuditAgentNewResponse), nil
	}
}

// DeleteAuditAgentNewInvoker 删除数据库Agent
func (c *DbssClient) DeleteAuditAgentNewInvoker(request *model.DeleteAuditAgentNewRequest) *DeleteAuditAgentNewInvoker {
	requestDef := GenReqDefForDeleteAuditAgentNew()
	return &DeleteAuditAgentNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadAuditAgentNew 下载审计Agent
//
// 下载审计Agent
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) DownloadAuditAgentNew(request *model.DownloadAuditAgentNewRequest) (*model.DownloadAuditAgentNewResponse, error) {
	requestDef := GenReqDefForDownloadAuditAgentNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadAuditAgentNewResponse), nil
	}
}

// DownloadAuditAgentNewInvoker 下载审计Agent
func (c *DbssClient) DownloadAuditAgentNewInvoker(request *model.DownloadAuditAgentNewRequest) *DownloadAuditAgentNewInvoker {
	requestDef := GenReqDefForDownloadAuditAgentNew()
	return &DownloadAuditAgentNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditAgentNew 查询数据库Agent列表
//
// 查询数据库Agent列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditAgentNew(request *model.ListAuditAgentNewRequest) (*model.ListAuditAgentNewResponse, error) {
	requestDef := GenReqDefForListAuditAgentNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditAgentNewResponse), nil
	}
}

// ListAuditAgentNewInvoker 查询数据库Agent列表
func (c *DbssClient) ListAuditAgentNewInvoker(request *model.ListAuditAgentNewRequest) *ListAuditAgentNewInvoker {
	requestDef := GenReqDefForListAuditAgentNew()
	return &ListAuditAgentNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchAgentNew 开启/关闭Agent
//
// 用于开启和关闭Agent审计的功能，当开启后，开始抓取用户的访问信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SwitchAgentNew(request *model.SwitchAgentNewRequest) (*model.SwitchAgentNewResponse, error) {
	requestDef := GenReqDefForSwitchAgentNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchAgentNewResponse), nil
	}
}

// SwitchAgentNewInvoker 开启/关闭Agent
func (c *DbssClient) SwitchAgentNewInvoker(request *model.SwitchAgentNewRequest) *SwitchAgentNewInvoker {
	requestDef := GenReqDefForSwitchAgentNew()
	return &SwitchAgentNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAddAuditWhitelist 批量添加白名单
//
// 批量添加白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) BatchAddAuditWhitelist(request *model.BatchAddAuditWhitelistRequest) (*model.BatchAddAuditWhitelistResponse, error) {
	requestDef := GenReqDefForBatchAddAuditWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddAuditWhitelistResponse), nil
	}
}

// BatchAddAuditWhitelistInvoker 批量添加白名单
func (c *DbssClient) BatchAddAuditWhitelistInvoker(request *model.BatchAddAuditWhitelistRequest) *BatchAddAuditWhitelistInvoker {
	requestDef := GenReqDefForBatchAddAuditWhitelist()
	return &BatchAddAuditWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAuditSqlRule 添加sql自定义规则
//
// 添加sql自定义规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) CreateAuditSqlRule(request *model.CreateAuditSqlRuleRequest) (*model.CreateAuditSqlRuleResponse, error) {
	requestDef := GenReqDefForCreateAuditSqlRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuditSqlRuleResponse), nil
	}
}

// CreateAuditSqlRuleInvoker 添加sql自定义规则
func (c *DbssClient) CreateAuditSqlRuleInvoker(request *model.CreateAuditSqlRuleRequest) *CreateAuditSqlRuleInvoker {
	requestDef := GenReqDefForCreateAuditSqlRule()
	return &CreateAuditSqlRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAuditRuleSql 删除sql自定义规则
//
// 删除sql自定义规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) DeleteAuditRuleSql(request *model.DeleteAuditRuleSqlRequest) (*model.DeleteAuditRuleSqlResponse, error) {
	requestDef := GenReqDefForDeleteAuditRuleSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAuditRuleSqlResponse), nil
	}
}

// DeleteAuditRuleSqlInvoker 删除sql自定义规则
func (c *DbssClient) DeleteAuditRuleSqlInvoker(request *model.DeleteAuditRuleSqlRequest) *DeleteAuditRuleSqlInvoker {
	requestDef := GenReqDefForDeleteAuditRuleSql()
	return &DeleteAuditRuleSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAuditWhitelist 批量删除白名单
//
// 批量删除白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) DeleteAuditWhitelist(request *model.DeleteAuditWhitelistRequest) (*model.DeleteAuditWhitelistResponse, error) {
	requestDef := GenReqDefForDeleteAuditWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAuditWhitelistResponse), nil
	}
}

// DeleteAuditWhitelistInvoker 批量删除白名单
func (c *DbssClient) DeleteAuditWhitelistInvoker(request *model.DeleteAuditWhitelistRequest) *DeleteAuditWhitelistInvoker {
	requestDef := GenReqDefForDeleteAuditWhitelist()
	return &DeleteAuditWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSqlInjectionRulesNew 查询SQL注入规则策略
//
// 查询SQL注入规则策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListSqlInjectionRulesNew(request *model.ListSqlInjectionRulesNewRequest) (*model.ListSqlInjectionRulesNewResponse, error) {
	requestDef := GenReqDefForListSqlInjectionRulesNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSqlInjectionRulesNewResponse), nil
	}
}

// ListSqlInjectionRulesNewInvoker 查询SQL注入规则策略
func (c *DbssClient) ListSqlInjectionRulesNewInvoker(request *model.ListSqlInjectionRulesNewRequest) *ListSqlInjectionRulesNewInvoker {
	requestDef := GenReqDefForListSqlInjectionRulesNew()
	return &ListSqlInjectionRulesNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWhitelists 查询白名单列表
//
// 查询白名单列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListWhitelists(request *model.ListWhitelistsRequest) (*model.ListWhitelistsResponse, error) {
	requestDef := GenReqDefForListWhitelists()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWhitelistsResponse), nil
	}
}

// ListWhitelistsInvoker 查询白名单列表
func (c *DbssClient) ListWhitelistsInvoker(request *model.ListWhitelistsRequest) *ListWhitelistsInvoker {
	requestDef := GenReqDefForListWhitelists()
	return &ListWhitelistsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetAuditSqlRuleSwitch 开启关闭sql注入策略
//
// 开启关闭sql注入策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SetAuditSqlRuleSwitch(request *model.SetAuditSqlRuleSwitchRequest) (*model.SetAuditSqlRuleSwitchResponse, error) {
	requestDef := GenReqDefForSetAuditSqlRuleSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetAuditSqlRuleSwitchResponse), nil
	}
}

// SetAuditSqlRuleSwitchInvoker 开启关闭sql注入策略
func (c *DbssClient) SetAuditSqlRuleSwitchInvoker(request *model.SetAuditSqlRuleSwitchRequest) *SetAuditSqlRuleSwitchInvoker {
	requestDef := GenReqDefForSetAuditSqlRuleSwitch()
	return &SetAuditSqlRuleSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetSqlRuleRank sql规则优先级排序
//
// sql规则优先级排序
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) SetSqlRuleRank(request *model.SetSqlRuleRankRequest) (*model.SetSqlRuleRankResponse, error) {
	requestDef := GenReqDefForSetSqlRuleRank()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetSqlRuleRankResponse), nil
	}
}

// SetSqlRuleRankInvoker sql规则优先级排序
func (c *DbssClient) SetSqlRuleRankInvoker(request *model.SetSqlRuleRankRequest) *SetSqlRuleRankInvoker {
	requestDef := GenReqDefForSetSqlRuleRank()
	return &SetSqlRuleRankInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAuditSqlRule 编辑sql自定义规则
//
// 编辑sql自定义规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) UpdateAuditSqlRule(request *model.UpdateAuditSqlRuleRequest) (*model.UpdateAuditSqlRuleResponse, error) {
	requestDef := GenReqDefForUpdateAuditSqlRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAuditSqlRuleResponse), nil
	}
}

// UpdateAuditSqlRuleInvoker 编辑sql自定义规则
func (c *DbssClient) UpdateAuditSqlRuleInvoker(request *model.UpdateAuditSqlRuleRequest) *UpdateAuditSqlRuleInvoker {
	requestDef := GenReqDefForUpdateAuditSqlRule()
	return &UpdateAuditSqlRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAuditWhitelist 修改白名单
//
// 修改白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) UpdateAuditWhitelist(request *model.UpdateAuditWhitelistRequest) (*model.UpdateAuditWhitelistResponse, error) {
	requestDef := GenReqDefForUpdateAuditWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAuditWhitelistResponse), nil
	}
}

// UpdateAuditWhitelistInvoker 修改白名单
func (c *DbssClient) UpdateAuditWhitelistInvoker(request *model.UpdateAuditWhitelistRequest) *UpdateAuditWhitelistInvoker {
	requestDef := GenReqDefForUpdateAuditWhitelist()
	return &UpdateAuditWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListAuditTags 获取实例标签
//
// 获取实例标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DbssClient) ListAuditTags(request *model.ListAuditTagsRequest) (*model.ListAuditTagsResponse, error) {
	requestDef := GenReqDefForListAuditTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditTagsResponse), nil
	}
}

// ListAuditTagsInvoker 获取实例标签
func (c *DbssClient) ListAuditTagsInvoker(request *model.ListAuditTagsRequest) *ListAuditTagsInvoker {
	requestDef := GenReqDefForListAuditTags()
	return &ListAuditTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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
