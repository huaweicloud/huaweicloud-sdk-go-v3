package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListResourceShareTagsRequest struct {

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListResourceShareTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceShareTagsRequest struct{}"
	}

	return strings.Join([]string{"ListResourceShareTagsRequest", string(data)}, " ")
}
