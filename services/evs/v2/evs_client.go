package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/evs/v2/model"
)

type EvsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewEvsClient(hcClient *httpclient.HcHttpClient) *EvsClient {
	return &EvsClient{HcClient: hcClient}
}

func EvsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AssociateAutoSnapshotPolicy 绑定自动快照策略
//
// 绑定自动快照策略。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) AssociateAutoSnapshotPolicy(request *model.AssociateAutoSnapshotPolicyRequest) (*model.AssociateAutoSnapshotPolicyResponse, error) {
	requestDef := GenReqDefForAssociateAutoSnapshotPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateAutoSnapshotPolicyResponse), nil
	}
}

// AssociateAutoSnapshotPolicyInvoker 绑定自动快照策略
func (c *EvsClient) AssociateAutoSnapshotPolicyInvoker(request *model.AssociateAutoSnapshotPolicyRequest) *AssociateAutoSnapshotPolicyInvoker {
	requestDef := GenReqDefForAssociateAutoSnapshotPolicy()
	return &AssociateAutoSnapshotPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateSnapshotGroupTag 为指定快照组添加标签
//
// 为指定快照组添加标签。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) BatchCreateSnapshotGroupTag(request *model.BatchCreateSnapshotGroupTagRequest) (*model.BatchCreateSnapshotGroupTagResponse, error) {
	requestDef := GenReqDefForBatchCreateSnapshotGroupTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateSnapshotGroupTagResponse), nil
	}
}

// BatchCreateSnapshotGroupTagInvoker 为指定快照组添加标签
func (c *EvsClient) BatchCreateSnapshotGroupTagInvoker(request *model.BatchCreateSnapshotGroupTagRequest) *BatchCreateSnapshotGroupTagInvoker {
	requestDef := GenReqDefForBatchCreateSnapshotGroupTag()
	return &BatchCreateSnapshotGroupTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateSnapshotTags 为指定快照添加标签
//
// 为指定快照添加标签。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) BatchCreateSnapshotTags(request *model.BatchCreateSnapshotTagsRequest) (*model.BatchCreateSnapshotTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateSnapshotTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateSnapshotTagsResponse), nil
	}
}

// BatchCreateSnapshotTagsInvoker 为指定快照添加标签
func (c *EvsClient) BatchCreateSnapshotTagsInvoker(request *model.BatchCreateSnapshotTagsRequest) *BatchCreateSnapshotTagsInvoker {
	requestDef := GenReqDefForBatchCreateSnapshotTags()
	return &BatchCreateSnapshotTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateVolumeTags 为指定云硬盘批量添加标签
//
// 为指定云硬盘批量添加标签。
//
// 添加标签时，如果云硬盘的标签已存在相同key，则会覆盖已有标签。
// 单个云硬盘最多支持创建10个标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) BatchCreateVolumeTags(request *model.BatchCreateVolumeTagsRequest) (*model.BatchCreateVolumeTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateVolumeTagsResponse), nil
	}
}

// BatchCreateVolumeTagsInvoker 为指定云硬盘批量添加标签
func (c *EvsClient) BatchCreateVolumeTagsInvoker(request *model.BatchCreateVolumeTagsRequest) *BatchCreateVolumeTagsInvoker {
	requestDef := GenReqDefForBatchCreateVolumeTags()
	return &BatchCreateVolumeTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteSnapshotGroupTag 为指定快照组删除标签
//
// 为指定快照组删除标签。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) BatchDeleteSnapshotGroupTag(request *model.BatchDeleteSnapshotGroupTagRequest) (*model.BatchDeleteSnapshotGroupTagResponse, error) {
	requestDef := GenReqDefForBatchDeleteSnapshotGroupTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteSnapshotGroupTagResponse), nil
	}
}

// BatchDeleteSnapshotGroupTagInvoker 为指定快照组删除标签
func (c *EvsClient) BatchDeleteSnapshotGroupTagInvoker(request *model.BatchDeleteSnapshotGroupTagRequest) *BatchDeleteSnapshotGroupTagInvoker {
	requestDef := GenReqDefForBatchDeleteSnapshotGroupTag()
	return &BatchDeleteSnapshotGroupTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteSnapshotTags 为指定快照删除标签
//
// 为指定快照删除标签。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) BatchDeleteSnapshotTags(request *model.BatchDeleteSnapshotTagsRequest) (*model.BatchDeleteSnapshotTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteSnapshotTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteSnapshotTagsResponse), nil
	}
}

// BatchDeleteSnapshotTagsInvoker 为指定快照删除标签
func (c *EvsClient) BatchDeleteSnapshotTagsInvoker(request *model.BatchDeleteSnapshotTagsRequest) *BatchDeleteSnapshotTagsInvoker {
	requestDef := GenReqDefForBatchDeleteSnapshotTags()
	return &BatchDeleteSnapshotTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteVolumeTags 为指定云硬盘批量删除标签
//
// 为指定云硬盘批量删除标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) BatchDeleteVolumeTags(request *model.BatchDeleteVolumeTagsRequest) (*model.BatchDeleteVolumeTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteVolumeTagsResponse), nil
	}
}

// BatchDeleteVolumeTagsInvoker 为指定云硬盘批量删除标签
func (c *EvsClient) BatchDeleteVolumeTagsInvoker(request *model.BatchDeleteVolumeTagsRequest) *BatchDeleteVolumeTagsInvoker {
	requestDef := GenReqDefForBatchDeleteVolumeTags()
	return &BatchDeleteVolumeTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchResizeVolumes 批量扩容云硬盘
//
// 对按需或者包周期云硬盘进行批量扩容。
// [在批量扩容存在包周期云硬盘的场景下：](tag:hws)
// - [如果您需要查看订单可用的优惠券，请参考\&quot;[查询订单可用优惠券](https://support.huaweicloud.com/api-oce/zh-cn_topic_0092953630.html)\&quot;。](tag:hws)
// - [如果您需要支付订单，请参考\&quot;[支付包周期产品订单](https://support.huaweicloud.com/api-oce/api_order_00030.html)\&quot;。](tag:hws)
// - [如果您需要退订该包周期资源，请参考“[退订包周期资源](https://support.huaweicloud.com/api-oce/zh-cn_topic_0082522030.html)”。](tag:hws)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) BatchResizeVolumes(request *model.BatchResizeVolumesRequest) (*model.BatchResizeVolumesResponse, error) {
	requestDef := GenReqDefForBatchResizeVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchResizeVolumesResponse), nil
	}
}

// BatchResizeVolumesInvoker 批量扩容云硬盘
func (c *EvsClient) BatchResizeVolumesInvoker(request *model.BatchResizeVolumesRequest) *BatchResizeVolumesInvoker {
	requestDef := GenReqDefForBatchResizeVolumes()
	return &BatchResizeVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeVolumeChargeMode 修改云硬盘计费模式
//
// 将挂载状态下的云硬盘的计费模式有按需转成包周期，且到期时间和挂载的虚拟机保持一致。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ChangeVolumeChargeMode(request *model.ChangeVolumeChargeModeRequest) (*model.ChangeVolumeChargeModeResponse, error) {
	requestDef := GenReqDefForChangeVolumeChargeMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeVolumeChargeModeResponse), nil
	}
}

// ChangeVolumeChargeModeInvoker 修改云硬盘计费模式
func (c *EvsClient) ChangeVolumeChargeModeInvoker(request *model.ChangeVolumeChargeModeRequest) *ChangeVolumeChargeModeInvoker {
	requestDef := GenReqDefForChangeVolumeChargeMode()
	return &ChangeVolumeChargeModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CinderAcceptVolumeTransfer 接受云硬盘过户
//
// 通过云硬盘过户记录ID以及身份认证密钥来接受云硬盘过户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CinderAcceptVolumeTransfer(request *model.CinderAcceptVolumeTransferRequest) (*model.CinderAcceptVolumeTransferResponse, error) {
	requestDef := GenReqDefForCinderAcceptVolumeTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderAcceptVolumeTransferResponse), nil
	}
}

// CinderAcceptVolumeTransferInvoker 接受云硬盘过户
func (c *EvsClient) CinderAcceptVolumeTransferInvoker(request *model.CinderAcceptVolumeTransferRequest) *CinderAcceptVolumeTransferInvoker {
	requestDef := GenReqDefForCinderAcceptVolumeTransfer()
	return &CinderAcceptVolumeTransferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CinderCreateVolumeTransfer 创建云硬盘过户
//
// 指定云硬盘来创建云硬盘过户记录，创建成功后，会返回过户记录ID以及身份认证密钥。
// 云硬盘在过户过程中的状态变化如下：创建云硬盘过户后，云硬盘状态由“available”变为“awaiting-transfer”。当云硬盘过户被接收后，云硬盘状态变为“available”。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CinderCreateVolumeTransfer(request *model.CinderCreateVolumeTransferRequest) (*model.CinderCreateVolumeTransferResponse, error) {
	requestDef := GenReqDefForCinderCreateVolumeTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderCreateVolumeTransferResponse), nil
	}
}

// CinderCreateVolumeTransferInvoker 创建云硬盘过户
func (c *EvsClient) CinderCreateVolumeTransferInvoker(request *model.CinderCreateVolumeTransferRequest) *CinderCreateVolumeTransferInvoker {
	requestDef := GenReqDefForCinderCreateVolumeTransfer()
	return &CinderCreateVolumeTransferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CinderDeleteVolumeTransfer 删除云硬盘过户
//
// 当云硬盘过户未被接受时，您可以删除云硬盘过户记录，接受后则无法执行删除操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CinderDeleteVolumeTransfer(request *model.CinderDeleteVolumeTransferRequest) (*model.CinderDeleteVolumeTransferResponse, error) {
	requestDef := GenReqDefForCinderDeleteVolumeTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderDeleteVolumeTransferResponse), nil
	}
}

// CinderDeleteVolumeTransferInvoker 删除云硬盘过户
func (c *EvsClient) CinderDeleteVolumeTransferInvoker(request *model.CinderDeleteVolumeTransferRequest) *CinderDeleteVolumeTransferInvoker {
	requestDef := GenReqDefForCinderDeleteVolumeTransfer()
	return &CinderDeleteVolumeTransferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CinderListAvailabilityZones 查询所有的可用分区信息
//
// 查询所有的可用分区信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CinderListAvailabilityZones(request *model.CinderListAvailabilityZonesRequest) (*model.CinderListAvailabilityZonesResponse, error) {
	requestDef := GenReqDefForCinderListAvailabilityZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderListAvailabilityZonesResponse), nil
	}
}

// CinderListAvailabilityZonesInvoker 查询所有的可用分区信息
func (c *EvsClient) CinderListAvailabilityZonesInvoker(request *model.CinderListAvailabilityZonesRequest) *CinderListAvailabilityZonesInvoker {
	requestDef := GenReqDefForCinderListAvailabilityZones()
	return &CinderListAvailabilityZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CinderListQuotas 查询租户的详细配额
//
// 查询租户的详细配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CinderListQuotas(request *model.CinderListQuotasRequest) (*model.CinderListQuotasResponse, error) {
	requestDef := GenReqDefForCinderListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderListQuotasResponse), nil
	}
}

// CinderListQuotasInvoker 查询租户的详细配额
func (c *EvsClient) CinderListQuotasInvoker(request *model.CinderListQuotasRequest) *CinderListQuotasInvoker {
	requestDef := GenReqDefForCinderListQuotas()
	return &CinderListQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CinderListVolumeTransfers 查询云硬盘过户记录列表概要
//
// 查询当前租户下所有云硬盘的过户记录列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CinderListVolumeTransfers(request *model.CinderListVolumeTransfersRequest) (*model.CinderListVolumeTransfersResponse, error) {
	requestDef := GenReqDefForCinderListVolumeTransfers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderListVolumeTransfersResponse), nil
	}
}

// CinderListVolumeTransfersInvoker 查询云硬盘过户记录列表概要
func (c *EvsClient) CinderListVolumeTransfersInvoker(request *model.CinderListVolumeTransfersRequest) *CinderListVolumeTransfersInvoker {
	requestDef := GenReqDefForCinderListVolumeTransfers()
	return &CinderListVolumeTransfersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CinderListVolumeTypes 查询云硬盘类型列表
//
// 查询云硬盘类型列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CinderListVolumeTypes(request *model.CinderListVolumeTypesRequest) (*model.CinderListVolumeTypesResponse, error) {
	requestDef := GenReqDefForCinderListVolumeTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderListVolumeTypesResponse), nil
	}
}

// CinderListVolumeTypesInvoker 查询云硬盘类型列表
func (c *EvsClient) CinderListVolumeTypesInvoker(request *model.CinderListVolumeTypesRequest) *CinderListVolumeTypesInvoker {
	requestDef := GenReqDefForCinderListVolumeTypes()
	return &CinderListVolumeTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CinderShowVolumeTransfer 查询单个云硬盘过户记录详情
//
// 查询单个云硬盘的过户记录详情，比如过户记录创建时间、ID以及名称等信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CinderShowVolumeTransfer(request *model.CinderShowVolumeTransferRequest) (*model.CinderShowVolumeTransferResponse, error) {
	requestDef := GenReqDefForCinderShowVolumeTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderShowVolumeTransferResponse), nil
	}
}

// CinderShowVolumeTransferInvoker 查询单个云硬盘过户记录详情
func (c *EvsClient) CinderShowVolumeTransferInvoker(request *model.CinderShowVolumeTransferRequest) *CinderShowVolumeTransferInvoker {
	requestDef := GenReqDefForCinderShowVolumeTransfer()
	return &CinderShowVolumeTransferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountAutoSnapshotPolicies 查询自动快照策略总数
//
// 查询自动快照策略总数。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CountAutoSnapshotPolicies(request *model.CountAutoSnapshotPoliciesRequest) (*model.CountAutoSnapshotPoliciesResponse, error) {
	requestDef := GenReqDefForCountAutoSnapshotPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountAutoSnapshotPoliciesResponse), nil
	}
}

// CountAutoSnapshotPoliciesInvoker 查询自动快照策略总数
func (c *EvsClient) CountAutoSnapshotPoliciesInvoker(request *model.CountAutoSnapshotPoliciesRequest) *CountAutoSnapshotPoliciesInvoker {
	requestDef := GenReqDefForCountAutoSnapshotPolicies()
	return &CountAutoSnapshotPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountSnapshotChains 查询快照链个数
//
// 查询快照链个数。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CountSnapshotChains(request *model.CountSnapshotChainsRequest) (*model.CountSnapshotChainsResponse, error) {
	requestDef := GenReqDefForCountSnapshotChains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountSnapshotChainsResponse), nil
	}
}

// CountSnapshotChainsInvoker 查询快照链个数
func (c *EvsClient) CountSnapshotChainsInvoker(request *model.CountSnapshotChainsRequest) *CountSnapshotChainsInvoker {
	requestDef := GenReqDefForCountSnapshotChains()
	return &CountSnapshotChainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountSnapshots 查询快照个数
//
// 查询快照个数。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CountSnapshots(request *model.CountSnapshotsRequest) (*model.CountSnapshotsResponse, error) {
	requestDef := GenReqDefForCountSnapshots()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountSnapshotsResponse), nil
	}
}

// CountSnapshotsInvoker 查询快照个数
func (c *EvsClient) CountSnapshotsInvoker(request *model.CountSnapshotsRequest) *CountSnapshotsInvoker {
	requestDef := GenReqDefForCountSnapshots()
	return &CountSnapshotsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountSnapshotsGroup 查询快照一致性组个数
//
// 查询快照一致性组个数。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CountSnapshotsGroup(request *model.CountSnapshotsGroupRequest) (*model.CountSnapshotsGroupResponse, error) {
	requestDef := GenReqDefForCountSnapshotsGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountSnapshotsGroupResponse), nil
	}
}

// CountSnapshotsGroupInvoker 查询快照一致性组个数
func (c *EvsClient) CountSnapshotsGroupInvoker(request *model.CountSnapshotsGroupRequest) *CountSnapshotsGroupInvoker {
	requestDef := GenReqDefForCountSnapshotsGroup()
	return &CountSnapshotsGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAutoSnapshotPolicy 创建自动快照策略
//
// 创建自动快照策略。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CreateAutoSnapshotPolicy(request *model.CreateAutoSnapshotPolicyRequest) (*model.CreateAutoSnapshotPolicyResponse, error) {
	requestDef := GenReqDefForCreateAutoSnapshotPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAutoSnapshotPolicyResponse), nil
	}
}

// CreateAutoSnapshotPolicyInvoker 创建自动快照策略
func (c *EvsClient) CreateAutoSnapshotPolicyInvoker(request *model.CreateAutoSnapshotPolicyRequest) *CreateAutoSnapshotPolicyInvoker {
	requestDef := GenReqDefForCreateAutoSnapshotPolicy()
	return &CreateAutoSnapshotPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSnapshot 创建云硬盘快照
//
// 创建云硬盘快照。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CreateSnapshot(request *model.CreateSnapshotRequest) (*model.CreateSnapshotResponse, error) {
	requestDef := GenReqDefForCreateSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSnapshotResponse), nil
	}
}

// CreateSnapshotInvoker 创建云硬盘快照
func (c *EvsClient) CreateSnapshotInvoker(request *model.CreateSnapshotRequest) *CreateSnapshotInvoker {
	requestDef := GenReqDefForCreateSnapshot()
	return &CreateSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSnapshotGroup 创建快照一致性组
//
// 创建快照一致性组。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CreateSnapshotGroup(request *model.CreateSnapshotGroupRequest) (*model.CreateSnapshotGroupResponse, error) {
	requestDef := GenReqDefForCreateSnapshotGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSnapshotGroupResponse), nil
	}
}

// CreateSnapshotGroupInvoker 创建快照一致性组
func (c *EvsClient) CreateSnapshotGroupInvoker(request *model.CreateSnapshotGroupRequest) *CreateSnapshotGroupInvoker {
	requestDef := GenReqDefForCreateSnapshotGroup()
	return &CreateSnapshotGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStandardSnapshot 创建快照
//
// 创建标准快照。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CreateStandardSnapshot(request *model.CreateStandardSnapshotRequest) (*model.CreateStandardSnapshotResponse, error) {
	requestDef := GenReqDefForCreateStandardSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStandardSnapshotResponse), nil
	}
}

// CreateStandardSnapshotInvoker 创建快照
func (c *EvsClient) CreateStandardSnapshotInvoker(request *model.CreateStandardSnapshotRequest) *CreateStandardSnapshotInvoker {
	requestDef := GenReqDefForCreateStandardSnapshot()
	return &CreateStandardSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVolume 创建云硬盘
//
// 创建按需或包周期云硬盘。
// 在创建包周期云硬盘的场景下：
// - 如果您需要查看订单可用的优惠券，请参考\&quot;[查询订单可用优惠券](https://support.huaweicloud.com/api-oce/zh-cn_topic_0092953630.html)\&quot;。
// - 如果您需要支付订单，请参考\&quot;[支付包周期产品订单](https://support.huaweicloud.com/api-oce/zh-cn_topic_0075746561.html)\&quot;。
// - 如果您需要查询订单的资源开通详情，请参考\&quot;[查询订单的资源开通详情](https://support.huaweicloud.com/api-oce/api_order_00001.html)\&quot;。
// - 如果您需要退订该包周期资源，请参考“[退订包周期资源](https://support.huaweicloud.com/api-oce/zh-cn_topic_0082522030.html)”。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) CreateVolume(request *model.CreateVolumeRequest) (*model.CreateVolumeResponse, error) {
	requestDef := GenReqDefForCreateVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVolumeResponse), nil
	}
}

// CreateVolumeInvoker 创建云硬盘
func (c *EvsClient) CreateVolumeInvoker(request *model.CreateVolumeRequest) *CreateVolumeInvoker {
	requestDef := GenReqDefForCreateVolume()
	return &CreateVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAutoSnapshotPolicy 删除单个自动快照策略
//
// 删除单个自动快照策略。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) DeleteAutoSnapshotPolicy(request *model.DeleteAutoSnapshotPolicyRequest) (*model.DeleteAutoSnapshotPolicyResponse, error) {
	requestDef := GenReqDefForDeleteAutoSnapshotPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAutoSnapshotPolicyResponse), nil
	}
}

// DeleteAutoSnapshotPolicyInvoker 删除单个自动快照策略
func (c *EvsClient) DeleteAutoSnapshotPolicyInvoker(request *model.DeleteAutoSnapshotPolicyRequest) *DeleteAutoSnapshotPolicyInvoker {
	requestDef := GenReqDefForDeleteAutoSnapshotPolicy()
	return &DeleteAutoSnapshotPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSnapshot 删除云硬盘快照
//
// 删除云硬盘快照。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) DeleteSnapshot(request *model.DeleteSnapshotRequest) (*model.DeleteSnapshotResponse, error) {
	requestDef := GenReqDefForDeleteSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSnapshotResponse), nil
	}
}

// DeleteSnapshotInvoker 删除云硬盘快照
func (c *EvsClient) DeleteSnapshotInvoker(request *model.DeleteSnapshotRequest) *DeleteSnapshotInvoker {
	requestDef := GenReqDefForDeleteSnapshot()
	return &DeleteSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSnapshotGroup 删除快照一致性组
//
// 删除快照一致性组。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) DeleteSnapshotGroup(request *model.DeleteSnapshotGroupRequest) (*model.DeleteSnapshotGroupResponse, error) {
	requestDef := GenReqDefForDeleteSnapshotGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSnapshotGroupResponse), nil
	}
}

// DeleteSnapshotGroupInvoker 删除快照一致性组
func (c *EvsClient) DeleteSnapshotGroupInvoker(request *model.DeleteSnapshotGroupRequest) *DeleteSnapshotGroupInvoker {
	requestDef := GenReqDefForDeleteSnapshotGroup()
	return &DeleteSnapshotGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteStandardSnapshot 删除快照
//
// 删除快照。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) DeleteStandardSnapshot(request *model.DeleteStandardSnapshotRequest) (*model.DeleteStandardSnapshotResponse, error) {
	requestDef := GenReqDefForDeleteStandardSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStandardSnapshotResponse), nil
	}
}

// DeleteStandardSnapshotInvoker 删除快照
func (c *EvsClient) DeleteStandardSnapshotInvoker(request *model.DeleteStandardSnapshotRequest) *DeleteStandardSnapshotInvoker {
	requestDef := GenReqDefForDeleteStandardSnapshot()
	return &DeleteStandardSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVolume 删除云硬盘
//
// 删除一个云硬盘。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) DeleteVolume(request *model.DeleteVolumeRequest) (*model.DeleteVolumeResponse, error) {
	requestDef := GenReqDefForDeleteVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVolumeResponse), nil
	}
}

// DeleteVolumeInvoker 删除云硬盘
func (c *EvsClient) DeleteVolumeInvoker(request *model.DeleteVolumeRequest) *DeleteVolumeInvoker {
	requestDef := GenReqDefForDeleteVolume()
	return &DeleteVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVolumeInRecycle 删除回收站中单个云硬盘
//
// 删除回收站中单个云硬盘。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) DeleteVolumeInRecycle(request *model.DeleteVolumeInRecycleRequest) (*model.DeleteVolumeInRecycleResponse, error) {
	requestDef := GenReqDefForDeleteVolumeInRecycle()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVolumeInRecycleResponse), nil
	}
}

// DeleteVolumeInRecycleInvoker 删除回收站中单个云硬盘
func (c *EvsClient) DeleteVolumeInRecycleInvoker(request *model.DeleteVolumeInRecycleRequest) *DeleteVolumeInRecycleInvoker {
	requestDef := GenReqDefForDeleteVolumeInRecycle()
	return &DeleteVolumeInRecycleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateAutoSnapshotPolicy 解绑自动快照策略
//
// 解绑自动快照策略。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) DisassociateAutoSnapshotPolicy(request *model.DisassociateAutoSnapshotPolicyRequest) (*model.DisassociateAutoSnapshotPolicyResponse, error) {
	requestDef := GenReqDefForDisassociateAutoSnapshotPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateAutoSnapshotPolicyResponse), nil
	}
}

// DisassociateAutoSnapshotPolicyInvoker 解绑自动快照策略
func (c *EvsClient) DisassociateAutoSnapshotPolicyInvoker(request *model.DisassociateAutoSnapshotPolicyRequest) *DisassociateAutoSnapshotPolicyInvoker {
	requestDef := GenReqDefForDisassociateAutoSnapshotPolicy()
	return &DisassociateAutoSnapshotPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAutoSnapshotPolicy 查询自动快照策略
//
// 查询自动快照策略。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ListAutoSnapshotPolicy(request *model.ListAutoSnapshotPolicyRequest) (*model.ListAutoSnapshotPolicyResponse, error) {
	requestDef := GenReqDefForListAutoSnapshotPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAutoSnapshotPolicyResponse), nil
	}
}

// ListAutoSnapshotPolicyInvoker 查询自动快照策略
func (c *EvsClient) ListAutoSnapshotPolicyInvoker(request *model.ListAutoSnapshotPolicyRequest) *ListAutoSnapshotPolicyInvoker {
	requestDef := GenReqDefForListAutoSnapshotPolicy()
	return &ListAutoSnapshotPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSnapshotChains 查询快照链列表
//
// 查询快照链的列表。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ListSnapshotChains(request *model.ListSnapshotChainsRequest) (*model.ListSnapshotChainsResponse, error) {
	requestDef := GenReqDefForListSnapshotChains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotChainsResponse), nil
	}
}

// ListSnapshotChainsInvoker 查询快照链列表
func (c *EvsClient) ListSnapshotChainsInvoker(request *model.ListSnapshotChainsRequest) *ListSnapshotChainsInvoker {
	requestDef := GenReqDefForListSnapshotChains()
	return &ListSnapshotChainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSnapshotGroup 查询快照一致性组列表
//
// 查询快照一致性组列表。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ListSnapshotGroup(request *model.ListSnapshotGroupRequest) (*model.ListSnapshotGroupResponse, error) {
	requestDef := GenReqDefForListSnapshotGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotGroupResponse), nil
	}
}

// ListSnapshotGroupInvoker 查询快照一致性组列表
func (c *EvsClient) ListSnapshotGroupInvoker(request *model.ListSnapshotGroupRequest) *ListSnapshotGroupInvoker {
	requestDef := GenReqDefForListSnapshotGroup()
	return &ListSnapshotGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSnapshotGroupTags 查询快照组的所有标签
//
// 查询快照组的所有标签。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ListSnapshotGroupTags(request *model.ListSnapshotGroupTagsRequest) (*model.ListSnapshotGroupTagsResponse, error) {
	requestDef := GenReqDefForListSnapshotGroupTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotGroupTagsResponse), nil
	}
}

// ListSnapshotGroupTagsInvoker 查询快照组的所有标签
func (c *EvsClient) ListSnapshotGroupTagsInvoker(request *model.ListSnapshotGroupTagsRequest) *ListSnapshotGroupTagsInvoker {
	requestDef := GenReqDefForListSnapshotGroupTags()
	return &ListSnapshotGroupTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSnapshotTags 查询快照资源的所有标签
//
// 查询快照资源的所有标签。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ListSnapshotTags(request *model.ListSnapshotTagsRequest) (*model.ListSnapshotTagsResponse, error) {
	requestDef := GenReqDefForListSnapshotTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotTagsResponse), nil
	}
}

// ListSnapshotTagsInvoker 查询快照资源的所有标签
func (c *EvsClient) ListSnapshotTagsInvoker(request *model.ListSnapshotTagsRequest) *ListSnapshotTagsInvoker {
	requestDef := GenReqDefForListSnapshotTags()
	return &ListSnapshotTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSnapshots 查询云硬盘快照详情列表
//
// 查询云硬盘快照详细列表信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ListSnapshots(request *model.ListSnapshotsRequest) (*model.ListSnapshotsResponse, error) {
	requestDef := GenReqDefForListSnapshots()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotsResponse), nil
	}
}

// ListSnapshotsInvoker 查询云硬盘快照详情列表
func (c *EvsClient) ListSnapshotsInvoker(request *model.ListSnapshotsRequest) *ListSnapshotsInvoker {
	requestDef := GenReqDefForListSnapshots()
	return &ListSnapshotsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSnapshotsDetails 查询快照详情列表
//
// 查询快照详情列表。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ListSnapshotsDetails(request *model.ListSnapshotsDetailsRequest) (*model.ListSnapshotsDetailsResponse, error) {
	requestDef := GenReqDefForListSnapshotsDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotsDetailsResponse), nil
	}
}

// ListSnapshotsDetailsInvoker 查询快照详情列表
func (c *EvsClient) ListSnapshotsDetailsInvoker(request *model.ListSnapshotsDetailsRequest) *ListSnapshotsDetailsInvoker {
	requestDef := GenReqDefForListSnapshotsDetails()
	return &ListSnapshotsDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVolumeTags 获取云硬盘资源的所有标签
//
// 获取某个租户的所有云硬盘资源的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ListVolumeTags(request *model.ListVolumeTagsRequest) (*model.ListVolumeTagsResponse, error) {
	requestDef := GenReqDefForListVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumeTagsResponse), nil
	}
}

// ListVolumeTagsInvoker 获取云硬盘资源的所有标签
func (c *EvsClient) ListVolumeTagsInvoker(request *model.ListVolumeTagsRequest) *ListVolumeTagsInvoker {
	requestDef := GenReqDefForListVolumeTags()
	return &ListVolumeTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVolumes 查询所有云硬盘详情
//
// 查询所有云硬盘的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ListVolumes(request *model.ListVolumesRequest) (*model.ListVolumesResponse, error) {
	requestDef := GenReqDefForListVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumesResponse), nil
	}
}

// ListVolumesInvoker 查询所有云硬盘详情
func (c *EvsClient) ListVolumesInvoker(request *model.ListVolumesRequest) *ListVolumesInvoker {
	requestDef := GenReqDefForListVolumes()
	return &ListVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVolumesByTags 通过标签查询云硬盘资源实例详情
//
// 通过标签查询云硬盘资源实例详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ListVolumesByTags(request *model.ListVolumesByTagsRequest) (*model.ListVolumesByTagsResponse, error) {
	requestDef := GenReqDefForListVolumesByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumesByTagsResponse), nil
	}
}

// ListVolumesByTagsInvoker 通过标签查询云硬盘资源实例详情
func (c *EvsClient) ListVolumesByTagsInvoker(request *model.ListVolumesByTagsRequest) *ListVolumesByTagsInvoker {
	requestDef := GenReqDefForListVolumesByTags()
	return &ListVolumesByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVolumesInRecycle 查询回收站中所有云硬盘详情
//
// 查询回收站中所有云硬盘的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ListVolumesInRecycle(request *model.ListVolumesInRecycleRequest) (*model.ListVolumesInRecycleResponse, error) {
	requestDef := GenReqDefForListVolumesInRecycle()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumesInRecycleResponse), nil
	}
}

// ListVolumesInRecycleInvoker 查询回收站中所有云硬盘详情
func (c *EvsClient) ListVolumesInRecycleInvoker(request *model.ListVolumesInRecycleRequest) *ListVolumesInRecycleInvoker {
	requestDef := GenReqDefForListVolumesInRecycle()
	return &ListVolumesInRecycleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyVolumeQoS 修改云硬盘QoS
//
// 调整云硬盘的iops或者吞吐量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ModifyVolumeQoS(request *model.ModifyVolumeQoSRequest) (*model.ModifyVolumeQoSResponse, error) {
	requestDef := GenReqDefForModifyVolumeQoS()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyVolumeQoSResponse), nil
	}
}

// ModifyVolumeQoSInvoker 修改云硬盘QoS
func (c *EvsClient) ModifyVolumeQoSInvoker(request *model.ModifyVolumeQoSRequest) *ModifyVolumeQoSInvoker {
	requestDef := GenReqDefForModifyVolumeQoS()
	return &ModifyVolumeQoSInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeVolume 扩容云硬盘
//
// 对按需或者包周期云硬盘进行扩容。
// 在扩容包周期云硬盘的场景下：
// - 如果您需要查看订单可用的优惠券，请参考\&quot;[查询订单可用优惠券](https://support.huaweicloud.com/api-oce/zh-cn_topic_0092953630.html)\&quot;。
// - 如果您需要支付订单，请参考\&quot;[支付包周期产品订单](https://support.huaweicloud.com/api-oce/zh-cn_topic_0075746561.html)\&quot;。
// - 如果您需要查询订单的资源开通详情，请参考\&quot;[查询订单的资源开通详情](https://support.huaweicloud.com/api-oce/api_order_00001.html)\&quot;。
// - 如果您需要退订该包周期资源，请参考“[退订包周期资源](https://support.huaweicloud.com/api-oce/zh-cn_topic_0082522030.html)”。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ResizeVolume(request *model.ResizeVolumeRequest) (*model.ResizeVolumeResponse, error) {
	requestDef := GenReqDefForResizeVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeVolumeResponse), nil
	}
}

// ResizeVolumeInvoker 扩容云硬盘
func (c *EvsClient) ResizeVolumeInvoker(request *model.ResizeVolumeRequest) *ResizeVolumeInvoker {
	requestDef := GenReqDefForResizeVolume()
	return &ResizeVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetypeVolume 磁盘类型变更
//
// 对按需或者包周期云硬盘进行磁盘类型变更。
// [在磁盘类型变更包周期云硬盘的场景下：](tag:hws)
// - [如果您需要查看订单可用的优惠券，请参考\&quot;[查询订单可用优惠券](https://support.huaweicloud.com/api-oce/zh-cn_topic_0092953630.html)\&quot;。](tag:hws)
// - [如果您需要支付订单，请参考\&quot;[支付包周期产品订单](https://support.huaweicloud.com/api-oce/api_order_00030.html)\&quot;。](tag:hws)
// - [如果您需要查询订单的资源开通详情，请参考\&quot;[查询订单的资源开通详情](https://support.huaweicloud.com/api-oce/api_order_00001.html)\&quot;。](tag:hws)
// - [如果您需要退订该包周期资源，请参考“[退订包周期资源](https://support.huaweicloud.com/api-oce/zh-cn_topic_0082522030.html)”。](tag:hws)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) RetypeVolume(request *model.RetypeVolumeRequest) (*model.RetypeVolumeResponse, error) {
	requestDef := GenReqDefForRetypeVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetypeVolumeResponse), nil
	}
}

// RetypeVolumeInvoker 磁盘类型变更
func (c *EvsClient) RetypeVolumeInvoker(request *model.RetypeVolumeRequest) *RetypeVolumeInvoker {
	requestDef := GenReqDefForRetypeVolume()
	return &RetypeVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RevertVolumeInRecycle 还原回收站中单个云硬盘
//
// 还原回收站中单个云硬盘。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) RevertVolumeInRecycle(request *model.RevertVolumeInRecycleRequest) (*model.RevertVolumeInRecycleResponse, error) {
	requestDef := GenReqDefForRevertVolumeInRecycle()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RevertVolumeInRecycleResponse), nil
	}
}

// RevertVolumeInRecycleInvoker 还原回收站中单个云硬盘
func (c *EvsClient) RevertVolumeInRecycleInvoker(request *model.RevertVolumeInRecycleRequest) *RevertVolumeInRecycleInvoker {
	requestDef := GenReqDefForRevertVolumeInRecycle()
	return &RevertVolumeInRecycleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RollbackSnapshot 回滚快照到云硬盘
//
// 将快照数据回滚到云硬盘。支持企业项目授权功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) RollbackSnapshot(request *model.RollbackSnapshotRequest) (*model.RollbackSnapshotResponse, error) {
	requestDef := GenReqDefForRollbackSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RollbackSnapshotResponse), nil
	}
}

// RollbackSnapshotInvoker 回滚快照到云硬盘
func (c *EvsClient) RollbackSnapshotInvoker(request *model.RollbackSnapshotRequest) *RollbackSnapshotInvoker {
	requestDef := GenReqDefForRollbackSnapshot()
	return &RollbackSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RollbackStandardSnapshot 快照回滚
//
// 将快照数据回滚到云硬盘。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) RollbackStandardSnapshot(request *model.RollbackStandardSnapshotRequest) (*model.RollbackStandardSnapshotResponse, error) {
	requestDef := GenReqDefForRollbackStandardSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RollbackStandardSnapshotResponse), nil
	}
}

// RollbackStandardSnapshotInvoker 快照回滚
func (c *EvsClient) RollbackStandardSnapshotInvoker(request *model.RollbackStandardSnapshotRequest) *RollbackStandardSnapshotInvoker {
	requestDef := GenReqDefForRollbackStandardSnapshot()
	return &RollbackStandardSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAutoSnapshotPolicy 查询单个自动快照策略
//
// 查询单个自动快照策略。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ShowAutoSnapshotPolicy(request *model.ShowAutoSnapshotPolicyRequest) (*model.ShowAutoSnapshotPolicyResponse, error) {
	requestDef := GenReqDefForShowAutoSnapshotPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoSnapshotPolicyResponse), nil
	}
}

// ShowAutoSnapshotPolicyInvoker 查询单个自动快照策略
func (c *EvsClient) ShowAutoSnapshotPolicyInvoker(request *model.ShowAutoSnapshotPolicyRequest) *ShowAutoSnapshotPolicyInvoker {
	requestDef := GenReqDefForShowAutoSnapshotPolicy()
	return &ShowAutoSnapshotPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJob 查询job的状态
//
// 查询Job的执行状态。
// 可用于查询创建云硬盘，扩容云硬盘，删除云硬盘等API的执行状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

// ShowJobInvoker 查询job的状态
func (c *EvsClient) ShowJobInvoker(request *model.ShowJobRequest) *ShowJobInvoker {
	requestDef := GenReqDefForShowJob()
	return &ShowJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRecyclePolicy 查询回收站策略
//
// 查询回收站策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ShowRecyclePolicy(request *model.ShowRecyclePolicyRequest) (*model.ShowRecyclePolicyResponse, error) {
	requestDef := GenReqDefForShowRecyclePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecyclePolicyResponse), nil
	}
}

// ShowRecyclePolicyInvoker 查询回收站策略
func (c *EvsClient) ShowRecyclePolicyInvoker(request *model.ShowRecyclePolicyRequest) *ShowRecyclePolicyInvoker {
	requestDef := GenReqDefForShowRecyclePolicy()
	return &ShowRecyclePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSnapshot 查询单个云硬盘快照详情
//
// 查询单个云硬盘快照信息。支持企业项目授权功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ShowSnapshot(request *model.ShowSnapshotRequest) (*model.ShowSnapshotResponse, error) {
	requestDef := GenReqDefForShowSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSnapshotResponse), nil
	}
}

// ShowSnapshotInvoker 查询单个云硬盘快照详情
func (c *EvsClient) ShowSnapshotInvoker(request *model.ShowSnapshotRequest) *ShowSnapshotInvoker {
	requestDef := GenReqDefForShowSnapshot()
	return &ShowSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSnapshotGroup 查询单个快照一致性组
//
// 查询单个快照一致性组。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ShowSnapshotGroup(request *model.ShowSnapshotGroupRequest) (*model.ShowSnapshotGroupResponse, error) {
	requestDef := GenReqDefForShowSnapshotGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSnapshotGroupResponse), nil
	}
}

// ShowSnapshotGroupInvoker 查询单个快照一致性组
func (c *EvsClient) ShowSnapshotGroupInvoker(request *model.ShowSnapshotGroupRequest) *ShowSnapshotGroupInvoker {
	requestDef := GenReqDefForShowSnapshotGroup()
	return &ShowSnapshotGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStandardSnapshot 查询单个快照信息
//
// 查询单个快照信息。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ShowStandardSnapshot(request *model.ShowStandardSnapshotRequest) (*model.ShowStandardSnapshotResponse, error) {
	requestDef := GenReqDefForShowStandardSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStandardSnapshotResponse), nil
	}
}

// ShowStandardSnapshotInvoker 查询单个快照信息
func (c *EvsClient) ShowStandardSnapshotInvoker(request *model.ShowStandardSnapshotRequest) *ShowStandardSnapshotInvoker {
	requestDef := GenReqDefForShowStandardSnapshot()
	return &ShowStandardSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVolume 查询单个云硬盘详情
//
// 查询单个云硬盘的详细信息。支持企业项目授权功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ShowVolume(request *model.ShowVolumeRequest) (*model.ShowVolumeResponse, error) {
	requestDef := GenReqDefForShowVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVolumeResponse), nil
	}
}

// ShowVolumeInvoker 查询单个云硬盘详情
func (c *EvsClient) ShowVolumeInvoker(request *model.ShowVolumeRequest) *ShowVolumeInvoker {
	requestDef := GenReqDefForShowVolume()
	return &ShowVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVolumeInRecycle 查询回收站中单个云硬盘详情
//
// 查询回收站中单个云硬盘的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ShowVolumeInRecycle(request *model.ShowVolumeInRecycleRequest) (*model.ShowVolumeInRecycleResponse, error) {
	requestDef := GenReqDefForShowVolumeInRecycle()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVolumeInRecycleResponse), nil
	}
}

// ShowVolumeInRecycleInvoker 查询回收站中单个云硬盘详情
func (c *EvsClient) ShowVolumeInRecycleInvoker(request *model.ShowVolumeInRecycleRequest) *ShowVolumeInRecycleInvoker {
	requestDef := GenReqDefForShowVolumeInRecycle()
	return &ShowVolumeInRecycleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVolumeTags 查询云硬盘标签
//
// 查询指定云硬盘的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ShowVolumeTags(request *model.ShowVolumeTagsRequest) (*model.ShowVolumeTagsResponse, error) {
	requestDef := GenReqDefForShowVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVolumeTagsResponse), nil
	}
}

// ShowVolumeTagsInvoker 查询云硬盘标签
func (c *EvsClient) ShowVolumeTagsInvoker(request *model.ShowVolumeTagsRequest) *ShowVolumeTagsInvoker {
	requestDef := GenReqDefForShowVolumeTags()
	return &ShowVolumeTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnsubscribePostpaidVolume 退订包周期计费模式的云硬盘
//
// 退订包周期计费模式的云硬盘，有如下约束：
// -  系统盘、启动盘不可使用当前接口退订，必须和弹性云服务器一起退订
// -  接口的请求body体最多可以传60个云硬盘id
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) UnsubscribePostpaidVolume(request *model.UnsubscribePostpaidVolumeRequest) (*model.UnsubscribePostpaidVolumeResponse, error) {
	requestDef := GenReqDefForUnsubscribePostpaidVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnsubscribePostpaidVolumeResponse), nil
	}
}

// UnsubscribePostpaidVolumeInvoker 退订包周期计费模式的云硬盘
func (c *EvsClient) UnsubscribePostpaidVolumeInvoker(request *model.UnsubscribePostpaidVolumeRequest) *UnsubscribePostpaidVolumeInvoker {
	requestDef := GenReqDefForUnsubscribePostpaidVolume()
	return &UnsubscribePostpaidVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAutoSnapshotPolicy 更新自动快照策略
//
// 更新自动快照策略。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) UpdateAutoSnapshotPolicy(request *model.UpdateAutoSnapshotPolicyRequest) (*model.UpdateAutoSnapshotPolicyResponse, error) {
	requestDef := GenReqDefForUpdateAutoSnapshotPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAutoSnapshotPolicyResponse), nil
	}
}

// UpdateAutoSnapshotPolicyInvoker 更新自动快照策略
func (c *EvsClient) UpdateAutoSnapshotPolicyInvoker(request *model.UpdateAutoSnapshotPolicyRequest) *UpdateAutoSnapshotPolicyInvoker {
	requestDef := GenReqDefForUpdateAutoSnapshotPolicy()
	return &UpdateAutoSnapshotPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRecyclePolicy 更新回收站策略
//
// 更新回收站策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) UpdateRecyclePolicy(request *model.UpdateRecyclePolicyRequest) (*model.UpdateRecyclePolicyResponse, error) {
	requestDef := GenReqDefForUpdateRecyclePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRecyclePolicyResponse), nil
	}
}

// UpdateRecyclePolicyInvoker 更新回收站策略
func (c *EvsClient) UpdateRecyclePolicyInvoker(request *model.UpdateRecyclePolicyRequest) *UpdateRecyclePolicyInvoker {
	requestDef := GenReqDefForUpdateRecyclePolicy()
	return &UpdateRecyclePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSnapshot 更新云硬盘快照
//
// 更新云硬盘快照。支持企业项目授权功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) UpdateSnapshot(request *model.UpdateSnapshotRequest) (*model.UpdateSnapshotResponse, error) {
	requestDef := GenReqDefForUpdateSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSnapshotResponse), nil
	}
}

// UpdateSnapshotInvoker 更新云硬盘快照
func (c *EvsClient) UpdateSnapshotInvoker(request *model.UpdateSnapshotRequest) *UpdateSnapshotInvoker {
	requestDef := GenReqDefForUpdateSnapshot()
	return &UpdateSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSnapshotGroup 更新快照一致性组
//
// 更新快照一致性组。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) UpdateSnapshotGroup(request *model.UpdateSnapshotGroupRequest) (*model.UpdateSnapshotGroupResponse, error) {
	requestDef := GenReqDefForUpdateSnapshotGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSnapshotGroupResponse), nil
	}
}

// UpdateSnapshotGroupInvoker 更新快照一致性组
func (c *EvsClient) UpdateSnapshotGroupInvoker(request *model.UpdateSnapshotGroupRequest) *UpdateSnapshotGroupInvoker {
	requestDef := GenReqDefForUpdateSnapshotGroup()
	return &UpdateSnapshotGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateStandardSnapshot 更新快照信息
//
// 更新快照信息。标准快照当前只在部分区域支持，具体的区域列表请通过控制台查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) UpdateStandardSnapshot(request *model.UpdateStandardSnapshotRequest) (*model.UpdateStandardSnapshotResponse, error) {
	requestDef := GenReqDefForUpdateStandardSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStandardSnapshotResponse), nil
	}
}

// UpdateStandardSnapshotInvoker 更新快照信息
func (c *EvsClient) UpdateStandardSnapshotInvoker(request *model.UpdateStandardSnapshotRequest) *UpdateStandardSnapshotInvoker {
	requestDef := GenReqDefForUpdateStandardSnapshot()
	return &UpdateStandardSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVolume 更新云硬盘
//
// 更新一个云硬盘的名称和描述。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) UpdateVolume(request *model.UpdateVolumeRequest) (*model.UpdateVolumeResponse, error) {
	requestDef := GenReqDefForUpdateVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVolumeResponse), nil
	}
}

// UpdateVolumeInvoker 更新云硬盘
func (c *EvsClient) UpdateVolumeInvoker(request *model.UpdateVolumeRequest) *UpdateVolumeInvoker {
	requestDef := GenReqDefForUpdateVolume()
	return &UpdateVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVersions 查询接口版本信息列表
//
// 查询接口版本信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ListVersions(request *model.ListVersionsRequest) (*model.ListVersionsResponse, error) {
	requestDef := GenReqDefForListVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVersionsResponse), nil
	}
}

// ListVersionsInvoker 查询接口版本信息列表
func (c *EvsClient) ListVersionsInvoker(request *model.ListVersionsRequest) *ListVersionsInvoker {
	requestDef := GenReqDefForListVersions()
	return &ListVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVersion 查询API接口的版本信息
//
// 查询接口的指定版本信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EvsClient) ShowVersion(request *model.ShowVersionRequest) (*model.ShowVersionResponse, error) {
	requestDef := GenReqDefForShowVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVersionResponse), nil
	}
}

// ShowVersionInvoker 查询API接口的版本信息
func (c *EvsClient) ShowVersionInvoker(request *model.ShowVersionRequest) *ShowVersionInvoker {
	requestDef := GenReqDefForShowVersion()
	return &ShowVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
