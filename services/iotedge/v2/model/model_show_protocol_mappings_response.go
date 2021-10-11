package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowProtocolMappingsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowProtocolMappingsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProtocolMappingsResponse struct{}"
	}

	return strings.Join([]string{"ShowProtocolMappingsResponse", string(data)}, " ")
}
