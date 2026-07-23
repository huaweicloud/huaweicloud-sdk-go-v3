package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StructuredDocViewDto struct {

	// **参数解释：**  创建者账号，标识创建该文档的用户。  **取值范围：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  创建时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  最后更新时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  系统版本号，用于数据版本控制。  **取值范围：**  不涉及。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	// **参数解释：**  软删除标识。  **取值范围：**  - 0：表示未删除。 - 1：表示已删除。
	RdmDeleteFlag *int32 `json:"rdmDeleteFlag,omitempty"`

	// **参数解释：**  扩展类型，标识对象的扩展类别。  **取值范围：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *TenantViewDto `json:"tenant,omitempty"`

	// **参数解释：**  类名，标识对象的Java类名称。  **取值范围：**  不涉及。
	ClassName *string `json:"className,omitempty"`

	// **参数解释：**  类名，标识对象的类类型。  **取值范围：**  不涉及。
	Clazz *string `json:"clazz,omitempty"`

	// **参数解释：**  唯一标识，系统生成的文档主键ID。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  KooPage文档ID，系统生成的KooPage文档唯一标识。  **取值范围：**  不涉及。
	DocumentId *string `json:"documentId,omitempty"`

	// **参数解释：**  文档标题。  **取值范围：**  不涉及。
	Title *string `json:"title,omitempty"`

	// **参数解释：**  文档类型。  **取值范围：**  - directory：目录。 - pageDocument：Page文档。 - boardDocument：Board文档。 - mindDocument：Mind文档。 - drawDocument：Draw文档。
	Type string `json:"type"`

	// **参数解释：**  模板ID。  **取值范围：**  不涉及。
	TemplateId *string `json:"templateId,omitempty"`

	// **参数解释：**  团队ID。  **取值范围：**  不涉及。
	OrganizationId *string `json:"organizationId,omitempty"`

	// **参数解释：**  知识库ID。  **取值范围：**  不涉及。
	WikiId *string `json:"wikiId,omitempty"`

	// **参数解释：**  父文档ID，标识该文档所属的父级目录。  **取值范围：**  不涉及。
	ParentDocumentId *string `json:"parentDocumentId,omitempty"`

	Instance *ObjectReferenceViewDto `json:"instance,omitempty"`

	// **参数解释：**  创建者ID，标识创建该文档的用户ID。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CreateUserId *string `json:"createUserId,omitempty"`

	// **参数解释：**  更新者账号，标识最后更新该文档的用户。  **取值范围：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o StructuredDocViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructuredDocViewDto struct{}"
	}

	return strings.Join([]string{"StructuredDocViewDto", string(data)}, " ")
}
