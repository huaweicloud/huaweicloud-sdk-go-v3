package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkItemFlowFieldConfigVo 工作项字段配置
type WorkItemFlowFieldConfigVo struct {

	// 字段编码
	FieldCode *string `json:"field_code,omitempty"`

	// 字段值类型
	ValueType *string `json:"value_type,omitempty"`

	// 字段操作类型
	FieldOperation *string `json:"field_operation,omitempty"`

	FieldValue *WorkItemFlowFieldValueVo `json:"field_value,omitempty"`

	// 是否必填
	Required *bool `json:"required,omitempty"`

	FieldRange *WorkItemFlowFieldRangeVo `json:"field_range,omitempty"`
}

func (o WorkItemFlowFieldConfigVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkItemFlowFieldConfigVo struct{}"
	}

	return strings.Join([]string{"WorkItemFlowFieldConfigVo", string(data)}, " ")
}
