package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTopicReqQueues struct {

	// **参数解释**： 关联的代理。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Broker *string `json:"broker,omitempty"`

	// **参数解释**： 队列数。 **约束限制**： 不涉及。 **取值范围**： 1~50。 **默认取值**： 不涉及。
	QueueNum float32 `json:"queue_num,omitempty"`
}

func (o CreateTopicReqQueues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTopicReqQueues struct{}"
	}

	return strings.Join([]string{"CreateTopicReqQueues", string(data)}, " ")
}
