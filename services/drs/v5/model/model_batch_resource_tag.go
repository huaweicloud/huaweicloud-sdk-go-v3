package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchResourceTag **参数解释**：  标签键值对信息，包含标签键和标签值。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
type BatchResourceTag struct {

	// **参数解释**：  标签键。  **约束限制**：  - 长度为1-128个unicode字符。 - 可以包含任何语种字母、数字、空格和_.:=+-@，但首尾不能含有空格，不能以_sys_开头。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Key string `json:"key"`

	// **参数解释**：  标签值，删除标签时非必填。  **约束限制**：  - 最大长度255个unicode字符。 - 可以包含任何语种字母、数字、空格和_.:=+-@。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Value *string `json:"value,omitempty"`
}

func (o BatchResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResourceTag struct{}"
	}

	return strings.Join([]string{"BatchResourceTag", string(data)}, " ")
}
