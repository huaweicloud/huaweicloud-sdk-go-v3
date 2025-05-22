package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkloadQueueUsersResponse Response Object
type ListWorkloadQueueUsersResponse struct {

	// **参数解释**： 结果状态码。 **取值范围**： 不涉及。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// **参数解释**： 结果描述。 **取值范围**： 不涉及。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`

	// **参数解释**： 队列名称。 **取值范围**： 不涉及。
	QueueName *string `json:"queue_name,omitempty"`

	// **参数解释**： 队列用户列表。 **取值范围**： 不涉及。
	UserList *[]WorkloadQueueUser `json:"user_list,omitempty"`

	// **参数解释**： 总数量。 **取值范围**： 不涉及。
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
