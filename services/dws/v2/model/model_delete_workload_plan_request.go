package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkloadPlanRequest Request Object
type DeleteWorkloadPlanRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 计划ID
	PlanId string `json:"plan_id"`
}

func (o DeleteWorkloadPlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkloadPlanRequest struct{}"
	}

	return strings.Join([]string{"DeleteWorkloadPlanRequest", string(data)}, " ")
}
