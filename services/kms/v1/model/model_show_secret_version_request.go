package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowSecretVersionRequest struct {
	SecretId string `json:"secret_id"`

	VersionId string `json:"version_id"`
}

func (o ShowSecretVersionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSecretVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowSecretVersionRequest", string(data)}, " ")
}
