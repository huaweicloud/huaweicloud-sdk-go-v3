package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowSubnetTagsRequest struct {
	// 子网ID

	SubnetId string `json:"subnet_id"`
}

func (o ShowSubnetTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSubnetTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowSubnetTagsRequest", string(data)}, " ")
}
