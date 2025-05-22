package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkloadPlanStageResponse Response Object
type UpdateWorkloadPlanStageResponse struct {

	// **参数解释**： 结果状态码。 **取值范围**： 不涉及。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// **参数解释**： 结果描述。 **取值范围**： 不涉及。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`

	WorkloadPlanStage *PlanStage `json:"workload_plan_stage,omitempty"`
	HttpStatusCode    int        `json:"-"`
}

func (o UpdateWorkloadPlanStageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkloadPlanStageResponse struct{}"
	}

	return strings.Join([]string{"UpdateWorkloadPlanStageResponse", string(data)}, " ")
}
