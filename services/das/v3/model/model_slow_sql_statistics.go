package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlowSqlStatistics struct {

	// 执行次数。
	ExecuteCount int64 `json:"execute_count"`

	// 平均执行耗时(s)。
	AvgExecuteTime float64 `json:"avg_execute_time"`

	// 最大执行耗时(s)。
	MaxExecuteTime float64 `json:"max_execute_time"`

	// 平均锁等待时间(s)。
	AvgLockWaitTime float64 `json:"avg_lock_wait_time"`

	// 最大锁等待时间(s)。
	MaxLockWaitTime float64 `json:"max_lock_wait_time"`

	// 平均返回文档数。
	AvgRowsSent float64 `json:"avg_rows_sent"`

	// 最大返回文档数。
	MaxRowsSent float64 `json:"max_rows_sent"`

	// 平均扫描文档数。
	AvgRowsExamined float64 `json:"avg_rows_examined"`

	// 最大扫描文档数。
	MaxRowsExamined float64 `json:"max_rows_examined"`

	// 平均扫描索引数。
	AvgKeyExamined float64 `json:"avg_key_examined"`

	// 最大扫描索引数。
	MaxKeyExamined float64 `json:"max_key_examined"`

	// 节点ID，按node_id统计时赋值。
	NodeId *string `json:"node_id,omitempty"`

	// 节点名称，按node_id统计时赋值。
	NodeName *string `json:"node_name,omitempty"`

	// 语句类型，按sql_type统计时赋值。
	SqlType *string `json:"sql_type,omitempty"`

	// 库名，按db_name、collection统计时赋值。
	DbName *string `json:"db_name,omitempty"`

	// 数据库表，按collection统计时赋值。
	Collection *string `json:"collection,omitempty"`

	// 用户名，按user统计时赋值。
	User *string `json:"user,omitempty"`

	// 客户端，按client统计时赋值。
	Client *string `json:"client,omitempty"`
}

func (o SlowSqlStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowSqlStatistics struct{}"
	}

	return strings.Join([]string{"SlowSqlStatistics", string(data)}, " ")
}
