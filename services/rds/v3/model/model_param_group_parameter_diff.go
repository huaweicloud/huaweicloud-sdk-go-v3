package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParamGroupParameterDiff 参数差异信息
type ParamGroupParameterDiff struct {

	// **参数解释**：  参数名称  **约束限制**：  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：  源模板参数值。  **约束限制**：  不涉及。
	SourceValue *string `json:"source_value,omitempty"`

	// **参数解释**：  目标模板参数值。  **约束限制**：  不涉及。
	TargetValue *string `json:"target_value,omitempty"`
}

func (o ParamGroupParameterDiff) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParamGroupParameterDiff struct{}"
	}

	return strings.Join([]string{"ParamGroupParameterDiff", string(data)}, " ")
}
