package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListColdTableTasksResponse Response Object
type ListColdTableTasksResponse struct {

	// **参数解释**：  总记录数。  **取值范围**：  大于等于0。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**：  已下发任务列表。
	Tasks          *[]ColdTableTask `json:"tasks,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListColdTableTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListColdTableTasksResponse struct{}"
	}

	return strings.Join([]string{"ListColdTableTasksResponse", string(data)}, " ")
}
