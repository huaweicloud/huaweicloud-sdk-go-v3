package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSqlRecommendRulesResponseResult struct {

	// **参数解释**: 推荐类型。 **取值范围**: - all：全部 - exec_count：执行次数 - avg_exec_time：平均执行时间 - max_exec_time：最大执行时间
	RecommendType *string `json:"recommend_type,omitempty"`

	// **参数解释**: SQL ID。 **取值范围**: 不涉及。
	SqlId *string `json:"sql_id,omitempty"`

	// **参数解释**: SQL模板。 **取值范围**: 不涉及。
	SqlModel *string `json:"sql_model,omitempty"`

	// **参数解释**: SQL关键字。 **取值范围**: 不涉及。
	SqlKeyword *string `json:"sql_keyword,omitempty"`

	// **参数解释**: SQL类型。 **取值范围**: - SELECT - INSERT - UPDATE - DELETE - MERGE - OTHER
	SqlType *string `json:"sql_type,omitempty"`

	// **参数解释**: 数据库名称。 **取值范围**: 不涉及。
	Database *string `json:"database,omitempty"`

	// **参数解释**: 平均执行时间。 **取值范围**: 不涉及。
	AvgExecTime *float64 `json:"avg_exec_time,omitempty"`

	// **参数解释**: 最长执行时间。 **取值范围**: 不涉及。
	MaxExecTime *float64 `json:"max_exec_time,omitempty"`

	// **参数解释**: 执行次数。 **取值范围**: 不涉及。
	ExecCount *int32 `json:"exec_count,omitempty"`
}

func (o ListSqlRecommendRulesResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlRecommendRulesResponseResult struct{}"
	}

	return strings.Join([]string{"ListSqlRecommendRulesResponseResult", string(data)}, " ")
}
