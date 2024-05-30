package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactLogicTablesResultDataValue value，统一的返回结果的外层数据结构。
type ListFactLogicTablesResultDataValue struct {

	// 总量。
	Total *int32 `json:"total,omitempty"`

	// FactLogicTableVO信息。
	Records *[]FactLogicTableVo `json:"records,omitempty"`
}

func (o ListFactLogicTablesResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactLogicTablesResultDataValue struct{}"
	}

	return strings.Join([]string{"ListFactLogicTablesResultDataValue", string(data)}, " ")
}
