package model

import (
	"encoding/json"

	"strings"
)

//
type KeystoneUpdateMappingRequestBody struct {
	Mapping *MappingOption `json:"mapping"`
}

func (o KeystoneUpdateMappingRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneUpdateMappingRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateMappingRequestBody", string(data)}, " ")
}
