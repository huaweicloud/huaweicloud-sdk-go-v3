package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateSecretRequest struct {
	Body *CreateSecretRequestBody `json:"body,omitempty"`
}

func (o CreateSecretRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSecretRequest struct{}"
	}

	return strings.Join([]string{"CreateSecretRequest", string(data)}, " ")
}
