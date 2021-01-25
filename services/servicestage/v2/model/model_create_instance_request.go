package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateInstanceRequest struct {
	ApplicationId string          `json:"application_id"`
	ComponentId   string          `json:"component_id"`
	Body          *InstanceCreate `json:"body,omitempty"`
}

func (o CreateInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceRequest", string(data)}, " ")
}
