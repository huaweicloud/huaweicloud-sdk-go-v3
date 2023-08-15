package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TraceTopologyNode 调用链拓扑图的节点。
type TraceTopologyNode struct {

	// 节点id。
	NodeId *int64 `json:"node_id,omitempty"`

	// 节点名称。
	NodeName *string `json:"node_name,omitempty"`

	// 节点提示字段。
	Hint *string `json:"hint,omitempty"`
}

func (o TraceTopologyNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TraceTopologyNode struct{}"
	}

	return strings.Join([]string{"TraceTopologyNode", string(data)}, " ")
}
