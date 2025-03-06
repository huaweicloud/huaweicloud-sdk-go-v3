package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerModelExtendSpecOsVolume 云盘服务器操作系统盘信息，本地盘服务器不返回
type ServerModelExtendSpecOsVolume struct {

	// 操作系统盘大小
	Size *int32 `json:"size,omitempty"`

	// 操作系统盘类型
	VolumeType *string `json:"volume_type,omitempty"`
}

func (o ServerModelExtendSpecOsVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerModelExtendSpecOsVolume struct{}"
	}

	return strings.Join([]string{"ServerModelExtendSpecOsVolume", string(data)}, " ")
}
