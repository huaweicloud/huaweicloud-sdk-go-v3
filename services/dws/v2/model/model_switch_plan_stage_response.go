package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchPlanStageResponse Response Object
type SwitchPlanStageResponse struct {

	// 响应编码。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// 响应信息。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchPlanStageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchPlanStageResponse struct{}"
	}

	return strings.Join([]string{"SwitchPlanStageResponse", string(data)}, " ")
}
