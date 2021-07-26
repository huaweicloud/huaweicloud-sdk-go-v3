package model

import (
	"encoding/json"

	"strings"
)

// 修改路由body
type RouteBody struct {
	// 名称

	Name *string `json:"name,omitempty"`
	// 路由信息

	Servers *[]RouteServerBody `json:"servers,omitempty"`
}

func (o RouteBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RouteBody struct{}"
	}

	return strings.Join([]string{"RouteBody", string(data)}, " ")
}
