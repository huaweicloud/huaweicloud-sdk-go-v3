package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FolderQueryViewDto struct {

	// **参数解释：**  编码。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	BusinessCode string `json:"businessCode"`

	// **参数解释：**  类名。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ClassName *string `json:"className,omitempty"`

	// **参数解释：**  创建时间。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  创建者。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  中文描述。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：**  英文描述。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	DescriptionEn *string `json:"descriptionEn,omitempty"`

	// **参数解释：**  失效标识。  **取值范围：**  - true：失效。  - false：未失效。  **默认取值：**  false。
	DisableFlag *bool `json:"disableFlag,omitempty"`

	// **参数解释：**  扩展属性映射集。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ExtAttrMap *interface{} `json:"extAttrMap,omitempty"`

	// **参数解释：**  扩展属性列表。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ExtAttrs *[]ExaValueViewDto `json:"extAttrs,omitempty"`

	// **参数解释：**  唯一标识。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  最后更新时间。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  修改人。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  中文名称。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Name string `json:"name"`

	// **参数解释：**  英文名称。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	NameEn *string `json:"nameEn,omitempty"`

	// **参数解释：**  扩展类型。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *TenantQueryViewDto `json:"tenant,omitempty"`

	// **参数解释：**  类别。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Type *string `json:"type,omitempty"`
}

func (o FolderQueryViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FolderQueryViewDto struct{}"
	}

	return strings.Join([]string{"FolderQueryViewDto", string(data)}, " ")
}
