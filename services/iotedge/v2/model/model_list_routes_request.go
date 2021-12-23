package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRoutesRequest struct {
	// 边缘节点ID

	EdgeNodeId string `json:"edge_node_id"`
	// 是否解析路由

	Parsed *bool `json:"parsed,omitempty"`
}

func (o ListRoutesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRoutesRequest struct{}"
	}

	return strings.Join([]string{"ListRoutesRequest", string(data)}, " ")
}
