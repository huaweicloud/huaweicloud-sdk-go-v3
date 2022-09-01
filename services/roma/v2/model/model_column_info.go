package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 表字段详情
type ColumnInfo struct {

	// 字段名称
	FieldName *string `json:"field_name,omitempty" xml:"field_name"`

	// 字段类型
	FieldType *string `json:"field_type,omitempty" xml:"field_type"`

	// 字段长度
	FieldLength *string `json:"field_length,omitempty" xml:"field_length"`

	// 是否允许为空
	NullAble *string `json:"null_able,omitempty" xml:"null_able"`

	// 是否是分区字段
	IsPartition *bool `json:"is_partition,omitempty" xml:"is_partition"`

	// 是否是主键字段
	Primary *string `json:"primary,omitempty" xml:"primary"`

	// 是否是唯一键字段
	Unique *string `json:"unique,omitempty" xml:"unique"`

	// 小数部分位数，非数字类型返回null
	DecimalDigits *string `json:"decimal_digits,omitempty" xml:"decimal_digits"`
}

func (o ColumnInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColumnInfo struct{}"
	}

	return strings.Join([]string{"ColumnInfo", string(data)}, " ")
}
