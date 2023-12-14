package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClusterNodesRequest Request Object
type DeleteClusterNodesRequest struct {

	// 指定待删除节点集群的ID
	ClusterId string `json:"cluster_id"`

	Body *DeleteClusterNodesRequestBody `json:"body,omitempty"`
}

func (o DeleteClusterNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterNodesRequest struct{}"
	}

	return strings.Join([]string{"DeleteClusterNodesRequest", string(data)}, " ")
}
