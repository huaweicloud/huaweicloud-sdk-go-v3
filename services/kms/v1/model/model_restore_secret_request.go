package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RestoreSecretRequest struct {
	SecretId string `json:"secret_id"`
}

func (o RestoreSecretRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestoreSecretRequest struct{}"
	}

	return strings.Join([]string{"RestoreSecretRequest", string(data)}, " ")
}
