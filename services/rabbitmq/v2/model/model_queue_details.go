package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueueDetails struct {

	// **参数解释**： Queue所属Vhost名称。 **取值范围**： 不涉及。
	Vhost *string `json:"vhost,omitempty"`

	// **参数解释**： Queue名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： Queue是否开启持久化。 **取值范围**： - true：开启持久化。 - false：未开启持久化。
	Durable *bool `json:"durable,omitempty"`

	// **参数解释**： Queue是否开启自动删除。 **取值范围**： - true：开启自动删除。 - false：未开启自动删除。
	AutoDelete *bool `json:"auto_delete,omitempty"`

	// **参数解释**： 待消费消息数。 **取值范围**： 不涉及。
	Messages *int32 `json:"messages,omitempty"`

	// **参数解释**： 连接的消费者数。 **取值范围**： 不涉及。
	Consumers *int32 `json:"consumers,omitempty"`

	// **参数解释**： 策略[（AMQP版本不支持policy，不涉及此参数）](tag:hws,hws_hk,,hws_eu,srg)。 **取值范围**： 不涉及。
	Policy *string `json:"policy,omitempty"`

	Arguments *QueueArguments `json:"arguments,omitempty"`
}

func (o QueueDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueueDetails struct{}"
	}

	return strings.Join([]string{"QueueDetails", string(data)}, " ")
}
