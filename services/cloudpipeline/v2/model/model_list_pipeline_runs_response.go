package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPipelineRunsResponse struct {

	// 起始偏移
	Offset *int32 `json:"offset,omitempty"`

	// 查询大小
	Limit *int32 `json:"limit,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 流水线运行信息
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
