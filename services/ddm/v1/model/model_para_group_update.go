package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ParaGroupUpdate struct {

	// **参数解释**：  参数组名称。  **约束限制**：  不涉及。  **取值范围**：  在1到64个字符之间，区分大小写，可包含字母、数字、中划线、下划线或句点，不能包含其他特殊字符。。  **默认取值**：  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：  参数组描述。  **约束限制**：  不涉及。  **取值范围**：  不能超过256位，且不能包含回车和特殊字符 ! < \" = ' > &。  **默认取值**：  不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：  修改的值。  **约束限制**：  不涉及  **取值范围**：  长度为1-64的a-z、A-Z、0-9、.、_ 和 -的字符。  **默认取值**：  不涉及。
	Values map[string]string `json:"values"`
}

func (o ParaGroupUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParaGroupUpdate struct{}"
	}

	return strings.Join([]string{"ParaGroupUpdate", string(data)}, " ")
}
