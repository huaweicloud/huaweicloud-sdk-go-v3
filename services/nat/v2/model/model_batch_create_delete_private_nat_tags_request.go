package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateDeletePrivateNatTagsRequest Request Object
type BatchCreateDeletePrivateNatTagsRequest struct {

	// 私网NAT网关的ID。
	ResourceId string `json:"resource_id"`

	Body *BatchOperateResourceTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateDeletePrivateNatTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDeletePrivateNatTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateDeletePrivateNatTagsRequest", string(data)}, " ")
}
