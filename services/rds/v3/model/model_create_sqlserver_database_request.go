package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateSqlserverDatabaseRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *SqlserverDatabaseForCreation `json:"body,omitempty"`
}

func (o CreateSqlserverDatabaseRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSqlserverDatabaseRequest struct{}"
	}

	return strings.Join([]string{"CreateSqlserverDatabaseRequest", string(data)}, " ")
}
