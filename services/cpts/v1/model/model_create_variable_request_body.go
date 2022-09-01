package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVariableRequestBody
type CreateVariableRequestBody struct {

	// id
	Id int32 `json:"id" xml:"id"`

	// name
	Name string `json:"name" xml:"name"`

	// variable_type
	VariableType int32 `json:"variable_type" xml:"variable_type"`

	// variable
	Variable []interface{} `json:"variable" xml:"variable"`

	// 是否被引用
	IsQuoted bool `json:"is_quoted" xml:"is_quoted"`
}

func (o CreateVariableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVariableRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVariableRequestBody", string(data)}, " ")
}
