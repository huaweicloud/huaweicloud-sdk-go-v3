package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlPlanActionResponse Response Object
type ListSqlPlanActionResponse struct {

	// **参数解释**: sql执行计划总数。 **取值范围**: 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**: 执行计划列表。 **取值范围**: 不涉及。
	SqlPlanBindStateList *[]SqlPlanStateListResponseSqlPlanBindStateList `json:"sql_plan_bind_state_list,omitempty"`
	HttpStatusCode       int                                             `json:"-"`
}

func (o ListSqlPlanActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlPlanActionResponse struct{}"
	}

	return strings.Join([]string{"ListSqlPlanActionResponse", string(data)}, " ")
}
