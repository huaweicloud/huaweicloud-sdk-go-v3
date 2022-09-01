package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VariableDetail struct {

	// file_size
	FileSize *int32 `json:"file_size,omitempty" xml:"file_size"`

	// id
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 是否被引用
	IsQuoted *bool `json:"is_quoted,omitempty" xml:"is_quoted"`

	// name
	Name *string `json:"name,omitempty" xml:"name"`

	// variable
	Variable *[]interface{} `json:"variable,omitempty" xml:"variable"`

	// variable_type
	VariableType *int32 `json:"variable_type,omitempty" xml:"variable_type"`
}

func (o VariableDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VariableDetail struct{}"
	}

	return strings.Join([]string{"VariableDetail", string(data)}, " ")
}
