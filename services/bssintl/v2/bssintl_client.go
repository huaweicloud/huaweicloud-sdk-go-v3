package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bssintl/v2/model"
)

type BssintlClient struct {
	HcClient *http_client.HcHttpClient
}

func NewBssintlClient(hcClient *http_client.HcHttpClient) *BssintlClient {
	return &BssintlClient{HcClient: hcClient}
}

func BssintlClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

//功能描述：客户可以设置包年/包月资源到期后转为按需资源计费
func (c *BssintlClient) AutoRenewalResources(request *model.AutoRenewalResourcesRequest) (*model.AutoRenewalResourcesResponse, error) {
	requestDef := GenReqDefForAutoRenewalResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AutoRenewalResourcesResponse), nil
	}
}

//功能描述：取消包年/包月资源自动续费
func (c *BssintlClient) CancelAutoRenewalResources(request *model.CancelAutoRenewalResourcesRequest) (*model.CancelAutoRenewalResourcesResponse, error) {
	requestDef := GenReqDefForCancelAutoRenewalResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelAutoRenewalResourcesResponse), nil
	}
}

//功能描述：客户可以对待支付的订单进行取消操作
func (c *BssintlClient) CancelCustomerOrder(request *model.CancelCustomerOrderRequest) (*model.CancelCustomerOrderResponse, error) {
	requestDef := GenReqDefForCancelCustomerOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelCustomerOrderResponse), nil
	}
}

//功能描述：客户购买包年/包月资源后，支持客户退订包年/包月实例。退订资源实例包括资源续费部分和当前正在使用的部分，退订后资源将无法使用
func (c *BssintlClient) CancelResourcesSubscription(request *model.CancelResourcesSubscriptionRequest) (*model.CancelResourcesSubscriptionResponse, error) {
	requestDef := GenReqDefForCancelResourcesSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelResourcesSubscriptionResponse), nil
	}
}

//功能描述：客户可以进行实名认证变更申请。
func (c *BssintlClient) ChangeEnterpriseRealnameAuthentication(request *model.ChangeEnterpriseRealnameAuthenticationRequest) (*model.ChangeEnterpriseRealnameAuthenticationResponse, error) {
	requestDef := GenReqDefForChangeEnterpriseRealnameAuthentication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeEnterpriseRealnameAuthenticationResponse), nil
	}
}

//功能描述：客户注册时可检查客户的登录名称、手机号或者邮箱是否可以用于注册。
func (c *BssintlClient) CheckUserIdentity(request *model.CheckUserIdentityRequest) (*model.CheckUserIdentityResponse, error) {
	requestDef := GenReqDefForCheckUserIdentity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckUserIdentityResponse), nil
	}
}

//功能描述：企业客户可以进行企业实名认证申请。
func (c *BssintlClient) CreateEnterpriseRealnameAuthentication(request *model.CreateEnterpriseRealnameAuthenticationRequest) (*model.CreateEnterpriseRealnameAuthenticationResponse, error) {
	requestDef := GenReqDefForCreateEnterpriseRealnameAuthentication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnterpriseRealnameAuthenticationResponse), nil
	}
}

//功能描述：个人客户可以进行个人实名认证申请。
func (c *BssintlClient) CreatePersonalRealnameAuth(request *model.CreatePersonalRealnameAuthRequest) (*model.CreatePersonalRealnameAuthResponse, error) {
	requestDef := GenReqDefForCreatePersonalRealnameAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePersonalRealnameAuthResponse), nil
	}
}

//功能描述：在伙伴销售平台创建客户时同步创建华为云账号，并将客户在伙伴销售平台上的账号与华为云账号进行映射。同时，创建的华为云账号与伙伴账号关联绑定。华为云伙伴能力中心（一级经销商）可以注册精英服务商伙伴（二级经销商）的子客户。注册完成后，子客户可以自动和精英服务商伙伴绑定。
func (c *BssintlClient) CreateSubCustomer(request *model.CreateSubCustomerRequest) (*model.CreateSubCustomerResponse, error) {
	requestDef := GenReqDefForCreateSubCustomer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubCustomerResponse), nil
	}
}

//功能描述：冻结伙伴子客户
func (c *BssintlClient) FreezeSubCustomers(request *model.FreezeSubCustomersRequest) (*model.FreezeSubCustomersResponse, error) {
	requestDef := GenReqDefForFreezeSubCustomers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.FreezeSubCustomersResponse), nil
	}
}

//功能描述：伙伴在伙伴销售平台上查询使用量单位的进制转换信息，用于不同度量单位之间的转换。
func (c *BssintlClient) ListConversions(request *model.ListConversionsRequest) (*model.ListConversionsResponse, error) {
	requestDef := GenReqDefForListConversions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConversionsResponse), nil
	}
}

//功能描述：客户在伙伴销售平台查询已开通的按需资源
func (c *BssintlClient) ListCustomerOnDemandResources(request *model.ListCustomerOnDemandResourcesRequest) (*model.ListCustomerOnDemandResourcesResponse, error) {
	requestDef := GenReqDefForListCustomerOnDemandResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerOnDemandResourcesResponse), nil
	}
}

//功能描述：客户购买包年包月资源后，可以查看待审核、处理中、已取消、已完成和待支付等状态的订单
func (c *BssintlClient) ListCustomerOrders(request *model.ListCustomerOrdersRequest) (*model.ListCustomerOrdersResponse, error) {
	requestDef := GenReqDefForListCustomerOrders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerOrdersResponse), nil
	}
}

//功能描述：客户在客户自建平台查询自己的资源详单，用于反映各类资源的消耗情况。
func (c *BssintlClient) ListCustomerselfResourceRecordDetails(request *model.ListCustomerselfResourceRecordDetailsRequest) (*model.ListCustomerselfResourceRecordDetailsResponse, error) {
	requestDef := GenReqDefForListCustomerselfResourceRecordDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerselfResourceRecordDetailsResponse), nil
	}
}

//功能描述：客户在客户自建平台查询每个资源的消费明细数据
func (c *BssintlClient) ListCustomerselfResourceRecords(request *model.ListCustomerselfResourceRecordsRequest) (*model.ListCustomerselfResourceRecordsResponse, error) {
	requestDef := GenReqDefForListCustomerselfResourceRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerselfResourceRecordsResponse), nil
	}
}

//功能描述：客户在自建平台查询资源包列表。
func (c *BssintlClient) ListFreeResourceInfos(request *model.ListFreeResourceInfosRequest) (*model.ListFreeResourceInfosResponse, error) {
	requestDef := GenReqDefForListFreeResourceInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFreeResourceInfosResponse), nil
	}
}

//功能描述：客户在自建平台查询客户自己的资源包列表
func (c *BssintlClient) ListFreeResourceUsages(request *model.ListFreeResourceUsagesRequest) (*model.ListFreeResourceUsagesResponse, error) {
	requestDef := GenReqDefForListFreeResourceUsages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFreeResourceUsagesResponse), nil
	}
}

//功能描述：查询发票列表
func (c *BssintlClient) ListInvoices(request *model.ListInvoicesRequest) (*model.ListInvoicesResponse, error) {
	requestDef := GenReqDefForListInvoices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInvoicesResponse), nil
	}
}

//功能描述：伙伴在伙伴销售平台上查询资源使用量的度量单位及名称，度量单位类型等。
func (c *BssintlClient) ListMeasureUnits(request *model.ListMeasureUnitsRequest) (*model.ListMeasureUnitsResponse, error) {
	requestDef := GenReqDefForListMeasureUnits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMeasureUnitsResponse), nil
	}
}

//功能描述：客户可以查询自身的消费汇总单的功能，消费按月汇总。
func (c *BssintlClient) ListMonthlyExpenditures(request *model.ListMonthlyExpendituresRequest) (*model.ListMonthlyExpendituresResponse, error) {
	requestDef := GenReqDefForListMonthlyExpenditures()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMonthlyExpendituresResponse), nil
	}
}

//功能描述：按需资源询价
func (c *BssintlClient) ListOnDemandResourceRatings(request *model.ListOnDemandResourceRatingsRequest) (*model.ListOnDemandResourceRatingsResponse, error) {
	requestDef := GenReqDefForListOnDemandResourceRatings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOnDemandResourceRatingsResponse), nil
	}
}

//功能描述：功能介绍客户在伙伴销售平台支付待支付订单时，查询可使用的折扣。只返回商务合同折扣和伙伴授权折扣客户在客户自建平台查看订单可用的优惠券列表。
func (c *BssintlClient) ListOrderDiscounts(request *model.ListOrderDiscountsRequest) (*model.ListOrderDiscountsResponse, error) {
	requestDef := GenReqDefForListOrderDiscounts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrderDiscountsResponse), nil
	}
}

//功能描述：客户在客户自建平台查询某个或所有的包年/包月资源
func (c *BssintlClient) ListPayPerUseCustomerResources(request *model.ListPayPerUseCustomerResourcesRequest) (*model.ListPayPerUseCustomerResourcesResponse, error) {
	requestDef := GenReqDefForListPayPerUseCustomerResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPayPerUseCustomerResourcesResponse), nil
	}
}

//功能描述：伙伴可以查询伙伴月度消费账单
func (c *BssintlClient) ListPostpaidBillSum(request *model.ListPostpaidBillSumRequest) (*model.ListPostpaidBillSumResponse, error) {
	requestDef := GenReqDefForListPostpaidBillSum()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPostpaidBillSumResponse), nil
	}
}

//功能描述：客户在自建平台按照条件查询包年/包月产品开通时候的价格
func (c *BssintlClient) ListRateOnPeriodDetail(request *model.ListRateOnPeriodDetailRequest) (*model.ListRateOnPeriodDetailResponse, error) {
	requestDef := GenReqDefForListRateOnPeriodDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRateOnPeriodDetailResponse), nil
	}
}

//伙伴在伙伴销售平台查询资源类型的列表。
func (c *BssintlClient) ListResourceTypes(request *model.ListResourceTypesRequest) (*model.ListResourceTypesResponse, error) {
	requestDef := GenReqDefForListResourceTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceTypesResponse), nil
	}
}

//功能描述：伙伴在伙伴销售平台根据云服务类型查询关联的资源类型编码和名称，用于查询按需产品的价格或包年/包月产品的价格。
func (c *BssintlClient) ListServiceResources(request *model.ListServiceResourcesRequest) (*model.ListServiceResourcesResponse, error) {
	requestDef := GenReqDefForListServiceResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceResourcesResponse), nil
	}
}

//伙伴在伙伴销售平台查询云服务类型的列表。
func (c *BssintlClient) ListServiceTypes(request *model.ListServiceTypesRequest) (*model.ListServiceTypesResponse, error) {
	requestDef := GenReqDefForListServiceTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceTypesResponse), nil
	}
}

//功能描述：伙伴/客户可以查询自身的优惠券信息。
func (c *BssintlClient) ListSubCustomerCoupons(request *model.ListSubCustomerCouponsRequest) (*model.ListSubCustomerCouponsResponse, error) {
	requestDef := GenReqDefForListSubCustomerCoupons()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubCustomerCouponsResponse), nil
	}
}

//功能描述：伙伴可以查询合作伙伴的客户信息列表。
func (c *BssintlClient) ListSubCustomers(request *model.ListSubCustomersRequest) (*model.ListSubCustomersResponse, error) {
	requestDef := GenReqDefForListSubCustomers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubCustomersResponse), nil
	}
}

//功能描述：伙伴在伙伴销售平台查询资源的使用量类型列表。
func (c *BssintlClient) ListUsageTypes(request *model.ListUsageTypesRequest) (*model.ListUsageTypesResponse, error) {
	requestDef := GenReqDefForListUsageTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsageTypesResponse), nil
	}
}

//功能描述：客户可以对待支付状态的包年/包月产品订单进行支付
func (c *BssintlClient) PayOrders(request *model.PayOrdersRequest) (*model.PayOrdersResponse, error) {
	requestDef := GenReqDefForPayOrders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PayOrdersResponse), nil
	}
}

//功能描述：客户的包年包/月资源即将到期时，可进行包年/包月资源的续订
func (c *BssintlClient) RenewalResources(request *model.RenewalResourcesRequest) (*model.RenewalResourcesResponse, error) {
	requestDef := GenReqDefForRenewalResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RenewalResourcesResponse), nil
	}
}

//功能描述：客户注册时，如果填写了邮箱，可以向对应的邮箱发送注册验证码，校验信息的正确性。
func (c *BssintlClient) SendVerificationMessageCode(request *model.SendVerificationMessageCodeRequest) (*model.SendVerificationMessageCodeResponse, error) {
	requestDef := GenReqDefForSendVerificationMessageCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendVerificationMessageCodeResponse), nil
	}
}

//功能描述：客户可以查询自身的账户余额。
func (c *BssintlClient) ShowCustomerAccountBalances(request *model.ShowCustomerAccountBalancesRequest) (*model.ShowCustomerAccountBalancesResponse, error) {
	requestDef := GenReqDefForShowCustomerAccountBalances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCustomerAccountBalancesResponse), nil
	}
}

//功能描述：客户可以查看订单详情
func (c *BssintlClient) ShowCustomerOrderDetails(request *model.ShowCustomerOrderDetailsRequest) (*model.ShowCustomerOrderDetailsResponse, error) {
	requestDef := GenReqDefForShowCustomerOrderDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCustomerOrderDetailsResponse), nil
	}
}

//功能描述：如果实名认证申请或实名认证变更申请的响应中，显示需要人工审核，使用该接口查询审核结果。
func (c *BssintlClient) ShowRealnameAuthenticationReviewResult(request *model.ShowRealnameAuthenticationReviewResultRequest) (*model.ShowRealnameAuthenticationReviewResultResponse, error) {
	requestDef := GenReqDefForShowRealnameAuthenticationReviewResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRealnameAuthenticationReviewResultResponse), nil
	}
}

//功能描述：客户在伙伴销售平台查询某次退订订单或者降配订单的退款金额来自哪些资源和对应订单
func (c *BssintlClient) ShowRefundOrderDetails(request *model.ShowRefundOrderDetailsRequest) (*model.ShowRefundOrderDetailsResponse, error) {
	requestDef := GenReqDefForShowRefundOrderDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRefundOrderDetailsResponse), nil
	}
}

//功能描述：查询客户预算
func (c *BssintlClient) ShowSubCustomerBudget(request *model.ShowSubCustomerBudgetRequest) (*model.ShowSubCustomerBudgetResponse, error) {
	requestDef := GenReqDefForShowSubCustomerBudget()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSubCustomerBudgetResponse), nil
	}
}

//功能描述：解冻伙伴子客户
func (c *BssintlClient) UnfreezeSubCustomers(request *model.UnfreezeSubCustomersRequest) (*model.UnfreezeSubCustomersResponse, error) {
	requestDef := GenReqDefForUnfreezeSubCustomers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnfreezeSubCustomersResponse), nil
	}
}

//功能描述：客户可以设置包年/包月资源到期后转为按需资源计费。包年/包月计费模式到期后，按需的计费模式即生效
func (c *BssintlClient) UpdatePeriodToOnDemand(request *model.UpdatePeriodToOnDemandRequest) (*model.UpdatePeriodToOnDemandResponse, error) {
	requestDef := GenReqDefForUpdatePeriodToOnDemand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePeriodToOnDemandResponse), nil
	}
}

//功能描述：设置客户预算
func (c *BssintlClient) UpdateSubCustomerBudget(request *model.UpdateSubCustomerBudgetRequest) (*model.UpdateSubCustomerBudgetResponse, error) {
	requestDef := GenReqDefForUpdateSubCustomerBudget()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubCustomerBudgetResponse), nil
	}
}
