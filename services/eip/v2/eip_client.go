package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v2/model"
)

type EipClient struct {
	HcClient *http_client.HcHttpClient
}

func NewEipClient(hcClient *http_client.HcHttpClient) *EipClient {
	return &EipClient{HcClient: hcClient}
}

func EipClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AddPublicipsIntoSharedBandwidth 共享带宽插入弹性公网IP
//
// 共享带宽插入弹性公网IP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) AddPublicipsIntoSharedBandwidth(request *model.AddPublicipsIntoSharedBandwidthRequest) (*model.AddPublicipsIntoSharedBandwidthResponse, error) {
	requestDef := GenReqDefForAddPublicipsIntoSharedBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddPublicipsIntoSharedBandwidthResponse), nil
	}
}

// AddPublicipsIntoSharedBandwidthInvoker 共享带宽插入弹性公网IP
func (c *EipClient) AddPublicipsIntoSharedBandwidthInvoker(request *model.AddPublicipsIntoSharedBandwidthRequest) *AddPublicipsIntoSharedBandwidthInvoker {
	requestDef := GenReqDefForAddPublicipsIntoSharedBandwidth()
	return &AddPublicipsIntoSharedBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateSharedBandwidths 批量创建共享带宽
//
// 批量创建共享带宽。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) BatchCreateSharedBandwidths(request *model.BatchCreateSharedBandwidthsRequest) (*model.BatchCreateSharedBandwidthsResponse, error) {
	requestDef := GenReqDefForBatchCreateSharedBandwidths()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateSharedBandwidthsResponse), nil
	}
}

// BatchCreateSharedBandwidthsInvoker 批量创建共享带宽
func (c *EipClient) BatchCreateSharedBandwidthsInvoker(request *model.BatchCreateSharedBandwidthsRequest) *BatchCreateSharedBandwidthsInvoker {
	requestDef := GenReqDefForBatchCreateSharedBandwidths()
	return &BatchCreateSharedBandwidthsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeBandwidthToPeriod 按需转包API
//
// 租户按需转包接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) ChangeBandwidthToPeriod(request *model.ChangeBandwidthToPeriodRequest) (*model.ChangeBandwidthToPeriodResponse, error) {
	requestDef := GenReqDefForChangeBandwidthToPeriod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeBandwidthToPeriodResponse), nil
	}
}

// ChangeBandwidthToPeriodInvoker 按需转包API
func (c *EipClient) ChangeBandwidthToPeriodInvoker(request *model.ChangeBandwidthToPeriodRequest) *ChangeBandwidthToPeriodInvoker {
	requestDef := GenReqDefForChangeBandwidthToPeriod()
	return &ChangeBandwidthToPeriodInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSharedBandwidth 创建共享带宽
//
// 创建共享带宽。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) CreateSharedBandwidth(request *model.CreateSharedBandwidthRequest) (*model.CreateSharedBandwidthResponse, error) {
	requestDef := GenReqDefForCreateSharedBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSharedBandwidthResponse), nil
	}
}

// CreateSharedBandwidthInvoker 创建共享带宽
func (c *EipClient) CreateSharedBandwidthInvoker(request *model.CreateSharedBandwidthRequest) *CreateSharedBandwidthInvoker {
	requestDef := GenReqDefForCreateSharedBandwidth()
	return &CreateSharedBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSharedBandwidth 删除共享带宽
//
// 删除共享带宽。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) DeleteSharedBandwidth(request *model.DeleteSharedBandwidthRequest) (*model.DeleteSharedBandwidthResponse, error) {
	requestDef := GenReqDefForDeleteSharedBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSharedBandwidthResponse), nil
	}
}

// DeleteSharedBandwidthInvoker 删除共享带宽
func (c *EipClient) DeleteSharedBandwidthInvoker(request *model.DeleteSharedBandwidthRequest) *DeleteSharedBandwidthInvoker {
	requestDef := GenReqDefForDeleteSharedBandwidth()
	return &DeleteSharedBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBandwidthPkg 查询带宽加油包列表
//
// 查询带宽加油包列表信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) ListBandwidthPkg(request *model.ListBandwidthPkgRequest) (*model.ListBandwidthPkgResponse, error) {
	requestDef := GenReqDefForListBandwidthPkg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBandwidthPkgResponse), nil
	}
}

// ListBandwidthPkgInvoker 查询带宽加油包列表
func (c *EipClient) ListBandwidthPkgInvoker(request *model.ListBandwidthPkgRequest) *ListBandwidthPkgInvoker {
	requestDef := GenReqDefForListBandwidthPkg()
	return &ListBandwidthPkgInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBandwidths 查询带宽列表
//
// 查询带宽列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) ListBandwidths(request *model.ListBandwidthsRequest) (*model.ListBandwidthsResponse, error) {
	requestDef := GenReqDefForListBandwidths()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBandwidthsResponse), nil
	}
}

// ListBandwidthsInvoker 查询带宽列表
func (c *EipClient) ListBandwidthsInvoker(request *model.ListBandwidthsRequest) *ListBandwidthsInvoker {
	requestDef := GenReqDefForListBandwidths()
	return &ListBandwidthsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotas 查询配额接口
//
// 查询配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

// ListQuotasInvoker 查询配额接口
func (c *EipClient) ListQuotasInvoker(request *model.ListQuotasRequest) *ListQuotasInvoker {
	requestDef := GenReqDefForListQuotas()
	return &ListQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemovePublicipsFromSharedBandwidth 共享带宽移除弹性公网IP
//
// 共享带宽移除弹性公网IP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) RemovePublicipsFromSharedBandwidth(request *model.RemovePublicipsFromSharedBandwidthRequest) (*model.RemovePublicipsFromSharedBandwidthResponse, error) {
	requestDef := GenReqDefForRemovePublicipsFromSharedBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemovePublicipsFromSharedBandwidthResponse), nil
	}
}

// RemovePublicipsFromSharedBandwidthInvoker 共享带宽移除弹性公网IP
func (c *EipClient) RemovePublicipsFromSharedBandwidthInvoker(request *model.RemovePublicipsFromSharedBandwidthRequest) *RemovePublicipsFromSharedBandwidthInvoker {
	requestDef := GenReqDefForRemovePublicipsFromSharedBandwidth()
	return &RemovePublicipsFromSharedBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBandwidth 查询带宽
//
// 查询带宽
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) ShowBandwidth(request *model.ShowBandwidthRequest) (*model.ShowBandwidthResponse, error) {
	requestDef := GenReqDefForShowBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBandwidthResponse), nil
	}
}

// ShowBandwidthInvoker 查询带宽
func (c *EipClient) ShowBandwidthInvoker(request *model.ShowBandwidthRequest) *ShowBandwidthInvoker {
	requestDef := GenReqDefForShowBandwidth()
	return &ShowBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBandwidth 更新带宽
//
// 更新带宽。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) UpdateBandwidth(request *model.UpdateBandwidthRequest) (*model.UpdateBandwidthResponse, error) {
	requestDef := GenReqDefForUpdateBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBandwidthResponse), nil
	}
}

// UpdateBandwidthInvoker 更新带宽
func (c *EipClient) UpdateBandwidthInvoker(request *model.UpdateBandwidthRequest) *UpdateBandwidthInvoker {
	requestDef := GenReqDefForUpdateBandwidth()
	return &UpdateBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePrePaidBandwidth 更新包周期带宽
//
// 更新带宽。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) UpdatePrePaidBandwidth(request *model.UpdatePrePaidBandwidthRequest) (*model.UpdatePrePaidBandwidthResponse, error) {
	requestDef := GenReqDefForUpdatePrePaidBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePrePaidBandwidthResponse), nil
	}
}

// UpdatePrePaidBandwidthInvoker 更新包周期带宽
func (c *EipClient) UpdatePrePaidBandwidthInvoker(request *model.UpdatePrePaidBandwidthRequest) *UpdatePrePaidBandwidthInvoker {
	requestDef := GenReqDefForUpdatePrePaidBandwidth()
	return &UpdatePrePaidBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreatePublicipTags 批量创建弹性公网IP资源标签
//
// 为指定的弹性公网IP资源实例批量添加标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) BatchCreatePublicipTags(request *model.BatchCreatePublicipTagsRequest) (*model.BatchCreatePublicipTagsResponse, error) {
	requestDef := GenReqDefForBatchCreatePublicipTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreatePublicipTagsResponse), nil
	}
}

// BatchCreatePublicipTagsInvoker 批量创建弹性公网IP资源标签
func (c *EipClient) BatchCreatePublicipTagsInvoker(request *model.BatchCreatePublicipTagsRequest) *BatchCreatePublicipTagsInvoker {
	requestDef := GenReqDefForBatchCreatePublicipTags()
	return &BatchCreatePublicipTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreatePublicips 批量创建弹性公网IP
//
// 批量创建弹性公网IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) BatchCreatePublicips(request *model.BatchCreatePublicipsRequest) (*model.BatchCreatePublicipsResponse, error) {
	requestDef := GenReqDefForBatchCreatePublicips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreatePublicipsResponse), nil
	}
}

// BatchCreatePublicipsInvoker 批量创建弹性公网IP
func (c *EipClient) BatchCreatePublicipsInvoker(request *model.BatchCreatePublicipsRequest) *BatchCreatePublicipsInvoker {
	requestDef := GenReqDefForBatchCreatePublicips()
	return &BatchCreatePublicipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeletePublicIp 批量删除弹性公网IP
//
// 批量删除弹性公网IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) BatchDeletePublicIp(request *model.BatchDeletePublicIpRequest) (*model.BatchDeletePublicIpResponse, error) {
	requestDef := GenReqDefForBatchDeletePublicIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeletePublicIpResponse), nil
	}
}

// BatchDeletePublicIpInvoker 批量删除弹性公网IP
func (c *EipClient) BatchDeletePublicIpInvoker(request *model.BatchDeletePublicIpRequest) *BatchDeletePublicIpInvoker {
	requestDef := GenReqDefForBatchDeletePublicIp()
	return &BatchDeletePublicIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeletePublicipTags 批量删除弹性公网IP资源标签
//
// 为指定的弹性公网IP资源实例批量删除标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) BatchDeletePublicipTags(request *model.BatchDeletePublicipTagsRequest) (*model.BatchDeletePublicipTagsResponse, error) {
	requestDef := GenReqDefForBatchDeletePublicipTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeletePublicipTagsResponse), nil
	}
}

// BatchDeletePublicipTagsInvoker 批量删除弹性公网IP资源标签
func (c *EipClient) BatchDeletePublicipTagsInvoker(request *model.BatchDeletePublicipTagsRequest) *BatchDeletePublicipTagsInvoker {
	requestDef := GenReqDefForBatchDeletePublicipTags()
	return &BatchDeletePublicipTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDisassociatePublicips 批量解绑弹性公网IP
//
// 批量解绑弹性公网IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) BatchDisassociatePublicips(request *model.BatchDisassociatePublicipsRequest) (*model.BatchDisassociatePublicipsResponse, error) {
	requestDef := GenReqDefForBatchDisassociatePublicips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDisassociatePublicipsResponse), nil
	}
}

// BatchDisassociatePublicipsInvoker 批量解绑弹性公网IP
func (c *EipClient) BatchDisassociatePublicipsInvoker(request *model.BatchDisassociatePublicipsRequest) *BatchDisassociatePublicipsInvoker {
	requestDef := GenReqDefForBatchDisassociatePublicips()
	return &BatchDisassociatePublicipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangePublicipToPeriod 按需转包接口
//
// 租户按需转包接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) ChangePublicipToPeriod(request *model.ChangePublicipToPeriodRequest) (*model.ChangePublicipToPeriodResponse, error) {
	requestDef := GenReqDefForChangePublicipToPeriod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangePublicipToPeriodResponse), nil
	}
}

// ChangePublicipToPeriodInvoker 按需转包接口
func (c *EipClient) ChangePublicipToPeriodInvoker(request *model.ChangePublicipToPeriodRequest) *ChangePublicipToPeriodInvoker {
	requestDef := GenReqDefForChangePublicipToPeriod()
	return &ChangePublicipToPeriodInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountPublicIp 查询PublicIp数量
//
// 查询PublicIp数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) CountPublicIp(request *model.CountPublicIpRequest) (*model.CountPublicIpResponse, error) {
	requestDef := GenReqDefForCountPublicIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountPublicIpResponse), nil
	}
}

// CountPublicIpInvoker 查询PublicIp数量
func (c *EipClient) CountPublicIpInvoker(request *model.CountPublicIpRequest) *CountPublicIpInvoker {
	requestDef := GenReqDefForCountPublicIp()
	return &CountPublicIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountPublicIpInstance 查询PublicIp实例数
//
// 查询PublicIp实例数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) CountPublicIpInstance(request *model.CountPublicIpInstanceRequest) (*model.CountPublicIpInstanceResponse, error) {
	requestDef := GenReqDefForCountPublicIpInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountPublicIpInstanceResponse), nil
	}
}

// CountPublicIpInstanceInvoker 查询PublicIp实例数
func (c *EipClient) CountPublicIpInstanceInvoker(request *model.CountPublicIpInstanceRequest) *CountPublicIpInstanceInvoker {
	requestDef := GenReqDefForCountPublicIpInstance()
	return &CountPublicIpInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePrePaidPublicip 申请包周期弹性公网IP
//
// 申请包年包月的弹性公网IP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) CreatePrePaidPublicip(request *model.CreatePrePaidPublicipRequest) (*model.CreatePrePaidPublicipResponse, error) {
	requestDef := GenReqDefForCreatePrePaidPublicip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrePaidPublicipResponse), nil
	}
}

// CreatePrePaidPublicipInvoker 申请包周期弹性公网IP
func (c *EipClient) CreatePrePaidPublicipInvoker(request *model.CreatePrePaidPublicipRequest) *CreatePrePaidPublicipInvoker {
	requestDef := GenReqDefForCreatePrePaidPublicip()
	return &CreatePrePaidPublicipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePublicip 申请弹性公网IP
//
// 申请弹性公网IP，支持IPv4和IPv6。
//  弹性公网IP（Elastic IP）提供独立的公网IP资源，包括公网IP地址与公网出口带宽服务。可以与弹性云服务器、裸金属服务器、虚拟IP、弹性负载均衡、NAT网关等资源灵活地绑定及解绑。拥有多种灵活的计费方式，可以满足各种业务场景的需要。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) CreatePublicip(request *model.CreatePublicipRequest) (*model.CreatePublicipResponse, error) {
	requestDef := GenReqDefForCreatePublicip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePublicipResponse), nil
	}
}

// CreatePublicipInvoker 申请弹性公网IP
func (c *EipClient) CreatePublicipInvoker(request *model.CreatePublicipRequest) *CreatePublicipInvoker {
	requestDef := GenReqDefForCreatePublicip()
	return &CreatePublicipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePublicipTag 创建弹性公网IP资源标签
//
// 给指定弹性IP资源实例增加标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) CreatePublicipTag(request *model.CreatePublicipTagRequest) (*model.CreatePublicipTagResponse, error) {
	requestDef := GenReqDefForCreatePublicipTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePublicipTagResponse), nil
	}
}

// CreatePublicipTagInvoker 创建弹性公网IP资源标签
func (c *EipClient) CreatePublicipTagInvoker(request *model.CreatePublicipTagRequest) *CreatePublicipTagInvoker {
	requestDef := GenReqDefForCreatePublicipTag()
	return &CreatePublicipTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePublicip 删除弹性公网IP
//
// 删除弹性公网IP,绑定状态eip不允许直接删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) DeletePublicip(request *model.DeletePublicipRequest) (*model.DeletePublicipResponse, error) {
	requestDef := GenReqDefForDeletePublicip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePublicipResponse), nil
	}
}

// DeletePublicipInvoker 删除弹性公网IP
func (c *EipClient) DeletePublicipInvoker(request *model.DeletePublicipRequest) *DeletePublicipInvoker {
	requestDef := GenReqDefForDeletePublicip()
	return &DeletePublicipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePublicipTag 删除弹性公网IP的标签
//
// 删除指定弹性公网IP的标签信息。其中project_id是项目ID，publicip_id 是要操作的弹性公网IP的id。key是要删除标签的键。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) DeletePublicipTag(request *model.DeletePublicipTagRequest) (*model.DeletePublicipTagResponse, error) {
	requestDef := GenReqDefForDeletePublicipTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePublicipTagResponse), nil
	}
}

// DeletePublicipTagInvoker 删除弹性公网IP的标签
func (c *EipClient) DeletePublicipTagInvoker(request *model.DeletePublicipTagRequest) *DeletePublicipTagInvoker {
	requestDef := GenReqDefForDeletePublicipTag()
	return &DeletePublicipTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPublicipTags 查询租户的弹性公网IP标签
//
// 查询租户在指定区域和实例类型的所有标签集合。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) ListPublicipTags(request *model.ListPublicipTagsRequest) (*model.ListPublicipTagsResponse, error) {
	requestDef := GenReqDefForListPublicipTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublicipTagsResponse), nil
	}
}

// ListPublicipTagsInvoker 查询租户的弹性公网IP标签
func (c *EipClient) ListPublicipTagsInvoker(request *model.ListPublicipTagsRequest) *ListPublicipTagsInvoker {
	requestDef := GenReqDefForListPublicipTags()
	return &ListPublicipTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPublicips 查询弹性公网IP列表
//
// 查询弹性公网IP列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) ListPublicips(request *model.ListPublicipsRequest) (*model.ListPublicipsResponse, error) {
	requestDef := GenReqDefForListPublicips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublicipsResponse), nil
	}
}

// ListPublicipsInvoker 查询弹性公网IP列表
func (c *EipClient) ListPublicipsInvoker(request *model.ListPublicipsRequest) *ListPublicipsInvoker {
	requestDef := GenReqDefForListPublicips()
	return &ListPublicipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPublicipsByTags 按标签查询弹性公网IP列表
//
// 使用标签过滤弹性公网IP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) ListPublicipsByTags(request *model.ListPublicipsByTagsRequest) (*model.ListPublicipsByTagsResponse, error) {
	requestDef := GenReqDefForListPublicipsByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublicipsByTagsResponse), nil
	}
}

// ListPublicipsByTagsInvoker 按标签查询弹性公网IP列表
func (c *EipClient) ListPublicipsByTagsInvoker(request *model.ListPublicipsByTagsRequest) *ListPublicipsByTagsInvoker {
	requestDef := GenReqDefForListPublicipsByTags()
	return &ListPublicipsByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPublicIpType 查询PublicIp类型
//
// 查询PublicIp类型
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) ShowPublicIpType(request *model.ShowPublicIpTypeRequest) (*model.ShowPublicIpTypeResponse, error) {
	requestDef := GenReqDefForShowPublicIpType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicIpTypeResponse), nil
	}
}

// ShowPublicIpTypeInvoker 查询PublicIp类型
func (c *EipClient) ShowPublicIpTypeInvoker(request *model.ShowPublicIpTypeRequest) *ShowPublicIpTypeInvoker {
	requestDef := GenReqDefForShowPublicIpType()
	return &ShowPublicIpTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPublicip 查询弹性公网IP
//
// 查询指定的弹性公网IP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) ShowPublicip(request *model.ShowPublicipRequest) (*model.ShowPublicipResponse, error) {
	requestDef := GenReqDefForShowPublicip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicipResponse), nil
	}
}

// ShowPublicipInvoker 查询弹性公网IP
func (c *EipClient) ShowPublicipInvoker(request *model.ShowPublicipRequest) *ShowPublicipInvoker {
	requestDef := GenReqDefForShowPublicip()
	return &ShowPublicipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPublicipTags 查询弹性公网IP的标签
//
// 查询指定弹性IP实例的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) ShowPublicipTags(request *model.ShowPublicipTagsRequest) (*model.ShowPublicipTagsResponse, error) {
	requestDef := GenReqDefForShowPublicipTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicipTagsResponse), nil
	}
}

// ShowPublicipTagsInvoker 查询弹性公网IP的标签
func (c *EipClient) ShowPublicipTagsInvoker(request *model.ShowPublicipTagsRequest) *ShowPublicipTagsInvoker {
	requestDef := GenReqDefForShowPublicipTags()
	return &ShowPublicipTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePublicip 更新弹性公网IP
//
// 更新弹性公网IP，将弹性公网IP跟一个网卡绑定或者解绑定，转换IP地址版本类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) UpdatePublicip(request *model.UpdatePublicipRequest) (*model.UpdatePublicipResponse, error) {
	requestDef := GenReqDefForUpdatePublicip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePublicipResponse), nil
	}
}

// UpdatePublicipInvoker 更新弹性公网IP
func (c *EipClient) UpdatePublicipInvoker(request *model.UpdatePublicipRequest) *UpdatePublicipInvoker {
	requestDef := GenReqDefForUpdatePublicip()
	return &UpdatePublicipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourcesJobDetail 查询Job状态接口
//
// 查询Job状态接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) ShowResourcesJobDetail(request *model.ShowResourcesJobDetailRequest) (*model.ShowResourcesJobDetailResponse, error) {
	requestDef := GenReqDefForShowResourcesJobDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourcesJobDetailResponse), nil
	}
}

// ShowResourcesJobDetailInvoker 查询Job状态接口
func (c *EipClient) ShowResourcesJobDetailInvoker(request *model.ShowResourcesJobDetailRequest) *ShowResourcesJobDetailInvoker {
	requestDef := GenReqDefForShowResourcesJobDetail()
	return &ShowResourcesJobDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronCreateFloatingIp 创建浮动IP
//
// 创建浮动IP的外部网络UUID，请使用GET /v2.0/networks?router:external&#x3D;True或neutron net-external-list方式获取。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) NeutronCreateFloatingIp(request *model.NeutronCreateFloatingIpRequest) (*model.NeutronCreateFloatingIpResponse, error) {
	requestDef := GenReqDefForNeutronCreateFloatingIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronCreateFloatingIpResponse), nil
	}
}

// NeutronCreateFloatingIpInvoker 创建浮动IP
func (c *EipClient) NeutronCreateFloatingIpInvoker(request *model.NeutronCreateFloatingIpRequest) *NeutronCreateFloatingIpInvoker {
	requestDef := GenReqDefForNeutronCreateFloatingIp()
	return &NeutronCreateFloatingIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronDeleteFloatingIp 删除浮动IP
//
// 删除指定的浮动IP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) NeutronDeleteFloatingIp(request *model.NeutronDeleteFloatingIpRequest) (*model.NeutronDeleteFloatingIpResponse, error) {
	requestDef := GenReqDefForNeutronDeleteFloatingIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronDeleteFloatingIpResponse), nil
	}
}

// NeutronDeleteFloatingIpInvoker 删除浮动IP
func (c *EipClient) NeutronDeleteFloatingIpInvoker(request *model.NeutronDeleteFloatingIpRequest) *NeutronDeleteFloatingIpInvoker {
	requestDef := GenReqDefForNeutronDeleteFloatingIp()
	return &NeutronDeleteFloatingIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronListFloatingIps 查询浮动IP列表
//
// 查询提交请求的租户有权限操作的所有浮动IP地址。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) NeutronListFloatingIps(request *model.NeutronListFloatingIpsRequest) (*model.NeutronListFloatingIpsResponse, error) {
	requestDef := GenReqDefForNeutronListFloatingIps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronListFloatingIpsResponse), nil
	}
}

// NeutronListFloatingIpsInvoker 查询浮动IP列表
func (c *EipClient) NeutronListFloatingIpsInvoker(request *model.NeutronListFloatingIpsRequest) *NeutronListFloatingIpsInvoker {
	requestDef := GenReqDefForNeutronListFloatingIps()
	return &NeutronListFloatingIpsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronShowFloatingIp 查询浮动IP
//
// 查询浮动IP详情，包括浮动IP状态，浮动IP所属路由器ID，浮动IP的外部网络ID等等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) NeutronShowFloatingIp(request *model.NeutronShowFloatingIpRequest) (*model.NeutronShowFloatingIpResponse, error) {
	requestDef := GenReqDefForNeutronShowFloatingIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronShowFloatingIpResponse), nil
	}
}

// NeutronShowFloatingIpInvoker 查询浮动IP
func (c *EipClient) NeutronShowFloatingIpInvoker(request *model.NeutronShowFloatingIpRequest) *NeutronShowFloatingIpInvoker {
	requestDef := GenReqDefForNeutronShowFloatingIp()
	return &NeutronShowFloatingIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronUpdateFloatingIp 更新浮动IP
//
// 更新浮动IP。
//  更新时需在URL中给出浮动IP地址的ID。
//  port_id 为空，则表示浮动IP从端口解绑。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) NeutronUpdateFloatingIp(request *model.NeutronUpdateFloatingIpRequest) (*model.NeutronUpdateFloatingIpResponse, error) {
	requestDef := GenReqDefForNeutronUpdateFloatingIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronUpdateFloatingIpResponse), nil
	}
}

// NeutronUpdateFloatingIpInvoker 更新浮动IP
func (c *EipClient) NeutronUpdateFloatingIpInvoker(request *model.NeutronUpdateFloatingIpRequest) *NeutronUpdateFloatingIpInvoker {
	requestDef := GenReqDefForNeutronUpdateFloatingIp()
	return &NeutronUpdateFloatingIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
