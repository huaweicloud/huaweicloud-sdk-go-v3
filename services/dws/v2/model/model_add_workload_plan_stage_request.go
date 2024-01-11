package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddWorkloadPlanStageRequest Request Object
type AddWorkloadPlanStageRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 计划ID
	PlanId string `json:"plan_id"`

	Body *WorkloadPlanStageReq `json:"body,omitempty"`
}

func (o AddWorkloadPlanStageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddWorkloadPlanStageRequest struct{}"
	}

	return strings.Join([]string{"AddWorkloadPlanStageRequest", string(data)}, " ")
}
