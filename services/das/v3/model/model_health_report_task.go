package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportTask struct {

	// 报告ID
	TaskId string `json:"task_id"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 创建时间（Unix timestamp），单位：毫秒。
	CreateAt int64 `json:"create_at"`

	// 诊断状态
	ReportStatus string `json:"report_status"`

	// 风险点数量
	RiskCount int32 `json:"risk_count"`

	// 触发源
	Origin string `json:"origin"`

	// 日报诊断区间的起始时间（Unix timestamp），单位：毫秒。
	StartAt int64 `json:"start_at"`

	// 日报诊断区间的结束时间（Unix timestamp），单位：毫秒。
	EndAt int64 `json:"end_at"`
}

func (o HealthReportTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportTask struct{}"
	}

	return strings.Join([]string{"HealthReportTask", string(data)}, " ")
}
