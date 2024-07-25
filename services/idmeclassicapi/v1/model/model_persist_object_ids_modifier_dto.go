package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistObjectIdsModifierDto struct {

	// **参数解释：**  数据实例ID列表。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Ids []string `json:"ids"`

	// **参数解释：**  修改人。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o PersistObjectIdsModifierDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistObjectIdsModifierDto struct{}"
	}

	return strings.Join([]string{"PersistObjectIdsModifierDto", string(data)}, " ")
}
