package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListInstanceConsumerGroupsResponse struct {

	// 消费组总数。
	Total float32 `json:"total,omitempty" xml:"total"`

	// 消费组列表。
	Groups *[]ConsumerGroup `json:"groups,omitempty" xml:"groups"`

	// 最大可创建消费组数量。
	Max float32 `json:"max,omitempty" xml:"max"`

	// 剩余可创建消费组数量。
	Remaining float32 `json:"remaining,omitempty" xml:"remaining"`

	// 下个分页的offset。
	NextOffset float32 `json:"next_offset,omitempty" xml:"next_offset"`

	// 上个分页的offset。
	PreviousOffset float32 `json:"previous_offset,omitempty" xml:"previous_offset"`
	HttpStatusCode int     `json:"-"`
}

func (o ListInstanceConsumerGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceConsumerGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceConsumerGroupsResponse", string(data)}, " ")
}
