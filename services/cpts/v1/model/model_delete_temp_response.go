package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteTempResponse struct {
	// code

	Code *string `json:"code,omitempty"`
	// message

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTempResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTempResponse struct{}"
	}

	return strings.Join([]string{"DeleteTempResponse", string(data)}, " ")
}
