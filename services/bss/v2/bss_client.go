package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"
)

type BssClient struct {
	HcClient *http_client.HcHttpClient
}

func NewBssClient(hcClient *http_client.HcHttpClient) *BssClient {
	return &BssClient{HcClient: hcClient}
}

func BssClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

//为防止资源到期被删除，客户可为长期使用的包年/包月资源开通自动续费。  客户在费用中心开通自动续费请参见[这里](https://support.huaweicloud.com/usermanual-billing/renewals_topic_20000003.html)。  >![](public_sys-resources/icon-note.gif) **说明：**  >-   首先要客户成功支付包年/包月资源订单，才能进行自动续费的开通。 >-   目前支持设置自动续费的包年/包月产品请参见[自动续费规则说明](https://support.huaweicloud.com/usermanual-billing/renewals_topic_20000002.html)。 >-   在调用本接口前，您可以调用“[查询客户包年/包月资源列表](查询客户包年-包月资源列表.md)”接口获取资源ID、资源过期时间以及资源过期后扣费策略等信息。 >-   自动续费将于产品到期前7天的凌晨3:00开始扣款，请保持账户余额充足。若由于账户中余额不足等原因导致第一次未扣费成功，系统将每天凌晨3:00尝试进行一次扣款，直到扣款成功或保留产品资源的最后一天。 >-   续费周期与原资源的购买周期一致。
func (c *BssClient) AutoRenewalResources(request *model.AutoRenewalResourcesRequest) (*model.AutoRenewalResourcesResponse, error) {
	requestDef := GenReqDefForAutoRenewalResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AutoRenewalResourcesResponse), nil
	}
}

//客户设置自动续费后，还可以执行取消自动续费的操作。关闭自动续费后，资源到期将不会被自动续费。  客户在费用中心取消包年/包月资源自动续费请参见[这里](https://support.huaweicloud.com/usermanual-billing/renewals_topic_20000005.html)。  >![](public_sys-resources/icon-note.gif) **说明：**  >-   前提是已经调用“[设置包年/包月资源自动续费](设置包年-包月资源自动续费.md)”接口设置自动续费或在调用“[续订包年/包月资源](续订包年-包月资源.md)”接口时设置到期策略为自动续订。 >-   目前支持取消自动续费的包年/包月产品同支持自动续费的包年/包月产品。 >-   在调用本接口前，您可以调用“[查询客户包年/包月资源列表](查询客户包年-包月资源列表.md)”接口获取资源ID、资源过期时间以及资源过期后扣费策略等信息。
func (c *BssClient) CancelAutoRenewalResources(request *model.CancelAutoRenewalResourcesRequest) (*model.CancelAutoRenewalResourcesResponse, error) {
	requestDef := GenReqDefForCancelAutoRenewalResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelAutoRenewalResourcesResponse), nil
	}
}

//客户可以对待支付的订单进行取消操作。  客户登录费用中心取消包年包月产品的待支付订单请单击[这里](https://support.huaweicloud.com/usermanual-billing/zh-cn_topic_0031465730.html)。  >![](public_sys-resources/icon-note.gif) **说明：**  >只有订单状态是“待支付”的时候，才能取消订单。
func (c *BssClient) CancelCustomerOrder(request *model.CancelCustomerOrderRequest) (*model.CancelCustomerOrderResponse, error) {
	requestDef := GenReqDefForCancelCustomerOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelCustomerOrderResponse), nil
	}
}

//客户购买包年/包月资源后，支持客户退订包年/包月实例。退订资源实例包括资源续费部分和当前正在使用的部分，退订后资源将无法使用。  客户在费用中心退订已购买的包年包月资源请参见[这里](https://support.huaweicloud.com/usermanual-billing/zh-cn_topic_0083138805.html)。  >![](public_sys-resources/icon-note.gif) **说明：**  >-   首先要成功支付包年/包月产品，产生一条开通成功的包年/包月资源，才能进行退订。 >-   调用接口后，如果某个主资源有对应的从资源，系统会将主资源和从资源一起退订，主资源的从资源信息可以通过调用[查询客户包年/包月资源列表](查询客户包年-包月资源列表.md)接口获取。 >-   注意：如ECS主机挂载新购的云硬盘，但此硬盘不是该ECS主资源的从资源，主从资源信息必须以调用[查询客户包年/包月资源列表](查询客户包年-包月资源列表.md)接口获取的信息为准。 >-   调用该接口后，您还可以调用“[查询退款订单的金额详情](查询退款订单的金额详情.md)”接口查询退订订单对应的金额来自哪些订单。 >-   该接口支持5天无理由全额退订，具体规则请参见“[退订规则说明](https://support.huaweicloud.com/usermanual-billing/unsubscription_topic_20000081.html)”。
func (c *BssClient) CancelResourcesSubscription(request *model.CancelResourcesSubscriptionRequest) (*model.CancelResourcesSubscriptionResponse, error) {
	requestDef := GenReqDefForCancelResourcesSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelResourcesSubscriptionResponse), nil
	}
}

//客户可以进行实名认证变更申请。  个人客户登录帐号中心通过实名认证变更为企业帐号的方式及流程请参见[这里](https://support.huaweicloud.com/usermanual-account/zh-cn_topic_0103532632.html)。
func (c *BssClient) ChangeEnterpriseRealnameAuthentication(request *model.ChangeEnterpriseRealnameAuthenticationRequest) (*model.ChangeEnterpriseRealnameAuthenticationResponse, error) {
	requestDef := GenReqDefForChangeEnterpriseRealnameAuthentication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeEnterpriseRealnameAuthenticationResponse), nil
	}
}

//客户注册时可检查客户的登录名称、手机号或者邮箱是否可以用于注册。  >![](public_sys-resources/icon-note.gif) **说明：**  >针对校验手机号场景，目前仅支持校验手机号注册数量是否超过上限。
func (c *BssClient) CheckUserIdentity(request *model.CheckUserIdentityRequest) (*model.CheckUserIdentityResponse, error) {
	requestDef := GenReqDefForCheckUserIdentity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckUserIdentityResponse), nil
	}
}

//客户在自建平台开通客户企业项目权限。
func (c *BssClient) CreateEnterpriseProjectAuth(request *model.CreateEnterpriseProjectAuthRequest) (*model.CreateEnterpriseProjectAuthResponse, error) {
	requestDef := GenReqDefForCreateEnterpriseProjectAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnterpriseProjectAuthResponse), nil
	}
}

//企业客户可以进行企业实名认证申请。  客户登录帐号中心进行企业实名认证的方式及流程请参见[这里](https://support.huaweicloud.com/usermanual-account/zh-cn_topic_0077914253.html)。
func (c *BssClient) CreateEnterpriseRealnameAuthentication(request *model.CreateEnterpriseRealnameAuthenticationRequest) (*model.CreateEnterpriseRealnameAuthenticationResponse, error) {
	requestDef := GenReqDefForCreateEnterpriseRealnameAuthentication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnterpriseRealnameAuthenticationResponse), nil
	}
}

//合作伙伴可以在拥有的代金券额度范围内为客户下发优惠券。  伙伴登录合作伙伴中心为客户发放代金券请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/espp_050502.html)。  >![](public_sys-resources/icon-note.gif) **说明：**  >只能给代售子客户发放优惠券。
func (c *BssClient) CreatePartnerCoupons(request *model.CreatePartnerCouponsRequest) (*model.CreatePartnerCouponsResponse, error) {
	requestDef := GenReqDefForCreatePartnerCoupons()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePartnerCouponsResponse), nil
	}
}

//个人客户可以进行个人实名认证申请。  客户登录帐号中心进行个人实名认证的方式及流程请参见[这里](https://support.huaweicloud.com/usermanual-account/zh-cn_topic_0077914254.html)。
func (c *BssClient) CreatePersonalRealnameAuth(request *model.CreatePersonalRealnameAuthRequest) (*model.CreatePersonalRealnameAuthResponse, error) {
	requestDef := GenReqDefForCreatePersonalRealnameAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePersonalRealnameAuthResponse), nil
	}
}

//伙伴可以新增自己的邮寄地址信息。  伙伴登录伙伴中心新增邮寄地址请参见[向华为云索取发票](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435143.html)，进入索取发票页面，选择纸质发票，即可设置邮件地址。
func (c *BssClient) CreatePostal(request *model.CreatePostalRequest) (*model.CreatePostalResponse, error) {
	requestDef := GenReqDefForCreatePostal()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePostalResponse), nil
	}
}

//在伙伴销售平台创建客户时同步创建华为云账号，并将客户在伙伴销售平台上的账号与华为云账号进行映射。同时，创建的华为云账号与伙伴账号关联绑定。  华为云伙伴能力中心（一级经销商）可以注册精英服务商伙伴（二级经销商）的子客户。注册完成后，子客户可以自动和精英服务商伙伴绑定。  >![](public_sys-resources/icon-caution.gif) **注意：**  >-   调用该接口为客户创建华为云账号后，如果想从合作伙伴销售平台跳转至华为云官网，还需要进行SAML认证，具体请参见“[Web UI方式](https://support.huaweicloud.com/api-bpconsole/jac_00001.html)”中的“SAML认证”。 >-   如果创建的时候不输入手机号，那么客户将无法收到华为云发出的任何提醒短信，需要客户自己登录到华为云平台补充手机号。 >-   调用“创建客户”接口时，华为云会同步创建华为云客户账号，将客户ID及账号名返回给伙伴平台，然后华为云异步完成客户与伙伴的关联。伙伴与客户的关联结果可通过“[查询客户列表](查询客户列表.md)”查询。
func (c *BssClient) CreateSubCustomer(request *model.CreateSubCustomerRequest) (*model.CreateSubCustomerResponse, error) {
	requestDef := GenReqDefForCreateSubCustomer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubCustomerResponse), nil
	}
}

//企业主账号在自建平台创建企业子账号。  企业主账号创建企业子账号请参见[这里](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0104194162.html)。
func (c *BssClient) CreateSubEnterpriseAccount(request *model.CreateSubEnterpriseAccountRequest) (*model.CreateSubEnterpriseAccountResponse, error) {
	requestDef := GenReqDefForCreateSubEnterpriseAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubEnterpriseAccountResponse), nil
	}
}

//伙伴可以删除自己的邮寄地址信息。  伙伴登录伙伴中心修改邮寄地址请参见[向华为云索取发票](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435143.html)，进入索取发票页面，选择删除邮寄地址，即可删除邮件地址。
func (c *BssClient) DeletePostal(request *model.DeletePostalRequest) (*model.DeletePostalResponse, error) {
	requestDef := GenReqDefForDeletePostal()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePostalResponse), nil
	}
}

//伙伴在伙伴销售平台上查询城市信息。
func (c *BssClient) ListCities(request *model.ListCitiesRequest) (*model.ListCitiesResponse, error) {
	requestDef := GenReqDefForListCities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCitiesResponse), nil
	}
}

//伙伴在伙伴销售平台上查询度量单位的进制转换信息，用于不同度量单位之间的转换。
func (c *BssClient) ListConversions(request *model.ListConversionsRequest) (*model.ListConversionsResponse, error) {
	requestDef := GenReqDefForListConversions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConversionsResponse), nil
	}
}

//伙伴在伙伴销售平台上查询区县信息。
func (c *BssClient) ListCounties(request *model.ListCountiesRequest) (*model.ListCountiesResponse, error) {
	requestDef := GenReqDefForListCounties()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCountiesResponse), nil
	}
}

//华为云伙伴能力中心（一级经销商）可以查看给精英服务商（二级经销商）发放或回收代金券额度的操作记录。  一级经销商可以登录伙伴中心，进入“客户业务** **\\> 代金券管理”，选择“代金券额度”页签，单击“操作记录”查看代金券额度的发放和回收记录。  >![](public_sys-resources/icon-note.gif) **说明：**  >伙伴也可以单击代金券额度所在行的“操作记录”，查看该代金券额度对应的操作记录日志。
func (c *BssClient) ListCouponQuotasRecords(request *model.ListCouponQuotasRecordsRequest) (*model.ListCouponQuotasRecordsResponse, error) {
	requestDef := GenReqDefForListCouponQuotasRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCouponQuotasRecordsResponse), nil
	}
}

//客户在自建平台查询自己的消费流水账单。  客户登录费用中心查询自己的消费流水账单请参见[这里](https://support.huaweicloud.com/usermanual-billing/bills-topic_80000001.html#bills-topic_80000001__zh-cn_topic_0000001162496407_s716e04d5d0ba4e9d9a76a8bcbfbcfe73)的“**查看流水账单**”。
func (c *BssClient) ListCustomerBillsFeeRecords(request *model.ListCustomerBillsFeeRecordsRequest) (*model.ListCustomerBillsFeeRecordsResponse, error) {
	requestDef := GenReqDefForListCustomerBillsFeeRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerBillsFeeRecordsResponse), nil
	}
}

//客户可以查询指定月份的月度摊销成本。当前仅支持查询近18个月的摊销成本。摊销成本计算规则请参见[成本摊销规则](https://support.huaweicloud.com/usermanual-cost/costcenter_000002_01.html)。  客户可查询的数据范围同成本中心提供的[数据范围](https://support.huaweicloud.com/usermanual-cost/costcenter_0000004.html)一致。  客户登录成本中心导出成本明细请参见[导出成本明细数据](https://support.huaweicloud.com/usermanual-cost/costcenter_000002_03.html)。  >![](public_sys-resources/icon-note.gif) **说明：**  >该接口仅面向已开通成本中心的客户开放。如何开启成本中心请参见[这里](https://support.huaweicloud.com/usermanual-cost/costcenter_000004.html)。
func (c *BssClient) ListCustomerBillsMonthlyBreakDown(request *model.ListCustomerBillsMonthlyBreakDownRequest) (*model.ListCustomerBillsMonthlyBreakDownResponse, error) {
	requestDef := GenReqDefForListCustomerBillsMonthlyBreakDown()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerBillsMonthlyBreakDownResponse), nil
	}
}

//合作伙伴可以查询关联的代售类客户已开通的按需资源。
func (c *BssClient) ListCustomerOnDemandResources(request *model.ListCustomerOnDemandResourcesRequest) (*model.ListCustomerOnDemandResourcesResponse, error) {
	requestDef := GenReqDefForListCustomerOnDemandResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerOnDemandResourcesResponse), nil
	}
}

//客户购买包年/包月资源后，可以查看待审核、处理中、已取消、已完成和待支付等状态的订单。  伙伴登录伙伴中心查看客户订单请单击[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0076200001.html)。  >![](public_sys-resources/icon-note.gif) **说明：**  >如果想查询某条订单下的资源信息，在调用本接口获取订单ID后，请调用“[查询客户包年/包月资源列表](查询客户包年-包月资源列表.md)”接口在请求参数输入订单号进行查询。
func (c *BssClient) ListCustomerOrders(request *model.ListCustomerOrdersRequest) (*model.ListCustomerOrdersResponse, error) {
	requestDef := GenReqDefForListCustomerOrders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerOrdersResponse), nil
	}
}

//合作伙伴可以查询关联的代售类客户的账户余额。  伙伴登录伙伴中心查询客户余额请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435115.html)。  >![](public_sys-resources/icon-note.gif) **说明：**  >顾问销售类客户是客户在华为云充值，合作伙伴无法调用此接口查询其账户余额。代售类客户的账户由合作伙伴拨款，所以可以查询到。
func (c *BssClient) ListCustomersBalancesDetail(request *model.ListCustomersBalancesDetailRequest) (*model.ListCustomersBalancesDetailResponse, error) {
	requestDef := GenReqDefForListCustomersBalancesDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomersBalancesDetailResponse), nil
	}
}

//客户在自建平台查询自己的资源详单，用于反映各类资源的消耗情况。  客户登录费用中心查询资源详单请参见[这里](https://support.huaweicloud.com/usermanual-billing/bills_topic_100000063.html)。  >![](public_sys-resources/icon-note.gif) **说明：**  >由于资源消费呈现的是资源维度的8位小数原始消费金额，实际从账户扣费时按2位小数进行扣费（即扣到分），会存在精度差异，所以，不推荐消费汇总和资源消费直接对账。
func (c *BssClient) ListCustomerselfResourceRecordDetails(request *model.ListCustomerselfResourceRecordDetailsRequest) (*model.ListCustomerselfResourceRecordDetailsResponse, error) {
	requestDef := GenReqDefForListCustomerselfResourceRecordDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerselfResourceRecordDetailsResponse), nil
	}
}

//客户在自建平台查询每个资源的消费明细数据。  客户登录费用中心查询资源消费记录请参见[这里](https://support.huaweicloud.com/usermanual-billing/bills_topic_100000061.html)。
func (c *BssClient) ListCustomerselfResourceRecords(request *model.ListCustomerselfResourceRecordsRequest) (*model.ListCustomerselfResourceRecordsResponse, error) {
	requestDef := GenReqDefForListCustomerselfResourceRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerselfResourceRecordsResponse), nil
	}
}

//企业主账号在自建平台查询企业子账号的可回收余额。
func (c *BssClient) ListEnterpriseMultiAccount(request *model.ListEnterpriseMultiAccountRequest) (*model.ListEnterpriseMultiAccountResponse, error) {
	requestDef := GenReqDefForListEnterpriseMultiAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnterpriseMultiAccountResponse), nil
	}
}

//企业主账号在自建平台查询企业组织结构。
func (c *BssClient) ListEnterpriseOrganizations(request *model.ListEnterpriseOrganizationsRequest) (*model.ListEnterpriseOrganizationsResponse, error) {
	requestDef := GenReqDefForListEnterpriseOrganizations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnterpriseOrganizationsResponse), nil
	}
}

//企业主账号在自建平台查询企业子账号信息列表。
func (c *BssClient) ListEnterpriseSubCustomers(request *model.ListEnterpriseSubCustomersRequest) (*model.ListEnterpriseSubCustomersResponse, error) {
	requestDef := GenReqDefForListEnterpriseSubCustomers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnterpriseSubCustomersResponse), nil
	}
}

//客户在伙伴销售平台查询客户的资源包列表。
func (c *BssClient) ListFreeResourceInfos(request *model.ListFreeResourceInfosRequest) (*model.ListFreeResourceInfosResponse, error) {
	requestDef := GenReqDefForListFreeResourceInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFreeResourceInfosResponse), nil
	}
}

//客户在伙伴销售平台查询客户的资源内使用量。
func (c *BssClient) ListFreeResourceUsages(request *model.ListFreeResourceUsagesRequest) (*model.ListFreeResourceUsagesResponse, error) {
	requestDef := GenReqDefForListFreeResourceUsages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFreeResourceUsagesResponse), nil
	}
}

//伙伴在伙伴销售平台上查询产品的折扣和激励策略。  伙伴登录合作伙伴中心查看产品的折扣和激励策略请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/dp_120400.html)。
func (c *BssClient) ListIncentiveDiscountPolicies(request *model.ListIncentiveDiscountPoliciesRequest) (*model.ListIncentiveDiscountPoliciesResponse, error) {
	requestDef := GenReqDefForListIncentiveDiscountPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIncentiveDiscountPoliciesResponse), nil
	}
}

//华为云伙伴能力中心（一级经销商）可以查询精英服务商（二级经销商）列表。  一级经销商在伙伴中心查询二级经销商列表的方式请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/dp_120210.html)。
func (c *BssClient) ListIndirectPartners(request *model.ListIndirectPartnersRequest) (*model.ListIndirectPartnersResponse, error) {
	requestDef := GenReqDefForListIndirectPartners()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIndirectPartnersResponse), nil
	}
}

//华为云伙伴能力中心（一级经销商）可以查看发放给精英服务商（二级经销商）的代金券额度列表。  一级经销商登录伙伴中心，进入“客户业务** **\\> 代金券管理”，选择“已发放代金券额度”可查看代金券额度列表。  >![](public_sys-resources/icon-note.gif) **说明：**  >调用该接口之前，需通过客户经理联系华为运营人员，为合作伙伴设置代金券发放额度。
func (c *BssClient) ListIssuedCouponQuotas(request *model.ListIssuedCouponQuotasRequest) (*model.ListIssuedCouponQuotasResponse, error) {
	requestDef := GenReqDefForListIssuedCouponQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIssuedCouponQuotasResponse), nil
	}
}

//合作伙伴可以查询已发放的优惠券列表。  伙伴登录伙伴中心，进入“客户业务** **\\> 代金券管理”，选择“已发放代金券”页签，即可查询已发放的代金券。
func (c *BssClient) ListIssuedPartnerCoupons(request *model.ListIssuedPartnerCouponsRequest) (*model.ListIssuedPartnerCouponsResponse, error) {
	requestDef := GenReqDefForListIssuedPartnerCoupons()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIssuedPartnerCouponsResponse), nil
	}
}

//伙伴在伙伴销售平台上查询资源使用量，包年包月资源的时长及金额的度量单位及名称，度量单位类型等。
func (c *BssClient) ListMeasureUnits(request *model.ListMeasureUnitsRequest) (*model.ListMeasureUnitsResponse, error) {
	requestDef := GenReqDefForListMeasureUnits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMeasureUnitsResponse), nil
	}
}

//伙伴在销售平台按照条件查询按需产品的价格。  如果购买该产品的租户享受折扣，可以在查询结果中返回折扣金额以及扣除折扣后的最后成交价。  如果该租户享受多种折扣，系统会优先返回客户享受的商务折扣的折扣金额和最终成交价。  >![](public_sys-resources/icon-caution.gif) **注意：**  >华为云根据云服务类型、资源类型、云服务区和资源规格四个条件来查询产品，查询时请确认这4个查询条件均输入正确，否则该接口会返回无法找到产品的错误。
func (c *BssClient) ListOnDemandResourceRatings(request *model.ListOnDemandResourceRatingsRequest) (*model.ListOnDemandResourceRatingsResponse, error) {
	requestDef := GenReqDefForListOnDemandResourceRatings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOnDemandResourceRatingsResponse), nil
	}
}

//客户在伙伴销售平台支付待支付订单时，查询可使用的优惠券列表。
func (c *BssClient) ListOrderCouponsByOrderId(request *model.ListOrderCouponsByOrderIdRequest) (*model.ListOrderCouponsByOrderIdResponse, error) {
	requestDef := GenReqDefForListOrderCouponsByOrderId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrderCouponsByOrderIdResponse), nil
	}
}

//客户在伙伴销售平台支付待支付订单时，查询可使用的折扣信息。
func (c *BssClient) ListOrderDiscounts(request *model.ListOrderDiscountsRequest) (*model.ListOrderDiscountsResponse, error) {
	requestDef := GenReqDefForListOrderDiscounts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrderDiscountsResponse), nil
	}
}

//伙伴在伙伴销售平台上查询自身的收支明细情况。
func (c *BssClient) ListPartnerAccountChangeRecords(request *model.ListPartnerAccountChangeRecordsRequest) (*model.ListPartnerAccountChangeRecordsResponse, error) {
	requestDef := GenReqDefForListPartnerAccountChangeRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPartnerAccountChangeRecordsResponse), nil
	}
}

//伙伴在伙伴销售平台查询向客户及关联的精英服务商（二级经销商）拨款或回收的调账记录。  伙伴登录伙伴中心，在“拨款”或“回收”页面，单击“调账记录”，可以查看一级经销商为二级经销商调账的记录。
func (c *BssClient) ListPartnerAdjustRecords(request *model.ListPartnerAdjustRecordsRequest) (*model.ListPartnerAdjustRecordsResponse, error) {
	requestDef := GenReqDefForListPartnerAdjustRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPartnerAdjustRecordsResponse), nil
	}
}

//合作伙伴可以查询伙伴的账户余额。
func (c *BssClient) ListPartnerBalances(request *model.ListPartnerBalancesRequest) (*model.ListPartnerBalancesResponse, error) {
	requestDef := GenReqDefForListPartnerBalances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPartnerBalancesResponse), nil
	}
}

//合作伙伴可查看给客户发放和回收优惠券的操作记录。  合作伙伴登录伙伴中心查看、导出代金券操作日志请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435103.html)。
func (c *BssClient) ListPartnerCouponsRecord(request *model.ListPartnerCouponsRecordRequest) (*model.ListPartnerCouponsRecordResponse, error) {
	requestDef := GenReqDefForListPartnerCouponsRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPartnerCouponsRecordResponse), nil
	}
}

//客户在伙伴销售平台查询某个或所有的包年/包月资源。  >![](public_sys-resources/icon-note.gif) **说明：**  >成功调用本接口后，如果您需要对已生效状态的资源进行续订，您可以调用“[查询包年/包月产品价格](查询包年-包月产品价格.md)”接口对查询到的包年/包月资源进行续订询价，然后再调用“[续订包年/包月资源](续订包年-包月资源.md)”接口进行续订。
func (c *BssClient) ListPayPerUseCustomerResources(request *model.ListPayPerUseCustomerResourcesRequest) (*model.ListPayPerUseCustomerResourcesResponse, error) {
	requestDef := GenReqDefForListPayPerUseCustomerResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPayPerUseCustomerResourcesResponse), nil
	}
}

//伙伴可以查询自己的邮寄地址信息。  伙伴登录伙伴中心查询邮寄地址请参见[向华为云索取发票](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435143.html)，进入索取发票页面，即可查看邮寄地址。
func (c *BssClient) ListPostalAddress(request *model.ListPostalAddressRequest) (*model.ListPostalAddressResponse, error) {
	requestDef := GenReqDefForListPostalAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPostalAddressResponse), nil
	}
}

//伙伴在伙伴销售平台上查询省份信息。
func (c *BssClient) ListProvinces(request *model.ListProvincesRequest) (*model.ListProvincesResponse, error) {
	requestDef := GenReqDefForListProvinces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProvincesResponse), nil
	}
}

//合作伙伴可以查看所拥有的优惠劵额度。  伙伴登录合作伙伴中心查看所拥有的代金券额度请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435100.html)。
func (c *BssClient) ListQuotaCoupons(request *model.ListQuotaCouponsRequest) (*model.ListQuotaCouponsResponse, error) {
	requestDef := GenReqDefForListQuotaCoupons()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotaCouponsResponse), nil
	}
}

//伙伴在销售平台按照条件查询包年/包月产品开通时候的价格。  如果购买该产品的客户享受折扣，可以在查询结果中返回折扣金额以及扣除折扣后的最后成交价。  如果该客户享受多种折扣，系统会返回每种折扣的批价结果。如果客户在下单的时候选择自动支付，则系统会优先应用商务折扣的批价结果。  >![](public_sys-resources/icon-caution.gif) **注意：**  >华为云根据云服务类型、资源类型、云服务区和资源规格四个条件来查询产品，查询时请确认这4个查询条件均输入正确，否则该接口会返回无法找到产品的错误。
func (c *BssClient) ListRateOnPeriodDetail(request *model.ListRateOnPeriodDetailRequest) (*model.ListRateOnPeriodDetailResponse, error) {
	requestDef := GenReqDefForListRateOnPeriodDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRateOnPeriodDetailResponse), nil
	}
}

//伙伴在伙伴销售平台查询资源类型的列表。
func (c *BssClient) ListResourceTypes(request *model.ListResourceTypesRequest) (*model.ListResourceTypesResponse, error) {
	requestDef := GenReqDefForListResourceTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceTypesResponse), nil
	}
}

//客户在自建平台查询自己的资源使用量明细。  >![](public_sys-resources/icon-note.gif) **说明：**  >当前仅支持查询CDN和OBS两种云服务类型的资源用量明细，仅针对95计费场景。
func (c *BssClient) ListResourceUsage(request *model.ListResourceUsageRequest) (*model.ListResourceUsageResponse, error) {
	requestDef := GenReqDefForListResourceUsage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceUsageResponse), nil
	}
}

//客户在自建平台查询自己的资源使用量汇总。  >![](public_sys-resources/icon-note.gif) **说明：**  >-   当前仅支持查询CDN和OBS两种云服务类型的资源用量汇总，仅针对95计费场景。 >-   使用量汇总列表只包含月汇总金额和资源ID，若要查询具体某个资源的用量明细，请调用[查询资源用量明细](查询95计费资源用量明细.md)接口获取。
func (c *BssClient) ListResourceUsageSummary(request *model.ListResourceUsageSummaryRequest) (*model.ListResourceUsageSummaryResponse, error) {
	requestDef := GenReqDefForListResourceUsageSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceUsageSummaryResponse), nil
	}
}

//伙伴在伙伴销售平台根据云服务类型查询关联的资源类型编码和名称，用于查询按需产品的价格或包年/包月产品的价格。
func (c *BssClient) ListServiceResources(request *model.ListServiceResourcesRequest) (*model.ListServiceResourcesResponse, error) {
	requestDef := GenReqDefForListServiceResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceResourcesResponse), nil
	}
}

//伙伴在伙伴销售平台查询云服务类型的列表。
func (c *BssClient) ListServiceTypes(request *model.ListServiceTypesRequest) (*model.ListServiceTypesResponse, error) {
	requestDef := GenReqDefForListServiceTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceTypesResponse), nil
	}
}

//客户可以查看自己已经购买的储值卡列表。
func (c *BssClient) ListStoredValueCards(request *model.ListStoredValueCardsRequest) (*model.ListStoredValueCardsResponse, error) {
	requestDef := GenReqDefForListStoredValueCards()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStoredValueCardsResponse), nil
	}
}

//伙伴在伙伴销售平台可实时查询子客户的消费记录，了解客户的资源消耗情况。  伙伴在伙伴中心查询客户消费明细请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435155.html)。  >![](public_sys-resources/icon-note.gif) **说明：**  >-   消费记录支持查询18个月内的记录。 >-   如果是客户经理主管来查询，只支持按照单个客户经理查询，必须输入客户经理ID。 >-   目前支持伙伴查询所有子客户（包含代售类和顾问销售类）的消费记录。
func (c *BssClient) ListSubCustomerBillDetail(request *model.ListSubCustomerBillDetailRequest) (*model.ListSubCustomerBillDetailResponse, error) {
	requestDef := GenReqDefForListSubCustomerBillDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubCustomerBillDetailResponse), nil
	}
}

//伙伴可以查询自身的优惠券信息。  伙伴登录伙伴中心查询已发放代金券列表请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435101.html)，查看已下发代金券的内容。
func (c *BssClient) ListSubCustomerCoupons(request *model.ListSubCustomerCouponsRequest) (*model.ListSubCustomerCouponsResponse, error) {
	requestDef := GenReqDefForListSubCustomerCoupons()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubCustomerCouponsResponse), nil
	}
}

//伙伴可以查询合作伙伴的客户信息列表。  伙伴登录合作伙伴中心查询客户信息列表请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435115.html)。
func (c *BssClient) ListSubCustomers(request *model.ListSubCustomersRequest) (*model.ListSubCustomersResponse, error) {
	requestDef := GenReqDefForListSubCustomers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubCustomersResponse), nil
	}
}

//合作伙伴可查询客户的消费汇总账单，消费按月汇总。  伙伴在伙伴中心查询客户月度消费账单请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435154.html)。  >![](public_sys-resources/icon-note.gif) **说明：**  >消费汇总数据仅包含前一天24点前的数据。每天刷新一次，更新前一天的数据。 >该接口用于合作伙伴查询其代售类客户在华为的消费情况，如果输入某个客户ID，则是查询单个客户的，否则是查询该伙伴下所有使用伙伴拨款消费的客户的消费记录（包括退订记录）。
func (c *BssClient) ListSubcustomerMonthlyBills(request *model.ListSubcustomerMonthlyBillsRequest) (*model.ListSubcustomerMonthlyBillsResponse, error) {
	requestDef := GenReqDefForListSubcustomerMonthlyBills()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubcustomerMonthlyBillsResponse), nil
	}
}

//伙伴在伙伴销售平台查询资源的使用量类型列表。
func (c *BssClient) ListUsageTypes(request *model.ListUsageTypesRequest) (*model.ListUsageTypesResponse, error) {
	requestDef := GenReqDefForListUsageTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsageTypesResponse), nil
	}
}

//客户可以对待支付状态的包年/包月产品订单进行支付。  客户登录费用中心支付包年包月产品的待支付订单请单击[这里](https://support.huaweicloud.com/usermanual-billing/zh-cn_topic_0031512547.html)。  >![](public_sys-resources/icon-note.gif) **说明：**  >-   API支持月度结算和余额支付两种支付方式，月度结算优先。 >-   余额支付包括现金账户和信用账户两种支付方式，如果两个账户都有余额，则优先现金账户支付。 >-   同时使用订单折扣和优惠券的互斥规则如下： >    -   如果优惠券的限制属性上存在simultaneousUseWithEmpowerDiscount字段，并且值为0，则折扣和优惠券不能同时使用。 >    -   如果优惠券的限制属性上存在minConsumeDiscount字段，当折扣ID包含的所有订单项上的折扣率discount\\_ratio都小于minConsumeDiscount字段时，则折扣ID和优惠券不能同时使用。
func (c *BssClient) PayOrders(request *model.PayOrdersRequest) (*model.PayOrdersResponse, error) {
	requestDef := GenReqDefForPayOrders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PayOrdersResponse), nil
	}
}

//华为云伙伴能力中心（一级经销商）可以回收已发放给精英服务商（二级经销商）的代金券额度。  一级经销商在伙伴中心回收已发放给二级经销商的代金券额度请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/dp_120206.html)。
func (c *BssClient) ReclaimCouponQuotas(request *model.ReclaimCouponQuotasRequest) (*model.ReclaimCouponQuotasResponse, error) {
	requestDef := GenReqDefForReclaimCouponQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReclaimCouponQuotasResponse), nil
	}
}

//华为云伙伴能力中心（一级经销商）可以回收精英服务商（二级经销商）的账户余额。  一级经销商在伙伴中心回收二级经销商账户拨款请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/dp_120205.html)。
func (c *BssClient) ReclaimIndirectPartnerAccount(request *model.ReclaimIndirectPartnerAccountRequest) (*model.ReclaimIndirectPartnerAccountResponse, error) {
	requestDef := GenReqDefForReclaimIndirectPartnerAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReclaimIndirectPartnerAccountResponse), nil
	}
}

//对于合作伙伴已经下发给客户的优惠券，如遇发错或其他特殊情况，合作伙伴有回收的权利。优惠券回收后，客户将不再拥有该优惠券。  伙伴登录合作伙伴中心回收为客户发放的代金券请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/espp_050503.html)。  >![](public_sys-resources/icon-note.gif) **说明：**  >只能回收代售类子客户的优惠券。
func (c *BssClient) ReclaimPartnerCoupons(request *model.ReclaimPartnerCouponsRequest) (*model.ReclaimPartnerCouponsResponse, error) {
	requestDef := GenReqDefForReclaimPartnerCoupons()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReclaimPartnerCouponsResponse), nil
	}
}

//企业主账号在自建平台回收给企业子账号的拨款。  如果回收的是企业子账户的信用账户，可以回收所有额度；如果回收金额大于子账户信用余额的时候，可能会导致子账户欠费，请慎重选择。
func (c *BssClient) ReclaimSubEnterpriseAmount(request *model.ReclaimSubEnterpriseAmountRequest) (*model.ReclaimSubEnterpriseAmountResponse, error) {
	requestDef := GenReqDefForReclaimSubEnterpriseAmount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReclaimSubEnterpriseAmountResponse), nil
	}
}

//当客户不再使用华为云产品时，合作伙伴可以回收代售类客户账户余额。  伙伴登录伙伴中心回收代售类客户账户余额请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435147.html)。
func (c *BssClient) ReclaimToPartnerAccount(request *model.ReclaimToPartnerAccountRequest) (*model.ReclaimToPartnerAccountResponse, error) {
	requestDef := GenReqDefForReclaimToPartnerAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReclaimToPartnerAccountResponse), nil
	}
}

//客户的包年/包月资源即将到期时，可进行包年/包月资源的续订。  客户在费用中心执行续订操作请参见[这里](https://support.huaweicloud.com/usermanual-billing/renewals_topic_10000003.html)。  >![](public_sys-resources/icon-note.gif) **说明：**  >-   调用接口后，如果某个主资源有对应的从资源，系统会将主资源和从资源一起续订，主资源的从资源信息可以通过调用[查询客户包年/包月资源列表](查询客户包年-包月资源列表.md)接口获取。 >-   注意：如ECS主机挂载新购的云硬盘，但此硬盘不是该ECS主资源的从资源，主从资源信息必须以调用[查询客户包年/包月资源列表](查询客户包年-包月资源列表.md)接口获取的信息为准。 >-   本接口支持自动支付，支付时使用折扣或优惠券的说明，请参见[支付使用折扣或优惠券说明](支付使用折扣或优惠券说明.md)。
func (c *BssClient) RenewalResources(request *model.RenewalResourcesRequest) (*model.RenewalResourcesResponse, error) {
	requestDef := GenReqDefForRenewalResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RenewalResourcesResponse), nil
	}
}

//企业主账号在自建平台发送短信验证码。
func (c *BssClient) SendSmsVerificationCode(request *model.SendSmsVerificationCodeRequest) (*model.SendSmsVerificationCodeResponse, error) {
	requestDef := GenReqDefForSendSmsVerificationCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendSmsVerificationCodeResponse), nil
	}
}

//客户注册时，如果填写了手机号，可以向对应的手机发送注册验证码，校验信息的正确性。使用个人银行卡方式进行实名认证时，通过该接口向指定的手机发送验证码。
func (c *BssClient) SendVerificationMessageCode(request *model.SendVerificationMessageCodeRequest) (*model.SendVerificationMessageCodeResponse, error) {
	requestDef := GenReqDefForSendVerificationMessageCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendVerificationMessageCodeResponse), nil
	}
}

//客户可以查询自身的账户余额。  客户可以登录费用中心进入“[总览](https://account.huaweicloud.com/usercenter/#/userindex/allview)”页面，在“可用额度”区域可以查询自身的账户余额。
func (c *BssClient) ShowCustomerAccountBalances(request *model.ShowCustomerAccountBalancesRequest) (*model.ShowCustomerAccountBalancesResponse, error) {
	requestDef := GenReqDefForShowCustomerAccountBalances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCustomerAccountBalancesResponse), nil
	}
}

//客户在自建平台查询自身的消费汇总账单，此账单按月汇总消费数据。  客户登录费用中心查询自身的消费汇总账单请参见[这里](https://support.huaweicloud.com/usermanual-billing/bills-topic_80000001.html#bills-topic_80000001__zh-cn_topic_0000001162496407_s620ce713baf04899a416d781d1817931)的“**查看汇总**”。  >![](public_sys-resources/icon-note.gif) **说明：**  >当前支持查看2019/01月份至今的费用账单。企业主账号展示的费用账单，包含关联的统一还款企业子账号的消费数据。
func (c *BssClient) ShowCustomerMonthlySum(request *model.ShowCustomerMonthlySumRequest) (*model.ShowCustomerMonthlySumResponse, error) {
	requestDef := GenReqDefForShowCustomerMonthlySum()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCustomerMonthlySumResponse), nil
	}
}

//客户可以在伙伴销售平台查看订单详情。  客户登录费用中心查看订单详情请单击[这里](https://support.huaweicloud.com/usermanual-billing/order_topic_9000001.html)。  >![](public_sys-resources/icon-note.gif) **说明：**  >如果想查询某条订单下的资源信息，请调用“[查询客户包年/包月资源列表](查询客户包年-包月资源列表.md)”接口在请求参数输入订单号进行查询。
func (c *BssClient) ShowCustomerOrderDetails(request *model.ShowCustomerOrderDetailsRequest) (*model.ShowCustomerOrderDetailsResponse, error) {
	requestDef := GenReqDefForShowCustomerOrderDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCustomerOrderDetailsResponse), nil
	}
}

//企业主账号在自建平台查询自己的可拨款余额。
func (c *BssClient) ShowMultiAccountTransferAmount(request *model.ShowMultiAccountTransferAmountRequest) (*model.ShowMultiAccountTransferAmountResponse, error) {
	requestDef := GenReqDefForShowMultiAccountTransferAmount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMultiAccountTransferAmountResponse), nil
	}
}

//如果实名认证申请或实名认证变更申请的响应中，显示需要人工审核，使用该接口查询审核结果。
func (c *BssClient) ShowRealnameAuthenticationReviewResult(request *model.ShowRealnameAuthenticationReviewResultRequest) (*model.ShowRealnameAuthenticationReviewResultResponse, error) {
	requestDef := GenReqDefForShowRealnameAuthenticationReviewResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRealnameAuthenticationReviewResultResponse), nil
	}
}

//客户在伙伴销售平台查询某次退订订单或者降配订单的退款金额来自哪些资源和对应订单。  >![](public_sys-resources/icon-note.gif) **说明：**  >-   可以在调用完“[退订包年/包月资源](退订包年-包月资源.md)”接口生成退订订单ID后，调用该接口查询退订订单对应的金额所属资源和订单。例如，调用“[退订包年/包月资源](退订包年-包月资源.md)”接口退订资源及其已续费周期后，您可以调用本小节的接口查询到退订金额归属的原开通订单ID和原续费订单ID。 >-   2018年5月份之后退订的订单才能查询到归属的原订单ID。
func (c *BssClient) ShowRefundOrderDetails(request *model.ShowRefundOrderDetailsRequest) (*model.ShowRefundOrderDetailsResponse, error) {
	requestDef := GenReqDefForShowRefundOrderDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRefundOrderDetailsResponse), nil
	}
}

//华为云伙伴能力中心（一级经销商）可以向精英服务商（二级经销商）发放代金券额度。  一级经销商在伙伴中心向二级经销商发放代金券额度请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/dp_120206.html)。
func (c *BssClient) UpdateCouponQuotas(request *model.UpdateCouponQuotasRequest) (*model.UpdateCouponQuotasResponse, error) {
	requestDef := GenReqDefForUpdateCouponQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCouponQuotasResponse), nil
	}
}

//合作伙伴可以为代售类客户账户拨款。  伙伴登录伙伴中心为代售类客户账户拨款请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435147.html)。
func (c *BssClient) UpdateCustomerAccountAmount(request *model.UpdateCustomerAccountAmountRequest) (*model.UpdateCustomerAccountAmountResponse, error) {
	requestDef := GenReqDefForUpdateCustomerAccountAmount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCustomerAccountAmountResponse), nil
	}
}

//华为云伙伴能力中心（一级经销商）可以向精英服务商（二级经销商）账户拨款。  一级经销商在伙伴中心向二级经销商拨款请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/dp_120205.html)。
func (c *BssClient) UpdateIndirectPartnerAccount(request *model.UpdateIndirectPartnerAccountRequest) (*model.UpdateIndirectPartnerAccountResponse, error) {
	requestDef := GenReqDefForUpdateIndirectPartnerAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIndirectPartnerAccountResponse), nil
	}
}

//客户可以设置包年/包月资源到期后转为按需资源计费。包年/包月计费模式到期后，按需的计费模式即生效。  客户在费用中心设置包年包月资源到期转按需请参见[这里](https://support.huaweicloud.com/usermanual-billing/renewals_topic_50000003.html)。  >![](public_sys-resources/icon-note.gif) **说明：**  >-   客户需要成功支付包年/包月资源订单后，才能设置资源的到期转按需。 >-   目前解决方案组合产品、按需套餐包不支持到期转按需。 >-   在调用本接口前，您可以调用“[查询客户包年/包月资源列表](查询客户包年-包月资源列表.md)”接口获取资源ID、资源过期时间以及资源过期后的扣费策略等信息。 >-   设置包年/包月资源到期转按需后，包年/包月资源到期后将自动变成按需计费。 >-   取消包年/包月资源到期转按需的前提是已经调用“[设置或取消包年/包月资源到期转按需](设置或取消包年-包月资源到期转按需.md)”接口设置包年/包月资源的到期转按需或在调用“[续订包年/包月资源](续订包年-包月资源.md)”接口时设置到期策略为到期转按需。
func (c *BssClient) UpdatePeriodToOnDemand(request *model.UpdatePeriodToOnDemandRequest) (*model.UpdatePeriodToOnDemandResponse, error) {
	requestDef := GenReqDefForUpdatePeriodToOnDemand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePeriodToOnDemandResponse), nil
	}
}

//伙伴可以修改自己的邮寄地址信息。  伙伴登录伙伴中心修改邮寄地址请参见[向华为云索取发票](https://support.huaweicloud.com/usermanual-bpconsole/zh-cn_topic_0072435143.html)，进入索取发票页面，选择修改邮寄地址，即可修改邮件地址。
func (c *BssClient) UpdatePostal(request *model.UpdatePostalRequest) (*model.UpdatePostalResponse, error) {
	requestDef := GenReqDefForUpdatePostal()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePostalResponse), nil
	}
}

//企业主账号在自建平台向企业子账号拨款。
func (c *BssClient) UpdateSubEnterpriseAmount(request *model.UpdateSubEnterpriseAmountRequest) (*model.UpdateSubEnterpriseAmountResponse, error) {
	requestDef := GenReqDefForUpdateSubEnterpriseAmount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubEnterpriseAmountResponse), nil
	}
}
