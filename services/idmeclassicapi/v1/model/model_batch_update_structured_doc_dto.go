package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateStructuredDocDto struct {

	// **参数解释：**  KooPage文档ID，系统生成的文档主键唯一标识，用于定位待更新的目标文档。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Id string `json:"id"`

	// **参数解释：**  文档标题，用于更新文档的名称信息。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Title *string `json:"title,omitempty"`

	// **参数解释：**  模板ID，用于更新文档所使用的模板。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	TemplateId *string `json:"template_id,omitempty"`

	// **参数解释：**  团队ID，用于更新文档所属的团队组织。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	OrganizationId *string `json:"organization_id,omitempty"`

	// **参数解释：**  知识库ID，用于更新文档所属的知识库。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	WikiId *string `json:"wiki_id,omitempty"`

	// **参数解释：**  父文档ID，用于更新文档的父级目录归属。 例如将多个文档批量迁移到新的目录下。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ParentDocumentId *string `json:"parent_document_id,omitempty"`

	// **参数解释：**  实例ID，用于更新文档关联的数据模型实例。 例如将文档从旧产品实例关联到新产品实例。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释：**  更新者账号，用于记录执行本次批量更新操作的用户信息。 若不指定，默认使用当前调用者账号。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  当前调用者账号。
	Modifier *string `json:"modifier,omitempty"`
}

func (o BatchUpdateStructuredDocDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateStructuredDocDto struct{}"
	}

	return strings.Join([]string{"BatchUpdateStructuredDocDto", string(data)}, " ")
}
