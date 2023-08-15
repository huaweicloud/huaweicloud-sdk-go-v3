package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VariableValidationResponse struct {

	// 校验表达式
	Condition *string `json:"condition,omitempty"`

	// 校验失败后的错误信息
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o VariableValidationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VariableValidationResponse struct{}"
	}

	return strings.Join([]string{"VariableValidationResponse", string(data)}, " ")
}
