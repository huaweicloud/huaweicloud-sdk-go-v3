package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelineRunsQuery 查询流水线运行历史请求体
type ListPipelineRunsQuery struct {

	// 状态
	Status *[]string `json:"status,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 起始偏移
	Offset *int64 `json:"offset,omitempty"`

	// 查询大小
	Limit *int64 `json:"limit,omitempty"`

	// 排序字段名称
	SortKey *string `json:"sort_key,omitempty"`

	// 排序规则
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListPipelineRunsQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineRunsQuery struct{}"
	}

	return strings.Join([]string{"ListPipelineRunsQuery", string(data)}, " ")
}
