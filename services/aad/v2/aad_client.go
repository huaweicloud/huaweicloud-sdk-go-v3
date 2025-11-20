package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/aad/v2/model"
)

type AadClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewAadClient(hcClient *httpclient.HcHttpClient) *AadClient {
	return &AadClient{HcClient: hcClient}
}

func AadClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// AddWafWhiteIpRule 防护策略web-cc黑白名单-创建黑白名单规则
//
// 防护策略web-cc黑白名单-创建黑白名单规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) AddWafWhiteIpRule(request *model.AddWafWhiteIpRuleRequest) (*model.AddWafWhiteIpRuleResponse, error) {
	requestDef := GenReqDefForAddWafWhiteIpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddWafWhiteIpRuleResponse), nil
	}
}

// AddWafWhiteIpRuleInvoker 防护策略web-cc黑白名单-创建黑白名单规则
func (c *AadClient) AddWafWhiteIpRuleInvoker(request *model.AddWafWhiteIpRuleRequest) *AddWafWhiteIpRuleInvoker {
	requestDef := GenReqDefForAddWafWhiteIpRule()
	return &AddWafWhiteIpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDomain 创建防护域名
//
// 创建防护域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) CreateDomain(request *model.CreateDomainRequest) (*model.CreateDomainResponse, error) {
	requestDef := GenReqDefForCreateDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDomainResponse), nil
	}
}

// CreateDomainInvoker 创建防护域名
func (c *AadClient) CreateDomainInvoker(request *model.CreateDomainRequest) *CreateDomainInvoker {
	requestDef := GenReqDefForCreateDomain()
	return &CreateDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDomain 删除防护域名
//
// 删除防护域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) DeleteDomain(request *model.DeleteDomainRequest) (*model.DeleteDomainResponse, error) {
	requestDef := GenReqDefForDeleteDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDomainResponse), nil
	}
}

// DeleteDomainInvoker 删除防护域名
func (c *AadClient) DeleteDomainInvoker(request *model.DeleteDomainRequest) *DeleteDomainInvoker {
	requestDef := GenReqDefForDeleteDomain()
	return &DeleteDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWafWhiteIpRule 防护策略web-cc黑白名单-删除黑白名单规则
//
// 防护策略web-cc黑白名单-删除黑白名单规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) DeleteWafWhiteIpRule(request *model.DeleteWafWhiteIpRuleRequest) (*model.DeleteWafWhiteIpRuleResponse, error) {
	requestDef := GenReqDefForDeleteWafWhiteIpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWafWhiteIpRuleResponse), nil
	}
}

// DeleteWafWhiteIpRuleInvoker 防护策略web-cc黑白名单-删除黑白名单规则
func (c *AadClient) DeleteWafWhiteIpRuleInvoker(request *model.DeleteWafWhiteIpRuleRequest) *DeleteWafWhiteIpRuleInvoker {
	requestDef := GenReqDefForDeleteWafWhiteIpRule()
	return &DeleteWafWhiteIpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDDoSAttackEvent 查询DDoS攻击事件列表
//
// 查询DDoS攻击事件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListDDoSAttackEvent(request *model.ListDDoSAttackEventRequest) (*model.ListDDoSAttackEventResponse, error) {
	requestDef := GenReqDefForListDDoSAttackEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDDoSAttackEventResponse), nil
	}
}

// ListDDoSAttackEventInvoker 查询DDoS攻击事件列表
func (c *AadClient) ListDDoSAttackEventInvoker(request *model.ListDDoSAttackEventRequest) *ListDDoSAttackEventInvoker {
	requestDef := GenReqDefForListDDoSAttackEvent()
	return &ListDDoSAttackEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDDoSBlackHoleEvent 黑洞事件列表
//
// 黑洞事件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListDDoSBlackHoleEvent(request *model.ListDDoSBlackHoleEventRequest) (*model.ListDDoSBlackHoleEventResponse, error) {
	requestDef := GenReqDefForListDDoSBlackHoleEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDDoSBlackHoleEventResponse), nil
	}
}

// ListDDoSBlackHoleEventInvoker 黑洞事件列表
func (c *AadClient) ListDDoSBlackHoleEventInvoker(request *model.ListDDoSBlackHoleEventRequest) *ListDDoSBlackHoleEventInvoker {
	requestDef := GenReqDefForListDDoSBlackHoleEvent()
	return &ListDDoSBlackHoleEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDDoSConnectionNumber 查询新建连接数和并发连接数
//
// 查询新建连接数和并发连接数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListDDoSConnectionNumber(request *model.ListDDoSConnectionNumberRequest) (*model.ListDDoSConnectionNumberResponse, error) {
	requestDef := GenReqDefForListDDoSConnectionNumber()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDDoSConnectionNumberResponse), nil
	}
}

// ListDDoSConnectionNumberInvoker 查询新建连接数和并发连接数
func (c *AadClient) ListDDoSConnectionNumberInvoker(request *model.ListDDoSConnectionNumberRequest) *ListDDoSConnectionNumberInvoker {
	requestDef := GenReqDefForListDDoSConnectionNumber()
	return &ListDDoSConnectionNumberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDDoSFlow 查询DDoS攻击防护BPS/PPS流量
//
// 查询DDoS攻击防护BPS/PPS流量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListDDoSFlow(request *model.ListDDoSFlowRequest) (*model.ListDDoSFlowResponse, error) {
	requestDef := GenReqDefForListDDoSFlow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDDoSFlowResponse), nil
	}
}

// ListDDoSFlowInvoker 查询DDoS攻击防护BPS/PPS流量
func (c *AadClient) ListDDoSFlowInvoker(request *model.ListDDoSFlowRequest) *ListDDoSFlowInvoker {
	requestDef := GenReqDefForListDDoSFlow()
	return &ListDDoSFlowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFrequencyControlRule 查询频率控制规则列表
//
// 查询频率控制规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListFrequencyControlRule(request *model.ListFrequencyControlRuleRequest) (*model.ListFrequencyControlRuleResponse, error) {
	requestDef := GenReqDefForListFrequencyControlRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFrequencyControlRuleResponse), nil
	}
}

// ListFrequencyControlRuleInvoker 查询频率控制规则列表
func (c *AadClient) ListFrequencyControlRuleInvoker(request *model.ListFrequencyControlRuleRequest) *ListFrequencyControlRuleInvoker {
	requestDef := GenReqDefForListFrequencyControlRule()
	return &ListFrequencyControlRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGlobalConfig 查询控制台WAF全局配置
//
// 查询控制台WAF全局配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListGlobalConfig(request *model.ListGlobalConfigRequest) (*model.ListGlobalConfigResponse, error) {
	requestDef := GenReqDefForListGlobalConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGlobalConfigResponse), nil
	}
}

// ListGlobalConfigInvoker 查询控制台WAF全局配置
func (c *AadClient) ListGlobalConfigInvoker(request *model.ListGlobalConfigRequest) *ListGlobalConfigInvoker {
	requestDef := GenReqDefForListGlobalConfig()
	return &ListGlobalConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceDomains 查询实例关联的域名信息
//
// 查询实例关联的域名信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListInstanceDomains(request *model.ListInstanceDomainsRequest) (*model.ListInstanceDomainsResponse, error) {
	requestDef := GenReqDefForListInstanceDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceDomainsResponse), nil
	}
}

// ListInstanceDomainsInvoker 查询实例关联的域名信息
func (c *AadClient) ListInstanceDomainsInvoker(request *model.ListInstanceDomainsRequest) *ListInstanceDomainsInvoker {
	requestDef := GenReqDefForListInstanceDomains()
	return &ListInstanceDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSourceIp 查询回源ip列表
//
// 查询回源ip列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListSourceIp(request *model.ListSourceIpRequest) (*model.ListSourceIpResponse, error) {
	requestDef := GenReqDefForListSourceIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSourceIpResponse), nil
	}
}

// ListSourceIpInvoker 查询回源ip列表
func (c *AadClient) ListSourceIpInvoker(request *model.ListSourceIpRequest) *ListSourceIpInvoker {
	requestDef := GenReqDefForListSourceIp()
	return &ListSourceIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWafAttackEvent 查询攻击事件列表
//
// 查询攻击事件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListWafAttackEvent(request *model.ListWafAttackEventRequest) (*model.ListWafAttackEventResponse, error) {
	requestDef := GenReqDefForListWafAttackEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWafAttackEventResponse), nil
	}
}

// ListWafAttackEventInvoker 查询攻击事件列表
func (c *AadClient) ListWafAttackEventInvoker(request *model.ListWafAttackEventRequest) *ListWafAttackEventInvoker {
	requestDef := GenReqDefForListWafAttackEvent()
	return &ListWafAttackEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWafBandwidth 带宽曲线
//
// 带宽曲线
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListWafBandwidth(request *model.ListWafBandwidthRequest) (*model.ListWafBandwidthResponse, error) {
	requestDef := GenReqDefForListWafBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWafBandwidthResponse), nil
	}
}

// ListWafBandwidthInvoker 带宽曲线
func (c *AadClient) ListWafBandwidthInvoker(request *model.ListWafBandwidthRequest) *ListWafBandwidthInvoker {
	requestDef := GenReqDefForListWafBandwidth()
	return &ListWafBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWafCustomRule 查询精准防护规则
//
// 查询精准防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListWafCustomRule(request *model.ListWafCustomRuleRequest) (*model.ListWafCustomRuleResponse, error) {
	requestDef := GenReqDefForListWafCustomRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWafCustomRuleResponse), nil
	}
}

// ListWafCustomRuleInvoker 查询精准防护规则
func (c *AadClient) ListWafCustomRuleInvoker(request *model.ListWafCustomRuleRequest) *ListWafCustomRuleInvoker {
	requestDef := GenReqDefForListWafCustomRule()
	return &ListWafCustomRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWafGeoIpRule 查询区域封禁规则
//
// 查询区域封禁规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListWafGeoIpRule(request *model.ListWafGeoIpRuleRequest) (*model.ListWafGeoIpRuleResponse, error) {
	requestDef := GenReqDefForListWafGeoIpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWafGeoIpRuleResponse), nil
	}
}

// ListWafGeoIpRuleInvoker 查询区域封禁规则
func (c *AadClient) ListWafGeoIpRuleInvoker(request *model.ListWafGeoIpRuleRequest) *ListWafGeoIpRuleInvoker {
	requestDef := GenReqDefForListWafGeoIpRule()
	return &ListWafGeoIpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWafQps 查询请求QPS
//
// 查询请求QPS
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListWafQps(request *model.ListWafQpsRequest) (*model.ListWafQpsResponse, error) {
	requestDef := GenReqDefForListWafQps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWafQpsResponse), nil
	}
}

// ListWafQpsInvoker 查询请求QPS
func (c *AadClient) ListWafQpsInvoker(request *model.ListWafQpsRequest) *ListWafQpsInvoker {
	requestDef := GenReqDefForListWafQps()
	return &ListWafQpsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWafWhiteIpRule 防护策略web-cc黑白名单-查询黑白名单规则
//
// 防护策略web-cc黑白名单-查询黑白名单规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListWafWhiteIpRule(request *model.ListWafWhiteIpRuleRequest) (*model.ListWafWhiteIpRuleResponse, error) {
	requestDef := GenReqDefForListWafWhiteIpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWafWhiteIpRuleResponse), nil
	}
}

// ListWafWhiteIpRuleInvoker 防护策略web-cc黑白名单-查询黑白名单规则
func (c *AadClient) ListWafWhiteIpRuleInvoker(request *model.ListWafWhiteIpRuleRequest) *ListWafWhiteIpRuleInvoker {
	requestDef := GenReqDefForListWafWhiteIpRule()
	return &ListWafWhiteIpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWhiteBlackIpRule 查询DDoS攻击防护的黑白名单列表
//
// 查询DDoS攻击防护的黑白名单列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ListWhiteBlackIpRule(request *model.ListWhiteBlackIpRuleRequest) (*model.ListWhiteBlackIpRuleResponse, error) {
	requestDef := GenReqDefForListWhiteBlackIpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWhiteBlackIpRuleResponse), nil
	}
}

// ListWhiteBlackIpRuleInvoker 查询DDoS攻击防护的黑白名单列表
func (c *AadClient) ListWhiteBlackIpRuleInvoker(request *model.ListWhiteBlackIpRuleRequest) *ListWhiteBlackIpRuleInvoker {
	requestDef := GenReqDefForListWhiteBlackIpRule()
	return &ListWhiteBlackIpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlarmConfig 查询告警设置
//
// 查询告警设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ShowAlarmConfig(request *model.ShowAlarmConfigRequest) (*model.ShowAlarmConfigResponse, error) {
	requestDef := GenReqDefForShowAlarmConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlarmConfigResponse), nil
	}
}

// ShowAlarmConfigInvoker 查询告警设置
func (c *AadClient) ShowAlarmConfigInvoker(request *model.ShowAlarmConfigRequest) *ShowAlarmConfigInvoker {
	requestDef := GenReqDefForShowAlarmConfig()
	return &ShowAlarmConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDDoSPeak 查询高防入流量峰值、攻击流量峰值、DDoS攻击次数
//
// 查询高防入流量峰值、攻击流量峰值、DDoS攻击次数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ShowDDoSPeak(request *model.ShowDDoSPeakRequest) (*model.ShowDDoSPeakResponse, error) {
	requestDef := GenReqDefForShowDDoSPeak()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDDoSPeakResponse), nil
	}
}

// ShowDDoSPeakInvoker 查询高防入流量峰值、攻击流量峰值、DDoS攻击次数
func (c *AadClient) ShowDDoSPeakInvoker(request *model.ShowDDoSPeakRequest) *ShowDDoSPeakInvoker {
	requestDef := GenReqDefForShowDDoSPeak()
	return &ShowDDoSPeakInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDomainCertificate 查询域名关联的证书信息
//
// 查询域名关联的证书信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ShowDomainCertificate(request *model.ShowDomainCertificateRequest) (*model.ShowDomainCertificateResponse, error) {
	requestDef := GenReqDefForShowDomainCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainCertificateResponse), nil
	}
}

// ShowDomainCertificateInvoker 查询域名关联的证书信息
func (c *AadClient) ShowDomainCertificateInvoker(request *model.ShowDomainCertificateRequest) *ShowDomainCertificateInvoker {
	requestDef := GenReqDefForShowDomainCertificate()
	return &ShowDomainCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDomainDetail 查询域名详情
//
// 查询域名详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ShowDomainDetail(request *model.ShowDomainDetailRequest) (*model.ShowDomainDetailResponse, error) {
	requestDef := GenReqDefForShowDomainDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainDetailResponse), nil
	}
}

// ShowDomainDetailInvoker 查询域名详情
func (c *AadClient) ShowDomainDetailInvoker(request *model.ShowDomainDetailRequest) *ShowDomainDetailInvoker {
	requestDef := GenReqDefForShowDomainDetail()
	return &ShowDomainDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDomainNameConfig 查看域名配置
//
// 查看域名配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ShowDomainNameConfig(request *model.ShowDomainNameConfigRequest) (*model.ShowDomainNameConfigResponse, error) {
	requestDef := GenReqDefForShowDomainNameConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainNameConfigResponse), nil
	}
}

// ShowDomainNameConfigInvoker 查看域名配置
func (c *AadClient) ShowDomainNameConfigInvoker(request *model.ShowDomainNameConfigRequest) *ShowDomainNameConfigInvoker {
	requestDef := GenReqDefForShowDomainNameConfig()
	return &ShowDomainNameConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlowBlock 查询流量封禁信息
//
// 查询流量封禁信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ShowFlowBlock(request *model.ShowFlowBlockRequest) (*model.ShowFlowBlockResponse, error) {
	requestDef := GenReqDefForShowFlowBlock()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlowBlockResponse), nil
	}
}

// ShowFlowBlockInvoker 查询流量封禁信息
func (c *AadClient) ShowFlowBlockInvoker(request *model.ShowFlowBlockRequest) *ShowFlowBlockInvoker {
	requestDef := GenReqDefForShowFlowBlock()
	return &ShowFlowBlockInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceByInstanceId 查询实例详情
//
// 查询实例详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ShowInstanceByInstanceId(request *model.ShowInstanceByInstanceIdRequest) (*model.ShowInstanceByInstanceIdResponse, error) {
	requestDef := GenReqDefForShowInstanceByInstanceId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceByInstanceIdResponse), nil
	}
}

// ShowInstanceByInstanceIdInvoker 查询实例详情
func (c *AadClient) ShowInstanceByInstanceIdInvoker(request *model.ShowInstanceByInstanceIdRequest) *ShowInstanceByInstanceIdInvoker {
	requestDef := GenReqDefForShowInstanceByInstanceId()
	return &ShowInstanceByInstanceIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWafPolicy 查询WEB防护策略配置
//
// 查询WEB防护策略配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ShowWafPolicy(request *model.ShowWafPolicyRequest) (*model.ShowWafPolicyResponse, error) {
	requestDef := GenReqDefForShowWafPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWafPolicyResponse), nil
	}
}

// ShowWafPolicyInvoker 查询WEB防护策略配置
func (c *AadClient) ShowWafPolicyInvoker(request *model.ShowWafPolicyRequest) *ShowWafPolicyInvoker {
	requestDef := GenReqDefForShowWafPolicy()
	return &ShowWafPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWafQps 查询CC攻击防护请求QPS
//
// 查询CC攻击防护请求QPS
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) ShowWafQps(request *model.ShowWafQpsRequest) (*model.ShowWafQpsResponse, error) {
	requestDef := GenReqDefForShowWafQps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWafQpsResponse), nil
	}
}

// ShowWafQpsInvoker 查询CC攻击防护请求QPS
func (c *AadClient) ShowWafQpsInvoker(request *model.ShowWafQpsRequest) *ShowWafQpsInvoker {
	requestDef := GenReqDefForShowWafQps()
	return &ShowWafQpsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDomainConfig 修改域名配置
//
// 修改域名配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) UpdateDomainConfig(request *model.UpdateDomainConfigRequest) (*model.UpdateDomainConfigResponse, error) {
	requestDef := GenReqDefForUpdateDomainConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainConfigResponse), nil
	}
}

// UpdateDomainConfigInvoker 修改域名配置
func (c *AadClient) UpdateDomainConfigInvoker(request *model.UpdateDomainConfigRequest) *UpdateDomainConfigInvoker {
	requestDef := GenReqDefForUpdateDomainConfig()
	return &UpdateDomainConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateForwardRule 修改转发规则中的源站IP
//
// 修改转发规则中的源站IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) UpdateForwardRule(request *model.UpdateForwardRuleRequest) (*model.UpdateForwardRuleResponse, error) {
	requestDef := GenReqDefForUpdateForwardRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateForwardRuleResponse), nil
	}
}

// UpdateForwardRuleInvoker 修改转发规则中的源站IP
func (c *AadClient) UpdateForwardRuleInvoker(request *model.UpdateForwardRuleRequest) *UpdateForwardRuleInvoker {
	requestDef := GenReqDefForUpdateForwardRule()
	return &UpdateForwardRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpgradeInstanceSpec 修改实例规格
//
// 修改实例规格
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AadClient) UpgradeInstanceSpec(request *model.UpgradeInstanceSpecRequest) (*model.UpgradeInstanceSpecResponse, error) {
	requestDef := GenReqDefForUpgradeInstanceSpec()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeInstanceSpecResponse), nil
	}
}

// UpgradeInstanceSpecInvoker 修改实例规格
func (c *AadClient) UpgradeInstanceSpecInvoker(request *model.UpgradeInstanceSpecRequest) *UpgradeInstanceSpecInvoker {
	requestDef := GenReqDefForUpgradeInstanceSpec()
	return &UpgradeInstanceSpecInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
