package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListQueriesStatus struct {

	// **参数解释**： 平均查询等待时间。 **取值范围**： 不涉及。
	AverageQueryWaitingTime *float64 `json:"average_query_waiting_time,omitempty"`

	// **参数解释**： 平均查询耗时。 **取值范围**： 不涉及。
	AverageTimeConsumptionOfQueries *float64 `json:"average_time_consumption_of_queries,omitempty"`

	// **参数解释**： 平均会话耗时。 **取值范围**： 不涉及。
	AverageTimeConsumptionOfSessions *float64 `json:"average_time_consumption_of_sessions,omitempty"`

	// **参数解释**： 查询数量。 **取值范围**： 不涉及。
	QueriesCount *int64 `json:"queries_count,omitempty"`

	// **参数解释**： 会话数量。 **取值范围**： 不涉及。
	SessionCount *int64 `json:"session_count,omitempty"`
}

func (o ListQueriesStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueriesStatus struct{}"
	}

	return strings.Join([]string{"ListQueriesStatus", string(data)}, " ")
}
