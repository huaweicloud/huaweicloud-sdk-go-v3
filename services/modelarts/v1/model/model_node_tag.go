package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeTag **参数解释**：节点资源标签。 **约束限制**：不涉及。
type NodeTag struct {

	// **参数解释**：键。标签的键可以包含任意语种的字母、数字和空格，以及_.:=+-@字符，但首尾不能包含空格，且不能以_sys_开头。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Key string `json:"key"`

	// **参数解释**：值。标签的值可以包含任意语种的字母、数字和空格，以及_.:=+-@字符，但首尾不能包含空格。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Value string `json:"value"`
}

func (o NodeTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeTag struct{}"
	}

	return strings.Join([]string{"NodeTag", string(data)}, " ")
}
