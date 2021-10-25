package model

import (
	"encoding/json"

	"strings"
)

// 保护实例元数据数据结构
type MetadataParams struct {
	// 保护实例元数据中资源冻结的字段。 true：表示资源被冻结。 空：表示资源没有被冻结。

	SystemFrozen string `json:"__system__frozen"`
}

func (o MetadataParams) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MetadataParams struct{}"
	}

	return strings.Join([]string{"MetadataParams", string(data)}, " ")
}
