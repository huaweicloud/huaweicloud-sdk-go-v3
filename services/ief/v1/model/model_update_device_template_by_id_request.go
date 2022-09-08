package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDeviceTemplateByIdRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 设备模板ID
	DeviceTemplateId string `json:"device_template_id"`

	Body *DeviceTemplateUpdate `json:"body,omitempty"`
}

func (o UpdateDeviceTemplateByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceTemplateByIdRequest struct{}"
	}

	return strings.Join([]string{"UpdateDeviceTemplateByIdRequest", string(data)}, " ")
}
