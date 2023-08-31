package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceTypesRequest Request Object
type ListResourceTypesRequest struct {

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListResourceTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceTypesRequest struct{}"
	}

	return strings.Join([]string{"ListResourceTypesRequest", string(data)}, " ")
}
