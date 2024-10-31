package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportFullSqlStat struct {

	// 是否收集全量SQL。
	CollectFullSql bool `json:"collect_full_sql"`

	// 全量SQL Top总执行次数列表。
	ExecuteTopTemplates []HealthReportSqlTemplate `json:"execute_top_templates"`

	// 全量SQL Top总扫描行数列表。
	SumRowsExaminedTopTemplates []HealthReportSqlTemplate `json:"sum_rows_examined_top_templates"`

	// 全量SQL Top平均执行耗时列表。
	AvgCostTopTemplates []HealthReportSqlTemplate `json:"avg_cost_top_templates"`

	// 统计分析是否成功。
	AnalyzeSuccess bool `json:"analyze_success"`

	// 错误信息。
	ErrorMessage string `json:"error_message"`
}

func (o HealthReportFullSqlStat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportFullSqlStat struct{}"
	}

	return strings.Join([]string{"HealthReportFullSqlStat", string(data)}, " ")
}
