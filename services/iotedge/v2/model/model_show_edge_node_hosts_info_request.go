package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEdgeNodeHostsInfoRequest Request Object
type ShowEdgeNodeHostsInfoRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`
}

func (o ShowEdgeNodeHostsInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeNodeHostsInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowEdgeNodeHostsInfoRequest", string(data)}, " ")
}
