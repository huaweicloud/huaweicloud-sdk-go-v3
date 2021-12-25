package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteCommandRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 服务ID

	ServiceId string `json:"service_id"`
	// 命令ID

	CommandId int32 `json:"command_id"`
}

func (o DeleteCommandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCommandRequest struct{}"
	}

	return strings.Join([]string{"DeleteCommandRequest", string(data)}, " ")
}
