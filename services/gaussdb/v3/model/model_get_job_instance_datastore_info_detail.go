package model

import (
	"encoding/json"

	"strings"
)

// 数据库信息。
type GetJobInstanceDatastoreInfoDetail struct {
	// 数据库引擎。

	Type string `json:"type"`
	// 数据库版本。

	Version string `json:"version"`
}

func (o GetJobInstanceDatastoreInfoDetail) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GetJobInstanceDatastoreInfoDetail struct{}"
	}

	return strings.Join([]string{"GetJobInstanceDatastoreInfoDetail", string(data)}, " ")
}
