package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceCreateRequestTags struct {

	// **参数解释：** 标签键字段。 **约束限制：** 不涉及。 **取值范围：** 长度为1~128，标签的键可以包含任意语种字母、数字、空格，以及_ . : = + - @特殊字符，但首尾不能含有空格，不能以_sys_开头。 **默认取值：** 不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释：** 标签value值。 **约束限制：** 最大长度为256。 **取值范围：** 任意数量的字母、数字、空格、下划线、点、冒号、斜杠、等号、加号、减号、@等字符开始和结束的字符串。 **默认取值：** 不涉及。
	Value *string `json:"value,omitempty"`
}

func (o ServiceCreateRequestTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceCreateRequestTags struct{}"
	}

	return strings.Join([]string{"ServiceCreateRequestTags", string(data)}, " ")
}
