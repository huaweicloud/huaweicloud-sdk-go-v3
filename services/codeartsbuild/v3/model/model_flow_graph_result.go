package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlowGraphResult 结果
type FlowGraphResult struct {

	// edges
	Edges *[]FlowGraphResultEdges `json:"edges,omitempty"`

	// record信息
	Vertices *[]Vertices `json:"vertices,omitempty"`
}

func (o FlowGraphResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowGraphResult struct{}"
	}

	return strings.Join([]string{"FlowGraphResult", string(data)}, " ")
}
