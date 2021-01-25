package model

import (
	"encoding/json"

	"strings"
)

type TemplateExtend struct {
	Audio *OutputPolicy `json:"audio,omitempty"`
}

func (o TemplateExtend) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TemplateExtend struct{}"
	}

	return strings.Join([]string{"TemplateExtend", string(data)}, " ")
}
