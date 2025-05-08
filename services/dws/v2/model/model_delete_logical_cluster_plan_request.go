package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLogicalClusterPlanRequest Request Object
type DeleteLogicalClusterPlanRequest struct {

	// 指定待删除集群的ID
	ClusterId string `json:"cluster_id"`

	// 计划ID
	PlanId string `json:"plan_id"`
}

func (o DeleteLogicalClusterPlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogicalClusterPlanRequest struct{}"
	}

	return strings.Join([]string{"DeleteLogicalClusterPlanRequest", string(data)}, " ")
}
