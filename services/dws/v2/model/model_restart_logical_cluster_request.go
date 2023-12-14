package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartLogicalClusterRequest Request Object
type RestartLogicalClusterRequest struct {

	// 指定重启集群的ID
	ClusterId string `json:"cluster_id"`

	// 指定待重启逻辑集群的ID
	LogicalClusterId string `json:"logical_cluster_id"`
}

func (o RestartLogicalClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartLogicalClusterRequest struct{}"
	}

	return strings.Join([]string{"RestartLogicalClusterRequest", string(data)}, " ")
}
