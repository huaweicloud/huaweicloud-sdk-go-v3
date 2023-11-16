package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CodeEvent struct {

	// 事件类型
	Type *string `json:"type,omitempty"`

	// 是否可用
	Enable *bool `json:"enable,omitempty"`
}

func (o CodeEvent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeEvent struct{}"
	}

	return strings.Join([]string{"CodeEvent", string(data)}, " ")
}
