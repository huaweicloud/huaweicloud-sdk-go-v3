package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteSecretRequest struct {
	SecretId string `json:"secret_id"`
}

func (o DeleteSecretRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSecretRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecretRequest", string(data)}, " ")
}
