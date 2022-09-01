package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 调用链拓扑图的节点
type TraceTopologyNode struct {

	// 节点id
	NodeId *int64 `json:"node_id,omitempty" xml:"node_id"`

	// 节点名称
	NodeName *string `json:"node_name,omitempty" xml:"node_name"`

	// 节点提示字段
	Hint *string `json:"hint,omitempty" xml:"hint"`
}

func (o TraceTopologyNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TraceTopologyNode struct{}"
	}

	return strings.Join([]string{"TraceTopologyNode", string(data)}, " ")
}
