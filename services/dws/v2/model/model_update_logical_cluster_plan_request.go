package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogicalClusterPlanRequest Request Object
type UpdateLogicalClusterPlanRequest struct {

	// 指定待编辑集群的ID
	ClusterId string `json:"cluster_id"`

	// 逻辑集群增删计划ID
	PlanId string `json:"plan_id"`

	Body *UpdateLogicalClusterPlanBo `json:"body,omitempty"`
}

func (o UpdateLogicalClusterPlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogicalClusterPlanRequest struct{}"
	}

	return strings.Join([]string{"UpdateLogicalClusterPlanRequest", string(data)}, " ")
}
