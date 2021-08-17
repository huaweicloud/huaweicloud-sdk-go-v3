package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateTemplateReq struct {
	Template *TemplateRequest `json:"template"`
}

func (o CreateTemplateReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTemplateReq struct{}"
	}

	return strings.Join([]string{"CreateTemplateReq", string(data)}, " ")
}
