package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeCloudPhoneServerRequestBodyServerShareDataVolume 共享存储磁盘结构体，2.0规格可选。
type ChangeCloudPhoneServerRequestBodyServerShareDataVolume struct {

	// 磁盘类型，只支持如下类型：   - SSD   - GPSSD
	VolumeType string `json:"volume_type"`

	// 磁盘大小，单位GB，取值范围[0，32768]。
	Size int32 `json:"size"`
}

func (o ChangeCloudPhoneServerRequestBodyServerShareDataVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeCloudPhoneServerRequestBodyServerShareDataVolume struct{}"
	}

	return strings.Join([]string{"ChangeCloudPhoneServerRequestBodyServerShareDataVolume", string(data)}, " ")
}
