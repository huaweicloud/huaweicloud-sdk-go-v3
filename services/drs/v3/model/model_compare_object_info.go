package model

import (
	"encoding/json"

	"strings"
)

//
type CompareObjectInfo struct {
	// 库名。

	DbName string `json:"db_name"`
	// 该库下的表名列表。

	TableName *[]string `json:"table_name,omitempty"`
}

func (o CompareObjectInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CompareObjectInfo struct{}"
	}

	return strings.Join([]string{"CompareObjectInfo", string(data)}, " ")
}
