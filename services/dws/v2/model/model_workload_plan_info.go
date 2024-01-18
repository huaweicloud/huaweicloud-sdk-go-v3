package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadPlanInfo 工作计划
type WorkloadPlanInfo struct {

	// 计划状态。
	Status *int32 `json:"status,omitempty"`

	// 项目ID。
	ProjectId string `json:"project_id"`

	// 集群ID。
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
