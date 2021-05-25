package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowSubnetRequest struct {
	// 子网ID

	SubnetId string `json:"subnet_id"`
}

func (o ShowSubnetRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSubnetRequest struct{}"
	}

	return strings.Join([]string{"ShowSubnetRequest", string(data)}, " ")
}
