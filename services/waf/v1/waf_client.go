package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
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

// ApplyCertificateToHost 绑定证书到域名
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

// ApplyCertificateToHostInvoker 绑定证书到域名
func (c *WafClient) ApplyCertificateToHostInvoker(request *model.ApplyCertificateToHostRequest) *ApplyCertificateToHostInvoker {
	requestDef := GenReqDefForApplyCertificateToHost()
	return &ApplyCertificateToHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAntiTamperRule 创建防篡改规则
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

// CreateAntiTamperRuleInvoker 创建防篡改规则
func (c *WafClient) CreateAntiTamperRuleInvoker(request *model.CreateAntiTamperRuleRequest) *CreateAntiTamperRuleInvoker {
	requestDef := GenReqDefForCreateAntiTamperRule()
	return &CreateAntiTamperRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCertificate 创建证书
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

// CreateCertificateInvoker 创建证书
func (c *WafClient) CreateCertificateInvoker(request *model.CreateCertificateRequest) *CreateCertificateInvoker {
	requestDef := GenReqDefForCreateCertificate()
	return &CreateCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGeoipRule 创建地理位置控制规则
//
// 创建地理位置控制规则
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

// CreateGeoipRuleInvoker 创建地理位置控制规则
func (c *WafClient) CreateGeoipRuleInvoker(request *model.CreateGeoipRuleRequest) *CreateGeoipRuleInvoker {
	requestDef := GenReqDefForCreateGeoipRule()
	return &CreateGeoipRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHost 创建云模式防护域名
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

// CreateHostInvoker 创建云模式防护域名
func (c *WafClient) CreateHostInvoker(request *model.CreateHostRequest) *CreateHostInvoker {
	requestDef := GenReqDefForCreateHost()
	return &CreateHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateIgnoreRule 创建全局白名单(原误报屏蔽)规则
//
// 创建全局白名单(原误报屏蔽)规则
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

// CreateIgnoreRuleInvoker 创建全局白名单(原误报屏蔽)规则
func (c *WafClient) CreateIgnoreRuleInvoker(request *model.CreateIgnoreRuleRequest) *CreateIgnoreRuleInvoker {
	requestDef := GenReqDefForCreateIgnoreRule()
	return &CreateIgnoreRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePolicy 创建防护策略
//
// 创建防护策略，系统会在生成策略时配置一些默认的配置项，如果需要修改策略的默认配置项需要通过调用更新防护策略接口实现
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

// CreatePolicyInvoker 创建防护策略
func (c *WafClient) CreatePolicyInvoker(request *model.CreatePolicyRequest) *CreatePolicyInvoker {
	requestDef := GenReqDefForCreatePolicy()
	return &CreatePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePremiumHost 创建独享模式域名
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

// CreatePremiumHostInvoker 创建独享模式域名
func (c *WafClient) CreatePremiumHostInvoker(request *model.CreatePremiumHostRequest) *CreatePremiumHostInvoker {
	requestDef := GenReqDefForCreatePremiumHost()
	return &CreatePremiumHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePrivacyRule 创建隐私屏蔽防护规则
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

// CreatePrivacyRuleInvoker 创建隐私屏蔽防护规则
func (c *WafClient) CreatePrivacyRuleInvoker(request *model.CreatePrivacyRuleRequest) *CreatePrivacyRuleInvoker {
	requestDef := GenReqDefForCreatePrivacyRule()
	return &CreatePrivacyRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateValueList 创建引用表
//
// 创建引用表，引用表能够被CC攻击防护规则和精准访问防护中的规则所引用。
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

// CreateValueListInvoker 创建引用表
func (c *WafClient) CreateValueListInvoker(request *model.CreateValueListRequest) *CreateValueListInvoker {
	requestDef := GenReqDefForCreateValueList()
	return &CreateValueListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWhiteblackipRule 创建黑白名单规则
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

// CreateWhiteblackipRuleInvoker 创建黑白名单规则
func (c *WafClient) CreateWhiteblackipRuleInvoker(request *model.CreateWhiteblackipRuleRequest) *CreateWhiteblackipRuleInvoker {
	requestDef := GenReqDefForCreateWhiteblackipRule()
	return &CreateWhiteblackipRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAntitamperRule 删除防篡改防护规则
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

// DeleteAntitamperRuleInvoker 删除防篡改防护规则
func (c *WafClient) DeleteAntitamperRuleInvoker(request *model.DeleteAntitamperRuleRequest) *DeleteAntitamperRuleInvoker {
	requestDef := GenReqDefForDeleteAntitamperRule()
	return &DeleteAntitamperRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCertificate 删除证书
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

// DeleteCertificateInvoker 删除证书
func (c *WafClient) DeleteCertificateInvoker(request *model.DeleteCertificateRequest) *DeleteCertificateInvoker {
	requestDef := GenReqDefForDeleteCertificate()
	return &DeleteCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGeoipRule 删除地理位置控制防护规则
//
// 删除地理位置控制防护规则
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

// DeleteGeoipRuleInvoker 删除地理位置控制防护规则
func (c *WafClient) DeleteGeoipRuleInvoker(request *model.DeleteGeoipRuleRequest) *DeleteGeoipRuleInvoker {
	requestDef := GenReqDefForDeleteGeoipRule()
	return &DeleteGeoipRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHost 删除云模式防护域名
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

// DeleteHostInvoker 删除云模式防护域名
func (c *WafClient) DeleteHostInvoker(request *model.DeleteHostRequest) *DeleteHostInvoker {
	requestDef := GenReqDefForDeleteHost()
	return &DeleteHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteIgnoreRule 删除全局白名单(原误报屏蔽)防护规则
//
// 删除全局白名单(原误报屏蔽)防护规则
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

// DeleteIgnoreRuleInvoker 删除全局白名单(原误报屏蔽)防护规则
func (c *WafClient) DeleteIgnoreRuleInvoker(request *model.DeleteIgnoreRuleRequest) *DeleteIgnoreRuleInvoker {
	requestDef := GenReqDefForDeleteIgnoreRule()
	return &DeleteIgnoreRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePolicy 删除防护策略
//
// 删除防护策略，若策略正在使用，则需要先接解除域名与策略的绑定关系才能删除策略。
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

// DeletePolicyInvoker 删除防护策略
func (c *WafClient) DeletePolicyInvoker(request *model.DeletePolicyRequest) *DeletePolicyInvoker {
	requestDef := GenReqDefForDeletePolicy()
	return &DeletePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePremiumHost 删除独享模式域名
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

// DeletePremiumHostInvoker 删除独享模式域名
func (c *WafClient) DeletePremiumHostInvoker(request *model.DeletePremiumHostRequest) *DeletePremiumHostInvoker {
	requestDef := GenReqDefForDeletePremiumHost()
	return &DeletePremiumHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePrivacyRule 删除隐私屏蔽防护规则
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

// DeletePrivacyRuleInvoker 删除隐私屏蔽防护规则
func (c *WafClient) DeletePrivacyRuleInvoker(request *model.DeletePrivacyRuleRequest) *DeletePrivacyRuleInvoker {
	requestDef := GenReqDefForDeletePrivacyRule()
	return &DeletePrivacyRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteValueList 删除引用表
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

// DeleteValueListInvoker 删除引用表
func (c *WafClient) DeleteValueListInvoker(request *model.DeleteValueListRequest) *DeleteValueListInvoker {
	requestDef := GenReqDefForDeleteValueList()
	return &DeleteValueListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWhiteBlackIpRule 删除黑白名单防护规则
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

// DeleteWhiteBlackIpRuleInvoker 删除黑白名单防护规则
func (c *WafClient) DeleteWhiteBlackIpRuleInvoker(request *model.DeleteWhiteBlackIpRuleRequest) *DeleteWhiteBlackIpRuleInvoker {
	requestDef := GenReqDefForDeleteWhiteBlackIpRule()
	return &DeleteWhiteBlackIpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAntitamperRule 查询防篡改规则列表
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

// ListAntitamperRuleInvoker 查询防篡改规则列表
func (c *WafClient) ListAntitamperRuleInvoker(request *model.ListAntitamperRuleRequest) *ListAntitamperRuleInvoker {
	requestDef := GenReqDefForListAntitamperRule()
	return &ListAntitamperRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBandwidthTimeline 查询安全统计带宽数据
//
// 查询安全统计带宽数据。需要注意的是，安全统计相关的接口，暂时不能支持任意时间的查询。只能支持 console上显示的 昨天，今天，3天，7天和30天 数据查询。
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

// ListBandwidthTimelineInvoker 查询安全统计带宽数据
func (c *WafClient) ListBandwidthTimelineInvoker(request *model.ListBandwidthTimelineRequest) *ListBandwidthTimelineInvoker {
	requestDef := GenReqDefForListBandwidthTimeline()
	return &ListBandwidthTimelineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCertificates 查询证书列表
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

// ListCertificatesInvoker 查询证书列表
func (c *WafClient) ListCertificatesInvoker(request *model.ListCertificatesRequest) *ListCertificatesInvoker {
	requestDef := GenReqDefForListCertificates()
	return &ListCertificatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCompositeHosts 查询全部防护域名列表
//
// 查询全部防护域名列表，包括云模式和独享模式
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

// ListCompositeHostsInvoker 查询全部防护域名列表
func (c *WafClient) ListCompositeHostsInvoker(request *model.ListCompositeHostsRequest) *ListCompositeHostsInvoker {
	requestDef := GenReqDefForListCompositeHosts()
	return &ListCompositeHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEvent 查询攻击事件列表
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

// ListEventInvoker 查询攻击事件列表
func (c *WafClient) ListEventInvoker(request *model.ListEventRequest) *ListEventInvoker {
	requestDef := GenReqDefForListEvent()
	return &ListEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGeoipRule 查询地理位置访问控制规则列表
//
// 查询地理位置访问控制规则列表
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

// ListGeoipRuleInvoker 查询地理位置访问控制规则列表
func (c *WafClient) ListGeoipRuleInvoker(request *model.ListGeoipRuleRequest) *ListGeoipRuleInvoker {
	requestDef := GenReqDefForListGeoipRule()
	return &ListGeoipRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHost 查询云模式防护域名列表
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

// ListHostInvoker 查询云模式防护域名列表
func (c *WafClient) ListHostInvoker(request *model.ListHostRequest) *ListHostInvoker {
	requestDef := GenReqDefForListHost()
	return &ListHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHostRoute 获取云模式域名路由信息
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

// ListHostRouteInvoker 获取云模式域名路由信息
func (c *WafClient) ListHostRouteInvoker(request *model.ListHostRouteRequest) *ListHostRouteInvoker {
	requestDef := GenReqDefForListHostRoute()
	return &ListHostRouteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIgnoreRule 查询全局白名单(原误报屏蔽)规则列表
//
// 查询全局白名单(原误报屏蔽)规则列表
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

// ListIgnoreRuleInvoker 查询全局白名单(原误报屏蔽)规则列表
func (c *WafClient) ListIgnoreRuleInvoker(request *model.ListIgnoreRuleRequest) *ListIgnoreRuleInvoker {
	requestDef := GenReqDefForListIgnoreRule()
	return &ListIgnoreRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOverviewsClassification 查询安全总览分类统计top信息
//
// 查询安全总览分类统计TOP信息，包含受攻击域名 、攻击源ip、受攻击URL、攻击来源区域、攻击事件分布。需要注意的是，安全总览相关的接口，暂时不能支持任意时间的查询。只能支持 console上显示的 昨天，今天，3天，7天和30天 数据查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WafClient) ListOverviewsClassification(request *model.ListOverviewsClassificationRequest) (*model.ListOverviewsClassificationResponse, error) {
	requestDef := GenReqDefForListOverviewsClassification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOverviewsClassificationResponse), nil
	}
}

// ListOverviewsClassificationInvoker 查询安全总览分类统计top信息
func (c *WafClient) ListOverviewsClassificationInvoker(request *model.ListOverviewsClassificationRequest) *ListOverviewsClassificationInvoker {
	requestDef := GenReqDefForListOverviewsClassification()
	return &ListOverviewsClassificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicy 查询防护策略列表
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

// ListPolicyInvoker 查询防护策略列表
func (c *WafClient) ListPolicyInvoker(request *model.ListPolicyRequest) *ListPolicyInvoker {
	requestDef := GenReqDefForListPolicy()
	return &ListPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPremiumHost 独享模式域名列表
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

// ListPremiumHostInvoker 独享模式域名列表
func (c *WafClient) ListPremiumHostInvoker(request *model.ListPremiumHostRequest) *ListPremiumHostInvoker {
	requestDef := GenReqDefForListPremiumHost()
	return &ListPremiumHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPrivacyRule 查询隐私屏蔽防护规则
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

// ListPrivacyRuleInvoker 查询隐私屏蔽防护规则
func (c *WafClient) ListPrivacyRuleInvoker(request *model.ListPrivacyRuleRequest) *ListPrivacyRuleInvoker {
	requestDef := GenReqDefForListPrivacyRule()
	return &ListPrivacyRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQpsTimeline 查询安全统计qps次数
//
// 查询安全统计qps次数。需要注意的是，安全统计相关的接口，暂时不能支持任意时间的查询。只能支持 console上显示的 昨天，今天，3天，7天和30天 数据查询。
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

// ListQpsTimelineInvoker 查询安全统计qps次数
func (c *WafClient) ListQpsTimelineInvoker(request *model.ListQpsTimelineRequest) *ListQpsTimelineInvoker {
	requestDef := GenReqDefForListQpsTimeline()
	return &ListQpsTimelineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStatistics 查询安全总览请求与攻击数量
//
// 查询安全总览请求与攻击数量。需要注意的是，安全总览相关的接口，暂时不能支持任意时间的查询。只能支持 console上显示的 昨天，今天，3天，7天和30天 数据查询。
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

// ListStatisticsInvoker 查询安全总览请求与攻击数量
func (c *WafClient) ListStatisticsInvoker(request *model.ListStatisticsRequest) *ListStatisticsInvoker {
	requestDef := GenReqDefForListStatistics()
	return &ListStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTopAbnormal 查询业务异常数量
//
// 查询业务异常TOP统计信息。需要注意的是，安全总览相关的接口，暂时不能支持任意时间的查询。只能支持 console上显示的 昨天，今天，3天，7天和30天 数据查询。
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

// ListTopAbnormalInvoker 查询业务异常数量
func (c *WafClient) ListTopAbnormalInvoker(request *model.ListTopAbnormalRequest) *ListTopAbnormalInvoker {
	requestDef := GenReqDefForListTopAbnormal()
	return &ListTopAbnormalInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListValueList 查询引用表列表
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

// ListValueListInvoker 查询引用表列表
func (c *WafClient) ListValueListInvoker(request *model.ListValueListRequest) *ListValueListInvoker {
	requestDef := GenReqDefForListValueList()
	return &ListValueListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWhiteblackipRule 查询黑白名单规则列表
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

// ListWhiteblackipRuleInvoker 查询黑白名单规则列表
func (c *WafClient) ListWhiteblackipRuleInvoker(request *model.ListWhiteblackipRuleRequest) *ListWhiteblackipRuleInvoker {
	requestDef := GenReqDefForListWhiteblackipRule()
	return &ListWhiteblackipRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCertificate 查询证书
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

// ShowCertificateInvoker 查询证书
func (c *WafClient) ShowCertificateInvoker(request *model.ShowCertificateRequest) *ShowCertificateInvoker {
	requestDef := GenReqDefForShowCertificate()
	return &ShowCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCompositeHost 根据Id查询防护域名
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

// ShowCompositeHostInvoker 根据Id查询防护域名
func (c *WafClient) ShowCompositeHostInvoker(request *model.ShowCompositeHostRequest) *ShowCompositeHostInvoker {
	requestDef := GenReqDefForShowCompositeHost()
	return &ShowCompositeHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConsoleConfig 局点支持特性查询
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

// ShowConsoleConfigInvoker 局点支持特性查询
func (c *WafClient) ShowConsoleConfigInvoker(request *model.ShowConsoleConfigRequest) *ShowConsoleConfigInvoker {
	requestDef := GenReqDefForShowConsoleConfig()
	return &ShowConsoleConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEvent 查询指定事件id的防护事件详情
//
// 查询指定事件id的防护事件详情
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

// ShowEventInvoker 查询指定事件id的防护事件详情
func (c *WafClient) ShowEventInvoker(request *model.ShowEventRequest) *ShowEventInvoker {
	requestDef := GenReqDefForShowEvent()
	return &ShowEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHost 根据防护域名Id查询云模式防护域名详细信息
//
// 根据防护域名Id查询云模式防护域名详细信息
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

// ShowHostInvoker 根据防护域名Id查询云模式防护域名详细信息
func (c *WafClient) ShowHostInvoker(request *model.ShowHostRequest) *ShowHostInvoker {
	requestDef := GenReqDefForShowHost()
	return &ShowHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPolicy 根据Id查询防护策略
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

// ShowPolicyInvoker 根据Id查询防护策略
func (c *WafClient) ShowPolicyInvoker(request *model.ShowPolicyRequest) *ShowPolicyInvoker {
	requestDef := GenReqDefForShowPolicy()
	return &ShowPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPremiumHost 查看独享模式域名配置
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

// ShowPremiumHostInvoker 查看独享模式域名配置
func (c *WafClient) ShowPremiumHostInvoker(request *model.ShowPremiumHostRequest) *ShowPremiumHostInvoker {
	requestDef := GenReqDefForShowPremiumHost()
	return &ShowPremiumHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCertificate 修改证书
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

// UpdateCertificateInvoker 修改证书
func (c *WafClient) UpdateCertificateInvoker(request *model.UpdateCertificateRequest) *UpdateCertificateInvoker {
	requestDef := GenReqDefForUpdateCertificate()
	return &UpdateCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGeoipRule 更新地理位置控制防护规则
//
// 更新地理位置控制防护规则
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

// UpdateGeoipRuleInvoker 更新地理位置控制防护规则
func (c *WafClient) UpdateGeoipRuleInvoker(request *model.UpdateGeoipRuleRequest) *UpdateGeoipRuleInvoker {
	requestDef := GenReqDefForUpdateGeoipRule()
	return &UpdateGeoipRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHost 更新云模式防护域名的配置
//
// 更新云模式防护域名配置，在没有填入源站信息server的原始数据的情况下，则新的源站信息server会覆盖源站信息，而不是新增源站。此外，请求体可只传需要更新的部分。
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

// UpdateHostInvoker 更新云模式防护域名的配置
func (c *WafClient) UpdateHostInvoker(request *model.UpdateHostRequest) *UpdateHostInvoker {
	requestDef := GenReqDefForUpdateHost()
	return &UpdateHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHostProtectStatus 修改域名防护状态
//
// 修改域名防护状态
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

// UpdateHostProtectStatusInvoker 修改域名防护状态
func (c *WafClient) UpdateHostProtectStatusInvoker(request *model.UpdateHostProtectStatusRequest) *UpdateHostProtectStatusInvoker {
	requestDef := GenReqDefForUpdateHostProtectStatus()
	return &UpdateHostProtectStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePolicy 更新防护策略
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

// UpdatePolicyInvoker 更新防护策略
func (c *WafClient) UpdatePolicyInvoker(request *model.UpdatePolicyRequest) *UpdatePolicyInvoker {
	requestDef := GenReqDefForUpdatePolicy()
	return &UpdatePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePolicyProtectHost 更新防护策略的域名
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

// UpdatePolicyProtectHostInvoker 更新防护策略的域名
func (c *WafClient) UpdatePolicyProtectHostInvoker(request *model.UpdatePolicyProtectHostRequest) *UpdatePolicyProtectHostInvoker {
	requestDef := GenReqDefForUpdatePolicyProtectHost()
	return &UpdatePolicyProtectHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePolicyRuleStatus 修改单条规则的状态
//
// 修改单条规则的状态，用于开启或者关闭单条规则，比如关闭精准防护中某一条规则
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

// UpdatePolicyRuleStatusInvoker 修改单条规则的状态
func (c *WafClient) UpdatePolicyRuleStatusInvoker(request *model.UpdatePolicyRuleStatusRequest) *UpdatePolicyRuleStatusInvoker {
	requestDef := GenReqDefForUpdatePolicyRuleStatus()
	return &UpdatePolicyRuleStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePremiumHost 修改独享模式域名配置
//
// 修改独享模式域名配置，在没有填入源站信息server的原始数据的情况下，则新的源站信息server会覆盖源站信息，而不是新增源站。此外，请求体可只传需要更新的部分。
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

// UpdatePremiumHostInvoker 修改独享模式域名配置
func (c *WafClient) UpdatePremiumHostInvoker(request *model.UpdatePremiumHostRequest) *UpdatePremiumHostInvoker {
	requestDef := GenReqDefForUpdatePremiumHost()
	return &UpdatePremiumHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePremiumHostProtectStatus 修改独享模式域名防护状态
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

// UpdatePremiumHostProtectStatusInvoker 修改独享模式域名防护状态
func (c *WafClient) UpdatePremiumHostProtectStatusInvoker(request *model.UpdatePremiumHostProtectStatusRequest) *UpdatePremiumHostProtectStatusInvoker {
	requestDef := GenReqDefForUpdatePremiumHostProtectStatus()
	return &UpdatePremiumHostProtectStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePrivacyRule 更新隐私屏蔽防护规则
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

// UpdatePrivacyRuleInvoker 更新隐私屏蔽防护规则
func (c *WafClient) UpdatePrivacyRuleInvoker(request *model.UpdatePrivacyRuleRequest) *UpdatePrivacyRuleInvoker {
	requestDef := GenReqDefForUpdatePrivacyRule()
	return &UpdatePrivacyRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateValueList 修改引用表
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

// UpdateValueListInvoker 修改引用表
func (c *WafClient) UpdateValueListInvoker(request *model.UpdateValueListRequest) *UpdateValueListInvoker {
	requestDef := GenReqDefForUpdateValueList()
	return &UpdateValueListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWhiteblackipRule 更新黑白名单防护规则
//
// 更新黑白名单防护规则，可以更新ip/ip段以及防护动作等信息
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

// UpdateWhiteblackipRuleInvoker 更新黑白名单防护规则
func (c *WafClient) UpdateWhiteblackipRuleInvoker(request *model.UpdateWhiteblackipRuleRequest) *UpdateWhiteblackipRuleInvoker {
	requestDef := GenReqDefForUpdateWhiteblackipRule()
	return &UpdateWhiteblackipRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
