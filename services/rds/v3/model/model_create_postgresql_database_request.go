package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePostgresqlDatabaseRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *PostgresqlDatabaseForCreation `json:"body,omitempty"`
}

func (o CreatePostgresqlDatabaseRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePostgresqlDatabaseRequest struct{}"
	}

	return strings.Join([]string{"CreatePostgresqlDatabaseRequest", string(data)}, " ")
}
