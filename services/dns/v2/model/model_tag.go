package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Tag struct {

	// **参数解释：** 标签键。 **取值范围：** - 不能为空。 - 对于同一资源键的取值唯一。 - 长度不超过128个字符。 - 取值可以包含任意语种字母、数字、空格，以及_ . : = + - @特殊字符，但首尾不能含有空格，不能以_sys_开头。
	Key string `json:"key"`

	// **参数解释：** 标签值。 **取值范围：** - 可以为空。 - 长度不超过255个字符。 - 取值可以包含任意语种字母、数字、空格，以及_ . : / = + - @特殊字符。
	Value *string `json:"value,omitempty"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
