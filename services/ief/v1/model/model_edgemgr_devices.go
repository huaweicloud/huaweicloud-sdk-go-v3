package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 终端设备属性
type EdgemgrDevices struct {
	Device *EdgemgrDevicesDetail `json:"device"`
}

func (o EdgemgrDevices) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgemgrDevices struct{}"
	}

	return strings.Join([]string{"EdgemgrDevices", string(data)}, " ")
}
