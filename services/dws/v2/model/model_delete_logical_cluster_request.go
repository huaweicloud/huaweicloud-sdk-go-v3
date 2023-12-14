package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLogicalClusterRequest Request Object
type DeleteLogicalClusterRequest struct {

	// 指定待删除集群的ID
	ClusterId string `json:"cluster_id"`

	// 指定待删除逻辑集群的ID
	LogicalClusterId string `json:"logical_cluster_id"`
}

func (o DeleteLogicalClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogicalClusterRequest struct{}"
	}

	return strings.Join([]string{"DeleteLogicalClusterRequest", string(data)}, " ")
}
