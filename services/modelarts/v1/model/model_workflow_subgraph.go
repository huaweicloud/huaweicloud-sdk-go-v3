package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowSubgraph 子图。
type WorkflowSubgraph struct {

	// 子图名称。
	Name *string `json:"name,omitempty"`

	// 子图step成员。
	Steps *[]string `json:"steps,omitempty"`
}

func (o WorkflowSubgraph) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowSubgraph struct{}"
	}

	return strings.Join([]string{"WorkflowSubgraph", string(data)}, " ")
}
