package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowTopicStatusRespQueues struct {

	// **参数解释**： 队列ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释**： 最小偏移量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MinOffset *int32 `json:"min_offset,omitempty"`

	// **参数解释**： 最大偏移量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MaxOffset *int32 `json:"max_offset,omitempty"`

	// **参数解释**： 最后一条消息的时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	LastMessageTime *int64 `json:"last_message_time,omitempty"`
}

func (o ShowTopicStatusRespQueues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopicStatusRespQueues struct{}"
	}

	return strings.Join([]string{"ShowTopicStatusRespQueues", string(data)}, " ")
}
