package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDocRequestDto struct {

	// **参数解释**：  kooPage文档ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	DocumentId *string `json:"document_id,omitempty"`

	// **参数解释**：  文档标题。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Title *string `json:"title,omitempty"`

	// **参数解释**：  实例ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释：**  修改人。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o UpdateDocRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDocRequestDto struct{}"
	}

	return strings.Join([]string{"UpdateDocRequestDto", string(data)}, " ")
}
