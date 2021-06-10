package model

import (
	"encoding/json"

	"strings"
)

//
type CompareTableInfoWithToken struct {
	// 表名。

	TableName string `json:"table_name"`
	// 该表的min token。

	MinToken *string `json:"min_token,omitempty"`
	// 该表的max token。

	MaxToken *string `json:"max_token,omitempty"`
}

func (o CompareTableInfoWithToken) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CompareTableInfoWithToken struct{}"
	}

	return strings.Join([]string{"CompareTableInfoWithToken", string(data)}, " ")
}
