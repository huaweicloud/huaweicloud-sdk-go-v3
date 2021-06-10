package model

import (
	"encoding/json"

	"strings"
)

type ResourceBody struct {
	// 源站域名或源站IP,IP仅支持IPv4,多个源站IP之间以英文逗号间隔,最多支持10个源站IP。

	Sources []SourceWithPort `json:"sources"`
}

func (o ResourceBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResourceBody struct{}"
	}

	return strings.Join([]string{"ResourceBody", string(data)}, " ")
}
