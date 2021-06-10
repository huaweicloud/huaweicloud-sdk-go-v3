package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ImportFunctionRequest struct {
	Body *ImportFunctionRequestBody `json:"body,omitempty"`
}

func (o ImportFunctionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImportFunctionRequest struct{}"
	}

	return strings.Join([]string{"ImportFunctionRequest", string(data)}, " ")
}
