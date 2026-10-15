package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkLabel **参数解释**：标签键值对，用于资源分类和筛选。 **约束限制**：不涉及。
type FlinkLabel struct {

	// **参数解释**：标签键。 **约束限制**：不涉及。 **取值范围**：1~128个字符，不支持输入等号。 **默认取值**：不涉及。
	Key string `json:"key"`

	// **参数解释**：标签值。 **约束限制**：不涉及。 **取值范围**：1~512个字符，字符组成不限制。 **默认取值**：不涉及。
	Value string `json:"value"`
}

func (o FlinkLabel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkLabel struct{}"
	}

	return strings.Join([]string{"FlinkLabel", string(data)}, " ")
}
