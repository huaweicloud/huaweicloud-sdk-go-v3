package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCommandRequest Request Object
type UpdateCommandRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 服务ID
	ServiceId string `json:"service_id"`

	// 命令ID
	CommandId int32 `json:"command_id"`

	Body *UpdateCommandRequestBody `json:"body,omitempty"`
}

func (o UpdateCommandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCommandRequest struct{}"
	}

	return strings.Join([]string{"UpdateCommandRequest", string(data)}, " ")
}
