package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HierarchicalValue **参数解释**： 多层级告警阈值。    **约束限制**： 当hierarchical_value和value同时使用时以hierarchical_value为准。 创建或修改告警规则以下2种场景只支持设置一个阈值：   1.告警类型为`指标告警`且告警策略为`所有策略都满足才告警`的场景。   2.告警类型为`事件告警`的场景。
type HierarchicalValue struct {

	// **参数解释**： 紧急级别的阈值。 **约束限制**： 不涉及。 **取值范围**： [-1.7976931348623157e+108, 1.7976931348623157e+108]           **默认取值**： 不涉及。
	Critical *float64 `json:"critical,omitempty"`

	// **参数解释**： 重要级别的阈值。 **约束限制**： 不涉及。 **取值范围**： [-1.7976931348623157e+108, 1.7976931348623157e+108]           **默认取值**： 不涉及。
	Major *float64 `json:"major,omitempty"`

	// **参数解释**： 次要级别的阈值。 **约束限制**： 不涉及。 **取值范围**： [-1.7976931348623157e+108, 1.7976931348623157e+108]           **默认取值**： 不涉及。
	Minor *float64 `json:"minor,omitempty"`

	// **参数解释**： 提示级别的阈值。 **约束限制**： 不涉及。 **取值范围**： [-1.7976931348623157e+108, 1.7976931348623157e+108]           **默认取值**： 不涉及。
	Info *float64 `json:"info,omitempty"`
}

func (o HierarchicalValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HierarchicalValue struct{}"
	}

	return strings.Join([]string{"HierarchicalValue", string(data)}, " ")
}
