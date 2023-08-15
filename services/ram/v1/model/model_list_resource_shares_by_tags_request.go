package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceSharesByTagsRequest Request Object
type ListResourceSharesByTagsRequest struct {

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Offset *string `json:"offset,omitempty"`

	Body *ResourceSharesByTagsReqBody `json:"body,omitempty"`
}

func (o ListResourceSharesByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceSharesByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListResourceSharesByTagsRequest", string(data)}, " ")
}
