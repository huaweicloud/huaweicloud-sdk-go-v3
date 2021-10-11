package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteModuleRequest struct {
	// 边缘节点ID

	EdgeNodeId string `json:"edge_node_id"`
	// 边缘模块ID

	ModuleId string `json:"module_id"`
}

func (o DeleteModuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteModuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteModuleRequest", string(data)}, " ")
}
