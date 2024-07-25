package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaselineSearchRequestBody BaselineSearchRequestBody
type BaselineSearchRequestBody struct {

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录
	Offset *int32 `json:"offset,omitempty"`

	// 排序关键字
	SortBy *string `json:"sort_by,omitempty"`

	// 降序或升序, DESC|ESC
	Order *string `json:"order,omitempty"`

	// 起始时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区
	FromDate *string `json:"from_date,omitempty"`

	// 截止时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区
	ToDate *string `json:"to_date,omitempty"`

	// 搜索条件表达式
	Condition *interface{} `json:"condition,omitempty"`
}

func (o BaselineSearchRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaselineSearchRequestBody struct{}"
	}

	return strings.Join([]string{"BaselineSearchRequestBody", string(data)}, " ")
}
