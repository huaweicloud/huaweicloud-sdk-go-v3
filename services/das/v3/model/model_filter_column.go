package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FilterColumn 筛选条件
type FilterColumn struct {

	// 筛选条件字段名
	ColumnName *string `json:"column_name,omitempty"`

	// 筛选条件值
	ColumnValue *string `json:"column_value,omitempty"`
}

func (o FilterColumn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilterColumn struct{}"
	}

	return strings.Join([]string{"FilterColumn", string(data)}, " ")
}
