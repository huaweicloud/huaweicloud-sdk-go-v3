package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Message **参数解释**： 消息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type Message struct {

	// **参数解释**： 消息ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MsgId *string `json:"msg_id,omitempty"`

	// **参数解释**： 实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 主题名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Topic *string `json:"topic,omitempty"`

	// **参数解释**： 存储消息的时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StoreTimestamp float32 `json:"store_timestamp,omitempty"`

	// **参数解释**： 产生消息的时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BornTimestamp float32 `json:"born_timestamp,omitempty"`

	// **参数解释**： 重试次数。 **取值范围**： 不涉及。
	ReconsumeTimes *int32 `json:"reconsume_times,omitempty"`

	// **参数解释**： 消息体。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Body *string `json:"body,omitempty"`

	// **参数解释**： 消息体校验和。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BodyCrc float32 `json:"body_crc,omitempty"`

	// **参数解释**： 存储大小。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StoreSize float32 `json:"store_size,omitempty"`

	// **参数解释**： 消息属性列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PropertyList *[]MessagePropertyList `json:"property_list,omitempty"`

	// **参数解释**： 产生消息的主机IP。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BornHost *string `json:"born_host,omitempty"`

	// **参数解释**： 存储消息的主机IP。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StoreHost *string `json:"store_host,omitempty"`

	// **参数解释**： 队列ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	QueueId *int32 `json:"queue_id,omitempty"`

	// **参数解释**： 在队列中的偏移量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	QueueOffset *int32 `json:"queue_offset,omitempty"`
}

func (o Message) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Message struct{}"
	}

	return strings.Join([]string{"Message", string(data)}, " ")
}
