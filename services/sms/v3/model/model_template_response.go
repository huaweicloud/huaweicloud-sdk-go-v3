package model

import (
	"encoding/json"

	"strings"
)

// 模板响应
type TemplateResponse struct {
	Template *TemplateResponseBody `json:"template,omitempty"`
}

func (o TemplateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TemplateResponse struct{}"
	}

	return strings.Join([]string{"TemplateResponse", string(data)}, " ")
}
