package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TmsTag TMS的标签结构体。
type TmsTag struct {

	// **参数解释**：TMS标签的key。 **约束限制**：长度限制为128个字符，支持任意语种字母、数字、空格，以及_ . : = + - @特殊字符，但首尾不能含有空格，不能以_sys_开头。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Key string `json:"key"`

	// **参数解释**：TMS标签的value。 **约束限制**：长度限制为255个字符，支持任意语种字母、数字、空格，以及_ . : / = + - @特殊字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Value *string `json:"value,omitempty"`
}

func (o TmsTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsTag struct{}"
	}

	return strings.Join([]string{"TmsTag", string(data)}, " ")
}
