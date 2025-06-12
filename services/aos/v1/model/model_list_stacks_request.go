package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStacksRequest Request Object
type ListStacksRequest struct {

	// 用户指定的，对于此请求的唯一ID，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 分页标记。当一页无法返回所有结果，上一次的请求将返回next_marker以指引还有更多页数，用户可以将next_marker中的值放到此处以查询下一页的信息。此marker只能用于与上一请求指定的相同参数的请求。不指定时默认从第一页开始查询。
	Marker *string `json:"marker,omitempty"`

	// 每页返回的最多结果数量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListStacksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStacksRequest struct{}"
	}

	return strings.Join([]string{"ListStacksRequest", string(data)}, " ")
}
