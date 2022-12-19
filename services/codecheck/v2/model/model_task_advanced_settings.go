package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskAdvancedSettings struct {

	// 高级选项对应的名称
	Key *string `json:"key,omitempty"`

	// 高级选项对应的取值
	Value *string `json:"value,omitempty"`

	// 高级选项对应的可选项
	OptionValue *string `json:"option_value,omitempty"`

	// 高级选项对应的中文描述
	Description *string `json:"description,omitempty"`
}

func (o TaskAdvancedSettings) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskAdvancedSettings struct{}"
	}

	return strings.Join([]string{"TaskAdvancedSettings", string(data)}, " ")
}
