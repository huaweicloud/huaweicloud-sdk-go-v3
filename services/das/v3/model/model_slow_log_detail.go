package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SlowLogDetail 慢日志明细
type SlowLogDetail struct {

	// 执行时间(sqlserver、mongodb：结束时间；其他引擎：开始时间)（Unix timestamp），单位：毫秒
	OccurrenceTime *int64 `json:"occurrence_time,omitempty"`

	// SQL模板ID
	SqlTemplateId *string `json:"sql_template_id,omitempty"`

	// 原始SQL语句
	OriginalSql *string `json:"original_sql,omitempty"`

	// 数据库名
	DbName *string `json:"db_name,omitempty"`

	// 客户端
	Client *string `json:"client,omitempty"`

	// 用户
	User *string `json:"user,omitempty"`

	// 执行耗时（秒）
	ExecuteTime *float64 `json:"execute_time,omitempty"`

	// 锁等待耗时（秒）
	LockWaitTime *float64 `json:"lock_wait_time,omitempty"`

	// 扫描行数
	RowsExamined *int64 `json:"rows_examined,omitempty"`

	// 返回行数
	RowsSent *int64 `json:"rows_sent,omitempty"`

	// 是否可诊断优化
	Tunable *bool `json:"tunable,omitempty"`

	// sqlserver：执行完成时间（Unix timestamp），单位：毫秒
	EndTime *int64 `json:"end_time,omitempty"`

	// sqlserver：应用名
	AppName *string `json:"app_name,omitempty"`

	// sqlserver：影响行数
	RowsAffected *int64 `json:"rows_affected,omitempty"`

	// sqlserver：CPU耗时（ms）
	CpuTime *float64 `json:"cpu_time,omitempty"`

	// sqlserver：IO逻辑读
	LogicalReads *int64 `json:"logical_reads,omitempty"`

	// sqlserver：IO物理读
	PhysicalReads *int64 `json:"physical_reads,omitempty"`

	// sqlserver：IO写
	Writes *int64 `json:"writes,omitempty"`

	// SQL操作类型
	SqlType *string `json:"sql_type,omitempty"`

	// mongodb：数据库表
	Collection *string `json:"collection,omitempty"`

	// mongodb：扫描索引数
	KeyExamined *int64 `json:"key_examined,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 节点名称
	NodeName *string `json:"node_name,omitempty"`

	// 执行状态
	Killed *string `json:"killed,omitempty"`
}

func (o SlowLogDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowLogDetail struct{}"
	}

	return strings.Join([]string{"SlowLogDetail", string(data)}, " ")
}
