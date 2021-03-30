package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateListenerTagsRequestBody struct {
	Tag *ResourceTag `json:"tag,omitempty"`
}

func (o CreateListenerTagsRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateListenerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateListenerTagsRequestBody", string(data)}, " ")
}
