package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceTag **参数解释**：资源标签。
type ResourceTag struct {

	// **参数解释**：资源标签的键。 **约束限制**：不能以空格开头或结尾，也不能以_sys_开头。 **取值范围**：长度1~128的多语言字母、空格、数字、下划线、点、冒号、斜杠、等号、加号、减号、@符号组合。 **默认取值**：不涉及。
	Key string `json:"key"`

	// **参数解释**：资源标签的值。 **约束限制**：不涉及。 **取值范围**：长度0~255的多语言字母、空格、数字、下划线、点、冒号、斜杠、等号、加号、减号、@符号组合。 **默认取值**：不涉及。
	Value *string `json:"value,omitempty"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
