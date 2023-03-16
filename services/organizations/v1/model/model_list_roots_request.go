package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRootsRequest struct {

	// 页面中最大结果数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListRootsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRootsRequest struct{}"
	}

	return strings.Join([]string{"ListRootsRequest", string(data)}, " ")
}
