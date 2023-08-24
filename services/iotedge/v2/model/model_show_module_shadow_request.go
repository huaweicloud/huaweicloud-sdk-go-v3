package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowModuleShadowRequest Request Object
type ShowModuleShadowRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 边缘模块ID
	ModuleId string `json:"module_id"`
}

func (o ShowModuleShadowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowModuleShadowRequest struct{}"
	}

	return strings.Join([]string{"ShowModuleShadowRequest", string(data)}, " ")
}
