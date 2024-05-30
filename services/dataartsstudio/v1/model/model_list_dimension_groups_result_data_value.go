package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDimensionGroupsResultDataValue value，统一的返回结果的外层数据结构。
type ListDimensionGroupsResultDataValue struct {
	MainTable *FactLogicTableVo `json:"main_table,omitempty"`

	// 维度。
	DimensionTables *[]DimensionLogicTableVo `json:"dimension_tables,omitempty"`
}

func (o ListDimensionGroupsResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDimensionGroupsResultDataValue struct{}"
	}

	return strings.Join([]string{"ListDimensionGroupsResultDataValue", string(data)}, " ")
}
