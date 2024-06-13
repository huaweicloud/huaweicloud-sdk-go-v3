package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEdgeNodeRequest Request Object
type UpdateEdgeNodeRequest struct {

	// 节点id
	EdgeNodeId string `json:"edge_node_id"`

	Body *UpdateNodeReqDto `json:"body,omitempty"`
}

func (o UpdateEdgeNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeNodeRequest struct{}"
	}

	return strings.Join([]string{"UpdateEdgeNodeRequest", string(data)}, " ")
}
