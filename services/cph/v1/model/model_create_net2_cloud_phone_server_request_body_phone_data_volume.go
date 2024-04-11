package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNet2CloudPhoneServerRequestBodyPhoneDataVolume 手机磁盘结构体，2.0规格必选。
type CreateNet2CloudPhoneServerRequestBodyPhoneDataVolume struct {

	// 磁盘类型，只支持如下类型： - SSD - GPSSD
	VolumeType string `json:"volume_type"`

	// 磁盘大小，单位GB，取值范围[0，32768]。
	Size int32 `json:"size"`
}

func (o CreateNet2CloudPhoneServerRequestBodyPhoneDataVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNet2CloudPhoneServerRequestBodyPhoneDataVolume struct{}"
	}

	return strings.Join([]string{"CreateNet2CloudPhoneServerRequestBodyPhoneDataVolume", string(data)}, " ")
}
