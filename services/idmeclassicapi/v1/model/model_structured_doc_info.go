package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StructuredDocInfo struct {

	// **参数解释**：  文档ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：  kooPage文档ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	DocumentId *string `json:"document_id,omitempty"`

	// **参数解释**：  文档标题。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Title *string `json:"title,omitempty"`

	// **参数解释**：  文档类型。  **约束限制**：  不涉及。  **取值范围**：  - directory：目录。 - pageDocument：Page文档。 - boardDocument：Board文档。 - mindDocument：Mind文档。 - drawDocument：Draw文档。  **默认取值**：  不涉及。
	Type *string `json:"type,omitempty"`

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

	// **参数解释**：  模型名称。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ModelName *string `json:"model_name,omitempty"`

	// **参数解释**：  创建者ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	CreateUserId *string `json:"create_user_id,omitempty"`

	// **参数解释**：  修改人。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o StructuredDocInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructuredDocInfo struct{}"
	}

	return strings.Join([]string{"StructuredDocInfo", string(data)}, " ")
}
