package model

import (
	"encoding/json"

	"strings"
)

type UpdateTempRequestBodyContent1 struct {
	// content_type

	ContentType *int32 `json:"content_type,omitempty"`

	Content *UpdateTempRequestBodyContent `json:"content,omitempty"`
}

func (o UpdateTempRequestBodyContent1) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTempRequestBodyContent1 struct{}"
	}

	return strings.Join([]string{"UpdateTempRequestBodyContent1", string(data)}, " ")
}
