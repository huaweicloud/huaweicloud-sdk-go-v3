package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkloadQueueResponse Response Object
type ListWorkloadQueueResponse struct {

	// 资源池队列详情
	QueueList *[]PlanStageQueue `json:"queue_list,omitempty"`

	// 资源池名称队列
	WorkloadQueueNameList *[]string `json:"workload_queue_name_list,omitempty"`

	// 资源池队列查询返回码
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`
	HttpStatusCode  int    `json:"-"`
}

func (o ListWorkloadQueueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkloadQueueResponse struct{}"
	}

	return strings.Join([]string{"ListWorkloadQueueResponse", string(data)}, " ")
}
