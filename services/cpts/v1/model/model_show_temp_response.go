package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowTempResponse struct {
	// code

	Code *string `json:"code,omitempty"`
	// message

	Message *string `json:"message,omitempty"`
	// temp_info

	TempInfo       *interface{} `json:"temp_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowTempResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTempResponse struct{}"
	}

	return strings.Join([]string{"ShowTempResponse", string(data)}, " ")
}
