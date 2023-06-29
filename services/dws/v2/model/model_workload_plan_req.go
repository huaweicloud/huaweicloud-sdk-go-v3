package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadPlanReq 资源管理计划
type WorkloadPlanReq struct {

	// 计划名称
	PlanName string `json:"plan_name"`

	// 逻辑集群名称
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`
}

func (o WorkloadPlanReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadPlanReq struct{}"
	}

	return strings.Join([]string{"WorkloadPlanReq", string(data)}, " ")
}
