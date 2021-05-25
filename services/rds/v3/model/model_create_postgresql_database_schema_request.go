package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePostgresqlDatabaseSchemaRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *PostgresqlDatabaseSchemaReq `json:"body,omitempty"`
}

func (o CreatePostgresqlDatabaseSchemaRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePostgresqlDatabaseSchemaRequest struct{}"
	}

	return strings.Join([]string{"CreatePostgresqlDatabaseSchemaRequest", string(data)}, " ")
}
