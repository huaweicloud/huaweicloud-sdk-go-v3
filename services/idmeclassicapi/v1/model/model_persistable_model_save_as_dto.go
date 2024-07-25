package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistableModelSaveAsDto struct {

	// **参数解释：**  唯一标识。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  修改者。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  最后更新时间。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  创建者。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  创建时间。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  扩展类型。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *ObjectReferenceParamDto `json:"tenant,omitempty"`

	// **参数解释：**  源模型编号。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	SourceEntityNumber *string `json:"sourceEntityNumber,omitempty"`

	// **参数解释：**  源实例的唯一标识（单实例为ID，版本实例为versionId）。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	SourceInstanceId string `json:"sourceInstanceId"`

	// **参数解释：**  置空字段数组。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	NeedSetNullAttrs *[]string `json:"needSetNullAttrs,omitempty"`

	// **参数解释：**  要保存的属性。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	EntityToSave *interface{} `json:"entityToSave,omitempty"`

	// **参数解释：**  要保存的结果。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	EntityToReturn *interface{} `json:"entityToReturn,omitempty"`

	// **参数解释：**  唯一键约束属性。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	UniqueKey *string `json:"uniqueKey,omitempty"`
}

func (o PersistableModelSaveAsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistableModelSaveAsDto struct{}"
	}

	return strings.Join([]string{"PersistableModelSaveAsDto", string(data)}, " ")
}
