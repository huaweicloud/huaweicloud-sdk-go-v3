package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelineQuery 查询流水线列表请求体
type ListPipelineQuery struct {

	// 项目ID列表
	ProjectIds *[]string `json:"project_ids,omitempty"`

	// 组件ID
	ComponentId *string `json:"component_id,omitempty"`

	// 流水线名称
	Name *string `json:"name,omitempty"`

	// 状态
	Status *[]string `json:"status,omitempty"`

	// 是否为变更流水线
	IsPublish *bool `json:"is_publish,omitempty"`

	// 创建人ID列表
	CreatorIds *[]string `json:"creator_ids,omitempty"`

	// 执行人ID列表
	ExecutorIds *[]string `json:"executor_ids,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 起始偏移
	Offset *int64 `json:"offset,omitempty"`

	// 查询数量
	Limit *int64 `json:"limit,omitempty"`

	// 排序字段名称
	SortKey *string `json:"sort_key,omitempty"`

	// 排序规则
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListPipelineQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineQuery struct{}"
	}

	return strings.Join([]string{"ListPipelineQuery", string(data)}, " ")
}
