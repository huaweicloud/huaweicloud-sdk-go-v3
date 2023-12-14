package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeClusterWithExistedNodesRequest Request Object
type ResizeClusterWithExistedNodesRequest struct {

	// 指定节点集群的ID
	ClusterId string `json:"cluster_id"`

	Body *ResizeClusterWithExistedNodesRequestBody `json:"body,omitempty"`
}

func (o ResizeClusterWithExistedNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeClusterWithExistedNodesRequest struct{}"
	}

	return strings.Join([]string{"ResizeClusterWithExistedNodesRequest", string(data)}, " ")
}
