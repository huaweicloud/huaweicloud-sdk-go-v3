package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CharacterSetEnum struct {

	// **参数解释：**  中文名。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CnName *string `json:"cnName,omitempty"`

	// **参数解释：**  编码。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释：**  英文名。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	EnName *string `json:"enName,omitempty"`
}

func (o CharacterSetEnum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CharacterSetEnum struct{}"
	}

	return strings.Join([]string{"CharacterSetEnum", string(data)}, " ")
}
