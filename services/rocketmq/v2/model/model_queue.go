package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Queue struct {

	// **参数解释**： 队列ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释**： 队列消费堆积总数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Lag *int64 `json:"lag,omitempty"`

	// **参数解释**： 队列消息总数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BrokerOffset *int64 `json:"broker_offset,omitempty"`

	// **参数解释**： 已消费消息数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ConsumerOffset *int64 `json:"consumer_offset,omitempty"`

	// **参数解释**： 最新消费消息的存储时间，Unix毫秒时间戳格式。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	LastMessageTime *int64 `json:"last_message_time,omitempty"`
}

func (o Queue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Queue struct{}"
	}

	return strings.Join([]string{"Queue", string(data)}, " ")
}
