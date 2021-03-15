package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type FreezeDeviceResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o FreezeDeviceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "FreezeDeviceResponse struct{}"
	}

	return strings.Join([]string{"FreezeDeviceResponse", string(data)}, " ")
}
