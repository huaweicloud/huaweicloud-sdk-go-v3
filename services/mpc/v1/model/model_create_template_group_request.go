package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateTemplateGroupRequest struct {
	Body *TransTemplateGroup `json:"body,omitempty"`
}

func (o CreateTemplateGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTemplateGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateTemplateGroupRequest", string(data)}, " ")
}
