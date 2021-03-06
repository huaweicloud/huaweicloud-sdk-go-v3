package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListTemplatesResponse struct {
	// 模板列表。

	Templates      *[]TemplateView `json:"templates,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListTemplatesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListTemplatesResponse", string(data)}, " ")
}
