package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlowGraphResultEdges struct {

	// 依赖子任务ID
	From *string `json:"from,omitempty"`

	// 被依赖的子任务ID
	To *string `json:"to,omitempty"`
}

func (o FlowGraphResultEdges) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowGraphResultEdges struct{}"
	}

	return strings.Join([]string{"FlowGraphResultEdges", string(data)}, " ")
}
