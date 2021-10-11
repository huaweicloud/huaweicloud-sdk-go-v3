package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteBackendInstanceV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBackendInstanceV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteBackendInstanceV2Response struct{}"
	}

	return strings.Join([]string{"DeleteBackendInstanceV2Response", string(data)}, " ")
}
