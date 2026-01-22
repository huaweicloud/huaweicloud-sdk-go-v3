package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetConsumeOffsetForRocketMqResponse Response Object
type ResetConsumeOffsetForRocketMqResponse struct {

	// **参数解释**： 重置的队列。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Queues         *[]ResetConsumeOffsetRespQueues `json:"queues,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ResetConsumeOffsetForRocketMqResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetConsumeOffsetForRocketMqResponse struct{}"
	}

	return strings.Join([]string{"ResetConsumeOffsetForRocketMqResponse", string(data)}, " ")
}
