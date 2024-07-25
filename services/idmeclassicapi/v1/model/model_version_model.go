package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModel struct {
	Branch *VersionModelBranch `json:"branch,omitempty"`

	// **参数解释：**  检出时间。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CheckOutTime *string `json:"checkOutTime,omitempty"`

	// **参数解释：**  检出用户名称。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CheckOutUserName *string `json:"checkOutUserName,omitempty"`

	// **参数解释：**  创建时间。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  创建者。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  描述信息。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：**  唯一标识。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  关键信息资产ID。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Kiaguid *string `json:"kiaguid,omitempty"`

	// **参数解释：**  最后更新时间。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	Master *VersionModelMaster `json:"master,omitempty"`

	// **参数解释：**  更新者。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  中文名称。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：**  设置NULL值的属性名称。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	NeedSetNullAttrs *[]string `json:"needSetNullAttrs,omitempty"`

	// **参数解释：**  扩展类型。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// **参数解释：**  安全密级。  **约束限制：**  不涉及。  **取值范围：**  - INTERNAL：内部公开。 - SECRET：秘密。 - CONFIDENTIAL：机密。 - TOP_SECRET：绝密。  **默认取值：**  不涉及。
	SecurityLevel *string `json:"securityLevel,omitempty"`
}

func (o VersionModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModel struct{}"
	}

	return strings.Join([]string{"VersionModel", string(data)}, " ")
}
