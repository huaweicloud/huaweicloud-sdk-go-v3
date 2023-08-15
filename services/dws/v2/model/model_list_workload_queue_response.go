package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkloadQueueResponse Response Object
type ListWorkloadQueueResponse struct {

	// 工作负载队列名称。
	WorkloadQueueNameList *[]string `json:"workload_queue_name_list,omitempty"`

	// 结果状态码。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// 结果描述。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListWorkloadQueueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkloadQueueResponse struct{}"
	}

	return strings.Join([]string{"ListWorkloadQueueResponse", string(data)}, " ")
}
