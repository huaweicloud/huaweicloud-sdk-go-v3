package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateSqlserverDatabaseResponse struct {
	// 操作结果。

	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSqlserverDatabaseResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSqlserverDatabaseResponse struct{}"
	}

	return strings.Join([]string{"CreateSqlserverDatabaseResponse", string(data)}, " ")
}
