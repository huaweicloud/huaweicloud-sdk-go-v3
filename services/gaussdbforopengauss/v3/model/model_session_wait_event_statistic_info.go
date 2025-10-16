package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SessionWaitEventStatisticInfo struct {

	// **参数解释**: 组件名称。 **取值范围**: 不涉及。
	NodeName *string `json:"node_name,omitempty"`

	// **参数解释**: 等待事件名称。 **取值范围**: 不涉及。
	EventName *string `json:"event_name,omitempty"`

	// **参数解释**: 等待事件数量。 **取值范围**: 不涉及。
	Count *string `json:"count,omitempty"`
}

func (o SessionWaitEventStatisticInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SessionWaitEventStatisticInfo struct{}"
	}

	return strings.Join([]string{"SessionWaitEventStatisticInfo", string(data)}, " ")
}
