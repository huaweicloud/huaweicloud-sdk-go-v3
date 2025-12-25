package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TableColumnDisplaySetting struct {

	// 表格列名称
	ColumnName *string `json:"column_name,omitempty"`

	// 表格列别名
	ColumnAlias *string `json:"column_alias,omitempty"`

	// 是否默认展示
	DisplayByDefault *bool `json:"display_by_default,omitempty"`
}

func (o TableColumnDisplaySetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableColumnDisplaySetting struct{}"
	}

	return strings.Join([]string{"TableColumnDisplaySetting", string(data)}, " ")
}
