package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v3/model"
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

// ListBandwidth 查询带宽列表
//
// 查询带宽列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) ListBandwidth(request *model.ListBandwidthRequest) (*model.ListBandwidthResponse, error) {
	requestDef := GenReqDefForListBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBandwidthResponse), nil
	}
}

// ListBandwidthInvoker 查询带宽列表
func (c *EipClient) ListBandwidthInvoker(request *model.ListBandwidthRequest) *ListBandwidthInvoker {
	requestDef := GenReqDefForListBandwidth()
	return &ListBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCommonPools 查询公共池列表
//
// 查询公共池列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) ListCommonPools(request *model.ListCommonPoolsRequest) (*model.ListCommonPoolsResponse, error) {
	requestDef := GenReqDefForListCommonPools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCommonPoolsResponse), nil
	}
}

// ListCommonPoolsInvoker 查询公共池列表
func (c *EipClient) ListCommonPoolsInvoker(request *model.ListCommonPoolsRequest) *ListCommonPoolsInvoker {
	requestDef := GenReqDefForListCommonPools()
	return &ListCommonPoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPublicBorderGroups 查询公共池分组列表
//
// 查询公共池分组列表，包含名称和位置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) ListPublicBorderGroups(request *model.ListPublicBorderGroupsRequest) (*model.ListPublicBorderGroupsResponse, error) {
	requestDef := GenReqDefForListPublicBorderGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublicBorderGroupsResponse), nil
	}
}

// ListPublicBorderGroupsInvoker 查询公共池分组列表
func (c *EipClient) ListPublicBorderGroupsInvoker(request *model.ListPublicBorderGroupsRequest) *ListPublicBorderGroupsInvoker {
	requestDef := GenReqDefForListPublicBorderGroups()
	return &ListPublicBorderGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPublicipPool 查询公网IP池列表
//
// 全量查询公网IP池列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) ListPublicipPool(request *model.ListPublicipPoolRequest) (*model.ListPublicipPoolResponse, error) {
	requestDef := GenReqDefForListPublicipPool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublicipPoolResponse), nil
	}
}

// ListPublicipPoolInvoker 查询公网IP池列表
func (c *EipClient) ListPublicipPoolInvoker(request *model.ListPublicipPoolRequest) *ListPublicipPoolInvoker {
	requestDef := GenReqDefForListPublicipPool()
	return &ListPublicipPoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListShareBandwidthTypes 查询指定租户下的共享带宽类型列表
//
// 查询指定租户下的共享带宽类型列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) ListShareBandwidthTypes(request *model.ListShareBandwidthTypesRequest) (*model.ListShareBandwidthTypesResponse, error) {
	requestDef := GenReqDefForListShareBandwidthTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListShareBandwidthTypesResponse), nil
	}
}

// ListShareBandwidthTypesInvoker 查询指定租户下的共享带宽类型列表
func (c *EipClient) ListShareBandwidthTypesInvoker(request *model.ListShareBandwidthTypesRequest) *ListShareBandwidthTypesInvoker {
	requestDef := GenReqDefForListShareBandwidthTypes()
	return &ListShareBandwidthTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPublicipPool 查询公网IP池详情
//
// 查询公网IP池详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) ShowPublicipPool(request *model.ShowPublicipPoolRequest) (*model.ShowPublicipPoolResponse, error) {
	requestDef := GenReqDefForShowPublicipPool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicipPoolResponse), nil
	}
}

// ShowPublicipPoolInvoker 查询公网IP池详情
func (c *EipClient) ShowPublicipPoolInvoker(request *model.ShowPublicipPoolRequest) *ShowPublicipPoolInvoker {
	requestDef := GenReqDefForShowPublicipPool()
	return &ShowPublicipPoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociatePublicips 绑定弹性公网IP
//
// 绑定弹性公网IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) AssociatePublicips(request *model.AssociatePublicipsRequest) (*model.AssociatePublicipsResponse, error) {
	requestDef := GenReqDefForAssociatePublicips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociatePublicipsResponse), nil
	}
}

// AssociatePublicipsInvoker 绑定弹性公网IP
func (c *EipClient) AssociatePublicipsInvoker(request *model.AssociatePublicipsRequest) *AssociatePublicipsInvoker {
	requestDef := GenReqDefForAssociatePublicips()
	return &AssociatePublicipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachBatchPublicIp 共享带宽批量加入弹性公网IP
//
// 共享带宽批量加入弹性公网IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) AttachBatchPublicIp(request *model.AttachBatchPublicIpRequest) (*model.AttachBatchPublicIpResponse, error) {
	requestDef := GenReqDefForAttachBatchPublicIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachBatchPublicIpResponse), nil
	}
}

// AttachBatchPublicIpInvoker 共享带宽批量加入弹性公网IP
func (c *EipClient) AttachBatchPublicIpInvoker(request *model.AttachBatchPublicIpRequest) *AttachBatchPublicIpInvoker {
	requestDef := GenReqDefForAttachBatchPublicIp()
	return &AttachBatchPublicIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachShareBandwidth 共享带宽加入弹性公网IP
//
// 共享带宽加入弹性公网IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) AttachShareBandwidth(request *model.AttachShareBandwidthRequest) (*model.AttachShareBandwidthResponse, error) {
	requestDef := GenReqDefForAttachShareBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachShareBandwidthResponse), nil
	}
}

// AttachShareBandwidthInvoker 共享带宽加入弹性公网IP
func (c *EipClient) AttachShareBandwidthInvoker(request *model.AttachShareBandwidthRequest) *AttachShareBandwidthInvoker {
	requestDef := GenReqDefForAttachShareBandwidth()
	return &AttachShareBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountEipAvailableResources 查询弹性公网IP可用数
//
// IP池用于查询公网可用ip个数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) CountEipAvailableResources(request *model.CountEipAvailableResourcesRequest) (*model.CountEipAvailableResourcesResponse, error) {
	requestDef := GenReqDefForCountEipAvailableResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountEipAvailableResourcesResponse), nil
	}
}

// CountEipAvailableResourcesInvoker 查询弹性公网IP可用数
func (c *EipClient) CountEipAvailableResourcesInvoker(request *model.CountEipAvailableResourcesRequest) *CountEipAvailableResourcesInvoker {
	requestDef := GenReqDefForCountEipAvailableResources()
	return &CountEipAvailableResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachBatchPublicIp 共享带宽批量移出弹性公网IP
//
// 共享带宽批量移出弹性公网IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) DetachBatchPublicIp(request *model.DetachBatchPublicIpRequest) (*model.DetachBatchPublicIpResponse, error) {
	requestDef := GenReqDefForDetachBatchPublicIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachBatchPublicIpResponse), nil
	}
}

// DetachBatchPublicIpInvoker 共享带宽批量移出弹性公网IP
func (c *EipClient) DetachBatchPublicIpInvoker(request *model.DetachBatchPublicIpRequest) *DetachBatchPublicIpInvoker {
	requestDef := GenReqDefForDetachBatchPublicIp()
	return &DetachBatchPublicIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachShareBandwidth 共享带宽移出弹性公网IP
//
// 共享带宽移出弹性公网IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) DetachShareBandwidth(request *model.DetachShareBandwidthRequest) (*model.DetachShareBandwidthResponse, error) {
	requestDef := GenReqDefForDetachShareBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachShareBandwidthResponse), nil
	}
}

// DetachShareBandwidthInvoker 共享带宽移出弹性公网IP
func (c *EipClient) DetachShareBandwidthInvoker(request *model.DetachShareBandwidthRequest) *DetachShareBandwidthInvoker {
	requestDef := GenReqDefForDetachShareBandwidth()
	return &DetachShareBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableNat64 弹性公网IP关闭NAT64
//
// 弹性公网IP关闭NAT64
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) DisableNat64(request *model.DisableNat64Request) (*model.DisableNat64Response, error) {
	requestDef := GenReqDefForDisableNat64()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableNat64Response), nil
	}
}

// DisableNat64Invoker 弹性公网IP关闭NAT64
func (c *EipClient) DisableNat64Invoker(request *model.DisableNat64Request) *DisableNat64Invoker {
	requestDef := GenReqDefForDisableNat64()
	return &DisableNat64Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociatePublicips 解绑弹性公网IP
//
// 解绑弹性公网IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) DisassociatePublicips(request *model.DisassociatePublicipsRequest) (*model.DisassociatePublicipsResponse, error) {
	requestDef := GenReqDefForDisassociatePublicips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociatePublicipsResponse), nil
	}
}

// DisassociatePublicipsInvoker 解绑弹性公网IP
func (c *EipClient) DisassociatePublicipsInvoker(request *model.DisassociatePublicipsRequest) *DisassociatePublicipsInvoker {
	requestDef := GenReqDefForDisassociatePublicips()
	return &DisassociatePublicipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableNat64 弹性公网IP开启NAT64
//
// 弹性公网IP开启NAT64
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) EnableNat64(request *model.EnableNat64Request) (*model.EnableNat64Response, error) {
	requestDef := GenReqDefForEnableNat64()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableNat64Response), nil
	}
}

// EnableNat64Invoker 弹性公网IP开启NAT64
func (c *EipClient) EnableNat64Invoker(request *model.EnableNat64Request) *EnableNat64Invoker {
	requestDef := GenReqDefForEnableNat64()
	return &EnableNat64Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPublicips 全量查询弹性公网IP列表
//
// 查询弹性公网IP列表信息
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

// ListPublicipsInvoker 全量查询弹性公网IP列表
func (c *EipClient) ListPublicipsInvoker(request *model.ListPublicipsRequest) *ListPublicipsInvoker {
	requestDef := GenReqDefForListPublicips()
	return &ListPublicipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPublicip 查询弹性公网IP详情
//
// 查询弹性公网IP详情
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

// ShowPublicipInvoker 查询弹性公网IP详情
func (c *EipClient) ShowPublicipInvoker(request *model.ShowPublicipRequest) *ShowPublicipInvoker {
	requestDef := GenReqDefForShowPublicip()
	return &ShowPublicipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAssociatePublicip 绑定弹性公网IP
//
// 绑定弹性公网IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) UpdateAssociatePublicip(request *model.UpdateAssociatePublicipRequest) (*model.UpdateAssociatePublicipResponse, error) {
	requestDef := GenReqDefForUpdateAssociatePublicip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAssociatePublicipResponse), nil
	}
}

// UpdateAssociatePublicipInvoker 绑定弹性公网IP
func (c *EipClient) UpdateAssociatePublicipInvoker(request *model.UpdateAssociatePublicipRequest) *UpdateAssociatePublicipInvoker {
	requestDef := GenReqDefForUpdateAssociatePublicip()
	return &UpdateAssociatePublicipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDisassociatePublicip 解绑弹性公网IP
//
// 解绑弹性公网IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EipClient) UpdateDisassociatePublicip(request *model.UpdateDisassociatePublicipRequest) (*model.UpdateDisassociatePublicipResponse, error) {
	requestDef := GenReqDefForUpdateDisassociatePublicip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDisassociatePublicipResponse), nil
	}
}

// UpdateDisassociatePublicipInvoker 解绑弹性公网IP
func (c *EipClient) UpdateDisassociatePublicipInvoker(request *model.UpdateDisassociatePublicipRequest) *UpdateDisassociatePublicipInvoker {
	requestDef := GenReqDefForUpdateDisassociatePublicip()
	return &UpdateDisassociatePublicipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
