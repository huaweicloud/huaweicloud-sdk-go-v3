package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListPostgresqlDatabasesResponse struct {
	// 数据库信息。

	Databases *[]PgListDatabase `json:"databases,omitempty"`
	// 总数。

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPostgresqlDatabasesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPostgresqlDatabasesResponse struct{}"
	}

	return strings.Join([]string{"ListPostgresqlDatabasesResponse", string(data)}, " ")
}
