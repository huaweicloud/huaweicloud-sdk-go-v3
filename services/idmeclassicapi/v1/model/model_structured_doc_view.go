package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StructuredDocView struct {

	// **参数解释：**  创建者。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  创建时间。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  最后更新时间。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  系统版本。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	// **参数解释：**  软删除标识，参数值为0或1。  **取值范围：**  - 0：表示未删除。 - 1：表示已删除。  **默认取值：**  不涉及。
	RdmDeleteFlag *int32 `json:"rdmDeleteFlag,omitempty"`

	// **参数解释：**  扩展类型。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *TenantViewDto `json:"tenant,omitempty"`

	// **参数解释：**  类名。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ClassName *string `json:"className,omitempty"`

	// **参数解释：**  类名。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Clazz *string `json:"clazz,omitempty"`

	// **参数解释：**  唯一标识。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  kooPage文档ID。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	DocumentId *string `json:"document_id,omitempty"`

	// **参数解释：**  文档标题。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Title *string `json:"title,omitempty"`

	// **参数解释：**  文档类型。  **取值范围：**  - directory：目录。 - pageDocument：Page文档。 - boardDocument：Board文档。 - mindDocument：Mind文档。 - drawDocument：Draw文档。  **默认取值：**  不涉及。
	Type string `json:"type"`

	// **参数解释：**  模板ID。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	TemplateId *string `json:"template_id,omitempty"`

	// **参数解释：**  团队ID。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	OrganizationId *string `json:"organization_id,omitempty"`

	// **参数解释：**  知识库ID。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	WikiId *string `json:"wiki_id,omitempty"`

	// **参数解释：**  父文档ID。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ParentDocumentId *string `json:"parent_document_id,omitempty"`

	// **参数解释：**  实例ID。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释：**  模型名称。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ModelName *string `json:"model_name,omitempty"`

	// **参数解释：**  创建者ID。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CreateUserId *string `json:"create_user_id,omitempty"`

	// **参数解释：**  修改人。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o StructuredDocView) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructuredDocView struct{}"
	}

	return strings.Join([]string{"StructuredDocView", string(data)}, " ")
}
