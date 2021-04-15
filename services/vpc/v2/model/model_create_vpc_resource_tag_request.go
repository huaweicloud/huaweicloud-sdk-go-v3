package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateVpcResourceTagRequest struct {
	VpcId string `json:"vpc_id"`

	Body *CreateVpcResourceTagRequestBody `json:"body,omitempty"`
}

func (o CreateVpcResourceTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateVpcResourceTagRequest struct{}"
	}

	return strings.Join([]string{"CreateVpcResourceTagRequest", string(data)}, " ")
}
