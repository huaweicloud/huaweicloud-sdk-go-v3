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

// ListHostStatus 查询云服务器列表
//
// 查询云服务器列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ListRiskConfigCheckRules 查询指定安全配置项的检查项列表
//
// 查询指定安全配置项的检查项列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ListVulnerabilities 查询漏洞列表
//
// 查询漏洞列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ShowRiskConfigDetail 查询指定安全配置项的检查结果
//
// 查询指定安全配置项的检查结果
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
