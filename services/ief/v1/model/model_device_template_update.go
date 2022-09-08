package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 设备模板
type DeviceTemplateUpdate struct {
	DeviceTemplate *DeviceTemplateUpdateDetail `json:"device_template"`
}

func (o DeviceTemplateUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceTemplateUpdate struct{}"
	}

	return strings.Join([]string{"DeviceTemplateUpdate", string(data)}, " ")
}
