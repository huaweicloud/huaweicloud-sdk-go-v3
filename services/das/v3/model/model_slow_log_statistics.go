package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SlowLogStatistics 慢日志统计信息
type SlowLogStatistics struct {

	// 执行次数
	ExecuteCount *int64 `json:"execute_count,omitempty"`

	// 平均执行耗时(s)
	AvgExecuteTime *float64 `json:"avg_execute_time,omitempty"`

	// 最大执行耗时(s)
	MaxExecuteTime *float64 `json:"max_execute_time,omitempty"`

	// 平均锁等待时间(s)
	AvgLockWaitTime *float64 `json:"avg_lock_wait_time,omitempty"`

	// 最大锁等待时间(s)
	MaxLockWaitTime *float64 `json:"max_lock_wait_time,omitempty"`

	// 平均扫描行数
	AvgRowsExamined *float64 `json:"avg_rows_examined,omitempty"`

	// 最大扫描行数
	MaxRowsExamined *float64 `json:"max_rows_examined,omitempty"`

	// 平均返回行数
	AvgRowsSent *float64 `json:"avg_rows_sent,omitempty"`

	// 最大返回行数
	MaxRowsSent *float64 `json:"max_rows_sent,omitempty"`

	// 平均扫描索引数
	AvgKeyExamined *float64 `json:"avg_key_examined,omitempty"`

	// 最大扫描索引数
	MaxKeyExamined *float64 `json:"max_key_examined,omitempty"`

	// 节点ID，按nodeId统计时赋值
	NodeId *string `json:"node_id,omitempty"`

	// 节点名称，按nodeId统计时赋值
	NodeName *string `json:"node_name,omitempty"`

	// 语句类型，按sqlType统计时赋值
	SqlType *string `json:"sql_type,omitempty"`

	// 库名，按dbName、collection统计时赋值
	DbName *string `json:"db_name,omitempty"`

	// 数据库表，按collection统计时赋值
	Collection *string `json:"collection,omitempty"`

	// 用户名，按user统计时赋值
	User *string `json:"user,omitempty"`

	// 客户端，按client统计时赋值
	Client *string `json:"client,omitempty"`
}

func (o SlowLogStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowLogStatistics struct{}"
	}

	return strings.Join([]string{"SlowLogStatistics", string(data)}, " ")
}
