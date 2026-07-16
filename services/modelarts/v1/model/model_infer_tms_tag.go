package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InferTmsTag **参数解释：** Tms的标签结构体。 **取值范围：** 不涉及。
type InferTmsTag struct {

	// **参数解释：** 标签键字段。 **约束限制：** 不涉及。 **取值范围：** 长度为1~128，标签的键可以包含任意语种字母、数字、空格，以及_ . : = + - @特殊字符，但首尾不能含有空格，不能以_sys_开头。 **默认取值：** 不涉及。
	Key string `json:"key"`

	// **参数解释：** 标签值字段。 **约束限制：** 字母、数字、下划线、点、斜杠、等号、加号、减号和@符号。 **取值范围：** 长度为0~255。 **默认取值：** 不涉及。
	Value string `json:"value"`
}

func (o InferTmsTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InferTmsTag struct{}"
	}

	return strings.Join([]string{"InferTmsTag", string(data)}, " ")
}
