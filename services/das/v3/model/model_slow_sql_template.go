package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlowSqlTemplate struct {

	// SQL模板。
	SqlTemplate string `json:"sql_template"`

	// SQL样本。
	SqlSample *string `json:"sql_sample,omitempty"`

	// SQL样本执行用户。
	SqlSampleUser *string `json:"sql_sample_user,omitempty"`

	// 库名。
	DbNames []string `json:"db_names"`

	// 执行次数。
	ExecuteCount int64 `json:"execute_count"`

	// 平均执行耗时(ms)。
	AvgExecuteTime float64 `json:"avg_execute_time"`

	// 最大执行耗时(ms)。
	MaxExecuteTime float64 `json:"max_execute_time"`

	// 平均锁等待时间(ms)。
	AvgLockWaitTime float64 `json:"avg_lock_wait_time"`

	// 最大锁等待时间(ms)。
	MaxLockWaitTime float64 `json:"max_lock_wait_time"`

	// 平均扫描行数。
	AvgRowsExamined float64 `json:"avg_rows_examined"`

	// 最大扫描行数。
	MaxRowsExamined float64 `json:"max_rows_examined"`

	// 平均返回行数。
	AvgRowsSent float64 `json:"avg_rows_sent"`

	// 最大返回行数。
	MaxRowsSent float64 `json:"max_rows_sent"`
}

func (o SlowSqlTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowSqlTemplate struct{}"
	}

	return strings.Join([]string{"SlowSqlTemplate", string(data)}, " ")
}
