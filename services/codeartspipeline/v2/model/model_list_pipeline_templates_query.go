package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListPipelineTemplatesQuery struct {

	// **参数解释**： 模板语言。 **约束限制**： 不涉及。 **取值范围**： - java。 - python。 - nodejs。 - go。 - net。 - cpp。 - php。 - other。 - none。 **默认取值**： 不涉及。
	Language *string `json:"language,omitempty"`

	// **参数解释**： 是否系统模板。 **约束限制**： 不涉及。 **取值范围**： - true：是系统模板 - false：不是系统模板。 **默认取值**： 不涉及。
	IsSystem *bool `json:"is_system,omitempty"`

	// **参数解释**： 模板名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 偏移量，表示从此偏移量开始查询。 **约束限制**： 不涉及。 **取值范围**： offset大于等于0。 **默认取值**： 默认为0。
	Offset *int64 `json:"offset,omitempty"`

	// **参数解释**： 每次查询的条目数量。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 默认为10。
	Limit *int64 `json:"limit,omitempty"`

	// **参数解释**： 用于排序的字段，非必选。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**：   排序类型，非必选。 **约束限制**： 不涉及。 **取值范围**： - asc：按排序字段升序。 - desc：按排序字段降序。 **默认取值**： 不涉及。
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListPipelineTemplatesQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineTemplatesQuery struct{}"
	}

	return strings.Join([]string{"ListPipelineTemplatesQuery", string(data)}, " ")
}
