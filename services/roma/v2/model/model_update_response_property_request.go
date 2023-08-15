package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResponsePropertyRequest Request Object
type UpdateResponsePropertyRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 服务ID
	ServiceId string `json:"service_id"`

	// 命令ID
	CommandId int32 `json:"command_id"`

	// 属性/请求属性/响应属性ID
	PropertyId int32 `json:"property_id"`

	Body *UpdatePropertyRequestBody `json:"body,omitempty"`
}

func (o UpdateResponsePropertyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResponsePropertyRequest struct{}"
	}

	return strings.Join([]string{"UpdateResponsePropertyRequest", string(data)}, " ")
}
