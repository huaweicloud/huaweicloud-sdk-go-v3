package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagsResponse **参数解释：** Tms的标签结构体。
type TagsResponse struct {

	// **参数解释：** 标签键字段。 **约束限制：** 不涉及。 **取值范围：** 长度为1~128，标签的键可以包含任意语种字母、数字、空格，以及_ . : = + - @特殊字符，但首尾不能含有空格，不能以_sys_开头。 **默认取值：** 不涉及。
	Key string `json:"key"`

	// **参数解释：** 标签值字段。 **取值范围：** 长度为255，包括字母、数字、下划线、点、斜杠、等号、加号、减号和@符号。
	Value string `json:"value"`
}

func (o TagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsResponse struct{}"
	}

	return strings.Join([]string{"TagsResponse", string(data)}, " ")
}
