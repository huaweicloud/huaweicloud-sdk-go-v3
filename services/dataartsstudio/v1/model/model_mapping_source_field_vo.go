package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MappingSourceFieldVo struct {

	// 目标字段ID,当前表的某个字段
	TargetFieldId *int64 `json:"target_field_id,omitempty"`

	// 目标字段编码
	TargetFieldName string `json:"target_field_name"`

	// 来源字段id,多个id以逗号分隔
	FieldIds *string `json:"field_ids,omitempty"`

	// 转换表达式
	TransformExpression *string `json:"transform_expression,omitempty"`

	// 来源字段名称列表
	FieldNames *[]string `json:"field_names,omitempty"`

	// 字段是否发生变化
	Changed *bool `json:"changed,omitempty"`
}

func (o MappingSourceFieldVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MappingSourceFieldVo struct{}"
	}

	return strings.Join([]string{"MappingSourceFieldVo", string(data)}, " ")
}
