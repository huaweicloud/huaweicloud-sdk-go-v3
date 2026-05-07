package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgentStatusInfo struct {

	// **参数解释**: Agent状态 **取值范围**: 包含如下3种。 - online：在线。 - offline：离线。 - agent_protect_interrupted：防护中断。
	AgentStatus *string `json:"agent_status,omitempty"`

	// **参数解释**： agent状态时间，采用时间戳，默认毫秒 **取值范围**： 0-4824695185000
	StatusTime *int64 `json:"status_time,omitempty"`

	// **参数解释**： 异常原因 **取值范围**： 字符长度0-512位
	AbnormalReason *string `json:"abnormal_reason,omitempty"`
}

func (o AgentStatusInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentStatusInfo struct{}"
	}

	return strings.Join([]string{"AgentStatusInfo", string(data)}, " ")
}
