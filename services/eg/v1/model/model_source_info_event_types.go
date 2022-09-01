package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SourceInfoEventTypes struct {

	// 事件类型名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 事件类型描述
	Description *string `json:"description,omitempty" xml:"description"`
}

func (o SourceInfoEventTypes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceInfoEventTypes struct{}"
	}

	return strings.Join([]string{"SourceInfoEventTypes", string(data)}, " ")
}
