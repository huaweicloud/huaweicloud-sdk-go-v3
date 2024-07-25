package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryCondition struct {

	// **参数解释：**  查询条件的名称（数据模型的属性英文名称）。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ConditionName *string `json:"conditionName,omitempty"`

	// **参数解释：**  查询条件值（已过时）。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ConditionValue *string `json:"conditionValue,omitempty"`

	// **参数解释：**  查询条件的值。operator为in时有多值，operator为其他操作符时均为单值。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ConditionValues *[]string `json:"conditionValues,omitempty"`

	// **参数解释：**  查询条件列表。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Conditions *[]QueryCondition `json:"conditions,omitempty"`

	// **参数解释：**  是否忽略大小写。  **约束限制：**  不涉及。  **取值范围：**  - true：表示忽略。 - false：表示不忽略。  **默认取值：**  false。
	IgnoreStr *bool `json:"ignoreStr,omitempty"`

	// **参数解释：**  关联查询时被关联表的别名。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	JoinTableAlias *string `json:"joinTableAlias,omitempty"`

	// **参数解释：**  连接符。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Joiner *string `json:"joiner,omitempty"`

	// **参数解释：**  操作符。  **约束限制：**  不涉及。  **取值范围：**  - =：等于查询。 - like：模糊查询。 - customLike：支持输入*或%的模糊查询。 - in：包含查询。 - <：小于查询。 - \\>：大于查询。 - \\>=：大于等于查询。 - <=：小于等于查询。 - <>：不等于查询。  **默认取值：**  不涉及。
	Operator *string `json:"operator,omitempty"`

	PreCondition *QueryCondition `json:"preCondition,omitempty"`
}

func (o QueryCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCondition struct{}"
	}

	return strings.Join([]string{"QueryCondition", string(data)}, " ")
}
