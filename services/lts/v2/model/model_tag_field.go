package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// tag信息。
type TagField struct {

	// 字段名称
	FieldName string `json:"fieldName" xml:"fieldName"`

	// 字段数据类型，例：string，long，float
	Type string `json:"type" xml:"type"`

	// 内容
	Content *string `json:"content,omitempty" xml:"content"`

	// 是否开启快速分析
	IsAnalysis *bool `json:"isAnalysis,omitempty" xml:"isAnalysis"`
}

func (o TagField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagField struct{}"
	}

	return strings.Join([]string{"TagField", string(data)}, " ")
}
