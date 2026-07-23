package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistableModelQueryViewDto struct {

	// **参数解释：**  类名，表示数据实例的具体类类型。  **取值范围：**  不涉及。
	ClassName *string `json:"className,omitempty"`

	// **参数解释：**  创建时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  创建者账号。  **取值范围：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  失效标识，用于标识实例是否已失效。   **取值范围：**  - true：失效。  - false：未失效。
	DisableFlag *bool `json:"disableFlag,omitempty"`

	// **参数解释：**  扩展属性映射集，以键值对形式存储扩展属性。   **取值范围：**  不涉及。
	ExtAttrMap *interface{} `json:"extAttrMap,omitempty"`

	// **参数解释：**  扩展属性列表，包含实例的所有扩展属性详情。  **取值范围：**  不涉及。
	ExtAttrs *[]ExaValueViewDto `json:"extAttrs,omitempty"`

	Folder *FolderQueryViewDto `json:"folder,omitempty"`

	// **参数解释：**  数据实例的唯一标识。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  最后更新时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  更新者账号。  **取值范围：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  扩展类型，表示数据模型的具体扩展类型。  **取值范围：**  不涉及。
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
