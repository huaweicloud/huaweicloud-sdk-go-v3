package model

import (
	"encoding/json"

	"strings"
)

// 创建数据库schema信息。
type DbSchemaReq struct {
	// 数据库名称。 数据库名称在1到63个字符之间，由字母、数字、或下划线组成，不能包含其他特殊字符，不能以“pg”和数字开头，且不能和RDS for PostgreSQL模板库重名。 RDS for PostgreSQL模板库包括postgres， template0 ，template1。
	DbName  string       `json:"db_name"`
	Schemas *[]SchemaReq `json:"schemas,omitempty"`
}

func (o DbSchemaReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DbSchemaReq struct{}"
	}

	return strings.Join([]string{"DbSchemaReq", string(data)}, " ")
}
