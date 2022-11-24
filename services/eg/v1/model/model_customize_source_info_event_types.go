package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomizeSourceInfoEventTypes struct {

	// 事件类型名称
	Name *string `json:"name,omitempty"`

	// 事件类型描述
	Description *string `json:"description,omitempty"`
}

func (o CustomizeSourceInfoEventTypes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomizeSourceInfoEventTypes struct{}"
	}

	return strings.Join([]string{"CustomizeSourceInfoEventTypes", string(data)}, " ")
}
