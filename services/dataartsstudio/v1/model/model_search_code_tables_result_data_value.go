package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCodeTablesResultDataValue value，统一的返回结果的外层数据结构。
type SearchCodeTablesResultDataValue struct {

	// 总量。
	Total *int32 `json:"total,omitempty"`

	// 查询到的码表对象（CodeTableVO）数组。
	Records *[]CodeTableVo `json:"records,omitempty"`
}

func (o SearchCodeTablesResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCodeTablesResultDataValue struct{}"
	}

	return strings.Join([]string{"SearchCodeTablesResultDataValue", string(data)}, " ")
}
