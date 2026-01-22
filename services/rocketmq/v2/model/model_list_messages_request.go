package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMessagesRequest Request Object
type ListMessagesRequest struct {

	// **参数解释**： 消息引擎。 **约束限制**： 不涉及。 **取值范围**： - rocketmq：RocketMQ消息引擎。 - reliability：RocketMQ消息引擎别称。 **默认取值**： 不涉及。
	Engine string `json:"engine"`

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 主题名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Topic string `json:"topic"`

	// **参数解释**： 队列。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Queue *string `json:"queue,omitempty"`

	// **参数解释**： 查询数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 消息的key。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释**： 开始时间。 **约束限制**： 不通过msg_id精确查询消息时，此参数必填。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 结束时间。 **约束限制**： 不通过msg_id精确查询消息时，此参数必填。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 消息ID。 **约束限制**： 不通过时间范围查询消息时，此参数必填。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MsgId *string `json:"msg_id,omitempty"`
}

func (o ListMessagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessagesRequest struct{}"
	}

	return strings.Join([]string{"ListMessagesRequest", string(data)}, " ")
}
