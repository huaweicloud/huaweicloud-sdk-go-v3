package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FullSql struct {

	// SQL语句。
	Sql string `json:"sql" xml:"sql"`

	// 操作类型。
	OperateType string `json:"operate_type" xml:"operate_type"`

	// 状态。
	Status string `json:"status" xml:"status"`

	// 错误码。
	ErrorNo string `json:"error_no" xml:"error_no"`

	// 数据库名。
	Database string `json:"database" xml:"database"`

	// 客户端。
	Client string `json:"client" xml:"client"`

	// 线程ID。
	ThreadId string `json:"thread_id" xml:"thread_id"`

	// 用户。
	User string `json:"user" xml:"user"`

	// 执行开始时间（Unix timestamp），单位：毫秒。
	ExecuteAt int64 `json:"execute_at" xml:"execute_at"`

	// 执行耗时（毫秒）。
	QueryTime float64 `json:"query_time" xml:"query_time"`

	// 锁等待耗时（毫秒）。
	LockTime float64 `json:"lock_time" xml:"lock_time"`

	// 扫描行数。
	RowsExamined int64 `json:"rows_examined" xml:"rows_examined"`

	// 返回行数。
	RowsSent int64 `json:"rows_sent" xml:"rows_sent"`

	// 更新行数。
	RowsAffected int64 `json:"rows_affected" xml:"rows_affected"`
}

func (o FullSql) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullSql struct{}"
	}

	return strings.Join([]string{"FullSql", string(data)}, " ")
}
