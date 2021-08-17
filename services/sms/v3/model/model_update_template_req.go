package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateTemplateReq struct {
	Template *TemplateRequest `json:"template,omitempty"`
}

func (o UpdateTemplateReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTemplateReq struct{}"
	}

	return strings.Join([]string{"UpdateTemplateReq", string(data)}, " ")
}
