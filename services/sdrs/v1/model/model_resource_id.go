package model

import (
	"encoding/json"

	"strings"
)

// 保护实例Id
type ResourceId struct {
	// 资源ID

	Id string `json:"id"`
}

func (o ResourceId) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResourceId struct{}"
	}

	return strings.Join([]string{"ResourceId", string(data)}, " ")
}
