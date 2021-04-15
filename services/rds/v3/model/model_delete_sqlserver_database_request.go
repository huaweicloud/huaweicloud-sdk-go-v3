package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteSqlserverDatabaseRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	DbName string `json:"db_name"`
}

func (o DeleteSqlserverDatabaseRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSqlserverDatabaseRequest struct{}"
	}

	return strings.Join([]string{"DeleteSqlserverDatabaseRequest", string(data)}, " ")
}
