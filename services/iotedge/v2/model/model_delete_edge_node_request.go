package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteEdgeNodeRequest struct {
	// 边缘节点ID

	EdgeNodeId string `json:"edge_node_id"`
	// 是否同时删除外部节点（仅对高级版有效），默认为false不删除IEF侧的边缘节点

	DeleteExternalNode *bool `json:"delete_external_node,omitempty"`
}

func (o DeleteEdgeNodeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteEdgeNodeRequest struct{}"
	}

	return strings.Join([]string{"DeleteEdgeNodeRequest", string(data)}, " ")
}
