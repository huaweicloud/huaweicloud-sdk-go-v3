package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 硬盘对象
type Volume struct {

	// 硬盘ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 磁盘状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 磁盘大小。
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 硬盘所属的AZ信息。
	AvailabilityZone *string `json:"availability_zone,omitempty" xml:"availability_zone"`

	// 硬盘的挂载信息。
	Attachments *[]Attachment `json:"attachments,omitempty" xml:"attachments"`

	// 磁盘名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 磁盘类型。
	VolumeType *string `json:"volume_type,omitempty" xml:"volume_type"`

	// 显示这个卷是否可启动。
	Bootable *string `json:"bootable,omitempty" xml:"bootable"`

	// 显示该卷是否已被加密。
	Encrypted *bool `json:"encrypted,omitempty" xml:"encrypted"`

	// 磁盘是否多挂载。
	Multiattach bool `json:"multiattach" xml:"multiattach"`

	// 硬盘的元数据。
	Metadata map[string]string `json:"metadata,omitempty" xml:"metadata"`
}

func (o Volume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Volume struct{}"
	}

	return strings.Join([]string{"Volume", string(data)}, " ")
}
