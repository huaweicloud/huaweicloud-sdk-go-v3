package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeviceTemplateRequest Request Object
type ShowDeviceTemplateRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 设备模板ID
	DeviceTemplateId string `json:"device_template_id"`
}

func (o ShowDeviceTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowDeviceTemplateRequest", string(data)}, " ")
}
