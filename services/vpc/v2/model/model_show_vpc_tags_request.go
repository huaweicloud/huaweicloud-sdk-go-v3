package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowVpcTagsRequest struct {
	VpcId string `json:"vpc_id"`
}

func (o ShowVpcTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowVpcTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowVpcTagsRequest", string(data)}, " ")
}
