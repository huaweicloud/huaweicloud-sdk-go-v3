package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bssintl/v2/model"
)

type BssintlClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewBssintlClient(hcClient *httpclient.HcHttpClient) *BssintlClient {
	return &BssintlClient{HcClient: hcClient}
}

func BssintlClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// AutoRenewalResources 设置包年/包月资源自动续费
//
// 功能描述：客户可以设置包年/包月资源到期后转为按需资源计费
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) AutoRenewalResources(request *model.AutoRenewalResourcesRequest) (*model.AutoRenewalResourcesResponse, error) {
	requestDef := GenReqDefForAutoRenewalResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AutoRenewalResourcesResponse), nil
	}
}

// AutoRenewalResourcesInvoker 设置包年/包月资源自动续费
func (c *BssintlClient) AutoRenewalResourcesInvoker(request *model.AutoRenewalResourcesRequest) *AutoRenewalResourcesInvoker {
	requestDef := GenReqDefForAutoRenewalResources()
	return &AutoRenewalResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelAutoRenewalResources 取消包年/包月资源自动续费
//
// 功能描述：取消包年/包月资源自动续费
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) CancelAutoRenewalResources(request *model.CancelAutoRenewalResourcesRequest) (*model.CancelAutoRenewalResourcesResponse, error) {
	requestDef := GenReqDefForCancelAutoRenewalResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelAutoRenewalResourcesResponse), nil
	}
}

// CancelAutoRenewalResourcesInvoker 取消包年/包月资源自动续费
func (c *BssintlClient) CancelAutoRenewalResourcesInvoker(request *model.CancelAutoRenewalResourcesRequest) *CancelAutoRenewalResourcesInvoker {
	requestDef := GenReqDefForCancelAutoRenewalResources()
	return &CancelAutoRenewalResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelCustomerOrder 取消待支付订单
//
// 功能描述：客户可以对待支付的订单进行取消操作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) CancelCustomerOrder(request *model.CancelCustomerOrderRequest) (*model.CancelCustomerOrderResponse, error) {
	requestDef := GenReqDefForCancelCustomerOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelCustomerOrderResponse), nil
	}
}

// CancelCustomerOrderInvoker 取消待支付订单
func (c *BssintlClient) CancelCustomerOrderInvoker(request *model.CancelCustomerOrderRequest) *CancelCustomerOrderInvoker {
	requestDef := GenReqDefForCancelCustomerOrder()
	return &CancelCustomerOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelResourcesSubscription 退订包年/包月资源
//
// 功能描述：客户购买包年/包月资源后，支持客户退订包年/包月实例。退订资源实例包括资源续费部分和当前正在使用的部分，退订后资源将无法使用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) CancelResourcesSubscription(request *model.CancelResourcesSubscriptionRequest) (*model.CancelResourcesSubscriptionResponse, error) {
	requestDef := GenReqDefForCancelResourcesSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelResourcesSubscriptionResponse), nil
	}
}

// CancelResourcesSubscriptionInvoker 退订包年/包月资源
func (c *BssintlClient) CancelResourcesSubscriptionInvoker(request *model.CancelResourcesSubscriptionRequest) *CancelResourcesSubscriptionInvoker {
	requestDef := GenReqDefForCancelResourcesSubscription()
	return &CancelResourcesSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeEnterpriseRealnameAuthentication 申请实名认证变更
//
// 功能描述：客户可以进行实名认证变更申请。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ChangeEnterpriseRealnameAuthentication(request *model.ChangeEnterpriseRealnameAuthenticationRequest) (*model.ChangeEnterpriseRealnameAuthenticationResponse, error) {
	requestDef := GenReqDefForChangeEnterpriseRealnameAuthentication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeEnterpriseRealnameAuthenticationResponse), nil
	}
}

// ChangeEnterpriseRealnameAuthenticationInvoker 申请实名认证变更
func (c *BssintlClient) ChangeEnterpriseRealnameAuthenticationInvoker(request *model.ChangeEnterpriseRealnameAuthenticationRequest) *ChangeEnterpriseRealnameAuthenticationInvoker {
	requestDef := GenReqDefForChangeEnterpriseRealnameAuthentication()
	return &ChangeEnterpriseRealnameAuthenticationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckUserIdentity 校验客户注册信息
//
// 功能描述：客户注册时可检查客户的登录名称、手机号或者邮箱是否可以用于注册。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) CheckUserIdentity(request *model.CheckUserIdentityRequest) (*model.CheckUserIdentityResponse, error) {
	requestDef := GenReqDefForCheckUserIdentity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckUserIdentityResponse), nil
	}
}

// CheckUserIdentityInvoker 校验客户注册信息
func (c *BssintlClient) CheckUserIdentityInvoker(request *model.CheckUserIdentityRequest) *CheckUserIdentityInvoker {
	requestDef := GenReqDefForCheckUserIdentity()
	return &CheckUserIdentityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEnterpriseProjectAuth 开通客户企业项目权限
//
// 客户在自建平台开通客户企业项目权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) CreateEnterpriseProjectAuth(request *model.CreateEnterpriseProjectAuthRequest) (*model.CreateEnterpriseProjectAuthResponse, error) {
	requestDef := GenReqDefForCreateEnterpriseProjectAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnterpriseProjectAuthResponse), nil
	}
}

// CreateEnterpriseProjectAuthInvoker 开通客户企业项目权限
func (c *BssintlClient) CreateEnterpriseProjectAuthInvoker(request *model.CreateEnterpriseProjectAuthRequest) *CreateEnterpriseProjectAuthInvoker {
	requestDef := GenReqDefForCreateEnterpriseProjectAuth()
	return &CreateEnterpriseProjectAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEnterpriseRealnameAuthentication 申请企业实名认证
//
// 功能描述：企业客户可以进行企业实名认证申请。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) CreateEnterpriseRealnameAuthentication(request *model.CreateEnterpriseRealnameAuthenticationRequest) (*model.CreateEnterpriseRealnameAuthenticationResponse, error) {
	requestDef := GenReqDefForCreateEnterpriseRealnameAuthentication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnterpriseRealnameAuthenticationResponse), nil
	}
}

// CreateEnterpriseRealnameAuthenticationInvoker 申请企业实名认证
func (c *BssintlClient) CreateEnterpriseRealnameAuthenticationInvoker(request *model.CreateEnterpriseRealnameAuthenticationRequest) *CreateEnterpriseRealnameAuthenticationInvoker {
	requestDef := GenReqDefForCreateEnterpriseRealnameAuthentication()
	return &CreateEnterpriseRealnameAuthenticationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePersonalRealnameAuth 申请个人实名认证
//
// 功能描述：个人客户可以进行个人实名认证申请。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) CreatePersonalRealnameAuth(request *model.CreatePersonalRealnameAuthRequest) (*model.CreatePersonalRealnameAuthResponse, error) {
	requestDef := GenReqDefForCreatePersonalRealnameAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePersonalRealnameAuthResponse), nil
	}
}

// CreatePersonalRealnameAuthInvoker 申请个人实名认证
func (c *BssintlClient) CreatePersonalRealnameAuthInvoker(request *model.CreatePersonalRealnameAuthRequest) *CreatePersonalRealnameAuthInvoker {
	requestDef := GenReqDefForCreatePersonalRealnameAuth()
	return &CreatePersonalRealnameAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSubCustomer 创建客户
//
// 功能描述：在伙伴销售平台创建客户时同步创建华为云账号，并将客户在伙伴销售平台上的账号与华为云账号进行映射。同时，创建的华为云账号与伙伴账号关联绑定。华为云伙伴能力中心（一级经销商）可以注册云经销商伙伴（二级经销商）的子客户。注册完成后，子客户可以自动和云经销商伙伴绑定。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) CreateSubCustomer(request *model.CreateSubCustomerRequest) (*model.CreateSubCustomerResponse, error) {
	requestDef := GenReqDefForCreateSubCustomer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubCustomerResponse), nil
	}
}

// CreateSubCustomerInvoker 创建客户
func (c *BssintlClient) CreateSubCustomerInvoker(request *model.CreateSubCustomerRequest) *CreateSubCustomerInvoker {
	requestDef := GenReqDefForCreateSubCustomer()
	return &CreateSubCustomerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// FreezeSubCustomers 冻结客户账号
//
// 功能描述：冻结伙伴子客户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) FreezeSubCustomers(request *model.FreezeSubCustomersRequest) (*model.FreezeSubCustomersResponse, error) {
	requestDef := GenReqDefForFreezeSubCustomers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.FreezeSubCustomersResponse), nil
	}
}

// FreezeSubCustomersInvoker 冻结客户账号
func (c *BssintlClient) FreezeSubCustomersInvoker(request *model.FreezeSubCustomersRequest) *FreezeSubCustomersInvoker {
	requestDef := GenReqDefForFreezeSubCustomers()
	return &FreezeSubCustomersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConversions 查询使用量单位进制
//
// 功能描述：伙伴在伙伴销售平台上查询使用量单位的进制转换信息，用于不同度量单位之间的转换。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListConversions(request *model.ListConversionsRequest) (*model.ListConversionsResponse, error) {
	requestDef := GenReqDefForListConversions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConversionsResponse), nil
	}
}

// ListConversionsInvoker 查询使用量单位进制
func (c *BssintlClient) ListConversionsInvoker(request *model.ListConversionsRequest) *ListConversionsInvoker {
	requestDef := GenReqDefForListConversions()
	return &ListConversionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCosts 查询成本数据
//
// 客户在自建平台查询成本分析数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListCosts(request *model.ListCostsRequest) (*model.ListCostsResponse, error) {
	requestDef := GenReqDefForListCosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCostsResponse), nil
	}
}

// ListCostsInvoker 查询成本数据
func (c *BssintlClient) ListCostsInvoker(request *model.ListCostsRequest) *ListCostsInvoker {
	requestDef := GenReqDefForListCosts()
	return &ListCostsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomerOnDemandResources 查询客户按需资源列表
//
// 功能描述：客户在伙伴销售平台查询已开通的按需资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListCustomerOnDemandResources(request *model.ListCustomerOnDemandResourcesRequest) (*model.ListCustomerOnDemandResourcesResponse, error) {
	requestDef := GenReqDefForListCustomerOnDemandResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerOnDemandResourcesResponse), nil
	}
}

// ListCustomerOnDemandResourcesInvoker 查询客户按需资源列表
func (c *BssintlClient) ListCustomerOnDemandResourcesInvoker(request *model.ListCustomerOnDemandResourcesRequest) *ListCustomerOnDemandResourcesInvoker {
	requestDef := GenReqDefForListCustomerOnDemandResources()
	return &ListCustomerOnDemandResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomerOrders 查询订单列表
//
// 功能描述：客户购买包年包月资源后，可以查看待审核、处理中、已取消、已完成和待支付等状态的订单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListCustomerOrders(request *model.ListCustomerOrdersRequest) (*model.ListCustomerOrdersResponse, error) {
	requestDef := GenReqDefForListCustomerOrders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerOrdersResponse), nil
	}
}

// ListCustomerOrdersInvoker 查询订单列表
func (c *BssintlClient) ListCustomerOrdersInvoker(request *model.ListCustomerOrdersRequest) *ListCustomerOrdersInvoker {
	requestDef := GenReqDefForListCustomerOrders()
	return &ListCustomerOrdersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomerselfResourceRecordDetails 查询资源详单
//
// 功能描述：客户在客户自建平台查询自己的资源详单，用于反映各类资源的消耗情况。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListCustomerselfResourceRecordDetails(request *model.ListCustomerselfResourceRecordDetailsRequest) (*model.ListCustomerselfResourceRecordDetailsResponse, error) {
	requestDef := GenReqDefForListCustomerselfResourceRecordDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerselfResourceRecordDetailsResponse), nil
	}
}

// ListCustomerselfResourceRecordDetailsInvoker 查询资源详单
func (c *BssintlClient) ListCustomerselfResourceRecordDetailsInvoker(request *model.ListCustomerselfResourceRecordDetailsRequest) *ListCustomerselfResourceRecordDetailsInvoker {
	requestDef := GenReqDefForListCustomerselfResourceRecordDetails()
	return &ListCustomerselfResourceRecordDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomerselfResourceRecords 查询资源消费记录
//
// 功能描述：客户在客户自建平台查询每个资源的消费明细数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListCustomerselfResourceRecords(request *model.ListCustomerselfResourceRecordsRequest) (*model.ListCustomerselfResourceRecordsResponse, error) {
	requestDef := GenReqDefForListCustomerselfResourceRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerselfResourceRecordsResponse), nil
	}
}

// ListCustomerselfResourceRecordsInvoker 查询资源消费记录
func (c *BssintlClient) ListCustomerselfResourceRecordsInvoker(request *model.ListCustomerselfResourceRecordsRequest) *ListCustomerselfResourceRecordsInvoker {
	requestDef := GenReqDefForListCustomerselfResourceRecords()
	return &ListCustomerselfResourceRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFreeResourceInfos 查询资源包列表
//
// 功能描述：客户在自建平台查询资源包列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListFreeResourceInfos(request *model.ListFreeResourceInfosRequest) (*model.ListFreeResourceInfosResponse, error) {
	requestDef := GenReqDefForListFreeResourceInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFreeResourceInfosResponse), nil
	}
}

// ListFreeResourceInfosInvoker 查询资源包列表
func (c *BssintlClient) ListFreeResourceInfosInvoker(request *model.ListFreeResourceInfosRequest) *ListFreeResourceInfosInvoker {
	requestDef := GenReqDefForListFreeResourceInfos()
	return &ListFreeResourceInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFreeResourceUsages 查询资源内使用量
//
// 功能描述：客户在自建平台查询客户自己的资源包列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListFreeResourceUsages(request *model.ListFreeResourceUsagesRequest) (*model.ListFreeResourceUsagesResponse, error) {
	requestDef := GenReqDefForListFreeResourceUsages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFreeResourceUsagesResponse), nil
	}
}

// ListFreeResourceUsagesInvoker 查询资源内使用量
func (c *BssintlClient) ListFreeResourceUsagesInvoker(request *model.ListFreeResourceUsagesRequest) *ListFreeResourceUsagesInvoker {
	requestDef := GenReqDefForListFreeResourceUsages()
	return &ListFreeResourceUsagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFreeResourcesUsageRecords 查询资源包使用明细
//
// 客户在自建平台查询资源包使用明细。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListFreeResourcesUsageRecords(request *model.ListFreeResourcesUsageRecordsRequest) (*model.ListFreeResourcesUsageRecordsResponse, error) {
	requestDef := GenReqDefForListFreeResourcesUsageRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFreeResourcesUsageRecordsResponse), nil
	}
}

// ListFreeResourcesUsageRecordsInvoker 查询资源包使用明细
func (c *BssintlClient) ListFreeResourcesUsageRecordsInvoker(request *model.ListFreeResourcesUsageRecordsRequest) *ListFreeResourcesUsageRecordsInvoker {
	requestDef := GenReqDefForListFreeResourcesUsageRecords()
	return &ListFreeResourcesUsageRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIndirectPartners 查询云经销商列表
//
// 华为云总经销商（一级经销商）可以查询云经销商（二级经销商）列表。
//
// 一级经销商在伙伴中心查询二级经销商列表的方式请参见[这里](https://support.huaweicloud.com/usermanual-bpconsole/dp_120210.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListIndirectPartners(request *model.ListIndirectPartnersRequest) (*model.ListIndirectPartnersResponse, error) {
	requestDef := GenReqDefForListIndirectPartners()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIndirectPartnersResponse), nil
	}
}

// ListIndirectPartnersInvoker 查询云经销商列表
func (c *BssintlClient) ListIndirectPartnersInvoker(request *model.ListIndirectPartnersRequest) *ListIndirectPartnersInvoker {
	requestDef := GenReqDefForListIndirectPartners()
	return &ListIndirectPartnersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInvoices 查询发票列表
//
// 功能描述：查询发票列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListInvoices(request *model.ListInvoicesRequest) (*model.ListInvoicesResponse, error) {
	requestDef := GenReqDefForListInvoices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInvoicesResponse), nil
	}
}

// ListInvoicesInvoker 查询发票列表
func (c *BssintlClient) ListInvoicesInvoker(request *model.ListInvoicesRequest) *ListInvoicesInvoker {
	requestDef := GenReqDefForListInvoices()
	return &ListInvoicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMeasureUnits 查询使用量单位列表
//
// 功能描述：伙伴在伙伴销售平台上查询资源使用量的度量单位及名称，度量单位类型等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListMeasureUnits(request *model.ListMeasureUnitsRequest) (*model.ListMeasureUnitsResponse, error) {
	requestDef := GenReqDefForListMeasureUnits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMeasureUnitsResponse), nil
	}
}

// ListMeasureUnitsInvoker 查询使用量单位列表
func (c *BssintlClient) ListMeasureUnitsInvoker(request *model.ListMeasureUnitsRequest) *ListMeasureUnitsInvoker {
	requestDef := GenReqDefForListMeasureUnits()
	return &ListMeasureUnitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMonthlyExpenditures 查询消费汇总(客户)
//
// 功能描述：客户可以查询自身的消费汇总单的功能，消费按月汇总。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListMonthlyExpenditures(request *model.ListMonthlyExpendituresRequest) (*model.ListMonthlyExpendituresResponse, error) {
	requestDef := GenReqDefForListMonthlyExpenditures()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMonthlyExpendituresResponse), nil
	}
}

// ListMonthlyExpendituresInvoker 查询消费汇总(客户)
func (c *BssintlClient) ListMonthlyExpendituresInvoker(request *model.ListMonthlyExpendituresRequest) *ListMonthlyExpendituresInvoker {
	requestDef := GenReqDefForListMonthlyExpenditures()
	return &ListMonthlyExpendituresInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOnDemandResourceRatings 查询按需产品价格
//
// 功能描述：按需资源询价
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListOnDemandResourceRatings(request *model.ListOnDemandResourceRatingsRequest) (*model.ListOnDemandResourceRatingsResponse, error) {
	requestDef := GenReqDefForListOnDemandResourceRatings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOnDemandResourceRatingsResponse), nil
	}
}

// ListOnDemandResourceRatingsInvoker 查询按需产品价格
func (c *BssintlClient) ListOnDemandResourceRatingsInvoker(request *model.ListOnDemandResourceRatingsRequest) *ListOnDemandResourceRatingsInvoker {
	requestDef := GenReqDefForListOnDemandResourceRatings()
	return &ListOnDemandResourceRatingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOrderDiscounts 查询订单可用折扣
//
// 功能描述：功能介绍客户在伙伴销售平台支付待支付订单时，查询可使用的折扣。只返回商务合同折扣和伙伴授权折扣客户在客户自建平台查看订单可用的优惠券列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListOrderDiscounts(request *model.ListOrderDiscountsRequest) (*model.ListOrderDiscountsResponse, error) {
	requestDef := GenReqDefForListOrderDiscounts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrderDiscountsResponse), nil
	}
}

// ListOrderDiscountsInvoker 查询订单可用折扣
func (c *BssintlClient) ListOrderDiscountsInvoker(request *model.ListOrderDiscountsRequest) *ListOrderDiscountsInvoker {
	requestDef := GenReqDefForListOrderDiscounts()
	return &ListOrderDiscountsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPayPerUseCustomerResources 查询客户包年/包月资源列表
//
// 功能描述：客户在客户自建平台查询某个或所有的包年/包月资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListPayPerUseCustomerResources(request *model.ListPayPerUseCustomerResourcesRequest) (*model.ListPayPerUseCustomerResourcesResponse, error) {
	requestDef := GenReqDefForListPayPerUseCustomerResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPayPerUseCustomerResourcesResponse), nil
	}
}

// ListPayPerUseCustomerResourcesInvoker 查询客户包年/包月资源列表
func (c *BssintlClient) ListPayPerUseCustomerResourcesInvoker(request *model.ListPayPerUseCustomerResourcesRequest) *ListPayPerUseCustomerResourcesInvoker {
	requestDef := GenReqDefForListPayPerUseCustomerResources()
	return &ListPayPerUseCustomerResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPostpaidBillSum 查询伙伴月度消费账单
//
// 功能描述：伙伴可以查询伙伴月度消费账单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListPostpaidBillSum(request *model.ListPostpaidBillSumRequest) (*model.ListPostpaidBillSumResponse, error) {
	requestDef := GenReqDefForListPostpaidBillSum()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPostpaidBillSumResponse), nil
	}
}

// ListPostpaidBillSumInvoker 查询伙伴月度消费账单
func (c *BssintlClient) ListPostpaidBillSumInvoker(request *model.ListPostpaidBillSumRequest) *ListPostpaidBillSumInvoker {
	requestDef := GenReqDefForListPostpaidBillSum()
	return &ListPostpaidBillSumInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRateOnPeriodDetail 查询包年/包月产品价格
//
// 功能描述：客户在自建平台按照条件查询包年/包月产品开通时候的价格
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListRateOnPeriodDetail(request *model.ListRateOnPeriodDetailRequest) (*model.ListRateOnPeriodDetailResponse, error) {
	requestDef := GenReqDefForListRateOnPeriodDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRateOnPeriodDetailResponse), nil
	}
}

// ListRateOnPeriodDetailInvoker 查询包年/包月产品价格
func (c *BssintlClient) ListRateOnPeriodDetailInvoker(request *model.ListRateOnPeriodDetailRequest) *ListRateOnPeriodDetailInvoker {
	requestDef := GenReqDefForListRateOnPeriodDetail()
	return &ListRateOnPeriodDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRenewRateOnPeriod 查询待续订包年包月资源的续订金额
//
// 功能描述：客户在自建平台按照条件查询待续订包年/包月资源续订时候的续订金额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListRenewRateOnPeriod(request *model.ListRenewRateOnPeriodRequest) (*model.ListRenewRateOnPeriodResponse, error) {
	requestDef := GenReqDefForListRenewRateOnPeriod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRenewRateOnPeriodResponse), nil
	}
}

// ListRenewRateOnPeriodInvoker 查询待续订包年包月资源的续订金额
func (c *BssintlClient) ListRenewRateOnPeriodInvoker(request *model.ListRenewRateOnPeriodRequest) *ListRenewRateOnPeriodInvoker {
	requestDef := GenReqDefForListRenewRateOnPeriod()
	return &ListRenewRateOnPeriodInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceTypes 查询资源类型列表
//
// 伙伴在伙伴销售平台查询资源类型的列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListResourceTypes(request *model.ListResourceTypesRequest) (*model.ListResourceTypesResponse, error) {
	requestDef := GenReqDefForListResourceTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceTypesResponse), nil
	}
}

// ListResourceTypesInvoker 查询资源类型列表
func (c *BssintlClient) ListResourceTypesInvoker(request *model.ListResourceTypesRequest) *ListResourceTypesInvoker {
	requestDef := GenReqDefForListResourceTypes()
	return &ListResourceTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServiceResources 根据云服务类型查询资源列表
//
// 功能描述：伙伴在伙伴销售平台根据云服务类型查询关联的资源类型编码和名称，用于查询按需产品的价格或包年/包月产品的价格。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListServiceResources(request *model.ListServiceResourcesRequest) (*model.ListServiceResourcesResponse, error) {
	requestDef := GenReqDefForListServiceResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceResourcesResponse), nil
	}
}

// ListServiceResourcesInvoker 根据云服务类型查询资源列表
func (c *BssintlClient) ListServiceResourcesInvoker(request *model.ListServiceResourcesRequest) *ListServiceResourcesInvoker {
	requestDef := GenReqDefForListServiceResources()
	return &ListServiceResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServiceTypes 查询云服务类型列表
//
// 伙伴在伙伴销售平台查询云服务类型的列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListServiceTypes(request *model.ListServiceTypesRequest) (*model.ListServiceTypesResponse, error) {
	requestDef := GenReqDefForListServiceTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceTypesResponse), nil
	}
}

// ListServiceTypesInvoker 查询云服务类型列表
func (c *BssintlClient) ListServiceTypesInvoker(request *model.ListServiceTypesRequest) *ListServiceTypesInvoker {
	requestDef := GenReqDefForListServiceTypes()
	return &ListServiceTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubCustomerBudget 批量查询客户预算
//
// 功能描述：批量查询客户预算
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListSubCustomerBudget(request *model.ListSubCustomerBudgetRequest) (*model.ListSubCustomerBudgetResponse, error) {
	requestDef := GenReqDefForListSubCustomerBudget()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubCustomerBudgetResponse), nil
	}
}

// ListSubCustomerBudgetInvoker 批量查询客户预算
func (c *BssintlClient) ListSubCustomerBudgetInvoker(request *model.ListSubCustomerBudgetRequest) *ListSubCustomerBudgetInvoker {
	requestDef := GenReqDefForListSubCustomerBudget()
	return &ListSubCustomerBudgetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubCustomerCoupons 查询优惠券列表
//
// 功能描述：伙伴/客户可以查询自身的优惠券信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListSubCustomerCoupons(request *model.ListSubCustomerCouponsRequest) (*model.ListSubCustomerCouponsResponse, error) {
	requestDef := GenReqDefForListSubCustomerCoupons()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubCustomerCouponsResponse), nil
	}
}

// ListSubCustomerCouponsInvoker 查询优惠券列表
func (c *BssintlClient) ListSubCustomerCouponsInvoker(request *model.ListSubCustomerCouponsRequest) *ListSubCustomerCouponsInvoker {
	requestDef := GenReqDefForListSubCustomerCoupons()
	return &ListSubCustomerCouponsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubCustomers 查询客户列表
//
// 功能描述：伙伴可以查询合作伙伴的客户信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListSubCustomers(request *model.ListSubCustomersRequest) (*model.ListSubCustomersResponse, error) {
	requestDef := GenReqDefForListSubCustomers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubCustomersResponse), nil
	}
}

// ListSubCustomersInvoker 查询客户列表
func (c *BssintlClient) ListSubCustomersInvoker(request *model.ListSubCustomersRequest) *ListSubCustomersInvoker {
	requestDef := GenReqDefForListSubCustomers()
	return &ListSubCustomersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsageTypes 查询使用量类型列表
//
// 功能描述：伙伴在伙伴销售平台查询资源的使用量类型列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ListUsageTypes(request *model.ListUsageTypesRequest) (*model.ListUsageTypesResponse, error) {
	requestDef := GenReqDefForListUsageTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsageTypesResponse), nil
	}
}

// ListUsageTypesInvoker 查询使用量类型列表
func (c *BssintlClient) ListUsageTypesInvoker(request *model.ListUsageTypesRequest) *ListUsageTypesInvoker {
	requestDef := GenReqDefForListUsageTypes()
	return &ListUsageTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PayOrders 支付包年/包月产品订单
//
// 客户可以对待支付状态的包年/包月产品订单进行支付
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) PayOrders(request *model.PayOrdersRequest) (*model.PayOrdersResponse, error) {
	requestDef := GenReqDefForPayOrders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PayOrdersResponse), nil
	}
}

// PayOrdersInvoker 支付包年/包月产品订单
func (c *BssintlClient) PayOrdersInvoker(request *model.PayOrdersRequest) *PayOrdersInvoker {
	requestDef := GenReqDefForPayOrders()
	return &PayOrdersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RenewalResources 续订包年/包月资源
//
// 功能描述：客户的包年包/月资源即将到期时，可进行包年/包月资源的续订
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) RenewalResources(request *model.RenewalResourcesRequest) (*model.RenewalResourcesResponse, error) {
	requestDef := GenReqDefForRenewalResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RenewalResourcesResponse), nil
	}
}

// RenewalResourcesInvoker 续订包年/包月资源
func (c *BssintlClient) RenewalResourcesInvoker(request *model.RenewalResourcesRequest) *RenewalResourcesInvoker {
	requestDef := GenReqDefForRenewalResources()
	return &RenewalResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendVerificationMessageCode 发送验证码
//
// 功能描述：客户注册时，如果填写了邮箱，可以向对应的邮箱发送注册验证码，校验信息的正确性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) SendVerificationMessageCode(request *model.SendVerificationMessageCodeRequest) (*model.SendVerificationMessageCodeResponse, error) {
	requestDef := GenReqDefForSendVerificationMessageCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendVerificationMessageCodeResponse), nil
	}
}

// SendVerificationMessageCodeInvoker 发送验证码
func (c *BssintlClient) SendVerificationMessageCodeInvoker(request *model.SendVerificationMessageCodeRequest) *SendVerificationMessageCodeInvoker {
	requestDef := GenReqDefForSendVerificationMessageCode()
	return &SendVerificationMessageCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCustomerAccountBalances 查询账户余额
//
// 功能描述：客户可以查询自身的账户余额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ShowCustomerAccountBalances(request *model.ShowCustomerAccountBalancesRequest) (*model.ShowCustomerAccountBalancesResponse, error) {
	requestDef := GenReqDefForShowCustomerAccountBalances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCustomerAccountBalancesResponse), nil
	}
}

// ShowCustomerAccountBalancesInvoker 查询账户余额
func (c *BssintlClient) ShowCustomerAccountBalancesInvoker(request *model.ShowCustomerAccountBalancesRequest) *ShowCustomerAccountBalancesInvoker {
	requestDef := GenReqDefForShowCustomerAccountBalances()
	return &ShowCustomerAccountBalancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCustomerOrderDetails 查询订单详情
//
// 功能描述：客户可以查看订单详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ShowCustomerOrderDetails(request *model.ShowCustomerOrderDetailsRequest) (*model.ShowCustomerOrderDetailsResponse, error) {
	requestDef := GenReqDefForShowCustomerOrderDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCustomerOrderDetailsResponse), nil
	}
}

// ShowCustomerOrderDetailsInvoker 查询订单详情
func (c *BssintlClient) ShowCustomerOrderDetailsInvoker(request *model.ShowCustomerOrderDetailsRequest) *ShowCustomerOrderDetailsInvoker {
	requestDef := GenReqDefForShowCustomerOrderDetails()
	return &ShowCustomerOrderDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPartnerConsumptionQuota 查询消费配额
//
// 功能描述：合作伙伴可以查询消费配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ShowPartnerConsumptionQuota(request *model.ShowPartnerConsumptionQuotaRequest) (*model.ShowPartnerConsumptionQuotaResponse, error) {
	requestDef := GenReqDefForShowPartnerConsumptionQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPartnerConsumptionQuotaResponse), nil
	}
}

// ShowPartnerConsumptionQuotaInvoker 查询消费配额
func (c *BssintlClient) ShowPartnerConsumptionQuotaInvoker(request *model.ShowPartnerConsumptionQuotaRequest) *ShowPartnerConsumptionQuotaInvoker {
	requestDef := GenReqDefForShowPartnerConsumptionQuota()
	return &ShowPartnerConsumptionQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRealnameAuthenticationReviewResult 查询实名认证审核结果
//
// 功能描述：如果实名认证申请或实名认证变更申请的响应中，显示需要人工审核，使用该接口查询审核结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ShowRealnameAuthenticationReviewResult(request *model.ShowRealnameAuthenticationReviewResultRequest) (*model.ShowRealnameAuthenticationReviewResultResponse, error) {
	requestDef := GenReqDefForShowRealnameAuthenticationReviewResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRealnameAuthenticationReviewResultResponse), nil
	}
}

// ShowRealnameAuthenticationReviewResultInvoker 查询实名认证审核结果
func (c *BssintlClient) ShowRealnameAuthenticationReviewResultInvoker(request *model.ShowRealnameAuthenticationReviewResultRequest) *ShowRealnameAuthenticationReviewResultInvoker {
	requestDef := GenReqDefForShowRealnameAuthenticationReviewResult()
	return &ShowRealnameAuthenticationReviewResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRefundOrderDetails 查询退款订单的金额详情
//
// 功能描述：客户在伙伴销售平台查询某次退订订单或者降配订单的退款金额来自哪些资源和对应订单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ShowRefundOrderDetails(request *model.ShowRefundOrderDetailsRequest) (*model.ShowRefundOrderDetailsResponse, error) {
	requestDef := GenReqDefForShowRefundOrderDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRefundOrderDetailsResponse), nil
	}
}

// ShowRefundOrderDetailsInvoker 查询退款订单的金额详情
func (c *BssintlClient) ShowRefundOrderDetailsInvoker(request *model.ShowRefundOrderDetailsRequest) *ShowRefundOrderDetailsInvoker {
	requestDef := GenReqDefForShowRefundOrderDetails()
	return &ShowRefundOrderDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSubCustomerBudget 查询客户预算
//
// 功能描述：查询客户预算
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) ShowSubCustomerBudget(request *model.ShowSubCustomerBudgetRequest) (*model.ShowSubCustomerBudgetResponse, error) {
	requestDef := GenReqDefForShowSubCustomerBudget()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSubCustomerBudgetResponse), nil
	}
}

// ShowSubCustomerBudgetInvoker 查询客户预算
func (c *BssintlClient) ShowSubCustomerBudgetInvoker(request *model.ShowSubCustomerBudgetRequest) *ShowSubCustomerBudgetInvoker {
	requestDef := GenReqDefForShowSubCustomerBudget()
	return &ShowSubCustomerBudgetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnfreezeSubCustomers 解冻客户账号
//
// 功能描述：解冻伙伴子客户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) UnfreezeSubCustomers(request *model.UnfreezeSubCustomersRequest) (*model.UnfreezeSubCustomersResponse, error) {
	requestDef := GenReqDefForUnfreezeSubCustomers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnfreezeSubCustomersResponse), nil
	}
}

// UnfreezeSubCustomersInvoker 解冻客户账号
func (c *BssintlClient) UnfreezeSubCustomersInvoker(request *model.UnfreezeSubCustomersRequest) *UnfreezeSubCustomersInvoker {
	requestDef := GenReqDefForUnfreezeSubCustomers()
	return &UnfreezeSubCustomersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePeriodToOnDemand 设置或者取消包年/包月资源到期转按需
//
// 功能描述：客户可以设置包年/包月资源到期后转为按需资源计费。包年/包月计费模式到期后，按需的计费模式即生效
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) UpdatePeriodToOnDemand(request *model.UpdatePeriodToOnDemandRequest) (*model.UpdatePeriodToOnDemandResponse, error) {
	requestDef := GenReqDefForUpdatePeriodToOnDemand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePeriodToOnDemandResponse), nil
	}
}

// UpdatePeriodToOnDemandInvoker 设置或者取消包年/包月资源到期转按需
func (c *BssintlClient) UpdatePeriodToOnDemandInvoker(request *model.UpdatePeriodToOnDemandRequest) *UpdatePeriodToOnDemandInvoker {
	requestDef := GenReqDefForUpdatePeriodToOnDemand()
	return &UpdatePeriodToOnDemandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSubCustomerBudget 设置客户预算
//
// 功能描述：设置客户预算
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *BssintlClient) UpdateSubCustomerBudget(request *model.UpdateSubCustomerBudgetRequest) (*model.UpdateSubCustomerBudgetResponse, error) {
	requestDef := GenReqDefForUpdateSubCustomerBudget()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubCustomerBudgetResponse), nil
	}
}

// UpdateSubCustomerBudgetInvoker 设置客户预算
func (c *BssintlClient) UpdateSubCustomerBudgetInvoker(request *model.UpdateSubCustomerBudgetRequest) *UpdateSubCustomerBudgetInvoker {
	requestDef := GenReqDefForUpdateSubCustomerBudget()
	return &UpdateSubCustomerBudgetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
