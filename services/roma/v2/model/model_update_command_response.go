package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCommandResponse Response Object
type UpdateCommandResponse struct {

	// 命令所属服务id
	ServiceId *int32 `json:"service_id,omitempty"`

	// 命令id
	CommandId *int32 `json:"command_id,omitempty"`

	// 命令名称
	CommandName *string `json:"command_name,omitempty"`

	// 命令描述
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateCommandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCommandResponse struct{}"
	}

	return strings.Join([]string{"UpdateCommandResponse", string(data)}, " ")
}
