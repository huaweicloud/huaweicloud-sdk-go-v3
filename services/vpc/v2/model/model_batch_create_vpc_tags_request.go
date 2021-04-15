package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchCreateVpcTagsRequest struct {
	VpcId string `json:"vpc_id"`

	Body *BatchCreateVpcTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateVpcTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateVpcTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateVpcTagsRequest", string(data)}, " ")
}
