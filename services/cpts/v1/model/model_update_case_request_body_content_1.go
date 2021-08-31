package model

import (
	"encoding/json"

	"strings"
)

type UpdateCaseRequestBodyContent1 struct {
	// content_type

	ContentType *int32 `json:"content_type,omitempty"`

	Content *UpdateCaseRequestBodyContent `json:"content,omitempty"`
}

func (o UpdateCaseRequestBodyContent1) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCaseRequestBodyContent1 struct{}"
	}

	return strings.Join([]string{"UpdateCaseRequestBodyContent1", string(data)}, " ")
}
