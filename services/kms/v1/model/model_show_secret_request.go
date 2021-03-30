package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowSecretRequest struct {
	SecretId string `json:"secret_id"`
}

func (o ShowSecretRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSecretRequest struct{}"
	}

	return strings.Join([]string{"ShowSecretRequest", string(data)}, " ")
}
