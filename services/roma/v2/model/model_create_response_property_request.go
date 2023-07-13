package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResponsePropertyRequest Request Object
type CreateResponsePropertyRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 服务ID
	ServiceId string `json:"service_id"`

	// 命令ID
	CommandId int32 `json:"command_id"`

	Body *CreatePropertyRequestBody `json:"body,omitempty"`
}

func (o CreateResponsePropertyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResponsePropertyRequest struct{}"
	}

	return strings.Join([]string{"CreateResponsePropertyRequest", string(data)}, " ")
}
