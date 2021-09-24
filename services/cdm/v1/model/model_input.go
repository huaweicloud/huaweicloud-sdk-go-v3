package model

import (
	"encoding/json"

	"strings"
)

type Input struct {
	// 参数名

	Name *string `json:"name,omitempty"`
	// 参数值

	Values *string `json:"values,omitempty"`
	// 值类型

	Type *string `json:"type,omitempty"`
}

func (o Input) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Input struct{}"
	}

	return strings.Join([]string{"Input", string(data)}, " ")
}
