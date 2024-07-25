package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistObjectIdDecryptDto struct {

	// **参数解释：**  是否加密。  **约束限制：**  不涉及。  **取值范围：**  - true：加密。 - false：不加密。  **默认取值：**  false。
	Decrypt *bool `json:"decrypt,omitempty"`

	// **参数解释：**  唯一标识。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id string `json:"id"`
}

func (o PersistObjectIdDecryptDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistObjectIdDecryptDto struct{}"
	}

	return strings.Join([]string{"PersistObjectIdDecryptDto", string(data)}, " ")
}
