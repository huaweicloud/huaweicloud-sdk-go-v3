package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigurationUpdateReqV3 struct {

	// **参数解释**：  描述。  **约束限制**：  不涉及  **取值范围**：  0-256的不是!、<、>、=、&、\" 或 ' 的字符。  **默认取值**：  不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：  修改的值。  **约束限制**：  不涉及  **取值范围**：  长度为1-64的a-z、A-Z、0-9、.、_ 和 -的字符。  **默认取值**：  不涉及。
	Values map[string]string `json:"values"`

	// **参数解释**：  参数的名称。  **约束限制**：  不涉及  **取值范围**：  长度为1-64的a-z、A-Z、0-9、.、_ 和 -的字符。  **默认取值**：  不涉及。
	Name *string `json:"name,omitempty"`
}

func (o ConfigurationUpdateReqV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationUpdateReqV3 struct{}"
	}

	return strings.Join([]string{"ConfigurationUpdateReqV3", string(data)}, " ")
}
