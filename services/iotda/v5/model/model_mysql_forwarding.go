package model

import (
	"encoding/json"

	"strings"
)

// MySql配置信息
type MysqlForwarding struct {
	Address *NetAddress `json:"address"`
	// 连接MYSQL数据库的库名
	DbName string `json:"db_name"`
	// 连接MYSQL数据库的用户名
	Username string `json:"username"`
	// 连接MYSQL数据库的密码
	Password string `json:"password"`
	// MYSQL数据库的表名
	TableName string `json:"table_name"`
	// MYSQL数据库的列和流转数据的对应关系列表。
	ColumnMappings []ColumnMapping `json:"column_mappings"`
}

func (o MysqlForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MysqlForwarding struct{}"
	}

	return strings.Join([]string{"MysqlForwarding", string(data)}, " ")
}
