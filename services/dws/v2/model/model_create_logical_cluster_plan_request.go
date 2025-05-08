package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogicalClusterPlanRequest Request Object
type CreateLogicalClusterPlanRequest struct {

	// 指定集群的ID
	ClusterId string `json:"cluster_id"`

	Body *LogicalClusterPlanBo `json:"body,omitempty"`
}

func (o CreateLogicalClusterPlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogicalClusterPlanRequest struct{}"
	}

	return strings.Join([]string{"CreateLogicalClusterPlanRequest", string(data)}, " ")
}
