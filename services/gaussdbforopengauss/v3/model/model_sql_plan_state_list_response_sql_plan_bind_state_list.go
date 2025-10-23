package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SqlPlanStateListResponseSqlPlanBindStateList struct {

	// **参数解释**: 当前计划的概览。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	Outline *string `json:"outline,omitempty"`

	// **参数解释**: SQL执行计划的开销。 **取值范围**: 不涉及。
	Cost *string `json:"cost,omitempty"`

	// **参数解释**: 绑定状态。 **取值范围**: - bind：绑定。 - drop：解绑。
	Status *string `json:"status,omitempty"`

	// **参数解释**: SQL文本的哈希值。 **取值范围**: 不涉及。
	SqlHash *string `json:"sql_hash,omitempty"`

	// **参数解释**: SQL执行计划ID。 **取值范围**: 不涉及。
	PlanId *string `json:"plan_id,omitempty"`
}

func (o SqlPlanStateListResponseSqlPlanBindStateList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlPlanStateListResponseSqlPlanBindStateList struct{}"
	}

	return strings.Join([]string{"SqlPlanStateListResponseSqlPlanBindStateList", string(data)}, " ")
}
