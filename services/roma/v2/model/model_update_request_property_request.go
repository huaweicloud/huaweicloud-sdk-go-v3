package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateRequestPropertyRequest struct {
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

func (o UpdateRequestPropertyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRequestPropertyRequest struct{}"
	}

	return strings.Join([]string{"UpdateRequestPropertyRequest", string(data)}, " ")
}
