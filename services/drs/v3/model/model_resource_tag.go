package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceTag 标签信息体
type ResourceTag struct {

	// **参数解释**：  标签键。  **约束限制**：  - 最大长度36个字符。 - 只能包含字母、数字、下划线、中划线和中文。 - 需要修改标签时填写。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释**：  标签值。  **约束限制**：  - 最大长度43个字符。 - 只能包含字母、数字、下划线、中划线和中文。 - 需要修改标签时填写。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Value *string `json:"value,omitempty"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
