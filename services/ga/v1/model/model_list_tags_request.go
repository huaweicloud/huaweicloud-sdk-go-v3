package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagsRequest Request Object
type ListTagsRequest struct {

	// 资源类型。
	ResourceType *ResourceType `json:"resource_type"`

	// 每页返回的个数。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTagsRequest", string(data)}, " ")
}
