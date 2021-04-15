package model

import (
	"encoding/json"

	"strings"
)

type ImageTaggingI18nTag struct {
	// 中文标签

	Zh *string `json:"zh,omitempty"`
	// 英文标签

	En *string `json:"en,omitempty"`
}

func (o ImageTaggingI18nTag) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImageTaggingI18nTag struct{}"
	}

	return strings.Join([]string{"ImageTaggingI18nTag", string(data)}, " ")
}
