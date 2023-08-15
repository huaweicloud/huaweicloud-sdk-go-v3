package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceTemplateUpdate 设备模板
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
