package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListProjectTemplatesResponse struct {
	// 模板列表

	Templates *[]ProjectTemplates `json:"templates,omitempty"`
	// 状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListProjectTemplatesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProjectTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListProjectTemplatesResponse", string(data)}, " ")
}
