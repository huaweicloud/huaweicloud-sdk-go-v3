package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConsumerListOrDetailsResponse Response Object
type ShowConsumerListOrDetailsResponse struct {

	// **参数解释**： Topic列表（当查询Topic消费“列表”时才显示此参数）。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Topics *[]string `json:"topics,omitempty"`

	// **参数解释**： Topic总数（当查询Topic消费“列表”时才显示此参数）。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 消费堆积总数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Lag *int64 `json:"lag,omitempty"`

	// **参数解释**： 消息总数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MaxOffset *int64 `json:"max_offset,omitempty"`

	// **参数解释**： 已消费消息数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ConsumerOffset *int64 `json:"consumer_offset,omitempty"`

	// **参数解释**： Topic关联代理（当查询Topic消费“详情”才显示此参数）。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Brokers        *[]Brokers `json:"brokers,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowConsumerListOrDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsumerListOrDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowConsumerListOrDetailsResponse", string(data)}, " ")
}
