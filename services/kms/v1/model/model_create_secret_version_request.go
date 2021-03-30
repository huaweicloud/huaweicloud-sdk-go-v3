package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateSecretVersionRequest struct {
	SecretId string `json:"secret_id"`

	Body *CreateSecretVersionRequestBody `json:"body,omitempty"`
}

func (o CreateSecretVersionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSecretVersionRequest struct{}"
	}

	return strings.Join([]string{"CreateSecretVersionRequest", string(data)}, " ")
}
