package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/waf/v1/model"
)

type WafClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewWafClient(hcClient *httpclient.HcHttpClient) *WafClient {
	return &WafClient{HcClient: hcClient}
}

func WafClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// ApplyCertificateToHost 绑定证书到域名
//
// 绑定证书到域名
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ChangePrepaidCloudWaf 变更包周期云模式waf规格
//
// 变更包周期云模式waf规格。注：
//   - 1.变更某产品规格的前提是必须已购买该产品
//   - 2.waf版本只支持升配，不支持降配；扩展包数量可以增加或者减少，但不支持数量减少为0
//   - 3.不支持同时升降配，如增加域名扩展包数量，同时减少规则扩展包数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ChangePrepaidCloudWaf(request *model.ChangePrepaidCloudWafRequest) (*model.ChangePrepaidCloudWafResponse, error) {
	requestDef := GenReqDefForChangePrepaidCloudWaf()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangePrepaidCloudWafResponse), nil
	}
}

// ChangePrepaidCloudWafInvoker 变更包周期云模式waf规格
func (c *WafClient) ChangePrepaidCloudWafInvoker(request *model.ChangePrepaidCloudWafRequest) *ChangePrepaidCloudWafInvoker {
	requestDef := GenReqDefForChangePrepaidCloudWaf()
	return &ChangePrepaidCloudWafInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAntiTamperRule 创建防篡改规则
//
// 创建防篡改规则
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// CreateAnticrawlerRule 创建JS脚本反爬虫规则
//
// 创建JS脚本反爬虫规则，在调用此接口创建防护规则前，需要调用更新JS脚本反爬虫规则防护模式（UpdateAnticrawlerRuleType）接口指定防护模式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) CreateAnticrawlerRule(request *model.CreateAnticrawlerRuleRequest) (*model.CreateAnticrawlerRuleResponse, error) {
	requestDef := GenReqDefForCreateAnticrawlerRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAnticrawlerRuleResponse), nil
	}
}

// CreateAnticrawlerRuleInvoker 创建JS脚本反爬虫规则
func (c *WafClient) CreateAnticrawlerRuleInvoker(request *model.CreateAnticrawlerRuleRequest) *CreateAnticrawlerRuleInvoker {
	requestDef := GenReqDefForCreateAnticrawlerRule()
	return &CreateAnticrawlerRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAntileakageRule 创建防敏感信息泄露规则
//
// 创建防敏感信息泄露规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) CreateAntileakageRule(request *model.CreateAntileakageRuleRequest) (*model.CreateAntileakageRuleResponse, error) {
	requestDef := GenReqDefForCreateAntileakageRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAntileakageRuleResponse), nil
	}
}

// CreateAntileakageRuleInvoker 创建防敏感信息泄露规则
func (c *WafClient) CreateAntileakageRuleInvoker(request *model.CreateAntileakageRuleRequest) *CreateAntileakageRuleInvoker {
	requestDef := GenReqDefForCreateAntileakageRule()
	return &CreateAntileakageRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCcRule 创建cc规则
//
// 创建cc规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) CreateCcRule(request *model.CreateCcRuleRequest) (*model.CreateCcRuleResponse, error) {
	requestDef := GenReqDefForCreateCcRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCcRuleResponse), nil
	}
}

// CreateCcRuleInvoker 创建cc规则
func (c *WafClient) CreateCcRuleInvoker(request *model.CreateCcRuleRequest) *CreateCcRuleInvoker {
	requestDef := GenReqDefForCreateCcRule()
	return &CreateCcRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCertificate 创建证书
//
// 创建证书
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// CreateCloudWafPostPaidResource 开通云模式按需计费接口
//
// 开通云模式按需计费接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) CreateCloudWafPostPaidResource(request *model.CreateCloudWafPostPaidResourceRequest) (*model.CreateCloudWafPostPaidResourceResponse, error) {
	requestDef := GenReqDefForCreateCloudWafPostPaidResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCloudWafPostPaidResourceResponse), nil
	}
}

// CreateCloudWafPostPaidResourceInvoker 开通云模式按需计费接口
func (c *WafClient) CreateCloudWafPostPaidResourceInvoker(request *model.CreateCloudWafPostPaidResourceRequest) *CreateCloudWafPostPaidResourceInvoker {
	requestDef := GenReqDefForCreateCloudWafPostPaidResource()
	return &CreateCloudWafPostPaidResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCustomRule 创建精准防护规则
//
// 创建精准防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) CreateCustomRule(request *model.CreateCustomRuleRequest) (*model.CreateCustomRuleResponse, error) {
	requestDef := GenReqDefForCreateCustomRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCustomRuleResponse), nil
	}
}

// CreateCustomRuleInvoker 创建精准防护规则
func (c *WafClient) CreateCustomRuleInvoker(request *model.CreateCustomRuleRequest) *CreateCustomRuleInvoker {
	requestDef := GenReqDefForCreateCustomRule()
	return &CreateCustomRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGeoipRule 创建地理位置控制规则
//
// 创建地理位置控制规则
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// CreateInstance 创建WAF独享引擎实例
//
// 创建WAF独享引擎实例。独享模式支持的局点包括：华东-青岛、中东-利雅得、华北-北京一、华北-北京四、华北-乌兰察布一、华东-上海一、华东-上海二、华南-广州、华南-深圳、中国-香港、西南-贵阳一、亚太-曼谷、 亚太-新加坡、非洲约翰内斯堡、土耳其-伊斯坦布尔；普通租户类独享支持的局点：华北-北京四、华东-上海一、华南-广州、中国-香港、亚太-曼谷、 亚太-新加坡。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

// CreateInstanceInvoker 创建WAF独享引擎实例
func (c *WafClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateIpGroup 创建ip地址组
//
// 创建ip地址组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) CreateIpGroup(request *model.CreateIpGroupRequest) (*model.CreateIpGroupResponse, error) {
	requestDef := GenReqDefForCreateIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIpGroupResponse), nil
	}
}

// CreateIpGroupInvoker 创建ip地址组
func (c *WafClient) CreateIpGroupInvoker(request *model.CreateIpGroupRequest) *CreateIpGroupInvoker {
	requestDef := GenReqDefForCreateIpGroup()
	return &CreateIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePolicy 创建防护策略
//
// 创建防护策略，系统会在生成策略时配置一些默认的配置项，如果需要修改策略的默认配置项需要通过调用更新防护策略接口实现
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// CreatePrepaidCloudWaf 购买包周期云模式waf
//
// 购买包周期云模式waf。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) CreatePrepaidCloudWaf(request *model.CreatePrepaidCloudWafRequest) (*model.CreatePrepaidCloudWafResponse, error) {
	requestDef := GenReqDefForCreatePrepaidCloudWaf()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrepaidCloudWafResponse), nil
	}
}

// CreatePrepaidCloudWafInvoker 购买包周期云模式waf
func (c *WafClient) CreatePrepaidCloudWafInvoker(request *model.CreatePrepaidCloudWafRequest) *CreatePrepaidCloudWafInvoker {
	requestDef := GenReqDefForCreatePrepaidCloudWaf()
	return &CreatePrepaidCloudWafInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePrivacyRule 创建隐私屏蔽防护规则
//
// 创建隐私屏蔽防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// CreatePunishmentRule 创建攻击惩罚规则
//
// 创建攻击惩罚规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) CreatePunishmentRule(request *model.CreatePunishmentRuleRequest) (*model.CreatePunishmentRuleResponse, error) {
	requestDef := GenReqDefForCreatePunishmentRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePunishmentRuleResponse), nil
	}
}

// CreatePunishmentRuleInvoker 创建攻击惩罚规则
func (c *WafClient) CreatePunishmentRuleInvoker(request *model.CreatePunishmentRuleRequest) *CreatePunishmentRuleInvoker {
	requestDef := GenReqDefForCreatePunishmentRule()
	return &CreatePunishmentRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateValueList 创建引用表
//
// 创建引用表，引用表能够被CC攻击防护规则和精准访问防护中的规则所引用。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// DeleteAnticrawlerRule 删除JS脚本反爬虫防护规则
//
// 删除JS脚本反爬虫防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) DeleteAnticrawlerRule(request *model.DeleteAnticrawlerRuleRequest) (*model.DeleteAnticrawlerRuleResponse, error) {
	requestDef := GenReqDefForDeleteAnticrawlerRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAnticrawlerRuleResponse), nil
	}
}

// DeleteAnticrawlerRuleInvoker 删除JS脚本反爬虫防护规则
func (c *WafClient) DeleteAnticrawlerRuleInvoker(request *model.DeleteAnticrawlerRuleRequest) *DeleteAnticrawlerRuleInvoker {
	requestDef := GenReqDefForDeleteAnticrawlerRule()
	return &DeleteAnticrawlerRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAntileakageRule 删除防敏感信息泄露防护规则
//
// 删除防敏感信息泄露防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) DeleteAntileakageRule(request *model.DeleteAntileakageRuleRequest) (*model.DeleteAntileakageRuleResponse, error) {
	requestDef := GenReqDefForDeleteAntileakageRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAntileakageRuleResponse), nil
	}
}

// DeleteAntileakageRuleInvoker 删除防敏感信息泄露防护规则
func (c *WafClient) DeleteAntileakageRuleInvoker(request *model.DeleteAntileakageRuleRequest) *DeleteAntileakageRuleInvoker {
	requestDef := GenReqDefForDeleteAntileakageRule()
	return &DeleteAntileakageRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAntitamperRule 删除防篡改防护规则
//
// 删除防篡改防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// DeleteCcRule 删除cc防护规则
//
// 删除cc防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) DeleteCcRule(request *model.DeleteCcRuleRequest) (*model.DeleteCcRuleResponse, error) {
	requestDef := GenReqDefForDeleteCcRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCcRuleResponse), nil
	}
}

// DeleteCcRuleInvoker 删除cc防护规则
func (c *WafClient) DeleteCcRuleInvoker(request *model.DeleteCcRuleRequest) *DeleteCcRuleInvoker {
	requestDef := GenReqDefForDeleteCcRule()
	return &DeleteCcRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCertificate 删除证书
//
// 删除证书
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// DeleteCloudWafPostPaidResource 关闭云模式按需计费接口
//
// 关闭云模式按需计费接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) DeleteCloudWafPostPaidResource(request *model.DeleteCloudWafPostPaidResourceRequest) (*model.DeleteCloudWafPostPaidResourceResponse, error) {
	requestDef := GenReqDefForDeleteCloudWafPostPaidResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCloudWafPostPaidResourceResponse), nil
	}
}

// DeleteCloudWafPostPaidResourceInvoker 关闭云模式按需计费接口
func (c *WafClient) DeleteCloudWafPostPaidResourceInvoker(request *model.DeleteCloudWafPostPaidResourceRequest) *DeleteCloudWafPostPaidResourceInvoker {
	requestDef := GenReqDefForDeleteCloudWafPostPaidResource()
	return &DeleteCloudWafPostPaidResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCustomRule 删除精准防护规则
//
// 删除精准防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) DeleteCustomRule(request *model.DeleteCustomRuleRequest) (*model.DeleteCustomRuleResponse, error) {
	requestDef := GenReqDefForDeleteCustomRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCustomRuleResponse), nil
	}
}

// DeleteCustomRuleInvoker 删除精准防护规则
func (c *WafClient) DeleteCustomRuleInvoker(request *model.DeleteCustomRuleRequest) *DeleteCustomRuleInvoker {
	requestDef := GenReqDefForDeleteCustomRule()
	return &DeleteCustomRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGeoipRule 删除地理位置控制防护规则
//
// 删除地理位置控制防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// DeleteInstance 删除WAF独享引擎信息
//
// 删除WAF独享引擎信息。独享模式只在部分局点支持，包括：华北-北京四、华东-上海一、华南-广州、华南-深圳  、中国-香港、亚太-曼谷、 亚太-新加坡。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

// DeleteInstanceInvoker 删除WAF独享引擎信息
func (c *WafClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteIpGroup 删除ip地址组
//
// 删除ip地址组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) DeleteIpGroup(request *model.DeleteIpGroupRequest) (*model.DeleteIpGroupResponse, error) {
	requestDef := GenReqDefForDeleteIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIpGroupResponse), nil
	}
}

// DeleteIpGroupInvoker 删除ip地址组
func (c *WafClient) DeleteIpGroupInvoker(request *model.DeleteIpGroupRequest) *DeleteIpGroupInvoker {
	requestDef := GenReqDefForDeleteIpGroup()
	return &DeleteIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePolicy 删除防护策略
//
// 删除防护策略，若策略正在使用，则需要先接解除域名与策略的绑定关系才能删除策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// DeletePunishmentRule 删除攻击惩罚规则
//
// 删除攻击惩罚规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) DeletePunishmentRule(request *model.DeletePunishmentRuleRequest) (*model.DeletePunishmentRuleResponse, error) {
	requestDef := GenReqDefForDeletePunishmentRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePunishmentRuleResponse), nil
	}
}

// DeletePunishmentRuleInvoker 删除攻击惩罚规则
func (c *WafClient) DeletePunishmentRuleInvoker(request *model.DeletePunishmentRuleRequest) *DeletePunishmentRuleInvoker {
	requestDef := GenReqDefForDeletePunishmentRule()
	return &DeletePunishmentRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteValueList 删除引用表
//
// 删除引用表
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListAnticrawlerRules 查询JS脚本反爬虫规则列表
//
// 查询JS脚本反爬虫规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ListAnticrawlerRules(request *model.ListAnticrawlerRulesRequest) (*model.ListAnticrawlerRulesResponse, error) {
	requestDef := GenReqDefForListAnticrawlerRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAnticrawlerRulesResponse), nil
	}
}

// ListAnticrawlerRulesInvoker 查询JS脚本反爬虫规则列表
func (c *WafClient) ListAnticrawlerRulesInvoker(request *model.ListAnticrawlerRulesRequest) *ListAnticrawlerRulesInvoker {
	requestDef := GenReqDefForListAnticrawlerRules()
	return &ListAnticrawlerRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAntileakageRules 查询防敏感信息泄露规则列表
//
// 查询防敏感信息泄露规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ListAntileakageRules(request *model.ListAntileakageRulesRequest) (*model.ListAntileakageRulesResponse, error) {
	requestDef := GenReqDefForListAntileakageRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAntileakageRulesResponse), nil
	}
}

// ListAntileakageRulesInvoker 查询防敏感信息泄露规则列表
func (c *WafClient) ListAntileakageRulesInvoker(request *model.ListAntileakageRulesRequest) *ListAntileakageRulesInvoker {
	requestDef := GenReqDefForListAntileakageRules()
	return &ListAntileakageRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAntitamperRule 查询防篡改规则列表
//
// 查询防篡改规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// 查询安全统计带宽数据，统计的带宽数据为平均值，单位为bit/s。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListCcRules 查询cc规则列表
//
// 查询cc规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ListCcRules(request *model.ListCcRulesRequest) (*model.ListCcRulesResponse, error) {
	requestDef := GenReqDefForListCcRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCcRulesResponse), nil
	}
}

// ListCcRulesInvoker 查询cc规则列表
func (c *WafClient) ListCcRulesInvoker(request *model.ListCcRulesRequest) *ListCcRulesInvoker {
	requestDef := GenReqDefForListCcRules()
	return &ListCcRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCertificates 查询证书列表
//
// 查询证书列表
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListCustomRules 查询精准防护规则列表
//
// 查询精准防护规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ListCustomRules(request *model.ListCustomRulesRequest) (*model.ListCustomRulesResponse, error) {
	requestDef := GenReqDefForListCustomRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomRulesResponse), nil
	}
}

// ListCustomRulesInvoker 查询精准防护规则列表
func (c *WafClient) ListCustomRulesInvoker(request *model.ListCustomRulesRequest) *ListCustomRulesInvoker {
	requestDef := GenReqDefForListCustomRules()
	return &ListCustomRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEvent 查询攻击事件列表
//
// 查询攻击事件列表，该API暂时不支持查询全部防护事件，pagesize参数不可设为-1，由于性能原因，数据量越大消耗的内存越大，后端最多限制查询10000条数据，例如：自定义时间段内的数据超过了10000条，就无法查出page为101，pagesize为100之后的数据，需要调整时间区间，再进行查询
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListInstance 查询WAF独享引擎列表
//
// 查询WAF独享引擎列表。独享模式只在部分局点支持，包括：华北-北京四、华东-上海一、华南-广州、华南-深圳  、中国-香港、亚太-曼谷、 亚太-新加坡。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ListInstance(request *model.ListInstanceRequest) (*model.ListInstanceResponse, error) {
	requestDef := GenReqDefForListInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceResponse), nil
	}
}

// ListInstanceInvoker 查询WAF独享引擎列表
func (c *WafClient) ListInstanceInvoker(request *model.ListInstanceRequest) *ListInstanceInvoker {
	requestDef := GenReqDefForListInstance()
	return &ListInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIpGroup 查询地址组列表
//
// 查询地址组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ListIpGroup(request *model.ListIpGroupRequest) (*model.ListIpGroupResponse, error) {
	requestDef := GenReqDefForListIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIpGroupResponse), nil
	}
}

// ListIpGroupInvoker 查询地址组列表
func (c *WafClient) ListIpGroupInvoker(request *model.ListIpGroupRequest) *ListIpGroupInvoker {
	requestDef := GenReqDefForListIpGroup()
	return &ListIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNoticeConfigs 查询告警通知配置
//
// 查询告警通知配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ListNoticeConfigs(request *model.ListNoticeConfigsRequest) (*model.ListNoticeConfigsResponse, error) {
	requestDef := GenReqDefForListNoticeConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNoticeConfigsResponse), nil
	}
}

// ListNoticeConfigsInvoker 查询告警通知配置
func (c *WafClient) ListNoticeConfigsInvoker(request *model.ListNoticeConfigsRequest) *ListNoticeConfigsInvoker {
	requestDef := GenReqDefForListNoticeConfigs()
	return &ListNoticeConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOverviewsClassification 查询安全总览分类统计top信息
//
// 查询安全总览分类统计TOP信息，包含受攻击域名 、攻击源ip、受攻击URL、攻击来源区域、攻击事件分布。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListPrivacyRule 查询隐私屏蔽防护规则列表
//
// 查询隐私屏蔽防护规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ListPrivacyRule(request *model.ListPrivacyRuleRequest) (*model.ListPrivacyRuleResponse, error) {
	requestDef := GenReqDefForListPrivacyRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivacyRuleResponse), nil
	}
}

// ListPrivacyRuleInvoker 查询隐私屏蔽防护规则列表
func (c *WafClient) ListPrivacyRuleInvoker(request *model.ListPrivacyRuleRequest) *ListPrivacyRuleInvoker {
	requestDef := GenReqDefForListPrivacyRule()
	return &ListPrivacyRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPunishmentRules 查询攻击惩罚规则列表
//
// 查询攻击惩罚规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ListPunishmentRules(request *model.ListPunishmentRulesRequest) (*model.ListPunishmentRulesResponse, error) {
	requestDef := GenReqDefForListPunishmentRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPunishmentRulesResponse), nil
	}
}

// ListPunishmentRulesInvoker 查询攻击惩罚规则列表
func (c *WafClient) ListPunishmentRulesInvoker(request *model.ListPunishmentRulesRequest) *ListPunishmentRulesInvoker {
	requestDef := GenReqDefForListPunishmentRules()
	return &ListPunishmentRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQpsTimeline 查询安全统计qps次数
//
// 查询安全统计qps次数。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListRequestTimeline 查询安全总览中请求次数时间线统计数据
//
// 查询安全总览中请求次数时间线统计数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ListRequestTimeline(request *model.ListRequestTimelineRequest) (*model.ListRequestTimelineResponse, error) {
	requestDef := GenReqDefForListRequestTimeline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRequestTimelineResponse), nil
	}
}

// ListRequestTimelineInvoker 查询安全总览中请求次数时间线统计数据
func (c *WafClient) ListRequestTimelineInvoker(request *model.ListRequestTimelineRequest) *ListRequestTimelineInvoker {
	requestDef := GenReqDefForListRequestTimeline()
	return &ListRequestTimelineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStatistics 查询安全总览请求与攻击数量
//
// 查询安全总览请求与攻击数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// 查询业务异常TOP统计信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// MigrateCompositeHosts 按企业项目迁移防护域名
//
// 按企业项目迁移防护域名，仅专业版与独享版支持该功能
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) MigrateCompositeHosts(request *model.MigrateCompositeHostsRequest) (*model.MigrateCompositeHostsResponse, error) {
	requestDef := GenReqDefForMigrateCompositeHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MigrateCompositeHostsResponse), nil
	}
}

// MigrateCompositeHostsInvoker 按企业项目迁移防护域名
func (c *WafClient) MigrateCompositeHostsInvoker(request *model.MigrateCompositeHostsRequest) *MigrateCompositeHostsInvoker {
	requestDef := GenReqDefForMigrateCompositeHosts()
	return &MigrateCompositeHostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RenameInstance 重命名WAF独享引擎
//
// 重命名WAF独享引擎。独享模式只在部分局点支持，包括：华北-北京四、华东-上海一、华南-广州、华南-深圳  、中国-香港、亚太-曼谷、 亚太-新加坡。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) RenameInstance(request *model.RenameInstanceRequest) (*model.RenameInstanceResponse, error) {
	requestDef := GenReqDefForRenameInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RenameInstanceResponse), nil
	}
}

// RenameInstanceInvoker 重命名WAF独享引擎
func (c *WafClient) RenameInstanceInvoker(request *model.RenameInstanceRequest) *RenameInstanceInvoker {
	requestDef := GenReqDefForRenameInstance()
	return &RenameInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAnticrawlerRule 查询JS脚本反爬虫防护规则
//
// 根据Id查询JS脚本反爬虫防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ShowAnticrawlerRule(request *model.ShowAnticrawlerRuleRequest) (*model.ShowAnticrawlerRuleResponse, error) {
	requestDef := GenReqDefForShowAnticrawlerRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAnticrawlerRuleResponse), nil
	}
}

// ShowAnticrawlerRuleInvoker 查询JS脚本反爬虫防护规则
func (c *WafClient) ShowAnticrawlerRuleInvoker(request *model.ShowAnticrawlerRuleRequest) *ShowAnticrawlerRuleInvoker {
	requestDef := GenReqDefForShowAnticrawlerRule()
	return &ShowAnticrawlerRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAntileakageRule 查询防敏感信息泄露防护规则
//
// 根据Id查询防敏感信息泄露防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ShowAntileakageRule(request *model.ShowAntileakageRuleRequest) (*model.ShowAntileakageRuleResponse, error) {
	requestDef := GenReqDefForShowAntileakageRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAntileakageRuleResponse), nil
	}
}

// ShowAntileakageRuleInvoker 查询防敏感信息泄露防护规则
func (c *WafClient) ShowAntileakageRuleInvoker(request *model.ShowAntileakageRuleRequest) *ShowAntileakageRuleInvoker {
	requestDef := GenReqDefForShowAntileakageRule()
	return &ShowAntileakageRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAntitamperRule 查询防篡改防护规则
//
// 查询防篡改防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ShowAntitamperRule(request *model.ShowAntitamperRuleRequest) (*model.ShowAntitamperRuleResponse, error) {
	requestDef := GenReqDefForShowAntitamperRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAntitamperRuleResponse), nil
	}
}

// ShowAntitamperRuleInvoker 查询防篡改防护规则
func (c *WafClient) ShowAntitamperRuleInvoker(request *model.ShowAntitamperRuleRequest) *ShowAntitamperRuleInvoker {
	requestDef := GenReqDefForShowAntitamperRule()
	return &ShowAntitamperRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCcRule 根据Id查询cc防护规则
//
// 根据Id查询cc防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ShowCcRule(request *model.ShowCcRuleRequest) (*model.ShowCcRuleResponse, error) {
	requestDef := GenReqDefForShowCcRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCcRuleResponse), nil
	}
}

// ShowCcRuleInvoker 根据Id查询cc防护规则
func (c *WafClient) ShowCcRuleInvoker(request *model.ShowCcRuleRequest) *ShowCcRuleInvoker {
	requestDef := GenReqDefForShowCcRule()
	return &ShowCcRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCertificate 查询证书
//
// 查询证书
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowCustomRule 根据Id查询精准防护规则
//
// 根据Id查询精准防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ShowCustomRule(request *model.ShowCustomRuleRequest) (*model.ShowCustomRuleResponse, error) {
	requestDef := GenReqDefForShowCustomRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCustomRuleResponse), nil
	}
}

// ShowCustomRuleInvoker 根据Id查询精准防护规则
func (c *WafClient) ShowCustomRuleInvoker(request *model.ShowCustomRuleRequest) *ShowCustomRuleInvoker {
	requestDef := GenReqDefForShowCustomRule()
	return &ShowCustomRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEvent 查询指定事件id的防护事件详情
//
// 查询指定事件id的防护事件详情
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowGeoipRule 删除地理位置控制防护规则
//
// 删除地理位置控制防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ShowGeoipRule(request *model.ShowGeoipRuleRequest) (*model.ShowGeoipRuleResponse, error) {
	requestDef := GenReqDefForShowGeoipRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGeoipRuleResponse), nil
	}
}

// ShowGeoipRuleInvoker 删除地理位置控制防护规则
func (c *WafClient) ShowGeoipRuleInvoker(request *model.ShowGeoipRuleRequest) *ShowGeoipRuleInvoker {
	requestDef := GenReqDefForShowGeoipRule()
	return &ShowGeoipRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHost 根据防护域名Id查询云模式防护域名详细信息
//
// 根据防护域名Id查询云模式防护域名详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowIgnoreRule 查询全局白名单(原误报屏蔽)防护规则
//
// 查询全局白名单(原误报屏蔽)防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ShowIgnoreRule(request *model.ShowIgnoreRuleRequest) (*model.ShowIgnoreRuleResponse, error) {
	requestDef := GenReqDefForShowIgnoreRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIgnoreRuleResponse), nil
	}
}

// ShowIgnoreRuleInvoker 查询全局白名单(原误报屏蔽)防护规则
func (c *WafClient) ShowIgnoreRuleInvoker(request *model.ShowIgnoreRuleRequest) *ShowIgnoreRuleInvoker {
	requestDef := GenReqDefForShowIgnoreRule()
	return &ShowIgnoreRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstance 查询WAF独享引擎信息
//
// 查询WAF独享引擎信息。独享模式只在部分局点支持，包括：华北-北京四、华东-上海一、华南-广州、华南-深圳  、中国-香港、亚太-曼谷、 亚太-新加坡。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

// ShowInstanceInvoker 查询WAF独享引擎信息
func (c *WafClient) ShowInstanceInvoker(request *model.ShowInstanceRequest) *ShowInstanceInvoker {
	requestDef := GenReqDefForShowInstance()
	return &ShowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIpGroup 查询ip地址组明细
//
// 查询ip地址组明细
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ShowIpGroup(request *model.ShowIpGroupRequest) (*model.ShowIpGroupResponse, error) {
	requestDef := GenReqDefForShowIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIpGroupResponse), nil
	}
}

// ShowIpGroupInvoker 查询ip地址组明细
func (c *WafClient) ShowIpGroupInvoker(request *model.ShowIpGroupRequest) *ShowIpGroupInvoker {
	requestDef := GenReqDefForShowIpGroup()
	return &ShowIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLtsInfoConfig 查询lts配置信息
//
// 查询lts配置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ShowLtsInfoConfig(request *model.ShowLtsInfoConfigRequest) (*model.ShowLtsInfoConfigResponse, error) {
	requestDef := GenReqDefForShowLtsInfoConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLtsInfoConfigResponse), nil
	}
}

// ShowLtsInfoConfigInvoker 查询lts配置信息
func (c *WafClient) ShowLtsInfoConfigInvoker(request *model.ShowLtsInfoConfigRequest) *ShowLtsInfoConfigInvoker {
	requestDef := GenReqDefForShowLtsInfoConfig()
	return &ShowLtsInfoConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPolicy 根据Id查询防护策略
//
// 根据Id查询防护策略
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowPrivacyRule 查询隐私屏蔽防护规则
//
// 删除隐私屏蔽防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ShowPrivacyRule(request *model.ShowPrivacyRuleRequest) (*model.ShowPrivacyRuleResponse, error) {
	requestDef := GenReqDefForShowPrivacyRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivacyRuleResponse), nil
	}
}

// ShowPrivacyRuleInvoker 查询隐私屏蔽防护规则
func (c *WafClient) ShowPrivacyRuleInvoker(request *model.ShowPrivacyRuleRequest) *ShowPrivacyRuleInvoker {
	requestDef := GenReqDefForShowPrivacyRule()
	return &ShowPrivacyRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPunishmentRule 根据Id查询攻击惩罚防护规则
//
// 根据Id查询攻击惩罚防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ShowPunishmentRule(request *model.ShowPunishmentRuleRequest) (*model.ShowPunishmentRuleResponse, error) {
	requestDef := GenReqDefForShowPunishmentRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPunishmentRuleResponse), nil
	}
}

// ShowPunishmentRuleInvoker 根据Id查询攻击惩罚防护规则
func (c *WafClient) ShowPunishmentRuleInvoker(request *model.ShowPunishmentRuleRequest) *ShowPunishmentRuleInvoker {
	requestDef := GenReqDefForShowPunishmentRule()
	return &ShowPunishmentRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSourceIp 查询WAF回源Ip信息
//
// 查询WAF回源Ip信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ShowSourceIp(request *model.ShowSourceIpRequest) (*model.ShowSourceIpResponse, error) {
	requestDef := GenReqDefForShowSourceIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSourceIpResponse), nil
	}
}

// ShowSourceIpInvoker 查询WAF回源Ip信息
func (c *WafClient) ShowSourceIpInvoker(request *model.ShowSourceIpRequest) *ShowSourceIpInvoker {
	requestDef := GenReqDefForShowSourceIp()
	return &ShowSourceIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSubscriptionInfo 查询租户订购信息
//
// 查询租户订购信息，包括云模式包周期、按需计费、独享模式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ShowSubscriptionInfo(request *model.ShowSubscriptionInfoRequest) (*model.ShowSubscriptionInfoResponse, error) {
	requestDef := GenReqDefForShowSubscriptionInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSubscriptionInfoResponse), nil
	}
}

// ShowSubscriptionInfoInvoker 查询租户订购信息
func (c *WafClient) ShowSubscriptionInfoInvoker(request *model.ShowSubscriptionInfoRequest) *ShowSubscriptionInfoInvoker {
	requestDef := GenReqDefForShowSubscriptionInfo()
	return &ShowSubscriptionInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowValueList 查询引用表
//
// 查询引用表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ShowValueList(request *model.ShowValueListRequest) (*model.ShowValueListResponse, error) {
	requestDef := GenReqDefForShowValueList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowValueListResponse), nil
	}
}

// ShowValueListInvoker 查询引用表
func (c *WafClient) ShowValueListInvoker(request *model.ShowValueListRequest) *ShowValueListInvoker {
	requestDef := GenReqDefForShowValueList()
	return &ShowValueListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWhiteBlackIpRule 查询黑白名单防护规则
//
// 查询黑白名单防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) ShowWhiteBlackIpRule(request *model.ShowWhiteBlackIpRuleRequest) (*model.ShowWhiteBlackIpRuleResponse, error) {
	requestDef := GenReqDefForShowWhiteBlackIpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWhiteBlackIpRuleResponse), nil
	}
}

// ShowWhiteBlackIpRuleInvoker 查询黑白名单防护规则
func (c *WafClient) ShowWhiteBlackIpRuleInvoker(request *model.ShowWhiteBlackIpRuleRequest) *ShowWhiteBlackIpRuleInvoker {
	requestDef := GenReqDefForShowWhiteBlackIpRule()
	return &ShowWhiteBlackIpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAlertNoticeConfig 更新告警通知配置
//
// 更新告警通知配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) UpdateAlertNoticeConfig(request *model.UpdateAlertNoticeConfigRequest) (*model.UpdateAlertNoticeConfigResponse, error) {
	requestDef := GenReqDefForUpdateAlertNoticeConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlertNoticeConfigResponse), nil
	}
}

// UpdateAlertNoticeConfigInvoker 更新告警通知配置
func (c *WafClient) UpdateAlertNoticeConfigInvoker(request *model.UpdateAlertNoticeConfigRequest) *UpdateAlertNoticeConfigInvoker {
	requestDef := GenReqDefForUpdateAlertNoticeConfig()
	return &UpdateAlertNoticeConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAntiTamperRuleRefresh 网页防篡改规则更新缓存
//
// 网页防篡改规则更新缓存
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) UpdateAntiTamperRuleRefresh(request *model.UpdateAntiTamperRuleRefreshRequest) (*model.UpdateAntiTamperRuleRefreshResponse, error) {
	requestDef := GenReqDefForUpdateAntiTamperRuleRefresh()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAntiTamperRuleRefreshResponse), nil
	}
}

// UpdateAntiTamperRuleRefreshInvoker 网页防篡改规则更新缓存
func (c *WafClient) UpdateAntiTamperRuleRefreshInvoker(request *model.UpdateAntiTamperRuleRefreshRequest) *UpdateAntiTamperRuleRefreshInvoker {
	requestDef := GenReqDefForUpdateAntiTamperRuleRefresh()
	return &UpdateAntiTamperRuleRefreshInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAnticrawlerRule 更新JS脚本反爬虫防护规则
//
// 更新JS脚本反爬虫防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) UpdateAnticrawlerRule(request *model.UpdateAnticrawlerRuleRequest) (*model.UpdateAnticrawlerRuleResponse, error) {
	requestDef := GenReqDefForUpdateAnticrawlerRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAnticrawlerRuleResponse), nil
	}
}

// UpdateAnticrawlerRuleInvoker 更新JS脚本反爬虫防护规则
func (c *WafClient) UpdateAnticrawlerRuleInvoker(request *model.UpdateAnticrawlerRuleRequest) *UpdateAnticrawlerRuleInvoker {
	requestDef := GenReqDefForUpdateAnticrawlerRule()
	return &UpdateAnticrawlerRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAnticrawlerRuleType 更新JS脚本反爬虫规则防护模式
//
// 更新JS脚本反爬虫规则防护模式，在创建JS脚本反爬虫规则前，需要调用该接口指定JS脚本反爬虫规则防护模式。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) UpdateAnticrawlerRuleType(request *model.UpdateAnticrawlerRuleTypeRequest) (*model.UpdateAnticrawlerRuleTypeResponse, error) {
	requestDef := GenReqDefForUpdateAnticrawlerRuleType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAnticrawlerRuleTypeResponse), nil
	}
}

// UpdateAnticrawlerRuleTypeInvoker 更新JS脚本反爬虫规则防护模式
func (c *WafClient) UpdateAnticrawlerRuleTypeInvoker(request *model.UpdateAnticrawlerRuleTypeRequest) *UpdateAnticrawlerRuleTypeInvoker {
	requestDef := GenReqDefForUpdateAnticrawlerRuleType()
	return &UpdateAnticrawlerRuleTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAntileakageRule 更新防敏感信息泄露防护规则
//
// 更新防敏感信息泄露防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) UpdateAntileakageRule(request *model.UpdateAntileakageRuleRequest) (*model.UpdateAntileakageRuleResponse, error) {
	requestDef := GenReqDefForUpdateAntileakageRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAntileakageRuleResponse), nil
	}
}

// UpdateAntileakageRuleInvoker 更新防敏感信息泄露防护规则
func (c *WafClient) UpdateAntileakageRuleInvoker(request *model.UpdateAntileakageRuleRequest) *UpdateAntileakageRuleInvoker {
	requestDef := GenReqDefForUpdateAntileakageRule()
	return &UpdateAntileakageRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCcRule 更新cc防护规则
//
// 更新cc防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) UpdateCcRule(request *model.UpdateCcRuleRequest) (*model.UpdateCcRuleResponse, error) {
	requestDef := GenReqDefForUpdateCcRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCcRuleResponse), nil
	}
}

// UpdateCcRuleInvoker 更新cc防护规则
func (c *WafClient) UpdateCcRuleInvoker(request *model.UpdateCcRuleRequest) *UpdateCcRuleInvoker {
	requestDef := GenReqDefForUpdateCcRule()
	return &UpdateCcRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCertificate 修改证书
//
// 修改证书
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// UpdateCustomRule 更新精准防护规则
//
// 更新精准防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) UpdateCustomRule(request *model.UpdateCustomRuleRequest) (*model.UpdateCustomRuleResponse, error) {
	requestDef := GenReqDefForUpdateCustomRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCustomRuleResponse), nil
	}
}

// UpdateCustomRuleInvoker 更新精准防护规则
func (c *WafClient) UpdateCustomRuleInvoker(request *model.UpdateCustomRuleRequest) *UpdateCustomRuleInvoker {
	requestDef := GenReqDefForUpdateCustomRule()
	return &UpdateCustomRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGeoipRule 更新地理位置控制防护规则
//
// 更新地理位置控制防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// UpdateIgnoreRule 更新全局白名单(原误报屏蔽)防护规则
//
// 更新全局白名单(原误报屏蔽)防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) UpdateIgnoreRule(request *model.UpdateIgnoreRuleRequest) (*model.UpdateIgnoreRuleResponse, error) {
	requestDef := GenReqDefForUpdateIgnoreRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIgnoreRuleResponse), nil
	}
}

// UpdateIgnoreRuleInvoker 更新全局白名单(原误报屏蔽)防护规则
func (c *WafClient) UpdateIgnoreRuleInvoker(request *model.UpdateIgnoreRuleRequest) *UpdateIgnoreRuleInvoker {
	requestDef := GenReqDefForUpdateIgnoreRule()
	return &UpdateIgnoreRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIpGroup 修改ip地址组
//
// 修改ip地址组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) UpdateIpGroup(request *model.UpdateIpGroupRequest) (*model.UpdateIpGroupResponse, error) {
	requestDef := GenReqDefForUpdateIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIpGroupResponse), nil
	}
}

// UpdateIpGroupInvoker 修改ip地址组
func (c *WafClient) UpdateIpGroupInvoker(request *model.UpdateIpGroupRequest) *UpdateIpGroupInvoker {
	requestDef := GenReqDefForUpdateIpGroup()
	return &UpdateIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLtsInfoConfig 配置全量日志lts
//
// 配置全量日志lts，该接口可用来开启与关闭waf全量日志以及配置日志组和日志流。日志组id和日志流id可前往云日志服务获取。配置的日志流id要属于所配置的日志组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) UpdateLtsInfoConfig(request *model.UpdateLtsInfoConfigRequest) (*model.UpdateLtsInfoConfigResponse, error) {
	requestDef := GenReqDefForUpdateLtsInfoConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLtsInfoConfigResponse), nil
	}
}

// UpdateLtsInfoConfigInvoker 配置全量日志lts
func (c *WafClient) UpdateLtsInfoConfigInvoker(request *model.UpdateLtsInfoConfigRequest) *UpdateLtsInfoConfigInvoker {
	requestDef := GenReqDefForUpdateLtsInfoConfig()
	return &UpdateLtsInfoConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePolicy 更新防护策略
//
// 更新防护策略，请求体可只传需要更新的部分
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// UpdatePunishmentRule 更新攻击惩罚规则
//
// 更新攻击惩罚规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WafClient) UpdatePunishmentRule(request *model.UpdatePunishmentRuleRequest) (*model.UpdatePunishmentRuleResponse, error) {
	requestDef := GenReqDefForUpdatePunishmentRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePunishmentRuleResponse), nil
	}
}

// UpdatePunishmentRuleInvoker 更新攻击惩罚规则
func (c *WafClient) UpdatePunishmentRuleInvoker(request *model.UpdatePunishmentRuleRequest) *UpdatePunishmentRuleInvoker {
	requestDef := GenReqDefForUpdatePunishmentRule()
	return &UpdatePunishmentRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateValueList 修改引用表
//
// 修改引用表
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
