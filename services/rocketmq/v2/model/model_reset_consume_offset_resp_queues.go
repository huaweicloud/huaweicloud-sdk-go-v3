package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResetConsumeOffsetRespQueues struct {

	// **参数解释**： 队列所在的broker。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BrokerName *string `json:"broker_name,omitempty"`

	// **参数解释**： 队列ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	QueueId *int32 `json:"queue_id,omitempty"`

	// **参数解释**： 重置消费进度。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TimestampOffset *int64 `json:"timestamp_offset,omitempty"`
}

func (o ResetConsumeOffsetRespQueues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetConsumeOffsetRespQueues struct{}"
	}

	return strings.Join([]string{"ResetConsumeOffsetRespQueues", string(data)}, " ")
}
