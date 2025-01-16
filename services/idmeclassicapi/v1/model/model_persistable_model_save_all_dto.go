package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistableModelSaveAllDto struct {

	// **参数解释：**  创建时间。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  创建者。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Creator *int64 `json:"creator,omitempty"`

	// **参数解释：**  唯一标识。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  最后更新时间。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  更新者。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  将自定义属性（包括基本属性、扩展属性和分类属性）设置为空值，其长度不能超过1000个字符。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	NeedSetNullAttrs *[]string `json:"needSetNullAttrs,omitempty"`

	// **参数解释：**  扩展类型。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *ObjectReferenceParamDto `json:"tenant,omitempty"`

	// **参数解释：**  示例模型的唯一键属性。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	UniqueKey *string `json:"uniqueKey,omitempty"`
}

func (o PersistableModelSaveAllDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistableModelSaveAllDto struct{}"
	}

	return strings.Join([]string{"PersistableModelSaveAllDto", string(data)}, " ")
}
