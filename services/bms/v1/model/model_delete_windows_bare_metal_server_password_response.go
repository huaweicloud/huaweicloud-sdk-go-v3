package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteWindowsBareMetalServerPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWindowsBareMetalServerPasswordResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteWindowsBareMetalServerPasswordResponse struct{}"
	}

	return strings.Join([]string{"DeleteWindowsBareMetalServerPasswordResponse", string(data)}, " ")
}
