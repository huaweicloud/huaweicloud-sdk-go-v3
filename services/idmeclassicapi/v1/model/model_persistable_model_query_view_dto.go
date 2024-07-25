package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistableModelQueryViewDto struct {

	// **参数解释：**  类名。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ClassName *string `json:"className,omitempty"`

	// **参数解释：**  创建时间。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  创建者。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  失效标识。   **取值范围：**  - true：失效。  - false：未失效。  **默认取值：**  false。
	DisableFlag *bool `json:"disableFlag,omitempty"`

	// **参数解释：**  扩展属性映射集。   **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ExtAttrMap *interface{} `json:"extAttrMap,omitempty"`

	// **参数解释：**  扩展属性列表。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ExtAttrs *[]ExaValueViewDto `json:"extAttrs,omitempty"`

	Folder *FolderQueryViewDto `json:"folder,omitempty"`

	// **参数解释：**  唯一标识。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  最后更新时间。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  修改人。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  扩展类型。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *TenantQueryViewDto `json:"tenant,omitempty"`
}

func (o PersistableModelQueryViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistableModelQueryViewDto struct{}"
	}

	return strings.Join([]string{"PersistableModelQueryViewDto", string(data)}, " ")
}
