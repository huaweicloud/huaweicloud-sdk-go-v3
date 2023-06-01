package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEventStreamingRequest struct {

	// 偏移量，表示从此偏移量开始查询，偏移量不能小于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，不能小于1或大于1000
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListEventStreamingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventStreamingRequest struct{}"
	}

	return strings.Join([]string{"ListEventStreamingRequest", string(data)}, " ")
}
