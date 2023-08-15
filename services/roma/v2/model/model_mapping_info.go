package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MappingInfo struct {

	// 源端字段
	SourceColumn *string `json:"source_column,omitempty"`

	// 源端字段类型
	SourceColumnType *string `json:"source_column_type,omitempty"`

	// 源端字段长度
	SourceColumnLength *string `json:"source_column_length,omitempty"`

	// 目标端字段
	TargetColumn *string `json:"target_column,omitempty"`

	// 目标端字段类型
	TargetColumnType *string `json:"target_column_type,omitempty"`

	// 目标端字段长度
	TargetColumnLength *string `json:"target_column_length,omitempty"`
}

func (o MappingInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MappingInfo struct{}"
	}

	return strings.Join([]string{"MappingInfo", string(data)}, " ")
}
