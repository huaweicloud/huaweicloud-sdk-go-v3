package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePostgresqlDatabaseSchemaRequest struct {
	XLanguage  *string      `json:"X-Language,omitempty"`
	InstanceId string       `json:"instance_id"`
	Body       *DbSchemaReq `json:"body,omitempty"`
}

func (o CreatePostgresqlDatabaseSchemaRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePostgresqlDatabaseSchemaRequest struct{}"
	}

	return strings.Join([]string{"CreatePostgresqlDatabaseSchemaRequest", string(data)}, " ")
}
