package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteMfaDeviceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMfaDeviceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteMfaDeviceResponse struct{}"
	}

	return strings.Join([]string{"DeleteMfaDeviceResponse", string(data)}, " ")
}
