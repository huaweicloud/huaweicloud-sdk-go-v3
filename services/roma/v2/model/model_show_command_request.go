package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowCommandRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 服务ID

	ServiceId string `json:"service_id"`
	// 命令ID

	CommandId int32 `json:"command_id"`
}

func (o ShowCommandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommandRequest struct{}"
	}

	return strings.Join([]string{"ShowCommandRequest", string(data)}, " ")
}
