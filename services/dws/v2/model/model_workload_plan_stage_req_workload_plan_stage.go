package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadPlanStageReqWorkloadPlanStage 资源管理计划阶段详情
type WorkloadPlanStageReqWorkloadPlanStage struct {

	// 日期
	Day *string `json:"day,omitempty"`

	// 月份
	Month *string `json:"month,omitempty"`

	// 计划阶段
	StageName *string `json:"stage_name,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 资源队列
	QueueList *[]QueueResourceItem `json:"queue_list,omitempty"`
}

func (o WorkloadPlanStageReqWorkloadPlanStage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadPlanStageReqWorkloadPlanStage struct{}"
	}

	return strings.Join([]string{"WorkloadPlanStageReqWorkloadPlanStage", string(data)}, " ")
}
