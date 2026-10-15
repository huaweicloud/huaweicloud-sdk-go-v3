package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReplayExecuteRecords 流量回放SQL的执行记录
type ReplayExecuteRecords struct {

	// 执行记录的时间，格式为秒级时间戳。
	ExecuteTime *int64 `json:"execute_time,omitempty"`

	// 执行完的SQL数量。
	FinishedSql *int64 `json:"finished_sql,omitempty"`

	// 执行的异常SQL数量。
	AbnormalSql *int64 `json:"abnormal_sql,omitempty"`

	// 执行的慢SQL数量。
	SlowSql *int64 `json:"slow_sql,omitempty"`

	// 执行记录的SQL总数。
	TotalSql *int64 `json:"total_sql,omitempty"`
}

func (o ReplayExecuteRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplayExecuteRecords struct{}"
	}

	return strings.Join([]string{"ReplayExecuteRecords", string(data)}, " ")
}
