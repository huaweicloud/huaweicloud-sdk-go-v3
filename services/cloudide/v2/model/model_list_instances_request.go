package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListInstancesRequest struct {
	// 每页显示的条目数量 10/15/30

	Limit *int64 `json:"limit,omitempty"`
	// 偏移量，表示从此偏移量开始查询

	Offset *int64 `json:"offset,omitempty"`
	// 是否页面显示（以标签配置为准）

	IsTemporary *bool `json:"is_temporary,omitempty"`
	// 标签

	Label *string `json:"label,omitempty"`
	// 关键字查询(根据实例名，描述模糊查询)

	Search *string `json:"search,omitempty"`
	// 排序方式 asc/desc

	SortDir *string `json:"sort_dir,omitempty"`
	// 排序字段。 display_name 实例名、status状态、pvc_quantity 存储容量、created_time 创建时间、stack_id 技术栈

	SortKey *string `json:"sort_key,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}
