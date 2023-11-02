package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopWorkloadPlanRequest Request Object
type StopWorkloadPlanRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 计划ID
	PlanId string `json:"plan_id"`
}

func (o StopWorkloadPlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopWorkloadPlanRequest struct{}"
	}

	return strings.Join([]string{"StopWorkloadPlanRequest", string(data)}, " ")
}
