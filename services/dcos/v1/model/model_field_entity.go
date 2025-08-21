package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FieldEntity 字段实体
type FieldEntity struct {

	// 字段编码
	FieldCode string `json:"field_code"`

	// 字段值
	FieldValue *interface{} `json:"field_value,omitempty"`
}

func (o FieldEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FieldEntity struct{}"
	}

	return strings.Join([]string{"FieldEntity", string(data)}, " ")
}
