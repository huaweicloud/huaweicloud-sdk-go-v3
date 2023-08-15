package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceTemplate 设备模板
type DeviceTemplate struct {
	DeviceTemplate *EdgemgrDeviceReq `json:"device_template"`
}

func (o DeviceTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceTemplate struct{}"
	}

	return strings.Join([]string{"DeviceTemplate", string(data)}, " ")
}
