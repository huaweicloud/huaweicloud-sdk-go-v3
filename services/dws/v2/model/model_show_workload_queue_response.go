package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkloadQueueResponse Response Object
type ShowWorkloadQueueResponse struct {

	// 结果状态码。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// 结果描述。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`

	WorkloadQueue  *WorkloadQueueItem `json:"workload_queue,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowWorkloadQueueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkloadQueueResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkloadQueueResponse", string(data)}, " ")
}
