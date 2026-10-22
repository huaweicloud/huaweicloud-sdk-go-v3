package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2ResourceTagParam **参数解释**：标签查询参数。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type V2ResourceTagParam struct {

	// **参数解释**：资源标签的键。 **约束限制**：不涉及。 **取值范围**：长度1~128的多语言字母、空格、数字、下划线、点、冒号、斜杠、等号、加号、减号、@符号组合。 **默认取值**：不涉及。
	Key string `json:"key"`

	// **参数解释**：资源标签的值列表。 **约束限制**：最大20个字符串。每个字符串最大255个字符。
	Values *[]string `json:"values,omitempty"`
}

func (o V2ResourceTagParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2ResourceTagParam struct{}"
	}

	return strings.Join([]string{"V2ResourceTagParam", string(data)}, " ")
}
