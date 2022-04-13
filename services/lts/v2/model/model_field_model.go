package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建结构化模板中的字段模型
type FieldModel struct {
	// 字段名称

	FieldName string `json:"field_name"`
	// 是否开启快速分析。

	IsAnalysis bool `json:"is_analysis"`
}

func (o FieldModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FieldModel struct{}"
	}

	return strings.Join([]string{"FieldModel", string(data)}, " ")
}
