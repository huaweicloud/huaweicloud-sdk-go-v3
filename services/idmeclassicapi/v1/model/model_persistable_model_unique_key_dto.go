package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistableModelUniqueKeyDto struct {

	// **参数解释：**  是否加密。  **约束限制：**  不涉及。  **取值范围：**  - true：加密。 - false：不加密。  **默认取值：**  false。
	Decrypt *bool `json:"decrypt,omitempty"`

	// **参数解释：**  示例模型的唯一键约束属性值。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	UniqueKey string `json:"uniqueKey"`
}

func (o PersistableModelUniqueKeyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistableModelUniqueKeyDto struct{}"
	}

	return strings.Join([]string{"PersistableModelUniqueKeyDto", string(data)}, " ")
}
