package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineTemplateDto struct {

	// **参数解释**： 模板名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 模板描述。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 模板语言。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Language string `json:"language"`

	// **参数解释**： 自定义参数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Variables *[]CustomVariable `json:"variables,omitempty"`

	// **参数解释**： 模板编排json，包含stages。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Definition string `json:"definition"`

	// **参数解释**： 是否系统模板。 **约束限制**： 不涉及。 **取值范围**： - true：是系统模板。 - false：不是系统模板。 **默认取值**： 不涉及。
	IsSystem bool `json:"is_system"`

	// **参数解释**： 租户id。 **约束限制**： 不涉及。 **取值范围**： 32位字符，由数字和字母组成。 **默认取值**： 不涉及。
	DomainId string `json:"domain_id"`

	// **参数解释**： 是否显示流水线源。 **约束限制**： 不涉及。 **取值范围**： - true：显示流水线源。 - false：不显示流水线源。 **默认取值**： 不涉及。
	IsShowSource bool `json:"is_show_source"`
}

func (o PipelineTemplateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineTemplateDto struct{}"
	}

	return strings.Join([]string{"PipelineTemplateDto", string(data)}, " ")
}
