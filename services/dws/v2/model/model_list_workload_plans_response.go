package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkloadPlansResponse Response Object
type ListWorkloadPlansResponse struct {

	// 结果状态码。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// 结果描述。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`

	// 资源池名称。
	PlanList *[]WorkloadPlanInfo `json:"plan_list,omitempty"`

	// 总数量
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListWorkloadPlansResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkloadPlansResponse struct{}"
	}

	return strings.Join([]string{"ListWorkloadPlansResponse", string(data)}, " ")
}
