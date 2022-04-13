package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StructFieldInfoReturn struct {
	// 字段名称

	FieldName *string `json:"fieldName,omitempty"`
	// 字段数据类型

	Type *string `json:"type,omitempty"`
	// 字段内容

	Content *string `json:"content,omitempty"`
	// 结构化方式

	IsAnalysis *bool `json:"isAnalysis,omitempty"`
	// 序号

	Index *int32 `json:"index,omitempty"`
}

func (o StructFieldInfoReturn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructFieldInfoReturn struct{}"
	}

	return strings.Join([]string{"StructFieldInfoReturn", string(data)}, " ")
}
