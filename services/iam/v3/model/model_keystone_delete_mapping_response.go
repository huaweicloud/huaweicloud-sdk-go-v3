package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneDeleteMappingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneDeleteMappingResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneDeleteMappingResponse struct{}"
	}

	return strings.Join([]string{"KeystoneDeleteMappingResponse", string(data)}, " ")
}
