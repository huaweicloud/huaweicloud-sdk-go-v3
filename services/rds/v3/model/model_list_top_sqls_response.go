package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopSqlsResponse Response Object
type ListTopSqlsResponse struct {

	// 总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// TOP SQL 信息列表。
	List *[]QueryTopSqlsResult `json:"list,omitempty"`

	// 平均CPU开销表TOP SQL列表。
	AvgCpuTimeTopValues *[]TopSqlsTimeResult `json:"avg_cpu_time_top_values,omitempty"`

	// 平均执行耗时TOP SQL列表。
	AvgDurationTimeTopValues *[]TopSqlsTimeResult `json:"avg_duration_time_top_values,omitempty"`

	// 平均返回行TOP SQL列表。
	AvgRowsTopValues *[]TopSqlsRowResult `json:"avg_rows_top_values,omitempty"`

	// 平均逻辑读TOP SQL列表。
	AvgLogicalTopValues *[]TopSqlsLogicalReadResult `json:"avg_logical_top_values,omitempty"`

	// 总CPU开销表TOP SQL列表。
	TotalCpuTimeTopValues *[]TopSqlsTimeResult `json:"total_cpu_time_top_values,omitempty"`

	// 总执行耗时TOP SQL列表。
	TotalDurationTimeTopValues *[]TopSqlsTimeResult `json:"total_duration_time_top_values,omitempty"`

	// 总返回行TOP SQL列表。
	TotalRowsTopValues *[]TopSqlsRowResult `json:"total_rows_top_values,omitempty"`

	// 总逻辑读TOP SQL列表。
	TotalLogicalReadsTopValues *[]TopSqlsLogicalReadResult `json:"total_logical_reads_top_values,omitempty"`
	HttpStatusCode             int                         `json:"-"`
}

func (o ListTopSqlsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopSqlsResponse struct{}"
	}

	return strings.Join([]string{"ListTopSqlsResponse", string(data)}, " ")
}
