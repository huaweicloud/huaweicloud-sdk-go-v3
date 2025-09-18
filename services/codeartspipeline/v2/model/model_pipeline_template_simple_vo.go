package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineTemplateSimpleVo struct {

	// **参数解释**： 模板ID。 **取值范围**： 32位字符，由数字和字母组成。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 模板名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 模板图标。 **取值范围**： 不涉及。
	Icon *string `json:"icon,omitempty"`

	// **参数解释**： 版本。 **取值范围**： 默认3.0。
	ManifestVersion *string `json:"manifest_version,omitempty"`

	// **参数解释**： 模板语言。 **取值范围**： - java。 - python。 - nodejs。 - go。 - net。 - cpp。 - php。 - other。 - none。
	Language *string `json:"language,omitempty"`

	// **参数解释**： 模板描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 是否系统模板。 **取值范围**： - true：是系统模板。 - false：不是系统模板。
	IsSystem *bool `json:"is_system,omitempty"`

	// **参数解释**： 模板局点。 **取值范围**： 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释**： 模板所属租户ID。 **取值范围**： 32位字符，由数字和字母组成。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**： 模板创建人ID。 **取值范围**： 32位字符，由数字和字母组成。
	CreatorId *string `json:"creator_id,omitempty"`

	// **参数解释**： 模板创建人名称。 **取值范围**： 不涉及。
	CreatorName *string `json:"creator_name,omitempty"`

	// **参数解释**： 模板更新人ID。 **取值范围**： 32位字符，由数字和字母组成。
	UpdaterId *string `json:"updater_id,omitempty"`

	// **参数解释**： 是否收藏。 **取值范围**： - true：收藏。 - false：不收藏。
	IsCollect *bool `json:"is_collect,omitempty"`

	// **参数解释**： 是否展示流水线源。 **取值范围**： - true：展示流水线源。 - false：不展示流水线源。
	IsShowSource *string `json:"is_show_source,omitempty"`

	// **参数解释**： 模板编排的阶段列表。 **约束限制**： 不涉及。
	Stages *[]PipelineTemplateSimpleVoStages `json:"stages,omitempty"`
}

func (o PipelineTemplateSimpleVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineTemplateSimpleVo struct{}"
	}

	return strings.Join([]string{"PipelineTemplateSimpleVo", string(data)}, " ")
}
