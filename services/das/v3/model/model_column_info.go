package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ColumnInfo 列信息
type ColumnInfo struct {

	// 列名称
	ColumnName *string `json:"column_name,omitempty"`

	// 数据类型
	DataType *string `json:"data_type,omitempty"`

	// 字符集名称
	CharacterSetName *string `json:"character_set_name,omitempty"`

	// 是否为主键
	PrimaryKey *bool `json:"primary_key,omitempty"`
}

func (o ColumnInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColumnInfo struct{}"
	}

	return strings.Join([]string{"ColumnInfo", string(data)}, " ")
}
