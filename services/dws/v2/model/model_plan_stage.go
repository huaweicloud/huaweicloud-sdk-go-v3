package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlanStage 工作计划阶段。
type PlanStage struct {

	// 计划月份。
	Month string `json:"month"`

	// 计划日期。
	Day string `json:"day"`

	// 计划ID。
	PlanId string `json:"plan_id"`

	// 计划阶段ID。
	StageId string `json:"stage_id"`

	// 计划阶段名称。
	StageName string `json:"stage_name"`

	// 计划开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 计划结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 下次校验时间
	NextValidTime *string `json:"next_valid_time,omitempty"`

	// 资源队列列表
	QueueList *[]QueueResourceItem `json:"queue_list,omitempty"`
}

func (o PlanStage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlanStage struct{}"
	}

	return strings.Join([]string{"PlanStage", string(data)}, " ")
}
