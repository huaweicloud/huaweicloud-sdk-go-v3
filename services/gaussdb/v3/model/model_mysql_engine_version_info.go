package model

import (
	"encoding/json"

	"strings"
)

type MysqlEngineVersionInfo struct {
	// 数据库版本ID，该字段不会有重复

	Id string `json:"id"`
	// 数据库版本号，只返回两位数的大版本号

	Name string `json:"name"`
}

func (o MysqlEngineVersionInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MysqlEngineVersionInfo struct{}"
	}

	return strings.Join([]string{"MysqlEngineVersionInfo", string(data)}, " ")
}
