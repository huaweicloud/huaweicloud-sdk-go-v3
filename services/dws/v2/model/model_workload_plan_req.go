package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadPlanReq **参数解释**： 资源管理计划请求。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type WorkloadPlanReq struct {

	// **参数解释**： 计划名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PlanName string `json:"plan_name"`

	// **参数解释**： 逻辑集群名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`
}

func (o WorkloadPlanReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadPlanReq struct{}"
	}

	return strings.Join([]string{"WorkloadPlanReq", string(data)}, " ")
}
