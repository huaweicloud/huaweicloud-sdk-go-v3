package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopObjectOverview top对象总览
type TopObjectOverview struct {

	// 平均cpu耗时(单位为毫秒)
	AvgCpuTime *float64 `json:"avg_cpu_time,omitempty"`

	// 平均cpu耗时百分比
	AvgCpuTimePercentage *float64 `json:"avg_cpu_time_percentage,omitempty"`

	// 平均执行耗时(单位为毫秒)
	AvgExecutionTime *float64 `json:"avg_execution_time,omitempty"`

	// 平均执行耗时百分比
	AvgExecutionTimePercentage *float64 `json:"avg_execution_time_percentage,omitempty"`

	// 平均逻辑io
	AvgLogicalIo *float64 `json:"avg_logical_io,omitempty"`

	// 平均逻辑io百分比
	AvgLogicalIoPercentage *float64 `json:"avg_logical_io_percentage,omitempty"`

	// 平均逻辑读
	AvgLogicalReads *float64 `json:"avg_logical_reads,omitempty"`

	// 平均逻辑读百分比
	AvgLogicalReadsPercentage *float64 `json:"avg_logical_reads_percentage,omitempty"`

	// 平均逻辑写
	AvgLogicalWrites *float64 `json:"avg_logical_writes,omitempty"`

	// 平均逻辑写百分比
	AvgLogicalWritesPercentage *float64 `json:"avg_logical_writes_percentage,omitempty"`

	// 平均物理读
	AvgPhysicalReads *float64 `json:"avg_physical_reads,omitempty"`

	// 平均物理读百分比
	AvgPhysicalReadsPercentage *float64 `json:"avg_physical_reads_percentage,omitempty"`

	// 平均返回行
	AvgRows *float64 `json:"avg_rows,omitempty"`

	// 平均返回行百分比
	AvgRowsPercentage *float64 `json:"avg_rows_percentage,omitempty"`

	// 数据库名
	DatabaseName *string `json:"database_name,omitempty"`

	// 对象id
	ObjectId *string `json:"object_id,omitempty"`

	// 对象名称
	ObjectName *string `json:"object_name,omitempty"`

	// id
	RowId *string `json:"row_id,omitempty"`

	// 对象类型
	ObjectType *string `json:"object_type,omitempty"`

	// 模式
	SchemaName *string `json:"schema_name,omitempty"`

	// 总cpu耗时(单位为毫秒)
	TotalCpuTime *float64 `json:"total_cpu_time,omitempty"`

	// 总cpu耗时百分比
	TotalCpuTimePercentage *float64 `json:"total_cpu_time_percentage,omitempty"`

	// 总执行耗时(单位为毫秒)
	TotalExecutionTime *float64 `json:"total_execution_time,omitempty"`

	// 总执行耗时百分比
	TotalExecutionTimePercentage *float64 `json:"total_execution_time_percentage,omitempty"`

	// 总执行次数
	TotalExecutionCount *float64 `json:"total_execution_count,omitempty"`

	// 总逻辑io
	TotalLogicalIo *float64 `json:"total_logical_io,omitempty"`

	// 总逻辑io百分比
	TotalLogicalIoPercentage *float64 `json:"total_logical_io_percentage,omitempty"`

	// 总逻辑读
	TotalLogicalReads *float64 `json:"total_logical_reads,omitempty"`

	// 总逻辑读百分比
	TotalLogicalReadsPercentage *float64 `json:"total_logical_reads_percentage,omitempty"`

	// 总逻辑写
	TotalLogicalWrites *float64 `json:"total_logical_writes,omitempty"`

	// 总逻辑写百分比
	TotalLogicalWritesPercentage *float64 `json:"total_logical_writes_percentage,omitempty"`

	// 总物理读
	TotalPhysicalReads *float64 `json:"total_physical_reads,omitempty"`

	// 总物理读百分比
	TotalPhysicalReadsPercentage *float64 `json:"total_physical_reads_percentage,omitempty"`

	// 总返回行
	TotalRows *float64 `json:"total_rows,omitempty"`

	// 总返回行百分比
	TotalRowsPercentage *float64 `json:"total_rows_percentage,omitempty"`
}

func (o TopObjectOverview) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopObjectOverview struct{}"
	}

	return strings.Join([]string{"TopObjectOverview", string(data)}, " ")
}
