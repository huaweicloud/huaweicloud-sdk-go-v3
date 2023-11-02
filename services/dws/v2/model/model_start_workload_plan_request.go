package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartWorkloadPlanRequest Request Object
type StartWorkloadPlanRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 计划ID
	PlanId string `json:"plan_id"`
}

func (o StartWorkloadPlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartWorkloadPlanRequest struct{}"
	}

	return strings.Join([]string{"StartWorkloadPlanRequest", string(data)}, " ")
}
