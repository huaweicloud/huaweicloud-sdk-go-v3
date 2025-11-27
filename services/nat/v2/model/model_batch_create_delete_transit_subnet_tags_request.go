package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateDeleteTransitSubnetTagsRequest Request Object
type BatchCreateDeleteTransitSubnetTagsRequest struct {

	// 中转子网的ID。
	ResourceId string `json:"resource_id"`

	Body *BatchOperateResourceTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateDeleteTransitSubnetTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDeleteTransitSubnetTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateDeleteTransitSubnetTagsRequest", string(data)}, " ")
}
