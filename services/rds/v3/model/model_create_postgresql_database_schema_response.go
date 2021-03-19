package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreatePostgresqlDatabaseSchemaResponse struct {
	// 操作结果。

	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePostgresqlDatabaseSchemaResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePostgresqlDatabaseSchemaResponse struct{}"
	}

	return strings.Join([]string{"CreatePostgresqlDatabaseSchemaResponse", string(data)}, " ")
}
