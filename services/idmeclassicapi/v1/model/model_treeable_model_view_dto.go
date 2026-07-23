package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TreeableModelViewDto struct {

	// **参数解释：**  唯一标识，系统生成的子节点主键ID。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  创建者账号，标识创建该子节点的用户。  **取值范围：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  更新者账号，标识最后更新该子节点的用户。  **取值范围：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  创建时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  最新更新时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
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

	RootNode *ObjectReferenceViewDto `json:"rootNode,omitempty"`

	ParentNode *ObjectReferenceViewDto `json:"parentNode,omitempty"`

	// **参数解释：**  是否为叶子节点，标识该节点是否拥有子节点。  **取值范围：**  - true：是叶子节点（无子节点）。 - false：不是叶子节点（有子节点）。  **默认取值：**  不涉及。
	LeafFlag *bool `json:"leafFlag,omitempty"`

	// **参数解释：**  用于存储当前节点全路径，格式为从根节点到当前节点的ID路径，以“/”分隔。 例如：根节点ID/父节点ID/当前节点ID/。  **取值范围：**  不涉及。
	FullPath *string `json:"fullPath,omitempty"`

	// **参数解释：**  用于存储当前节点原始全路径，格式与fullPath一致，记录节点在树形结构中的原始层级位置。  **取值范围：**  不涉及。
	RawFullPath *string `json:"rawFullPath,omitempty"`
}

func (o TreeableModelViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TreeableModelViewDto struct{}"
	}

	return strings.Join([]string{"TreeableModelViewDto", string(data)}, " ")
}
