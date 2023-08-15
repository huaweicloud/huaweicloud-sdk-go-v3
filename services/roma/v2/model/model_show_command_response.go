package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCommandResponse Response Object
type ShowCommandResponse struct {

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

func (o ShowCommandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommandResponse struct{}"
	}

	return strings.Join([]string{"ShowCommandResponse", string(data)}, " ")
}
