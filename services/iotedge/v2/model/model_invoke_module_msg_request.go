package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokeModuleMsgRequest Request Object
type InvokeModuleMsgRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 边缘模块ID
	ModuleId string `json:"module_id"`

	Body *interface{} `json:"body,omitempty"`
}

func (o InvokeModuleMsgRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokeModuleMsgRequest struct{}"
	}

	return strings.Join([]string{"InvokeModuleMsgRequest", string(data)}, " ")
}
