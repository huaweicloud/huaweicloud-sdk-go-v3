package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudPhoneSingleServerRequestBodyDataVolume 数据盘结构体。
type CreateCloudPhoneSingleServerRequestBodyDataVolume struct {

	// 数据盘类型，只支持如下类型： SSD GPSSD
	VolumeType string `json:"volume_type"`

	// 数据盘大小，单位GiB，取值范围[1，32768]
	Size int32 `json:"size"`

	// 数据盘数量，默认值为1
	Count *int32 `json:"count,omitempty"`
}

func (o CreateCloudPhoneSingleServerRequestBodyDataVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudPhoneSingleServerRequestBodyDataVolume struct{}"
	}

	return strings.Join([]string{"CreateCloudPhoneSingleServerRequestBodyDataVolume", string(data)}, " ")
}
