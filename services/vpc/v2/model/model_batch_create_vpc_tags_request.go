package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchCreateVpcTagsRequest struct {
	// 功能说明：虚拟私有云唯一标识 取值范围：合法UUID 约束：ID对应的VPC必须存在

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
