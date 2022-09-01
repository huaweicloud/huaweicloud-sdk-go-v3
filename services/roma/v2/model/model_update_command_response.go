package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateCommandResponse struct {

	// 命令所属服务id
	ServiceId *int32 `json:"service_id,omitempty" xml:"service_id"`

	// 命令id
	CommandId *int32 `json:"command_id,omitempty" xml:"command_id"`

	// 命令名称
	CommandName *string `json:"command_name,omitempty" xml:"command_name"`

	// 命令描述
	Description    *string `json:"description,omitempty" xml:"description"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateCommandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCommandResponse struct{}"
	}

	return strings.Join([]string{"UpdateCommandResponse", string(data)}, " ")
}
