package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ParamGroupParameterDiffV3 struct {

	// **参数解释**：  参数名称。  **参数范围**：  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：  源参数值。  **参数范围**：  不涉及。
	SourceValue *string `json:"source_value,omitempty"`

	// **参数解释**：  目标参数值。  **参数范围**：  不涉及。
	TargetValue *string `json:"target_value,omitempty"`
}

func (o ParamGroupParameterDiffV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParamGroupParameterDiffV3 struct{}"
	}

	return strings.Join([]string{"ParamGroupParameterDiffV3", string(data)}, " ")
}
