package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportSlowLogStat struct {

	// 是否收集慢SQL。
	CollectSlowLog bool `json:"collect_slow_log"`

	// 慢SQL Top执行次数列表。
	TopExecuteSlowLogs []HealthReportSqlTemplate `json:"top_execute_slow_logs"`

	// 慢SQL Top平均执行时间列表。
	TopAvgQueryTimeSlowLogs []HealthReportSqlTemplate `json:"top_avg_query_time_slow_logs"`

	// 慢SQL Top最大执行时间列表。
	TopMaxQueryTimeSlowLogs []HealthReportSqlTemplate `json:"top_max_query_time_slow_logs"`

	// 慢SQL Top扫描返回比列表。
	RowsExaminedExceeding []HealthReportSqlTemplate `json:"rows_examined_exceeding"`

	// 统计分析是否成功。
	AnalyzeSuccess bool `json:"analyze_success"`

	// 错误信息。
	ErrorMessage string `json:"error_message"`
}

func (o HealthReportSlowLogStat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportSlowLogStat struct{}"
	}

	return strings.Join([]string{"HealthReportSlowLogStat", string(data)}, " ")
}
