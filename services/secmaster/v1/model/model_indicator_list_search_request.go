package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IndicatorListSearchRequest 情报列表查询请求体
type IndicatorListSearchRequest struct {

	// 威胁情报ID列表
	Ids *[]string `json:"ids,omitempty"`

	// 数据类ID
	DataclassId *string `json:"dataclass_id,omitempty"`

	Condition *DataobjectSearchCondition `json:"condition"`

	// 分页查询参数。用于指定查询结果的起始位置，从0开始。
	Offset int32 `json:"offset"`

	// 分页查询参数，用于指定一次查询最多的结果数，从1开始。
	Limit int32 `json:"limit"`

	// 排序字段
	SortBy *string `json:"sort_by,omitempty"`

	// 查询起始时间，例如：2024-01-20T00:00:00.000Z+0800
	FromDate *string `json:"from_date,omitempty"`

	// 查询截止时间，例如：2024-01-26T23:59:59.999Z+0800
	ToDate *string `json:"to_date,omitempty"`
}

func (o IndicatorListSearchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndicatorListSearchRequest struct{}"
	}

	return strings.Join([]string{"IndicatorListSearchRequest", string(data)}, " ")
}
