package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FullSql struct {

	// SQL语句。
	Sql string `json:"sql"`

	// 操作类型。
	OperateType string `json:"operate_type"`

	// 状态。
	Status string `json:"status"`

	// 错误码。
	ErrorNo string `json:"error_no"`

	// 数据库名。
	Database string `json:"database"`

	// 客户端。
	Client string `json:"client"`

	// 线程ID。
	ThreadId string `json:"thread_id"`

	// 用户。
	User string `json:"user"`

	// 执行开始时间（Unix timestamp），单位：毫秒。
	ExecuteAt int64 `json:"execute_at"`

	// 执行耗时（毫秒）。
	QueryTime float64 `json:"query_time"`

	// 锁等待耗时（毫秒）。
	LockTime float64 `json:"lock_time"`

	// 扫描行数。
	RowsExamined int64 `json:"rows_examined"`

	// 返回行数。
	RowsSent int64 `json:"rows_sent"`

	// 更新行数。
	RowsAffected int64 `json:"rows_affected"`
}

func (o FullSql) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullSql struct{}"
	}

	return strings.Join([]string{"FullSql", string(data)}, " ")
}
