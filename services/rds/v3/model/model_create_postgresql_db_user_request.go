package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePostgresqlDbUserRequest struct {
	XLanguage  *string                    `json:"X-Language,omitempty"`
	InstanceId string                     `json:"instance_id"`
	Body       *PostgresqlUserForCreation `json:"body,omitempty"`
}

func (o CreatePostgresqlDbUserRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePostgresqlDbUserRequest struct{}"
	}

	return strings.Join([]string{"CreatePostgresqlDbUserRequest", string(data)}, " ")
}
