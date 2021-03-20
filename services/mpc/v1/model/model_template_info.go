package model

import (
	"encoding/json"

	"strings"
)

type TemplateInfo struct {
	// 转码模板ID。

	TemplateId *int32 `json:"template_id,omitempty"`

	Template *QueryTransTemplate `json:"template,omitempty"`

	Error *XCodeError `json:"error,omitempty"`
}

func (o TemplateInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TemplateInfo struct{}"
	}

	return strings.Join([]string{"TemplateInfo", string(data)}, " ")
}
