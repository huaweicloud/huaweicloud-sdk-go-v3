package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IcebergStructField struct {

	// 字段id。
	Id int32 `json:"id"`

	// 字段名称。
	Name string `json:"name"`

	// 字段类型, 包含 BOOLEAN, INTEGER, LONG, FLOAT, DOUBLE, DATE, TIME, TIMESTAMP, TIMESTAMP_NANO, STRING, UUID, FIXED, BINARY, DECIMAL, STRUCT, LIST, MAP
	Type string `json:"type"`

	// 如果type字段输入了STRUCT, LIST, MAP以外的类型，此处应该为空。如果type类型输入了STRUCT, LIST, MAP，此处应该为嵌套类型json体的字符串。
	TypeJson *string `json:"type_json,omitempty"`

	// 表示这个字段是必须的还是可选的
	Required bool `json:"required"`

	// comment
	Doc *string `json:"doc,omitempty"`

	// 字段写入默认值, 在每次写入操作时生效，覆盖显式写入的null值。
	WriteDefault *string `json:"write_default,omitempty"`
}

func (o IcebergStructField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IcebergStructField struct{}"
	}

	return strings.Join([]string{"IcebergStructField", string(data)}, " ")
}
