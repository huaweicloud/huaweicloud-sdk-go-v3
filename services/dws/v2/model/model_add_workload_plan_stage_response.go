package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddWorkloadPlanStageResponse Response Object
type AddWorkloadPlanStageResponse struct {

	// **参数解释**： 响应编码。 **取值范围**： 不涉及。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// **参数解释**： 响应信息。 **取值范围**： 不涉及。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddWorkloadPlanStageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddWorkloadPlanStageResponse struct{}"
	}

	return strings.Join([]string{"AddWorkloadPlanStageResponse", string(data)}, " ")
}
