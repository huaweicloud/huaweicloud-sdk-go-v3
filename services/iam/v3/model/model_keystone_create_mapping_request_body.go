package model

import (
	"encoding/json"

	"strings"
)

//
type KeystoneCreateMappingRequestBody struct {
	Mapping *MappingOption `json:"mapping"`
}

func (o KeystoneCreateMappingRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateMappingRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneCreateMappingRequestBody", string(data)}, " ")
}
