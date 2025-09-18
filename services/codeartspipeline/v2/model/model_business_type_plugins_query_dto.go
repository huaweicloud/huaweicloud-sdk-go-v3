package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BusinessTypePluginsQueryDto struct {

	// **参数解释**： 用于区分插件为流水线可使用/模板可使用。 **约束限制**： 不涉及。 **取值范围**： - pipeline：流水线可使用。 - template：模板可使用。 **默认取值**： 不涉及。
	UseCondition *string `json:"use_condition,omitempty"`

	// **参数解释**： 用于区分源的代码仓类型codehub/gitlab/github等。 **约束限制**： 不涉及。 **取值范围**： - codehub。 - gitee。 - github。 - gitcode。 - gitlab。 **默认取值**： 不涉及。
	InputRepoType *string `json:"input_repo_type,omitempty"`

	// **参数解释**： 用于区分单源/多源的情况。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InputSourceType *string `json:"input_source_type,omitempty"`

	// **参数解释**： 业务类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BusinessType *string `json:"business_type,omitempty"`

	// **参数解释**： 匹配名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RegexName *string `json:"regex_name,omitempty"`
}

func (o BusinessTypePluginsQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessTypePluginsQueryDto struct{}"
	}

	return strings.Join([]string{"BusinessTypePluginsQueryDto", string(data)}, " ")
}
