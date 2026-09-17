package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InsTopSlowLogInfo InsTopSlowLogInfo响应
type InsTopSlowLogInfo struct {

	// 模板ID
	TemplateId *string `json:"template_id,omitempty"`

	// SQL模板
	Template *string `json:"template,omitempty"`

	// 数据库名
	Databases *[]string `json:"databases,omitempty"`

	// 执行次数
	Time *int64 `json:"time,omitempty"`

	// 平均执行时间
	AvgQueryTime *float64 `json:"avg_query_time,omitempty"`

	// 最大执行时间
	MaxQueryTime *float64 `json:"max_query_time,omitempty"`

	// 平均扫描行
	AvgRowsExamined *float64 `json:"avg_rows_examined,omitempty"`

	// 最大扫描行
	MaxRowsExamined *float64 `json:"max_rows_examined,omitempty"`

	// 总扫描行
	SumRowsExamined *float64 `json:"sum_rows_examined,omitempty"`

	// 平均返回行
	AvgRowsSent *float64 `json:"avg_rows_sent,omitempty"`

	// 最大返回行
	MaxRowsSent *float64 `json:"max_rows_sent,omitempty"`
}

func (o InsTopSlowLogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InsTopSlowLogInfo struct{}"
	}

	return strings.Join([]string{"InsTopSlowLogInfo", string(data)}, " ")
}
