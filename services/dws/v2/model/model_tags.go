package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tags **参数解释**： 标签列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type Tags struct {

	// **参数解释**： 标签的键。 **约束限制**： 不涉及。 **取值范围**： - 输入标签键的最大长度为128个unicode字符，不能为空字符串，且首尾字符不能为空格； - 不能包含“=”,“*”,“<”,“>”,“\\\\”,“,”,“|”,“/”； - 只能包含大写字母（A-Z）、小写字母（a-z）、数字（0-9）和特殊字符（中划线-、下划线_）以及中文字符；  **默认取值**： 不涉及。
	Key string `json:"key"`

	// **参数解释**： 标签的值。 **约束限制**： 不涉及。 **取值范围**： - 输入标签值的最大长度为256个字符，首尾字符不能为空格，可以为空字符串。 - 不能包含“=”,“*”,“<”,“>”,“\\\\”,“,”,“|”,“/”。 - 只能包含大写字母（A-Z）、小写字母（a-z）、数字（0-9）和特殊字符（中划线-、下划线_）以及中文字符。  **默认取值**： 不涉及。
	Value string `json:"value"`
}

func (o Tags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tags struct{}"
	}

	return strings.Join([]string{"Tags", string(data)}, " ")
}
