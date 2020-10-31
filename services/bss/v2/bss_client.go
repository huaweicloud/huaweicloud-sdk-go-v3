package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"
)

type BssClient struct {
	hcClient *http_client.HcHttpClient
}

func NewBssClient(hcClient *http_client.HcHttpClient) *BssClient {
	return &BssClient{hcClient: hcClient}
}

func BssClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

//功能描述：设置包周期资源自动续费
func (c *BssClient) AutoRenewalResources(request *model.AutoRenewalResourcesRequest) (*model.AutoRenewalResourcesResponse, error) {
	requestDef := GenReqDefForAutoRenewalResources()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AutoRenewalResourcesResponse), nil
	}
}

//功能描述：设置伙伴折扣
func (c *BssClient) BatchSetSubCustomerDiscount(request *model.BatchSetSubCustomerDiscountRequest) (*model.BatchSetSubCustomerDiscountResponse, error) {
	requestDef := GenReqDefForBatchSetSubCustomerDiscount()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetSubCustomerDiscountResponse), nil
	}
}

//功能描述：取消包年/包月资源自动续费
func (c *BssClient) CancelAutoRenewalResources(request *model.CancelAutoRenewalResourcesRequest) (*model.CancelAutoRenewalResourcesResponse, error) {
	requestDef := GenReqDefForCancelAutoRenewalResources()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelAutoRenewalResourcesResponse), nil
	}
}

//功能描述：取消包周期订单
func (c *BssClient) CancelCustomerOrder(request *model.CancelCustomerOrderRequest) (*model.CancelCustomerOrderResponse, error) {
	requestDef := GenReqDefForCancelCustomerOrder()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelCustomerOrderResponse), nil
	}
}

//功能描述：退订包周期资源
func (c *BssClient) CancelResourcesSubscription(request *model.CancelResourcesSubscriptionRequest) (*model.CancelResourcesSubscriptionResponse, error) {
	requestDef := GenReqDefForCancelResourcesSubscription()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelResourcesSubscriptionResponse), nil
	}
}

//功能描述：实名认证变更申请
func (c *BssClient) ChangeEnterpriseRealnameAuthentication(request *model.ChangeEnterpriseRealnameAuthenticationRequest) (*model.ChangeEnterpriseRealnameAuthenticationResponse, error) {
	requestDef := GenReqDefForChangeEnterpriseRealnameAuthentication()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeEnterpriseRealnameAuthenticationResponse), nil
	}
}

//功能描述：校验客户的注册信息
func (c *BssClient) CheckUserIdentity(request *model.CheckUserIdentityRequest) (*model.CheckUserIdentityResponse, error) {
	requestDef := GenReqDefForCheckUserIdentity()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckUserIdentityResponse), nil
	}
}

//功能描述：客户开通自身的企业项目功能
func (c *BssClient) CreateEnterpriseProjectAuth(request *model.CreateEnterpriseProjectAuthRequest) (*model.CreateEnterpriseProjectAuthResponse, error) {
	requestDef := GenReqDefForCreateEnterpriseProjectAuth()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnterpriseProjectAuthResponse), nil
	}
}

//功能描述：企业实名认证申请V2
func (c *BssClient) CreateEnterpriseRealnameAuthentication(request *model.CreateEnterpriseRealnameAuthenticationRequest) (*model.CreateEnterpriseRealnameAuthenticationResponse, error) {
	requestDef := GenReqDefForCreateEnterpriseRealnameAuthentication()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnterpriseRealnameAuthenticationResponse), nil
	}
}

//功能描述：伙伴给子客户发券
func (c *BssClient) CreatePartnerCoupons(request *model.CreatePartnerCouponsRequest) (*model.CreatePartnerCouponsResponse, error) {
	requestDef := GenReqDefForCreatePartnerCoupons()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePartnerCouponsResponse), nil
	}
}

//功能描述：个人实名认证申请
func (c *BssClient) CreatePersonalRealnameAuth(request *model.CreatePersonalRealnameAuthRequest) (*model.CreatePersonalRealnameAuthResponse, error) {
	requestDef := GenReqDefForCreatePersonalRealnameAuth()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePersonalRealnameAuthResponse), nil
	}
}

//功能描述：新增邮寄地址
func (c *BssClient) CreatePostal(request *model.CreatePostalRequest) (*model.CreatePostalResponse, error) {
	requestDef := GenReqDefForCreatePostal()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePostalResponse), nil
	}
}

//功能描述：在伙伴销售平台创建客户时同步创建华为云账号，并将客户在伙伴销售平台上的账号与华为云账号进行映射。同时，创建的华为云账号与伙伴账号关联绑定。
func (c *BssClient) CreateSubCustomer(request *model.CreateSubCustomerRequest) (*model.CreateSubCustomerResponse, error) {
	requestDef := GenReqDefForCreateSubCustomer()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubCustomerResponse), nil
	}
}

//功能描述：创建企业子账号
func (c *BssClient) CreateSubEnterpriseAccount(request *model.CreateSubEnterpriseAccountRequest) (*model.CreateSubEnterpriseAccountResponse, error) {
	requestDef := GenReqDefForCreateSubEnterpriseAccount()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubEnterpriseAccountResponse), nil
	}
}

//功能描述：删除邮寄地址
func (c *BssClient) DeletePostal(request *model.DeletePostalRequest) (*model.DeletePostalResponse, error) {
	requestDef := GenReqDefForDeletePostal()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePostalResponse), nil
	}
}

//功能描述：根据省份查询城市列表
func (c *BssClient) ListCities(request *model.ListCitiesRequest) (*model.ListCitiesResponse, error) {
	requestDef := GenReqDefForListCities()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCitiesResponse), nil
	}
}

//功能描述：根据城市查询区县列表
func (c *BssClient) ListCounties(request *model.ListCountiesRequest) (*model.ListCountiesResponse, error) {
	requestDef := GenReqDefForListCounties()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCountiesResponse), nil
	}
}

//功能描述：查询优惠券额度发放回收记录
func (c *BssClient) ListCouponQuotasRecords(request *model.ListCouponQuotasRecordsRequest) (*model.ListCouponQuotasRecordsResponse, error) {
	requestDef := GenReqDefForListCouponQuotasRecords()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCouponQuotasRecordsResponse), nil
	}
}

//功能描述：客户在客户自建平台查询自己的流水账单
func (c *BssClient) ListCustomerBillsFeeRecords(request *model.ListCustomerBillsFeeRecordsRequest) (*model.ListCustomerBillsFeeRecordsResponse, error) {
	requestDef := GenReqDefForListCustomerBillsFeeRecords()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerBillsFeeRecordsResponse), nil
	}
}

//功能描述：查询客户按需资源列表
func (c *BssClient) ListCustomerOnDemandResources(request *model.ListCustomerOnDemandResourcesRequest) (*model.ListCustomerOnDemandResourcesResponse, error) {
	requestDef := GenReqDefForListCustomerOnDemandResources()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerOnDemandResourcesResponse), nil
	}
}

//功能描述：查询订单列表
func (c *BssClient) ListCustomerOrders(request *model.ListCustomerOrdersRequest) (*model.ListCustomerOrdersResponse, error) {
	requestDef := GenReqDefForListCustomerOrders()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerOrdersResponse), nil
	}
}

//功能描述：批量查询伙伴子客户账户余额
func (c *BssClient) ListCustomersBalancesDetail(request *model.ListCustomersBalancesDetailRequest) (*model.ListCustomersBalancesDetailResponse, error) {
	requestDef := GenReqDefForListCustomersBalancesDetail()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomersBalancesDetailResponse), nil
	}
}

//功能描述：客户在客户自建平台查询自己的资源详单，用于反映各类资源的消耗情况。资源详单数据有延迟，最大延迟24小时。
func (c *BssClient) ListCustomerselfResourceRecordDetails(request *model.ListCustomerselfResourceRecordDetailsRequest) (*model.ListCustomerselfResourceRecordDetailsResponse, error) {
	requestDef := GenReqDefForListCustomerselfResourceRecordDetails()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerselfResourceRecordDetailsResponse), nil
	}
}

//功能描述：查询资源消费记录（客户）
func (c *BssClient) ListCustomerselfResourceRecords(request *model.ListCustomerselfResourceRecordsRequest) (*model.ListCustomerselfResourceRecordsResponse, error) {
	requestDef := GenReqDefForListCustomerselfResourceRecords()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerselfResourceRecordsResponse), nil
	}
}

//功能描述：查询企业子可回收余额
func (c *BssClient) ListEnterpriseMultiAccount(request *model.ListEnterpriseMultiAccountRequest) (*model.ListEnterpriseMultiAccountResponse, error) {
	requestDef := GenReqDefForListEnterpriseMultiAccount()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnterpriseMultiAccountResponse), nil
	}
}

//功能描述：企业主账号在客户自建平台查询企业组织结构
func (c *BssClient) ListEnterpriseOrganizations(request *model.ListEnterpriseOrganizationsRequest) (*model.ListEnterpriseOrganizationsResponse, error) {
	requestDef := GenReqDefForListEnterpriseOrganizations()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnterpriseOrganizationsResponse), nil
	}
}

//功能描述：企业主账号在客户自建平台查询企业子账号信息列表
func (c *BssClient) ListEnterpriseSubCustomers(request *model.ListEnterpriseSubCustomersRequest) (*model.ListEnterpriseSubCustomersResponse, error) {
	requestDef := GenReqDefForListEnterpriseSubCustomers()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnterpriseSubCustomersResponse), nil
	}
}

//功能描述：查询精英服务商列表
func (c *BssClient) ListIndirectPartners(request *model.ListIndirectPartnersRequest) (*model.ListIndirectPartnersResponse, error) {
	requestDef := GenReqDefForListIndirectPartners()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIndirectPartnersResponse), nil
	}
}

//功能描述：一级经销商查询发给二级经销商的额度
func (c *BssClient) ListIssuedCouponQuotas(request *model.ListIssuedCouponQuotasRequest) (*model.ListIssuedCouponQuotasResponse, error) {
	requestDef := GenReqDefForListIssuedCouponQuotas()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIssuedCouponQuotasResponse), nil
	}
}

//功能描述：查询已发放的优惠券列表
func (c *BssClient) ListIssuedPartnerCoupons(request *model.ListIssuedPartnerCouponsRequest) (*model.ListIssuedPartnerCouponsResponse, error) {
	requestDef := GenReqDefForListIssuedPartnerCoupons()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIssuedPartnerCouponsResponse), nil
	}
}

//功能描述：按需资源询价
func (c *BssClient) ListOnDemandResourceRatings(request *model.ListOnDemandResourceRatingsRequest) (*model.ListOnDemandResourceRatingsResponse, error) {
	requestDef := GenReqDefForListOnDemandResourceRatings()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOnDemandResourceRatingsResponse), nil
	}
}

//功能描述：查询订单详情
func (c *BssClient) ListOrderCouponsByOrderId(request *model.ListOrderCouponsByOrderIdRequest) (*model.ListOrderCouponsByOrderIdResponse, error) {
	requestDef := GenReqDefForListOrderCouponsByOrderId()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrderCouponsByOrderIdResponse), nil
	}
}

//功能描述：查询调账回收记录
func (c *BssClient) ListPartnerAdjustRecords(request *model.ListPartnerAdjustRecordsRequest) (*model.ListPartnerAdjustRecordsResponse, error) {
	requestDef := GenReqDefForListPartnerAdjustRecords()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPartnerAdjustRecordsResponse), nil
	}
}

//功能描述：查询伙伴账户余额
func (c *BssClient) ListPartnerBalances(request *model.ListPartnerBalancesRequest) (*model.ListPartnerBalancesResponse, error) {
	requestDef := GenReqDefForListPartnerBalances()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPartnerBalancesResponse), nil
	}
}

//功能描述：查询优惠券的发放回收记录
func (c *BssClient) ListPartnerCouponsRecord(request *model.ListPartnerCouponsRecordRequest) (*model.ListPartnerCouponsRecordResponse, error) {
	requestDef := GenReqDefForListPartnerCouponsRecord()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPartnerCouponsRecordResponse), nil
	}
}

//功能描述：查询伙伴代付订单列表
func (c *BssClient) ListPartnerPayOrders(request *model.ListPartnerPayOrdersRequest) (*model.ListPartnerPayOrdersResponse, error) {
	requestDef := GenReqDefForListPartnerPayOrders()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPartnerPayOrdersResponse), nil
	}
}

//功能描述：查询客户包年/包月资源列表
func (c *BssClient) ListPayPerUseCustomerResources(request *model.ListPayPerUseCustomerResourcesRequest) (*model.ListPayPerUseCustomerResourcesResponse, error) {
	requestDef := GenReqDefForListPayPerUseCustomerResources()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPayPerUseCustomerResourcesResponse), nil
	}
}

//功能描述：查询邮寄地址
func (c *BssClient) ListPostalAddress(request *model.ListPostalAddressRequest) (*model.ListPostalAddressResponse, error) {
	requestDef := GenReqDefForListPostalAddress()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPostalAddressResponse), nil
	}
}

//功能描述：根据国家查询省份列表
func (c *BssClient) ListProvinces(request *model.ListProvincesRequest) (*model.ListProvincesResponse, error) {
	requestDef := GenReqDefForListProvinces()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProvincesResponse), nil
	}
}

//功能描述：查询优惠券额度列表
func (c *BssClient) ListQuotaCoupons(request *model.ListQuotaCouponsRequest) (*model.ListQuotaCouponsResponse, error) {
	requestDef := GenReqDefForListQuotaCoupons()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotaCouponsResponse), nil
	}
}

//功能描述：包周期资源订购询价
func (c *BssClient) ListRateOnPeriodDetail(request *model.ListRateOnPeriodDetailRequest) (*model.ListRateOnPeriodDetailResponse, error) {
	requestDef := GenReqDefForListRateOnPeriodDetail()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRateOnPeriodDetailResponse), nil
	}
}

//功能描述：查询资源类型
func (c *BssClient) ListResourceTypes(request *model.ListResourceTypesRequest) (*model.ListResourceTypesResponse, error) {
	requestDef := GenReqDefForListResourceTypes()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceTypesResponse), nil
	}
}

//功能描述：查询套餐内使用量
func (c *BssClient) ListResourceUsages(request *model.ListResourceUsagesRequest) (*model.ListResourceUsagesResponse, error) {
	requestDef := GenReqDefForListResourceUsages()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceUsagesResponse), nil
	}
}

//功能描述：根据云服务类型查询资源列表
func (c *BssClient) ListServiceResources(request *model.ListServiceResourcesRequest) (*model.ListServiceResourcesResponse, error) {
	requestDef := GenReqDefForListServiceResources()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceResourcesResponse), nil
	}
}

//功能描述：查询云服务类型列表
func (c *BssClient) ListServiceTypes(request *model.ListServiceTypesRequest) (*model.ListServiceTypesResponse, error) {
	requestDef := GenReqDefForListServiceTypes()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceTypesResponse), nil
	}
}

//功能描述：查询硬件库存
func (c *BssClient) ListSkuInventories(request *model.ListSkuInventoriesRequest) (*model.ListSkuInventoriesResponse, error) {
	requestDef := GenReqDefForListSkuInventories()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSkuInventoriesResponse), nil
	}
}

//功能描述：查询优惠券列表
func (c *BssClient) ListSubCustomerCoupons(request *model.ListSubCustomerCouponsRequest) (*model.ListSubCustomerCouponsResponse, error) {
	requestDef := GenReqDefForListSubCustomerCoupons()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubCustomerCouponsResponse), nil
	}
}

//功能描述：查询伙伴折扣
func (c *BssClient) ListSubCustomerDiscounts(request *model.ListSubCustomerDiscountsRequest) (*model.ListSubCustomerDiscountsResponse, error) {
	requestDef := GenReqDefForListSubCustomerDiscounts()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubCustomerDiscountsResponse), nil
	}
}

//功能描述：查询客户消费记录
func (c *BssClient) ListSubCustomerResFeeRecords(request *model.ListSubCustomerResFeeRecordsRequest) (*model.ListSubCustomerResFeeRecordsResponse, error) {
	requestDef := GenReqDefForListSubCustomerResFeeRecords()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubCustomerResFeeRecordsResponse), nil
	}
}

//功能描述：查询客户列表
func (c *BssClient) ListSubCustomers(request *model.ListSubCustomersRequest) (*model.ListSubCustomersResponse, error) {
	requestDef := GenReqDefForListSubCustomers()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubCustomersResponse), nil
	}
}

//功能描述：合作伙伴可查询客户的消费汇总账单，消费按月汇总
func (c *BssClient) ListSubcustomerMonthlyBills(request *model.ListSubcustomerMonthlyBillsRequest) (*model.ListSubcustomerMonthlyBillsResponse, error) {
	requestDef := GenReqDefForListSubcustomerMonthlyBills()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubcustomerMonthlyBillsResponse), nil
	}
}

//功能描述：支付包周期订单
func (c *BssClient) PayOrders(request *model.PayOrdersRequest) (*model.PayOrdersResponse, error) {
	requestDef := GenReqDefForPayOrders()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PayOrdersResponse), nil
	}
}

//功能描述：一级经销商给二级经销商回收额度
func (c *BssClient) ReclaimCouponQuotas(request *model.ReclaimCouponQuotasRequest) (*model.ReclaimCouponQuotasResponse, error) {
	requestDef := GenReqDefForReclaimCouponQuotas()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReclaimCouponQuotasResponse), nil
	}
}

//功能描述：合作伙伴可以回收二级渠道账户余额
func (c *BssClient) ReclaimIndirectPartnerAccount(request *model.ReclaimIndirectPartnerAccountRequest) (*model.ReclaimIndirectPartnerAccountResponse, error) {
	requestDef := GenReqDefForReclaimIndirectPartnerAccount()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReclaimIndirectPartnerAccountResponse), nil
	}
}

//功能描述：伙伴回收子客户优惠券
func (c *BssClient) ReclaimPartnerCoupons(request *model.ReclaimPartnerCouponsRequest) (*model.ReclaimPartnerCouponsResponse, error) {
	requestDef := GenReqDefForReclaimPartnerCoupons()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReclaimPartnerCouponsResponse), nil
	}
}

//功能描述：企业主账号从企业子账号回收拨款
func (c *BssClient) ReclaimSubEnterpriseAmount(request *model.ReclaimSubEnterpriseAmountRequest) (*model.ReclaimSubEnterpriseAmountResponse, error) {
	requestDef := GenReqDefForReclaimSubEnterpriseAmount()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReclaimSubEnterpriseAmountResponse), nil
	}
}

//功能描述：回收子客户余额（支持一级回收二级的子客户余额）
func (c *BssClient) ReclaimToPartnerAccount(request *model.ReclaimToPartnerAccountRequest) (*model.ReclaimToPartnerAccountResponse, error) {
	requestDef := GenReqDefForReclaimToPartnerAccount()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReclaimToPartnerAccountResponse), nil
	}
}

//功能描述：续订包周期资源
func (c *BssClient) RenewalResources(request *model.RenewalResourcesRequest) (*model.RenewalResourcesResponse, error) {
	requestDef := GenReqDefForRenewalResources()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RenewalResourcesResponse), nil
	}
}

//功能描述：发送短信验证码
func (c *BssClient) SendSmsVerificationCode(request *model.SendSmsVerificationCodeRequest) (*model.SendSmsVerificationCodeResponse, error) {
	requestDef := GenReqDefForSendSmsVerificationCode()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendSmsVerificationCodeResponse), nil
	}
}

//功能描述：发送验证码
func (c *BssClient) SendVerificationMessageCode(request *model.SendVerificationMessageCodeRequest) (*model.SendVerificationMessageCodeResponse, error) {
	requestDef := GenReqDefForSendVerificationMessageCode()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendVerificationMessageCodeResponse), nil
	}
}

//功能描述：查询客户账户余额
func (c *BssClient) ShowCusotmerAccountBalances(request *model.ShowCusotmerAccountBalancesRequest) (*model.ShowCusotmerAccountBalancesResponse, error) {
	requestDef := GenReqDefForShowCusotmerAccountBalances()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCusotmerAccountBalancesResponse), nil
	}
}

//功能描述：客户可以查询自身的消费汇总单的功能，消费按月汇总。每天刷新一次，更新前一天的数据。
func (c *BssClient) ShowCustomerMonthlySum(request *model.ShowCustomerMonthlySumRequest) (*model.ShowCustomerMonthlySumResponse, error) {
	requestDef := GenReqDefForShowCustomerMonthlySum()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCustomerMonthlySumResponse), nil
	}
}

//功能描述：查询订单详情
func (c *BssClient) ShowCustomerOrderDetails(request *model.ShowCustomerOrderDetailsRequest) (*model.ShowCustomerOrderDetailsResponse, error) {
	requestDef := GenReqDefForShowCustomerOrderDetails()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCustomerOrderDetailsResponse), nil
	}
}

//功能描述：查询企业主的可拨款余额
func (c *BssClient) ShowMultiAccountTransferAmount(request *model.ShowMultiAccountTransferAmountRequest) (*model.ShowMultiAccountTransferAmountResponse, error) {
	requestDef := GenReqDefForShowMultiAccountTransferAmount()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMultiAccountTransferAmountResponse), nil
	}
}

//功能描述：查询实名认证审核结果
func (c *BssClient) ShowRealnameAuthenticationReviewResult(request *model.ShowRealnameAuthenticationReviewResultRequest) (*model.ShowRealnameAuthenticationReviewResultResponse, error) {
	requestDef := GenReqDefForShowRealnameAuthenticationReviewResult()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRealnameAuthenticationReviewResultResponse), nil
	}
}

//功能描述：查询退款订单的金额详情
func (c *BssClient) ShowRefundOrderDetails(request *model.ShowRefundOrderDetailsRequest) (*model.ShowRefundOrderDetailsResponse, error) {
	requestDef := GenReqDefForShowRefundOrderDetails()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRefundOrderDetailsResponse), nil
	}
}

//功能描述：向精英服务商发放代金券额度
func (c *BssClient) UpdateCouponQuotas(request *model.UpdateCouponQuotasRequest) (*model.UpdateCouponQuotasResponse, error) {
	requestDef := GenReqDefForUpdateCouponQuotas()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCouponQuotasResponse), nil
	}
}

//功能描述：从伙伴账户调账给子客户
func (c *BssClient) UpdateCustomerAccountAmount(request *model.UpdateCustomerAccountAmountRequest) (*model.UpdateCustomerAccountAmountResponse, error) {
	requestDef := GenReqDefForUpdateCustomerAccountAmount()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCustomerAccountAmountResponse), nil
	}
}

//功能描述：从伙伴账户调账给二级渠道
func (c *BssClient) UpdateIndirectPartnerAccount(request *model.UpdateIndirectPartnerAccountRequest) (*model.UpdateIndirectPartnerAccountResponse, error) {
	requestDef := GenReqDefForUpdateIndirectPartnerAccount()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIndirectPartnerAccountResponse), nil
	}
}

//功能描述：修改邮寄地址
func (c *BssClient) UpdatePostal(request *model.UpdatePostalRequest) (*model.UpdatePostalResponse, error) {
	requestDef := GenReqDefForUpdatePostal()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePostalResponse), nil
	}
}

//功能描述：企业主账号向企业子账号拨款
func (c *BssClient) UpdateSubEnterpriseAmount(request *model.UpdateSubEnterpriseAmountRequest) (*model.UpdateSubEnterpriseAmountResponse, error) {
	requestDef := GenReqDefForUpdateSubEnterpriseAmount()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubEnterpriseAmountResponse), nil
	}
}
