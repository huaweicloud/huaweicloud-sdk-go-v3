package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FullSqlStat 全量SQL统计分析
type FullSqlStat struct {

	// 是否收集全量SQL
	CollectFullSql *bool `json:"collect_full_sql,omitempty"`

	// 全量SQL Top总执行次数列表
	ExecuteTopTemplates *[]HealthReportSqlTemplate `json:"execute_top_templates,omitempty"`

	// 全量SQL Top总扫描行数列表
	SumRowsExaminedTopTemplates *[]HealthReportSqlTemplate `json:"sum_rows_examined_top_templates,omitempty"`

	// 全量SQL Top平均执行耗时列表
	AvgCostTopTemplates *[]HealthReportSqlTemplate `json:"avg_cost_top_templates,omitempty"`

	// 统计分析是否成功
	AnalyzeSuccess *bool `json:"analyze_success,omitempty"`

	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o FullSqlStat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullSqlStat struct{}"
	}

	return strings.Join([]string{"FullSqlStat", string(data)}, " ")
}
