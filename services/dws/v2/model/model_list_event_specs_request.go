package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventSpecsRequest Request Object
type ListEventSpecsRequest struct {

	// **参数解释**： 事件配置名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SpecName *string `json:"spec_name,omitempty"`

	// **参数解释**： 事件类别。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Category *string `json:"category,omitempty"`

	// **参数解释**： 事件级别。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Severity *string `json:"severity,omitempty"`

	// **参数解释**： 事件源类别。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SourceType *string `json:"source_type,omitempty"`

	// **参数解释**： 事件标签。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Tag *string `json:"tag,omitempty"`

	// **参数解释**： 分页偏移量，从0开始，页数减1。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 0
	Offset *string `json:"offset,omitempty"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 大于0。 **默认取值**： 1000
	Limit *string `json:"limit,omitempty"`
}

func (o ListEventSpecsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventSpecsRequest struct{}"
	}

	return strings.Join([]string{"ListEventSpecsRequest", string(data)}, " ")
}
