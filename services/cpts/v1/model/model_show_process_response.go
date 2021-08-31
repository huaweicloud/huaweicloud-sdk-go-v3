package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowProcessResponse struct {
	// code

	Code *string `json:"code,omitempty"`
	// message

	Message *string `json:"message,omitempty"`

	Json *ShowProcessResqJson `json:"json,omitempty"`
	// extend

	Extend         *string `json:"extend,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowProcessResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProcessResponse struct{}"
	}

	return strings.Join([]string{"ShowProcessResponse", string(data)}, " ")
}
