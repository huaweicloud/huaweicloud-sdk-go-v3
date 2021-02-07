package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteDomainMappingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDomainMappingResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDomainMappingResponse struct{}"
	}

	return strings.Join([]string{"DeleteDomainMappingResponse", string(data)}, " ")
}
