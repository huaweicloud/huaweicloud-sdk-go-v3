package model

import (
	"encoding/json"

	"strings"
)

type ComponentList struct {
	// 组件名称

	ComponentName string `json:"component_name"`
}

func (o ComponentList) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ComponentList struct{}"
	}

	return strings.Join([]string{"ComponentList", string(data)}, " ")
}
