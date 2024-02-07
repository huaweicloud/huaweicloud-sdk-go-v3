package v3

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/geip/v3/model"
)

type GeipClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewGeipClient(hcClient *httpclient.HcHttpClient) *GeipClient {
	return &GeipClient{HcClient: hcClient}
}

func GeipClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// AddInternetBandwidthTags 添加全域公网带宽标签
//
// 添加全域公网带宽标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) AddInternetBandwidthTags(request *model.AddInternetBandwidthTagsRequest) (*model.AddInternetBandwidthTagsResponse, error) {
	requestDef := GenReqDefForAddInternetBandwidthTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddInternetBandwidthTagsResponse), nil
	}
}

// AddInternetBandwidthTagsInvoker 添加全域公网带宽标签
func (c *GeipClient) AddInternetBandwidthTagsInvoker(request *model.AddInternetBandwidthTagsRequest) *AddInternetBandwidthTagsInvoker {
	requestDef := GenReqDefForAddInternetBandwidthTags()
	return &AddInternetBandwidthTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateInternetBandwidth 批量创建全域公网带宽
//
// 批量创建全域公网带宽
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) BatchCreateInternetBandwidth(request *model.BatchCreateInternetBandwidthRequest) (*model.BatchCreateInternetBandwidthResponse, error) {
	requestDef := GenReqDefForBatchCreateInternetBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateInternetBandwidthResponse), nil
	}
}

// BatchCreateInternetBandwidthInvoker 批量创建全域公网带宽
func (c *GeipClient) BatchCreateInternetBandwidthInvoker(request *model.BatchCreateInternetBandwidthRequest) *BatchCreateInternetBandwidthInvoker {
	requestDef := GenReqDefForBatchCreateInternetBandwidth()
	return &BatchCreateInternetBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateInternetBandwidthTags 批量添加全域公网带宽标签
//
// 批量添加全域公网带宽标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) BatchCreateInternetBandwidthTags(request *model.BatchCreateInternetBandwidthTagsRequest) (*model.BatchCreateInternetBandwidthTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateInternetBandwidthTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateInternetBandwidthTagsResponse), nil
	}
}

// BatchCreateInternetBandwidthTagsInvoker 批量添加全域公网带宽标签
func (c *GeipClient) BatchCreateInternetBandwidthTagsInvoker(request *model.BatchCreateInternetBandwidthTagsRequest) *BatchCreateInternetBandwidthTagsInvoker {
	requestDef := GenReqDefForBatchCreateInternetBandwidthTags()
	return &BatchCreateInternetBandwidthTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteInternetBandwidthTags 批量删除全域公网带宽标签
//
// 批量删除全域公网带宽标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) BatchDeleteInternetBandwidthTags(request *model.BatchDeleteInternetBandwidthTagsRequest) (*model.BatchDeleteInternetBandwidthTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteInternetBandwidthTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteInternetBandwidthTagsResponse), nil
	}
}

// BatchDeleteInternetBandwidthTagsInvoker 批量删除全域公网带宽标签
func (c *GeipClient) BatchDeleteInternetBandwidthTagsInvoker(request *model.BatchDeleteInternetBandwidthTagsRequest) *BatchDeleteInternetBandwidthTagsInvoker {
	requestDef := GenReqDefForBatchDeleteInternetBandwidthTags()
	return &BatchDeleteInternetBandwidthTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountInternetBandwidth 查询全域公网带宽个数
//
// 查询全域公网带宽个数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) CountInternetBandwidth(request *model.CountInternetBandwidthRequest) (*model.CountInternetBandwidthResponse, error) {
	requestDef := GenReqDefForCountInternetBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountInternetBandwidthResponse), nil
	}
}

// CountInternetBandwidthInvoker 查询全域公网带宽个数
func (c *GeipClient) CountInternetBandwidthInvoker(request *model.CountInternetBandwidthRequest) *CountInternetBandwidthInvoker {
	requestDef := GenReqDefForCountInternetBandwidth()
	return &CountInternetBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInternetBandwidth 创建全域公网带宽
//
// 创建全域公网带宽
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) CreateInternetBandwidth(request *model.CreateInternetBandwidthRequest) (*model.CreateInternetBandwidthResponse, error) {
	requestDef := GenReqDefForCreateInternetBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInternetBandwidthResponse), nil
	}
}

// CreateInternetBandwidthInvoker 创建全域公网带宽
func (c *GeipClient) CreateInternetBandwidthInvoker(request *model.CreateInternetBandwidthRequest) *CreateInternetBandwidthInvoker {
	requestDef := GenReqDefForCreateInternetBandwidth()
	return &CreateInternetBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateUserDisclaimer 创建租户签署免责条款
//
// 创建租户签署免责条款
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) CreateUserDisclaimer(request *model.CreateUserDisclaimerRequest) (*model.CreateUserDisclaimerResponse, error) {
	requestDef := GenReqDefForCreateUserDisclaimer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUserDisclaimerResponse), nil
	}
}

// CreateUserDisclaimerInvoker 创建租户签署免责条款
func (c *GeipClient) CreateUserDisclaimerInvoker(request *model.CreateUserDisclaimerRequest) *CreateUserDisclaimerInvoker {
	requestDef := GenReqDefForCreateUserDisclaimer()
	return &CreateUserDisclaimerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInternetBandwidth 删除全域公网带宽
//
// 删除全域公网带宽
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) DeleteInternetBandwidth(request *model.DeleteInternetBandwidthRequest) (*model.DeleteInternetBandwidthResponse, error) {
	requestDef := GenReqDefForDeleteInternetBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInternetBandwidthResponse), nil
	}
}

// DeleteInternetBandwidthInvoker 删除全域公网带宽
func (c *GeipClient) DeleteInternetBandwidthInvoker(request *model.DeleteInternetBandwidthRequest) *DeleteInternetBandwidthInvoker {
	requestDef := GenReqDefForDeleteInternetBandwidth()
	return &DeleteInternetBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInternetBandwidthTag 删除全域公网带宽标签
//
// 删除全域公网带宽标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) DeleteInternetBandwidthTag(request *model.DeleteInternetBandwidthTagRequest) (*model.DeleteInternetBandwidthTagResponse, error) {
	requestDef := GenReqDefForDeleteInternetBandwidthTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInternetBandwidthTagResponse), nil
	}
}

// DeleteInternetBandwidthTagInvoker 删除全域公网带宽标签
func (c *GeipClient) DeleteInternetBandwidthTagInvoker(request *model.DeleteInternetBandwidthTagRequest) *DeleteInternetBandwidthTagInvoker {
	requestDef := GenReqDefForDeleteInternetBandwidthTag()
	return &DeleteInternetBandwidthTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteUserDisclaimer 删除租户撤销免责条款
//
// 删除租户撤销免责条款
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) DeleteUserDisclaimer(request *model.DeleteUserDisclaimerRequest) (*model.DeleteUserDisclaimerResponse, error) {
	requestDef := GenReqDefForDeleteUserDisclaimer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserDisclaimerResponse), nil
	}
}

// DeleteUserDisclaimerInvoker 删除租户撤销免责条款
func (c *GeipClient) DeleteUserDisclaimerInvoker(request *model.DeleteUserDisclaimerRequest) *DeleteUserDisclaimerInvoker {
	requestDef := GenReqDefForDeleteUserDisclaimer()
	return &DeleteUserDisclaimerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccessSites 查询接入点列表
//
// 查询接入点列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListAccessSites(request *model.ListAccessSitesRequest) (*model.ListAccessSitesResponse, error) {
	requestDef := GenReqDefForListAccessSites()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccessSitesResponse), nil
	}
}

// ListAccessSitesInvoker 查询接入点列表
func (c *GeipClient) ListAccessSitesInvoker(request *model.ListAccessSitesRequest) *ListAccessSitesInvoker {
	requestDef := GenReqDefForListAccessSites()
	return &ListAccessSitesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGeipResourceQuotas 查询租户全域弹性公网IP配额
//
// 查询租户全域弹性公网IP配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListGeipResourceQuotas(request *model.ListGeipResourceQuotasRequest) (*model.ListGeipResourceQuotasResponse, error) {
	requestDef := GenReqDefForListGeipResourceQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGeipResourceQuotasResponse), nil
	}
}

// ListGeipResourceQuotasInvoker 查询租户全域弹性公网IP配额
func (c *GeipClient) ListGeipResourceQuotasInvoker(request *model.ListGeipResourceQuotasRequest) *ListGeipResourceQuotasInvoker {
	requestDef := GenReqDefForListGeipResourceQuotas()
	return &ListGeipResourceQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInternetBandwidthCountFilterTags 查询资源实例列表数目
//
// 查询资源实例列表数目
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListInternetBandwidthCountFilterTags(request *model.ListInternetBandwidthCountFilterTagsRequest) (*model.ListInternetBandwidthCountFilterTagsResponse, error) {
	requestDef := GenReqDefForListInternetBandwidthCountFilterTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInternetBandwidthCountFilterTagsResponse), nil
	}
}

// ListInternetBandwidthCountFilterTagsInvoker 查询资源实例列表数目
func (c *GeipClient) ListInternetBandwidthCountFilterTagsInvoker(request *model.ListInternetBandwidthCountFilterTagsRequest) *ListInternetBandwidthCountFilterTagsInvoker {
	requestDef := GenReqDefForListInternetBandwidthCountFilterTags()
	return &ListInternetBandwidthCountFilterTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInternetBandwidthDomainTags 查询全域公网带宽项目标签
//
// 查询全域公网带宽项目标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListInternetBandwidthDomainTags(request *model.ListInternetBandwidthDomainTagsRequest) (*model.ListInternetBandwidthDomainTagsResponse, error) {
	requestDef := GenReqDefForListInternetBandwidthDomainTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInternetBandwidthDomainTagsResponse), nil
	}
}

// ListInternetBandwidthDomainTagsInvoker 查询全域公网带宽项目标签
func (c *GeipClient) ListInternetBandwidthDomainTagsInvoker(request *model.ListInternetBandwidthDomainTagsRequest) *ListInternetBandwidthDomainTagsInvoker {
	requestDef := GenReqDefForListInternetBandwidthDomainTags()
	return &ListInternetBandwidthDomainTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInternetBandwidthFilterTags 查询资源实例列表
//
// 查询资源实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListInternetBandwidthFilterTags(request *model.ListInternetBandwidthFilterTagsRequest) (*model.ListInternetBandwidthFilterTagsResponse, error) {
	requestDef := GenReqDefForListInternetBandwidthFilterTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInternetBandwidthFilterTagsResponse), nil
	}
}

// ListInternetBandwidthFilterTagsInvoker 查询资源实例列表
func (c *GeipClient) ListInternetBandwidthFilterTagsInvoker(request *model.ListInternetBandwidthFilterTagsRequest) *ListInternetBandwidthFilterTagsInvoker {
	requestDef := GenReqDefForListInternetBandwidthFilterTags()
	return &ListInternetBandwidthFilterTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInternetBandwidthLimits 全域公网带宽限制列表
//
// 查询全域公网带宽限制列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListInternetBandwidthLimits(request *model.ListInternetBandwidthLimitsRequest) (*model.ListInternetBandwidthLimitsResponse, error) {
	requestDef := GenReqDefForListInternetBandwidthLimits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInternetBandwidthLimitsResponse), nil
	}
}

// ListInternetBandwidthLimitsInvoker 全域公网带宽限制列表
func (c *GeipClient) ListInternetBandwidthLimitsInvoker(request *model.ListInternetBandwidthLimitsRequest) *ListInternetBandwidthLimitsInvoker {
	requestDef := GenReqDefForListInternetBandwidthLimits()
	return &ListInternetBandwidthLimitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInternetBandwidths 查询全域公网带宽列表
//
// 查询全域公网带宽列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListInternetBandwidths(request *model.ListInternetBandwidthsRequest) (*model.ListInternetBandwidthsResponse, error) {
	requestDef := GenReqDefForListInternetBandwidths()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInternetBandwidthsResponse), nil
	}
}

// ListInternetBandwidthsInvoker 查询全域公网带宽列表
func (c *GeipClient) ListInternetBandwidthsInvoker(request *model.ListInternetBandwidthsRequest) *ListInternetBandwidthsInvoker {
	requestDef := GenReqDefForListInternetBandwidths()
	return &ListInternetBandwidthsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSupportMasks 查询全域弹性公网IP段支持的掩码列表
//
// 查询全域弹性公网IP段支持的掩码列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListSupportMasks(request *model.ListSupportMasksRequest) (*model.ListSupportMasksResponse, error) {
	requestDef := GenReqDefForListSupportMasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSupportMasksResponse), nil
	}
}

// ListSupportMasksInvoker 查询全域弹性公网IP段支持的掩码列表
func (c *GeipClient) ListSupportMasksInvoker(request *model.ListSupportMasksRequest) *ListSupportMasksInvoker {
	requestDef := GenReqDefForListSupportMasks()
	return &ListSupportMasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInternetBandwidth 查询全域公网带宽详情
//
// 查询全域公网带宽详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ShowInternetBandwidth(request *model.ShowInternetBandwidthRequest) (*model.ShowInternetBandwidthResponse, error) {
	requestDef := GenReqDefForShowInternetBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInternetBandwidthResponse), nil
	}
}

// ShowInternetBandwidthInvoker 查询全域公网带宽详情
func (c *GeipClient) ShowInternetBandwidthInvoker(request *model.ShowInternetBandwidthRequest) *ShowInternetBandwidthInvoker {
	requestDef := GenReqDefForShowInternetBandwidth()
	return &ShowInternetBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInternetBandwidthTags 查询全域公网带宽标签
//
// 查询全域公网带宽标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ShowInternetBandwidthTags(request *model.ShowInternetBandwidthTagsRequest) (*model.ShowInternetBandwidthTagsResponse, error) {
	requestDef := GenReqDefForShowInternetBandwidthTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInternetBandwidthTagsResponse), nil
	}
}

// ShowInternetBandwidthTagsInvoker 查询全域公网带宽标签
func (c *GeipClient) ShowInternetBandwidthTagsInvoker(request *model.ShowInternetBandwidthTagsRequest) *ShowInternetBandwidthTagsInvoker {
	requestDef := GenReqDefForShowInternetBandwidthTags()
	return &ShowInternetBandwidthTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserDisclaimer 查询租户签署免责条款详情
//
// 查询租户签署免责条款详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ShowUserDisclaimer(request *model.ShowUserDisclaimerRequest) (*model.ShowUserDisclaimerResponse, error) {
	requestDef := GenReqDefForShowUserDisclaimer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserDisclaimerResponse), nil
	}
}

// ShowUserDisclaimerInvoker 查询租户签署免责条款详情
func (c *GeipClient) ShowUserDisclaimerInvoker(request *model.ShowUserDisclaimerRequest) *ShowUserDisclaimerInvoker {
	requestDef := GenReqDefForShowUserDisclaimer()
	return &ShowUserDisclaimerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInternetBandwidth 更新全域公网带宽
//
// 更新全域公网带宽
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) UpdateInternetBandwidth(request *model.UpdateInternetBandwidthRequest) (*model.UpdateInternetBandwidthResponse, error) {
	requestDef := GenReqDefForUpdateInternetBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInternetBandwidthResponse), nil
	}
}

// UpdateInternetBandwidthInvoker 更新全域公网带宽
func (c *GeipClient) UpdateInternetBandwidthInvoker(request *model.UpdateInternetBandwidthRequest) *UpdateInternetBandwidthInvoker {
	requestDef := GenReqDefForUpdateInternetBandwidth()
	return &UpdateInternetBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddGeipSegmentTags 添加全域弹性公网IP段标签
//
// 添加全域弹性公网IP段的标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) AddGeipSegmentTags(request *model.AddGeipSegmentTagsRequest) (*model.AddGeipSegmentTagsResponse, error) {
	requestDef := GenReqDefForAddGeipSegmentTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddGeipSegmentTagsResponse), nil
	}
}

// AddGeipSegmentTagsInvoker 添加全域弹性公网IP段标签
func (c *GeipClient) AddGeipSegmentTagsInvoker(request *model.AddGeipSegmentTagsRequest) *AddGeipSegmentTagsInvoker {
	requestDef := GenReqDefForAddGeipSegmentTags()
	return &AddGeipSegmentTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddGlobalEipTags 添加全域弹性公网IP标签
//
// 添加全域弹性公网IP的标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) AddGlobalEipTags(request *model.AddGlobalEipTagsRequest) (*model.AddGlobalEipTagsResponse, error) {
	requestDef := GenReqDefForAddGlobalEipTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddGlobalEipTagsResponse), nil
	}
}

// AddGlobalEipTagsInvoker 添加全域弹性公网IP标签
func (c *GeipClient) AddGlobalEipTagsInvoker(request *model.AddGlobalEipTagsRequest) *AddGlobalEipTagsInvoker {
	requestDef := GenReqDefForAddGlobalEipTags()
	return &AddGlobalEipTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateGeipSegmentInstance 全域弹性公网IP段绑定后端实例
//
// 全域弹性公网IP段绑定后端实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) AssociateGeipSegmentInstance(request *model.AssociateGeipSegmentInstanceRequest) (*model.AssociateGeipSegmentInstanceResponse, error) {
	requestDef := GenReqDefForAssociateGeipSegmentInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateGeipSegmentInstanceResponse), nil
	}
}

// AssociateGeipSegmentInstanceInvoker 全域弹性公网IP段绑定后端实例
func (c *GeipClient) AssociateGeipSegmentInstanceInvoker(request *model.AssociateGeipSegmentInstanceRequest) *AssociateGeipSegmentInstanceInvoker {
	requestDef := GenReqDefForAssociateGeipSegmentInstance()
	return &AssociateGeipSegmentInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateInstance 绑定后端实例
//
// 绑定后端实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) AssociateInstance(request *model.AssociateInstanceRequest) (*model.AssociateInstanceResponse, error) {
	requestDef := GenReqDefForAssociateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateInstanceResponse), nil
	}
}

// AssociateInstanceInvoker 绑定后端实例
func (c *GeipClient) AssociateInstanceInvoker(request *model.AssociateInstanceRequest) *AssociateInstanceInvoker {
	requestDef := GenReqDefForAssociateInstance()
	return &AssociateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachInternetBandwidth 绑定全域公网带宽
//
// 绑定全域公网带宽
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) AttachInternetBandwidth(request *model.AttachInternetBandwidthRequest) (*model.AttachInternetBandwidthResponse, error) {
	requestDef := GenReqDefForAttachInternetBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachInternetBandwidthResponse), nil
	}
}

// AttachInternetBandwidthInvoker 绑定全域公网带宽
func (c *GeipClient) AttachInternetBandwidthInvoker(request *model.AttachInternetBandwidthRequest) *AttachInternetBandwidthInvoker {
	requestDef := GenReqDefForAttachInternetBandwidth()
	return &AttachInternetBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAttachGeipSegmentInternetBandwidth 全域弹性公网IP段批量绑定全域公网带宽
//
// 全域弹性公网IP段批量绑定全域公网带宽
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) BatchAttachGeipSegmentInternetBandwidth(request *model.BatchAttachGeipSegmentInternetBandwidthRequest) (*model.BatchAttachGeipSegmentInternetBandwidthResponse, error) {
	requestDef := GenReqDefForBatchAttachGeipSegmentInternetBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAttachGeipSegmentInternetBandwidthResponse), nil
	}
}

// BatchAttachGeipSegmentInternetBandwidthInvoker 全域弹性公网IP段批量绑定全域公网带宽
func (c *GeipClient) BatchAttachGeipSegmentInternetBandwidthInvoker(request *model.BatchAttachGeipSegmentInternetBandwidthRequest) *BatchAttachGeipSegmentInternetBandwidthInvoker {
	requestDef := GenReqDefForBatchAttachGeipSegmentInternetBandwidth()
	return &BatchAttachGeipSegmentInternetBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAttachInternetBandwidth 批量绑定全域公网带宽
//
// 批量绑定全域公网带宽
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) BatchAttachInternetBandwidth(request *model.BatchAttachInternetBandwidthRequest) (*model.BatchAttachInternetBandwidthResponse, error) {
	requestDef := GenReqDefForBatchAttachInternetBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAttachInternetBandwidthResponse), nil
	}
}

// BatchAttachInternetBandwidthInvoker 批量绑定全域公网带宽
func (c *GeipClient) BatchAttachInternetBandwidthInvoker(request *model.BatchAttachInternetBandwidthRequest) *BatchAttachInternetBandwidthInvoker {
	requestDef := GenReqDefForBatchAttachInternetBandwidth()
	return &BatchAttachInternetBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateGeipSegmentTags 批量添加全域弹性公网IP段标签
//
// 批量添加全域弹性公网IP段的标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) BatchCreateGeipSegmentTags(request *model.BatchCreateGeipSegmentTagsRequest) (*model.BatchCreateGeipSegmentTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateGeipSegmentTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateGeipSegmentTagsResponse), nil
	}
}

// BatchCreateGeipSegmentTagsInvoker 批量添加全域弹性公网IP段标签
func (c *GeipClient) BatchCreateGeipSegmentTagsInvoker(request *model.BatchCreateGeipSegmentTagsRequest) *BatchCreateGeipSegmentTagsInvoker {
	requestDef := GenReqDefForBatchCreateGeipSegmentTags()
	return &BatchCreateGeipSegmentTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateGlobalEip 批量创建全域弹性公网IP
//
// 批量创建全域弹性公网IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) BatchCreateGlobalEip(request *model.BatchCreateGlobalEipRequest) (*model.BatchCreateGlobalEipResponse, error) {
	requestDef := GenReqDefForBatchCreateGlobalEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateGlobalEipResponse), nil
	}
}

// BatchCreateGlobalEipInvoker 批量创建全域弹性公网IP
func (c *GeipClient) BatchCreateGlobalEipInvoker(request *model.BatchCreateGlobalEipRequest) *BatchCreateGlobalEipInvoker {
	requestDef := GenReqDefForBatchCreateGlobalEip()
	return &BatchCreateGlobalEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateGlobalEipTags 批量添加全域弹性公网IP标签
//
// 批量添加全域弹性公网IP的标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) BatchCreateGlobalEipTags(request *model.BatchCreateGlobalEipTagsRequest) (*model.BatchCreateGlobalEipTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateGlobalEipTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateGlobalEipTagsResponse), nil
	}
}

// BatchCreateGlobalEipTagsInvoker 批量添加全域弹性公网IP标签
func (c *GeipClient) BatchCreateGlobalEipTagsInvoker(request *model.BatchCreateGlobalEipTagsRequest) *BatchCreateGlobalEipTagsInvoker {
	requestDef := GenReqDefForBatchCreateGlobalEipTags()
	return &BatchCreateGlobalEipTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteGeipSegmentTags 批量删除全域弹性公网IP段标签
//
// 批量删除全域弹性公网IP段的标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) BatchDeleteGeipSegmentTags(request *model.BatchDeleteGeipSegmentTagsRequest) (*model.BatchDeleteGeipSegmentTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteGeipSegmentTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteGeipSegmentTagsResponse), nil
	}
}

// BatchDeleteGeipSegmentTagsInvoker 批量删除全域弹性公网IP段标签
func (c *GeipClient) BatchDeleteGeipSegmentTagsInvoker(request *model.BatchDeleteGeipSegmentTagsRequest) *BatchDeleteGeipSegmentTagsInvoker {
	requestDef := GenReqDefForBatchDeleteGeipSegmentTags()
	return &BatchDeleteGeipSegmentTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteGlobalEipTags 批量删除全域弹性公网IP标签
//
// 批量删除全域弹性公网IP的标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) BatchDeleteGlobalEipTags(request *model.BatchDeleteGlobalEipTagsRequest) (*model.BatchDeleteGlobalEipTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteGlobalEipTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteGlobalEipTagsResponse), nil
	}
}

// BatchDeleteGlobalEipTagsInvoker 批量删除全域弹性公网IP标签
func (c *GeipClient) BatchDeleteGlobalEipTagsInvoker(request *model.BatchDeleteGlobalEipTagsRequest) *BatchDeleteGlobalEipTagsInvoker {
	requestDef := GenReqDefForBatchDeleteGlobalEipTags()
	return &BatchDeleteGlobalEipTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDetachGeipSegmentInternetBandwidth 全域弹性公网IP段批量解绑全域公网带宽
//
// 全域弹性公网IP段批量解绑全域公网带宽
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) BatchDetachGeipSegmentInternetBandwidth(request *model.BatchDetachGeipSegmentInternetBandwidthRequest) (*model.BatchDetachGeipSegmentInternetBandwidthResponse, error) {
	requestDef := GenReqDefForBatchDetachGeipSegmentInternetBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDetachGeipSegmentInternetBandwidthResponse), nil
	}
}

// BatchDetachGeipSegmentInternetBandwidthInvoker 全域弹性公网IP段批量解绑全域公网带宽
func (c *GeipClient) BatchDetachGeipSegmentInternetBandwidthInvoker(request *model.BatchDetachGeipSegmentInternetBandwidthRequest) *BatchDetachGeipSegmentInternetBandwidthInvoker {
	requestDef := GenReqDefForBatchDetachGeipSegmentInternetBandwidth()
	return &BatchDetachGeipSegmentInternetBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDetachInternetBandwidth 批量解绑全域公网带宽
//
// 批量解绑全域公网带宽
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) BatchDetachInternetBandwidth(request *model.BatchDetachInternetBandwidthRequest) (*model.BatchDetachInternetBandwidthResponse, error) {
	requestDef := GenReqDefForBatchDetachInternetBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDetachInternetBandwidthResponse), nil
	}
}

// BatchDetachInternetBandwidthInvoker 批量解绑全域公网带宽
func (c *GeipClient) BatchDetachInternetBandwidthInvoker(request *model.BatchDetachInternetBandwidthRequest) *BatchDetachInternetBandwidthInvoker {
	requestDef := GenReqDefForBatchDetachInternetBandwidth()
	return &BatchDetachInternetBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountGlobalEipSegment 查询全域弹性公网IP段个数
//
// 查询全域弹性公网IP段个数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) CountGlobalEipSegment(request *model.CountGlobalEipSegmentRequest) (*model.CountGlobalEipSegmentResponse, error) {
	requestDef := GenReqDefForCountGlobalEipSegment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountGlobalEipSegmentResponse), nil
	}
}

// CountGlobalEipSegmentInvoker 查询全域弹性公网IP段个数
func (c *GeipClient) CountGlobalEipSegmentInvoker(request *model.CountGlobalEipSegmentRequest) *CountGlobalEipSegmentInvoker {
	requestDef := GenReqDefForCountGlobalEipSegment()
	return &CountGlobalEipSegmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountGlobalEips 查询全域弹性公网IP个数
//
// 查询全域弹性公网IP个数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) CountGlobalEips(request *model.CountGlobalEipsRequest) (*model.CountGlobalEipsResponse, error) {
	requestDef := GenReqDefForCountGlobalEips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountGlobalEipsResponse), nil
	}
}

// CountGlobalEipsInvoker 查询全域弹性公网IP个数
func (c *GeipClient) CountGlobalEipsInvoker(request *model.CountGlobalEipsRequest) *CountGlobalEipsInvoker {
	requestDef := GenReqDefForCountGlobalEips()
	return &CountGlobalEipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGlobalEip 创建全域弹性公网IP
//
// 创建全域弹性公网IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) CreateGlobalEip(request *model.CreateGlobalEipRequest) (*model.CreateGlobalEipResponse, error) {
	requestDef := GenReqDefForCreateGlobalEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGlobalEipResponse), nil
	}
}

// CreateGlobalEipInvoker 创建全域弹性公网IP
func (c *GeipClient) CreateGlobalEipInvoker(request *model.CreateGlobalEipRequest) *CreateGlobalEipInvoker {
	requestDef := GenReqDefForCreateGlobalEip()
	return &CreateGlobalEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGlobalEipSegment 创建全域弹性公网IP段
//
// 创建全域弹性公网IP段
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) CreateGlobalEipSegment(request *model.CreateGlobalEipSegmentRequest) (*model.CreateGlobalEipSegmentResponse, error) {
	requestDef := GenReqDefForCreateGlobalEipSegment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGlobalEipSegmentResponse), nil
	}
}

// CreateGlobalEipSegmentInvoker 创建全域弹性公网IP段
func (c *GeipClient) CreateGlobalEipSegmentInvoker(request *model.CreateGlobalEipSegmentRequest) *CreateGlobalEipSegmentInvoker {
	requestDef := GenReqDefForCreateGlobalEipSegment()
	return &CreateGlobalEipSegmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGeipSegmentTag 删除全域弹性公网IP段标签
//
// 删除全域弹性公网IP段的标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) DeleteGeipSegmentTag(request *model.DeleteGeipSegmentTagRequest) (*model.DeleteGeipSegmentTagResponse, error) {
	requestDef := GenReqDefForDeleteGeipSegmentTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGeipSegmentTagResponse), nil
	}
}

// DeleteGeipSegmentTagInvoker 删除全域弹性公网IP段标签
func (c *GeipClient) DeleteGeipSegmentTagInvoker(request *model.DeleteGeipSegmentTagRequest) *DeleteGeipSegmentTagInvoker {
	requestDef := GenReqDefForDeleteGeipSegmentTag()
	return &DeleteGeipSegmentTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGlobalEip 删除全域弹性公网IP
//
// 删除全域弹性公网IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) DeleteGlobalEip(request *model.DeleteGlobalEipRequest) (*model.DeleteGlobalEipResponse, error) {
	requestDef := GenReqDefForDeleteGlobalEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGlobalEipResponse), nil
	}
}

// DeleteGlobalEipInvoker 删除全域弹性公网IP
func (c *GeipClient) DeleteGlobalEipInvoker(request *model.DeleteGlobalEipRequest) *DeleteGlobalEipInvoker {
	requestDef := GenReqDefForDeleteGlobalEip()
	return &DeleteGlobalEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGlobalEipSegment 删除全域弹性公网IP段
//
// 删除全域弹性公网IP段
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) DeleteGlobalEipSegment(request *model.DeleteGlobalEipSegmentRequest) (*model.DeleteGlobalEipSegmentResponse, error) {
	requestDef := GenReqDefForDeleteGlobalEipSegment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGlobalEipSegmentResponse), nil
	}
}

// DeleteGlobalEipSegmentInvoker 删除全域弹性公网IP段
func (c *GeipClient) DeleteGlobalEipSegmentInvoker(request *model.DeleteGlobalEipSegmentRequest) *DeleteGlobalEipSegmentInvoker {
	requestDef := GenReqDefForDeleteGlobalEipSegment()
	return &DeleteGlobalEipSegmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGlobalEipTag 删除全域弹性公网IP标签
//
// 删除全域弹性公网IP的标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) DeleteGlobalEipTag(request *model.DeleteGlobalEipTagRequest) (*model.DeleteGlobalEipTagResponse, error) {
	requestDef := GenReqDefForDeleteGlobalEipTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGlobalEipTagResponse), nil
	}
}

// DeleteGlobalEipTagInvoker 删除全域弹性公网IP标签
func (c *GeipClient) DeleteGlobalEipTagInvoker(request *model.DeleteGlobalEipTagRequest) *DeleteGlobalEipTagInvoker {
	requestDef := GenReqDefForDeleteGlobalEipTag()
	return &DeleteGlobalEipTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachInternetBandwidth 解绑全域公网带宽
//
// 解绑全域公网带宽
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) DetachInternetBandwidth(request *model.DetachInternetBandwidthRequest) (*model.DetachInternetBandwidthResponse, error) {
	requestDef := GenReqDefForDetachInternetBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachInternetBandwidthResponse), nil
	}
}

// DetachInternetBandwidthInvoker 解绑全域公网带宽
func (c *GeipClient) DetachInternetBandwidthInvoker(request *model.DetachInternetBandwidthRequest) *DetachInternetBandwidthInvoker {
	requestDef := GenReqDefForDetachInternetBandwidth()
	return &DetachInternetBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateGeipSegmentInstance 全域弹性公网IP段解绑后端实例
//
// 全域弹性公网IP段解绑后端实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) DisassociateGeipSegmentInstance(request *model.DisassociateGeipSegmentInstanceRequest) (*model.DisassociateGeipSegmentInstanceResponse, error) {
	requestDef := GenReqDefForDisassociateGeipSegmentInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateGeipSegmentInstanceResponse), nil
	}
}

// DisassociateGeipSegmentInstanceInvoker 全域弹性公网IP段解绑后端实例
func (c *GeipClient) DisassociateGeipSegmentInstanceInvoker(request *model.DisassociateGeipSegmentInstanceRequest) *DisassociateGeipSegmentInstanceInvoker {
	requestDef := GenReqDefForDisassociateGeipSegmentInstance()
	return &DisassociateGeipSegmentInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateInstance 解绑后端实例
//
// 解绑后端实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) DisassociateInstance(request *model.DisassociateInstanceRequest) (*model.DisassociateInstanceResponse, error) {
	requestDef := GenReqDefForDisassociateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateInstanceResponse), nil
	}
}

// DisassociateInstanceInvoker 解绑后端实例
func (c *GeipClient) DisassociateInstanceInvoker(request *model.DisassociateInstanceRequest) *DisassociateInstanceInvoker {
	requestDef := GenReqDefForDisassociateInstance()
	return &DisassociateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGeipPools 查询全域弹性公网IP池列表
//
// 查询全域弹性公网IP池列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListGeipPools(request *model.ListGeipPoolsRequest) (*model.ListGeipPoolsResponse, error) {
	requestDef := GenReqDefForListGeipPools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGeipPoolsResponse), nil
	}
}

// ListGeipPoolsInvoker 查询全域弹性公网IP池列表
func (c *GeipClient) ListGeipPoolsInvoker(request *model.ListGeipPoolsRequest) *ListGeipPoolsInvoker {
	requestDef := GenReqDefForListGeipPools()
	return &ListGeipPoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGeipSegmentCountFilterTags 查询资源实例列表数目
//
// 查询资源实例列表的数目
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListGeipSegmentCountFilterTags(request *model.ListGeipSegmentCountFilterTagsRequest) (*model.ListGeipSegmentCountFilterTagsResponse, error) {
	requestDef := GenReqDefForListGeipSegmentCountFilterTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGeipSegmentCountFilterTagsResponse), nil
	}
}

// ListGeipSegmentCountFilterTagsInvoker 查询资源实例列表数目
func (c *GeipClient) ListGeipSegmentCountFilterTagsInvoker(request *model.ListGeipSegmentCountFilterTagsRequest) *ListGeipSegmentCountFilterTagsInvoker {
	requestDef := GenReqDefForListGeipSegmentCountFilterTags()
	return &ListGeipSegmentCountFilterTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGeipSegmentDomainTags 查询全域弹性公网IP段项目标签
//
// 查询全域弹性公网IP段的项目标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListGeipSegmentDomainTags(request *model.ListGeipSegmentDomainTagsRequest) (*model.ListGeipSegmentDomainTagsResponse, error) {
	requestDef := GenReqDefForListGeipSegmentDomainTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGeipSegmentDomainTagsResponse), nil
	}
}

// ListGeipSegmentDomainTagsInvoker 查询全域弹性公网IP段项目标签
func (c *GeipClient) ListGeipSegmentDomainTagsInvoker(request *model.ListGeipSegmentDomainTagsRequest) *ListGeipSegmentDomainTagsInvoker {
	requestDef := GenReqDefForListGeipSegmentDomainTags()
	return &ListGeipSegmentDomainTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGeipSegmentFilterTags 查询资源实例列表
//
// 查询资源实例的列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListGeipSegmentFilterTags(request *model.ListGeipSegmentFilterTagsRequest) (*model.ListGeipSegmentFilterTagsResponse, error) {
	requestDef := GenReqDefForListGeipSegmentFilterTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGeipSegmentFilterTagsResponse), nil
	}
}

// ListGeipSegmentFilterTagsInvoker 查询资源实例列表
func (c *GeipClient) ListGeipSegmentFilterTagsInvoker(request *model.ListGeipSegmentFilterTagsRequest) *ListGeipSegmentFilterTagsInvoker {
	requestDef := GenReqDefForListGeipSegmentFilterTags()
	return &ListGeipSegmentFilterTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGlobalEipCountFilterTags 查询资源实例列表数目
//
// 查询资源实例列表数目
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListGlobalEipCountFilterTags(request *model.ListGlobalEipCountFilterTagsRequest) (*model.ListGlobalEipCountFilterTagsResponse, error) {
	requestDef := GenReqDefForListGlobalEipCountFilterTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGlobalEipCountFilterTagsResponse), nil
	}
}

// ListGlobalEipCountFilterTagsInvoker 查询资源实例列表数目
func (c *GeipClient) ListGlobalEipCountFilterTagsInvoker(request *model.ListGlobalEipCountFilterTagsRequest) *ListGlobalEipCountFilterTagsInvoker {
	requestDef := GenReqDefForListGlobalEipCountFilterTags()
	return &ListGlobalEipCountFilterTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGlobalEipDomainTags 查询全域弹性公网IP项目标签
//
// 查询全域弹性公网IP的项目标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListGlobalEipDomainTags(request *model.ListGlobalEipDomainTagsRequest) (*model.ListGlobalEipDomainTagsResponse, error) {
	requestDef := GenReqDefForListGlobalEipDomainTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGlobalEipDomainTagsResponse), nil
	}
}

// ListGlobalEipDomainTagsInvoker 查询全域弹性公网IP项目标签
func (c *GeipClient) ListGlobalEipDomainTagsInvoker(request *model.ListGlobalEipDomainTagsRequest) *ListGlobalEipDomainTagsInvoker {
	requestDef := GenReqDefForListGlobalEipDomainTags()
	return &ListGlobalEipDomainTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGlobalEipFilterTags 查询资源实例列表
//
// 查询资源实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListGlobalEipFilterTags(request *model.ListGlobalEipFilterTagsRequest) (*model.ListGlobalEipFilterTagsResponse, error) {
	requestDef := GenReqDefForListGlobalEipFilterTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGlobalEipFilterTagsResponse), nil
	}
}

// ListGlobalEipFilterTagsInvoker 查询资源实例列表
func (c *GeipClient) ListGlobalEipFilterTagsInvoker(request *model.ListGlobalEipFilterTagsRequest) *ListGlobalEipFilterTagsInvoker {
	requestDef := GenReqDefForListGlobalEipFilterTags()
	return &ListGlobalEipFilterTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGlobalEipSegments 查询全域弹性公网IP段列表
//
// 查询全域弹性公网IP段列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListGlobalEipSegments(request *model.ListGlobalEipSegmentsRequest) (*model.ListGlobalEipSegmentsResponse, error) {
	requestDef := GenReqDefForListGlobalEipSegments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGlobalEipSegmentsResponse), nil
	}
}

// ListGlobalEipSegmentsInvoker 查询全域弹性公网IP段列表
func (c *GeipClient) ListGlobalEipSegmentsInvoker(request *model.ListGlobalEipSegmentsRequest) *ListGlobalEipSegmentsInvoker {
	requestDef := GenReqDefForListGlobalEipSegments()
	return &ListGlobalEipSegmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGlobalEips 查询全域弹性公网IP列表
//
// 查询全域弹性公网IP列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListGlobalEips(request *model.ListGlobalEipsRequest) (*model.ListGlobalEipsResponse, error) {
	requestDef := GenReqDefForListGlobalEips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGlobalEipsResponse), nil
	}
}

// ListGlobalEipsInvoker 查询全域弹性公网IP列表
func (c *GeipClient) ListGlobalEipsInvoker(request *model.ListGlobalEipsRequest) *ListGlobalEipsInvoker {
	requestDef := GenReqDefForListGlobalEips()
	return &ListGlobalEipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGeipSegmentTags 查询全域弹性公网IP段标签
//
// 查询全域弹性公网IP段的标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ShowGeipSegmentTags(request *model.ShowGeipSegmentTagsRequest) (*model.ShowGeipSegmentTagsResponse, error) {
	requestDef := GenReqDefForShowGeipSegmentTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGeipSegmentTagsResponse), nil
	}
}

// ShowGeipSegmentTagsInvoker 查询全域弹性公网IP段标签
func (c *GeipClient) ShowGeipSegmentTagsInvoker(request *model.ShowGeipSegmentTagsRequest) *ShowGeipSegmentTagsInvoker {
	requestDef := GenReqDefForShowGeipSegmentTags()
	return &ShowGeipSegmentTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGlobalEip 查询全域弹性公网IP详情
//
// 查询全域弹性公网IP详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ShowGlobalEip(request *model.ShowGlobalEipRequest) (*model.ShowGlobalEipResponse, error) {
	requestDef := GenReqDefForShowGlobalEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGlobalEipResponse), nil
	}
}

// ShowGlobalEipInvoker 查询全域弹性公网IP详情
func (c *GeipClient) ShowGlobalEipInvoker(request *model.ShowGlobalEipRequest) *ShowGlobalEipInvoker {
	requestDef := GenReqDefForShowGlobalEip()
	return &ShowGlobalEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGlobalEipSegment 查询全域弹性公网IP段详情
//
// 查询全域弹性公网IP段详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ShowGlobalEipSegment(request *model.ShowGlobalEipSegmentRequest) (*model.ShowGlobalEipSegmentResponse, error) {
	requestDef := GenReqDefForShowGlobalEipSegment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGlobalEipSegmentResponse), nil
	}
}

// ShowGlobalEipSegmentInvoker 查询全域弹性公网IP段详情
func (c *GeipClient) ShowGlobalEipSegmentInvoker(request *model.ShowGlobalEipSegmentRequest) *ShowGlobalEipSegmentInvoker {
	requestDef := GenReqDefForShowGlobalEipSegment()
	return &ShowGlobalEipSegmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGlobalEipTags 查询全域弹性公网IP标签
//
// 查询全域弹性公网IP的标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ShowGlobalEipTags(request *model.ShowGlobalEipTagsRequest) (*model.ShowGlobalEipTagsResponse, error) {
	requestDef := GenReqDefForShowGlobalEipTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGlobalEipTagsResponse), nil
	}
}

// ShowGlobalEipTagsInvoker 查询全域弹性公网IP标签
func (c *GeipClient) ShowGlobalEipTagsInvoker(request *model.ShowGlobalEipTagsRequest) *ShowGlobalEipTagsInvoker {
	requestDef := GenReqDefForShowGlobalEipTags()
	return &ShowGlobalEipTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGlobalEip 更新全域弹性公网IP信息
//
// 更新全域弹性公网IP信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) UpdateGlobalEip(request *model.UpdateGlobalEipRequest) (*model.UpdateGlobalEipResponse, error) {
	requestDef := GenReqDefForUpdateGlobalEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGlobalEipResponse), nil
	}
}

// UpdateGlobalEipInvoker 更新全域弹性公网IP信息
func (c *GeipClient) UpdateGlobalEipInvoker(request *model.UpdateGlobalEipRequest) *UpdateGlobalEipInvoker {
	requestDef := GenReqDefForUpdateGlobalEip()
	return &UpdateGlobalEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGlobalEipSegment 更新全域弹性公网IP段
//
// 更新全域弹性公网IP段
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) UpdateGlobalEipSegment(request *model.UpdateGlobalEipSegmentRequest) (*model.UpdateGlobalEipSegmentResponse, error) {
	requestDef := GenReqDefForUpdateGlobalEipSegment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGlobalEipSegmentResponse), nil
	}
}

// UpdateGlobalEipSegmentInvoker 更新全域弹性公网IP段
func (c *GeipClient) UpdateGlobalEipSegmentInvoker(request *model.UpdateGlobalEipSegmentRequest) *UpdateGlobalEipSegmentInvoker {
	requestDef := GenReqDefForUpdateGlobalEipSegment()
	return &UpdateGlobalEipSegmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobs 查询Job列表
//
// 查询Job列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListJobs(request *model.ListJobsRequest) (*model.ListJobsResponse, error) {
	requestDef := GenReqDefForListJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobsResponse), nil
	}
}

// ListJobsInvoker 查询Job列表
func (c *GeipClient) ListJobsInvoker(request *model.ListJobsRequest) *ListJobsInvoker {
	requestDef := GenReqDefForListJobs()
	return &ListJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobs 查询Job详情
//
// 查询Job详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ShowJobs(request *model.ShowJobsRequest) (*model.ShowJobsResponse, error) {
	requestDef := GenReqDefForShowJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobsResponse), nil
	}
}

// ShowJobsInvoker 查询Job详情
func (c *GeipClient) ShowJobsInvoker(request *model.ShowJobsRequest) *ShowJobsInvoker {
	requestDef := GenReqDefForShowJobs()
	return &ShowJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSupportRegions 全域弹性公网IP支持绑定的Region限制
//
// 全域弹性公网IP支持绑定的Region限制
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListSupportRegions(request *model.ListSupportRegionsRequest) (*model.ListSupportRegionsResponse, error) {
	requestDef := GenReqDefForListSupportRegions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSupportRegionsResponse), nil
	}
}

// ListSupportRegionsInvoker 全域弹性公网IP支持绑定的Region限制
func (c *GeipClient) ListSupportRegionsInvoker(request *model.ListSupportRegionsRequest) *ListSupportRegionsInvoker {
	requestDef := GenReqDefForListSupportRegions()
	return &ListSupportRegionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTenantGeipSupportInstances 查询指定站点允许绑定的Region信息
//
// console通过此接口获取指定pop点的全域弹性公网IP允许绑定的region实例信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GeipClient) ListTenantGeipSupportInstances(request *model.ListTenantGeipSupportInstancesRequest) (*model.ListTenantGeipSupportInstancesResponse, error) {
	requestDef := GenReqDefForListTenantGeipSupportInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTenantGeipSupportInstancesResponse), nil
	}
}

// ListTenantGeipSupportInstancesInvoker 查询指定站点允许绑定的Region信息
func (c *GeipClient) ListTenantGeipSupportInstancesInvoker(request *model.ListTenantGeipSupportInstancesRequest) *ListTenantGeipSupportInstancesInvoker {
	requestDef := GenReqDefForListTenantGeipSupportInstances()
	return &ListTenantGeipSupportInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
