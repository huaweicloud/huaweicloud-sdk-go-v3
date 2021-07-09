package model

import (
	"encoding/json"

	"strings"
)

type MysqlUpdateInstanceNameRequest struct {
	// 修改后的实例名称

	Name string `json:"name"`
}

func (o MysqlUpdateInstanceNameRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MysqlUpdateInstanceNameRequest struct{}"
	}

	return strings.Join([]string{"MysqlUpdateInstanceNameRequest", string(data)}, " ")
}
