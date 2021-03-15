package model

import (
	"encoding/json"

	"strings"
)

// 数据库信息。
type DatabaseForListSchema struct {
	// schema名称。
	SchemaName string `json:"schema_name"`
	// schema所属用户。
	Owner string `json:"owner"`
}

func (o DatabaseForListSchema) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DatabaseForListSchema struct{}"
	}

	return strings.Join([]string{"DatabaseForListSchema", string(data)}, " ")
}
