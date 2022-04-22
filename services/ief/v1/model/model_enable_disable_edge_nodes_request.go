package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type EnableDisableEdgeNodesRequest struct {

	// 节点ID
	NodeId string `json:"node_id"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	Body *NodeAction `json:"body,omitempty"`
}

func (o EnableDisableEdgeNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDisableEdgeNodesRequest struct{}"
	}

	return strings.Join([]string{"EnableDisableEdgeNodesRequest", string(data)}, " ")
}
