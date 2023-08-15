package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateDeleteTransitIpTagsRequest Request Object
type BatchCreateDeleteTransitIpTagsRequest struct {

	// 中转IP的ID。
	ResourceId string `json:"resource_id"`

	Body *BatchOperateResourceTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateDeleteTransitIpTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDeleteTransitIpTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateDeleteTransitIpTagsRequest", string(data)}, " ")
}
