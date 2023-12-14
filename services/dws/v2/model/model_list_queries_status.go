package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListQueriesStatus struct {

	// 平均查询等待时间。
	AverageQueryWaitingTime *float64 `json:"average_query_waiting_time,omitempty"`

	// 平均查询耗时。
	AverageTimeConsumptionOfQueries *float64 `json:"average_time_consumption_of_queries,omitempty"`

	// 平均会话耗时。
	AverageTimeConsumptionOfSessions *float64 `json:"average_time_consumption_of_sessions,omitempty"`

	// 查询数量。
	QueriesCount *int64 `json:"queries_count,omitempty"`

	// 会话数量。
	SessionCount *int64 `json:"session_count,omitempty"`
}

func (o ListQueriesStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueriesStatus struct{}"
	}

	return strings.Join([]string{"ListQueriesStatus", string(data)}, " ")
}
