package model

import (
	"encoding/json"

	"strings"
)

type BasePathDto struct {
	// 节点日志根目录

	LogBasePath *string `json:"log_base_path,omitempty"`
	// 节点配置根目录

	ConfigBasePath *string `json:"config_base_path,omitempty"`
	// 节点数据存储根目录

	DbBasePath *string `json:"db_base_path,omitempty"`
}

func (o BasePathDto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BasePathDto struct{}"
	}

	return strings.Join([]string{"BasePathDto", string(data)}, " ")
}
