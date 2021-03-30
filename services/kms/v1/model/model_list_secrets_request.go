package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListSecretsRequest struct {
	Limit *string `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`
}

func (o ListSecretsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSecretsRequest struct{}"
	}

	return strings.Join([]string{"ListSecretsRequest", string(data)}, " ")
}
