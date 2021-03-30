package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateSecretStageRequest struct {
	SecretId string `json:"secret_id"`

	StageName string `json:"stage_name"`

	Body *UpdateSecretStageRequestBody `json:"body,omitempty"`
}

func (o UpdateSecretStageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateSecretStageRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecretStageRequest", string(data)}, " ")
}
