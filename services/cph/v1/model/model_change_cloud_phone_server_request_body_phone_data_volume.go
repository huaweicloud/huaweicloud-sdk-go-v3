package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeCloudPhoneServerRequestBodyPhoneDataVolume 手机磁盘结构体，2.0规格必选。
type ChangeCloudPhoneServerRequestBodyPhoneDataVolume struct {

	// 磁盘类型，只支持如下类型： - SSD - GPSSD
	VolumeType string `json:"volume_type"`

	// 手机磁盘大小，单位GiB，取值范围[10，32768]。
	Size int32 `json:"size"`
}

func (o ChangeCloudPhoneServerRequestBodyPhoneDataVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeCloudPhoneServerRequestBodyPhoneDataVolume struct{}"
	}

	return strings.Join([]string{"ChangeCloudPhoneServerRequestBodyPhoneDataVolume", string(data)}, " ")
}
