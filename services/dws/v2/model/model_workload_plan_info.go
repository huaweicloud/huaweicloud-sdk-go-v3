package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadPlanInfo 工作计划
type WorkloadPlanInfo struct {

	// 计划状态。
	Status *int32 `json:"status,omitempty"`

	// **参数解释**： 项目ID。获取方式方法请参见[获取项目ID](dws_02_0011.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProjectId string `json:"project_id"`

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// 计划ID。
	PlanId string `json:"plan_id"`

	// 计划名称。
	PlanName string `json:"plan_name"`

	// 当前计划阶段。
	CurrentStage *string `json:"current_stage,omitempty"`

	// 逻辑集群名称。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// 计划阶段列表。
	StageList *[]PlanStage `json:"stage_list,omitempty"`
}

func (o WorkloadPlanInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadPlanInfo struct{}"
	}

	return strings.Join([]string{"WorkloadPlanInfo", string(data)}, " ")
}
