package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowVpcRequest struct {
	// 虚拟私有云ID

	VpcId string `json:"vpc_id"`
}

func (o ShowVpcRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowVpcRequest struct{}"
	}

	return strings.Join([]string{"ShowVpcRequest", string(data)}, " ")
}
