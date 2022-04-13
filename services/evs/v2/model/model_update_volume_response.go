package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateVolumeResponse struct {
	// 是否挂载信息。

	Attachments *[]Attachment `json:"attachments,omitempty"`
	// 云硬盘所属AZ。

	AvailabilityZone *string `json:"availability_zone,omitempty"`
	// 是否为可启动云硬盘。

	Bootable *string `json:"bootable,omitempty"`
	// 创建云硬盘的时间。

	CreatedAt *string `json:"created_at,omitempty"`
	// 云硬盘ID。

	Id *string `json:"id,omitempty"`
	// 云硬盘uri自描述信息

	Links *[]Link `json:"links,omitempty"`

	Metadata *VolumeMetadata `json:"metadata,omitempty"`
	// 是否为可共享云硬盘。

	Multiattach *bool `json:"multiattach,omitempty"`
	// 云硬盘名称

	Name *string `json:"name,omitempty"`
	// 预留属性。

	OsVolHostAttrhost *string `json:"os-vol-host-attr:host,omitempty"`
	// 云硬盘所属的项目ID。

	OsVolTenantAttrtenantId *string `json:"os-vol-tenant-attr:tenant_id,omitempty"`
	// 是否为共享云硬盘。

	Shareable *string `json:"shareable,omitempty"`
	// 云硬盘大小。

	Size *int32 `json:"size,omitempty"`
	// 快照ID。

	SnapshotId *string `json:"snapshot_id,omitempty"`
	// 预留字段。

	SourceVolid *string `json:"source_volid,omitempty"`
	// 云硬盘状态。

	Status *string `json:"status,omitempty"`
	// 云硬盘镜像的元数据。 > 说明： >  > 关于“volume_image_metadata”字段的详细说明，具体请参见：\"[查询镜像详情](https://support.huaweicloud.com/api-ims/ims_03_0703.html)\"。

	VolumeImageMetadata *interface{} `json:"volume_image_metadata,omitempty"`
	// 云硬盘类型。

	VolumeType *string `json:"volume_type,omitempty"`
	// 云硬盘描述。

	Description *string `json:"description,omitempty"`
	// 预留属性。

	OsVolumeReplicationextendedStatus *string `json:"os-volume-replication:extended_status,omitempty"`
	HttpStatusCode                    int     `json:"-"`
}

func (o UpdateVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVolumeResponse struct{}"
	}

	return strings.Join([]string{"UpdateVolumeResponse", string(data)}, " ")
}
