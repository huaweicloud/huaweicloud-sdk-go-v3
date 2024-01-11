package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkloadQueueUsersResponse Response Object
type ListWorkloadQueueUsersResponse struct {

	// 结果状态码。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// 结果描述。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`

	// 队列名称。
	QueueName *string `json:"queue_name,omitempty"`

	// 队列用户列表。
	UserList *[]WorkloadQueueUser `json:"user_list,omitempty"`

	// 总数量
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListWorkloadQueueUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkloadQueueUsersResponse struct{}"
	}

	return strings.Join([]string{"ListWorkloadQueueUsersResponse", string(data)}, " ")
}
