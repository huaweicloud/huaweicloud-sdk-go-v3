package model

import (
	"encoding/json"

	"strings"
)

// 目标实例根节点库信息体
type DefaultRootDb struct {
	// 库名。

	DbName *string `json:"db_name,omitempty"`
	// 编码格式。

	DbEncoding *string `json:"db_encoding,omitempty"`
}

func (o DefaultRootDb) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DefaultRootDb struct{}"
	}

	return strings.Join([]string{"DefaultRootDb", string(data)}, " ")
}
