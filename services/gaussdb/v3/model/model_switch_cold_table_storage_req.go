package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchColdTableStorageReq struct {

	// **参数解释**：  操作类型。  **约束限制**：  不涉及。  **取值范围**：  - archive：创建冷表。 - unarchive：回迁冷表。  **默认取值**：  不涉及。
	Action string `json:"action"`

	// **参数解释**：  需要操作的表信息。  **约束限制**：  数组长度小于等于50。
	Tables []TableTarget `json:"tables"`
}

func (o SwitchColdTableStorageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchColdTableStorageReq struct{}"
	}

	return strings.Join([]string{"SwitchColdTableStorageReq", string(data)}, " ")
}
