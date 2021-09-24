package model

import (
	"encoding/json"

	"strings"
)

// 数据库信息。
type ListDatastore struct {
	// 数据库引擎。

	Type string `json:"type"`
	// 数据库版本。

	Version string `json:"version"`
}

func (o ListDatastore) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListDatastore struct{}"
	}

	return strings.Join([]string{"ListDatastore", string(data)}, " ")
}
