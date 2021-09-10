package model

import (
	"encoding/json"

	"strings"
)

type UpdateDatabaseReq struct {
	// 数据库名称。

	Name string `json:"name"`
	// 数据库备注。

	Comment string `json:"comment"`
}

func (o UpdateDatabaseReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDatabaseReq struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseReq", string(data)}, " ")
}
