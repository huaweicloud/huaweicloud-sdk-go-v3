package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDevServerFlavorsResponse Response Object
type ListDevServerFlavorsResponse struct {

	// **参数解释**：当前页数。 **约束限制**：不涉及。 **默认取值**：不涉及。
	Current *int32 `json:"current,omitempty"`

	// **参数解释**：Lite Server规格列表。
	Data *[]ServerFlavorinstanceResponse `json:"data,omitempty"`

	// **参数解释**：总的页数。 **约束限制**：不涉及。 **默认取值**：不涉及。
	Pages *int32 `json:"pages,omitempty"`

	// **参数解释**：每一页的数量。 **约束限制**：不涉及。 **默认取值**：不涉及。
	Size *int32 `json:"size,omitempty"`

	// **参数解释**：总的记录数量。 **约束限制**：不涉及。 **默认取值**：不涉及。
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDevServerFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevServerFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListDevServerFlavorsResponse", string(data)}, " ")
}
