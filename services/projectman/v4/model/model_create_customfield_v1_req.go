package model

import (
	"encoding/json"

	"strings"
)

type CreateCustomfieldV1Req struct {
	// 字段名称
	Name string `json:"name"`
	// 自定义字段类型 可选类型  textArea|select|radio|text|checkbox|date|time_date|number
	Type string `json:"type"`
	// 字段选项
	Options string `json:"options"`
	// 描述
	Memo string `json:"memo"`
	// 工作项类型
	ScrumType string `json:"scrum_type"`
}

func (o CreateCustomfieldV1Req) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCustomfieldV1Req struct{}"
	}

	return strings.Join([]string{"CreateCustomfieldV1Req", string(data)}, " ")
}
