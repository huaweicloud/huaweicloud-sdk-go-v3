package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExaValueViewDto struct {

	// **参数解释：**  类名。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ClassName *string `json:"className,omitempty"`

	// **参数解释：**  创建时间。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  创建者。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  唯一标识。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  最后更新时间。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  修改人。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  中文名称。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：**  软删除标识。  **取值范围：**  - 0：表示未删除。 - 1：表示已删除。  **默认取值：**  0。
	RdmDeleteFlag *int32 `json:"rdmDeleteFlag,omitempty"`

	// **参数解释：**  扩展类型。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// **参数解释：**  系统版本。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	Tenant *TenantViewDto `json:"tenant,omitempty"`

	// **参数解释：**  扩展属性类型。  **约束限制：**  不涉及。  **取值范围：**  - STRING：文本 - TEXT：长文本 - INTEGER：整型 - LONG：长整型 - DECIMAL：浮点型 - DECIMAL_WITH_PRECISION：浮点型（自定义精度） - FILE：文件 - BOOLEAN：布尔值 - DATE：日期 - ENUM：枚举 - CATEGORY：分类 - URL：URL - REFERENCE_OBJECT：参考对象  **默认取值：**  不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释：**  值。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Value *interface{} `json:"value,omitempty"`
}

func (o ExaValueViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExaValueViewDto struct{}"
	}

	return strings.Join([]string{"ExaValueViewDto", string(data)}, " ")
}
