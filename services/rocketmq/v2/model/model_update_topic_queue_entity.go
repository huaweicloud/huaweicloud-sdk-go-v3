package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTopicQueueEntity struct {

	// **参数解释**： Broker名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Broker *string `json:"broker,omitempty"`

	// **参数解释**： 读队列个数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ReadQueueNum float32 `json:"read_queue_num,omitempty"`

	// **参数解释**： 写队列个数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	WriteQueueNum float32 `json:"write_queue_num,omitempty"`
}

func (o UpdateTopicQueueEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicQueueEntity struct{}"
	}

	return strings.Join([]string{"UpdateTopicQueueEntity", string(data)}, " ")
}
