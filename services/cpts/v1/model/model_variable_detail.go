package model

import (
	"encoding/json"

	"strings"
)

type VariableDetail struct {
	// file_size

	FileSize *int32 `json:"file_size,omitempty"`
	// id

	Id *int32 `json:"id,omitempty"`
	// is_quoted

	IsQuoted *bool `json:"is_quoted,omitempty"`
	// name

	Name *string `json:"name,omitempty"`
	// variable

	Variable *[]interface{} `json:"variable,omitempty"`
	// variable_type

	VariableType *int32 `json:"variable_type,omitempty"`
}

func (o VariableDetail) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VariableDetail struct{}"
	}

	return strings.Join([]string{"VariableDetail", string(data)}, " ")
}
