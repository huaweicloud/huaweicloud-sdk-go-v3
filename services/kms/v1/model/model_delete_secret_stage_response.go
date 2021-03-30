package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteSecretStageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecretStageResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSecretStageResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecretStageResponse", string(data)}, " ")
}
