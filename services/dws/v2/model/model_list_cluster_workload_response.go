package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterWorkloadResponse Response Object
type ListClusterWorkloadResponse struct {

	// **参数解释**： 结果状态码。 **取值范围**： 不涉及。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// **参数解释**： 结果描述。 **取值范围**： 不涉及。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`

	WorkloadStatus *WorkloadStatus `json:"workload_status,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListClusterWorkloadResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterWorkloadResponse struct{}"
	}

	return strings.Join([]string{"ListClusterWorkloadResponse", string(data)}, " ")
}
