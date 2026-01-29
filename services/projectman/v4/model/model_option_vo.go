package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OptionVo struct {

	// **参数解释：**  选项Id。 **取值范围：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  选项名称。 **取值范围：**  不涉及。
	DisplayValue *string `json:"display_value,omitempty"`

	// **参数解释：**  选项的唯一标识，自定义的选项id和value相同。 **取值范围：**  不涉及。
	Value *string `json:"value,omitempty"`

	// **参数解释：**  选项code。 **取值范围：**  不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释：**  选项名称的中文拼音首字母。 **取值范围：**  不涉及。
	ValuePy *string `json:"value_py,omitempty"`

	// **参数解释：**  选项在选项列表中的排序。 **取值范围：**  不涉及。
	Sequence *int32 `json:"sequence,omitempty"`

	// **参数解释：**  层级选项的分层标识，单选列表和多选列表值都为1，层级字段按照层级依次为1,2,3,4。 **取值范围：**  不涉及。
	Level *int32 `json:"level,omitempty"`

	// **参数解释：**  选项所在的项目空间Id。 **取值范围：**  不涉及。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释：**  选项归属的定义级别。1,2,3为系统级，4为租户自定义，5为项目自定义。 **取值范围：**  不涉及。
	BelongDefinitionType *string `json:"belong_definition_type,omitempty"`
}

func (o OptionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OptionVo struct{}"
	}

	return strings.Join([]string{"OptionVo", string(data)}, " ")
}
