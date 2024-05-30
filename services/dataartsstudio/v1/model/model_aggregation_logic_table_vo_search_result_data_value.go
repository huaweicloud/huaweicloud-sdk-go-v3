package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AggregationLogicTableVoSearchResultDataValue 返回的数据信息。
type AggregationLogicTableVoSearchResultDataValue struct {

	// 查询到的汇总表值对象（AggregationLogicTableVO）数组。
	Records *[]AggregationLogicTableVo `json:"records,omitempty"`

	// 符合搜索条件的记录总数。
	Total *int32 `json:"total,omitempty"`
}

func (o AggregationLogicTableVoSearchResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregationLogicTableVoSearchResultDataValue struct{}"
	}

	return strings.Join([]string{"AggregationLogicTableVoSearchResultDataValue", string(data)}, " ")
}
