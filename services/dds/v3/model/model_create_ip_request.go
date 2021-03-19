package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateIpRequest struct {
	InstanceId string `json:"instance_id"`

	Body *CreateIpRequestBody `json:"body,omitempty"`
}

func (o CreateIpRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateIpRequest struct{}"
	}

	return strings.Join([]string{"CreateIpRequest", string(data)}, " ")
}
