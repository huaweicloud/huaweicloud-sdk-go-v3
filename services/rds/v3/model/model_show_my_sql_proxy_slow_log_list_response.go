package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMySqlProxySlowLogListResponse Response Object
type ShowMySqlProxySlowLogListResponse struct {

	// **参数解释**：  数据库代理慢日志信息列表。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	SlowLogList *[]ProxySlowLogDetail `json:"slow_log_list,omitempty"`

	// **参数解释**：  慢日志展示列表，该字段定义slow_log_list返回哪些字段信息，line_num字段一定返回。  **约束限制**：  不涉及。  **取值范围**：  - source_ip：客户端IP。 - desc_ip：后端数据库IP回。 - user：数据库用户。 - reaction_time：响应时长，单位ms。 - trace_id：SQL执行跟踪ID。 - sql：执行语句。 - start_time：SQL语句执行开始时间，毫秒级时间戳。 - end_time：SQL语句执行结束时间，毫秒级时间戳。 - database：数据库名称，默认不返回。 - log_time：日志上报时间，毫秒级时间戳，默认不返回。  **默认取值**：  不涉及。
	SlowLogColumn *[]string `json:"slow_log_column,omitempty"`

	// **参数解释**：  慢日志阈值，单位ms。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	SlowLogQueryTime *string `json:"slow_log_query_time,omitempty"`

	// **参数解释**：  慢日志上报开关状态。  **约束限制**：  不涉及。  **取值范围**：  - on：开启。 - off：关闭。  **默认取值**：  不涉及。
	LtsSlowLogEnabled *string `json:"lts_slow_log_enabled,omitempty"`

	// **参数解释**：  数据库代理版本是否支持慢日志上报。  **约束限制**：  不涉及。  **取值范围**：  - true：支持。 - false：不支持。  **默认取值**：  不涉及。
	SupportSwitchLtsSlowLog *bool `json:"support_switch_lts_slow_log,omitempty"`

	// **参数解释**：  每次查询到的慢日志数量。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	TotalCount     *string `json:"total_count,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMySqlProxySlowLogListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMySqlProxySlowLogListResponse struct{}"
	}

	return strings.Join([]string{"ShowMySqlProxySlowLogListResponse", string(data)}, " ")
}
