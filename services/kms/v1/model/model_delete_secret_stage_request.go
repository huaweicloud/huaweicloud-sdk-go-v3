package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteSecretStageRequest struct {
	SecretId string `json:"secret_id"`

	StageName string `json:"stage_name"`
}

func (o DeleteSecretStageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSecretStageRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecretStageRequest", string(data)}, " ")
}
