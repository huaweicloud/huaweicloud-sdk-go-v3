package model

import (
	"encoding/json"

	"strings"
)

// 用于资源对象id列表，{“id”：id值}
type ResourceList struct {
	// 资源ID

	Id string `json:"id"`
}

func (o ResourceList) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResourceList struct{}"
	}

	return strings.Join([]string{"ResourceList", string(data)}, " ")
}
