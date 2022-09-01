package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MappingInfo struct {

	// 源端字段
	SourceColumn *string `json:"source_column,omitempty" xml:"source_column"`

	// 源端字段类型
	SourceColumnType *string `json:"source_column_type,omitempty" xml:"source_column_type"`

	// 源端字段长度
	SourceColumnLength *string `json:"source_column_length,omitempty" xml:"source_column_length"`

	// 目标端字段
	TargetColumn *string `json:"target_column,omitempty" xml:"target_column"`

	// 目标端字段类型
	TargetColumnType *string `json:"target_column_type,omitempty" xml:"target_column_type"`

	// 目标端字段长度
	TargetColumnLength *string `json:"target_column_length,omitempty" xml:"target_column_length"`
}

func (o MappingInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MappingInfo struct{}"
	}

	return strings.Join([]string{"MappingInfo", string(data)}, " ")
}
