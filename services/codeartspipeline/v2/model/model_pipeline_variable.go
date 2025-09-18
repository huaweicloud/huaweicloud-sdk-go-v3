package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineVariable 流水线自定义参数
type PipelineVariable struct {

	// **参数解释**： 自定义参数名称。 **取值范围**： 仅支持大小写英文字母、数字、“_”，不超过128个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 序号。 **取值范围**： [1, 2147483647]。
	Sequence *int32 `json:"sequence,omitempty"`

	// **参数解释**： 自定义参数的类型。 **取值范围**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 自定义参数的默认值。 **取值范围**： 不涉及。
	Value *string `json:"value,omitempty"`

	// **参数解释**： 是否为私密参数。 **取值范围**： - true：是私密参数。 - false：不是私密参数。
	IsSecret *bool `json:"is_secret,omitempty"`

	// **参数解释**： 自定义参数描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 是否运行时设置。 **取值范围**： - true：运行时设置。 - false：非运行时设置。
	IsRuntime *bool `json:"is_runtime,omitempty"`

	// **参数解释**： 限定枚举值列表。 **取值范围**： 不涉及。
	Limits *[]string `json:"limits,omitempty"`

	// **参数解释**： 自增长参数是否被重置。 **取值范围**： - true：被重置。 - false：不被重置。
	IsReset *bool `json:"is_reset,omitempty"`

	// **参数解释**： 自增长参数最新值。 **取值范围**： 不涉及。
	LatestValue *string `json:"latest_value,omitempty"`
}

func (o PipelineVariable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineVariable struct{}"
	}

	return strings.Join([]string{"PipelineVariable", string(data)}, " ")
}
