package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkloadPlanStageResponse Response Object
type ShowWorkloadPlanStageResponse struct {

	// 结果状态码。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// 结果描述。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`

	WorkloadPlanStage *PlanStage `json:"workload_plan_stage,omitempty"`
	HttpStatusCode    int        `json:"-"`
}

func (o ShowWorkloadPlanStageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkloadPlanStageResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkloadPlanStageResponse", string(data)}, " ")
}
