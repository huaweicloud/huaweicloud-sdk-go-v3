package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNodeRequest Request Object
type UpdateNodeRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 节点ID
	NodeId string `json:"node_id"`

	Body *NodeExpansionInfo `json:"body,omitempty"`
}

func (o UpdateNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodeRequest struct{}"
	}

	return strings.Join([]string{"UpdateNodeRequest", string(data)}, " ")
}
