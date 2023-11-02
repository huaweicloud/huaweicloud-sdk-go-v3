package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkloadPlanResponse Response Object
type DeleteWorkloadPlanResponse struct {

	// 响应编码。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// 响应信息。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteWorkloadPlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkloadPlanResponse struct{}"
	}

	return strings.Join([]string{"DeleteWorkloadPlanResponse", string(data)}, " ")
}
