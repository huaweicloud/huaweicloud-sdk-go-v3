package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListWorkloadQueueResponse struct {

	// 工作负载队列名称。
	WorkloadQueueNameList *[]string `json:"workload_queue_name_list,omitempty"`
	HttpStatusCode        int       `json:"-"`
}

func (o ListWorkloadQueueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkloadQueueResponse struct{}"
	}

	return strings.Join([]string{"ListWorkloadQueueResponse", string(data)}, " ")
}
