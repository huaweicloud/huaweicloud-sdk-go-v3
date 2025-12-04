package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNet2CloudPhoneServerRequestBodyServerShareDataVolume 共享存储磁盘结构体，2.0规格可选。注：rx7服务器+AOSP11镜像当前不支持共享应用。
type CreateNet2CloudPhoneServerRequestBodyServerShareDataVolume struct {

	// 磁盘类型，只支持如下类型：   - SSD   - GPSSD
	VolumeType string `json:"volume_type"`

	// 共享磁盘大小，单位GiB，取值范围[10，32768]。
	Size int32 `json:"size"`
}

func (o CreateNet2CloudPhoneServerRequestBodyServerShareDataVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNet2CloudPhoneServerRequestBodyServerShareDataVolume struct{}"
	}

	return strings.Join([]string{"CreateNet2CloudPhoneServerRequestBodyServerShareDataVolume", string(data)}, " ")
}
