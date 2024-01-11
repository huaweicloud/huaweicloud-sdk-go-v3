package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadPlanStageReq 资源管理计划阶段
type WorkloadPlanStageReq struct {
	WorkloadPlanStage *WorkloadPlanStageReqWorkloadPlanStage `json:"workload_plan_stage,omitempty"`
}

func (o WorkloadPlanStageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadPlanStageReq struct{}"
	}

	return strings.Join([]string{"WorkloadPlanStageReq", string(data)}, " ")
}
