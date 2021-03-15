package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type TagDeviceResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o TagDeviceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TagDeviceResponse struct{}"
	}

	return strings.Join([]string{"TagDeviceResponse", string(data)}, " ")
}
