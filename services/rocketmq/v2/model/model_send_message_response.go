package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendMessageResponse Response Object
type SendMessageResponse struct {

	// **参数解释**： 主题名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	Topic *string `json:"topic,omitempty"`

	// **参数解释**： 消息内容。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	Body *string `json:"body,omitempty"`

	// **参数解释**： 特性列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	PropertyList *[]SendMessageProperties `json:"property_list,omitempty"`

	// **参数解释**： 消息ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	MsgId *string `json:"msg_id,omitempty"`

	// **参数解释**： 队列ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	QueueId float32 `json:"queue_id,omitempty"`

	// **参数解释**： 队列offset。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	QueueOffset float32 `json:"queue_offset,omitempty"`

	// **参数解释**： Broker名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	BrokerName     *string `json:"broker_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SendMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendMessageResponse struct{}"
	}

	return strings.Join([]string{"SendMessageResponse", string(data)}, " ")
}
