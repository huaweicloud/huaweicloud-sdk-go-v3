package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowSubgraphResp 子图。
type WorkflowSubgraphResp struct {

	// **参数解释**：子图名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：子图step成员。
	Steps *[]string `json:"steps,omitempty"`
}

func (o WorkflowSubgraphResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowSubgraphResp struct{}"
	}

	return strings.Join([]string{"WorkflowSubgraphResp", string(data)}, " ")
}
