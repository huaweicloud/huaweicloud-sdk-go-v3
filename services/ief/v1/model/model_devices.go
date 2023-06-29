package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Devices 终端设备属性
type Devices struct {
	Devices *DevicesDevices `json:"devices"`
}

func (o Devices) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Devices struct{}"
	}

	return strings.Join([]string{"Devices", string(data)}, " ")
}
