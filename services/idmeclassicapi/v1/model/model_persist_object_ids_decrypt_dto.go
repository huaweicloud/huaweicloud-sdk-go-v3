package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistObjectIdsDecryptDto struct {

	// **参数解释：**  是否加密。  **约束限制：**  不涉及。  **取值范围：**  - true：加密。 - false：不加密。  **默认取值：**  false。
	Decrypt *bool `json:"decrypt,omitempty"`

	// **参数解释：**  ID列表。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Ids []string `json:"ids"`
}

func (o PersistObjectIdsDecryptDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistObjectIdsDecryptDto struct{}"
	}

	return strings.Join([]string{"PersistObjectIdsDecryptDto", string(data)}, " ")
}
