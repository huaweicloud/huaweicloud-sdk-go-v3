package model

import (
	"encoding/json"

	"strings"
)

// 数据库信息。
type MysqlDatastore struct {
	// 数据库引擎，现在只支持gaussdb-mysql

	Type string `json:"type"`
	// 数据库版本。  数据库支持的详细版本信息，可调用查询数据库引擎的版本接口获取。

	Version string `json:"version"`
}

func (o MysqlDatastore) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MysqlDatastore struct{}"
	}

	return strings.Join([]string{"MysqlDatastore", string(data)}, " ")
}
