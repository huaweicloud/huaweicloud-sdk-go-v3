package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRequestPropertyRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 服务ID
	ServiceId string `json:"service_id" xml:"service_id"`

	// 命令ID
	CommandId int32 `json:"command_id" xml:"command_id"`

	// 属性/请求属性/响应属性ID
	PropertyId int32 `json:"property_id" xml:"property_id"`
}

func (o ShowRequestPropertyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRequestPropertyRequest struct{}"
	}

	return strings.Join([]string{"ShowRequestPropertyRequest", string(data)}, " ")
}
