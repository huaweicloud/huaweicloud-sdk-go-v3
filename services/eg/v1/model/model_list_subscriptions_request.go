package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSubscriptionsRequest struct {

	// 事件通道ID
	ChannelId *string `json:"channel_id,omitempty"`

	// 偏移量，表示从此偏移量开始查询，偏移量不能小于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，不能小于0。
	Limit *int32 `json:"limit,omitempty"`

	// 指定查询排序
	Sort *string `json:"sort,omitempty"`

	// 指定查询的事件订阅名称，精准匹配
	Name *string `json:"name,omitempty"`

	// 指定查询的事件订阅名称，模糊匹配
	FuzzyName *string `json:"fuzzy_name,omitempty"`
}

func (o ListSubscriptionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionsRequest struct{}"
	}

	return strings.Join([]string{"ListSubscriptionsRequest", string(data)}, " ")
}
