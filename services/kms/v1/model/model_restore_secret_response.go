package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RestoreSecretResponse struct {
	Secret         *Secret `json:"secret,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreSecretResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestoreSecretResponse struct{}"
	}

	return strings.Join([]string{"RestoreSecretResponse", string(data)}, " ")
}
