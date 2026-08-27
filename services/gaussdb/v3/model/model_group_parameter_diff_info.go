package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GroupParameterDiffInfo **参数解释**：  参数差异信息。
type GroupParameterDiffInfo struct {

	// **参数解释**：  参数名称。  **取值范围**：  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：  源参数值（当前模板中的值）。  **取值范围**：  不涉及。
	SourceValue *string `json:"source_value,omitempty"`

	// **参数解释**：  目标参数值（系统默认模板中的值）。  **取值范围**：  不涉及。
	TargetValue *string `json:"target_value,omitempty"`
}

func (o GroupParameterDiffInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupParameterDiffInfo struct{}"
	}

	return strings.Join([]string{"GroupParameterDiffInfo", string(data)}, " ")
}
