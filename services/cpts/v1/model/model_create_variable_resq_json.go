package model

import (
	"encoding/json"

	"strings"
)

// json
type CreateVariableResqJson struct {
	// variable_id

	VariableId *int32 `json:"variable_id,omitempty"`
}

func (o CreateVariableResqJson) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateVariableResqJson struct{}"
	}

	return strings.Join([]string{"CreateVariableResqJson", string(data)}, " ")
}
