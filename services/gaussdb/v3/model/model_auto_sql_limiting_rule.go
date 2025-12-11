package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutoSqlLimitingRule struct {

	// **参数解释**：  节点ID。  获取方法请参见[查询实例详情](https://support.huaweicloud.com/api-taurusdb/ShowGaussMySqlInstanceInfoUnifyStatus.html)。  **约束限制**：  节点角色必须为主节点。  **取值范围**：  只能由英文字母、数字组成，前面为UUID，后缀为no07，长度为36个字符。  **默认取值**：  不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**：  自治限流规则每天生效开始时间。  **约束限制**：  格式为 \"hh:mm\"。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	StartTime string `json:"start_time"`

	// **参数解释**：  自治限流规则每天生效结束时间。  **约束限制**：  格式为 \"hh:mm\"。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EndTime string `json:"end_time"`

	// **参数解释**：  限流策略CPU利用率和活跃会话数的关联关系。  **约束限制**：  不涉及。  **取值范围**：    - and：且。   - or：或。  **默认取值**：  不涉及。
	Condition string `json:"condition"`

	// **参数解释**：  限流策略CPU利用率。  **约束限制**：  不涉及。  **取值范围**：  [70~100]。  **默认取值**：  不涉及。
	CpuUsage int32 `json:"cpu_usage"`

	// **参数解释**：  限流策略活跃会话数。  **约束限制**：  不涉及。  **取值范围**：  [1~5000]。  **默认取值**：  不涉及。
	ActiveSessions int32 `json:"active_sessions"`

	// **参数解释**：  每次最大限流时长（分钟）。  **约束限制**：  不涉及。  **取值范围**：  [1~1440]。  **默认取值**：  不涉及。
	ClearTime int32 `json:"clear_time"`

	// **参数解释**：  限流策略满足限流条件的事件持续时间（分钟）。  **约束限制**：  不涉及。  **取值范围**：  [2~60]。  **默认取值**：  不涉及。
	Duration int32 `json:"duration"`

	// **参数解释**：  最大并发数。  **约束限制**：  不涉及。  **取值范围**：  [0~1000000000]。  **默认取值**：  不涉及。
	MaxConcurrency int32 `json:"max_concurrency"`

	// **参数解释**：  是否保留已有的SQL限流规则。  **约束限制**：  不涉及。  **取值范围**：    - true：保留已有的SQL限流规则。   - false：清理已有的SQL限流规则。  **默认取值**：  不涉及。
	RetainSqlRule bool `json:"retain_sql_rule"`
}

func (o AutoSqlLimitingRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoSqlLimitingRule struct{}"
	}

	return strings.Join([]string{"AutoSqlLimitingRule", string(data)}, " ")
}
