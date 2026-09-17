package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEdgeNodeSoftwareVersionRequest Request Object
type ShowEdgeNodeSoftwareVersionRequest struct {

	// 节点id
	EdgeNodeId string `json:"edge_node_id"`
}

func (o ShowEdgeNodeSoftwareVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeNodeSoftwareVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowEdgeNodeSoftwareVersionRequest", string(data)}, " ")
}
