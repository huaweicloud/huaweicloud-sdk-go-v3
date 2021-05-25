package model

import (
	"encoding/json"

	"strings"
)

// 链接信息。
type Link struct {
	// 链接地址。

	Href *string `json:"href,omitempty"`
	// 取值为“self”，表示href为本地链接。

	Rel *string `json:"rel,omitempty"`
}

func (o Link) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Link struct{}"
	}

	return strings.Join([]string{"Link", string(data)}, " ")
}
