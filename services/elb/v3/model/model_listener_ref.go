package model

import (
	"encoding/json"

	"strings"
)

// 监听器
type ListenerRef struct {
	// 监听器ID。

	Id string `json:"id"`
}

func (o ListenerRef) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListenerRef struct{}"
	}

	return strings.Join([]string{"ListenerRef", string(data)}, " ")
}
