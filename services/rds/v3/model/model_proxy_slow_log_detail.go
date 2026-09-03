package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProxySlowLogDetail 数据库代理慢日志详情
type ProxySlowLogDetail struct {

	// **参数解释**：  客户端IP。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	SourceIp *string `json:"source_ip,omitempty"`

	// **参数解释**：  后端数据库IP。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	DescIp *string `json:"desc_ip,omitempty"`

	// **参数解释**：  数据库用户。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	User *string `json:"user,omitempty"`

	// **参数解释**：  响应时长，单位ms。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ReactionTime *string `json:"reaction_time,omitempty"`

	// **参数解释**：  SQL执行跟踪ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	TraceId *string `json:"trace_id,omitempty"`

	// **参数解释**：  执行语句。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Sql *string `json:"sql,omitempty"`

	// **参数解释**：  SQL语句执行开始时间，毫秒级时间戳。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**：  SQL语句执行结束时间，毫秒级时间戳。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**：  每次查询起始记录ID，用于获取后续数据时在请求中传入。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	LineNum *string `json:"line_num,omitempty"`

	// **参数解释**：  数据库名称。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Database *string `json:"database,omitempty"`

	// **参数解释**：  日志上报时间，毫秒级时间戳。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	LogTime *string `json:"log_time,omitempty"`
}

func (o ProxySlowLogDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxySlowLogDetail struct{}"
	}

	return strings.Join([]string{"ProxySlowLogDetail", string(data)}, " ")
}
