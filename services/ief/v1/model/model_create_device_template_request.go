package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDeviceTemplateRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty" xml:"ief-instance-id"`

	Body *DeviceTemplate `json:"body,omitempty" xml:"body"`
}

func (o CreateDeviceTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeviceTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateDeviceTemplateRequest", string(data)}, " ")
}
