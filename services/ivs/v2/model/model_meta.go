package model

import (
	"encoding/json"

	"strings"
)

type Meta struct {
	// 唯一标识此次请求的ID，用户自定义，不超过64位。

	Uuid *string `json:"uuid,omitempty"`
}

func (o Meta) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Meta struct{}"
	}

	return strings.Join([]string{"Meta", string(data)}, " ")
}
