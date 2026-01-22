package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateOrUpdateConsumerGroup struct {

	// **参数解释**： 消费组名称。 **约束限制**： 只能由英文字母、数字、百分号、竖线、中划线、下划线组成，长度3~64个字符。 当创建消费组时必填。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 关联的代理列表。 **约束限制**： 仅RocketMQ实例4.8.0版本需要填写此参数。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Brokers *[]string `json:"brokers,omitempty"`

	// **参数解释**： 是否设置为广播消费。 **约束限制**： 不涉及。 **取值范围**： - true：使用广播消费。 - false：不使用广播消费。 **默认取值**： 不涉及。
	Broadcast *bool `json:"broadcast,omitempty"`

	// **参数解释**： 最大重试次数。 **约束限制**： 不涉及。 **取值范围**： 1~16。 **默认取值**： 不涉及。
	RetryMaxTime *int32 `json:"retry_max_time,omitempty"`

	// **参数解释**： 是否可以消费。 **约束限制**： 不涉及。 **取值范围**： - true：可以消费。 - false：不可以消费。 **默认取值**： 不涉及。
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**： 是否按顺序消费。 **约束限制**： 仅RocketMQ实例5.x版本需要填写此参数。[华为云Stack不支持此参数。](tag:hcs,hcs_oemout) **取值范围**： - true：顺序消费。 - false：不按顺序消费。 **默认取值**： 不涉及。
	ConsumeOrderly *bool `json:"consume_orderly,omitempty"`

	// **参数解释**： 消费组描述。 **约束限制**： 不涉及。 **取值范围**： 0~200。 **默认取值**： 不涉及。
	GroupDesc *string `json:"group_desc,omitempty"`
}

func (o CreateOrUpdateConsumerGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateConsumerGroup struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateConsumerGroup", string(data)}, " ")
}
