package model

import (
	"encoding/json"

	"strings"
)

// 数据库信息。
type PgListDatabase struct {
	// 数据库名称。
	Name *string `json:"name,omitempty"`
	// 数据库所属用户。
	Owner *string `json:"owner,omitempty"`
	// 数据库使用的字符集，例如UTF8。
	CharacterSet *string `json:"character_set,omitempty"`
	// 数据库排序集，例如en_US.UTF-8等。
	CollateSet *string `json:"collate_set,omitempty"`
	// 数据库大小（单位：字节）。
	Size *int32 `json:"size,omitempty"`
}

func (o PgListDatabase) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PgListDatabase struct{}"
	}

	return strings.Join([]string{"PgListDatabase", string(data)}, " ")
}
