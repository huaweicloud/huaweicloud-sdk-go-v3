package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryRequestStaticsVo struct {

	// **参数解释：**  SQL编码类型，用于指定数据统计查询的字符编码格式。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  UTF8。
	CharacterSet *string `json:"characterSet,omitempty"`

	// **参数解释：**  查询条件组。  此参数已废弃，不建议继续使用，建议使用替代参数filter。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Conditions *[]QueryCondition `json:"conditions,omitempty"`

	// **参数解释：**  是否对查询过程中的敏感数据进行加密。  **约束限制：**  不涉及。  **取值范围：**  - true：加密。 - false：不加密。  **默认取值：**  false。
	Decrypt *bool `json:"decrypt,omitempty"`

	// **参数解释：**  实体类型，用于区分数据实体与关系实体。  **约束限制：**  不涉及。  **取值范围：**  - ENTITY：数据实体。 - RRELATION：关系实体。  **默认取值：**  不涉及。
	EntityType *string `json:"entityType,omitempty"`

	Filter *QueryCondition `json:"filter,omitempty"`

	// **参数解释：**  指定聚合函数信息数组，用于定义对数据执行的统计计算方式。支持同时配置多个函数，如同时计算平均值和最大值。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Funcs []FuncInfo `json:"funcs"`

	// **参数解释：**  分组属性名称，用于指定按哪个数据属性进行分组统计。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	GroupBy *string `json:"groupBy,omitempty"`

	// **参数解释：**  是否需要查询总记录数。  **约束限制：**  不涉及。  **取值范围：**  - true：需要。 - false：不需要。  **默认取值：**  false。
	IsNeedTotal *bool `json:"isNeedTotal,omitempty"`

	// **参数解释：**  是否需要展示所有参考对象信息。  **约束限制：**  不涉及。  **取值范围：**  - true：需要。 - false：不需要。  **默认取值：**  不涉及。
	IsPresentAll *bool `json:"isPresentAll,omitempty"`

	// **参数解释：**  需要展示详细信息的参考对象名称列表。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	NeedPresentDetail *[]string `json:"needPresentDetail,omitempty"`

	// **参数解释：**  按某个字段进行排序。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	OrderBy *string `json:"orderBy,omitempty"`

	// **参数解释：**  排序字段的表别名。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	OrderByTableAlias *string `json:"orderByTableAlias,omitempty"`

	// **参数解释：**  多租查询参数，用于控制公共数据（如标准工艺库、通用物料库）的统计范围。  **约束限制：**  不涉及。  **取值范围：**  - EXCLUDE_PUBLIC_DATA：仅统计当前租户私有数据，不包括公共数据。 - INCLUDE_PUBLIC_DATA：同时统计当前租户私有数据与公共数据。 - ONLY_NEED_PUBLIC_DATA：仅统计公共数据。  **默认取值：**  不涉及。
	PublicData *string `json:"publicData,omitempty"`

	// **参数解释：**  排序方向。  **约束限制：**  不涉及。  **取值范围：**  - ASC：表示升序。 - DESC：表示降序。  **默认取值：**  ASC。
	Sort *string `json:"sort,omitempty"`

	// **参数解释：**  排序。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Sorts *[]SortInfoVo `json:"sorts,omitempty"`
}

func (o QueryRequestStaticsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRequestStaticsVo struct{}"
	}

	return strings.Join([]string{"QueryRequestStaticsVo", string(data)}, " ")
}
