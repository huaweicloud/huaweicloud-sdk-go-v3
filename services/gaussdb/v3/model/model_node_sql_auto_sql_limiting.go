package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeSqlAutoSqlLimiting struct {

	// **参数解释**：  节点ID。  **取值范围**：  只能由英文字母、数字组成，前面为UUID，后缀为no07，长度为36个字符。  **默认取值**：  不涉及。
	NodeId *string `json:"node_id,omitempty"`

	// **参数解释**：  自治限流规则每天生效开始时间。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**：  自治限流规则每天生效开始时间。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**：  限流策略CPU利用率和活跃会话数的关联关系。  **取值范围**：    - and：且。   - or：或。  **默认取值**：  不涉及。
	Condition *string `json:"condition,omitempty"`

	// **参数解释**：  限流策略CPU利用率。  **约束限制**：  不涉及。  **取值范围**：  [70~100]。  **默认取值**：  不涉及。
	CpuUsage *int32 `json:"cpu_usage,omitempty"`

	// **参数解释**：  限流策略活跃会话数。  **约束限制**：  不涉及。  **取值范围**：  [1~5000]。  **默认取值**：  不涉及。
	ActiveSessions *int32 `json:"active_sessions,omitempty"`

	// **参数解释**：  每次最大限流时长（分钟）。  **取值范围**：  [1~1440]。  **默认取值**：  不涉及。
	ClearTime *int32 `json:"clear_time,omitempty"`

	// **参数解释**：  限流策略满足限流条件的事件持续时间（分钟）。  **取值范围**：  [2~60]。  **默认取值**：  不涉及。
	Duration *int32 `json:"duration,omitempty"`

	// **参数解释**：  最大并发数。  **取值范围**：  [0~1000000000]。  **默认取值**：  不涉及。
	MaxConcurrency *int32 `json:"max_concurrency,omitempty"`
}

func (o NodeSqlAutoSqlLimiting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeSqlAutoSqlLimiting struct{}"
	}

	return strings.Join([]string{"NodeSqlAutoSqlLimiting", string(data)}, " ")
}
