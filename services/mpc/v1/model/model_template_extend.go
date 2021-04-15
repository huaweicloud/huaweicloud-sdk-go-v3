package model

import (
	"encoding/json"

	"strings"
)

type TemplateExtend struct {
	Audio *AudioExtendSettings `json:"audio,omitempty"`

	Video *VideoExtendSettings `json:"video,omitempty"`
}

func (o TemplateExtend) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TemplateExtend struct{}"
	}

	return strings.Join([]string{"TemplateExtend", string(data)}, " ")
}
