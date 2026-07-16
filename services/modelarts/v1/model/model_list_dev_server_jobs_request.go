package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDevServerJobsRequest Request Object
type ListDevServerJobsRequest struct {

	// **参数解释**：Lite Server job id。 **约束限制**：无。 **取值范围**：1 - 64字符。 **默认取值**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：Lite Server job的name。 **约束限制**：无。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：Lite Server job的类型。 **约束限制**：无。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**：Lite Server job的状态。 **约束限制**：无。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**：是否可见。 **约束限制**：bool。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Visible *bool `json:"visible,omitempty"`
}

func (o ListDevServerJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevServerJobsRequest struct{}"
	}

	return strings.Join([]string{"ListDevServerJobsRequest", string(data)}, " ")
}
