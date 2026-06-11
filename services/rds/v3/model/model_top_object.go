package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopObject top对象详情
type TopObject struct {

	// id
	RowId string `json:"row_id"`

	// 数据库名
	DatabaseName string `json:"database_name"`

	// 模式
	SchemaName string `json:"schema_name"`

	// 对象名
	ObjectName *string `json:"object_name,omitempty"`

	// 对象id
	ObjectId *string `json:"object_id,omitempty"`

	// 对象类型
	ObjectType *string `json:"object_type,omitempty"`

	// sql文本
	SqlStatement *string `json:"sql_statement,omitempty"`

	// 执行次数
	ExecutionCount *string `json:"execution_count,omitempty"`

	// 获取执行计划次数
	PlanGenerationNum *string `json:"plan_generation_num,omitempty"`

	// 最近执行时间
	LastExecutionTime *string `json:"last_execution_time,omitempty"`

	// 平均cpu耗时(单位为毫秒)
	AvgWorkerTime *string `json:"avg_worker_time,omitempty"`

	// 总cpu耗时(单位为毫秒)
	TotalWorkerTime *string `json:"total_worker_time,omitempty"`

	// 最近cpu耗时(单位为毫秒)
	LastWorkerTime *string `json:"last_worker_time,omitempty"`

	// 最小cpu耗时(单位为毫秒)
	MinWorkerTime *string `json:"min_worker_time,omitempty"`

	// 最大cpu耗时(单位为毫秒)
	MaxWorkerTime *string `json:"max_worker_time,omitempty"`

	// 平均逻辑读
	AvgLogicalReads *string `json:"avg_logical_reads,omitempty"`

	// 总共逻辑读
	TotalLogicalReads *string `json:"total_logical_reads,omitempty"`

	// 最近逻辑读
	LastLogicalReads *string `json:"last_logical_reads,omitempty"`

	// 最小逻辑读
	MinLogicalReads *string `json:"min_logical_reads,omitempty"`

	// 最大逻辑读
	MaxLogicalReads *string `json:"max_logical_reads,omitempty"`

	// 平均逻辑写
	AvgLogicalWrites *string `json:"avg_logical_writes,omitempty"`

	// 总共逻辑写
	TotalLogicalWrites *string `json:"total_logical_writes,omitempty"`

	// 最近逻辑写
	LastLogicalWrites *string `json:"last_logical_writes,omitempty"`

	// 最小逻辑写
	MinLogicalWrites *string `json:"min_logical_writes,omitempty"`

	// 最大逻辑写
	MaxLogicalWrites *string `json:"max_logical_writes,omitempty"`

	// 平均逻辑io
	AvgLogicalIo *string `json:"avg_logical_io,omitempty"`

	// 总共逻辑io
	TotalLogicalIo *string `json:"total_logical_io,omitempty"`

	// 最近逻辑io
	LastLogicalIo *string `json:"last_logical_io,omitempty"`

	// 最小逻辑io
	MinLogicalIo *string `json:"min_logical_io,omitempty"`

	// 最大逻辑io
	MaxLogicalIo *string `json:"max_logical_io,omitempty"`

	// 平均物理读
	AvgPhysicalReads *string `json:"avg_physical_reads,omitempty"`

	// 总共物理读
	TotalPhysicalReads *string `json:"total_physical_reads,omitempty"`

	// 最近物理读
	LastPhysicalReads *string `json:"last_physical_reads,omitempty"`

	// 最小物理读
	MinPhysicalReads *string `json:"min_physical_reads,omitempty"`

	// 最大物理读
	MaxPhysicalReads *string `json:"max_physical_reads,omitempty"`

	// 平均执行耗时
	AvgElapsedTime *string `json:"avg_elapsed_time,omitempty"`

	// 总共执行耗时
	TotalElapsedTime *string `json:"total_elapsed_time,omitempty"`

	// 最近执行耗时
	LastElapsedTime *string `json:"last_elapsed_time,omitempty"`

	// 最小执行耗时
	MinElapsedTime *string `json:"min_elapsed_time,omitempty"`

	// 最大执行耗时
	MaxElapsedTime *string `json:"max_elapsed_time,omitempty"`

	// 平均返回行
	AvgRows *string `json:"avg_rows,omitempty"`

	// 总返回行
	TotalRows *string `json:"total_rows,omitempty"`

	// 最近返回行
	LastRows *string `json:"last_rows,omitempty"`

	// 最小返回行
	MinRows *string `json:"min_rows,omitempty"`

	// 最大返回行
	MaxRows *string `json:"max_rows,omitempty"`
}

func (o TopObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopObject struct{}"
	}

	return strings.Join([]string{"TopObject", string(data)}, " ")
}
