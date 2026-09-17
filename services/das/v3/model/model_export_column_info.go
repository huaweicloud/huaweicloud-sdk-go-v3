package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportColumnInfo 导出列信息
type ExportColumnInfo struct {

	// 列名称
	ColumnName *string `json:"column_name,omitempty"`

	// 列值
	ColumnValue *string `json:"column_value,omitempty"`
}

func (o ExportColumnInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportColumnInfo struct{}"
	}

	return strings.Join([]string{"ExportColumnInfo", string(data)}, " ")
}
