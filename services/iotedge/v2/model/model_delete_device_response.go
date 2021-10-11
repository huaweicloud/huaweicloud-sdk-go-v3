package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteDeviceResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDeviceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDeviceResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeviceResponse", string(data)}, " ")
}
