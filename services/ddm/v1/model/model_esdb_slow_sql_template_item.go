package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EsdbSlowSqlTemplateItem struct {

	// **参数解释**：  示例。  **参数范围**：  不涉及。
	QuerySample *string `json:"query_sample,omitempty"`

	// **参数解释**：  数据库。  **参数范围**：  不涉及。
	Database *string `json:"database,omitempty"`

	// **参数解释**：  执行时间。  **参数范围**：  不涉及。
	ExecuteTimes float32 `json:"execute_times,omitempty"`

	// **参数解释**：  平均检查行数。  **参数范围**：  不涉及。
	AvgRowsExamined float32 `json:"avg_rows_examined,omitempty"`

	// **参数解释**：  最大时间。  **参数范围**：  不涉及。
	MaxTime float32 `json:"max_time,omitempty"`

	// **参数解释**：  平均时间。  **参数范围**：  不涉及。
	AvgTime float32 `json:"avg_time,omitempty"`

	// **参数解释**：  最大检查行数。  **参数范围**：  不涉及。
	MaxRowsExamined float32 `json:"max_rows_examined,omitempty"`

	// **参数解释**：  总检查行数。  **参数范围**：  不涉及。
	RowsExaminedSum float32 `json:"rows_examined_sum,omitempty"`

	// **参数解释**：  总耗时。  **参数范围**：  不涉及。
	CostTimeSum float32 `json:"cost_time_sum,omitempty"`
}

func (o EsdbSlowSqlTemplateItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EsdbSlowSqlTemplateItem struct{}"
	}

	return strings.Join([]string{"EsdbSlowSqlTemplateItem", string(data)}, " ")
}
