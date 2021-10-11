package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowEdgeNodeRequest struct {
	// 边缘节点ID

	EdgeNodeId string `json:"edge_node_id"`
}

func (o ShowEdgeNodeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowEdgeNodeRequest struct{}"
	}

	return strings.Join([]string{"ShowEdgeNodeRequest", string(data)}, " ")
}
