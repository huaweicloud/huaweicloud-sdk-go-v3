package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkloadPlanRequest Request Object
type CreateWorkloadPlanRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	Body *WorkloadPlanReq `json:"body,omitempty"`
}

func (o CreateWorkloadPlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkloadPlanRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkloadPlanRequest", string(data)}, " ")
}
