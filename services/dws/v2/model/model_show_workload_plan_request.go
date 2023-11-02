package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkloadPlanRequest Request Object
type ShowWorkloadPlanRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 计划ID
	PlanId string `json:"plan_id"`
}

func (o ShowWorkloadPlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkloadPlanRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkloadPlanRequest", string(data)}, " ")
}
