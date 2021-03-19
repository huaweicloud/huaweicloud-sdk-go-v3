package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListTemplateGroupRequest struct {
	GroupId *[]string `json:"group_id,omitempty"`

	GroupName *[]string `json:"group_name,omitempty"`

	Page *int32 `json:"page,omitempty"`

	Size *int32 `json:"size,omitempty"`
}

func (o ListTemplateGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTemplateGroupRequest struct{}"
	}

	return strings.Join([]string{"ListTemplateGroupRequest", string(data)}, " ")
}
