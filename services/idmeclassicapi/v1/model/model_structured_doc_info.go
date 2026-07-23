package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StructuredDocInfo struct {

	// **参数解释：**  文档ID，用于指定自定义文档唯一标识。 若不指定，系统会自动生成。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  KooPage文档ID，用于关联已有的KooPage文档。 若创建新文档，可不指定。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	DocumentId *string `json:"document_id,omitempty"`

	// **参数解释：**  文档标题，用于标识结构化文档的名称。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Title *string `json:"title,omitempty"`

	// **参数解释：**  文档类型，指定结构化文档的类型。  **约束限制：**  不涉及。  **取值范围：**  - directory：目录，用于组织和管理文档层级结构。 - pageDocument：Page文档，适用于富文本编辑场景，如设计说明书、技术文档等。 - boardDocument：Board文档，适用于白板协作场景，如工艺评审、方案讨论等。 - mindDocument：Mind文档，适用于思维导图场景，如产品结构分析、流程梳理等。 - drawDocument：Draw文档，适用于绘图场景，如工艺流程图、设备布局图等。  **默认取值：**  不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释：**  模板ID，用于指定文档创建时所使用的模板。 若指定模板，创建的文档将继承模板的格式和内容。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	TemplateId *string `json:"template_id,omitempty"`

	// **参数解释：**  团队ID，用于指定文档所属的团队组织。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	OrganizationId *string `json:"organization_id,omitempty"`

	// **参数解释：**  知识库ID，用于指定文档所属的知识库。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	WikiId *string `json:"wiki_id,omitempty"`

	// **参数解释：**  父文档ID，用于指定文档的父级目录。 若指定，创建的文档将作为该父文档的子文档。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ParentDocumentId *string `json:"parent_document_id,omitempty"`

	// **参数解释：**  实例ID，用于将结构化文档关联到指定的数据模型实例。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释：**  模型名称，用于指定文档关联的数据模型名称。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ModelName *string `json:"model_name,omitempty"`

	// **参数解释：**  创建者ID，用于指定文档的创建者。 若不指定，默认使用当前调用者ID。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  当前调用者ID。
	CreateUserId *string `json:"create_user_id,omitempty"`

	// **参数解释：**  更新者账号，用于指定文档的最后更新者信息。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o StructuredDocInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructuredDocInfo struct{}"
	}

	return strings.Join([]string{"StructuredDocInfo", string(data)}, " ")
}
