package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkloadPlanResponse Response Object
type ShowWorkloadPlanResponse struct {

	// **参数解释**： 结果状态码。 **取值范围**： 不涉及。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// **参数解释**： 结果描述。 **取值范围**： 不涉及。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`

	WorkloadPlan   *WorkloadPlanInfo `json:"workload_plan,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowWorkloadPlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkloadPlanResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkloadPlanResponse", string(data)}, " ")
}
