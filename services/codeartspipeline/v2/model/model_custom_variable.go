package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomVariable struct {

	// **参数解释**： 流水线ID，可以通过[查询流水线列表](ListPipelines.xml)接口，其中pipelines.pipelineId即为流水线ID。 **约束限制**： 不涉及。 **取值范围**： 32位字符，由数字和字母组成。 **默认取值**： 不涉及。
	PipelineId *string `json:"pipeline_id,omitempty"`

	// **参数解释**： 自定义参数名称。 **约束限制**： 不涉及。 **取值范围**： 仅支持大小写英文字母、数字、“_”，不超过128个字符。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 参数序号，从1开始。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Sequence *int32 `json:"sequence,omitempty"`

	// **参数解释**： 自定义参数类型。 **约束限制**： 不涉及。 **取值范围**： - autoIncrement：自增长参数。 - enum：枚举参数。 - string：字符串参数。 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 自定义参数默认值。 **约束限制**： 不涉及。 **取值范围**： 最长8192字符。 **默认取值**： 不涉及。
	Value *string `json:"value,omitempty"`

	// **参数解释**： 是否私密参数。 **约束限制**： 不涉及。 **取值范围**： - true：是私密参数。 - false：不是私密参数。 **默认取值**： false。
	IsSecret *bool `json:"is_secret,omitempty"`

	// **参数解释**： 参数描述。 **约束限制**： 不涉及。 **取值范围**： 最长1024字符。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 是否运行时设置参数。 **约束限制**： 不涉及。 **取值范围**： - true：是运行时设置参数。 - false：不是运行时设置参数。 **默认取值**： false。
	IsRuntime *bool `json:"is_runtime,omitempty"`

	// **参数解释**： 枚举值列表。 **约束限制**： 不涉及。 **取值范围**： 每个枚举值不超过1024字符。 **默认取值**： 不涉及。
	Limits *[]interface{} `json:"limits,omitempty"`

	// **参数解释**： 是否重置。自增长参数被编辑，则使用编辑后的值，否则进行末位数字递增。 **约束限制**： 不涉及。 **取值范围**： - true：使用编辑后的参数值。 - false：使用自增长参数。 **默认取值**： false。
	IsReset *bool `json:"is_reset,omitempty"`

	// **参数解释**： 最近一次运行的参数值。 **约束限制**： 不涉及。 **取值范围**： 最长8192字符。 **默认取值**： 不涉及。
	LatestValue *string `json:"latest_value,omitempty"`

	// **参数解释**： 流水线运行时设置参数的传入值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RuntimeValue *string `json:"runtime_value,omitempty"`
}

func (o CustomVariable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomVariable struct{}"
	}

	return strings.Join([]string{"CustomVariable", string(data)}, " ")
}
