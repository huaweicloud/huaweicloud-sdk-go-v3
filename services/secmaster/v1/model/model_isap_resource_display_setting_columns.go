package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapResourceDisplaySettingColumns 列信息
type IsapResourceDisplaySettingColumns struct {

	// 列别名
	ColumnAlias *string `json:"column_alias,omitempty"`

	// 列名称
	ColumnName *string `json:"column_name,omitempty"`

	// 是否默认显示
	DisplayByDefault *bool `json:"display_by_default,omitempty"`
}

func (o IsapResourceDisplaySettingColumns) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapResourceDisplaySettingColumns struct{}"
	}

	return strings.Join([]string{"IsapResourceDisplaySettingColumns", string(data)}, " ")
}
