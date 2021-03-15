package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UnfreezeDeviceResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UnfreezeDeviceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UnfreezeDeviceResponse struct{}"
	}

	return strings.Join([]string{"UnfreezeDeviceResponse", string(data)}, " ")
}
