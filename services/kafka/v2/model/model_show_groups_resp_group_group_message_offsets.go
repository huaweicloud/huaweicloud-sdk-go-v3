package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowGroupsRespGroupGroupMessageOffsets struct {

	// **参数解释**： 分区编号。 **取值范围**： 不涉及。
	Partition *int32 `json:"partition,omitempty"`

	// **参数解释**： 剩余可消费消息数，即消息堆积数。 **取值范围**： 不涉及。
	Lag *int64 `json:"lag,omitempty"`

	// **参数解释**： Topic名称。 **取值范围**： 不涉及。
	Topic *string `json:"topic,omitempty"`

	// **参数解释**： 当前消费进度。 **取值范围**： 不涉及。
	MessageCurrentOffset *int64 `json:"message_current_offset,omitempty"`

	// **参数解释**： 最大消息位置（LEO）。 **取值范围**： 不涉及。
	MessageLogEndOffset *int64 `json:"message_log_end_offset,omitempty"`
}

func (o ShowGroupsRespGroupGroupMessageOffsets) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupsRespGroupGroupMessageOffsets struct{}"
	}

	return strings.Join([]string{"ShowGroupsRespGroupGroupMessageOffsets", string(data)}, " ")
}
