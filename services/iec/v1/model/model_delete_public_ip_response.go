package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeletePublicIpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePublicIpResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePublicIpResponse struct{}"
	}

	return strings.Join([]string{"DeletePublicIpResponse", string(data)}, " ")
}
