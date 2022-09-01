package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiTaskColumnInfo struct {

	// 字段名
	FieldName *string `json:"field_name,omitempty" xml:"field_name"`

	// 字段类型
	FieldType *string `json:"field_type,omitempty" xml:"field_type"`

	// 字段长度
	FieldLength *string `json:"field_length,omitempty" xml:"field_length"`
}

func (o MultiTaskColumnInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiTaskColumnInfo struct{}"
	}

	return strings.Join([]string{"MultiTaskColumnInfo", string(data)}, " ")
}
