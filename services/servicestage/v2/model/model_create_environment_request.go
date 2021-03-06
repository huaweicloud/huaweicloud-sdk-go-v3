package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateEnvironmentRequest struct {
	Body *EnvironmentCreate `json:"body,omitempty"`
}

func (o CreateEnvironmentRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateEnvironmentRequest struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentRequest", string(data)}, " ")
}
