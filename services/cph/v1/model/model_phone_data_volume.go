package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PhoneDataVolume 云手机数据盘信息。
type PhoneDataVolume struct {

	// 云手机数据盘类型。
	VolumeType *string `json:"volume_type,omitempty"`

	// 云手机数据盘大小。
	VolumeSize *int32 `json:"volume_size,omitempty"`
}

func (o PhoneDataVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhoneDataVolume struct{}"
	}

	return strings.Join([]string{"PhoneDataVolume", string(data)}, " ")
}
