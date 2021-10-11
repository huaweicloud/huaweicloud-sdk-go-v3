package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowExternalEntityRequest struct {
	// 边缘节点ID

	EdgeNodeId string `json:"edge_node_id"`
	// 外部实体ID

	ExternalId string `json:"external_id"`
}

func (o ShowExternalEntityRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowExternalEntityRequest struct{}"
	}

	return strings.Join([]string{"ShowExternalEntityRequest", string(data)}, " ")
}
