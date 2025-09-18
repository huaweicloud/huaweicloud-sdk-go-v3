package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelineRunsResponse Response Object
type ListPipelineRunsResponse struct {

	// **参数解释**： 起始偏移。 **取值范围**： 不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 查询大小。 **取值范围**： 不涉及。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 执行记录总数。 **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 流水线运行信息列表。 **取值范围**： 不涉及。
	PipelineRuns   *[]ListPipelineRunsPagePipelineRuns `json:"pipeline_runs,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListPipelineRunsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineRunsResponse struct{}"
	}

	return strings.Join([]string{"ListPipelineRunsResponse", string(data)}, " ")
}
