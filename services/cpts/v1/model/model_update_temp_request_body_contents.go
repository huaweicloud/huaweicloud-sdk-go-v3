package model

import (
	"encoding/json"

	"strings"
)

type UpdateTempRequestBodyContents struct {
	// content

	Content *[]UpdateTempRequestBodyContent1 `json:"content,omitempty"`
	// content_id

	ContentId *int32 `json:"content_id,omitempty"`
	// data_type

	DataType *int32 `json:"data_type,omitempty"`
}

func (o UpdateTempRequestBodyContents) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTempRequestBodyContents struct{}"
	}

	return strings.Join([]string{"UpdateTempRequestBodyContents", string(data)}, " ")
}
