package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Command struct {
	// 命令所属服务id

	ServiceId *int32 `json:"service_id,omitempty"`
	// 命令id

	CommandId *int32 `json:"command_id,omitempty"`
	// 命令名称

	CommandName *string `json:"command_name,omitempty"`
	// 命令描述

	Description *string `json:"description,omitempty"`
}

func (o Command) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Command struct{}"
	}

	return strings.Join([]string{"Command", string(data)}, " ")
}
