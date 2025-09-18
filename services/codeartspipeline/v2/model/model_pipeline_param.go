package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineParam **参数解释**： 流水线参数 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type PipelineParam struct {

	// **参数解释**： 流水线参数名字 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 流水线参数值 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Value string `json:"value"`

	// **参数解释**： 流水线参数描述 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Description string `json:"description"`

	// **参数解释**： 流水线参数类型 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ParamType string `json:"param_type"`

	// **参数解释**： 是否静态参数。 **约束限制**： 不涉及。 **取值范围**： - true：是静态参数。 - false：不是静态参数。 **默认取值**： 不涉及。
	IsStatic bool `json:"is_static"`

	// **参数解释**： 是否默认参数。 **约束限制**： 不涉及。 **取值范围**： - true：是默认参数。 - false：不是默认参数。 **默认取值**： 不涉及。
	IsDefault bool `json:"is_default"`
}

func (o PipelineParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineParam struct{}"
	}

	return strings.Join([]string{"PipelineParam", string(data)}, " ")
}
