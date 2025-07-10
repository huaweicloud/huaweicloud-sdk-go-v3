package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobRequest Request Object
type ListJobRequest struct {

	// 分页页码，表示从此页开始查询，page_index大于等于0
	PageIndex *int32 `json:"page_index,omitempty"`

	// 每页显示的条目数量，page_size小于等于100
	PageSize *int32 `json:"page_size,omitempty"`

	// 查询关键字
	Search *string `json:"search,omitempty"`

	// 排序字段
	SortField *string `json:"sort_field,omitempty"`

	// 排序方式（ASC|DESC）
	SortOrder *string `json:"sort_order,omitempty"`

	// 创建人
	CreatorId *string `json:"creator_id,omitempty"`

	// 构建状态
	BuildStatus *string `json:"build_status,omitempty"`
}

func (o ListJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobRequest struct{}"
	}

	return strings.Join([]string{"ListJobRequest", string(data)}, " ")
}
