package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 云手机服务器卷信息
type Volume struct {

	// 云手机服务器的硬盘名称
	VolumeName string `json:"volume_name"`

	// 云手机服务器的硬盘唯一标识
	VolumeId string `json:"volume_id"`

	// 云手机服务器的硬盘大小，单位G
	VolumeSize int32 `json:"volume_size"`

	// 云手机服务器的硬盘类型
	VolumeType string `json:"volume_type"`

	// 硬盘创建时间  时间格式为UTC，YYYY-MM-DDTHH:MM:SSZ
	CreateTime string `json:"create_time"`

	// 硬盘更新时间  时间格式为UTC，YYYY-MM-DDTHH:MM:SSZ
	UpdateTime string `json:"update_time"`
}

func (o Volume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Volume struct{}"
	}

	return strings.Join([]string{"Volume", string(data)}, " ")
}
