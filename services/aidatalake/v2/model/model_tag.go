package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tag **参数解释**：标签。 **约束限制**：不涉及。
type Tag struct {

	// **参数解释**：标签键。 **约束限制**：不涉及。 **取值范围**：长度为1~128个英文字符、数字、下划线、中划线、句点的组合。 **默认取值**：不涉及。
	Key string `json:"key"`

	// **参数解释**：标签值。 **约束限制**：不涉及。 **取值范围**：长度为0~255个字符。 **默认取值**：不涉及。
	Value *string `json:"value,omitempty"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
