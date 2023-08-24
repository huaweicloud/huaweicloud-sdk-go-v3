package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateModuleShadowRequest Request Object
type UpdateModuleShadowRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 边缘模块ID
	ModuleId string `json:"module_id"`

	Body *UpdateModuleShadowsRequestBody `json:"body,omitempty"`
}

func (o UpdateModuleShadowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModuleShadowRequest struct{}"
	}

	return strings.Join([]string{"UpdateModuleShadowRequest", string(data)}, " ")
}
