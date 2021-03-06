package model

import (
	"encoding/json"

	"strings"
)

// 属性值
type AdditionalProperties struct {
	// 类型

	Type string `json:"type"`
}

func (o AdditionalProperties) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AdditionalProperties struct{}"
	}

	return strings.Join([]string{"AdditionalProperties", string(data)}, " ")
}
