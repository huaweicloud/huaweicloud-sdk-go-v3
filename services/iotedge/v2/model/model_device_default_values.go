package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceDefaultValues 设备属性默认值
type DeviceDefaultValues struct {

	// 设备ID
	DeviceId string `json:"device_id"`

	// 服务id，可选
	ServiceId *string `json:"service_id,omitempty"`

	// 属性key和value的map，用于设置属性的值
	Properties *interface{} `json:"properties"`
}

func (o DeviceDefaultValues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceDefaultValues struct{}"
	}

	return strings.Join([]string{"DeviceDefaultValues", string(data)}, " ")
}
