package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateSecretResponse struct {
	Secret         *Secret `json:"secret,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSecretResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateSecretResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecretResponse", string(data)}, " ")
}
