package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListPostgresqlDbUserPaginatedResponse struct {
	// 列表中每个元素表示一个数据库用户。

	Users *[]PostgresqlUserForList `json:"users,omitempty"`
	// 数据库用户总数。

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPostgresqlDbUserPaginatedResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPostgresqlDbUserPaginatedResponse struct{}"
	}

	return strings.Join([]string{"ListPostgresqlDbUserPaginatedResponse", string(data)}, " ")
}
