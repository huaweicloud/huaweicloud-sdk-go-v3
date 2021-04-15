package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteVpcTagRequest struct {
	VpcId string `json:"vpc_id"`

	Key string `json:"key"`
}

func (o DeleteVpcTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteVpcTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteVpcTagRequest", string(data)}, " ")
}
