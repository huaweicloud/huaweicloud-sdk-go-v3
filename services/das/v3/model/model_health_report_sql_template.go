package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportSqlTemplate struct {

	// 模版ID。
	TemplateId string `json:"template_id"`

	// 模版内容。
	Template string `json:"template"`

	// 数据库列表。
	Databases []string `json:"databases"`

	// 执行次数。
	Times int64 `json:"times"`

	// 平均执行时间。
	AvgQueryTime float64 `json:"avg_query_time"`

	// 最大执行时间。
	MaxQueryTime float64 `json:"max_query_time"`

	// 平均扫描行数。
	AvgRowsExamined float64 `json:"avg_rows_examined"`

	// 最大扫描行数。
	MaxRowsExamined float64 `json:"max_rows_examined"`

	// 总扫描行数。
	SumRowsExamined float64 `json:"sum_rows_examined"`

	// 平均返回行数。
	AvgRowsSent float64 `json:"avg_rows_sent"`

	// 最大返回行数。
	MaxRowsSent float64 `json:"max_rows_sent"`
}

func (o HealthReportSqlTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportSqlTemplate struct{}"
	}

	return strings.Join([]string{"HealthReportSqlTemplate", string(data)}, " ")
}
