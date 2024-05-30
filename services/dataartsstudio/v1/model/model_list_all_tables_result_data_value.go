package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllTablesResultDataValue value，统一的返回结果的外层数据结构。
type ListAllTablesResultDataValue struct {

	// 总量。
	Total *int32 `json:"total,omitempty"`

	// 查询到的审批单对象（AllTableVO）数组。
	Records *[]AllTableVo `json:"records,omitempty"`
}

func (o ListAllTablesResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllTablesResultDataValue struct{}"
	}

	return strings.Join([]string{"ListAllTablesResultDataValue", string(data)}, " ")
}
