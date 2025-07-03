package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryCondition 查询条件构建对象。
type QueryCondition struct {

	// **参数解释：**  用于定义一组嵌套查询条件，其逻辑关系由同级的joiner属性决定。每个子条件可包含基础条件或嵌套条件，形成多层级查询逻辑。  - 基础条件结构：由conditionName（属性英文名称）、operator（操作符）、conditionValues（条件值）组成单一查询条件。 - 嵌套条件结构：包含子级joiner和conditions数组，支持多层级逻辑嵌套。  **约束限制：**  当一个对象同时符合基础条件和嵌套条件时，系统将优先考虑嵌套条件，并忽略基础条件。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Conditions *[]QueryCondition `json:"conditions,omitempty"`

	// **参数解释：**  表示查询条件的名称（对应数据模型中属性的英文名称）。  **约束限制：**  该参数仅在conditions数组中生效，在filter中不生效。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ConditionName *string `json:"conditionName,omitempty"`

	// **参数解释：**  操作符。  **约束限制：**  - 仅在conditions数组中生效，在filter中不生效。 - “多值”为“是”的属性支持使用contains、=、like、customLike、startWith、endWith、ISNULL。 - “多值”为“否”的属性不支持使用contains。   **取值范围：**  - =：等于查询。 - <：小于查询。 - \\>：大于查询。 - \\>=：大于等于查询。 - <=：小于等于查询。 - <>：不等于查询。 - startWith：头匹配查询。 - endWith：尾匹配查询。 - like：模糊查询。 - customLike：支持输入*或%的模糊查询。 - in：包含查询。 - not in：排除查询。 - ISNULL：值为空查询。 - NOTNULL：值不为空查询。 - contains：包含查询。  **默认取值：**  不涉及。
	Operator *string `json:"operator,omitempty"`

	// **参数解释：**  查询条件的值。  **约束限制：**  仅在conditions数组中生效，在filter中不生效。  **取值范围：**  - 非多值约束属性：operator为in时有多个值，operator为其他操作符时均为单个值。 - 多值约束属性：operator为contains或=时，取值为多个值。operator为其他操作符时均为单个值。  **默认取值：**  不涉及。
	ConditionValues *[]string `json:"conditionValues,omitempty"`

	// **参数解释：**  查询条件值（已过时）。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ConditionValue *string `json:"conditionValue,omitempty"`

	// **参数解释：**  是否忽略大小写。  **约束限制：**  仅在conditions数组中生效，在filter中不生效。  **取值范围：**  - true：表示忽略。 - false：表示不忽略。  **默认取值：**  false。
	IgnoreStr *bool `json:"ignoreStr,omitempty"`

	// **参数解释：**  关联查询时被关联表的别名。  **约束限制：**  仅在conditions数组中生效，在filter中不生效。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	JoinTableAlias *string `json:"joinTableAlias,omitempty"`

	// **参数解释：**  连接符，连接conditions中的条件对象。  **约束限制：**  不涉及。  **取值范围：**  - and：与连接。 - or：或连接。  **默认取值：**  不涉及。
	Joiner *string `json:"joiner,omitempty"`

	PreCondition *QueryCondition `json:"preCondition,omitempty"`
}

func (o QueryCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCondition struct{}"
	}

	return strings.Join([]string{"QueryCondition", string(data)}, " ")
}
