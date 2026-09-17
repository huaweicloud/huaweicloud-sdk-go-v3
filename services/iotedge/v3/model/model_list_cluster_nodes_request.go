package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterNodesRequest Request Object
type ListClusterNodesRequest struct {

	// 边缘集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ListClusterNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterNodesRequest struct{}"
	}

	return strings.Join([]string{"ListClusterNodesRequest", string(data)}, " ")
}
