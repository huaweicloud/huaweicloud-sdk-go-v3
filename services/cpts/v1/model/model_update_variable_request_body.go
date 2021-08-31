package model

import (
	"encoding/json"

	"strings"
)

// UpdateVariableRequestBody
type UpdateVariableRequestBody struct {
	// id

	Id int32 `json:"id"`
	// name

	Name string `json:"name"`
	// variable_type

	VariableType int32 `json:"variable_type"`
	// variable

	Variable []interface{} `json:"variable"`
}

func (o UpdateVariableRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateVariableRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVariableRequestBody", string(data)}, " ")
}
