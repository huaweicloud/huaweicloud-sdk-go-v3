package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateStructuredDocDto struct {

	// **参数解释**：  文档ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Id string `json:"id"`

	// **参数解释**：  文档标题。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Title *string `json:"title,omitempty"`

	// **参数解释**：  模板ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	TemplateId *string `json:"template_id,omitempty"`

	// **参数解释**：  团队ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	OrganizationId *string `json:"organization_id,omitempty"`

	// **参数解释**：  知识库ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	WikiId *string `json:"wiki_id,omitempty"`

	// **参数解释**：  父文档ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ParentDocumentId *string `json:"parent_document_id,omitempty"`

	// **参数解释**：  实例ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  修改人。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o BatchUpdateStructuredDocDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateStructuredDocDto struct{}"
	}

	return strings.Join([]string{"BatchUpdateStructuredDocDto", string(data)}, " ")
}
