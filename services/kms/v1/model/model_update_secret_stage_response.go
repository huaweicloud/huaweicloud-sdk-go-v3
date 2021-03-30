package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateSecretStageResponse struct {
	Stage          *Stage `json:"stage,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateSecretStageResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateSecretStageResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecretStageResponse", string(data)}, " ")
}
