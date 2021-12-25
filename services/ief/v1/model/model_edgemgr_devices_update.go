package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 终端设备属性
type EdgemgrDevicesUpdate struct {
	Device *EdgemgrDevicesPara `json:"device"`
}

func (o EdgemgrDevicesUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgemgrDevicesUpdate struct{}"
	}

	return strings.Join([]string{"EdgemgrDevicesUpdate", string(data)}, " ")
}
