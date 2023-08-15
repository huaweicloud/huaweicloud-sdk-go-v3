package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateModuleStateRequest Request Object
type UpdateModuleStateRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 边缘模块ID
	ModuleId string `json:"module_id"`

	Body *UpdateEdgeModuleStateReqDto `json:"body,omitempty"`
}

func (o UpdateModuleStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModuleStateRequest struct{}"
	}

	return strings.Join([]string{"UpdateModuleStateRequest", string(data)}, " ")
}
