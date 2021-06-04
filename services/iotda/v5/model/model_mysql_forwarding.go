package model

import (
	"encoding/json"

	"strings"
)

// MySql配置信息
type MysqlForwarding struct {
	Address *NetAddress `json:"address"`
	// **参数说明**：连接MYSQL数据库的库名。 **取值范围**：长度不超过64，只允许字母、数字、下划线（_）、连接符（-）的组合。

	DbName string `json:"db_name"`
	// **参数说明**：连接MYSQL数据库的用户名

	Username string `json:"username"`
	// **参数说明**：连接MYSQL数据库的密码

	Password string `json:"password"`
	// **参数说明**：MYSQL数据库的表名

	TableName string `json:"table_name"`
	// **参数说明**：MYSQL数据库的列和流转数据的对应关系列表。

	ColumnMappings []ColumnMapping `json:"column_mappings"`
}

func (o MysqlForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MysqlForwarding struct{}"
	}

	return strings.Join([]string{"MysqlForwarding", string(data)}, " ")
}
