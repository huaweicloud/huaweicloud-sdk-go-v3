package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListPostgresqlDatabaseSchemasResponse struct {
	// 列表中每个元素表示一个数据库schema。
	DatabaseSchemas *[]DatabaseForListSchema `json:"database_schemas,omitempty"`
	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPostgresqlDatabaseSchemasResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPostgresqlDatabaseSchemasResponse struct{}"
	}

	return strings.Join([]string{"ListPostgresqlDatabaseSchemasResponse", string(data)}, " ")
}
