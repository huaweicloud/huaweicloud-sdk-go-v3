package model

import (
	"encoding/json"

	"strings"
)

type TransTemplate struct {
	// 转码模板名称。

	TemplateName string `json:"template_name"`

	Video *Video `json:"video,omitempty"`

	Audio *Audio `json:"audio,omitempty"`

	Common *Common `json:"common"`
}

func (o TransTemplate) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TransTemplate struct{}"
	}

	return strings.Join([]string{"TransTemplate", string(data)}, " ")
}
