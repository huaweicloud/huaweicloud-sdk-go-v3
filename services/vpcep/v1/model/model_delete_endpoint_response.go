package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteEndpointResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEndpointResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteEndpointResponse struct{}"
	}

	return strings.Join([]string{"DeleteEndpointResponse", string(data)}, " ")
}
