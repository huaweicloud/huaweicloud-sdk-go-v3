package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkloadPlanResponse Response Object
type CreateWorkloadPlanResponse struct {

	// **参数解释**： 响应编码。 **取值范围**： 不涉及。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// **参数解释**： 响应信息。 **取值范围**： 不涉及。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateWorkloadPlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkloadPlanResponse struct{}"
	}

	return strings.Join([]string{"CreateWorkloadPlanResponse", string(data)}, " ")
}
