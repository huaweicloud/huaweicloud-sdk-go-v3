package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DemoFieldAccess struct {

	// 字段名称需和keys中字段保持一致
	FieldName *string `json:"field_name,omitempty"`

	// 字段值
	FieldValue *string `json:"field_value,omitempty"`
}

func (o DemoFieldAccess) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DemoFieldAccess struct{}"
	}

	return strings.Join([]string{"DemoFieldAccess", string(data)}, " ")
}
