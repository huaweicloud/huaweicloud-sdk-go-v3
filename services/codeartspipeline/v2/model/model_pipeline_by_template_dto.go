package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineByTemplateDto struct {

	// **参数解释**： 流水线名称。 **约束限制**： 不涉及。 **取值范围**： 仅包含中文、大小写英文字母、数字、'-'和'_'，且长度为[1,128]个字符。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 流水线描述。 **约束限制**： 不涉及。 **取值范围**： 不超过1024字符。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 是否为变更流水线。 **约束限制**： 不涉及。 **取值范围**： - true：是变更流水线。 - false：不是变更流水线。 **默认取值**： 不涉及。
	IsPublish bool `json:"is_publish"`

	// **参数解释**： 流水线源列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Sources []CodeSource `json:"sources"`

	// **参数解释**： 流水线涉密等级，非涉密场景不涉及，涉密场景必填。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SecurityLevel *int32 `json:"security_level,omitempty"`

	// **参数解释**： 流水线涉密等级编码，非涉密场景不涉及。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ConfidentialityCode *string `json:"confidentiality_code,omitempty"`

	// **参数解释**： 流水线参数列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Variables *[]PipelineByTemplateDtoVariables `json:"variables,omitempty"`
}

func (o PipelineByTemplateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineByTemplateDto struct{}"
	}

	return strings.Join([]string{"PipelineByTemplateDto", string(data)}, " ")
}
