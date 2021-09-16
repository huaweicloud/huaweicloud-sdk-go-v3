package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteSubnetRequest struct {
	// 子网ID。

	SubnetId string `json:"subnet_id"`
}

func (o DeleteSubnetRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSubnetRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubnetRequest", string(data)}, " ")
}
