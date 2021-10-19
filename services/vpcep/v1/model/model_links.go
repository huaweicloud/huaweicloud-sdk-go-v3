package model

import (
	"encoding/json"

	"strings"
)

// API的url地址
type Links struct {
	// 当前API版本的引用地址。

	Href *string `json:"href,omitempty"`
	// 发送的实体的MIME类型，取值为application/json。

	Type *string `json:"type,omitempty"`
	// 当前API版本和被引用地址的关系。

	Rel *string `json:"rel,omitempty"`
}

func (o Links) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Links struct{}"
	}

	return strings.Join([]string{"Links", string(data)}, " ")
}
