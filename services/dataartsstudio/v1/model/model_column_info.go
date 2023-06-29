package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ColumnInfo 概要信息
type ColumnInfo struct {

	// 列名称
	ColumnName *string `json:"column_name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`
}

func (o ColumnInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColumnInfo struct{}"
	}

	return strings.Join([]string{"ColumnInfo", string(data)}, " ")
}
