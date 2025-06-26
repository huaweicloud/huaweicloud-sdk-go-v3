package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadPlanInfo **参数解释**： 资源管理计划。 **取值范围**： 不涉及。
type WorkloadPlanInfo struct {

	// **参数解释**： 计划状态。 **取值范围**： 不涉及。
	Status *int32 `json:"status,omitempty"`

	// **参数解释**： 项目ID。 **取值范围**： 不涉及。
	ProjectId string `json:"project_id"`

	// **参数解释**： 集群ID。 **取值范围**： 36位UUID。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 计划ID。 **取值范围**： 不涉及。
	PlanId string `json:"plan_id"`

	// **参数解释**： 计划名称。 **取值范围**： 不涉及。
	PlanName string `json:"plan_name"`

	// **参数解释**： 当前计划阶段。 **取值范围**： 不涉及。
	CurrentStage *string `json:"current_stage,omitempty"`

	// **参数解释**： 逻辑集群名称。 **取值范围**： 不涉及。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// **参数解释**： 计划阶段列表。 **取值范围**： 不涉及。
	StageList *[]PlanStage `json:"stage_list,omitempty"`
}

func (o WorkloadPlanInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadPlanInfo struct{}"
	}

	return strings.Join([]string{"WorkloadPlanInfo", string(data)}, " ")
}
