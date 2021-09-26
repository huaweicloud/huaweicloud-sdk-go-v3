package model

import (
	"encoding/json"

	"strings"
)

// 批量启动边缘实例对象
type BatchStart struct {
	// 待启动的边缘实例列表。

	Servers *[]BaseId `json:"servers,omitempty"`
}

func (o BatchStart) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchStart struct{}"
	}

	return strings.Join([]string{"BatchStart", string(data)}, " ")
}