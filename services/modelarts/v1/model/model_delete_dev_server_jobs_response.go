package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDevServerJobsResponse Response Object
type DeleteDevServerJobsResponse struct {

	// **参数解释**：当前页数。 **取值范围**：不涉及。
	Current *int32 `json:"current,omitempty"`

	// **参数解释**：job实例列表。
	Data *[]DevServerJobListResponse `json:"data,omitempty"`

	// **参数解释**：总页数。 **取值范围**：不涉及。
	Pages *int32 `json:"pages,omitempty"`

	// **参数解释**：每一页的数量。 **取值范围**：不涉及。
	Size *int32 `json:"size,omitempty"`

	// **参数解释**：总的记录数量。 **取值范围**：不涉及。
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteDevServerJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDevServerJobsResponse struct{}"
	}

	return strings.Join([]string{"DeleteDevServerJobsResponse", string(data)}, " ")
}
