package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkloadPlanStageRequest Request Object
type ShowWorkloadPlanStageRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 计划ID
	PlanId string `json:"plan_id"`

	// 计划阶段ID
	StageId string `json:"stage_id"`
}

func (o ShowWorkloadPlanStageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkloadPlanStageRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkloadPlanStageRequest", string(data)}, " ")
}
