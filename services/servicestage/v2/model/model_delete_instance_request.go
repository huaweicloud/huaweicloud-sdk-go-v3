package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteInstanceRequest struct {
	ApplicationId string `json:"application_id"`
	ComponentId   string `json:"component_id"`
	InstanceId    string `json:"instance_id"`
	Force         *bool  `json:"force,omitempty"`
}

func (o DeleteInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceRequest", string(data)}, " ")
}
