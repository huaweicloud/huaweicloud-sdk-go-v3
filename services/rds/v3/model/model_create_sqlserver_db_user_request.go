package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateSqlserverDbUserRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *SqlserverUserForCreation `json:"body,omitempty"`
}

func (o CreateSqlserverDbUserRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSqlserverDbUserRequest struct{}"
	}

	return strings.Join([]string{"CreateSqlserverDbUserRequest", string(data)}, " ")
}
