package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModel struct {
	Branch *VersionModelBranch `json:"branch,omitempty"`

	// **参数解释：**  检出时间，使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。 通常由系统自动维护，管理员更新时无需传入。如需强制置空（如解除检出状态），可通过needSetNullAttrs指定。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CheckOutTime *string `json:"checkOutTime,omitempty"`

	// **参数解释：**  检出用户名称，记录当前检出该实例的用户账号。 通常由系统自动维护，管理员更新时无需传入。如需强制置空（如解除检出状态），可通过needSetNullAttrs指定。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CheckOutUserName *string `json:"checkOutUserName,omitempty"`

	// **参数解释：**  实例创建时间，使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。 由系统自动生成，管理员更新时不可修改。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  创建者账号。由系统自动记录，管理员更新时不可修改。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  版本实例的描述信息，用于记录工艺变更说明、设备规格备注、BOM版本变更原因等场景描述。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：**  版本实例的唯一标识，用于定位待更新的M-V模型实例。 若指定的id不存在，系统将不做任何更新操作。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  关键信息资产ID，用于数据资产的安全标识与密级关联。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Kiaguid *string `json:"kiaguid,omitempty"`

	// **参数解释：**  实例最后更新时间，使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。由系统自动更新，管理员更新时无需传入。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	Master *VersionModelMaster `json:"master,omitempty"`

	// **参数解释：**  实例更新者账号。由系统自动记录为当前调用用户，管理员更新时无需传入。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  版本实例的中文名称，如“焊接工艺V2.1”、“数控机床规格A版”等。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：**  需要置为NULL值的属性名称列表，用于清除不再使用的属性值。支持基本属性、扩展属性和分类属性。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	NeedSetNullAttrs *[]string `json:"needSetNullAttrs,omitempty"`

	// **参数解释：**  扩展类型。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// **参数解释：**  安全密级。  **约束限制：**  不涉及。  **取值范围：**  - internal：内部公开。 - secret：秘密。 - confidential：机密。 - top_secret：绝密。  **默认取值：**  不涉及。
	SecurityLevel *string `json:"securityLevel,omitempty"`
}

func (o VersionModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModel struct{}"
	}

	return strings.Join([]string{"VersionModel", string(data)}, " ")
}
