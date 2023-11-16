package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtensionValidation struct {

	// 消息
	RequiredMessage *string `json:"required_message,omitempty"`

	// 正则
	Regex *string `json:"regex,omitempty"`

	// 正则消息
	RegexMessage *string `json:"regex_message,omitempty"`

	// 最大长度
	MaxLength *int32 `json:"max_length,omitempty"`

	// 最小长度
	MinLength *int32 `json:"min_length,omitempty"`
}

func (o ExtensionValidation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionValidation struct{}"
	}

	return strings.Join([]string{"ExtensionValidation", string(data)}, " ")
}
