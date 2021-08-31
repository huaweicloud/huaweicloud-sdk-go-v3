package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateCaseResponse struct {
	// code

	Code *string `json:"code,omitempty"`

	Json *CreateCaseResqJson `json:"json,omitempty"`
	// message

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCaseResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCaseResponse struct{}"
	}

	return strings.Join([]string{"CreateCaseResponse", string(data)}, " ")
}
