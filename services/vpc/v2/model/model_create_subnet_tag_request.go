package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateSubnetTagRequest struct {
	SubnetId string `json:"subnet_id"`

	Body *CreateSubnetTagRequestBody `json:"body,omitempty"`
}

func (o CreateSubnetTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSubnetTagRequest struct{}"
	}

	return strings.Join([]string{"CreateSubnetTagRequest", string(data)}, " ")
}
