package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteVpcRequest struct {
	// 虚拟私有云ID。

	VpcId string `json:"vpc_id"`
}

func (o DeleteVpcRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteVpcRequest struct{}"
	}

	return strings.Join([]string{"DeleteVpcRequest", string(data)}, " ")
}
