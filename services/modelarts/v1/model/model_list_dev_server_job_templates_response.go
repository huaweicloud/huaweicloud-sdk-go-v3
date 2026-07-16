package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDevServerJobTemplatesResponse Response Object
type ListDevServerJobTemplatesResponse struct {

	// **参数解释**：当前页数。 **取值范围**：不涉及。
	Current *int32 `json:"current,omitempty"`

	// **参数解释**：模板列表。 **取值范围**：不涉及。
	Data *[]DevServerTemplateListResponse `json:"data,omitempty"`

	// **参数解释**：总的页数。 **取值范围**：不涉及。
	Pages *int32 `json:"pages,omitempty"`

	// **参数解释**：每一页的数量。 **取值范围**：不涉及。
	Size *int32 `json:"size,omitempty"`

	// **参数解释**：总的记录数量。 **取值范围**：不涉及。
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDevServerJobTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevServerJobTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListDevServerJobTemplatesResponse", string(data)}, " ")
}
