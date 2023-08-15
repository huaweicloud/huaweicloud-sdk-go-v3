package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagFieldsInfo 调用成功时的返回值。
type TagFieldsInfo struct {

	// 字段名称
	FieldName *string `json:"fieldName,omitempty"`

	// 字段类型
	Type *string `json:"type,omitempty"`

	// 内容
	Content *string `json:"content,omitempty"`

	// 是否解析
	IsAnalysis *bool `json:"isAnalysis,omitempty"`

	// 字段名称
	Index *int32 `json:"index,omitempty"`
}

func (o TagFieldsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagFieldsInfo struct{}"
	}

	return strings.Join([]string{"TagFieldsInfo", string(data)}, " ")
}
