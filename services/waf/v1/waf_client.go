package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/waf/v1/model"
)

type WafClient struct {
	HcClient *http_client.HcHttpClient
}

func NewWafClient(hcClient *http_client.HcHttpClient) *WafClient {
	return &WafClient{HcClient: hcClient}
}

func WafClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 绑定证书到域名
//
// 绑定证书到域名
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ApplyCertificateToHost(request *model.ApplyCertificateToHostRequest) (*model.ApplyCertificateToHostResponse, error) {
	requestDef := GenReqDefForApplyCertificateToHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApplyCertificateToHostResponse), nil
	}
}

// 创建防篡改规则
//
// 创建防篡改规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) CreateAntiTamperRule(request *model.CreateAntiTamperRuleRequest) (*model.CreateAntiTamperRuleResponse, error) {
	requestDef := GenReqDefForCreateAntiTamperRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAntiTamperRuleResponse), nil
	}
}

// 创建证书
//
// 创建证书
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) CreateCertificate(request *model.CreateCertificateRequest) (*model.CreateCertificateResponse, error) {
	requestDef := GenReqDefForCreateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCertificateResponse), nil
	}
}

// 创建地理位置规则
//
// 创建地理位置规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) CreateGeoipRule(request *model.CreateGeoipRuleRequest) (*model.CreateGeoipRuleResponse, error) {
	requestDef := GenReqDefForCreateGeoipRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGeoipRuleResponse), nil
	}
}

// 创建云模式防护域名
//
// 创建云模式防护域名
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) CreateHost(request *model.CreateHostRequest) (*model.CreateHostResponse, error) {
	requestDef := GenReqDefForCreateHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHostResponse), nil
	}
}

// 创建误报屏蔽规则
//
// 创建误报屏蔽规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) CreateIgnoreRule(request *model.CreateIgnoreRuleRequest) (*model.CreateIgnoreRuleResponse, error) {
	requestDef := GenReqDefForCreateIgnoreRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIgnoreRuleResponse), nil
	}
}

// 创建防护策略
//
// 创建防护策略
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) CreatePolicy(request *model.CreatePolicyRequest) (*model.CreatePolicyResponse, error) {
	requestDef := GenReqDefForCreatePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePolicyResponse), nil
	}
}

// 创建独享模式域名
//
// 创建独享模式域名
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) CreatePremiumHost(request *model.CreatePremiumHostRequest) (*model.CreatePremiumHostResponse, error) {
	requestDef := GenReqDefForCreatePremiumHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePremiumHostResponse), nil
	}
}

// 创建隐私屏蔽防护规则
//
// 创建隐私屏蔽防护规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) CreatePrivacyRule(request *model.CreatePrivacyRuleRequest) (*model.CreatePrivacyRuleResponse, error) {
	requestDef := GenReqDefForCreatePrivacyRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrivacyRuleResponse), nil
	}
}

// 创建引用表
//
// 创建引用表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) CreateValueList(request *model.CreateValueListRequest) (*model.CreateValueListResponse, error) {
	requestDef := GenReqDefForCreateValueList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateValueListResponse), nil
	}
}

// 创建黑白名单规则
//
// 创建黑白名单规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) CreateWhiteblackipRule(request *model.CreateWhiteblackipRuleRequest) (*model.CreateWhiteblackipRuleResponse, error) {
	requestDef := GenReqDefForCreateWhiteblackipRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWhiteblackipRuleResponse), nil
	}
}

// 删除防篡改防护规则
//
// 删除防篡改防护规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) DeleteAntitamperRule(request *model.DeleteAntitamperRuleRequest) (*model.DeleteAntitamperRuleResponse, error) {
	requestDef := GenReqDefForDeleteAntitamperRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAntitamperRuleResponse), nil
	}
}

// 删除证书
//
// 删除证书
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) DeleteCertificate(request *model.DeleteCertificateRequest) (*model.DeleteCertificateResponse, error) {
	requestDef := GenReqDefForDeleteCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCertificateResponse), nil
	}
}

// 删除地理位置防护规则
//
// 删除地理位置防护规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) DeleteGeoipRule(request *model.DeleteGeoipRuleRequest) (*model.DeleteGeoipRuleResponse, error) {
	requestDef := GenReqDefForDeleteGeoipRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGeoipRuleResponse), nil
	}
}

// 删除云模式防护域名
//
// 删除云模式防护域名
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) DeleteHost(request *model.DeleteHostRequest) (*model.DeleteHostResponse, error) {
	requestDef := GenReqDefForDeleteHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHostResponse), nil
	}
}

// 删除误报屏蔽防护规则
//
// 删除误报屏蔽防护规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) DeleteIgnoreRule(request *model.DeleteIgnoreRuleRequest) (*model.DeleteIgnoreRuleResponse, error) {
	requestDef := GenReqDefForDeleteIgnoreRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIgnoreRuleResponse), nil
	}
}

// 删除防护策略
//
// 删除防护策略
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) DeletePolicy(request *model.DeletePolicyRequest) (*model.DeletePolicyResponse, error) {
	requestDef := GenReqDefForDeletePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePolicyResponse), nil
	}
}

// 删除独享模式域名
//
// 删除独享模式域名
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) DeletePremiumHost(request *model.DeletePremiumHostRequest) (*model.DeletePremiumHostResponse, error) {
	requestDef := GenReqDefForDeletePremiumHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePremiumHostResponse), nil
	}
}

// 删除隐私屏蔽防护规则
//
// 删除隐私屏蔽防护规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) DeletePrivacyRule(request *model.DeletePrivacyRuleRequest) (*model.DeletePrivacyRuleResponse, error) {
	requestDef := GenReqDefForDeletePrivacyRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePrivacyRuleResponse), nil
	}
}

// 删除引用表
//
// 删除引用表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) DeleteValueList(request *model.DeleteValueListRequest) (*model.DeleteValueListResponse, error) {
	requestDef := GenReqDefForDeleteValueList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteValueListResponse), nil
	}
}

// 删除黑白名单防护规则
//
// 删除黑白名单防护规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) DeleteWhiteBlackIpRule(request *model.DeleteWhiteBlackIpRuleRequest) (*model.DeleteWhiteBlackIpRuleResponse, error) {
	requestDef := GenReqDefForDeleteWhiteBlackIpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWhiteBlackIpRuleResponse), nil
	}
}

// 查询防篡改规则列表
//
// 查询防篡改规则列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ListAntitamperRule(request *model.ListAntitamperRuleRequest) (*model.ListAntitamperRuleResponse, error) {
	requestDef := GenReqDefForListAntitamperRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAntitamperRuleResponse), nil
	}
}

// 查询安全统计带宽数据
//
// 查询安全统计带宽数据
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ListBandwidthTimeline(request *model.ListBandwidthTimelineRequest) (*model.ListBandwidthTimelineResponse, error) {
	requestDef := GenReqDefForListBandwidthTimeline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBandwidthTimelineResponse), nil
	}
}

// 查询证书列表
//
// 查询证书列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ListCertificates(request *model.ListCertificatesRequest) (*model.ListCertificatesResponse, error) {
	requestDef := GenReqDefForListCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCertificatesResponse), nil
	}
}

// 查询全部防护域名列表
//
// 查询全部防护域名列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ListCompositeHosts(request *model.ListCompositeHostsRequest) (*model.ListCompositeHostsResponse, error) {
	requestDef := GenReqDefForListCompositeHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCompositeHostsResponse), nil
	}
}

// 查询攻击事件列表
//
// 查询攻击事件列表，该API暂时不支持查询全部防护事件，pagesize参数不可设为-1，由于性能原因，数据量越大消耗的内存越大，后端最多限制查询10000条数据，例如：自定义时间段内的数据超过了10000条，就无法查出page为101，pagesize为100之后的数据，需要调整时间区间，再进行查询
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ListEvent(request *model.ListEventRequest) (*model.ListEventResponse, error) {
	requestDef := GenReqDefForListEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventResponse), nil
	}
}

// 查询地理位置规则列表
//
// 查询地理位置规则列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ListGeoipRule(request *model.ListGeoipRuleRequest) (*model.ListGeoipRuleResponse, error) {
	requestDef := GenReqDefForListGeoipRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGeoipRuleResponse), nil
	}
}

// 查询云模式防护域名列表
//
// 查询云模式防护域名列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ListHost(request *model.ListHostRequest) (*model.ListHostResponse, error) {
	requestDef := GenReqDefForListHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostResponse), nil
	}
}

// 获取云模式域名路由信息
//
// 返回路由信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ListHostRoute(request *model.ListHostRouteRequest) (*model.ListHostRouteResponse, error) {
	requestDef := GenReqDefForListHostRoute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostRouteResponse), nil
	}
}

// 查询误报屏蔽规则列表
//
// 查询误报屏蔽规则列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ListIgnoreRule(request *model.ListIgnoreRuleRequest) (*model.ListIgnoreRuleResponse, error) {
	requestDef := GenReqDefForListIgnoreRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIgnoreRuleResponse), nil
	}
}

// 查询防护策略列表
//
// 查询防护策略列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ListPolicy(request *model.ListPolicyRequest) (*model.ListPolicyResponse, error) {
	requestDef := GenReqDefForListPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyResponse), nil
	}
}

// 独享模式域名列表
//
// 独享模式域名列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ListPremiumHost(request *model.ListPremiumHostRequest) (*model.ListPremiumHostResponse, error) {
	requestDef := GenReqDefForListPremiumHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPremiumHostResponse), nil
	}
}

// 查询隐私屏蔽防护规则
//
// 查询隐私屏蔽防护规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ListPrivacyRule(request *model.ListPrivacyRuleRequest) (*model.ListPrivacyRuleResponse, error) {
	requestDef := GenReqDefForListPrivacyRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivacyRuleResponse), nil
	}
}

// 查询安全统计qps次数
//
// 查询安全统计qps次数
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ListQpsTimeline(request *model.ListQpsTimelineRequest) (*model.ListQpsTimelineResponse, error) {
	requestDef := GenReqDefForListQpsTimeline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQpsTimelineResponse), nil
	}
}

// 查询安全总览请求数据
//
// 查询安全总览请求数据
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ListStatistics(request *model.ListStatisticsRequest) (*model.ListStatisticsResponse, error) {
	requestDef := GenReqDefForListStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStatisticsResponse), nil
	}
}

// 查询业务异常数量
//
// 查询业务异常数量
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ListTopAbnormal(request *model.ListTopAbnormalRequest) (*model.ListTopAbnormalResponse, error) {
	requestDef := GenReqDefForListTopAbnormal()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopAbnormalResponse), nil
	}
}

// 查询引用表列表
//
// 查询引用表列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ListValueList(request *model.ListValueListRequest) (*model.ListValueListResponse, error) {
	requestDef := GenReqDefForListValueList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListValueListResponse), nil
	}
}

// 查询黑白名单规则列表
//
// 查询黑白名单规则列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ListWhiteblackipRule(request *model.ListWhiteblackipRuleRequest) (*model.ListWhiteblackipRuleResponse, error) {
	requestDef := GenReqDefForListWhiteblackipRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWhiteblackipRuleResponse), nil
	}
}

// 查询证书
//
// 查询证书
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ShowCertificate(request *model.ShowCertificateRequest) (*model.ShowCertificateResponse, error) {
	requestDef := GenReqDefForShowCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificateResponse), nil
	}
}

// 根据Id查询防护域名
//
// 根据Id查询防护域名
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ShowCompositeHost(request *model.ShowCompositeHostRequest) (*model.ShowCompositeHostResponse, error) {
	requestDef := GenReqDefForShowCompositeHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCompositeHostResponse), nil
	}
}

// 局点支持特性查询
//
// 局点支持特性查询
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ShowConsoleConfig(request *model.ShowConsoleConfigRequest) (*model.ShowConsoleConfigResponse, error) {
	requestDef := GenReqDefForShowConsoleConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConsoleConfigResponse), nil
	}
}

// 查询攻击事件详情
//
// 查询攻击事件详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ShowEvent(request *model.ShowEventRequest) (*model.ShowEventResponse, error) {
	requestDef := GenReqDefForShowEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEventResponse), nil
	}
}

// 根据Id查询云模式防护域名
//
// 根据Id查询云模式防护域名
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ShowHost(request *model.ShowHostRequest) (*model.ShowHostResponse, error) {
	requestDef := GenReqDefForShowHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHostResponse), nil
	}
}

// 根据Id查询防护策略
//
// 根据Id查询防护策略
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ShowPolicy(request *model.ShowPolicyRequest) (*model.ShowPolicyResponse, error) {
	requestDef := GenReqDefForShowPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPolicyResponse), nil
	}
}

// 查看独享模式域名配置
//
// 查看独享模式域名配置
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ShowPremiumHost(request *model.ShowPremiumHostRequest) (*model.ShowPremiumHostResponse, error) {
	requestDef := GenReqDefForShowPremiumHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPremiumHostResponse), nil
	}
}

// 修改证书
//
// 修改证书
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) UpdateCertificate(request *model.UpdateCertificateRequest) (*model.UpdateCertificateResponse, error) {
	requestDef := GenReqDefForUpdateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCertificateResponse), nil
	}
}

// 更新地理位置防护规则
//
// 更新地理位置防护规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) UpdateGeoipRule(request *model.UpdateGeoipRuleRequest) (*model.UpdateGeoipRuleResponse, error) {
	requestDef := GenReqDefForUpdateGeoipRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGeoipRuleResponse), nil
	}
}

// 更新云模式防护域名
//
// 更新云模式防护域名配置，在没有填入源站信息server的原始数据的情况下，则新的源站信息server会覆盖源站信息，而不是新增源站
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) UpdateHost(request *model.UpdateHostRequest) (*model.UpdateHostResponse, error) {
	requestDef := GenReqDefForUpdateHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHostResponse), nil
	}
}

// 修改域名防护状态
//
// 返回路由信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) UpdateHostProtectStatus(request *model.UpdateHostProtectStatusRequest) (*model.UpdateHostProtectStatusResponse, error) {
	requestDef := GenReqDefForUpdateHostProtectStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHostProtectStatusResponse), nil
	}
}

// 更新防护策略
//
// 更新防护策略，请求体可只传需要更新的部分
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) UpdatePolicy(request *model.UpdatePolicyRequest) (*model.UpdatePolicyResponse, error) {
	requestDef := GenReqDefForUpdatePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePolicyResponse), nil
	}
}

// 更新防护策略的域名
//
// 更新防护策略的防护域名
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) UpdatePolicyProtectHost(request *model.UpdatePolicyProtectHostRequest) (*model.UpdatePolicyProtectHostResponse, error) {
	requestDef := GenReqDefForUpdatePolicyProtectHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePolicyProtectHostResponse), nil
	}
}

// 修改单条规则的状态
//
// 修改单条规则的状态
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) UpdatePolicyRuleStatus(request *model.UpdatePolicyRuleStatusRequest) (*model.UpdatePolicyRuleStatusResponse, error) {
	requestDef := GenReqDefForUpdatePolicyRuleStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePolicyRuleStatusResponse), nil
	}
}

// 修改独享模式域名配置
//
// 修改独享模式域名配置
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) UpdatePremiumHost(request *model.UpdatePremiumHostRequest) (*model.UpdatePremiumHostResponse, error) {
	requestDef := GenReqDefForUpdatePremiumHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePremiumHostResponse), nil
	}
}

// 修改独享模式域名防护状态
//
// 修改独享模式域名防护状态
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) UpdatePremiumHostProtectStatus(request *model.UpdatePremiumHostProtectStatusRequest) (*model.UpdatePremiumHostProtectStatusResponse, error) {
	requestDef := GenReqDefForUpdatePremiumHostProtectStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePremiumHostProtectStatusResponse), nil
	}
}

// 更新隐私屏蔽防护规则
//
// 更新隐私屏蔽防护规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) UpdatePrivacyRule(request *model.UpdatePrivacyRuleRequest) (*model.UpdatePrivacyRuleResponse, error) {
	requestDef := GenReqDefForUpdatePrivacyRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePrivacyRuleResponse), nil
	}
}

// 修改引用表
//
// 修改引用表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) UpdateValueList(request *model.UpdateValueListRequest) (*model.UpdateValueListResponse, error) {
	requestDef := GenReqDefForUpdateValueList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateValueListResponse), nil
	}
}

// 更新黑白名单防护规则
//
// 更新黑白名单防护规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) UpdateWhiteblackipRule(request *model.UpdateWhiteblackipRuleRequest) (*model.UpdateWhiteblackipRuleResponse, error) {
	requestDef := GenReqDefForUpdateWhiteblackipRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWhiteblackipRuleResponse), nil
	}
}
