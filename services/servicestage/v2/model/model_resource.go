package model

import (
	"encoding/json"

	"strings"
)

type Resource struct {
	// 资源ID

	Id string `json:"id"`

	Type *ResourceType `json:"type"`
}

func (o Resource) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
