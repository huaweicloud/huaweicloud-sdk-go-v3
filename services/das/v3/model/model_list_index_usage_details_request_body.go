package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIndexUsageDetailsRequestBody 获取索引使用详情列表请求体
type ListIndexUsageDetailsRequestBody struct {

	// 过滤条件
	Conditions []IndexUsageCondition `json:"conditions"`

	// 表名称
	ObjectName *string `json:"object_name,omitempty"`

	// 排序字段
	SortField *string `json:"sort_field,omitempty"`

	// 排序是否升序
	SortAsc *bool `json:"sort_asc,omitempty"`

	// 当前页
	CurPage *int32 `json:"cur_page,omitempty"`

	// 页大小
	PerPage *int32 `json:"per_page,omitempty"`
}

func (o ListIndexUsageDetailsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIndexUsageDetailsRequestBody struct{}"
	}

	return strings.Join([]string{"ListIndexUsageDetailsRequestBody", string(data)}, " ")
}
