package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrainingExperimentsResponse Response Object
type ListTrainingExperimentsResponse struct {

	// **参数解释**：查询到所有训练实验总数。 **取值范围**：不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**：查询到所有符合查询条件的训练实验总数。 **取值范围**：不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**：查询到所有训练实验限制个数。 **取值范围**：不涉及。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：查询作业的页数，最小为0。例如设置为0，则表示从第一页开始查询。  **取值范围**：不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：查询到所有训练实验排序依赖字段。 **取值范围**：不涉及。
	SortBy *string `json:"sort_by,omitempty"`

	// **参数解释**：查询到所有训练实验排序方式。  **取值范围**： - asc：升序 - desc：降序
	Order *string `json:"order,omitempty"`

	// **参数解释**：查询到所有符合查询条件的训练实验详情。
	Items          *[]TrainingExperimentResponse `json:"items,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListTrainingExperimentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrainingExperimentsResponse struct{}"
	}

	return strings.Join([]string{"ListTrainingExperimentsResponse", string(data)}, " ")
}
