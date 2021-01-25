package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateTemplateGroupRequest struct {
	Body *ModifyTransTemplateGroup `json:"body,omitempty"`
}

func (o UpdateTemplateGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTemplateGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateTemplateGroupRequest", string(data)}, " ")
}
