package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateModuleRequest Request Object
type UpdateModuleRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 边缘模块ID
	ModuleId string `json:"module_id"`

	Body *UpdateEdgeModuleReqDto `json:"body,omitempty"`
}

func (o UpdateModuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateModuleRequest", string(data)}, " ")
}
