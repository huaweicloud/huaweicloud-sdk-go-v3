package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubscriptionEntity struct {

	// **参数解释**： Topic名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TopicName *string `json:"topic_name,omitempty"`

	// **参数解释**： 消费者标签列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ConsumersInTags *[]ConsumersInTagEntity `json:"consumers_in_tags,omitempty"`
}

func (o SubscriptionEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionEntity struct{}"
	}

	return strings.Join([]string{"SubscriptionEntity", string(data)}, " ")
}
