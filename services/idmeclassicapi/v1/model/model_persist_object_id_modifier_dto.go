package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistObjectIdModifierDto struct {

	// **参数解释：**  唯一标识。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id string `json:"id"`

	// **参数解释：**  请求参数对象。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o PersistObjectIdModifierDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistObjectIdModifierDto struct{}"
	}

	return strings.Join([]string{"PersistObjectIdModifierDto", string(data)}, " ")
}
