package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteBindingDeviceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBindingDeviceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteBindingDeviceResponse struct{}"
	}

	return strings.Join([]string{"DeleteBindingDeviceResponse", string(data)}, " ")
}
