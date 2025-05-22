package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlowGraph2Result 结果
type FlowGraph2Result struct {

	// edges
	Edges *[]FlowGraph2ResultEdges `json:"edges,omitempty"`

	// record信息
	Vertices *[]Vertices `json:"vertices,omitempty"`
}

func (o FlowGraph2Result) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowGraph2Result struct{}"
	}

	return strings.Join([]string{"FlowGraph2Result", string(data)}, " ")
}
