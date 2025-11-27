package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchInspectionReport struct {

	// 报告ID
	TaskId string `json:"task_id"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 实例名称
	InstanceName string `json:"instance_name"`

	// CPU大小
	Cpu *int32 `json:"cpu,omitempty"`

	// 内存大小（单位:GB）
	Mem *int32 `json:"mem,omitempty"`

	// 磁盘大小（单位:GB）
	DiskSize *int32 `json:"disk_size,omitempty"`

	// 诊断报告生成时间（Unix timestamp），单位：毫秒。
	CreateAt *int64 `json:"create_at,omitempty"`

	// 诊断起始时间（Unix timestamp），单位：毫秒。
	StartAt int64 `json:"start_at"`

	// 诊断终止时间（Unix timestamp），单位：毫秒。
	EndAt int64 `json:"end_at"`

	// 健康等级
	HealthRank *string `json:"health_rank,omitempty"`

	// 评分
	Score *float64 `json:"score,omitempty"`

	// 扣分详情
	LostPointsDetailList *[]BatchInspectionLostPointsDetail `json:"lost_points_detail_list,omitempty"`
}

func (o BatchInspectionReport) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchInspectionReport struct{}"
	}

	return strings.Join([]string{"BatchInspectionReport", string(data)}, " ")
}
