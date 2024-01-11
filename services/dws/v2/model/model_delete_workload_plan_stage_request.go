package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkloadPlanStageRequest Request Object
type DeleteWorkloadPlanStageRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 计划ID
	PlanId string `json:"plan_id"`

	// 计划阶段ID
	StageId string `json:"stage_id"`
}

func (o DeleteWorkloadPlanStageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkloadPlanStageRequest struct{}"
	}

	return strings.Join([]string{"DeleteWorkloadPlanStageRequest", string(data)}, " ")
}
