package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteEndpointServiceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEndpointServiceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteEndpointServiceResponse struct{}"
	}

	return strings.Join([]string{"DeleteEndpointServiceResponse", string(data)}, " ")
}
