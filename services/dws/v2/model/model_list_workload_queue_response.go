package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkloadQueueResponse Response Object
type ListWorkloadQueueResponse struct {

	// **参数解释**： 资源池名称。 **取值范围**： 不涉及。
	WorkloadQueueNameList *[]string `json:"workload_queue_name_list,omitempty"`

	// **参数解释**： 结果状态码。 **取值范围**： 不涉及。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// **参数解释**： 结果描述。 **取值范围**： 不涉及。
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
