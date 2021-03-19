package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowInstanceDetailRequest struct {
	ApplicationId string `json:"application_id"`

	ComponentId string `json:"component_id"`

	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceDetailRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowInstanceDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceDetailRequest", string(data)}, " ")
}
