package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteDocRequestDto struct {

	// **参数解释**：  文档ID列表。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Ids *[]string `json:"ids,omitempty"`

	// **参数解释**：  模型名称。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ModelName *string `json:"model_name,omitempty"`

	// **参数解释**：  是否检查文档删除权限。  **约束限制**：  不涉及。  **取值范围**：  - true：检查。 - false：不检查。  **默认取值**：  true。
	IsCheck *bool `json:"is_check,omitempty"`
}

func (o BatchDeleteDocRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDocRequestDto struct{}"
	}

	return strings.Join([]string{"BatchDeleteDocRequestDto", string(data)}, " ")
}
