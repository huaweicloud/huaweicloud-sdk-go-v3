package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistableModelListViewDto struct {

	// **参数解释：**  访问控制列表。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	AclEntry *string `json:"aclEntry,omitempty"`

	// **参数解释：**  类名。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ClassName *string `json:"className,omitempty"`

	// **参数解释：**  分类属性。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ClsAttrs *[]interface{} `json:"clsAttrs,omitempty"`

	// **参数解释：**  创建时间。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  创建者。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  失效标识。   **取值范围：**  - true：失效。  - false：未失效。  **默认取值：**  不涉及。
	DisableFlag *bool `json:"disableFlag,omitempty"`

	Folder *ObjectReferenceViewDto `json:"folder,omitempty"`

	// **参数解释：**  用于存储当前节点全路径。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	FullPath *string `json:"fullPath,omitempty"`

	// **参数解释：**  唯一标识。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  最后更新时间。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  是否为叶子节点。  **取值范围：**  - true：是叶子节点。 - false：不是叶子节点。  **默认取值：**  false。
	LeafFlag *bool `json:"leafFlag,omitempty"`

	LifecycleState *ObjectReferenceViewDto `json:"lifecycleState,omitempty"`

	LifecycleTemplate *ObjectReferenceViewDto `json:"lifecycleTemplate,omitempty"`

	// **参数解释：**  更新者。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  拥有者。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Owner *string `json:"owner,omitempty"`

	ParentNode *ObjectReferenceViewDto `json:"parentNode,omitempty"`

	// **参数解释：**  用于存储当前节点原始全路径。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RawFullPath *string `json:"rawFullPath,omitempty"`

	// **参数解释：**  软删除标识。  **取值范围：**  - 0：表示未删除。 - 1：表示已删除。  **默认取值：**  0。
	RdmDeleteFlag *int32 `json:"rdmDeleteFlag,omitempty"`

	// **参数解释：**  扩展类型。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// **参数解释：**  系统版本。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	RootNode *ObjectReferenceViewDto `json:"rootNode,omitempty"`

	Tenant *ObjectReferenceViewDto `json:"tenant,omitempty"`

	// **参数解释：**  示例模型的唯一键约束属性。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	UniqueKey *string `json:"uniqueKey,omitempty"`
}

func (o PersistableModelListViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistableModelListViewDto struct{}"
	}

	return strings.Join([]string{"PersistableModelListViewDto", string(data)}, " ")
}
