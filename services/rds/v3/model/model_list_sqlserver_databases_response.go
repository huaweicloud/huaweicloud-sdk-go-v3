package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListSqlserverDatabasesResponse struct {
	// 数据库信息。

	Databases *[]SqlserverDatabaseForDetail `json:"databases,omitempty"`
	// 总数。

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSqlserverDatabasesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSqlserverDatabasesResponse struct{}"
	}

	return strings.Join([]string{"ListSqlserverDatabasesResponse", string(data)}, " ")
}
