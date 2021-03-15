package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListPostgresqlDbUserPaginatedResponse struct {
	// 用户信息。
	Users *[]PgUserForList `json:"users,omitempty"`
	// 总条数。
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
