package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceConsumerGroupsResponse Response Object
type ListInstanceConsumerGroupsResponse struct {

	// 消费组总数。
	Total float32 `json:"total,omitempty"`

	// 消费组列表。
	Groups *[]ConsumerGroup `json:"groups,omitempty"`

	// 最大可创建消费组数量。
	Max float32 `json:"max,omitempty"`

	// 剩余可创建消费组数量。
	Remaining float32 `json:"remaining,omitempty"`

	// 下个分页的offset。
	NextOffset float32 `json:"next_offset,omitempty"`

	// 上个分页的offset。
	PreviousOffset float32 `json:"previous_offset,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListInstanceConsumerGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceConsumerGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceConsumerGroupsResponse", string(data)}, " ")
}
