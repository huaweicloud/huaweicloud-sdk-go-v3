package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/edgesec/v2/model"
)

type EdgeSecClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewEdgeSecClient(hcClient *httpclient.HcHttpClient) *EdgeSecClient {
	return &EdgeSecClient{HcClient: hcClient}
}

func EdgeSecClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateDomain 创建防护域名
//
// 创建防护域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) CreateDomain(request *model.CreateDomainRequest) (*model.CreateDomainResponse, error) {
	requestDef := GenReqDefForCreateDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDomainResponse), nil
	}
}

// CreateDomainInvoker 创建防护域名
func (c *EdgeSecClient) CreateDomainInvoker(request *model.CreateDomainRequest) *CreateDomainInvoker {
	requestDef := GenReqDefForCreateDomain()
	return &CreateDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHttpReferenceTable 创建引用表
//
// 创建引用表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) CreateHttpReferenceTable(request *model.CreateHttpReferenceTableRequest) (*model.CreateHttpReferenceTableResponse, error) {
	requestDef := GenReqDefForCreateHttpReferenceTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHttpReferenceTableResponse), nil
	}
}

// CreateHttpReferenceTableInvoker 创建引用表
func (c *EdgeSecClient) CreateHttpReferenceTableInvoker(request *model.CreateHttpReferenceTableRequest) *CreateHttpReferenceTableInvoker {
	requestDef := GenReqDefForCreateHttpReferenceTable()
	return &CreateHttpReferenceTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDomain 删除防护域名
//
// 删除防护域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) DeleteDomain(request *model.DeleteDomainRequest) (*model.DeleteDomainResponse, error) {
	requestDef := GenReqDefForDeleteDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDomainResponse), nil
	}
}

// DeleteDomainInvoker 删除防护域名
func (c *EdgeSecClient) DeleteDomainInvoker(request *model.DeleteDomainRequest) *DeleteDomainInvoker {
	requestDef := GenReqDefForDeleteDomain()
	return &DeleteDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHttpReferenceTable 删除引用表
//
// 删除引用表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) DeleteHttpReferenceTable(request *model.DeleteHttpReferenceTableRequest) (*model.DeleteHttpReferenceTableResponse, error) {
	requestDef := GenReqDefForDeleteHttpReferenceTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHttpReferenceTableResponse), nil
	}
}

// DeleteHttpReferenceTableInvoker 删除引用表
func (c *EdgeSecClient) DeleteHttpReferenceTableInvoker(request *model.DeleteHttpReferenceTableRequest) *DeleteHttpReferenceTableInvoker {
	requestDef := GenReqDefForDeleteHttpReferenceTable()
	return &DeleteHttpReferenceTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDomainDetail 查询防护域名详情
//
// 查询防护域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowDomainDetail(request *model.ShowDomainDetailRequest) (*model.ShowDomainDetailResponse, error) {
	requestDef := GenReqDefForShowDomainDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainDetailResponse), nil
	}
}

// ShowDomainDetailInvoker 查询防护域名详情
func (c *EdgeSecClient) ShowDomainDetailInvoker(request *model.ShowDomainDetailRequest) *ShowDomainDetailInvoker {
	requestDef := GenReqDefForShowDomainDetail()
	return &ShowDomainDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDomains 查询防护域名列表
//
// 查询防护域名列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowDomains(request *model.ShowDomainsRequest) (*model.ShowDomainsResponse, error) {
	requestDef := GenReqDefForShowDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainsResponse), nil
	}
}

// ShowDomainsInvoker 查询防护域名列表
func (c *EdgeSecClient) ShowDomainsInvoker(request *model.ShowDomainsRequest) *ShowDomainsInvoker {
	requestDef := GenReqDefForShowDomains()
	return &ShowDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpOverviews 查询攻击域名
//
// 查询安全总览top信息，包含攻击域名、攻击源ip、攻击url、攻击来源区域、防护类型
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpOverviews(request *model.ShowHttpOverviewsRequest) (*model.ShowHttpOverviewsResponse, error) {
	requestDef := GenReqDefForShowHttpOverviews()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpOverviewsResponse), nil
	}
}

// ShowHttpOverviewsInvoker 查询攻击域名
func (c *EdgeSecClient) ShowHttpOverviewsInvoker(request *model.ShowHttpOverviewsRequest) *ShowHttpOverviewsInvoker {
	requestDef := GenReqDefForShowHttpOverviews()
	return &ShowHttpOverviewsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpReferenceTables 查询引用表列表
//
// 查询引用表列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpReferenceTables(request *model.ShowHttpReferenceTablesRequest) (*model.ShowHttpReferenceTablesResponse, error) {
	requestDef := GenReqDefForShowHttpReferenceTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpReferenceTablesResponse), nil
	}
}

// ShowHttpReferenceTablesInvoker 查询引用表列表
func (c *EdgeSecClient) ShowHttpReferenceTablesInvoker(request *model.ShowHttpReferenceTablesRequest) *ShowHttpReferenceTablesInvoker {
	requestDef := GenReqDefForShowHttpReferenceTables()
	return &ShowHttpReferenceTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpStatistics 查询安全总览请求数据
//
// 查询安全总览请求数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpStatistics(request *model.ShowHttpStatisticsRequest) (*model.ShowHttpStatisticsResponse, error) {
	requestDef := GenReqDefForShowHttpStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpStatisticsResponse), nil
	}
}

// ShowHttpStatisticsInvoker 查询安全总览请求数据
func (c *EdgeSecClient) ShowHttpStatisticsInvoker(request *model.ShowHttpStatisticsRequest) *ShowHttpStatisticsInvoker {
	requestDef := GenReqDefForShowHttpStatistics()
	return &ShowHttpStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDomain 更新防护域名
//
// 更新防护域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) UpdateDomain(request *model.UpdateDomainRequest) (*model.UpdateDomainResponse, error) {
	requestDef := GenReqDefForUpdateDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainResponse), nil
	}
}

// UpdateDomainInvoker 更新防护域名
func (c *EdgeSecClient) UpdateDomainInvoker(request *model.UpdateDomainRequest) *UpdateDomainInvoker {
	requestDef := GenReqDefForUpdateDomain()
	return &UpdateDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHttpReferenceTable 更新引用表
//
// 更新引用表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) UpdateHttpReferenceTable(request *model.UpdateHttpReferenceTableRequest) (*model.UpdateHttpReferenceTableResponse, error) {
	requestDef := GenReqDefForUpdateHttpReferenceTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHttpReferenceTableResponse), nil
	}
}

// UpdateHttpReferenceTableInvoker 更新引用表
func (c *EdgeSecClient) UpdateHttpReferenceTableInvoker(request *model.UpdateHttpReferenceTableRequest) *UpdateHttpReferenceTableInvoker {
	requestDef := GenReqDefForUpdateHttpReferenceTable()
	return &UpdateHttpReferenceTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadDdosAttackLogs Ddos攻击日志下载
//
// # Ddos攻击日志下载
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) DownloadDdosAttackLogs(request *model.DownloadDdosAttackLogsRequest) (*model.DownloadDdosAttackLogsResponse, error) {
	requestDef := GenReqDefForDownloadDdosAttackLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadDdosAttackLogsResponse), nil
	}
}

// DownloadDdosAttackLogsInvoker Ddos攻击日志下载
func (c *EdgeSecClient) DownloadDdosAttackLogsInvoker(request *model.DownloadDdosAttackLogsRequest) *DownloadDdosAttackLogsInvoker {
	requestDef := GenReqDefForDownloadDdosAttackLogs()
	return &DownloadDdosAttackLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDdosAttackLogs 查询ddos攻击日志列表
//
// 查询ddos攻击日志列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowDdosAttackLogs(request *model.ShowDdosAttackLogsRequest) (*model.ShowDdosAttackLogsResponse, error) {
	requestDef := GenReqDefForShowDdosAttackLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDdosAttackLogsResponse), nil
	}
}

// ShowDdosAttackLogsInvoker 查询ddos攻击日志列表
func (c *EdgeSecClient) ShowDdosAttackLogsInvoker(request *model.ShowDdosAttackLogsRequest) *ShowDdosAttackLogsInvoker {
	requestDef := GenReqDefForShowDdosAttackLogs()
	return &ShowDdosAttackLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDdosAttackTimelineStats 查询DDoS攻击统计时间线数据
//
// 查询DDoS攻击统计时间线数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowDdosAttackTimelineStats(request *model.ShowDdosAttackTimelineStatsRequest) (*model.ShowDdosAttackTimelineStatsResponse, error) {
	requestDef := GenReqDefForShowDdosAttackTimelineStats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDdosAttackTimelineStatsResponse), nil
	}
}

// ShowDdosAttackTimelineStatsInvoker 查询DDoS攻击统计时间线数据
func (c *EdgeSecClient) ShowDdosAttackTimelineStatsInvoker(request *model.ShowDdosAttackTimelineStatsRequest) *ShowDdosAttackTimelineStatsInvoker {
	requestDef := GenReqDefForShowDdosAttackTimelineStats()
	return &ShowDdosAttackTimelineStatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ApplyHttpPolicy 更新防护策略的域名
//
// 更新防护策略的域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ApplyHttpPolicy(request *model.ApplyHttpPolicyRequest) (*model.ApplyHttpPolicyResponse, error) {
	requestDef := GenReqDefForApplyHttpPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApplyHttpPolicyResponse), nil
	}
}

// ApplyHttpPolicyInvoker 更新防护策略的域名
func (c *EdgeSecClient) ApplyHttpPolicyInvoker(request *model.ApplyHttpPolicyRequest) *ApplyHttpPolicyInvoker {
	requestDef := GenReqDefForApplyHttpPolicy()
	return &ApplyHttpPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHttpAccessControlRule 创建精准防护规则
//
// 创建精准防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) CreateHttpAccessControlRule(request *model.CreateHttpAccessControlRuleRequest) (*model.CreateHttpAccessControlRuleResponse, error) {
	requestDef := GenReqDefForCreateHttpAccessControlRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHttpAccessControlRuleResponse), nil
	}
}

// CreateHttpAccessControlRuleInvoker 创建精准防护规则
func (c *EdgeSecClient) CreateHttpAccessControlRuleInvoker(request *model.CreateHttpAccessControlRuleRequest) *CreateHttpAccessControlRuleInvoker {
	requestDef := GenReqDefForCreateHttpAccessControlRule()
	return &CreateHttpAccessControlRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHttpGeoIpRule 创建地理位置规则
//
// 创建地理位置规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) CreateHttpGeoIpRule(request *model.CreateHttpGeoIpRuleRequest) (*model.CreateHttpGeoIpRuleResponse, error) {
	requestDef := GenReqDefForCreateHttpGeoIpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHttpGeoIpRuleResponse), nil
	}
}

// CreateHttpGeoIpRuleInvoker 创建地理位置规则
func (c *EdgeSecClient) CreateHttpGeoIpRuleInvoker(request *model.CreateHttpGeoIpRuleRequest) *CreateHttpGeoIpRuleInvoker {
	requestDef := GenReqDefForCreateHttpGeoIpRule()
	return &CreateHttpGeoIpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHttpIgnoreRule 创建误报屏蔽规则
//
// 创建误报屏蔽规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) CreateHttpIgnoreRule(request *model.CreateHttpIgnoreRuleRequest) (*model.CreateHttpIgnoreRuleResponse, error) {
	requestDef := GenReqDefForCreateHttpIgnoreRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHttpIgnoreRuleResponse), nil
	}
}

// CreateHttpIgnoreRuleInvoker 创建误报屏蔽规则
func (c *EdgeSecClient) CreateHttpIgnoreRuleInvoker(request *model.CreateHttpIgnoreRuleRequest) *CreateHttpIgnoreRuleInvoker {
	requestDef := GenReqDefForCreateHttpIgnoreRule()
	return &CreateHttpIgnoreRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHttpPolicy 创建防护策略
//
// 创建防护策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) CreateHttpPolicy(request *model.CreateHttpPolicyRequest) (*model.CreateHttpPolicyResponse, error) {
	requestDef := GenReqDefForCreateHttpPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHttpPolicyResponse), nil
	}
}

// CreateHttpPolicyInvoker 创建防护策略
func (c *EdgeSecClient) CreateHttpPolicyInvoker(request *model.CreateHttpPolicyRequest) *CreateHttpPolicyInvoker {
	requestDef := GenReqDefForCreateHttpPolicy()
	return &CreateHttpPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHttpPunishmentRule 创建攻击惩罚规则
//
// 创建攻击惩罚规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) CreateHttpPunishmentRule(request *model.CreateHttpPunishmentRuleRequest) (*model.CreateHttpPunishmentRuleResponse, error) {
	requestDef := GenReqDefForCreateHttpPunishmentRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHttpPunishmentRuleResponse), nil
	}
}

// CreateHttpPunishmentRuleInvoker 创建攻击惩罚规则
func (c *EdgeSecClient) CreateHttpPunishmentRuleInvoker(request *model.CreateHttpPunishmentRuleRequest) *CreateHttpPunishmentRuleInvoker {
	requestDef := GenReqDefForCreateHttpPunishmentRule()
	return &CreateHttpPunishmentRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHttpAccessControlRule 删除精准防护规则
//
// 删除精准防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) DeleteHttpAccessControlRule(request *model.DeleteHttpAccessControlRuleRequest) (*model.DeleteHttpAccessControlRuleResponse, error) {
	requestDef := GenReqDefForDeleteHttpAccessControlRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHttpAccessControlRuleResponse), nil
	}
}

// DeleteHttpAccessControlRuleInvoker 删除精准防护规则
func (c *EdgeSecClient) DeleteHttpAccessControlRuleInvoker(request *model.DeleteHttpAccessControlRuleRequest) *DeleteHttpAccessControlRuleInvoker {
	requestDef := GenReqDefForDeleteHttpAccessControlRule()
	return &DeleteHttpAccessControlRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHttpGeoIpRule 删除地理位置规则
//
// 删除地理位置规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) DeleteHttpGeoIpRule(request *model.DeleteHttpGeoIpRuleRequest) (*model.DeleteHttpGeoIpRuleResponse, error) {
	requestDef := GenReqDefForDeleteHttpGeoIpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHttpGeoIpRuleResponse), nil
	}
}

// DeleteHttpGeoIpRuleInvoker 删除地理位置规则
func (c *EdgeSecClient) DeleteHttpGeoIpRuleInvoker(request *model.DeleteHttpGeoIpRuleRequest) *DeleteHttpGeoIpRuleInvoker {
	requestDef := GenReqDefForDeleteHttpGeoIpRule()
	return &DeleteHttpGeoIpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHttpIgnoreRule 删除误报屏蔽规则
//
// 删除误报屏蔽规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) DeleteHttpIgnoreRule(request *model.DeleteHttpIgnoreRuleRequest) (*model.DeleteHttpIgnoreRuleResponse, error) {
	requestDef := GenReqDefForDeleteHttpIgnoreRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHttpIgnoreRuleResponse), nil
	}
}

// DeleteHttpIgnoreRuleInvoker 删除误报屏蔽规则
func (c *EdgeSecClient) DeleteHttpIgnoreRuleInvoker(request *model.DeleteHttpIgnoreRuleRequest) *DeleteHttpIgnoreRuleInvoker {
	requestDef := GenReqDefForDeleteHttpIgnoreRule()
	return &DeleteHttpIgnoreRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHttpPolicy 删除防护策略
//
// 删除防护策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) DeleteHttpPolicy(request *model.DeleteHttpPolicyRequest) (*model.DeleteHttpPolicyResponse, error) {
	requestDef := GenReqDefForDeleteHttpPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHttpPolicyResponse), nil
	}
}

// DeleteHttpPolicyInvoker 删除防护策略
func (c *EdgeSecClient) DeleteHttpPolicyInvoker(request *model.DeleteHttpPolicyRequest) *DeleteHttpPolicyInvoker {
	requestDef := GenReqDefForDeleteHttpPolicy()
	return &DeleteHttpPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHttpPunishmentRule 删除攻击惩罚规则
//
// 删除攻击惩罚规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) DeleteHttpPunishmentRule(request *model.DeleteHttpPunishmentRuleRequest) (*model.DeleteHttpPunishmentRuleResponse, error) {
	requestDef := GenReqDefForDeleteHttpPunishmentRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHttpPunishmentRuleResponse), nil
	}
}

// DeleteHttpPunishmentRuleInvoker 删除攻击惩罚规则
func (c *EdgeSecClient) DeleteHttpPunishmentRuleInvoker(request *model.DeleteHttpPunishmentRuleRequest) *DeleteHttpPunishmentRuleInvoker {
	requestDef := GenReqDefForDeleteHttpPunishmentRule()
	return &DeleteHttpPunishmentRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetHttpIgnoreRule 重置误报屏蔽规则
//
// 重置误报屏蔽规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ResetHttpIgnoreRule(request *model.ResetHttpIgnoreRuleRequest) (*model.ResetHttpIgnoreRuleResponse, error) {
	requestDef := GenReqDefForResetHttpIgnoreRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetHttpIgnoreRuleResponse), nil
	}
}

// ResetHttpIgnoreRuleInvoker 重置误报屏蔽规则
func (c *EdgeSecClient) ResetHttpIgnoreRuleInvoker(request *model.ResetHttpIgnoreRuleRequest) *ResetHttpIgnoreRuleInvoker {
	requestDef := GenReqDefForResetHttpIgnoreRule()
	return &ResetHttpIgnoreRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpAccessControlRule 查询精准防护规则
//
// 查询精准防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpAccessControlRule(request *model.ShowHttpAccessControlRuleRequest) (*model.ShowHttpAccessControlRuleResponse, error) {
	requestDef := GenReqDefForShowHttpAccessControlRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpAccessControlRuleResponse), nil
	}
}

// ShowHttpAccessControlRuleInvoker 查询精准防护规则
func (c *EdgeSecClient) ShowHttpAccessControlRuleInvoker(request *model.ShowHttpAccessControlRuleRequest) *ShowHttpAccessControlRuleInvoker {
	requestDef := GenReqDefForShowHttpAccessControlRule()
	return &ShowHttpAccessControlRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpAccessControlRules 查询精准防护规则列表
//
// 查询精准防护规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpAccessControlRules(request *model.ShowHttpAccessControlRulesRequest) (*model.ShowHttpAccessControlRulesResponse, error) {
	requestDef := GenReqDefForShowHttpAccessControlRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpAccessControlRulesResponse), nil
	}
}

// ShowHttpAccessControlRulesInvoker 查询精准防护规则列表
func (c *EdgeSecClient) ShowHttpAccessControlRulesInvoker(request *model.ShowHttpAccessControlRulesRequest) *ShowHttpAccessControlRulesInvoker {
	requestDef := GenReqDefForShowHttpAccessControlRules()
	return &ShowHttpAccessControlRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpAttackDistributionStats 查询HTTP攻击分布数据
//
// 查询HTTP攻击分布数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpAttackDistributionStats(request *model.ShowHttpAttackDistributionStatsRequest) (*model.ShowHttpAttackDistributionStatsResponse, error) {
	requestDef := GenReqDefForShowHttpAttackDistributionStats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpAttackDistributionStatsResponse), nil
	}
}

// ShowHttpAttackDistributionStatsInvoker 查询HTTP攻击分布数据
func (c *EdgeSecClient) ShowHttpAttackDistributionStatsInvoker(request *model.ShowHttpAttackDistributionStatsRequest) *ShowHttpAttackDistributionStatsInvoker {
	requestDef := GenReqDefForShowHttpAttackDistributionStats()
	return &ShowHttpAttackDistributionStatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpAttackTimelineStats 查询HTTP攻击统计时间线数据
//
// 查询HTTP攻击统计时间线数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpAttackTimelineStats(request *model.ShowHttpAttackTimelineStatsRequest) (*model.ShowHttpAttackTimelineStatsResponse, error) {
	requestDef := GenReqDefForShowHttpAttackTimelineStats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpAttackTimelineStatsResponse), nil
	}
}

// ShowHttpAttackTimelineStatsInvoker 查询HTTP攻击统计时间线数据
func (c *EdgeSecClient) ShowHttpAttackTimelineStatsInvoker(request *model.ShowHttpAttackTimelineStatsRequest) *ShowHttpAttackTimelineStatsInvoker {
	requestDef := GenReqDefForShowHttpAttackTimelineStats()
	return &ShowHttpAttackTimelineStatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpAttackTopStats 查询HTTP攻击TOP数据
//
// 查询HTTP攻击TOP数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpAttackTopStats(request *model.ShowHttpAttackTopStatsRequest) (*model.ShowHttpAttackTopStatsResponse, error) {
	requestDef := GenReqDefForShowHttpAttackTopStats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpAttackTopStatsResponse), nil
	}
}

// ShowHttpAttackTopStatsInvoker 查询HTTP攻击TOP数据
func (c *EdgeSecClient) ShowHttpAttackTopStatsInvoker(request *model.ShowHttpAttackTopStatsRequest) *ShowHttpAttackTopStatsInvoker {
	requestDef := GenReqDefForShowHttpAttackTopStats()
	return &ShowHttpAttackTopStatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpGeoIpRule 查询地理位置规则
//
// 查询地理位置规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpGeoIpRule(request *model.ShowHttpGeoIpRuleRequest) (*model.ShowHttpGeoIpRuleResponse, error) {
	requestDef := GenReqDefForShowHttpGeoIpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpGeoIpRuleResponse), nil
	}
}

// ShowHttpGeoIpRuleInvoker 查询地理位置规则
func (c *EdgeSecClient) ShowHttpGeoIpRuleInvoker(request *model.ShowHttpGeoIpRuleRequest) *ShowHttpGeoIpRuleInvoker {
	requestDef := GenReqDefForShowHttpGeoIpRule()
	return &ShowHttpGeoIpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpGeoIpRules 查询地理位置规则列表
//
// 查询地理位置规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpGeoIpRules(request *model.ShowHttpGeoIpRulesRequest) (*model.ShowHttpGeoIpRulesResponse, error) {
	requestDef := GenReqDefForShowHttpGeoIpRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpGeoIpRulesResponse), nil
	}
}

// ShowHttpGeoIpRulesInvoker 查询地理位置规则列表
func (c *EdgeSecClient) ShowHttpGeoIpRulesInvoker(request *model.ShowHttpGeoIpRulesRequest) *ShowHttpGeoIpRulesInvoker {
	requestDef := GenReqDefForShowHttpGeoIpRules()
	return &ShowHttpGeoIpRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpIgnoreRule 查询误报屏蔽规则
//
// 查询误报屏蔽规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpIgnoreRule(request *model.ShowHttpIgnoreRuleRequest) (*model.ShowHttpIgnoreRuleResponse, error) {
	requestDef := GenReqDefForShowHttpIgnoreRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpIgnoreRuleResponse), nil
	}
}

// ShowHttpIgnoreRuleInvoker 查询误报屏蔽规则
func (c *EdgeSecClient) ShowHttpIgnoreRuleInvoker(request *model.ShowHttpIgnoreRuleRequest) *ShowHttpIgnoreRuleInvoker {
	requestDef := GenReqDefForShowHttpIgnoreRule()
	return &ShowHttpIgnoreRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpIgnoreRules 查询误报屏蔽规则列表
//
// 查询误报屏蔽规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpIgnoreRules(request *model.ShowHttpIgnoreRulesRequest) (*model.ShowHttpIgnoreRulesResponse, error) {
	requestDef := GenReqDefForShowHttpIgnoreRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpIgnoreRulesResponse), nil
	}
}

// ShowHttpIgnoreRulesInvoker 查询误报屏蔽规则列表
func (c *EdgeSecClient) ShowHttpIgnoreRulesInvoker(request *model.ShowHttpIgnoreRulesRequest) *ShowHttpIgnoreRulesInvoker {
	requestDef := GenReqDefForShowHttpIgnoreRules()
	return &ShowHttpIgnoreRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpPolicies 查询防护策略列表
//
// 查询防护策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpPolicies(request *model.ShowHttpPoliciesRequest) (*model.ShowHttpPoliciesResponse, error) {
	requestDef := GenReqDefForShowHttpPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpPoliciesResponse), nil
	}
}

// ShowHttpPoliciesInvoker 查询防护策略列表
func (c *EdgeSecClient) ShowHttpPoliciesInvoker(request *model.ShowHttpPoliciesRequest) *ShowHttpPoliciesInvoker {
	requestDef := GenReqDefForShowHttpPolicies()
	return &ShowHttpPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpPolicy 查询防护策略
//
// 查询防护策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpPolicy(request *model.ShowHttpPolicyRequest) (*model.ShowHttpPolicyResponse, error) {
	requestDef := GenReqDefForShowHttpPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpPolicyResponse), nil
	}
}

// ShowHttpPolicyInvoker 查询防护策略
func (c *EdgeSecClient) ShowHttpPolicyInvoker(request *model.ShowHttpPolicyRequest) *ShowHttpPolicyInvoker {
	requestDef := GenReqDefForShowHttpPolicy()
	return &ShowHttpPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpPunishmentRule 查询攻击惩罚规则
//
// 查询攻击惩罚规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpPunishmentRule(request *model.ShowHttpPunishmentRuleRequest) (*model.ShowHttpPunishmentRuleResponse, error) {
	requestDef := GenReqDefForShowHttpPunishmentRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpPunishmentRuleResponse), nil
	}
}

// ShowHttpPunishmentRuleInvoker 查询攻击惩罚规则
func (c *EdgeSecClient) ShowHttpPunishmentRuleInvoker(request *model.ShowHttpPunishmentRuleRequest) *ShowHttpPunishmentRuleInvoker {
	requestDef := GenReqDefForShowHttpPunishmentRule()
	return &ShowHttpPunishmentRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpPunishmentRules 查询攻击惩罚规则列表
//
// 查询攻击惩罚规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpPunishmentRules(request *model.ShowHttpPunishmentRulesRequest) (*model.ShowHttpPunishmentRulesResponse, error) {
	requestDef := GenReqDefForShowHttpPunishmentRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpPunishmentRulesResponse), nil
	}
}

// ShowHttpPunishmentRulesInvoker 查询攻击惩罚规则列表
func (c *EdgeSecClient) ShowHttpPunishmentRulesInvoker(request *model.ShowHttpPunishmentRulesRequest) *ShowHttpPunishmentRulesInvoker {
	requestDef := GenReqDefForShowHttpPunishmentRules()
	return &ShowHttpPunishmentRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHttpAccessControlRule 更新精准防护规则
//
// 更新精准防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) UpdateHttpAccessControlRule(request *model.UpdateHttpAccessControlRuleRequest) (*model.UpdateHttpAccessControlRuleResponse, error) {
	requestDef := GenReqDefForUpdateHttpAccessControlRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHttpAccessControlRuleResponse), nil
	}
}

// UpdateHttpAccessControlRuleInvoker 更新精准防护规则
func (c *EdgeSecClient) UpdateHttpAccessControlRuleInvoker(request *model.UpdateHttpAccessControlRuleRequest) *UpdateHttpAccessControlRuleInvoker {
	requestDef := GenReqDefForUpdateHttpAccessControlRule()
	return &UpdateHttpAccessControlRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHttpGeoIpRule 更新地理位置规则
//
// 更新地理位置规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) UpdateHttpGeoIpRule(request *model.UpdateHttpGeoIpRuleRequest) (*model.UpdateHttpGeoIpRuleResponse, error) {
	requestDef := GenReqDefForUpdateHttpGeoIpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHttpGeoIpRuleResponse), nil
	}
}

// UpdateHttpGeoIpRuleInvoker 更新地理位置规则
func (c *EdgeSecClient) UpdateHttpGeoIpRuleInvoker(request *model.UpdateHttpGeoIpRuleRequest) *UpdateHttpGeoIpRuleInvoker {
	requestDef := GenReqDefForUpdateHttpGeoIpRule()
	return &UpdateHttpGeoIpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHttpIgnoreRule 更新误报屏蔽规则
//
// 更新误报屏蔽规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) UpdateHttpIgnoreRule(request *model.UpdateHttpIgnoreRuleRequest) (*model.UpdateHttpIgnoreRuleResponse, error) {
	requestDef := GenReqDefForUpdateHttpIgnoreRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHttpIgnoreRuleResponse), nil
	}
}

// UpdateHttpIgnoreRuleInvoker 更新误报屏蔽规则
func (c *EdgeSecClient) UpdateHttpIgnoreRuleInvoker(request *model.UpdateHttpIgnoreRuleRequest) *UpdateHttpIgnoreRuleInvoker {
	requestDef := GenReqDefForUpdateHttpIgnoreRule()
	return &UpdateHttpIgnoreRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHttpPolicy 更新防护策略
//
// 更新防护策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) UpdateHttpPolicy(request *model.UpdateHttpPolicyRequest) (*model.UpdateHttpPolicyResponse, error) {
	requestDef := GenReqDefForUpdateHttpPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHttpPolicyResponse), nil
	}
}

// UpdateHttpPolicyInvoker 更新防护策略
func (c *EdgeSecClient) UpdateHttpPolicyInvoker(request *model.UpdateHttpPolicyRequest) *UpdateHttpPolicyInvoker {
	requestDef := GenReqDefForUpdateHttpPolicy()
	return &UpdateHttpPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHttpPolicyRuleStatus 更新防护策略规则开关
//
// 更新防护策略规则开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) UpdateHttpPolicyRuleStatus(request *model.UpdateHttpPolicyRuleStatusRequest) (*model.UpdateHttpPolicyRuleStatusResponse, error) {
	requestDef := GenReqDefForUpdateHttpPolicyRuleStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHttpPolicyRuleStatusResponse), nil
	}
}

// UpdateHttpPolicyRuleStatusInvoker 更新防护策略规则开关
func (c *EdgeSecClient) UpdateHttpPolicyRuleStatusInvoker(request *model.UpdateHttpPolicyRuleStatusRequest) *UpdateHttpPolicyRuleStatusInvoker {
	requestDef := GenReqDefForUpdateHttpPolicyRuleStatus()
	return &UpdateHttpPolicyRuleStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHttpPunishmentRule 更新攻击惩罚规则
//
// 更新攻击惩罚规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) UpdateHttpPunishmentRule(request *model.UpdateHttpPunishmentRuleRequest) (*model.UpdateHttpPunishmentRuleResponse, error) {
	requestDef := GenReqDefForUpdateHttpPunishmentRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHttpPunishmentRuleResponse), nil
	}
}

// UpdateHttpPunishmentRuleInvoker 更新攻击惩罚规则
func (c *EdgeSecClient) UpdateHttpPunishmentRuleInvoker(request *model.UpdateHttpPunishmentRuleRequest) *UpdateHttpPunishmentRuleInvoker {
	requestDef := GenReqDefForUpdateHttpPunishmentRule()
	return &UpdateHttpPunishmentRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHttpCcRule 创建cc规则
//
// 创建cc规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) CreateHttpCcRule(request *model.CreateHttpCcRuleRequest) (*model.CreateHttpCcRuleResponse, error) {
	requestDef := GenReqDefForCreateHttpCcRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHttpCcRuleResponse), nil
	}
}

// CreateHttpCcRuleInvoker 创建cc规则
func (c *EdgeSecClient) CreateHttpCcRuleInvoker(request *model.CreateHttpCcRuleRequest) *CreateHttpCcRuleInvoker {
	requestDef := GenReqDefForCreateHttpCcRule()
	return &CreateHttpCcRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHttpCcRule 删除cc规则
//
// 删除cc规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) DeleteHttpCcRule(request *model.DeleteHttpCcRuleRequest) (*model.DeleteHttpCcRuleResponse, error) {
	requestDef := GenReqDefForDeleteHttpCcRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHttpCcRuleResponse), nil
	}
}

// DeleteHttpCcRuleInvoker 删除cc规则
func (c *EdgeSecClient) DeleteHttpCcRuleInvoker(request *model.DeleteHttpCcRuleRequest) *DeleteHttpCcRuleInvoker {
	requestDef := GenReqDefForDeleteHttpCcRule()
	return &DeleteHttpCcRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpCcRule 查询cc规则
//
// 查询cc规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpCcRule(request *model.ShowHttpCcRuleRequest) (*model.ShowHttpCcRuleResponse, error) {
	requestDef := GenReqDefForShowHttpCcRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpCcRuleResponse), nil
	}
}

// ShowHttpCcRuleInvoker 查询cc规则
func (c *EdgeSecClient) ShowHttpCcRuleInvoker(request *model.ShowHttpCcRuleRequest) *ShowHttpCcRuleInvoker {
	requestDef := GenReqDefForShowHttpCcRule()
	return &ShowHttpCcRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpCcRules 查询cc规则列表
//
// 查询cc规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpCcRules(request *model.ShowHttpCcRulesRequest) (*model.ShowHttpCcRulesResponse, error) {
	requestDef := GenReqDefForShowHttpCcRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpCcRulesResponse), nil
	}
}

// ShowHttpCcRulesInvoker 查询cc规则列表
func (c *EdgeSecClient) ShowHttpCcRulesInvoker(request *model.ShowHttpCcRulesRequest) *ShowHttpCcRulesInvoker {
	requestDef := GenReqDefForShowHttpCcRules()
	return &ShowHttpCcRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHttpCcRule 更新cc规则
//
// 更新cc规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) UpdateHttpCcRule(request *model.UpdateHttpCcRuleRequest) (*model.UpdateHttpCcRuleResponse, error) {
	requestDef := GenReqDefForUpdateHttpCcRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHttpCcRuleResponse), nil
	}
}

// UpdateHttpCcRuleInvoker 更新cc规则
func (c *EdgeSecClient) UpdateHttpCcRuleInvoker(request *model.UpdateHttpCcRuleRequest) *UpdateHttpCcRuleInvoker {
	requestDef := GenReqDefForUpdateHttpCcRule()
	return &UpdateHttpCcRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHttpBlockTrustIpRule 创建IP黑白名单规则
//
// 创建IP黑白名单规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) CreateHttpBlockTrustIpRule(request *model.CreateHttpBlockTrustIpRuleRequest) (*model.CreateHttpBlockTrustIpRuleResponse, error) {
	requestDef := GenReqDefForCreateHttpBlockTrustIpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHttpBlockTrustIpRuleResponse), nil
	}
}

// CreateHttpBlockTrustIpRuleInvoker 创建IP黑白名单规则
func (c *EdgeSecClient) CreateHttpBlockTrustIpRuleInvoker(request *model.CreateHttpBlockTrustIpRuleRequest) *CreateHttpBlockTrustIpRuleInvoker {
	requestDef := GenReqDefForCreateHttpBlockTrustIpRule()
	return &CreateHttpBlockTrustIpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHttpBlockTrustIpRule 删除IP黑白名单规则
//
// 删除IP黑白名单规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) DeleteHttpBlockTrustIpRule(request *model.DeleteHttpBlockTrustIpRuleRequest) (*model.DeleteHttpBlockTrustIpRuleResponse, error) {
	requestDef := GenReqDefForDeleteHttpBlockTrustIpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHttpBlockTrustIpRuleResponse), nil
	}
}

// DeleteHttpBlockTrustIpRuleInvoker 删除IP黑白名单规则
func (c *EdgeSecClient) DeleteHttpBlockTrustIpRuleInvoker(request *model.DeleteHttpBlockTrustIpRuleRequest) *DeleteHttpBlockTrustIpRuleInvoker {
	requestDef := GenReqDefForDeleteHttpBlockTrustIpRule()
	return &DeleteHttpBlockTrustIpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpBlockTrustIpRule 查询IP黑白名单规则
//
// 查询IP黑白名单规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpBlockTrustIpRule(request *model.ShowHttpBlockTrustIpRuleRequest) (*model.ShowHttpBlockTrustIpRuleResponse, error) {
	requestDef := GenReqDefForShowHttpBlockTrustIpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpBlockTrustIpRuleResponse), nil
	}
}

// ShowHttpBlockTrustIpRuleInvoker 查询IP黑白名单规则
func (c *EdgeSecClient) ShowHttpBlockTrustIpRuleInvoker(request *model.ShowHttpBlockTrustIpRuleRequest) *ShowHttpBlockTrustIpRuleInvoker {
	requestDef := GenReqDefForShowHttpBlockTrustIpRule()
	return &ShowHttpBlockTrustIpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpBlockTrustIpRules 查询IP黑白名单规则列表
//
// 查询IP黑白名单规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpBlockTrustIpRules(request *model.ShowHttpBlockTrustIpRulesRequest) (*model.ShowHttpBlockTrustIpRulesResponse, error) {
	requestDef := GenReqDefForShowHttpBlockTrustIpRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpBlockTrustIpRulesResponse), nil
	}
}

// ShowHttpBlockTrustIpRulesInvoker 查询IP黑白名单规则列表
func (c *EdgeSecClient) ShowHttpBlockTrustIpRulesInvoker(request *model.ShowHttpBlockTrustIpRulesRequest) *ShowHttpBlockTrustIpRulesInvoker {
	requestDef := GenReqDefForShowHttpBlockTrustIpRules()
	return &ShowHttpBlockTrustIpRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHttpBlockTrustIpRule 更新IP黑白名单规则
//
// 更新IP黑白名单规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) UpdateHttpBlockTrustIpRule(request *model.UpdateHttpBlockTrustIpRuleRequest) (*model.UpdateHttpBlockTrustIpRuleResponse, error) {
	requestDef := GenReqDefForUpdateHttpBlockTrustIpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHttpBlockTrustIpRuleResponse), nil
	}
}

// UpdateHttpBlockTrustIpRuleInvoker 更新IP黑白名单规则
func (c *EdgeSecClient) UpdateHttpBlockTrustIpRuleInvoker(request *model.UpdateHttpBlockTrustIpRuleRequest) *UpdateHttpBlockTrustIpRuleInvoker {
	requestDef := GenReqDefForUpdateHttpBlockTrustIpRule()
	return &UpdateHttpBlockTrustIpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHttpIpGroup 创建IP地址组
//
// 创建IP地址组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) CreateHttpIpGroup(request *model.CreateHttpIpGroupRequest) (*model.CreateHttpIpGroupResponse, error) {
	requestDef := GenReqDefForCreateHttpIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHttpIpGroupResponse), nil
	}
}

// CreateHttpIpGroupInvoker 创建IP地址组
func (c *EdgeSecClient) CreateHttpIpGroupInvoker(request *model.CreateHttpIpGroupRequest) *CreateHttpIpGroupInvoker {
	requestDef := GenReqDefForCreateHttpIpGroup()
	return &CreateHttpIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHttpIpGroup 删除IP地址组
//
// 删除IP地址组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) DeleteHttpIpGroup(request *model.DeleteHttpIpGroupRequest) (*model.DeleteHttpIpGroupResponse, error) {
	requestDef := GenReqDefForDeleteHttpIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHttpIpGroupResponse), nil
	}
}

// DeleteHttpIpGroupInvoker 删除IP地址组
func (c *EdgeSecClient) DeleteHttpIpGroupInvoker(request *model.DeleteHttpIpGroupRequest) *DeleteHttpIpGroupInvoker {
	requestDef := GenReqDefForDeleteHttpIpGroup()
	return &DeleteHttpIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpIpGroup 查询IP地址组
//
// 查询IP地址组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpIpGroup(request *model.ShowHttpIpGroupRequest) (*model.ShowHttpIpGroupResponse, error) {
	requestDef := GenReqDefForShowHttpIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpIpGroupResponse), nil
	}
}

// ShowHttpIpGroupInvoker 查询IP地址组
func (c *EdgeSecClient) ShowHttpIpGroupInvoker(request *model.ShowHttpIpGroupRequest) *ShowHttpIpGroupInvoker {
	requestDef := GenReqDefForShowHttpIpGroup()
	return &ShowHttpIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpIpGroups 查询IP地址组列表
//
// 查询IP地址组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpIpGroups(request *model.ShowHttpIpGroupsRequest) (*model.ShowHttpIpGroupsResponse, error) {
	requestDef := GenReqDefForShowHttpIpGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpIpGroupsResponse), nil
	}
}

// ShowHttpIpGroupsInvoker 查询IP地址组列表
func (c *EdgeSecClient) ShowHttpIpGroupsInvoker(request *model.ShowHttpIpGroupsRequest) *ShowHttpIpGroupsInvoker {
	requestDef := GenReqDefForShowHttpIpGroups()
	return &ShowHttpIpGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHttpIpGroup 更新IP地址组
//
// 更新IP地址组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) UpdateHttpIpGroup(request *model.UpdateHttpIpGroupRequest) (*model.UpdateHttpIpGroupResponse, error) {
	requestDef := GenReqDefForUpdateHttpIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHttpIpGroupResponse), nil
	}
}

// UpdateHttpIpGroupInvoker 更新IP地址组
func (c *EdgeSecClient) UpdateHttpIpGroupInvoker(request *model.UpdateHttpIpGroupRequest) *UpdateHttpIpGroupInvoker {
	requestDef := GenReqDefForUpdateHttpIpGroup()
	return &UpdateHttpIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
