package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowModuleRequest struct {
	// 边缘节点ID

	EdgeNodeId string `json:"edge_node_id"`
	// 边缘模块ID

	ModuleId string `json:"module_id"`
}

func (o ShowModuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowModuleRequest struct{}"
	}

	return strings.Join([]string{"ShowModuleRequest", string(data)}, " ")
}
