package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListSecretStageResponse struct {
	Stage          *Stage `json:"stage,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSecretStageResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSecretStageResponse struct{}"
	}

	return strings.Join([]string{"ListSecretStageResponse", string(data)}, " ")
}
