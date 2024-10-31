package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportInspectionScore struct {

	// 得分。
	Score float64 `json:"score"`

	// 严重事件。
	Critical int32 `json:"critical"`

	// 警告事件。
	Medium int32 `json:"medium"`

	// 优化事件。
	Light int32 `json:"light"`

	// CPU使用率。
	CpuUsage float64 `json:"cpu_usage"`

	// 内存使用率。
	MemUsage float64 `json:"mem_usage"`

	// 空间使用率。
	SpaceUsage float64 `json:"space_usage"`

	// 连接使用率。
	ConnectionRate float64 `json:"connection_rate"`

	// IOPS使用率。
	IopsUsage float64 `json:"iops_usage"`

	// 活跃会话。
	ThreadRunning float64 `json:"thread_running"`

	// 慢SQL数量。
	SlowSqlTotal int64 `json:"slow_sql_total"`

	// 扣分详情。
	LostPointsDetailList []HealthReportLostPointsDetail `json:"lost_points_detail_list"`
}

func (o HealthReportInspectionScore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportInspectionScore struct{}"
	}

	return strings.Join([]string{"HealthReportInspectionScore", string(data)}, " ")
}
