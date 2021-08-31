package model

import (
	"encoding/json"

	"strings"
)

type UpdateCaseRequestBodyContents struct {
	// content

	Content *[]UpdateCaseRequestBodyContent1 `json:"content,omitempty"`
	// content_id

	ContentId *int32 `json:"content_id,omitempty"`
	// data_type

	DataType *int32 `json:"data_type,omitempty"`
	// index

	Index *int32 `json:"index,omitempty"`
}

func (o UpdateCaseRequestBodyContents) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCaseRequestBodyContents struct{}"
	}

	return strings.Join([]string{"UpdateCaseRequestBodyContents", string(data)}, " ")
}
