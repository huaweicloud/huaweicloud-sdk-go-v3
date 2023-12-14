package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterNodesResponse Response Object
type ListClusterNodesResponse struct {

	// 逻辑集群节点列表
	NodeList *[]ClusterNodeInfo `json:"node_list,omitempty"`

	// 逻辑集群节点总数
	Count *int32 `json:"count,omitempty"`

	// 逻辑集群节点失败总数
	FailedCount    *int32 `json:"failed_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListClusterNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterNodesResponse struct{}"
	}

	return strings.Join([]string{"ListClusterNodesResponse", string(data)}, " ")
}
