package model

import (
	"encoding/json"

	"strings"
)

// 数据库信息。
type ListInstancesDatastoreResult struct {
	// 数据库引擎。

	Type string `json:"type"`
	// 数据库版本号。

	Version string `json:"version"`
}

func (o ListInstancesDatastoreResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListInstancesDatastoreResult struct{}"
	}

	return strings.Join([]string{"ListInstancesDatastoreResult", string(data)}, " ")
}
