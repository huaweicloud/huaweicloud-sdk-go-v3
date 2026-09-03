package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceHealthReport4ApiResponse Response Object
type ShowInstanceHealthReport4ApiResponse struct {

	// 日报诊断是否成功
	Success *bool `json:"success,omitempty"`

	// 日报诊断区间的起始时间（Unix timestamp），单位：毫秒
	StartAt *int64 `json:"start_at,omitempty"`

	// 日报诊断区间的结束时间（Unix timestamp），单位：毫秒
	EndAt *int64 `json:"end_at,omitempty"`

	// 报告ID
	TaskId *string `json:"task_id,omitempty"`

	SummaryInfo *SummaryInfo `json:"summary_info,omitempty"`

	InstanceInfo *HealthReportInstanceInfo `json:"instance_info,omitempty"`

	PerformanceStat *PerformanceStat `json:"performance_stat,omitempty"`

	DiskStat *DiskStat `json:"disk_stat,omitempty"`

	TableSpaceStat *TableSpaceStat `json:"table_space_stat,omitempty"`

	SlowLogStat *SlowLogStat `json:"slow_log_stat,omitempty"`

	FullSqlStat *FullSqlStat `json:"full_sql_stat,omitempty"`

	InspectionStat *InspectionStat `json:"inspection_stat,omitempty"`

	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty"`

	// 报告链接
	SuffixUri      *string `json:"suffix_uri,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceHealthReport4ApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceHealthReport4ApiResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceHealthReport4ApiResponse", string(data)}, " ")
}
