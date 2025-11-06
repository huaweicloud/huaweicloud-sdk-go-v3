package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagsRequestBody 标签字段信息
type TagsRequestBody struct {

	// **参数解释：** 标签键。 **约束限制：** - 标签键名称不可重复。 - 标签是以键值对（key-value）的形式表示，key和value为一一对应关系。 **取值范围：** 标签键可以包含任意语种的字母、数字和空格，以及_.:=+-@字符，但首尾不能包含空格，且不能以_sys_开头；长度不能超过128个字符。 **默认取值：** 不涉及。
	Key string `json:"key"`

	// **参数解释：** 标签值。 **约束限制：** 标签是以键值对（key-value）的形式表示，key和value为一一对应关系。 **取值范围：** - 标签值可以包含任意语种的字母、数字和空格，以及_.:=+-@字符。 - 标签值长度不能超过255个字符。 **默认取值：** 不涉及。
	Value *string `json:"value,omitempty"`
}

func (o TagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsRequestBody struct{}"
	}

	return strings.Join([]string{"TagsRequestBody", string(data)}, " ")
}
