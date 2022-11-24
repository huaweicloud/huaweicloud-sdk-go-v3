package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 组件节点。
type TopoNode struct {

	// 节点类型。
	NodeType *string `json:"node_type,omitempty"`

	// 节点名称。
	NodeName *string `json:"node_name,omitempty"`

	// 节点id。
	NodeId *string `json:"node_id,omitempty"`

	// 环境id。
	EnvId *int64 `json:"env_id,omitempty"`
}

func (o TopoNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopoNode struct{}"
	}

	return strings.Join([]string{"TopoNode", string(data)}, " ")
}
