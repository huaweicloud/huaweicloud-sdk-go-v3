package model

import (
	"encoding/json"

	"strings"
)

type ResourceBody struct {
	// 源站域名或源站IP，IP仅支持IPv4，多个源站IP以多个对象传入，多个对象的origin_type都必须为ipaddr，最多支持10个源站IP对象。

	Sources []SourceWithPort `json:"sources"`
}

func (o ResourceBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResourceBody struct{}"
	}

	return strings.Join([]string{"ResourceBody", string(data)}, " ")
}
