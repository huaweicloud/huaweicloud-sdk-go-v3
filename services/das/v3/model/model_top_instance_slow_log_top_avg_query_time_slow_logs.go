package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopInstanceSlowLogTopAvgQueryTimeSlowLogs struct {

	// SQL模板ID
	TemplateId *string `json:"template_id,omitempty"`

	// SQL模板
	Template *string `json:"template,omitempty"`

	// 数据库名称
	Databases *[]string `json:"databases,omitempty"`

	// 执行次数
	Times *int32 `json:"times,omitempty"`

	// 平均执行时间
	AvgQueryTime *float64 `json:"avg_query_time,omitempty"`

	// 最大执行时间
	MaxQueryTime *float64 `json:"max_query_time,omitempty"`

	// 平均扫描行数
	AvgRowsExamined *float64 `json:"avg_rows_examined,omitempty"`

	// 扫描行数总和
	SumRowsExamined *float64 `json:"sum_rows_examined,omitempty"`

	// 平均返回行数
	AvgRowsSent *float64 `json:"avg_rows_sent,omitempty"`
}

func (o TopInstanceSlowLogTopAvgQueryTimeSlowLogs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopInstanceSlowLogTopAvgQueryTimeSlowLogs struct{}"
	}

	return strings.Join([]string{"TopInstanceSlowLogTopAvgQueryTimeSlowLogs", string(data)}, " ")
}
