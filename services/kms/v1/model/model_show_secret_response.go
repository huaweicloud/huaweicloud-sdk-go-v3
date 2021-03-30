package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowSecretResponse struct {
	Secret         *Secret `json:"secret,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSecretResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSecretResponse struct{}"
	}

	return strings.Join([]string{"ShowSecretResponse", string(data)}, " ")
}
