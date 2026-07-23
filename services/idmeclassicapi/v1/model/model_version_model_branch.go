package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelBranch struct {

	// **参数解释：**  分支对象创建时间，使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。由系统自动生成，管理员更新时不可修改。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  分支对象的创建者账号。由系统自动记录，管理员更新时不可修改。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  分支对象的唯一标识，用于定位待更新的分支对象。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  分支对象最后更新时间，使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。由系统自动更新，管理员更新时无需传入。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  分支对象更新者账号。由系统自动记录，管理员更新时无需传入。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  需要将分支对象的自定义属性（包括基本属性、扩展属性和分类属性）设置为空值的属性名称列表。  **约束限制：**  不涉及。  **取值范围：**  长度不能超过1000个字符。  **默认取值：**  不涉及。
	NeedSetNullAttrs *[]string `json:"needSetNullAttrs,omitempty"`

	// **参数解释：**  分支对象扩展类型。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`
}

func (o VersionModelBranch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelBranch struct{}"
	}

	return strings.Join([]string{"VersionModelBranch", string(data)}, " ")
}
