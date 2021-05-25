package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteSubnetTagRequest struct {
	// 子网ID

	SubnetId string `json:"subnet_id"`
	// 功能说明：键值

	Key string `json:"key"`
}

func (o DeleteSubnetTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSubnetTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubnetTagRequest", string(data)}, " ")
}
