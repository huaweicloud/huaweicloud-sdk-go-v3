package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 组件节点。
type TxNode struct {

	// 节点id。
	TxNodeId *string `json:"tx_node_id,omitempty"`

	// 节点名称。
	TxNodeName *string `json:"tx_node_name,omitempty"`

	// 节点类型。
	TxNodeType *string `json:"tx_node_type,omitempty"`
}

func (o TxNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TxNode struct{}"
	}

	return strings.Join([]string{"TxNode", string(data)}, " ")
}
