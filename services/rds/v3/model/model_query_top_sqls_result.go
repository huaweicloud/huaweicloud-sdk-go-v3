package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryTopSqlsResult struct {

	// 对查询计算的二进制哈希值，用于标识具有类似逻辑的查询。
	Id *string `json:"id,omitempty"`

	// 当前执行的SQL语句。
	Statement *string `json:"statement,omitempty"`

	// SQL全文。
	Query *string `json:"query,omitempty"`

	// 数据库名称。
	DbName *string `json:"db_name,omitempty"`

	// 执行总次数。
	ExecutionCount *string `json:"execution_count,omitempty"`

	// 执行总次数百分比，范围0.0000-1.0000。
	ExecutionCountPercent *string `json:"execution_count_percent,omitempty"`

	// 总CPU耗时，单位ms。
	TotalCpuTime *string `json:"total_cpu_time,omitempty"`

	// 总CPU耗时百分比，范围0.0000-1.0000。
	TotalCpuTimePercent *string `json:"total_cpu_time_percent,omitempty"`

	// 平均CPU耗时，单位ms。
	AvgCpuTime *string `json:"avg_cpu_time,omitempty"`

	// 平均CPU耗时百分比，范围0.0000-1.0000。
	AvgCpuTimePercent *string `json:"avg_cpu_time_percent,omitempty"`

	// 总执行耗时。
	TotalDurationTime *string `json:"total_duration_time,omitempty"`

	// 总执行耗时百分比，范围0.0000-1.0000。
	TotalDurationTimePercent *string `json:"total_duration_time_percent,omitempty"`

	// 平均执行耗时。
	AvgDurationTime *string `json:"avg_duration_time,omitempty"`

	// 平均执行耗时百分比，范围0.0000-1.0000。
	AvgDurationTimePercent *string `json:"avg_duration_time_percent,omitempty"`

	// 总返回行。
	TotalRows *string `json:"total_rows,omitempty"`

	// 总返回行百分比，范围0.0000-1.0000。
	TotalRowsPercent *string `json:"total_rows_percent,omitempty"`

	// 平均返回行。
	AvgRows *string `json:"avg_rows,omitempty"`

	// 平均返回行百分比，范围0.0000-1.0000。
	AvgRowsPercent *string `json:"avg_rows_percent,omitempty"`

	// 总逻辑读消耗。
	TotalLogicalReads *string `json:"total_logical_reads,omitempty"`

	// 总逻辑读百分比，范围0.0000-1.0000。
	TotalLogicalReadsPercent *string `json:"total_logical_reads_percent,omitempty"`

	// 平均逻辑读消耗。
	AvgLogicalReads *string `json:"avg_logical_reads,omitempty"`

	// 平均逻辑读百分比，范围0.0000-1.0000。
	AvgLogicalReadsPercent *string `json:"avg_logical_reads_percent,omitempty"`

	// 平均逻辑写消耗。
	AvgLogicalWrite *string `json:"avg_logical_write,omitempty"`

	// 平均逻辑写百分比，范围0.0000-1.0000。
	AvgLogicalWritePercent *string `json:"avg_logical_write_percent,omitempty"`

	// 总逻辑写消耗。
	TotalLogicalWrite *string `json:"total_logical_write,omitempty"`

	// 总逻辑写百分比，范围0.0000-1.0000。
	TotalLogicalWritePercent *string `json:"total_logical_write_percent,omitempty"`

	// 平均物理读消耗。
	AvgPhysicalReads *string `json:"avg_physical_reads,omitempty"`

	// 平均物理读百分比，范围0.0000-1.0000。
	AvgPhysicalReadsPercent *string `json:"avg_physical_reads_percent,omitempty"`

	// 总物理读消耗。
	TotalPhysicalReads *string `json:"total_physical_reads,omitempty"`

	// 总物理读百分比，范围0.0000-1.0000。
	TotalPhysicalReadsPercent *string `json:"total_physical_reads_percent,omitempty"`

	// 上次执行时间。
	LastExecutionTime *string `json:"last_execution_time,omitempty"`
}

func (o QueryTopSqlsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTopSqlsResult struct{}"
	}

	return strings.Join([]string{"QueryTopSqlsResult", string(data)}, " ")
}
