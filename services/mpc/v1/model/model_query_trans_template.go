package model

import (
	"encoding/json"

	"strings"
)

type QueryTransTemplate struct {
	// 转码模板名称。

	TemplateName string `json:"template_name"`
	// 租户ID。

	TenantId *string `json:"tenant_id,omitempty"`

	Video *Video `json:"video"`

	Audio *Audio `json:"audio,omitempty"`

	Common *Common `json:"common,omitempty"`
}

func (o QueryTransTemplate) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryTransTemplate struct{}"
	}

	return strings.Join([]string{"QueryTransTemplate", string(data)}, " ")
}
