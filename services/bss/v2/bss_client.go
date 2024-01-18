package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"
)

type BssClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewBssClient(hcClient *httpclient.HcHttpClient) *BssClient {
	return &BssClient{HcClient: hcClient}
}

func BssClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// AutoRenewalResources 设置包年/包月资源自动续费
//
// 为防止资源到期被删除，客户可为长期使用的包年/包月资源开通自动续费。
//
// 客户在费用中心开通自动续费请参见[这里](https://support.huaweicloud.com/usermanual-billing/renewals_topic_20000003.html)。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;-   首先要客户成功支付包年/包月资源订单，才能进行自动续费的开通。
// &gt;-   目前支持设置自动续费的包年/包月产品请参见[自动续费规则说明](https://support.huaweicloud.com/usermanual-billing/renewals_topic_20000002.html)。
// &gt;-   在调用本接口前，您可以调用“[查询客户包年/包月资源列表](https://support.huaweicloud.com/api-oce/api_order_00021.html)”接口获取资源ID、资源过期时间以及资源过期后扣费策略等信息。
// &gt;-   自动续费将于产品到期前7天的凌晨3:00开始扣款，请保持账户余额充足。若由于账户中余额不足等原因导致第一次未扣费成功，系统将每天凌晨3:00尝试进行一次扣款，直到扣款成功或保留产品资源的最后一天。
// &gt;-   续费周期与原资源的购买周期一致。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) AutoRenewalResources(request *model.AutoRenewalResourcesRequest) (*model.AutoRenewalResourcesResponse, error) {
	requestDef := GenReqDefForAutoRenewalResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AutoRenewalResourcesResponse), nil
	}
}

// AutoRenewalResourcesInvoker 设置包年/包月资源自动续费
func (c *BssClient) AutoRenewalResourcesInvoker(request *model.AutoRenewalResourcesRequest) *AutoRenewalResourcesInvoker {
	requestDef := GenReqDefForAutoRenewalResources()
	return &AutoRenewalResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelAutoRenewalResources 取消包年/包月资源自动续费
//
// 客户设置自动续费后，还可以执行取消自动续费的操作。关闭自动续费后，资源到期将不会被自动续费。
//
// 客户在费用中心取消包年/包月资源自动续费请参见[这里](https://support.huaweicloud.com/usermanual-billing/renewals_topic_20000005.html)。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;-   前提是已经调用“[设置包年/包月资源自动续费](https://support.huaweicloud.com/api-oce/api_order_00022.html)”接口设置自动续费或在调用“[续订包年/包月资源](https://support.huaweicloud.com/api-oce/api_order_00018.html)”接口时设置到期策略为自动续订。
// &gt;-   目前支持取消自动续费的包年/包月产品同支持自动续费的包年/包月产品。
// &gt;-   在调用本接口前，您可以调用“[查询客户包年/包月资源列表](https://support.huaweicloud.com/api-bpconsole/api_order_00021.html)”接口获取资源ID、资源过期时间以及资源过期后扣费策略等信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) CancelAutoRenewalResources(request *model.CancelAutoRenewalResourcesRequest) (*model.CancelAutoRenewalResourcesResponse, error) {
	requestDef := GenReqDefForCancelAutoRenewalResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelAutoRenewalResourcesResponse), nil
	}
}

// CancelAutoRenewalResourcesInvoker 取消包年/包月资源自动续费
func (c *BssClient) CancelAutoRenewalResourcesInvoker(request *model.CancelAutoRenewalResourcesRequest) *CancelAutoRenewalResourcesInvoker {
	requestDef := GenReqDefForCancelAutoRenewalResources()
	return &CancelAutoRenewalResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelCustomerOrder 取消待支付订单
//
// 客户可以对待支付的订单进行取消操作。
//
// 客户登录费用中心取消包年包月产品的待支付订单请单击[这里](https://support.huaweicloud.com/usermanual-billing/zh-cn_topic_0031465730.html)。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;只有订单状态是“待支付”的时候，才能取消订单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) CancelCustomerOrder(request *model.CancelCustomerOrderRequest) (*model.CancelCustomerOrderResponse, error) {
	requestDef := GenReqDefForCancelCustomerOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelCustomerOrderResponse), nil
	}
}

// CancelCustomerOrderInvoker 取消待支付订单
func (c *BssClient) CancelCustomerOrderInvoker(request *model.CancelCustomerOrderRequest) *CancelCustomerOrderInvoker {
	requestDef := GenReqDefForCancelCustomerOrder()
	return &CancelCustomerOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelResourcesSubscription 退订包年/包月资源
//
// 客户购买包年/包月资源后，支持客户退订包年/包月实例。退订资源实例包括资源续费部分和当前正在使用的部分，退订后资源将无法使用。
//
// 客户在费用中心退订已购买的包年包月资源请参见[这里](https://support.huaweicloud.com/usermanual-billing/zh-cn_topic_0083138805.html)。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;-   首先要成功支付包年/包月产品，产生一条开通成功的包年/包月资源，才能进行退订。
// &gt;-   调用接口后，如果某个主资源有对应的从资源，系统会将主资源和从资源一起退订，主资源的从资源信息可以通过调用[查询客户包年/包月资源列表](https://support.huaweicloud.com/api-oce/api_order_00021.html)接口获取。
// &gt;-   注意：如ECS主机挂载新购的云硬盘，但此硬盘不是该ECS主资源的从资源，主从资源信息必须以调用[查询客户包年/包月资源列表](https://support.huaweicloud.com/api-oce/api_order_00021.html)接口获取的信息为准。
// &gt;-   调用该接口后，您还可以调用“[查询退款订单的金额详情](查询退款订单的金额详情.md)”接口查询退订订单对应的金额来自哪些订单。
// &gt;-   该接口支持5天无理由全额退订，具体规则请参见“[退订规则说明](https://support.huaweicloud.com/usermanual-billing/unsubscription_topic_20000081.html)”。
// &gt;-   您正在退订使用中的资源，请仔细确认资源信息和退款信息。未放入回收站的资源退订后无法恢复，若您要保留资源，仅退订未使用的续费周期，请退订续费周期。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) CancelResourcesSubscription(request *model.CancelResourcesSubscriptionRequest) (*model.CancelResourcesSubscriptionResponse, error) {
	requestDef := GenReqDefForCancelResourcesSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelResourcesSubscriptionResponse), nil
	}
}

// CancelResourcesSubscriptionInvoker 退订包年/包月资源
func (c *BssClient) CancelResourcesSubscriptionInvoker(request *model.CancelResourcesSubscriptionRequest) *CancelResourcesSubscriptionInvoker {
	requestDef := GenReqDefForCancelResourcesSubscription()
	return &CancelResourcesSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeEnterpriseRealnameAuthentication 申请实名认证变更
//
// 客户可以进行实名认证变更申请。
//
// 个人客户登录帐号中心通过实名认证变更为企业帐号的方式及流程请参见[这里](https://support.huaweicloud.com/usermanual-account/zh-cn_topic_0103532632.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ChangeEnterpriseRealnameAuthentication(request *model.ChangeEnterpriseRealnameAuthenticationRequest) (*model.ChangeEnterpriseRealnameAuthenticationResponse, error) {
	requestDef := GenReqDefForChangeEnterpriseRealnameAuthentication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeEnterpriseRealnameAuthenticationResponse), nil
	}
}

// ChangeEnterpriseRealnameAuthenticationInvoker 申请实名认证变更
func (c *BssClient) ChangeEnterpriseRealnameAuthenticationInvoker(request *model.ChangeEnterpriseRealnameAuthenticationRequest) *ChangeEnterpriseRealnameAuthenticationInvoker {
	requestDef := GenReqDefForChangeEnterpriseRealnameAuthentication()
	return &ChangeEnterpriseRealnameAuthenticationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckUserIdentity 校验客户注册信息
//
// 客户注册时可检查客户的登录名称、手机号或者邮箱是否可以用于注册。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;针对校验手机号场景，目前仅支持校验手机号注册数量是否超过上限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) CheckUserIdentity(request *model.CheckUserIdentityRequest) (*model.CheckUserIdentityResponse, error) {
	requestDef := GenReqDefForCheckUserIdentity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckUserIdentityResponse), nil
	}
}

// CheckUserIdentityInvoker 校验客户注册信息
func (c *BssClient) CheckUserIdentityInvoker(request *model.CheckUserIdentityRequest) *CheckUserIdentityInvoker {
	requestDef := GenReqDefForCheckUserIdentity()
	return &CheckUserIdentityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ClaimEnterpriseMultiAccountCoupon 企业主账号向企业子账号拨款优惠券
//
// 企业主账号在自建平台向企业子账号拨款优惠券。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;-   仅支持华为发放的测试类、商务类、活动类代金券。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ClaimEnterpriseMultiAccountCoupon(request *model.ClaimEnterpriseMultiAccountCouponRequest) (*model.ClaimEnterpriseMultiAccountCouponResponse, error) {
	requestDef := GenReqDefForClaimEnterpriseMultiAccountCoupon()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ClaimEnterpriseMultiAccountCouponResponse), nil
	}
}

// ClaimEnterpriseMultiAccountCouponInvoker 企业主账号向企业子账号拨款优惠券
func (c *BssClient) ClaimEnterpriseMultiAccountCouponInvoker(request *model.ClaimEnterpriseMultiAccountCouponRequest) *ClaimEnterpriseMultiAccountCouponInvoker {
	requestDef := GenReqDefForClaimEnterpriseMultiAccountCoupon()
	return &ClaimEnterpriseMultiAccountCouponInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEnterpriseProjectAuth 开通客户企业项目权限
//
// 客户在自建平台开通客户企业项目权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) CreateEnterpriseProjectAuth(request *model.CreateEnterpriseProjectAuthRequest) (*model.CreateEnterpriseProjectAuthResponse, error) {
	requestDef := GenReqDefForCreateEnterpriseProjectAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnterpriseProjectAuthResponse), nil
	}
}

// CreateEnterpriseProjectAuthInvoker 开通客户企业项目权限
func (c *BssClient) CreateEnterpriseProjectAuthInvoker(request *model.CreateEnterpriseProjectAuthRequest) *CreateEnterpriseProjectAuthInvoker {
	requestDef := GenReqDefForCreateEnterpriseProjectAuth()
	return &CreateEnterpriseProjectAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEnterpriseRealnameAuthentication 申请企业实名认证
//
// 企业客户可以进行企业实名认证申请。
//
// 客户登录帐号中心进行企业实名认证的方式及流程请参见[这里](https://support.huaweicloud.com/usermanual-account/zh-cn_topic_0077914253.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) CreateEnterpriseRealnameAuthentication(request *model.CreateEnterpriseRealnameAuthenticationRequest) (*model.CreateEnterpriseRealnameAuthenticationResponse, error) {
	requestDef := GenReqDefForCreateEnterpriseRealnameAuthentication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnterpriseRealnameAuthenticationResponse), nil
	}
}

// CreateEnterpriseRealnameAuthenticationInvoker 申请企业实名认证
func (c *BssClient) CreateEnterpriseRealnameAuthenticationInvoker(request *model.CreateEnterpriseRealnameAuthenticationRequest) *CreateEnterpriseRealnameAuthenticationInvoker {
	requestDef := GenReqDefForCreateEnterpriseRealnameAuthentication()
	return &CreateEnterpriseRealnameAuthenticationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePartnerCoupons 发放优惠券
//
// 合作伙伴可以在拥有的代金券额度范围内为客户下发优惠券。
//
// 伙伴登录合作伙伴中心为客户发放代金券请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/espp_050502.html)。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;只能给代售子客户发放优惠券。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) CreatePartnerCoupons(request *model.CreatePartnerCouponsRequest) (*model.CreatePartnerCouponsResponse, error) {
	requestDef := GenReqDefForCreatePartnerCoupons()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePartnerCouponsResponse), nil
	}
}

// CreatePartnerCouponsInvoker 发放优惠券
func (c *BssClient) CreatePartnerCouponsInvoker(request *model.CreatePartnerCouponsRequest) *CreatePartnerCouponsInvoker {
	requestDef := GenReqDefForCreatePartnerCoupons()
	return &CreatePartnerCouponsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePersonalRealnameAuth 申请个人实名认证
//
// 个人客户可以进行个人实名认证申请。
//
// 客户登录帐号中心进行个人实名认证的方式及流程请参见[这里](https://support.huaweicloud.com/usermanual-account/zh-cn_topic_0077914254.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) CreatePersonalRealnameAuth(request *model.CreatePersonalRealnameAuthRequest) (*model.CreatePersonalRealnameAuthResponse, error) {
	requestDef := GenReqDefForCreatePersonalRealnameAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePersonalRealnameAuthResponse), nil
	}
}

// CreatePersonalRealnameAuthInvoker 申请个人实名认证
func (c *BssClient) CreatePersonalRealnameAuthInvoker(request *model.CreatePersonalRealnameAuthRequest) *CreatePersonalRealnameAuthInvoker {
	requestDef := GenReqDefForCreatePersonalRealnameAuth()
	return &CreatePersonalRealnameAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSubCustomer 创建客户
//
// 在伙伴销售平台创建客户时同步创建华为云账号，并将客户在伙伴销售平台上的账号与华为云账号进行映射。同时，创建的华为云账号与伙伴账号关联绑定。
//
// 华为云总经销商（一级经销商）可以注册云经销商（二级经销商）的子客户。注册完成后，子客户可以自动和云经销商绑定。
//
// &gt;![](public_sys-resources/icon-caution.gif) **注意：**
// &gt;-   调用该接口为客户创建华为云账号后，如果想从合作伙伴销售平台跳转至华为云官网，还需要进行SAML认证，具体请参见“[Web UI方式](https://support.huaweicloud.com/api-bpconsole/jac_00001.html)”中的“SAML认证”。
// &gt;-   如果创建的时候不输入手机号，那么客户将无法收到华为云发出的任何提醒短信，需要客户自己登录到华为云平台补充手机号。
// &gt;-   调用“创建客户”接口时，华为云会同步创建华为云客户账号，将客户ID及账号名返回给伙伴平台，然后华为云异步完成客户与伙伴的关联。伙伴与客户的关联结果可通过“[查询客户列表](https://support.huaweicloud.com/api-bpconsole/mc_00021.html)”查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) CreateSubCustomer(request *model.CreateSubCustomerRequest) (*model.CreateSubCustomerResponse, error) {
	requestDef := GenReqDefForCreateSubCustomer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubCustomerResponse), nil
	}
}

// CreateSubCustomerInvoker 创建客户
func (c *BssClient) CreateSubCustomerInvoker(request *model.CreateSubCustomerRequest) *CreateSubCustomerInvoker {
	requestDef := GenReqDefForCreateSubCustomer()
	return &CreateSubCustomerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSubEnterpriseAccount 创建企业子账号
//
// 企业主账号在自建平台创建企业子账号。
//
// 企业主账号创建企业子账号请参见[这里](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0104194162.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) CreateSubEnterpriseAccount(request *model.CreateSubEnterpriseAccountRequest) (*model.CreateSubEnterpriseAccountResponse, error) {
	requestDef := GenReqDefForCreateSubEnterpriseAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubEnterpriseAccountResponse), nil
	}
}

// CreateSubEnterpriseAccountInvoker 创建企业子账号
func (c *BssClient) CreateSubEnterpriseAccountInvoker(request *model.CreateSubEnterpriseAccountRequest) *CreateSubEnterpriseAccountInvoker {
	requestDef := GenReqDefForCreateSubEnterpriseAccount()
	return &CreateSubEnterpriseAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCities 查询城市信息
//
// 伙伴在伙伴销售平台上查询城市信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListCities(request *model.ListCitiesRequest) (*model.ListCitiesResponse, error) {
	requestDef := GenReqDefForListCities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCitiesResponse), nil
	}
}

// ListCitiesInvoker 查询城市信息
func (c *BssClient) ListCitiesInvoker(request *model.ListCitiesRequest) *ListCitiesInvoker {
	requestDef := GenReqDefForListCities()
	return &ListCitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConsumeSubCustomers 查询伙伴消费子客户列表
//
// 伙伴在伙伴销售平台可实时查询消费账期内的子客户列表，可以用于查询子客户的消费记录。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;若当前子客户与伙伴无关联关系，仍可查询账期内处于关联状态且有消费的子客户列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListConsumeSubCustomers(request *model.ListConsumeSubCustomersRequest) (*model.ListConsumeSubCustomersResponse, error) {
	requestDef := GenReqDefForListConsumeSubCustomers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConsumeSubCustomersResponse), nil
	}
}

// ListConsumeSubCustomersInvoker 查询伙伴消费子客户列表
func (c *BssClient) ListConsumeSubCustomersInvoker(request *model.ListConsumeSubCustomersRequest) *ListConsumeSubCustomersInvoker {
	requestDef := GenReqDefForListConsumeSubCustomers()
	return &ListConsumeSubCustomersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConversions 查询度量单位进制
//
// 伙伴在伙伴销售平台上查询度量单位的进制转换信息，用于不同度量单位之间的转换。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListConversions(request *model.ListConversionsRequest) (*model.ListConversionsResponse, error) {
	requestDef := GenReqDefForListConversions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConversionsResponse), nil
	}
}

// ListConversionsInvoker 查询度量单位进制
func (c *BssClient) ListConversionsInvoker(request *model.ListConversionsRequest) *ListConversionsInvoker {
	requestDef := GenReqDefForListConversions()
	return &ListConversionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCosts 查询成本数据
//
// 客户在自建平台查询成本分析数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListCosts(request *model.ListCostsRequest) (*model.ListCostsResponse, error) {
	requestDef := GenReqDefForListCosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCostsResponse), nil
	}
}

// ListCostsInvoker 查询成本数据
func (c *BssClient) ListCostsInvoker(request *model.ListCostsRequest) *ListCostsInvoker {
	requestDef := GenReqDefForListCosts()
	return &ListCostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCounties 查询区县信息
//
// 伙伴在伙伴销售平台上查询区县信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListCounties(request *model.ListCountiesRequest) (*model.ListCountiesResponse, error) {
	requestDef := GenReqDefForListCounties()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCountiesResponse), nil
	}
}

// ListCountiesInvoker 查询区县信息
func (c *BssClient) ListCountiesInvoker(request *model.ListCountiesRequest) *ListCountiesInvoker {
	requestDef := GenReqDefForListCounties()
	return &ListCountiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCouponQuotasRecords 查询代金券额度的发放回收记录
//
// 华为云总经销商（一级经销商）可以查看给云经销商（二级经销商）发放或回收代金券额度的操作记录。
//
// 一级经销商可以登录伙伴中心，进入“客户业务** **\\&gt; 代金券管理”，选择“代金券额度”页签，单击“操作记录”查看代金券额度的发放和回收记录。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;伙伴也可以单击代金券额度所在行的“操作记录”，查看该代金券额度对应的操作记录日志。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListCouponQuotasRecords(request *model.ListCouponQuotasRecordsRequest) (*model.ListCouponQuotasRecordsResponse, error) {
	requestDef := GenReqDefForListCouponQuotasRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCouponQuotasRecordsResponse), nil
	}
}

// ListCouponQuotasRecordsInvoker 查询代金券额度的发放回收记录
func (c *BssClient) ListCouponQuotasRecordsInvoker(request *model.ListCouponQuotasRecordsRequest) *ListCouponQuotasRecordsInvoker {
	requestDef := GenReqDefForListCouponQuotasRecords()
	return &ListCouponQuotasRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomerAccountChangeRecords 查询收支明细(客户)
//
// 功能描述：客户可以查询自身的收支明细情况(此接口不适用于伙伴的代售类、转售类客户。)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListCustomerAccountChangeRecords(request *model.ListCustomerAccountChangeRecordsRequest) (*model.ListCustomerAccountChangeRecordsResponse, error) {
	requestDef := GenReqDefForListCustomerAccountChangeRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerAccountChangeRecordsResponse), nil
	}
}

// ListCustomerAccountChangeRecordsInvoker 查询收支明细(客户)
func (c *BssClient) ListCustomerAccountChangeRecordsInvoker(request *model.ListCustomerAccountChangeRecordsRequest) *ListCustomerAccountChangeRecordsInvoker {
	requestDef := GenReqDefForListCustomerAccountChangeRecords()
	return &ListCustomerAccountChangeRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomerBillsFeeRecords 查询流水账单
//
// 客户在自建平台查询自己的消费流水账单。
//
// 客户登录费用中心查询自己的消费流水账单请参见[这里](https://support.huaweicloud.com/usermanual-billing/bills-topic_80000001.html#bills-topic_80000001__zh-cn_topic_0000001162496407_s716e04d5d0ba4e9d9a76a8bcbfbcfe73)的“**查看流水账单**”。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListCustomerBillsFeeRecords(request *model.ListCustomerBillsFeeRecordsRequest) (*model.ListCustomerBillsFeeRecordsResponse, error) {
	requestDef := GenReqDefForListCustomerBillsFeeRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerBillsFeeRecordsResponse), nil
	}
}

// ListCustomerBillsFeeRecordsInvoker 查询流水账单
func (c *BssClient) ListCustomerBillsFeeRecordsInvoker(request *model.ListCustomerBillsFeeRecordsRequest) *ListCustomerBillsFeeRecordsInvoker {
	requestDef := GenReqDefForListCustomerBillsFeeRecords()
	return &ListCustomerBillsFeeRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomerBillsMonthlyBreakDown 查询月度成本
//
// 客户可以查询指定月份的月度摊销成本。当前仅支持查询近18个月的摊销成本。摊销成本计算规则请参见[成本摊销规则](https://support.huaweicloud.com/usermanual-cost/costcenter_000002_01.html)。
//
// 客户可查询的数据范围同成本中心提供的[数据范围](https://support.huaweicloud.com/usermanual-cost/costcenter_0000004.html)一致。
//
// 客户登录成本中心导出成本明细请参见[导出成本明细数据](https://support.huaweicloud.com/usermanual-cost/costcenter_000002_03.html)。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;该接口仅面向已开通成本中心的客户开放。如何开启成本中心请参见[这里](https://support.huaweicloud.com/usermanual-cost/costcenter_000004.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListCustomerBillsMonthlyBreakDown(request *model.ListCustomerBillsMonthlyBreakDownRequest) (*model.ListCustomerBillsMonthlyBreakDownResponse, error) {
	requestDef := GenReqDefForListCustomerBillsMonthlyBreakDown()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerBillsMonthlyBreakDownResponse), nil
	}
}

// ListCustomerBillsMonthlyBreakDownInvoker 查询月度成本
func (c *BssClient) ListCustomerBillsMonthlyBreakDownInvoker(request *model.ListCustomerBillsMonthlyBreakDownRequest) *ListCustomerBillsMonthlyBreakDownInvoker {
	requestDef := GenReqDefForListCustomerBillsMonthlyBreakDown()
	return &ListCustomerBillsMonthlyBreakDownInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomerOnDemandResources 查询客户按需资源列表
//
// 合作伙伴可以查询关联的代售类客户已开通的按需资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListCustomerOnDemandResources(request *model.ListCustomerOnDemandResourcesRequest) (*model.ListCustomerOnDemandResourcesResponse, error) {
	requestDef := GenReqDefForListCustomerOnDemandResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerOnDemandResourcesResponse), nil
	}
}

// ListCustomerOnDemandResourcesInvoker 查询客户按需资源列表
func (c *BssClient) ListCustomerOnDemandResourcesInvoker(request *model.ListCustomerOnDemandResourcesRequest) *ListCustomerOnDemandResourcesInvoker {
	requestDef := GenReqDefForListCustomerOnDemandResources()
	return &ListCustomerOnDemandResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomerOrders 查询订单列表
//
// 客户购买包年/包月资源后，可以查看待审核、处理中、已取消、已完成和待支付等状态的订单。
//
// 伙伴登录伙伴中心查看客户订单请单击[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0076200001.html)。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;如果想查询某条订单下的资源信息，在调用本接口获取订单ID后，请调用“[查询客户包年/包月资源列表](https://support.huaweicloud.com/api-oce/api_order_00021.html)”接口在请求参数输入订单号进行查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListCustomerOrders(request *model.ListCustomerOrdersRequest) (*model.ListCustomerOrdersResponse, error) {
	requestDef := GenReqDefForListCustomerOrders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerOrdersResponse), nil
	}
}

// ListCustomerOrdersInvoker 查询订单列表
func (c *BssClient) ListCustomerOrdersInvoker(request *model.ListCustomerOrdersRequest) *ListCustomerOrdersInvoker {
	requestDef := GenReqDefForListCustomerOrders()
	return &ListCustomerOrdersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomersBalancesDetail 查询客户账户余额
//
// 合作伙伴可以查询关联的代售类客户的账户余额。
//
// 伙伴登录伙伴中心查询客户余额请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435115.html)。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;顾问销售类客户是客户在华为云充值，合作伙伴无法调用此接口查询其账户余额。代售类客户的账户由合作伙伴拨款，所以可以查询到。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListCustomersBalancesDetail(request *model.ListCustomersBalancesDetailRequest) (*model.ListCustomersBalancesDetailResponse, error) {
	requestDef := GenReqDefForListCustomersBalancesDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomersBalancesDetailResponse), nil
	}
}

// ListCustomersBalancesDetailInvoker 查询客户账户余额
func (c *BssClient) ListCustomersBalancesDetailInvoker(request *model.ListCustomersBalancesDetailRequest) *ListCustomersBalancesDetailInvoker {
	requestDef := GenReqDefForListCustomersBalancesDetail()
	return &ListCustomersBalancesDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomerselfResourceRecordDetails 查询资源详单
//
// 客户在自建平台查询自己的资源详单，用于反映各类资源的消耗情况。
//
// 客户登录费用中心查询资源详单请参见[这里](https://support.huaweicloud.com/usermanual-billing/bills_topic_100000063.html)。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;由于资源消费呈现的是资源维度的8位小数原始消费金额，实际从账户扣费时按2位小数进行扣费（即扣到分），会存在精度差异，所以，不推荐消费汇总和资源消费直接对账。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListCustomerselfResourceRecordDetails(request *model.ListCustomerselfResourceRecordDetailsRequest) (*model.ListCustomerselfResourceRecordDetailsResponse, error) {
	requestDef := GenReqDefForListCustomerselfResourceRecordDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerselfResourceRecordDetailsResponse), nil
	}
}

// ListCustomerselfResourceRecordDetailsInvoker 查询资源详单
func (c *BssClient) ListCustomerselfResourceRecordDetailsInvoker(request *model.ListCustomerselfResourceRecordDetailsRequest) *ListCustomerselfResourceRecordDetailsInvoker {
	requestDef := GenReqDefForListCustomerselfResourceRecordDetails()
	return &ListCustomerselfResourceRecordDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomerselfResourceRecords 查询资源消费记录
//
// 客户在自建平台查询每个资源的消费明细数据。
//
// 客户登录费用中心查询资源消费记录请参见[这里](https://support.huaweicloud.com/usermanual-billing/bills_topic_100000061.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListCustomerselfResourceRecords(request *model.ListCustomerselfResourceRecordsRequest) (*model.ListCustomerselfResourceRecordsResponse, error) {
	requestDef := GenReqDefForListCustomerselfResourceRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerselfResourceRecordsResponse), nil
	}
}

// ListCustomerselfResourceRecordsInvoker 查询资源消费记录
func (c *BssClient) ListCustomerselfResourceRecordsInvoker(request *model.ListCustomerselfResourceRecordsRequest) *ListCustomerselfResourceRecordsInvoker {
	requestDef := GenReqDefForListCustomerselfResourceRecords()
	return &ListCustomerselfResourceRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnterpriseMultiAccount 查询企业子账号可回收余额
//
// 企业主账号在自建平台查询企业子账号的可回收余额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListEnterpriseMultiAccount(request *model.ListEnterpriseMultiAccountRequest) (*model.ListEnterpriseMultiAccountResponse, error) {
	requestDef := GenReqDefForListEnterpriseMultiAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnterpriseMultiAccountResponse), nil
	}
}

// ListEnterpriseMultiAccountInvoker 查询企业子账号可回收余额
func (c *BssClient) ListEnterpriseMultiAccountInvoker(request *model.ListEnterpriseMultiAccountRequest) *ListEnterpriseMultiAccountInvoker {
	requestDef := GenReqDefForListEnterpriseMultiAccount()
	return &ListEnterpriseMultiAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnterpriseOrganizations 查询企业组织结构
//
// 企业主账号在自建平台查询企业组织结构。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListEnterpriseOrganizations(request *model.ListEnterpriseOrganizationsRequest) (*model.ListEnterpriseOrganizationsResponse, error) {
	requestDef := GenReqDefForListEnterpriseOrganizations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnterpriseOrganizationsResponse), nil
	}
}

// ListEnterpriseOrganizationsInvoker 查询企业组织结构
func (c *BssClient) ListEnterpriseOrganizationsInvoker(request *model.ListEnterpriseOrganizationsRequest) *ListEnterpriseOrganizationsInvoker {
	requestDef := GenReqDefForListEnterpriseOrganizations()
	return &ListEnterpriseOrganizationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnterpriseSubCustomers 查询企业子账号列表
//
// 企业主账号在自建平台查询企业子账号信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListEnterpriseSubCustomers(request *model.ListEnterpriseSubCustomersRequest) (*model.ListEnterpriseSubCustomersResponse, error) {
	requestDef := GenReqDefForListEnterpriseSubCustomers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnterpriseSubCustomersResponse), nil
	}
}

// ListEnterpriseSubCustomersInvoker 查询企业子账号列表
func (c *BssClient) ListEnterpriseSubCustomersInvoker(request *model.ListEnterpriseSubCustomersRequest) *ListEnterpriseSubCustomersInvoker {
	requestDef := GenReqDefForListEnterpriseSubCustomers()
	return &ListEnterpriseSubCustomersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFreeResourceInfos 查询资源包列表
//
// 客户在伙伴销售平台查询客户的资源包列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListFreeResourceInfos(request *model.ListFreeResourceInfosRequest) (*model.ListFreeResourceInfosResponse, error) {
	requestDef := GenReqDefForListFreeResourceInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFreeResourceInfosResponse), nil
	}
}

// ListFreeResourceInfosInvoker 查询资源包列表
func (c *BssClient) ListFreeResourceInfosInvoker(request *model.ListFreeResourceInfosRequest) *ListFreeResourceInfosInvoker {
	requestDef := GenReqDefForListFreeResourceInfos()
	return &ListFreeResourceInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFreeResourceUsages 查询资源包使用量
//
// 客户在伙伴销售平台根据资源项维度查询客户的资源包使用量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListFreeResourceUsages(request *model.ListFreeResourceUsagesRequest) (*model.ListFreeResourceUsagesResponse, error) {
	requestDef := GenReqDefForListFreeResourceUsages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFreeResourceUsagesResponse), nil
	}
}

// ListFreeResourceUsagesInvoker 查询资源包使用量
func (c *BssClient) ListFreeResourceUsagesInvoker(request *model.ListFreeResourceUsagesRequest) *ListFreeResourceUsagesInvoker {
	requestDef := GenReqDefForListFreeResourceUsages()
	return &ListFreeResourceUsagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFreeResourcesUsageRecords 查询资源包使用明细
//
// 客户在自建平台查询资源包使用明细。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListFreeResourcesUsageRecords(request *model.ListFreeResourcesUsageRecordsRequest) (*model.ListFreeResourcesUsageRecordsResponse, error) {
	requestDef := GenReqDefForListFreeResourcesUsageRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFreeResourcesUsageRecordsResponse), nil
	}
}

// ListFreeResourcesUsageRecordsInvoker 查询资源包使用明细
func (c *BssClient) ListFreeResourcesUsageRecordsInvoker(request *model.ListFreeResourcesUsageRecordsRequest) *ListFreeResourcesUsageRecordsInvoker {
	requestDef := GenReqDefForListFreeResourcesUsageRecords()
	return &ListFreeResourcesUsageRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIncentiveDiscountPolicies 查询产品的折扣和激励策略
//
// 伙伴在伙伴销售平台上查询产品的折扣和激励策略。
//
// 伙伴登录合作伙伴中心查看产品的折扣和激励策略请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/dp_120400.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListIncentiveDiscountPolicies(request *model.ListIncentiveDiscountPoliciesRequest) (*model.ListIncentiveDiscountPoliciesResponse, error) {
	requestDef := GenReqDefForListIncentiveDiscountPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIncentiveDiscountPoliciesResponse), nil
	}
}

// ListIncentiveDiscountPoliciesInvoker 查询产品的折扣和激励策略
func (c *BssClient) ListIncentiveDiscountPoliciesInvoker(request *model.ListIncentiveDiscountPoliciesRequest) *ListIncentiveDiscountPoliciesInvoker {
	requestDef := GenReqDefForListIncentiveDiscountPolicies()
	return &ListIncentiveDiscountPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIndirectPartners 查询云经销商列表
//
// 华为云总经销商（一级经销商）可以查询云经销商（二级经销商）列表。
//
// 一级经销商在伙伴中心查询二级经销商列表的方式请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/dp_120210.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListIndirectPartners(request *model.ListIndirectPartnersRequest) (*model.ListIndirectPartnersResponse, error) {
	requestDef := GenReqDefForListIndirectPartners()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIndirectPartnersResponse), nil
	}
}

// ListIndirectPartnersInvoker 查询云经销商列表
func (c *BssClient) ListIndirectPartnersInvoker(request *model.ListIndirectPartnersRequest) *ListIndirectPartnersInvoker {
	requestDef := GenReqDefForListIndirectPartners()
	return &ListIndirectPartnersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIssuedCouponQuotas 查询已发放的代金券额度
//
// 华为云总经销商（一级经销商）可以查看发放给云经销商（二级经销商）的代金券额度列表。
//
// 一级经销商登录伙伴中心，进入“客户业务** **\\&gt; 代金券管理”，选择“已发放代金券额度”可查看代金券额度列表。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;调用该接口之前，需通过客户经理联系华为运营人员，为合作伙伴设置代金券发放额度。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListIssuedCouponQuotas(request *model.ListIssuedCouponQuotasRequest) (*model.ListIssuedCouponQuotasResponse, error) {
	requestDef := GenReqDefForListIssuedCouponQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIssuedCouponQuotasResponse), nil
	}
}

// ListIssuedCouponQuotasInvoker 查询已发放的代金券额度
func (c *BssClient) ListIssuedCouponQuotasInvoker(request *model.ListIssuedCouponQuotasRequest) *ListIssuedCouponQuotasInvoker {
	requestDef := GenReqDefForListIssuedCouponQuotas()
	return &ListIssuedCouponQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIssuedPartnerCoupons 查询已发放的优惠券
//
// 合作伙伴可以查询已发放的优惠券列表。
//
// 伙伴登录伙伴中心，进入“客户业务** **\\&gt; 代金券管理”，选择“已发放代金券”页签，即可查询已发放的代金券。
//
// 伙伴登录伙伴中心，进入“客户业务 \\&gt; 现金券管理”，选择“已发放现金券”页签，即可查询已发放的现金券。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;-   只可查到失效时间在12个月内的代金/现金劵。
// &gt;-   在API只可以查询代售子客户已发放的代金/现金劵，在伙伴中心可以查询代售和顾问销售已发放的代金/现金劵，对比一致性时需要注意关联模式是否一致。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListIssuedPartnerCoupons(request *model.ListIssuedPartnerCouponsRequest) (*model.ListIssuedPartnerCouponsResponse, error) {
	requestDef := GenReqDefForListIssuedPartnerCoupons()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIssuedPartnerCouponsResponse), nil
	}
}

// ListIssuedPartnerCouponsInvoker 查询已发放的优惠券
func (c *BssClient) ListIssuedPartnerCouponsInvoker(request *model.ListIssuedPartnerCouponsRequest) *ListIssuedPartnerCouponsInvoker {
	requestDef := GenReqDefForListIssuedPartnerCoupons()
	return &ListIssuedPartnerCouponsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMeasureUnits 查询度量单位列表
//
// 伙伴在伙伴销售平台上查询资源使用量，包年包月资源的时长及金额的度量单位及名称，度量单位类型等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListMeasureUnits(request *model.ListMeasureUnitsRequest) (*model.ListMeasureUnitsResponse, error) {
	requestDef := GenReqDefForListMeasureUnits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMeasureUnitsResponse), nil
	}
}

// ListMeasureUnitsInvoker 查询度量单位列表
func (c *BssClient) ListMeasureUnitsInvoker(request *model.ListMeasureUnitsRequest) *ListMeasureUnitsInvoker {
	requestDef := GenReqDefForListMeasureUnits()
	return &ListMeasureUnitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMultiAccountRetrieveCoupons 查询企业子账号可回收优惠券列表
//
// 企业主账号在自建平台查询企业子账号的可回收优惠券。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;-   仅支持华为发放的测试类、商务类、活动类代金券。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListMultiAccountRetrieveCoupons(request *model.ListMultiAccountRetrieveCouponsRequest) (*model.ListMultiAccountRetrieveCouponsResponse, error) {
	requestDef := GenReqDefForListMultiAccountRetrieveCoupons()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMultiAccountRetrieveCouponsResponse), nil
	}
}

// ListMultiAccountRetrieveCouponsInvoker 查询企业子账号可回收优惠券列表
func (c *BssClient) ListMultiAccountRetrieveCouponsInvoker(request *model.ListMultiAccountRetrieveCouponsRequest) *ListMultiAccountRetrieveCouponsInvoker {
	requestDef := GenReqDefForListMultiAccountRetrieveCoupons()
	return &ListMultiAccountRetrieveCouponsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMultiAccountTransferCoupons 查询企业主账号可拨款优惠券列表
//
// 企业主账号在自建平台查询自己的可拨款优惠券列表。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;-   仅支持华为发放的测试类、商务类、活动类代金券。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListMultiAccountTransferCoupons(request *model.ListMultiAccountTransferCouponsRequest) (*model.ListMultiAccountTransferCouponsResponse, error) {
	requestDef := GenReqDefForListMultiAccountTransferCoupons()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMultiAccountTransferCouponsResponse), nil
	}
}

// ListMultiAccountTransferCouponsInvoker 查询企业主账号可拨款优惠券列表
func (c *BssClient) ListMultiAccountTransferCouponsInvoker(request *model.ListMultiAccountTransferCouponsRequest) *ListMultiAccountTransferCouponsInvoker {
	requestDef := GenReqDefForListMultiAccountTransferCoupons()
	return &ListMultiAccountTransferCouponsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOnDemandResourceRatings 查询按需产品价格
//
// 伙伴在销售平台按照条件查询按需产品的价格。
//
// 如果购买该产品的租户享受折扣，可以在查询结果中返回折扣金额以及扣除折扣后的最后成交价。
//
// 如果该租户享受多种折扣，系统会优先返回客户享受的商务折扣的折扣金额和最终成交价。
//
// &gt;![](public_sys-resources/icon-caution.gif) **注意：**
// &gt;华为云根据云服务类型、资源类型、云服务区和资源规格四个条件来查询产品，查询时请确认这4个查询条件均输入正确，否则该接口会返回无法找到产品的错误。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListOnDemandResourceRatings(request *model.ListOnDemandResourceRatingsRequest) (*model.ListOnDemandResourceRatingsResponse, error) {
	requestDef := GenReqDefForListOnDemandResourceRatings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOnDemandResourceRatingsResponse), nil
	}
}

// ListOnDemandResourceRatingsInvoker 查询按需产品价格
func (c *BssClient) ListOnDemandResourceRatingsInvoker(request *model.ListOnDemandResourceRatingsRequest) *ListOnDemandResourceRatingsInvoker {
	requestDef := GenReqDefForListOnDemandResourceRatings()
	return &ListOnDemandResourceRatingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOrderCouponsByOrderId 查询订单可用优惠券
//
// 客户在伙伴销售平台支付待支付订单时，查询可使用的优惠券列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListOrderCouponsByOrderId(request *model.ListOrderCouponsByOrderIdRequest) (*model.ListOrderCouponsByOrderIdResponse, error) {
	requestDef := GenReqDefForListOrderCouponsByOrderId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrderCouponsByOrderIdResponse), nil
	}
}

// ListOrderCouponsByOrderIdInvoker 查询订单可用优惠券
func (c *BssClient) ListOrderCouponsByOrderIdInvoker(request *model.ListOrderCouponsByOrderIdRequest) *ListOrderCouponsByOrderIdInvoker {
	requestDef := GenReqDefForListOrderCouponsByOrderId()
	return &ListOrderCouponsByOrderIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOrderDiscounts 查询订单可用折扣
//
// 客户在伙伴销售平台支付待支付订单时，查询可使用的折扣信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListOrderDiscounts(request *model.ListOrderDiscountsRequest) (*model.ListOrderDiscountsResponse, error) {
	requestDef := GenReqDefForListOrderDiscounts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrderDiscountsResponse), nil
	}
}

// ListOrderDiscountsInvoker 查询订单可用折扣
func (c *BssClient) ListOrderDiscountsInvoker(request *model.ListOrderDiscountsRequest) *ListOrderDiscountsInvoker {
	requestDef := GenReqDefForListOrderDiscounts()
	return &ListOrderDiscountsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPartnerAccountChangeRecords 查询收支明细
//
// 伙伴在伙伴销售平台上查询自身的收支明细情况。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListPartnerAccountChangeRecords(request *model.ListPartnerAccountChangeRecordsRequest) (*model.ListPartnerAccountChangeRecordsResponse, error) {
	requestDef := GenReqDefForListPartnerAccountChangeRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPartnerAccountChangeRecordsResponse), nil
	}
}

// ListPartnerAccountChangeRecordsInvoker 查询收支明细
func (c *BssClient) ListPartnerAccountChangeRecordsInvoker(request *model.ListPartnerAccountChangeRecordsRequest) *ListPartnerAccountChangeRecordsInvoker {
	requestDef := GenReqDefForListPartnerAccountChangeRecords()
	return &ListPartnerAccountChangeRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPartnerAdjustRecords 查询调账记录
//
// 伙伴在伙伴销售平台查询向客户及关联的云经销商（二级经销商）拨款或回收的调账记录。
//
// 伙伴登录伙伴中心，在“拨款”或“回收”页面，单击“调账记录”，可以查看一级经销商为二级经销商调账的记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListPartnerAdjustRecords(request *model.ListPartnerAdjustRecordsRequest) (*model.ListPartnerAdjustRecordsResponse, error) {
	requestDef := GenReqDefForListPartnerAdjustRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPartnerAdjustRecordsResponse), nil
	}
}

// ListPartnerAdjustRecordsInvoker 查询调账记录
func (c *BssClient) ListPartnerAdjustRecordsInvoker(request *model.ListPartnerAdjustRecordsRequest) *ListPartnerAdjustRecordsInvoker {
	requestDef := GenReqDefForListPartnerAdjustRecords()
	return &ListPartnerAdjustRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPartnerBalances 查询云经销商账户余额
//
// 华为云总经销商（一级经销商）可以查询关联的云经销商（二级经销商）的账户余额；云经销商伙伴可以查询自己的账户余额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListPartnerBalances(request *model.ListPartnerBalancesRequest) (*model.ListPartnerBalancesResponse, error) {
	requestDef := GenReqDefForListPartnerBalances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPartnerBalancesResponse), nil
	}
}

// ListPartnerBalancesInvoker 查询云经销商账户余额
func (c *BssClient) ListPartnerBalancesInvoker(request *model.ListPartnerBalancesRequest) *ListPartnerBalancesInvoker {
	requestDef := GenReqDefForListPartnerBalances()
	return &ListPartnerBalancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPartnerCouponsRecord 查询优惠券的发放回收记录
//
// 合作伙伴可查看给客户发放和回收优惠券的操作记录。
//
// 合作伙伴登录伙伴中心查看、导出代金券操作日志请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435103.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListPartnerCouponsRecord(request *model.ListPartnerCouponsRecordRequest) (*model.ListPartnerCouponsRecordResponse, error) {
	requestDef := GenReqDefForListPartnerCouponsRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPartnerCouponsRecordResponse), nil
	}
}

// ListPartnerCouponsRecordInvoker 查询优惠券的发放回收记录
func (c *BssClient) ListPartnerCouponsRecordInvoker(request *model.ListPartnerCouponsRecordRequest) *ListPartnerCouponsRecordInvoker {
	requestDef := GenReqDefForListPartnerCouponsRecord()
	return &ListPartnerCouponsRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPayPerUseCustomerResources 查询客户包年/包月资源列表
//
// 客户在伙伴销售平台查询某个或所有的包年/包月资源。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;成功调用本接口后，如果您需要对已生效状态的资源进行续订，您可以调用“[查询包年/包月产品价格](https://support.huaweicloud.com/api-bpconsole/bcloud_01002.html)”接口对查询到的包年/包月资源进行续订询价，然后再调用“[续订包年/包月资源](https://support.huaweicloud.com/api-bpconsole/api_order_00018.html)”接口进行续订。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListPayPerUseCustomerResources(request *model.ListPayPerUseCustomerResourcesRequest) (*model.ListPayPerUseCustomerResourcesResponse, error) {
	requestDef := GenReqDefForListPayPerUseCustomerResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPayPerUseCustomerResourcesResponse), nil
	}
}

// ListPayPerUseCustomerResourcesInvoker 查询客户包年/包月资源列表
func (c *BssClient) ListPayPerUseCustomerResourcesInvoker(request *model.ListPayPerUseCustomerResourcesRequest) *ListPayPerUseCustomerResourcesInvoker {
	requestDef := GenReqDefForListPayPerUseCustomerResources()
	return &ListPayPerUseCustomerResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProvinces 查询省份信息
//
// 伙伴在伙伴销售平台上查询省份信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListProvinces(request *model.ListProvincesRequest) (*model.ListProvincesResponse, error) {
	requestDef := GenReqDefForListProvinces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProvincesResponse), nil
	}
}

// ListProvincesInvoker 查询省份信息
func (c *BssClient) ListProvincesInvoker(request *model.ListProvincesRequest) *ListProvincesInvoker {
	requestDef := GenReqDefForListProvinces()
	return &ListProvincesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotaCoupons 查询优惠券额度
//
// 合作伙伴可以查看所拥有的优惠劵额度。
//
// 伙伴登录合作伙伴中心查看所拥有的代金券额度请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435100.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListQuotaCoupons(request *model.ListQuotaCouponsRequest) (*model.ListQuotaCouponsResponse, error) {
	requestDef := GenReqDefForListQuotaCoupons()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotaCouponsResponse), nil
	}
}

// ListQuotaCouponsInvoker 查询优惠券额度
func (c *BssClient) ListQuotaCouponsInvoker(request *model.ListQuotaCouponsRequest) *ListQuotaCouponsInvoker {
	requestDef := GenReqDefForListQuotaCoupons()
	return &ListQuotaCouponsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRateOnPeriodDetail 查询包年/包月产品价格
//
// 伙伴在销售平台按照条件查询包年/包月产品开通时候的价格。
//
// 如果购买该产品的客户享受折扣，可以在查询结果中返回折扣金额以及扣除折扣后的最后成交价。
//
// 如果该客户享受多种折扣，系统会返回每种折扣的批价结果。如果客户在下单的时候选择自动支付，则系统会优先应用商务折扣的批价结果。
//
// &gt;![](public_sys-resources/icon-caution.gif) **注意：**
// &gt;华为云根据云服务类型、资源类型、云服务区和资源规格四个条件来查询产品，查询时请确认这4个查询条件均输入正确，否则该接口会返回无法找到产品的错误。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListRateOnPeriodDetail(request *model.ListRateOnPeriodDetailRequest) (*model.ListRateOnPeriodDetailResponse, error) {
	requestDef := GenReqDefForListRateOnPeriodDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRateOnPeriodDetailResponse), nil
	}
}

// ListRateOnPeriodDetailInvoker 查询包年/包月产品价格
func (c *BssClient) ListRateOnPeriodDetailInvoker(request *model.ListRateOnPeriodDetailRequest) *ListRateOnPeriodDetailInvoker {
	requestDef := GenReqDefForListRateOnPeriodDetail()
	return &ListRateOnPeriodDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRenewRateOnPeriod 查询待续订包年包月资源的续订金额
//
// 功能描述：客户在自建平台按照条件查询待续订包年/包月资源续订时候的续订金额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListRenewRateOnPeriod(request *model.ListRenewRateOnPeriodRequest) (*model.ListRenewRateOnPeriodResponse, error) {
	requestDef := GenReqDefForListRenewRateOnPeriod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRenewRateOnPeriodResponse), nil
	}
}

// ListRenewRateOnPeriodInvoker 查询待续订包年包月资源的续订金额
func (c *BssClient) ListRenewRateOnPeriodInvoker(request *model.ListRenewRateOnPeriodRequest) *ListRenewRateOnPeriodInvoker {
	requestDef := GenReqDefForListRenewRateOnPeriod()
	return &ListRenewRateOnPeriodInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceTypes 查询资源类型列表
//
// 伙伴在伙伴销售平台查询资源类型的列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListResourceTypes(request *model.ListResourceTypesRequest) (*model.ListResourceTypesResponse, error) {
	requestDef := GenReqDefForListResourceTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceTypesResponse), nil
	}
}

// ListResourceTypesInvoker 查询资源类型列表
func (c *BssClient) ListResourceTypesInvoker(request *model.ListResourceTypesRequest) *ListResourceTypesInvoker {
	requestDef := GenReqDefForListResourceTypes()
	return &ListResourceTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceUsage 查询95计费资源用量明细
//
// 客户在自建平台查询自己的资源使用量明细。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;当前仅支持查询CDN、OBS、IEC和VPC四种云服务类型的资源用量明细，仅针对95计费场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListResourceUsage(request *model.ListResourceUsageRequest) (*model.ListResourceUsageResponse, error) {
	requestDef := GenReqDefForListResourceUsage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceUsageResponse), nil
	}
}

// ListResourceUsageInvoker 查询95计费资源用量明细
func (c *BssClient) ListResourceUsageInvoker(request *model.ListResourceUsageRequest) *ListResourceUsageInvoker {
	requestDef := GenReqDefForListResourceUsage()
	return &ListResourceUsageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceUsageSummary 查询95计费资源用量汇总
//
// 客户在自建平台查询自己的资源使用量汇总。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;-   当前仅支持查询CDN、OBS、IEC和VPC四种云服务类型的资源用量汇总，仅针对95计费场景。
// &gt;-   使用量汇总列表只包含月汇总金额和资源ID，若要查询具体某个资源的用量明细，请调用[查询资源用量明细](https://support.huaweicloud.com/api-oce/zh-cn_topic_0000001190606757.html)接口获取。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListResourceUsageSummary(request *model.ListResourceUsageSummaryRequest) (*model.ListResourceUsageSummaryResponse, error) {
	requestDef := GenReqDefForListResourceUsageSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceUsageSummaryResponse), nil
	}
}

// ListResourceUsageSummaryInvoker 查询95计费资源用量汇总
func (c *BssClient) ListResourceUsageSummaryInvoker(request *model.ListResourceUsageSummaryRequest) *ListResourceUsageSummaryInvoker {
	requestDef := GenReqDefForListResourceUsageSummary()
	return &ListResourceUsageSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServiceResources 根据云服务类型查询资源列表
//
// 伙伴在伙伴销售平台根据云服务类型查询关联的资源类型编码和名称，用于查询按需产品的价格或包年/包月产品的价格。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListServiceResources(request *model.ListServiceResourcesRequest) (*model.ListServiceResourcesResponse, error) {
	requestDef := GenReqDefForListServiceResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceResourcesResponse), nil
	}
}

// ListServiceResourcesInvoker 根据云服务类型查询资源列表
func (c *BssClient) ListServiceResourcesInvoker(request *model.ListServiceResourcesRequest) *ListServiceResourcesInvoker {
	requestDef := GenReqDefForListServiceResources()
	return &ListServiceResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServiceTypes 查询云服务类型列表
//
// 伙伴在伙伴销售平台查询云服务类型的列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListServiceTypes(request *model.ListServiceTypesRequest) (*model.ListServiceTypesResponse, error) {
	requestDef := GenReqDefForListServiceTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceTypesResponse), nil
	}
}

// ListServiceTypesInvoker 查询云服务类型列表
func (c *BssClient) ListServiceTypesInvoker(request *model.ListServiceTypesRequest) *ListServiceTypesInvoker {
	requestDef := GenReqDefForListServiceTypes()
	return &ListServiceTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStoredValueCards 查询储值卡列表
//
// 客户可以查询已购买的储值卡列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListStoredValueCards(request *model.ListStoredValueCardsRequest) (*model.ListStoredValueCardsResponse, error) {
	requestDef := GenReqDefForListStoredValueCards()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStoredValueCardsResponse), nil
	}
}

// ListStoredValueCardsInvoker 查询储值卡列表
func (c *BssClient) ListStoredValueCardsInvoker(request *model.ListStoredValueCardsRequest) *ListStoredValueCardsInvoker {
	requestDef := GenReqDefForListStoredValueCards()
	return &ListStoredValueCardsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubCustomerBillDetail 查询伙伴子客户消费记录
//
// 伙伴在伙伴销售平台可实时查询子客户的消费记录，了解客户的资源消耗情况。
//
// 伙伴在伙伴中心查询客户消费明细请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435155.html)。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;-   消费记录支持查询18个月内的记录。
// &gt;-   如果是客户经理主管来查询，只支持按照单个客户经理查询，必须输入客户经理ID。
// &gt;-   目前支持伙伴查询所有子客户（包含代售类和顾问销售类）的消费记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListSubCustomerBillDetail(request *model.ListSubCustomerBillDetailRequest) (*model.ListSubCustomerBillDetailResponse, error) {
	requestDef := GenReqDefForListSubCustomerBillDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubCustomerBillDetailResponse), nil
	}
}

// ListSubCustomerBillDetailInvoker 查询伙伴子客户消费记录
func (c *BssClient) ListSubCustomerBillDetailInvoker(request *model.ListSubCustomerBillDetailRequest) *ListSubCustomerBillDetailInvoker {
	requestDef := GenReqDefForListSubCustomerBillDetail()
	return &ListSubCustomerBillDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubCustomerCoupons 查询优惠券列表
//
// 伙伴可以查询自身的优惠券信息。
//
// 伙伴登录伙伴中心查询已发放代金券列表请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435101.html)，查看已下发代金券的内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListSubCustomerCoupons(request *model.ListSubCustomerCouponsRequest) (*model.ListSubCustomerCouponsResponse, error) {
	requestDef := GenReqDefForListSubCustomerCoupons()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubCustomerCouponsResponse), nil
	}
}

// ListSubCustomerCouponsInvoker 查询优惠券列表
func (c *BssClient) ListSubCustomerCouponsInvoker(request *model.ListSubCustomerCouponsRequest) *ListSubCustomerCouponsInvoker {
	requestDef := GenReqDefForListSubCustomerCoupons()
	return &ListSubCustomerCouponsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubCustomerNewTag 查询客户新客标签
//
// 伙伴通过该接口可以查询客户的新客标签。
//
// &gt;![](public_sys-resources/icon-caution.gif) **注意：**
// &gt;-   新客标签失效后，new\\_customer\\_tag会变成N，且有效期过期。
// &gt;-   客户如果没有实名认证，则新客标签为空。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListSubCustomerNewTag(request *model.ListSubCustomerNewTagRequest) (*model.ListSubCustomerNewTagResponse, error) {
	requestDef := GenReqDefForListSubCustomerNewTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubCustomerNewTagResponse), nil
	}
}

// ListSubCustomerNewTagInvoker 查询客户新客标签
func (c *BssClient) ListSubCustomerNewTagInvoker(request *model.ListSubCustomerNewTagRequest) *ListSubCustomerNewTagInvoker {
	requestDef := GenReqDefForListSubCustomerNewTag()
	return &ListSubCustomerNewTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubCustomers 查询客户列表
//
// 伙伴可以查询合作伙伴的客户信息列表。
//
// 伙伴登录合作伙伴中心查询客户信息列表请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435115.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListSubCustomers(request *model.ListSubCustomersRequest) (*model.ListSubCustomersResponse, error) {
	requestDef := GenReqDefForListSubCustomers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubCustomersResponse), nil
	}
}

// ListSubCustomersInvoker 查询客户列表
func (c *BssClient) ListSubCustomersInvoker(request *model.ListSubCustomersRequest) *ListSubCustomersInvoker {
	requestDef := GenReqDefForListSubCustomers()
	return &ListSubCustomersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubcustomerMonthlyBills 查询客户月度消费账单
//
// 合作伙伴可查询客户的消费汇总账单，消费按月汇总。
//
// 伙伴在伙伴中心查询客户月度消费账单请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435154.html)。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;消费汇总数据仅包含前一天24点前的数据。每天刷新一次，更新前一天的数据。
// &gt;该接口用于合作伙伴查询其代售类客户在华为的消费情况，如果输入某个客户ID，则是查询单个客户的，否则是查询该伙伴下所有使用伙伴拨款消费的客户的消费记录（包括退订记录）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListSubcustomerMonthlyBills(request *model.ListSubcustomerMonthlyBillsRequest) (*model.ListSubcustomerMonthlyBillsResponse, error) {
	requestDef := GenReqDefForListSubcustomerMonthlyBills()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubcustomerMonthlyBillsResponse), nil
	}
}

// ListSubcustomerMonthlyBillsInvoker 查询客户月度消费账单
func (c *BssClient) ListSubcustomerMonthlyBillsInvoker(request *model.ListSubcustomerMonthlyBillsRequest) *ListSubcustomerMonthlyBillsInvoker {
	requestDef := GenReqDefForListSubcustomerMonthlyBills()
	return &ListSubcustomerMonthlyBillsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsageTypes 查询使用量类型列表
//
// 伙伴在伙伴销售平台查询资源的使用量类型列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ListUsageTypes(request *model.ListUsageTypesRequest) (*model.ListUsageTypesResponse, error) {
	requestDef := GenReqDefForListUsageTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsageTypesResponse), nil
	}
}

// ListUsageTypesInvoker 查询使用量类型列表
func (c *BssClient) ListUsageTypesInvoker(request *model.ListUsageTypesRequest) *ListUsageTypesInvoker {
	requestDef := GenReqDefForListUsageTypes()
	return &ListUsageTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PayOrders 支付包年/包月产品订单
//
// 客户可以对待支付状态的包年/包月产品订单进行支付。
//
// 客户登录费用中心支付包年包月产品的待支付订单请单击[这里](https://support.huaweicloud.com/usermanual-billing/zh-cn_topic_0031512547.html)。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;-   API支持月度结算和余额支付两种支付方式，月度结算优先。
// &gt;-   余额支付包括现金账户和信用账户两种支付方式，如果两个账户都有余额，则优先现金账户支付。
// &gt;-   同时使用订单折扣和优惠券的互斥规则如下：
// &gt;    -   如果优惠券的限制属性上存在simultaneousUseWithEmpowerDiscount字段，并且值为0，则折扣和优惠券不能同时使用。
// &gt;    -   如果优惠券的限制属性上存在minConsumeDiscount字段，当折扣ID包含的所有订单项上的折扣率discount\\_ratio都小于minConsumeDiscount字段时，则折扣ID和优惠券不能同时使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) PayOrders(request *model.PayOrdersRequest) (*model.PayOrdersResponse, error) {
	requestDef := GenReqDefForPayOrders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PayOrdersResponse), nil
	}
}

// PayOrdersInvoker 支付包年/包月产品订单
func (c *BssClient) PayOrdersInvoker(request *model.PayOrdersRequest) *PayOrdersInvoker {
	requestDef := GenReqDefForPayOrders()
	return &PayOrdersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ReclaimCouponQuotas 回收云经销商的代金券额度
//
// 华为云总经销商（一级经销商）可以回收已发放给云经销商（二级经销商）的代金券额度。
//
// 一级经销商在伙伴中心回收已发放给二级经销商的代金券额度请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/dp_120206.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ReclaimCouponQuotas(request *model.ReclaimCouponQuotasRequest) (*model.ReclaimCouponQuotasResponse, error) {
	requestDef := GenReqDefForReclaimCouponQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReclaimCouponQuotasResponse), nil
	}
}

// ReclaimCouponQuotasInvoker 回收云经销商的代金券额度
func (c *BssClient) ReclaimCouponQuotasInvoker(request *model.ReclaimCouponQuotasRequest) *ReclaimCouponQuotasInvoker {
	requestDef := GenReqDefForReclaimCouponQuotas()
	return &ReclaimCouponQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ReclaimEnterpriseMultiAccountCoupon 企业主账号从企业子账号回收优惠券
//
// 企业主账号在自建平台回收给企业子账号的拨款优惠券。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;-   仅支持华为发放的测试类、商务类、活动类代金券。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ReclaimEnterpriseMultiAccountCoupon(request *model.ReclaimEnterpriseMultiAccountCouponRequest) (*model.ReclaimEnterpriseMultiAccountCouponResponse, error) {
	requestDef := GenReqDefForReclaimEnterpriseMultiAccountCoupon()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReclaimEnterpriseMultiAccountCouponResponse), nil
	}
}

// ReclaimEnterpriseMultiAccountCouponInvoker 企业主账号从企业子账号回收优惠券
func (c *BssClient) ReclaimEnterpriseMultiAccountCouponInvoker(request *model.ReclaimEnterpriseMultiAccountCouponRequest) *ReclaimEnterpriseMultiAccountCouponInvoker {
	requestDef := GenReqDefForReclaimEnterpriseMultiAccountCoupon()
	return &ReclaimEnterpriseMultiAccountCouponInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ReclaimIndirectPartnerAccount 回收云经销商账户拨款
//
// 华为云总经销商（一级经销商）可以回收云经销商（二级经销商）的账户余额。
//
// 一级经销商在伙伴中心回收二级经销商账户拨款请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/dp_120205.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ReclaimIndirectPartnerAccount(request *model.ReclaimIndirectPartnerAccountRequest) (*model.ReclaimIndirectPartnerAccountResponse, error) {
	requestDef := GenReqDefForReclaimIndirectPartnerAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReclaimIndirectPartnerAccountResponse), nil
	}
}

// ReclaimIndirectPartnerAccountInvoker 回收云经销商账户拨款
func (c *BssClient) ReclaimIndirectPartnerAccountInvoker(request *model.ReclaimIndirectPartnerAccountRequest) *ReclaimIndirectPartnerAccountInvoker {
	requestDef := GenReqDefForReclaimIndirectPartnerAccount()
	return &ReclaimIndirectPartnerAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ReclaimPartnerCoupons 回收优惠券
//
// 对于合作伙伴已经下发给客户的优惠券，如遇发错或其他特殊情况，合作伙伴有回收的权利。优惠券回收后，客户将不再拥有该优惠券。
//
// 伙伴登录合作伙伴中心回收为客户发放的代金券请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/espp_050503.html)。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;只能回收代售类子客户的优惠券。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ReclaimPartnerCoupons(request *model.ReclaimPartnerCouponsRequest) (*model.ReclaimPartnerCouponsResponse, error) {
	requestDef := GenReqDefForReclaimPartnerCoupons()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReclaimPartnerCouponsResponse), nil
	}
}

// ReclaimPartnerCouponsInvoker 回收优惠券
func (c *BssClient) ReclaimPartnerCouponsInvoker(request *model.ReclaimPartnerCouponsRequest) *ReclaimPartnerCouponsInvoker {
	requestDef := GenReqDefForReclaimPartnerCoupons()
	return &ReclaimPartnerCouponsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ReclaimSubEnterpriseAmount 企业主账号从企业子账号回收拨款
//
// 企业主账号在自建平台回收给企业子账号的拨款。
//
// 如果回收的是企业子账户的信用账户，可以回收所有额度；如果回收金额大于子账户信用余额的时候，可能会导致子账户欠费，请慎重选择。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ReclaimSubEnterpriseAmount(request *model.ReclaimSubEnterpriseAmountRequest) (*model.ReclaimSubEnterpriseAmountResponse, error) {
	requestDef := GenReqDefForReclaimSubEnterpriseAmount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReclaimSubEnterpriseAmountResponse), nil
	}
}

// ReclaimSubEnterpriseAmountInvoker 企业主账号从企业子账号回收拨款
func (c *BssClient) ReclaimSubEnterpriseAmountInvoker(request *model.ReclaimSubEnterpriseAmountRequest) *ReclaimSubEnterpriseAmountInvoker {
	requestDef := GenReqDefForReclaimSubEnterpriseAmount()
	return &ReclaimSubEnterpriseAmountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ReclaimToPartnerAccount 回收客户账户余额
//
// 当客户不再使用华为云产品时，合作伙伴可以回收代售类客户账户余额。
//
// 伙伴登录伙伴中心回收代售类客户账户余额请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435147.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ReclaimToPartnerAccount(request *model.ReclaimToPartnerAccountRequest) (*model.ReclaimToPartnerAccountResponse, error) {
	requestDef := GenReqDefForReclaimToPartnerAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReclaimToPartnerAccountResponse), nil
	}
}

// ReclaimToPartnerAccountInvoker 回收客户账户余额
func (c *BssClient) ReclaimToPartnerAccountInvoker(request *model.ReclaimToPartnerAccountRequest) *ReclaimToPartnerAccountInvoker {
	requestDef := GenReqDefForReclaimToPartnerAccount()
	return &ReclaimToPartnerAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RenewalResources 续订包年/包月资源
//
// 客户的包年/包月资源即将到期时，可进行包年/包月资源的续订。
//
// 客户在费用中心执行续订操作请参见[这里](https://support.huaweicloud.com/usermanual-billing/renewals_topic_10000003.html)。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;-   调用接口后，如果某个主资源有对应的从资源，系统会将主资源和从资源一起续订，主资源的从资源信息可以通过调用[查询客户包年/包月资源列表](https://support.huaweicloud.com/api-oce/api_order_00021.html)接口获取。
// &gt;-   注意：如ECS主机挂载新购的云硬盘，但此硬盘不是该ECS主资源的从资源，主从资源信息必须以调用[查询客户包年/包月资源列表](https://support.huaweicloud.com/api-oce/api_order_00021.html)接口获取的信息为准。
// &gt;-   本接口支持自动支付，支付时使用折扣或优惠券的说明，请参见[支付使用折扣或优惠券说明](https://support.huaweicloud.com/api-oce/appendix_00001.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) RenewalResources(request *model.RenewalResourcesRequest) (*model.RenewalResourcesResponse, error) {
	requestDef := GenReqDefForRenewalResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RenewalResourcesResponse), nil
	}
}

// RenewalResourcesInvoker 续订包年/包月资源
func (c *BssClient) RenewalResourcesInvoker(request *model.RenewalResourcesRequest) *RenewalResourcesInvoker {
	requestDef := GenReqDefForRenewalResources()
	return &RenewalResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendSmsVerificationCode 发送短信验证码
//
// 企业主账号在自建平台发送短信验证码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) SendSmsVerificationCode(request *model.SendSmsVerificationCodeRequest) (*model.SendSmsVerificationCodeResponse, error) {
	requestDef := GenReqDefForSendSmsVerificationCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendSmsVerificationCodeResponse), nil
	}
}

// SendSmsVerificationCodeInvoker 发送短信验证码
func (c *BssClient) SendSmsVerificationCodeInvoker(request *model.SendSmsVerificationCodeRequest) *SendSmsVerificationCodeInvoker {
	requestDef := GenReqDefForSendSmsVerificationCode()
	return &SendSmsVerificationCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendVerificationMessageCode 发送验证码
//
// 客户注册时，如果填写了手机号，可以向对应的手机发送注册验证码，校验信息的正确性。使用个人银行卡方式进行实名认证时，通过该接口向指定的手机发送验证码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) SendVerificationMessageCode(request *model.SendVerificationMessageCodeRequest) (*model.SendVerificationMessageCodeResponse, error) {
	requestDef := GenReqDefForSendVerificationMessageCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendVerificationMessageCodeResponse), nil
	}
}

// SendVerificationMessageCodeInvoker 发送验证码
func (c *BssClient) SendVerificationMessageCodeInvoker(request *model.SendVerificationMessageCodeRequest) *SendVerificationMessageCodeInvoker {
	requestDef := GenReqDefForSendVerificationMessageCode()
	return &SendVerificationMessageCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCustomerAccountBalances 查询账户余额
//
// 客户可以查询自身的账户余额。
//
// 客户可以登录费用中心进入“[总览](https://account.huaweicloud.com/usercenter/#/userindex/allview)”页面，在“可用额度”区域可以查询自身的账户余额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ShowCustomerAccountBalances(request *model.ShowCustomerAccountBalancesRequest) (*model.ShowCustomerAccountBalancesResponse, error) {
	requestDef := GenReqDefForShowCustomerAccountBalances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCustomerAccountBalancesResponse), nil
	}
}

// ShowCustomerAccountBalancesInvoker 查询账户余额
func (c *BssClient) ShowCustomerAccountBalancesInvoker(request *model.ShowCustomerAccountBalancesRequest) *ShowCustomerAccountBalancesInvoker {
	requestDef := GenReqDefForShowCustomerAccountBalances()
	return &ShowCustomerAccountBalancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCustomerMonthlySum 查询汇总账单
//
// 客户在自建平台查询自身的消费汇总账单，此账单按月汇总消费数据。
//
// 客户登录费用中心查询自身的消费汇总账单请参见[这里](https://support.huaweicloud.com/usermanual-billing/bills-topic_80000001.html#bills-topic_80000001__zh-cn_topic_0000001162496407_s620ce713baf04899a416d781d1817931)的“**查看汇总**”。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;当前支持查看2019/01月份至今的费用账单。企业主账号展示的费用账单，包含关联的统一还款企业子账号的消费数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ShowCustomerMonthlySum(request *model.ShowCustomerMonthlySumRequest) (*model.ShowCustomerMonthlySumResponse, error) {
	requestDef := GenReqDefForShowCustomerMonthlySum()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCustomerMonthlySumResponse), nil
	}
}

// ShowCustomerMonthlySumInvoker 查询汇总账单
func (c *BssClient) ShowCustomerMonthlySumInvoker(request *model.ShowCustomerMonthlySumRequest) *ShowCustomerMonthlySumInvoker {
	requestDef := GenReqDefForShowCustomerMonthlySum()
	return &ShowCustomerMonthlySumInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCustomerOrderDetails 查询订单详情
//
// 客户可以在伙伴销售平台查看订单详情。
//
// 客户登录费用中心查看订单详情请单击[这里](https://support.huaweicloud.com/usermanual-billing/order_topic_9000001.html)。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;如果想查询某条订单下的资源信息，请调用“[查询客户包年/包月资源列表](https://support.huaweicloud.com/api-oce/api_order_00021.html)”接口在请求参数输入订单号进行查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ShowCustomerOrderDetails(request *model.ShowCustomerOrderDetailsRequest) (*model.ShowCustomerOrderDetailsResponse, error) {
	requestDef := GenReqDefForShowCustomerOrderDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCustomerOrderDetailsResponse), nil
	}
}

// ShowCustomerOrderDetailsInvoker 查询订单详情
func (c *BssClient) ShowCustomerOrderDetailsInvoker(request *model.ShowCustomerOrderDetailsRequest) *ShowCustomerOrderDetailsInvoker {
	requestDef := GenReqDefForShowCustomerOrderDetails()
	return &ShowCustomerOrderDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMultiAccountTransferAmount 查询企业主账号可拨款余额
//
// 企业主账号在自建平台查询自己的可拨款余额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ShowMultiAccountTransferAmount(request *model.ShowMultiAccountTransferAmountRequest) (*model.ShowMultiAccountTransferAmountResponse, error) {
	requestDef := GenReqDefForShowMultiAccountTransferAmount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMultiAccountTransferAmountResponse), nil
	}
}

// ShowMultiAccountTransferAmountInvoker 查询企业主账号可拨款余额
func (c *BssClient) ShowMultiAccountTransferAmountInvoker(request *model.ShowMultiAccountTransferAmountRequest) *ShowMultiAccountTransferAmountInvoker {
	requestDef := GenReqDefForShowMultiAccountTransferAmount()
	return &ShowMultiAccountTransferAmountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRealnameAuthenticationReviewResult 查询实名认证审核结果
//
// 如果实名认证申请或实名认证变更申请的响应中，显示需要人工审核，使用该接口查询审核结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ShowRealnameAuthenticationReviewResult(request *model.ShowRealnameAuthenticationReviewResultRequest) (*model.ShowRealnameAuthenticationReviewResultResponse, error) {
	requestDef := GenReqDefForShowRealnameAuthenticationReviewResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRealnameAuthenticationReviewResultResponse), nil
	}
}

// ShowRealnameAuthenticationReviewResultInvoker 查询实名认证审核结果
func (c *BssClient) ShowRealnameAuthenticationReviewResultInvoker(request *model.ShowRealnameAuthenticationReviewResultRequest) *ShowRealnameAuthenticationReviewResultInvoker {
	requestDef := GenReqDefForShowRealnameAuthenticationReviewResult()
	return &ShowRealnameAuthenticationReviewResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRefundOrderDetails 查询退款订单的金额详情
//
// 客户在伙伴销售平台查询某次退订订单或者降配订单的退款金额来自哪些资源和对应订单。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;-   可以在调用完“[退订包年/包月资源](https://support.huaweicloud.com/api-oce/api_order_00019.html)”接口生成退订订单ID后，调用该接口查询退订订单对应的金额所属资源和订单。例如，调用“[退订包年/包月资源](https://support.huaweicloud.com/api-oce/api_order_00019.html)”接口退订资源及其已续费周期后，您可以调用本小节的接口查询到退订金额归属的原开通订单ID和原续费订单ID。
// &gt;-   2018年5月份之后退订的订单才能查询到归属的原订单ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) ShowRefundOrderDetails(request *model.ShowRefundOrderDetailsRequest) (*model.ShowRefundOrderDetailsResponse, error) {
	requestDef := GenReqDefForShowRefundOrderDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRefundOrderDetailsResponse), nil
	}
}

// ShowRefundOrderDetailsInvoker 查询退款订单的金额详情
func (c *BssClient) ShowRefundOrderDetailsInvoker(request *model.ShowRefundOrderDetailsRequest) *ShowRefundOrderDetailsInvoker {
	requestDef := GenReqDefForShowRefundOrderDetails()
	return &ShowRefundOrderDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCouponQuotas 向云经销商发放代金券额度
//
// 华为云总经销商（一级经销商）可以向云经销商（二级经销商）发放代金券额度。
//
// 一级经销商在伙伴中心向二级经销商发放代金券额度请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/dp_120206.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) UpdateCouponQuotas(request *model.UpdateCouponQuotasRequest) (*model.UpdateCouponQuotasResponse, error) {
	requestDef := GenReqDefForUpdateCouponQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCouponQuotasResponse), nil
	}
}

// UpdateCouponQuotasInvoker 向云经销商发放代金券额度
func (c *BssClient) UpdateCouponQuotasInvoker(request *model.UpdateCouponQuotasRequest) *UpdateCouponQuotasInvoker {
	requestDef := GenReqDefForUpdateCouponQuotas()
	return &UpdateCouponQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCustomerAccountAmount 向客户账户拨款
//
// 合作伙伴可以为代售类客户账户拨款。
//
// 伙伴登录伙伴中心为代售类客户账户拨款请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435147.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) UpdateCustomerAccountAmount(request *model.UpdateCustomerAccountAmountRequest) (*model.UpdateCustomerAccountAmountResponse, error) {
	requestDef := GenReqDefForUpdateCustomerAccountAmount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCustomerAccountAmountResponse), nil
	}
}

// UpdateCustomerAccountAmountInvoker 向客户账户拨款
func (c *BssClient) UpdateCustomerAccountAmountInvoker(request *model.UpdateCustomerAccountAmountRequest) *UpdateCustomerAccountAmountInvoker {
	requestDef := GenReqDefForUpdateCustomerAccountAmount()
	return &UpdateCustomerAccountAmountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIndirectPartnerAccount 向云经销商账户拨款
//
// 华为云总经销商（一级经销商）可以向云经销商（二级经销商）账户拨款。
//
// 一级经销商在伙伴中心向二级经销商拨款请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/dp_120205.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) UpdateIndirectPartnerAccount(request *model.UpdateIndirectPartnerAccountRequest) (*model.UpdateIndirectPartnerAccountResponse, error) {
	requestDef := GenReqDefForUpdateIndirectPartnerAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIndirectPartnerAccountResponse), nil
	}
}

// UpdateIndirectPartnerAccountInvoker 向云经销商账户拨款
func (c *BssClient) UpdateIndirectPartnerAccountInvoker(request *model.UpdateIndirectPartnerAccountRequest) *UpdateIndirectPartnerAccountInvoker {
	requestDef := GenReqDefForUpdateIndirectPartnerAccount()
	return &UpdateIndirectPartnerAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePeriodToOnDemand 设置或取消包年/包月资源到期转按需
//
// 客户可以设置包年/包月资源到期后转为按需资源计费。包年/包月计费模式到期后，按需的计费模式即生效。
//
// 客户在费用中心设置包年包月资源到期转按需请参见[这里](https://support.huaweicloud.com/usermanual-billing/renewals_topic_50000003.html)。
//
// &gt;![](public_sys-resources/icon-note.gif) **说明：**
// &gt;-   客户需要成功支付包年/包月资源订单后，才能设置资源的到期转按需。
// &gt;-   目前解决方案组合产品、按需套餐包不支持到期转按需。
// &gt;-   在调用本接口前，您可以调用“[查询客户包年/包月资源列表](https://support.huaweicloud.com/api-oce/api_order_00021.html)”接口获取资源ID、资源过期时间以及资源过期后的扣费策略等信息。
// &gt;-   设置包年/包月资源到期转按需后，包年/包月资源到期后将自动变成按需计费。
// &gt;-   取消包年/包月资源到期转按需的前提是已经调用“[设置或取消包年/包月资源到期转按需](https://support.huaweicloud.com/api-oce/api_order_00024.html)”接口设置包年/包月资源的到期转按需或在调用“[续订包年/包月资源](https://support.huaweicloud.com/api-oce/api_order_00018.html)”接口时设置到期策略为到期转按需。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) UpdatePeriodToOnDemand(request *model.UpdatePeriodToOnDemandRequest) (*model.UpdatePeriodToOnDemandResponse, error) {
	requestDef := GenReqDefForUpdatePeriodToOnDemand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePeriodToOnDemandResponse), nil
	}
}

// UpdatePeriodToOnDemandInvoker 设置或取消包年/包月资源到期转按需
func (c *BssClient) UpdatePeriodToOnDemandInvoker(request *model.UpdatePeriodToOnDemandRequest) *UpdatePeriodToOnDemandInvoker {
	requestDef := GenReqDefForUpdatePeriodToOnDemand()
	return &UpdatePeriodToOnDemandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSubEnterpriseAmount 企业主账号向企业子账号拨款
//
// 企业主账号在自建平台向企业子账号拨款。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssClient) UpdateSubEnterpriseAmount(request *model.UpdateSubEnterpriseAmountRequest) (*model.UpdateSubEnterpriseAmountResponse, error) {
	requestDef := GenReqDefForUpdateSubEnterpriseAmount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubEnterpriseAmountResponse), nil
	}
}

// UpdateSubEnterpriseAmountInvoker 企业主账号向企业子账号拨款
func (c *BssClient) UpdateSubEnterpriseAmountInvoker(request *model.UpdateSubEnterpriseAmountRequest) *UpdateSubEnterpriseAmountInvoker {
	requestDef := GenReqDefForUpdateSubEnterpriseAmount()
	return &UpdateSubEnterpriseAmountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
