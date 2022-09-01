package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEventTargetRequest struct {

	// 偏移量，表示从此偏移量开始查询，偏移量不能小于0
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量，不能小于0。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 指定查询排序
	Sort *string `json:"sort,omitempty" xml:"sort"`

	// 指定查询的事件目标标签，模糊匹配
	FuzzyLabel *string `json:"fuzzy_label,omitempty" xml:"fuzzy_label"`
}

func (o ListEventTargetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventTargetRequest struct{}"
	}

	return strings.Join([]string{"ListEventTargetRequest", string(data)}, " ")
}
