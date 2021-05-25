package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListMigrationTaskRequest struct {
	// 偏移量，表示从此偏移量开始查询， offset大于等于0。

	Offset *int32 `json:"offset,omitempty"`
	// 每页显示的条目数量。

	Limit *int32 `json:"limit,omitempty"`
	// 迁移任务名称。

	Name *string `json:"name,omitempty"`
}

func (o ListMigrationTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"ListMigrationTaskRequest", string(data)}, " ")
}
