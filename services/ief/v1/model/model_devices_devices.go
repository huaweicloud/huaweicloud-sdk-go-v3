package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DevicesDevices 终端设备信息
type DevicesDevices struct {
	Added *DevicesDevicesAdded `json:"added,omitempty"`

	// 要解绑的终端设备ID
	Removed *[]string `json:"removed,omitempty"`
}

func (o DevicesDevices) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DevicesDevices struct{}"
	}

	return strings.Join([]string{"DevicesDevices", string(data)}, " ")
}
