package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareConditionOption 组合比较查询条件，可针对某个给定过滤字段，进行区间范围、大小、小于等条件合并查询。
type CompareConditionOption struct {

	// **参数解释**: 查询字段名称，当前仅支持特定的数值字段。 **约束限制**: 不涉及。 **取值范围**: - total_sql_time: 总SQL耗时。 - sql_time：SQL执行次数。  **默认取值**: 不涉及。
	Name string `json:"name"`

	// **参数解释**: 是否使能包含等于，如果为true，则表示包含边界条件（min或max）的取值。 **约束限制**: 不涉及。 **取值范围**: - true - false  **默认取值**: 不涉及。
	EnableEqual *bool `json:"enable_equal,omitempty"`

	// **参数解释**: 最小值判断条件对应取值（大于条件）。 **约束限制**: 不涉及。 **取值范围**: [0, 2^63-1] **默认取值**: 不涉及。
	Min *int64 `json:"min,omitempty"`

	// **参数解释**: 最大值判断条件对应取值（小于条件）。 **约束限制**: 不涉及。 **取值范围**: [0, 2^63-1] **默认取值**: 不涉及。
	Max *int64 `json:"max,omitempty"`

	// **参数解释**: 等值判断条件对应取值（等于条件）。value的优先级最高，如果value不为空，则忽略min和max的取值设置；value为空时，才使能min和max的条件过滤。 **约束限制**: 不涉及。 **取值范围**: [0, 2^63-1] **默认取值**: 不涉及。
	Value *int64 `json:"value,omitempty"`
}

func (o CompareConditionOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareConditionOption struct{}"
	}

	return strings.Join([]string{"CompareConditionOption", string(data)}, " ")
}
