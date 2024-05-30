package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDimensionLogicTablesResultDataValue value，统一的返回结果的外层数据结构。
type ListDimensionLogicTablesResultDataValue struct {

	// 总量。
	Total *int32 `json:"total,omitempty"`

	// DimensionVO信息。
	Records *[]DimensionLogicTableVo `json:"records,omitempty"`
}

func (o ListDimensionLogicTablesResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDimensionLogicTablesResultDataValue struct{}"
	}

	return strings.Join([]string{"ListDimensionLogicTablesResultDataValue", string(data)}, " ")
}
