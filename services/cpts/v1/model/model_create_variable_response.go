package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateVariableResponse struct {
	// code

	Code *string `json:"code,omitempty"`

	Json *CreateVariableResqJson `json:"json,omitempty"`
	// message

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVariableResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateVariableResponse struct{}"
	}

	return strings.Join([]string{"CreateVariableResponse", string(data)}, " ")
}
