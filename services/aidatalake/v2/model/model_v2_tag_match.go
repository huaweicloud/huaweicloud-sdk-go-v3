package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2TagMatch **参数解释**：搜索字段，key为要匹配的字段，如resource_name等。value为匹配的值。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type V2TagMatch struct {

	// **参数解释**：固定字典值。 **约束限制**：不涉及。 **取值范围**：resource_name，表示资源名称。 **默认取值**：不涉及。
	Key string `json:"key"`

	// **参数解释**：基于“key”对应字段的前缀匹配值。 **约束限制**：最大255个字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Value *string `json:"value,omitempty"`
}

func (o V2TagMatch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2TagMatch struct{}"
	}

	return strings.Join([]string{"V2TagMatch", string(data)}, " ")
}
