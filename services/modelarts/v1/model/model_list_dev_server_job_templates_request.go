package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDevServerJobTemplatesRequest Request Object
type ListDevServerJobTemplatesRequest struct {

	// **参数解释**：DevServerJob的模板id。 **约束限制**：1 - 64字符，字母、数字和中划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：DevServerJob的模板name。 **约束限制**：字母、数字和中划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：DevServerJob的模板类型。 **约束限制**：字母、数字和中划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Type *string `json:"type,omitempty"`
}

func (o ListDevServerJobTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevServerJobTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListDevServerJobTemplatesRequest", string(data)}, " ")
}
