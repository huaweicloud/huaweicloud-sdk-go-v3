package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateSubnetRequest struct {
	// 子网对应的vpc_id

	VpcId string `json:"vpc_id"`
	// 子网ID

	SubnetId string `json:"subnet_id"`

	Body *UpdateSubnetRequestBody `json:"body,omitempty"`
}

func (o UpdateSubnetRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateSubnetRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubnetRequest", string(data)}, " ")
}
