package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ChangeEnvironmentRequest struct {
	EnvironmentId string             `json:"environment_id"`
	Body          *EnvironmentModify `json:"body,omitempty"`
}

func (o ChangeEnvironmentRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChangeEnvironmentRequest struct{}"
	}

	return strings.Join([]string{"ChangeEnvironmentRequest", string(data)}, " ")
}
