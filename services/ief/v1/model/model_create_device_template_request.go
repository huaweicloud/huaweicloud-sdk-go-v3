package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDeviceTemplateRequest struct {
	// 铂金版实例ID，专业版实例为空值

	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	Body *DeviceTemplate `json:"body,omitempty"`
}

func (o CreateDeviceTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeviceTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateDeviceTemplateRequest", string(data)}, " ")
}
