package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchPlanStageRequest Request Object
type SwitchPlanStageRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 计划ID
	PlanId string `json:"plan_id"`

	Body *WorkloadPlanStageIdReq `json:"body,omitempty"`
}

func (o SwitchPlanStageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchPlanStageRequest struct{}"
	}

	return strings.Join([]string{"SwitchPlanStageRequest", string(data)}, " ")
}
