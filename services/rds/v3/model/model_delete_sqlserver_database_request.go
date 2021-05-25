package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteSqlserverDatabaseRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// 需要删除的数据库名。

	DbName string `json:"db_name"`
}

func (o DeleteSqlserverDatabaseRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSqlserverDatabaseRequest struct{}"
	}

	return strings.Join([]string{"DeleteSqlserverDatabaseRequest", string(data)}, " ")
}
