package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtensionParameterValidation 参数校验规则
type ExtensionParameterValidation struct {

	// 是否必填
	IsRequired *bool `json:"isRequired,omitempty"`

	// 正则校验
	Pattern *string `json:"pattern,omitempty"`

	// 校验说明
	PatternDescription *string `json:"patternDescription,omitempty"`
}

func (o ExtensionParameterValidation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionParameterValidation struct{}"
	}

	return strings.Join([]string{"ExtensionParameterValidation", string(data)}, " ")
}
