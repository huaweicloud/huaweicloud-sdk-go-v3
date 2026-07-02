package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IntelligentKillSessionStatistic struct {

	// **参数解释**：  预览智能Kill会话SQL模板关键字。  **取值范围**：  不涉及。
	Keyword *string `json:"keyword,omitempty"`

	// **参数解释**：  预览智能Kill会话中的SQL模板命中的首个会话正在执行的SQL语句。  **取值范围**：  不涉及。
	ExampleSqlText *string `json:"example_sql_text,omitempty"`

	// **参数解释**：  预览智能Kill会话中的SQL模板命中的会话线程ID列表。
	Ids *[]int64 `json:"ids,omitempty"`

	// **参数解释**：  预览智能Kill会话中的SQL模板命中的会话个数。  **取值范围**：  >=0。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**：  预览智能Kill会话中的SQL模板命中的会话总执行时间，单位为秒。  **取值范围**：  >=0。
	TotalTime *float64 `json:"total_time,omitempty"`

	// **参数解释**：  预览智能Kill会话中的SQL模板命中的会话平均执行时间，单位为秒。  **取值范围**：  >=0。
	AvgTime *float64 `json:"avg_time,omitempty"`

	// **参数解释**：  预览智能Kill会话中的SQL模板命中的会话中最长会话执行时间，单位为秒。  **取值范围**：  >=0。
	MaxTime *float64 `json:"max_time,omitempty"`

	// **参数解释**：  预览智能Kill会话中的SQL模板命中Kill会话策略。  **取值范围**：  - top3_time: 以每组内会话最长的执行时长排序，选择排名前三的组内会话进行Kill。 - top3_count: 以每组内会话数量排序，选择排名前三的组内会话进行Kill。 - top3_avg_time: 以每组内会话平均执行时长排序，选择排名前三的组内会话进行Kill。
	Strategy *string `json:"strategy,omitempty"`
}

func (o IntelligentKillSessionStatistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntelligentKillSessionStatistic struct{}"
	}

	return strings.Join([]string{"IntelligentKillSessionStatistic", string(data)}, " ")
}
