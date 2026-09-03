package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HealthReportTaskInfo 健康报告任务信息
type HealthReportTaskInfo struct {

	// 报告ID
	TaskId *string `json:"task_id,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 创建时间(Unix timestamp),单位:毫秒
	CreateAt *int64 `json:"create_at,omitempty"`

	// 诊断状态
	ReportStatus *string `json:"report_status,omitempty"`

	// 风险点数量
	RiskCount *int32 `json:"risk_count,omitempty"`

	// 触发源
	Origin *string `json:"origin,omitempty"`

	// 日报诊断区间的起始时间(Unix timestamp),单位:毫秒
	StartAt *int64 `json:"start_at,omitempty"`

	// 日报诊断区间的结束时间(Unix timestamp),单位:毫秒
	EndAt *int64 `json:"end_at,omitempty"`
}

func (o HealthReportTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportTaskInfo struct{}"
	}

	return strings.Join([]string{"HealthReportTaskInfo", string(data)}, " ")
}
