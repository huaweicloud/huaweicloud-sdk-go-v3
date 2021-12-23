package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVariableRequestBody
type CreateVariableRequestBody struct {
	// id

	Id int32 `json:"id"`
	// name

	Name string `json:"name"`
	// variable_type

	VariableType int32 `json:"variable_type"`
	// variable

	Variable []interface{} `json:"variable"`
}

func (o CreateVariableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVariableRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVariableRequestBody", string(data)}, " ")
}
