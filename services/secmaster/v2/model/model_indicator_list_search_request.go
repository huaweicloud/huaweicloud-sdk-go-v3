package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IndicatorListSearchRequest indicator list query request
type IndicatorListSearchRequest struct {

	// 指标ID列表
	Ids *[]string `json:"ids,omitempty"`

	// 指标名称
	Name *string `json:"name,omitempty"`

	// 数据类ID
	DataclassId *string `json:"dataclass_id,omitempty"`

	// 查询条件
	Condition string `json:"condition"`

	// request offset, from 0
	Offset int32 `json:"offset"`

	// request limit size
	Limit int32 `json:"limit"`

	// sort by property, create_time.
	SortBy *string `json:"sort_by,omitempty"`

	// 查询起始时间
	FromDate *string `json:"from_date,omitempty"`

	// 查询截止时间
	ToDate *string `json:"to_date,omitempty"`
}

func (o IndicatorListSearchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndicatorListSearchRequest struct{}"
	}

	return strings.Join([]string{"IndicatorListSearchRequest", string(data)}, " ")
}
