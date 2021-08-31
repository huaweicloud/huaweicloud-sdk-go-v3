package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowProjectResponse struct {
	// code

	Code *string `json:"code,omitempty"`
	// message

	Message *string `json:"message,omitempty"`

	Project        *ShowProjectResqProject `json:"project,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowProjectResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProjectResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectResponse", string(data)}, " ")
}
