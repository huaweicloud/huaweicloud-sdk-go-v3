package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListTemplateRequest struct {
	TemplateId *[]int32 `json:"template_id,omitempty"`

	Page *int32 `json:"page,omitempty"`

	Size *int32 `json:"size,omitempty"`
}

func (o ListTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListTemplateRequest", string(data)}, " ")
}
