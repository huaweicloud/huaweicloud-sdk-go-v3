package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateModuleRequest struct {
	// 边缘节点ID

	EdgeNodeId string `json:"edge_node_id"`

	Body *CreateEdgeModuleReqDto `json:"body,omitempty"`
}

func (o CreateModuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateModuleRequest struct{}"
	}

	return strings.Join([]string{"CreateModuleRequest", string(data)}, " ")
}
