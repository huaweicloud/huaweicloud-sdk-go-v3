package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesByTagRequest Request Object
type ListResourcesByTagRequest struct {

	// 资源类型。
	ResourceType *ResourceType `json:"resource_type"`

	// 每页返回的个数。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量。
	Offset *int32 `json:"offset,omitempty"`

	Body *ListResourcesByTagRequestBody `json:"body,omitempty"`
}

func (o ListResourcesByTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesByTagRequest struct{}"
	}

	return strings.Join([]string{"ListResourcesByTagRequest", string(data)}, " ")
}
