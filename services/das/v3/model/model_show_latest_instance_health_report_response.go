package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLatestInstanceHealthReportResponse Response Object
type ShowLatestInstanceHealthReportResponse struct {

	// 日报诊断是否成功
	Success *bool `json:"success,omitempty"`

	// 日报诊断区间的起始时间（Unix timestamp），单位：毫秒
	StartAt float32 `json:"start_at,omitempty"`

	// 日报诊断区间的结束时间（Unix timestamp），单位：毫秒
	EndAt float32 `json:"end_at,omitempty"`

	// 报告ID
	TaskId *string `json:"task_id,omitempty"`

	SummaryInfo *HealthReportSummaryInfo `json:"summary_info,omitempty"`

	InstanceInfo *HealthReportInstanceInfo `json:"instance_info,omitempty"`

	PerformanceStat *HealthReportPerformanceStat `json:"performance_stat,omitempty"`

	DiskStat *HealthReportDiskStat `json:"disk_stat,omitempty"`

	TableSpaceStat *HealthReportTableSpaceNewStat `json:"table_space_stat,omitempty"`

	SlowLogStat *HealthReportSlowLogStat `json:"slow_log_stat,omitempty"`

	FullSqlStat *HealthReportFullSqlStat `json:"full_sql_stat,omitempty"`

	InspectionStat *HealthReportInspectionStat `json:"inspection_stat,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowLatestInstanceHealthReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLatestInstanceHealthReportResponse struct{}"
	}

	return strings.Join([]string{"ShowLatestInstanceHealthReportResponse", string(data)}, " ")
}
