package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteRequestPropertyRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 服务ID

	ServiceId string `json:"service_id"`
	// 命令ID

	CommandId int32 `json:"command_id"`
	// 属性/请求属性/响应属性ID

	PropertyId int32 `json:"property_id"`
}

func (o DeleteRequestPropertyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRequestPropertyRequest struct{}"
	}

	return strings.Join([]string{"DeleteRequestPropertyRequest", string(data)}, " ")
}
