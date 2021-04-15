package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Response Object
type ImageTaggingBody struct {
	// 标签列表集合。

	Tags *[]ImageTaggingItemBody `json:"tags,omitempty"`
}

func (o ImageTaggingBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImageTaggingBody struct{}"
	}

	return strings.Join([]string{"ImageTaggingBody", string(data)}, " ")
}
