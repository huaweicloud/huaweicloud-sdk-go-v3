package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceConsumerGroupsResponse Response Object
type ListInstanceConsumerGroupsResponse struct {

	// **参数解释**： 消费组总数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Total float32 `json:"total,omitempty"`

	// **参数解释**： 消费组列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Groups *[]ConsumerGroup `json:"groups,omitempty"`

	// **参数解释**： 最大可创建消费组数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Max *int32 `json:"max,omitempty"`

	// **参数解释**： 剩余可创建消费组数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Remaining *int32 `json:"remaining,omitempty"`

	// **参数解释**： 下个分页的offset。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NextOffset *int32 `json:"next_offset,omitempty"`

	// **参数解释**： 上个分页的offset。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PreviousOffset *int32 `json:"previous_offset,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceConsumerGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceConsumerGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceConsumerGroupsResponse", string(data)}, " ")
}
