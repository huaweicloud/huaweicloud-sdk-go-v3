package model

import (
	"encoding/json"

	"strings"
)

// 标签结构体。
type TagV5Dto struct {
	// 标签键，在同一资源下标签键唯一。
	TagKey string `json:"tag_key"`
	// 标签值。
	TagValue *string `json:"tag_value,omitempty"`
}

func (o TagV5Dto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TagV5Dto struct{}"
	}

	return strings.Join([]string{"TagV5Dto", string(data)}, " ")
}
