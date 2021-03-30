package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateSecretResponse struct {
	Secret         *Secret `json:"secret,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSecretResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSecretResponse struct{}"
	}

	return strings.Join([]string{"CreateSecretResponse", string(data)}, " ")
}
