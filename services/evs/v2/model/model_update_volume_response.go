package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateVolumeResponse struct {

	// 是否挂载信息。
	Attachments *[]Attachment `json:"attachments,omitempty" xml:"attachments"`

	// 云硬盘所属AZ。
	AvailabilityZone *string `json:"availability_zone,omitempty" xml:"availability_zone"`

	// 是否为可启动云硬盘。
	Bootable *string `json:"bootable,omitempty" xml:"bootable"`

	// 创建云硬盘的时间。
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 云硬盘ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 云硬盘uri自描述信息
	Links *[]Link `json:"links,omitempty" xml:"links"`

	Metadata *VolumeMetadata `json:"metadata,omitempty" xml:"metadata"`

	// 是否为可共享云硬盘。
	Multiattach *bool `json:"multiattach,omitempty" xml:"multiattach"`

	// 云硬盘名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 预留属性。
	OsVolHostAttrhost *string `json:"os-vol-host-attr:host,omitempty" xml:"os-vol-host-attr:host"`

	// 云硬盘所属的项目ID。
	OsVolTenantAttrtenantId *string `json:"os-vol-tenant-attr:tenant_id,omitempty" xml:"os-vol-tenant-attr:tenant_id"`

	// 是否为共享云硬盘。
	Shareable *string `json:"shareable,omitempty" xml:"shareable"`

	// 云硬盘大小。
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 快照ID。
	SnapshotId *string `json:"snapshot_id,omitempty" xml:"snapshot_id"`

	// 预留字段。
	SourceVolid *string `json:"source_volid,omitempty" xml:"source_volid"`

	// 云硬盘状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 云硬盘镜像的元数据。 > 说明： >  > 关于“volume_image_metadata”字段的详细说明，具体请参见：\"[查询镜像详情](https://support.huaweicloud.com/api-ims/ims_03_0703.html)\"。
	VolumeImageMetadata *interface{} `json:"volume_image_metadata,omitempty" xml:"volume_image_metadata"`

	// 云硬盘类型。
	VolumeType *string `json:"volume_type,omitempty" xml:"volume_type"`

	// 云硬盘描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 预留属性。
	OsVolumeReplicationextendedStatus *string `json:"os-volume-replication:extended_status,omitempty" xml:"os-volume-replication:extended_status"`
	HttpStatusCode                    int     `json:"-"`
}

func (o UpdateVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVolumeResponse struct{}"
	}

	return strings.Join([]string{"UpdateVolumeResponse", string(data)}, " ")
}
