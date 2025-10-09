package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplicableInstanceInfo **参数解释**：  可应用参数模板的实例详情  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
type ApplicableInstanceInfo struct {

	// **参数解释**：  实例id  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EntityId string `json:"entity_id"`

	// **参数解释**：  实例名称  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EntityName string `json:"entity_name"`
}

func (o ApplicableInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicableInstanceInfo struct{}"
	}

	return strings.Join([]string{"ApplicableInstanceInfo", string(data)}, " ")
}
