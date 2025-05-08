package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableLogicalClusterPlanRequest Request Object
type DisableLogicalClusterPlanRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 计划ID
	PlanId string `json:"plan_id"`
}

func (o DisableLogicalClusterPlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableLogicalClusterPlanRequest struct{}"
	}

	return strings.Join([]string{"DisableLogicalClusterPlanRequest", string(data)}, " ")
}
