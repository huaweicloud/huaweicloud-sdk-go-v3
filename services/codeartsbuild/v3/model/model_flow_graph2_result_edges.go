package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlowGraph2ResultEdges struct {

	// 依赖子任务ID
	From *string `json:"from,omitempty"`

	// 被依赖的子任务ID
	To *string `json:"to,omitempty"`
}

func (o FlowGraph2ResultEdges) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowGraph2ResultEdges struct{}"
	}

	return strings.Join([]string{"FlowGraph2ResultEdges", string(data)}, " ")
}
