package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteClusterUpgradeActionRequest Request Object
type ExecuteClusterUpgradeActionRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	Body *ExecuteClusterUpgradeActionRequestBody `json:"body,omitempty"`
}

func (o ExecuteClusterUpgradeActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteClusterUpgradeActionRequest struct{}"
	}

	return strings.Join([]string{"ExecuteClusterUpgradeActionRequest", string(data)}, " ")
}
