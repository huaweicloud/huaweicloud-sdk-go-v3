package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Volume 云手机服务器卷信息。
type Volume struct {

	// 云手机服务器的硬盘名称。
	VolumeName *string `json:"volume_name,omitempty"`

	// 云手机服务器的硬盘唯一标识。
	VolumeId *string `json:"volume_id,omitempty"`

	// 云手机服务器的硬盘大小，单位G。
	VolumeSize *int32 `json:"volume_size,omitempty"`

	// 云手机服务器的硬盘类型。
	VolumeType *string `json:"volume_type,omitempty"`

	// 硬盘创建时间，  时间格式为UTC，YYYY-MM-DDTHH:MM:SSZ。
	CreateTime *string `json:"create_time,omitempty"`

	// 硬盘更新时间，  时间格式为UTC，YYYY-MM-DDTHH:MM:SSZ。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o Volume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Volume struct{}"
	}

	return strings.Join([]string{"Volume", string(data)}, " ")
}
