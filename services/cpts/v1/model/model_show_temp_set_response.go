package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowTempSetResponse struct {
	// code

	Code *string `json:"code,omitempty"`
	// message

	Message *string `json:"message,omitempty"`
	// temps

	Temps          *[]interface{} `json:"temps,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowTempSetResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTempSetResponse struct{}"
	}

	return strings.Join([]string{"ShowTempSetResponse", string(data)}, " ")
}
