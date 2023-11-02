package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartWorkloadPlanResponse Response Object
type StartWorkloadPlanResponse struct {

	// 响应编码。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// 响应信息。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartWorkloadPlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartWorkloadPlanResponse struct{}"
	}

	return strings.Join([]string{"StartWorkloadPlanResponse", string(data)}, " ")
}
