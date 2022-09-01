package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 调用成功时的返回值。
type TagFieldsInfo struct {

	// 字段名称
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName"`

	// 字段类型
	Type *string `json:"type,omitempty" xml:"type"`

	// 内容
	Content *string `json:"content,omitempty" xml:"content"`

	// 是否解析
	IsAnalysis *bool `json:"isAnalysis,omitempty" xml:"isAnalysis"`

	// 字段名称
	Index *int32 `json:"index,omitempty" xml:"index"`
}

func (o TagFieldsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagFieldsInfo struct{}"
	}

	return strings.Join([]string{"TagFieldsInfo", string(data)}, " ")
}
