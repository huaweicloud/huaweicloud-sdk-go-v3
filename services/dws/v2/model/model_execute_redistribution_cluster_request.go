package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteRedistributionClusterRequest Request Object
type ExecuteRedistributionClusterRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	Body *RedistributionReq `json:"body,omitempty"`
}

func (o ExecuteRedistributionClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteRedistributionClusterRequest struct{}"
	}

	return strings.Join([]string{"ExecuteRedistributionClusterRequest", string(data)}, " ")
}
