package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteResponsePropertyRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 服务ID

	ServiceId string `json:"service_id"`
	// 命令ID

	CommandId int32 `json:"command_id"`
	// 属性/请求属性/响应属性ID

	PropertyId int32 `json:"property_id"`
}

func (o DeleteResponsePropertyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResponsePropertyRequest struct{}"
	}

	return strings.Join([]string{"DeleteResponsePropertyRequest", string(data)}, " ")
}
