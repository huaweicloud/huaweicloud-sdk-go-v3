package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/evs/v2/model"
)

type EvsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewEvsClient(hcClient *http_client.HcHttpClient) *EvsClient {
	return &EvsClient{HcClient: hcClient}
}

func EvsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 为指定云硬盘批量添加标签
//
// 为指定云硬盘批量添加标签。
//
// 添加标签时，如果云硬盘的标签已存在相同key，则会覆盖已有标签。
// 单个云硬盘最多支持创建10个标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) BatchCreateVolumeTags(request *model.BatchCreateVolumeTagsRequest) (*model.BatchCreateVolumeTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateVolumeTagsResponse), nil
	}
}

// 为指定云硬盘批量删除标签
//
// 为指定云硬盘批量删除标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) BatchDeleteVolumeTags(request *model.BatchDeleteVolumeTagsRequest) (*model.BatchDeleteVolumeTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteVolumeTagsResponse), nil
	}
}

// 查询所有的可用分区信息
//
// 查询所有的可用分区信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) CinderListAvailabilityZones(request *model.CinderListAvailabilityZonesRequest) (*model.CinderListAvailabilityZonesResponse, error) {
	requestDef := GenReqDefForCinderListAvailabilityZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderListAvailabilityZonesResponse), nil
	}
}

// 查询租户的详细配额
//
// 查询租户的详细配额。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) CinderListQuotas(request *model.CinderListQuotasRequest) (*model.CinderListQuotasResponse, error) {
	requestDef := GenReqDefForCinderListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderListQuotasResponse), nil
	}
}

// 查询云硬盘类型列表
//
// 查询云硬盘类型列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) CinderListVolumeTypes(request *model.CinderListVolumeTypesRequest) (*model.CinderListVolumeTypesResponse, error) {
	requestDef := GenReqDefForCinderListVolumeTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderListVolumeTypesResponse), nil
	}
}

// 创建云硬盘快照
//
// 创建云硬盘快照。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) CreateSnapshot(request *model.CreateSnapshotRequest) (*model.CreateSnapshotResponse, error) {
	requestDef := GenReqDefForCreateSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSnapshotResponse), nil
	}
}

// 创建云硬盘
//
// 创建按需或包周期云硬盘。
// 在创建包周期云硬盘的场景下：
// - 如果您需要查看订单可用的优惠券，请参考\&quot;[查询订单可用优惠券](https://support.huaweicloud.com/api-oce/zh-cn_topic_0092953630.html)\&quot;。
// - 如果您需要支付订单，请参考\&quot;[支付包周期产品订单](https://support.huaweicloud.com/api-oce/zh-cn_topic_0075746561.html)\&quot;。
// - 如果您需要查询订单的资源开通详情，请参考\&quot;[查询订单的资源开通详情](https://support.huaweicloud.com/api-oce/api_order_00001.html)\&quot;。
// - 如果您需要退订该包周期资源，请参考“[退订包周期资源](https://support.huaweicloud.com/api-oce/zh-cn_topic_0082522030.html)”。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) CreateVolume(request *model.CreateVolumeRequest) (*model.CreateVolumeResponse, error) {
	requestDef := GenReqDefForCreateVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVolumeResponse), nil
	}
}

// 删除云硬盘快照
//
// 删除云硬盘快照。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) DeleteSnapshot(request *model.DeleteSnapshotRequest) (*model.DeleteSnapshotResponse, error) {
	requestDef := GenReqDefForDeleteSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSnapshotResponse), nil
	}
}

// 删除云硬盘
//
// 删除一个云硬盘。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) DeleteVolume(request *model.DeleteVolumeRequest) (*model.DeleteVolumeResponse, error) {
	requestDef := GenReqDefForDeleteVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVolumeResponse), nil
	}
}

// 查询云硬盘快照详细列表信息
//
// 查询云硬盘快照详细列表信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) ListSnapshots(request *model.ListSnapshotsRequest) (*model.ListSnapshotsResponse, error) {
	requestDef := GenReqDefForListSnapshots()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotsResponse), nil
	}
}

// 获取云硬盘资源的所有标签
//
// 获取某个租户的所有云硬盘资源的标签信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) ListVolumeTags(request *model.ListVolumeTagsRequest) (*model.ListVolumeTagsResponse, error) {
	requestDef := GenReqDefForListVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumeTagsResponse), nil
	}
}

// 查询所有云硬盘详情
//
// 查询所有云硬盘的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) ListVolumes(request *model.ListVolumesRequest) (*model.ListVolumesResponse, error) {
	requestDef := GenReqDefForListVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumesResponse), nil
	}
}

// 通过标签查询云硬盘资源实例详情
//
// 通过标签查询云硬盘资源实例详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) ListVolumesByTags(request *model.ListVolumesByTagsRequest) (*model.ListVolumesByTagsResponse, error) {
	requestDef := GenReqDefForListVolumesByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumesByTagsResponse), nil
	}
}

// 扩容云硬盘
//
// 对按需或者包周期云硬盘进行扩容。
// 在扩容包周期云硬盘的场景下：
// - 如果您需要查看订单可用的优惠券，请参考\&quot;[查询订单可用优惠券](https://support.huaweicloud.com/api-oce/zh-cn_topic_0092953630.html)\&quot;。
// - 如果您需要支付订单，请参考\&quot;[支付包周期产品订单](https://support.huaweicloud.com/api-oce/zh-cn_topic_0075746561.html)\&quot;。
// - 如果您需要查询订单的资源开通详情，请参考\&quot;[查询订单的资源开通详情](https://support.huaweicloud.com/api-oce/api_order_00001.html)\&quot;。
// - 如果您需要退订该包周期资源，请参考“[退订包周期资源](https://support.huaweicloud.com/api-oce/zh-cn_topic_0082522030.html)”。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) ResizeVolume(request *model.ResizeVolumeRequest) (*model.ResizeVolumeResponse, error) {
	requestDef := GenReqDefForResizeVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeVolumeResponse), nil
	}
}

// 回滚快照到云硬盘
//
// 将快照数据回滚到云硬盘。支持企业项目授权功能。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) RollbackSnapshot(request *model.RollbackSnapshotRequest) (*model.RollbackSnapshotResponse, error) {
	requestDef := GenReqDefForRollbackSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RollbackSnapshotResponse), nil
	}
}

// 查询job的状态
//
// 查询Job的执行状态。
// 可用于查询创建云硬盘，扩容云硬盘，删除云硬盘等API的执行状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

// 查询单个云硬盘快照详细信息
//
// 查询单个云硬盘快照信息。支持企业项目授权功能。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) ShowSnapshot(request *model.ShowSnapshotRequest) (*model.ShowSnapshotResponse, error) {
	requestDef := GenReqDefForShowSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSnapshotResponse), nil
	}
}

// 查询单个云硬盘详情
//
// 查询单个云硬盘的详细信息。支持企业项目授权功能。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) ShowVolume(request *model.ShowVolumeRequest) (*model.ShowVolumeResponse, error) {
	requestDef := GenReqDefForShowVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVolumeResponse), nil
	}
}

// 查询云硬盘标签
//
// 查询指定云硬盘的标签信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) ShowVolumeTags(request *model.ShowVolumeTagsRequest) (*model.ShowVolumeTagsResponse, error) {
	requestDef := GenReqDefForShowVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVolumeTagsResponse), nil
	}
}

// 更新云硬盘快照
//
// 更新云硬盘快照。支持企业项目授权功能。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) UpdateSnapshot(request *model.UpdateSnapshotRequest) (*model.UpdateSnapshotResponse, error) {
	requestDef := GenReqDefForUpdateSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSnapshotResponse), nil
	}
}

// 更新云硬盘
//
// 更新一个云硬盘的名称和描述。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EvsClient) UpdateVolume(request *model.UpdateVolumeRequest) (*model.UpdateVolumeResponse, error) {
	requestDef := GenReqDefForUpdateVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVolumeResponse), nil
	}
}
