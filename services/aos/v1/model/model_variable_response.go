package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// variable response
type VariableResponse struct {

	// 参数的名字
	Name *string `json:"name,omitempty"`

	// 参数类型
	Type *string `json:"type,omitempty"`

	// 描述参数的用途
	Description *string `json:"description,omitempty"`

	// 参数默认值。该类型与type保持一致
	Default *interface{} `json:"default,omitempty"`

	// 参数是否为敏感字段
	Sensitive *bool `json:"sensitive,omitempty"`

	// 参数是否可设置为null
	Nullable *bool `json:"nullable,omitempty"`

	// 参数校验模块
	Validations *[]VariableValidationResponse `json:"validations,omitempty"`
}

func (o VariableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VariableResponse struct{}"
	}

	return strings.Join([]string{"VariableResponse", string(data)}, " ")
}
