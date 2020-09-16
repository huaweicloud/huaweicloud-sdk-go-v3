package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"
)

type CSBPartnerOpenAPIClient struct {
	hcClient *http_client.HcHttpClient
}

func NewCSBPartnerOpenAPIClient(hcClient *http_client.HcHttpClient) *CSBPartnerOpenAPIClient {
	return &CSBPartnerOpenAPIClient{hcClient: hcClient}
}

func CSBPartnerOpenAPIClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

//功能描述：设置包周期资源自动续费
func (c *CSBPartnerOpenAPIClient) AutoRenewalResources(request *model.AutoRenewalResourcesRequest) (*model.AutoRenewalResourcesResponse, error) {
	requestDef := GenReqDefForAutoRenewalResources(request)
	resp, responseDef := GenRespForAutoRenewalResources()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：设置伙伴折扣
func (c *CSBPartnerOpenAPIClient) BatchSetSubCustomerDiscount(request *model.BatchSetSubCustomerDiscountRequest) (*model.BatchSetSubCustomerDiscountResponse, error) {
	requestDef := GenReqDefForBatchSetSubCustomerDiscount(request)
	resp, responseDef := GenRespForBatchSetSubCustomerDiscount()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：取消包年/包月资源自动续费
func (c *CSBPartnerOpenAPIClient) CancelAutoRenewalResources(request *model.CancelAutoRenewalResourcesRequest) (*model.CancelAutoRenewalResourcesResponse, error) {
	requestDef := GenReqDefForCancelAutoRenewalResources(request)
	resp, responseDef := GenRespForCancelAutoRenewalResources()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：取消包周期订单
func (c *CSBPartnerOpenAPIClient) CancelCustomerOrder(request *model.CancelCustomerOrderRequest) (*model.CancelCustomerOrderResponse, error) {
	requestDef := GenReqDefForCancelCustomerOrder(request)
	resp, responseDef := GenRespForCancelCustomerOrder()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：退订包周期资源
func (c *CSBPartnerOpenAPIClient) CancelResourcesSubscription(request *model.CancelResourcesSubscriptionRequest) (*model.CancelResourcesSubscriptionResponse, error) {
	requestDef := GenReqDefForCancelResourcesSubscription(request)
	resp, responseDef := GenRespForCancelResourcesSubscription()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：实名认证变更申请
func (c *CSBPartnerOpenAPIClient) ChangeEnterpriseRealnameAuthentication(request *model.ChangeEnterpriseRealnameAuthenticationRequest) (*model.ChangeEnterpriseRealnameAuthenticationResponse, error) {
	requestDef := GenReqDefForChangeEnterpriseRealnameAuthentication(request)
	resp, responseDef := GenRespForChangeEnterpriseRealnameAuthentication()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：校验客户的注册信息
func (c *CSBPartnerOpenAPIClient) CheckUserIdentity(request *model.CheckUserIdentityRequest) (*model.CheckUserIdentityResponse, error) {
	requestDef := GenReqDefForCheckUserIdentity(request)
	resp, responseDef := GenRespForCheckUserIdentity()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：客户开通自身的企业项目功能
func (c *CSBPartnerOpenAPIClient) CreateEnterpriseProjectAuth(request *model.CreateEnterpriseProjectAuthRequest) (*model.CreateEnterpriseProjectAuthResponse, error) {
	requestDef := GenReqDefForCreateEnterpriseProjectAuth(request)
	resp, responseDef := GenRespForCreateEnterpriseProjectAuth()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：企业实名认证申请V2
func (c *CSBPartnerOpenAPIClient) CreateEnterpriseRealnameAuthentication(request *model.CreateEnterpriseRealnameAuthenticationRequest) (*model.CreateEnterpriseRealnameAuthenticationResponse, error) {
	requestDef := GenReqDefForCreateEnterpriseRealnameAuthentication(request)
	resp, responseDef := GenRespForCreateEnterpriseRealnameAuthentication()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：伙伴给子客户发券
func (c *CSBPartnerOpenAPIClient) CreatePartnerCoupons(request *model.CreatePartnerCouponsRequest) (*model.CreatePartnerCouponsResponse, error) {
	requestDef := GenReqDefForCreatePartnerCoupons(request)
	resp, responseDef := GenRespForCreatePartnerCoupons()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：个人实名认证申请
func (c *CSBPartnerOpenAPIClient) CreatePersonalRealnameAuth(request *model.CreatePersonalRealnameAuthRequest) (*model.CreatePersonalRealnameAuthResponse, error) {
	requestDef := GenReqDefForCreatePersonalRealnameAuth(request)
	resp, responseDef := GenRespForCreatePersonalRealnameAuth()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：新增邮寄地址
func (c *CSBPartnerOpenAPIClient) CreatePostal(request *model.CreatePostalRequest) (*model.CreatePostalResponse, error) {
	requestDef := GenReqDefForCreatePostal(request)
	resp, responseDef := GenRespForCreatePostal()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：在伙伴销售平台创建客户时同步创建华为云账号，并将客户在伙伴销售平台上的账号与华为云账号进行映射。同时，创建的华为云账号与伙伴账号关联绑定。
func (c *CSBPartnerOpenAPIClient) CreateSubCustomer(request *model.CreateSubCustomerRequest) (*model.CreateSubCustomerResponse, error) {
	requestDef := GenReqDefForCreateSubCustomer(request)
	resp, responseDef := GenRespForCreateSubCustomer()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：创建企业子账号
func (c *CSBPartnerOpenAPIClient) CreateSubEnterpriseAccount(request *model.CreateSubEnterpriseAccountRequest) (*model.CreateSubEnterpriseAccountResponse, error) {
	requestDef := GenReqDefForCreateSubEnterpriseAccount(request)
	resp, responseDef := GenRespForCreateSubEnterpriseAccount()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：删除邮寄地址
func (c *CSBPartnerOpenAPIClient) DeletePostal(request *model.DeletePostalRequest) (*model.DeletePostalResponse, error) {
	requestDef := GenReqDefForDeletePostal(request)
	resp, responseDef := GenRespForDeletePostal()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：根据省份查询城市列表
func (c *CSBPartnerOpenAPIClient) ListCities(request *model.ListCitiesRequest) (*model.ListCitiesResponse, error) {
	requestDef := GenReqDefForListCities(request)
	resp, responseDef := GenRespForListCities()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：根据城市查询区县列表
func (c *CSBPartnerOpenAPIClient) ListCounties(request *model.ListCountiesRequest) (*model.ListCountiesResponse, error) {
	requestDef := GenReqDefForListCounties(request)
	resp, responseDef := GenRespForListCounties()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询优惠券额度发放回收记录
func (c *CSBPartnerOpenAPIClient) ListCouponQuotasRecords(request *model.ListCouponQuotasRecordsRequest) (*model.ListCouponQuotasRecordsResponse, error) {
	requestDef := GenReqDefForListCouponQuotasRecords(request)
	resp, responseDef := GenRespForListCouponQuotasRecords()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询客户按需资源列表
func (c *CSBPartnerOpenAPIClient) ListCustomerOnDemandResources(request *model.ListCustomerOnDemandResourcesRequest) (*model.ListCustomerOnDemandResourcesResponse, error) {
	requestDef := GenReqDefForListCustomerOnDemandResources(request)
	resp, responseDef := GenRespForListCustomerOnDemandResources()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询订单列表
func (c *CSBPartnerOpenAPIClient) ListCustomerOrders(request *model.ListCustomerOrdersRequest) (*model.ListCustomerOrdersResponse, error) {
	requestDef := GenReqDefForListCustomerOrders(request)
	resp, responseDef := GenRespForListCustomerOrders()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：批量查询伙伴子客户账户余额
func (c *CSBPartnerOpenAPIClient) ListCustomersBalancesDetail(request *model.ListCustomersBalancesDetailRequest) (*model.ListCustomersBalancesDetailResponse, error) {
	requestDef := GenReqDefForListCustomersBalancesDetail(request)
	resp, responseDef := GenRespForListCustomersBalancesDetail()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：客户在客户自建平台查询自己的资源详单，用于反映各类资源的消耗情况。资源详单数据有延迟，最大延迟24小时。
func (c *CSBPartnerOpenAPIClient) ListCustomerselfResourceRecordDetails(request *model.ListCustomerselfResourceRecordDetailsRequest) (*model.ListCustomerselfResourceRecordDetailsResponse, error) {
	requestDef := GenReqDefForListCustomerselfResourceRecordDetails(request)
	resp, responseDef := GenRespForListCustomerselfResourceRecordDetails()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询资源消费记录（客户）
func (c *CSBPartnerOpenAPIClient) ListCustomerselfResourceRecords(request *model.ListCustomerselfResourceRecordsRequest) (*model.ListCustomerselfResourceRecordsResponse, error) {
	requestDef := GenReqDefForListCustomerselfResourceRecords(request)
	resp, responseDef := GenRespForListCustomerselfResourceRecords()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询企业子可回收余额
func (c *CSBPartnerOpenAPIClient) ListEnterpriseMultiAccount(request *model.ListEnterpriseMultiAccountRequest) (*model.ListEnterpriseMultiAccountResponse, error) {
	requestDef := GenReqDefForListEnterpriseMultiAccount(request)
	resp, responseDef := GenRespForListEnterpriseMultiAccount()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询精英服务商列表
func (c *CSBPartnerOpenAPIClient) ListIndirectPartners(request *model.ListIndirectPartnersRequest) (*model.ListIndirectPartnersResponse, error) {
	requestDef := GenReqDefForListIndirectPartners(request)
	resp, responseDef := GenRespForListIndirectPartners()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：一级经销商查询发给二级经销商的额度
func (c *CSBPartnerOpenAPIClient) ListIssuedCouponQuotas(request *model.ListIssuedCouponQuotasRequest) (*model.ListIssuedCouponQuotasResponse, error) {
	requestDef := GenReqDefForListIssuedCouponQuotas(request)
	resp, responseDef := GenRespForListIssuedCouponQuotas()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询已发放的优惠券列表
func (c *CSBPartnerOpenAPIClient) ListIssuedPartnerCoupons(request *model.ListIssuedPartnerCouponsRequest) (*model.ListIssuedPartnerCouponsResponse, error) {
	requestDef := GenReqDefForListIssuedPartnerCoupons(request)
	resp, responseDef := GenRespForListIssuedPartnerCoupons()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：按需资源询价
func (c *CSBPartnerOpenAPIClient) ListOnDemandResourceRatings(request *model.ListOnDemandResourceRatingsRequest) (*model.ListOnDemandResourceRatingsResponse, error) {
	requestDef := GenReqDefForListOnDemandResourceRatings(request)
	resp, responseDef := GenRespForListOnDemandResourceRatings()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询订单详情
func (c *CSBPartnerOpenAPIClient) ListOrderCouponsByOrderId(request *model.ListOrderCouponsByOrderIdRequest) (*model.ListOrderCouponsByOrderIdResponse, error) {
	requestDef := GenReqDefForListOrderCouponsByOrderId(request)
	resp, responseDef := GenRespForListOrderCouponsByOrderId()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询调账回收记录
func (c *CSBPartnerOpenAPIClient) ListPartnerAdjustRecords(request *model.ListPartnerAdjustRecordsRequest) (*model.ListPartnerAdjustRecordsResponse, error) {
	requestDef := GenReqDefForListPartnerAdjustRecords(request)
	resp, responseDef := GenRespForListPartnerAdjustRecords()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询伙伴账户余额
func (c *CSBPartnerOpenAPIClient) ListPartnerBalances(request *model.ListPartnerBalancesRequest) (*model.ListPartnerBalancesResponse, error) {
	requestDef := GenReqDefForListPartnerBalances(request)
	resp, responseDef := GenRespForListPartnerBalances()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询优惠券的发放回收记录
func (c *CSBPartnerOpenAPIClient) ListPartnerCouponsRecord(request *model.ListPartnerCouponsRecordRequest) (*model.ListPartnerCouponsRecordResponse, error) {
	requestDef := GenReqDefForListPartnerCouponsRecord(request)
	resp, responseDef := GenRespForListPartnerCouponsRecord()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询伙伴代付订单列表
func (c *CSBPartnerOpenAPIClient) ListPartnerPayOrders(request *model.ListPartnerPayOrdersRequest) (*model.ListPartnerPayOrdersResponse, error) {
	requestDef := GenReqDefForListPartnerPayOrders(request)
	resp, responseDef := GenRespForListPartnerPayOrders()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询客户包年/包月资源列表
func (c *CSBPartnerOpenAPIClient) ListPayPerUseCustomerResources(request *model.ListPayPerUseCustomerResourcesRequest) (*model.ListPayPerUseCustomerResourcesResponse, error) {
	requestDef := GenReqDefForListPayPerUseCustomerResources(request)
	resp, responseDef := GenRespForListPayPerUseCustomerResources()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询邮寄地址
func (c *CSBPartnerOpenAPIClient) ListPostalAddress(request *model.ListPostalAddressRequest) (*model.ListPostalAddressResponse, error) {
	requestDef := GenReqDefForListPostalAddress(request)
	resp, responseDef := GenRespForListPostalAddress()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：根据国家查询省份列表
func (c *CSBPartnerOpenAPIClient) ListProvinces(request *model.ListProvincesRequest) (*model.ListProvincesResponse, error) {
	requestDef := GenReqDefForListProvinces(request)
	resp, responseDef := GenRespForListProvinces()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询优惠券额度列表
func (c *CSBPartnerOpenAPIClient) ListQuotaCoupons(request *model.ListQuotaCouponsRequest) (*model.ListQuotaCouponsResponse, error) {
	requestDef := GenReqDefForListQuotaCoupons(request)
	resp, responseDef := GenRespForListQuotaCoupons()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：包周期资源订购询价
func (c *CSBPartnerOpenAPIClient) ListRateOnPeriodDetail(request *model.ListRateOnPeriodDetailRequest) (*model.ListRateOnPeriodDetailResponse, error) {
	requestDef := GenReqDefForListRateOnPeriodDetail(request)
	resp, responseDef := GenRespForListRateOnPeriodDetail()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询套餐内使用量
func (c *CSBPartnerOpenAPIClient) ListResourceUsages(request *model.ListResourceUsagesRequest) (*model.ListResourceUsagesResponse, error) {
	requestDef := GenReqDefForListResourceUsages(request)
	resp, responseDef := GenRespForListResourceUsages()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询硬件库存
func (c *CSBPartnerOpenAPIClient) ListSkuInventories(request *model.ListSkuInventoriesRequest) (*model.ListSkuInventoriesResponse, error) {
	requestDef := GenReqDefForListSkuInventories(request)
	resp, responseDef := GenRespForListSkuInventories()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询优惠券列表
func (c *CSBPartnerOpenAPIClient) ListSubCustomerCoupons(request *model.ListSubCustomerCouponsRequest) (*model.ListSubCustomerCouponsResponse, error) {
	requestDef := GenReqDefForListSubCustomerCoupons(request)
	resp, responseDef := GenRespForListSubCustomerCoupons()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询伙伴折扣
func (c *CSBPartnerOpenAPIClient) ListSubCustomerDiscounts(request *model.ListSubCustomerDiscountsRequest) (*model.ListSubCustomerDiscountsResponse, error) {
	requestDef := GenReqDefForListSubCustomerDiscounts(request)
	resp, responseDef := GenRespForListSubCustomerDiscounts()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询客户消费记录
func (c *CSBPartnerOpenAPIClient) ListSubCustomerResFeeRecords(request *model.ListSubCustomerResFeeRecordsRequest) (*model.ListSubCustomerResFeeRecordsResponse, error) {
	requestDef := GenReqDefForListSubCustomerResFeeRecords(request)
	resp, responseDef := GenRespForListSubCustomerResFeeRecords()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询客户列表
func (c *CSBPartnerOpenAPIClient) ListSubCustomers(request *model.ListSubCustomersRequest) (*model.ListSubCustomersResponse, error) {
	requestDef := GenReqDefForListSubCustomers(request)
	resp, responseDef := GenRespForListSubCustomers()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：合作伙伴可查询客户的消费汇总账单，消费按月汇总
func (c *CSBPartnerOpenAPIClient) ListSubcustomerMonthlyBills(request *model.ListSubcustomerMonthlyBillsRequest) (*model.ListSubcustomerMonthlyBillsResponse, error) {
	requestDef := GenReqDefForListSubcustomerMonthlyBills(request)
	resp, responseDef := GenRespForListSubcustomerMonthlyBills()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：支付包周期订单
func (c *CSBPartnerOpenAPIClient) PayOrders(request *model.PayOrdersRequest) (*model.PayOrdersResponse, error) {
	requestDef := GenReqDefForPayOrders(request)
	resp, responseDef := GenRespForPayOrders()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：一级经销商给二级经销商回收额度
func (c *CSBPartnerOpenAPIClient) ReclaimCouponQuotas(request *model.ReclaimCouponQuotasRequest) (*model.ReclaimCouponQuotasResponse, error) {
	requestDef := GenReqDefForReclaimCouponQuotas(request)
	resp, responseDef := GenRespForReclaimCouponQuotas()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：合作伙伴可以回收二级渠道账户余额
func (c *CSBPartnerOpenAPIClient) ReclaimIndirectPartnerAccount(request *model.ReclaimIndirectPartnerAccountRequest) (*model.ReclaimIndirectPartnerAccountResponse, error) {
	requestDef := GenReqDefForReclaimIndirectPartnerAccount(request)
	resp, responseDef := GenRespForReclaimIndirectPartnerAccount()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：伙伴回收子客户优惠券
func (c *CSBPartnerOpenAPIClient) ReclaimPartnerCoupons(request *model.ReclaimPartnerCouponsRequest) (*model.ReclaimPartnerCouponsResponse, error) {
	requestDef := GenReqDefForReclaimPartnerCoupons(request)
	resp, responseDef := GenRespForReclaimPartnerCoupons()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：企业主账号从企业子账号回收拨款
func (c *CSBPartnerOpenAPIClient) ReclaimSubEnterpriseAmount(request *model.ReclaimSubEnterpriseAmountRequest) (*model.ReclaimSubEnterpriseAmountResponse, error) {
	requestDef := GenReqDefForReclaimSubEnterpriseAmount(request)
	resp, responseDef := GenRespForReclaimSubEnterpriseAmount()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：回收子客户余额（支持一级回收二级的子客户余额）
func (c *CSBPartnerOpenAPIClient) ReclaimToPartnerAccount(request *model.ReclaimToPartnerAccountRequest) (*model.ReclaimToPartnerAccountResponse, error) {
	requestDef := GenReqDefForReclaimToPartnerAccount(request)
	resp, responseDef := GenRespForReclaimToPartnerAccount()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：续订包周期资源
func (c *CSBPartnerOpenAPIClient) RenewalResources(request *model.RenewalResourcesRequest) (*model.RenewalResourcesResponse, error) {
	requestDef := GenReqDefForRenewalResources(request)
	resp, responseDef := GenRespForRenewalResources()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：发送短信验证码
func (c *CSBPartnerOpenAPIClient) SendSmsVerificationCode(request *model.SendSmsVerificationCodeRequest) (*model.SendSmsVerificationCodeResponse, error) {
	requestDef := GenReqDefForSendSmsVerificationCode(request)
	resp, responseDef := GenRespForSendSmsVerificationCode()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：发送验证码
func (c *CSBPartnerOpenAPIClient) SendVerificationMessageCode(request *model.SendVerificationMessageCodeRequest) (*model.SendVerificationMessageCodeResponse, error) {
	requestDef := GenReqDefForSendVerificationMessageCode(request)
	resp, responseDef := GenRespForSendVerificationMessageCode()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询客户账户余额
func (c *CSBPartnerOpenAPIClient) ShowCusotmerAccountBalances(request *model.ShowCusotmerAccountBalancesRequest) (*model.ShowCusotmerAccountBalancesResponse, error) {
	requestDef := GenReqDefForShowCusotmerAccountBalances(request)
	resp, responseDef := GenRespForShowCusotmerAccountBalances()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：客户可以查询自身的消费汇总单的功能，消费按月汇总。每天刷新一次，更新前一天的数据。
func (c *CSBPartnerOpenAPIClient) ShowCustomerMonthlySum(request *model.ShowCustomerMonthlySumRequest) (*model.ShowCustomerMonthlySumResponse, error) {
	requestDef := GenReqDefForShowCustomerMonthlySum(request)
	resp, responseDef := GenRespForShowCustomerMonthlySum()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询订单详情
func (c *CSBPartnerOpenAPIClient) ShowCustomerOrderDetails(request *model.ShowCustomerOrderDetailsRequest) (*model.ShowCustomerOrderDetailsResponse, error) {
	requestDef := GenReqDefForShowCustomerOrderDetails(request)
	resp, responseDef := GenRespForShowCustomerOrderDetails()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询企业主的可拨款余额
func (c *CSBPartnerOpenAPIClient) ShowMultiAccountTransferAmount(request *model.ShowMultiAccountTransferAmountRequest) (*model.ShowMultiAccountTransferAmountResponse, error) {
	requestDef := GenReqDefForShowMultiAccountTransferAmount(request)
	resp, responseDef := GenRespForShowMultiAccountTransferAmount()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询实名认证审核结果
func (c *CSBPartnerOpenAPIClient) ShowRealnameAuthenticationReviewResult(request *model.ShowRealnameAuthenticationReviewResultRequest) (*model.ShowRealnameAuthenticationReviewResultResponse, error) {
	requestDef := GenReqDefForShowRealnameAuthenticationReviewResult(request)
	resp, responseDef := GenRespForShowRealnameAuthenticationReviewResult()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：查询退款订单的金额详情
func (c *CSBPartnerOpenAPIClient) ShowRefundOrderDetails(request *model.ShowRefundOrderDetailsRequest) (*model.ShowRefundOrderDetailsResponse, error) {
	requestDef := GenReqDefForShowRefundOrderDetails(request)
	resp, responseDef := GenRespForShowRefundOrderDetails()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：向精英服务商发放代金券额度
func (c *CSBPartnerOpenAPIClient) UpdateCouponQuotas(request *model.UpdateCouponQuotasRequest) (*model.UpdateCouponQuotasResponse, error) {
	requestDef := GenReqDefForUpdateCouponQuotas(request)
	resp, responseDef := GenRespForUpdateCouponQuotas()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：从伙伴账户调账给子客户
func (c *CSBPartnerOpenAPIClient) UpdateCustomerAccountAmount(request *model.UpdateCustomerAccountAmountRequest) (*model.UpdateCustomerAccountAmountResponse, error) {
	requestDef := GenReqDefForUpdateCustomerAccountAmount(request)
	resp, responseDef := GenRespForUpdateCustomerAccountAmount()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：从伙伴账户调账给二级渠道
func (c *CSBPartnerOpenAPIClient) UpdateIndirectPartnerAccount(request *model.UpdateIndirectPartnerAccountRequest) (*model.UpdateIndirectPartnerAccountResponse, error) {
	requestDef := GenReqDefForUpdateIndirectPartnerAccount(request)
	resp, responseDef := GenRespForUpdateIndirectPartnerAccount()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：修改邮寄地址
func (c *CSBPartnerOpenAPIClient) UpdatePostal(request *model.UpdatePostalRequest) (*model.UpdatePostalResponse, error) {
	requestDef := GenReqDefForUpdatePostal(request)
	resp, responseDef := GenRespForUpdatePostal()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//功能描述：企业主账号向企业子账号拨款
func (c *CSBPartnerOpenAPIClient) UpdateSubEnterpriseAmount(request *model.UpdateSubEnterpriseAmountRequest) (*model.UpdateSubEnterpriseAmountResponse, error) {
	requestDef := GenReqDefForUpdateSubEnterpriseAmount(request)
	resp, responseDef := GenRespForUpdateSubEnterpriseAmount()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
