package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelinesResponse Response Object
type ListPipelinesResponse struct {

	// **参数解释**： 起始偏移。 **取值范围**： 不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 查询大小。 **取值范围**： 不涉及。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 记录总数。 **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 流水线。 **取值范围**： 不涉及。
	Pipelines      *[]ListPipelinesPagePipelines `json:"pipelines,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListPipelinesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelinesResponse struct{}"
	}

	return strings.Join([]string{"ListPipelinesResponse", string(data)}, " ")
}
