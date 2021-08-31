package model

import (
	"encoding/json"

	"strings"
)

type UpdateCaseRequestBodyContentHeaders struct {
	// key

	Key *string `json:"key,omitempty"`
	// value

	Value *string `json:"value,omitempty"`
}

func (o UpdateCaseRequestBodyContentHeaders) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCaseRequestBodyContentHeaders struct{}"
	}

	return strings.Join([]string{"UpdateCaseRequestBodyContentHeaders", string(data)}, " ")
}
