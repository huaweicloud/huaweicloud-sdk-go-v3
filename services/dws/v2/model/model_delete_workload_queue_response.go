package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkloadQueueResponse Response Object
type DeleteWorkloadQueueResponse struct {

	// 响应编码。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// 响应信息。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteWorkloadQueueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkloadQueueResponse struct{}"
	}

	return strings.Join([]string{"DeleteWorkloadQueueResponse", string(data)}, " ")
}
