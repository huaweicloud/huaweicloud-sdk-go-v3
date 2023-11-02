package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkloadPlanResponse Response Object
type ShowWorkloadPlanResponse struct {

	// 结果状态码。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// 结果描述。
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
