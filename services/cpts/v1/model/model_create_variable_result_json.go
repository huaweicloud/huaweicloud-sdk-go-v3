package model

import (
	"encoding/json"

	"strings"
)

type CreateVariableResultJson struct {
	// variable_id

	VariableId *int32 `json:"variable_id,omitempty"`
}

func (o CreateVariableResultJson) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateVariableResultJson struct{}"
	}

	return strings.Join([]string{"CreateVariableResultJson", string(data)}, " ")
}
