package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapTableColumnDisplaySettingDto 表格列展示设置数据传输对象
type IsapTableColumnDisplaySettingDto struct {

	// 表格列名称
	ColumnName *string `json:"column_name,omitempty"`

	// 表格列别名
	ColumnAlias *string `json:"column_alias,omitempty"`

	// 表格列别名
	ColumnAliasEn *string `json:"column_alias_en,omitempty"`

	// 是否默认展示
	DisplayByDefault *bool `json:"display_by_default,omitempty"`
}

func (o IsapTableColumnDisplaySettingDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapTableColumnDisplaySettingDto struct{}"
	}

	return strings.Join([]string{"IsapTableColumnDisplaySettingDto", string(data)}, " ")
}
