package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UntagDeviceResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UntagDeviceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UntagDeviceResponse struct{}"
	}

	return strings.Join([]string{"UntagDeviceResponse", string(data)}, " ")
}
