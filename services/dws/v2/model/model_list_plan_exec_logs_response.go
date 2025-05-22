package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlanExecLogsResponse Response Object
type ListPlanExecLogsResponse struct {

	// **参数解释**： 结果状态码。 **取值范围**： 不涉及。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// **参数解释**： 结果描述。 **取值范围**： 不涉及。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`

	// **参数解释**： 资源池名称。 **取值范围**： 不涉及。
	PlanLogs *[]PlanLog `json:"plan_logs,omitempty"`

	// **参数解释**： 总数量。 **取值范围**： 不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPlanExecLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlanExecLogsResponse struct{}"
	}

	return strings.Join([]string{"ListPlanExecLogsResponse", string(data)}, " ")
}
