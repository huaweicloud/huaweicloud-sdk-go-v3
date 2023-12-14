package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogicalClusterRequest Request Object
type UpdateLogicalClusterRequest struct {

	// 指定待编辑集群的ID
	ClusterId string `json:"cluster_id"`

	// 指定待编辑逻辑集群的ID
	LogicalClusterId string `json:"logical_cluster_id"`

	Body *UpdateLogicalClusterRequestBody `json:"body,omitempty"`
}

func (o UpdateLogicalClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogicalClusterRequest struct{}"
	}

	return strings.Join([]string{"UpdateLogicalClusterRequest", string(data)}, " ")
}
