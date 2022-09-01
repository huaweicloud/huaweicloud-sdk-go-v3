package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlowLog struct {

	// SQL语句。
	Sql string `json:"sql" xml:"sql"`

	// 数据库名。
	Database string `json:"database" xml:"database"`

	// 客户端。
	Client string `json:"client" xml:"client"`

	// 用户。
	User string `json:"user" xml:"user"`

	// 执行开始时间（Unix timestamp），单位：毫秒。
	ExecuteAt int64 `json:"execute_at" xml:"execute_at"`

	// 执行耗时（秒）。
	QueryTime float64 `json:"query_time" xml:"query_time"`

	// 锁等待耗时（秒）。
	LockTime float64 `json:"lock_time" xml:"lock_time"`

	// 扫描行数。
	RowsExamined int64 `json:"rows_examined" xml:"rows_examined"`

	// 返回行数。
	RowsSent int64 `json:"rows_sent" xml:"rows_sent"`
}

func (o SlowLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowLog struct{}"
	}

	return strings.Join([]string{"SlowLog", string(data)}, " ")
}
