package v5

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/hss/v5/model"
)

type HssClient struct {
	HcClient *http_client.HcHttpClient
}

func NewHssClient(hcClient *http_client.HcHttpClient) *HssClient {
	return &HssClient{HcClient: hcClient}
}

func HssClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// BatchCreateTags 批量创建标签
//
// 批量创建标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) BatchCreateTags(request *model.BatchCreateTagsRequest) (*model.BatchCreateTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateTagsResponse), nil
	}
}

// BatchCreateTagsInvoker 批量创建标签
func (c *HssClient) BatchCreateTagsInvoker(request *model.BatchCreateTagsRequest) *BatchCreateTagsInvoker {
	requestDef := GenReqDefForBatchCreateTags()
	return &BatchCreateTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResourceInstanceTag 删除资源标签
//
// 删除单个资源下的标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) DeleteResourceInstanceTag(request *model.DeleteResourceInstanceTagRequest) (*model.DeleteResourceInstanceTagResponse, error) {
	requestDef := GenReqDefForDeleteResourceInstanceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceInstanceTagResponse), nil
	}
}

// DeleteResourceInstanceTagInvoker 删除资源标签
func (c *HssClient) DeleteResourceInstanceTagInvoker(request *model.DeleteResourceInstanceTagRequest) *DeleteResourceInstanceTagInvoker {
	requestDef := GenReqDefForDeleteResourceInstanceTag()
	return &DeleteResourceInstanceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHostStatus 查询云服务器列表
//
// 查询云服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListHostStatus(request *model.ListHostStatusRequest) (*model.ListHostStatusResponse, error) {
	requestDef := GenReqDefForListHostStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostStatusResponse), nil
	}
}

// ListHostStatusInvoker 查询云服务器列表
func (c *HssClient) ListHostStatusInvoker(request *model.ListHostStatusRequest) *ListHostStatusInvoker {
	requestDef := GenReqDefForListHostStatus()
	return &ListHostStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPasswordComplexity 查询口令复杂度策略检测报告
//
// 查询口令复杂度策略检测报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListPasswordComplexity(request *model.ListPasswordComplexityRequest) (*model.ListPasswordComplexityResponse, error) {
	requestDef := GenReqDefForListPasswordComplexity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPasswordComplexityResponse), nil
	}
}

// ListPasswordComplexityInvoker 查询口令复杂度策略检测报告
func (c *HssClient) ListPasswordComplexityInvoker(request *model.ListPasswordComplexityRequest) *ListPasswordComplexityInvoker {
	requestDef := GenReqDefForListPasswordComplexity()
	return &ListPasswordComplexityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotasDetail 查询配额详情
//
// 查询配额详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListQuotasDetail(request *model.ListQuotasDetailRequest) (*model.ListQuotasDetailResponse, error) {
	requestDef := GenReqDefForListQuotasDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasDetailResponse), nil
	}
}

// ListQuotasDetailInvoker 查询配额详情
func (c *HssClient) ListQuotasDetailInvoker(request *model.ListQuotasDetailRequest) *ListQuotasDetailInvoker {
	requestDef := GenReqDefForListQuotasDetail()
	return &ListQuotasDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRiskConfigCheckRules 查询指定安全配置项的检查项列表
//
// 查询指定安全配置项的检查项列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListRiskConfigCheckRules(request *model.ListRiskConfigCheckRulesRequest) (*model.ListRiskConfigCheckRulesResponse, error) {
	requestDef := GenReqDefForListRiskConfigCheckRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRiskConfigCheckRulesResponse), nil
	}
}

// ListRiskConfigCheckRulesInvoker 查询指定安全配置项的检查项列表
func (c *HssClient) ListRiskConfigCheckRulesInvoker(request *model.ListRiskConfigCheckRulesRequest) *ListRiskConfigCheckRulesInvoker {
	requestDef := GenReqDefForListRiskConfigCheckRules()
	return &ListRiskConfigCheckRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRiskConfigHosts 查询指定安全配置项的受影响服务器列表
//
// 查询指定安全配置项的受影响服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListRiskConfigHosts(request *model.ListRiskConfigHostsRequest) (*model.ListRiskConfigHostsResponse, error) {
	requestDef := GenReqDefForListRiskConfigHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRiskConfigHostsResponse), nil
	}
}

// ListRiskConfigHostsInvoker 查询指定安全配置项的受影响服务器列表
func (c *HssClient) ListRiskConfigHostsInvoker(request *model.ListRiskConfigHostsRequest) *ListRiskConfigHostsInvoker {
	requestDef := GenReqDefForListRiskConfigHosts()
	return &ListRiskConfigHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRiskConfigs 查询租户的服务器安全配置检测结果列表
//
// 查询租户的服务器安全配置检测结果列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListRiskConfigs(request *model.ListRiskConfigsRequest) (*model.ListRiskConfigsResponse, error) {
	requestDef := GenReqDefForListRiskConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRiskConfigsResponse), nil
	}
}

// ListRiskConfigsInvoker 查询租户的服务器安全配置检测结果列表
func (c *HssClient) ListRiskConfigsInvoker(request *model.ListRiskConfigsRequest) *ListRiskConfigsInvoker {
	requestDef := GenReqDefForListRiskConfigs()
	return &ListRiskConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityEvents 查入侵事件列表
//
// 查入侵事件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListSecurityEvents(request *model.ListSecurityEventsRequest) (*model.ListSecurityEventsResponse, error) {
	requestDef := GenReqDefForListSecurityEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityEventsResponse), nil
	}
}

// ListSecurityEventsInvoker 查入侵事件列表
func (c *HssClient) ListSecurityEventsInvoker(request *model.ListSecurityEventsRequest) *ListSecurityEventsInvoker {
	requestDef := GenReqDefForListSecurityEvents()
	return &ListSecurityEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserChangeHistories 获取账户变动历史信息
//
// 获取账户变动历史记录信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListUserChangeHistories(request *model.ListUserChangeHistoriesRequest) (*model.ListUserChangeHistoriesResponse, error) {
	requestDef := GenReqDefForListUserChangeHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserChangeHistoriesResponse), nil
	}
}

// ListUserChangeHistoriesInvoker 获取账户变动历史信息
func (c *HssClient) ListUserChangeHistoriesInvoker(request *model.ListUserChangeHistoriesRequest) *ListUserChangeHistoriesInvoker {
	requestDef := GenReqDefForListUserChangeHistories()
	return &ListUserChangeHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsers 获取资产的账号列表
//
// 获取资产的账号列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListUsers(request *model.ListUsersRequest) (*model.ListUsersResponse, error) {
	requestDef := GenReqDefForListUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsersResponse), nil
	}
}

// ListUsersInvoker 获取资产的账号列表
func (c *HssClient) ListUsersInvoker(request *model.ListUsersRequest) *ListUsersInvoker {
	requestDef := GenReqDefForListUsers()
	return &ListUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVulnerabilities 查询漏洞列表
//
// 查询漏洞列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListVulnerabilities(request *model.ListVulnerabilitiesRequest) (*model.ListVulnerabilitiesResponse, error) {
	requestDef := GenReqDefForListVulnerabilities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVulnerabilitiesResponse), nil
	}
}

// ListVulnerabilitiesInvoker 查询漏洞列表
func (c *HssClient) ListVulnerabilitiesInvoker(request *model.ListVulnerabilitiesRequest) *ListVulnerabilitiesInvoker {
	requestDef := GenReqDefForListVulnerabilities()
	return &ListVulnerabilitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWeakPasswordUsers 查询弱口令检测结果列表
//
// 查询弱口令检测结果列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ListWeakPasswordUsers(request *model.ListWeakPasswordUsersRequest) (*model.ListWeakPasswordUsersResponse, error) {
	requestDef := GenReqDefForListWeakPasswordUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWeakPasswordUsersResponse), nil
	}
}

// ListWeakPasswordUsersInvoker 查询弱口令检测结果列表
func (c *HssClient) ListWeakPasswordUsersInvoker(request *model.ListWeakPasswordUsersRequest) *ListWeakPasswordUsersInvoker {
	requestDef := GenReqDefForListWeakPasswordUsers()
	return &ListWeakPasswordUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCheckRuleDetail 查询配置检查项检测报告
//
// 查询配置检查项检测报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowCheckRuleDetail(request *model.ShowCheckRuleDetailRequest) (*model.ShowCheckRuleDetailResponse, error) {
	requestDef := GenReqDefForShowCheckRuleDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCheckRuleDetailResponse), nil
	}
}

// ShowCheckRuleDetailInvoker 查询配置检查项检测报告
func (c *HssClient) ShowCheckRuleDetailInvoker(request *model.ShowCheckRuleDetailRequest) *ShowCheckRuleDetailInvoker {
	requestDef := GenReqDefForShowCheckRuleDetail()
	return &ShowCheckRuleDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceQuotas 查询配额信息
//
// 查询配额信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowResourceQuotas(request *model.ShowResourceQuotasRequest) (*model.ShowResourceQuotasResponse, error) {
	requestDef := GenReqDefForShowResourceQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceQuotasResponse), nil
	}
}

// ShowResourceQuotasInvoker 查询配额信息
func (c *HssClient) ShowResourceQuotasInvoker(request *model.ShowResourceQuotasRequest) *ShowResourceQuotasInvoker {
	requestDef := GenReqDefForShowResourceQuotas()
	return &ShowResourceQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRiskConfigDetail 查询指定安全配置项的检查结果
//
// 查询指定安全配置项的检查结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) ShowRiskConfigDetail(request *model.ShowRiskConfigDetailRequest) (*model.ShowRiskConfigDetailResponse, error) {
	requestDef := GenReqDefForShowRiskConfigDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRiskConfigDetailResponse), nil
	}
}

// ShowRiskConfigDetailInvoker 查询指定安全配置项的检查结果
func (c *HssClient) ShowRiskConfigDetailInvoker(request *model.ShowRiskConfigDetailRequest) *ShowRiskConfigDetailInvoker {
	requestDef := GenReqDefForShowRiskConfigDetail()
	return &ShowRiskConfigDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchHostsProtectStatus 切换防护状态
//
// 切换防护状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HssClient) SwitchHostsProtectStatus(request *model.SwitchHostsProtectStatusRequest) (*model.SwitchHostsProtectStatusResponse, error) {
	requestDef := GenReqDefForSwitchHostsProtectStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchHostsProtectStatusResponse), nil
	}
}

// SwitchHostsProtectStatusInvoker 切换防护状态
func (c *HssClient) SwitchHostsProtectStatusInvoker(request *model.SwitchHostsProtectStatusRequest) *SwitchHostsProtectStatusInvoker {
	requestDef := GenReqDefForSwitchHostsProtectStatus()
	return &SwitchHostsProtectStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
