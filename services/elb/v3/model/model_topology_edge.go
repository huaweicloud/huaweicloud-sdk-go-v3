package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopologyEdge
type TopologyEdge struct {

	// **参数解释**：边起点id。  **取值范围**：不涉及
	Source string `json:"source"`

	// **参数解释**：边终点id。  **取值范围**：不涉及
	Target string `json:"target"`

	// **参数解释**：边起点类型。  **取值范围**：不涉及
	SourceType string `json:"source_type"`

	// **参数解释**：边终点类型。  **取值范围**：不涉及
	TargetType string `json:"target_type"`

	// **参数解释**：边label信息。  **取值范围**：不涉及
	Label *string `json:"label,omitempty"`

	// **参数解释**：label id。  **取值范围**：不涉及
	LabelId *string `json:"label_id,omitempty"`
}

func (o TopologyEdge) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopologyEdge struct{}"
	}

	return strings.Join([]string{"TopologyEdge", string(data)}, " ")
}
