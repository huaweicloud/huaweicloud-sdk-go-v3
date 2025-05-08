package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableLogicalClusterPlanRequest Request Object
type EnableLogicalClusterPlanRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 计划ID
	PlanId string `json:"plan_id"`
}

func (o EnableLogicalClusterPlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableLogicalClusterPlanRequest struct{}"
	}

	return strings.Join([]string{"EnableLogicalClusterPlanRequest", string(data)}, " ")
}
