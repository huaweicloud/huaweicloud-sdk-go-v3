package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2ResourceTag **参数解释**：资源标签。
type V2ResourceTag struct {

	// **参数解释**：资源标签的键。 **约束限制**：不涉及。 **取值范围**：长度1~128的多语言字母、空格、数字、下划线、点、冒号、斜杠、等号、加号、减号、@符号组合。 **默认取值**：不涉及。
	Key string `json:"key"`

	// **参数解释**：资源标签的值。 **约束限制**：不涉及。 **取值范围**：长度0~255的多语言字母、空格、数字、下划线、点、冒号、斜杠、等号、加号、减号、@符号组合。 **默认取值**：不涉及。
	Value *string `json:"value,omitempty"`
}

func (o V2ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2ResourceTag struct{}"
	}

	return strings.Join([]string{"V2ResourceTag", string(data)}, " ")
}
