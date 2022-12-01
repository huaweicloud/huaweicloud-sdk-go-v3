package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExecuteRedistributionClusterRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	Body *interface{} `json:"body,omitempty"`
}

func (o ExecuteRedistributionClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteRedistributionClusterRequest struct{}"
	}

	return strings.Join([]string{"ExecuteRedistributionClusterRequest", string(data)}, " ")
}
