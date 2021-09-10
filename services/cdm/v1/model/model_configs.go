package model

import (
	"encoding/json"

	"strings"
)

type Configs struct {
	Inputs []interface{} `json:"inputs"`
	// 配置名称：源端作业的配置名称为“fromJobConfig”。目的端作业的配置名称为“toJobConfig”,连接的配置名称固定为“linkConfig”。

	Name string `json:"name"`
}

func (o Configs) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Configs struct{}"
	}

	return strings.Join([]string{"Configs", string(data)}, " ")
}
