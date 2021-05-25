package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchDeleteSubnetTagsRequest struct {
	// 子网ID

	SubnetId string `json:"subnet_id"`

	Body *BatchDeleteSubnetTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteSubnetTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteSubnetTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubnetTagsRequest", string(data)}, " ")
}
