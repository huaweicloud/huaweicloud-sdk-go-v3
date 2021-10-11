package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateRoutesRequest struct {
	// 边缘节点ID

	EdgeNodeId string `json:"edge_node_id"`

	Body *[]CreateRouterReqDto `json:"body,omitempty"`
}

func (o UpdateRoutesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRoutesRequest struct{}"
	}

	return strings.Join([]string{"UpdateRoutesRequest", string(data)}, " ")
}
