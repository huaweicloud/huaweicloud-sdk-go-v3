package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ColumnInfo 表字段详情
type ColumnInfo struct {

	// 字段名称
	FieldName *string `json:"field_name,omitempty"`

	// 字段类型
	FieldType *string `json:"field_type,omitempty"`

	// 字段长度
	FieldLength *string `json:"field_length,omitempty"`

	// 是否允许为空
	NullAble *string `json:"null_able,omitempty"`

	// 是否是分区字段
	IsPartition *bool `json:"is_partition,omitempty"`

	// 是否是主键字段
	Primary *string `json:"primary,omitempty"`

	// 是否是唯一键字段
	Unique *string `json:"unique,omitempty"`

	// 小数部分位数，非数字类型返回null
	DecimalDigits *string `json:"decimal_digits,omitempty"`
}

func (o ColumnInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColumnInfo struct{}"
	}

	return strings.Join([]string{"ColumnInfo", string(data)}, " ")
}
