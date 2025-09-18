package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineByTemplateDtoVariables struct {

	// **参数解释**： 参数名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *interface{} `json:"name,omitempty"`

	// **参数解释**： 参数序号，从1开始递增。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Sequence *int32 `json:"sequence,omitempty"`

	// **参数解释**： 参数类型。 **约束限制**： 不涉及。 **取值范围**： - autoIncrement：自增长参数。 - enum：枚举参数。 - string：字符串参数。 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 参数值。 **约束限制**： 不涉及。 **取值范围**： 不超过8192字符。 **默认取值**： 不涉及。
	Value *string `json:"value,omitempty"`

	// **参数解释**： 是否私密参数。 **约束限制**： 不涉及。 **取值范围**： - true：是私密参数。 - false：不是私密参数。 **默认取值**： false。
	IsSecret *bool `json:"is_secret,omitempty"`

	// **参数解释**： 参数描述。 **约束限制**： 不涉及。 **取值范围**： 不超过512字符。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 是否运行时设置参数。 **约束限制**： 不涉及。 **取值范围**： - true：是运行时设置参数。 - false：不是运行时设置参数。 **默认取值**： false。
	IsRuntime *bool `json:"is_runtime,omitempty"`

	// **参数解释**： 是否重置。自增长参数被编辑，则使用编辑后的值，否则进行末位数字递增。 **约束限制**： 不涉及。 **取值范围**： - true：使用编辑后的参数值。 - false：使用自增长参数。 **默认取值**： false
	IsReset *bool `json:"is_reset,omitempty"`

	// **参数解释**： 最近一次运行的参数值。 **约束限制**： 不涉及。 **取值范围**： 不超过8192字符。 **默认取值**： 不涉及。
	LatestValue *string `json:"latest_value,omitempty"`

	// **参数解释**： 枚举值列表。 **约束限制**： 不涉及。 **取值范围**： 每个枚举值不超过1024字符。 **默认取值**： 不涉及。
	Limits *[]string `json:"limits,omitempty"`
}

func (o PipelineByTemplateDtoVariables) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineByTemplateDtoVariables struct{}"
	}

	return strings.Join([]string{"PipelineByTemplateDtoVariables", string(data)}, " ")
}
