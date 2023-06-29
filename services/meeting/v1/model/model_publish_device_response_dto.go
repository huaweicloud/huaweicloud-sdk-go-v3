package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishDeviceResponseDto 发布的设备信息。
type PublishDeviceResponseDto struct {

	// 设备用户ID。
	DeviceUserId *string `json:"deviceUserId,omitempty"`

	// 设备名称。
	DeviceName *string `json:"deviceName,omitempty"`
}

func (o PublishDeviceResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishDeviceResponseDto struct{}"
	}

	return strings.Join([]string{"PublishDeviceResponseDto", string(data)}, " ")
}
